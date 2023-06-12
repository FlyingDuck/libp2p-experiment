package cmdadd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"

	"github.com/FlyingDuck/libp2p-experiment/linglongzone/cmds"
	"github.com/FlyingDuck/libp2p-experiment/linglongzone/core/env"
	coreiface "github.com/ipfs/boxo/coreiface"
	"github.com/ipfs/boxo/coreiface/options"
	"github.com/ipfs/boxo/files"
	"github.com/multiformats/go-multihash"
	"github.com/spf13/cobra"
)

func NewExecutor(cmd *cobra.Command, args []string) *executor {
	return &executor{
		cmd:  cmd,
		args: args,
	}
}

type executor struct {
	cmd  *cobra.Command
	args []string
}

func (ex *executor) Execute() error {
	ctx := context.Background()

	// 解析要存储的文件

	argDef := cmds.FileArg("path", true, true, "The path to a file to be added to IPFS.") //.EnableRecursive().EnableStdin()
	inputs := ex.args

	fileArgs := make([]files.DirEntry, 0)
	// Each file argument's import directory name is recorded under its base name
	// to reject two files with the same name but different import directories
	// (same directory just means the _exact_ same file, so we can skip it):
	//    file base name -> file directory name
	fileImportDirName := make(map[string]string)

	for _, input := range inputs {
		fpath := input
		fpath = filepath.Clean(fpath)

		var (
			// todo 这里先默认支持有限类型的参数，后续通过命令的 Flag 来控制
			rulesFile   string
			ignoreRules []string
			hidden      bool
			recursive   bool
		)
		filter, err := files.NewFilter(rulesFile, ignoreRules, hidden)
		if err != nil {
			return err
		}

		nf, err := appendFile(fpath, &argDef, recursive, filter)
		if err != nil {
			return err
		}
		curDir := filepath.Dir(fpath)
		fname := filepath.Base(fpath)
		prevDir, ok := fileImportDirName[fname]
		if !ok {
			fileImportDirName[fname] = curDir
		} else {
			if prevDir != curDir {
				return fmt.Errorf("file name %s repeated under different import directories: %s and %s", fpath, curDir, prevDir)
			}
			continue // Skip repeated files.
		}
		fileArgs = append(fileArgs, files.FileEntry(fpath, nf))
	}

	var toadd files.Directory
	if len(fileArgs) > 0 {
		sort.Slice(fileArgs, func(i, j int) bool {
			return fileArgs[i].Name() < fileArgs[j].Name()
		})
		toadd = files.NewSliceDirectory(fileArgs)
	}

	//

	coreAPI, err := env.GetAPI()
	if err != nil {
		return err
	}
	enc, err := env.GetCidEncoder()
	if err != nil {
		return err
	}

	hashFunCode := multihash.SHA2_256
	opts := []options.UnixfsAddOption{
		options.Unixfs.Hash(uint64(hashFunCode)),
	}

	adderOutChanSize := 8

	var added int
	addit := toadd.Entries()
	for addit.Next() {
		_, dir := addit.Node().(files.Directory)
		errCh := make(chan error, 1)
		events := make(chan interface{}, adderOutChanSize)
		opts[len(opts)-1] = options.Unixfs.Events(events)

		go func() {
			var err error
			defer close(events)
			pathAdded, err := coreAPI.Unixfs().Add(ctx, addit.Node(), opts...)
			if err != nil {
				errCh <- err
				return
			}
			fmt.Printf("pathAdded: %s\n", pathAdded.Cid())

			errCh <- err
		}()

		for event := range events {
			output, ok := event.(*coreiface.AddEvent)
			if !ok {
				return errors.New("unknown event type")
			}

			h := ""
			if output.Path != nil {
				h = enc.Encode(output.Path.Cid())
			}

			if !dir && addit.Name() != "" {
				output.Name = addit.Name()
			} else {
				output.Name = path.Join(addit.Name(), output.Name)
			}

			addEvent := &AddEvent{
				Name:  output.Name,
				Hash:  h,
				Bytes: output.Bytes,
				Size:  output.Size,
			}
			fmt.Printf("AddEvent: %+v", addEvent)
		}

		if err := <-errCh; err != nil {
			return err
		}
		added++
	}

	if addit.Err() != nil {
		return addit.Err()
	}

	if added == 0 {
		return fmt.Errorf("expected a file argument")
	}

	return nil
}

var (
	returnDirNotSupportedErr = func(fpath string, arg *cmds.Argument) error {
		return fmt.Errorf("invalid path '%s', argument '%s' does not support directories", fpath, arg.Name)
	}
	returnNotRecursiveErr = func(fpath string) error {
		return fmt.Errorf("'%s' is a directory, use the '-%s' flag to specify directories", fpath, cmds.RecShort)
	}
)

//const dirNotSupportedFmtStr = "invalid path '%s', argument '%s' does not support directories"
//const notRecursiveFmtStr = "'%s' is a directory, use the '-%s' flag to specify directories"

func appendFile(fpath string, argDef *cmds.Argument, recursive bool, filter *files.Filter) (files.Node, error) {
	stat, err := os.Lstat(fpath)
	if err != nil {
		return nil, err
	}

	if stat.IsDir() {
		if !argDef.Recursive {
			return nil, returnDirNotSupportedErr(fpath, argDef)
		}
		if !recursive {
			return nil, returnNotRecursiveErr(fpath)
		}
	} else if (stat.Mode() & os.ModeNamedPipe) != 0 {
		// Special case pipes that are provided directly on the command line
		// We do this here instead of go-ipfs-files, as we need to differentiate between
		// recursive(unsupported) and direct(supported) mode
		file, err := os.Open(fpath)
		if err != nil {
			return nil, err
		}

		return files.NewReaderFile(file), nil
	}
	return files.NewSerialFileWithFilter(fpath, filter, stat)
}

type AddEvent struct {
	Name  string
	Hash  string `json:",omitempty"`
	Bytes int64  `json:",omitempty"`
	Size  string `json:",omitempty"`
}
