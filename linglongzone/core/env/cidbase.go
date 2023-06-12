package env

import (
	"github.com/ipfs/go-cidutil/cidenc"
)

func GetCidEncoder() (cidenc.Encoder, error) {
	//base, _ := req.Options[OptionCidBase.Name()].(string)
	//upgrade, upgradeDefined := req.Options[OptionUpgradeCidV0InOutput.Name()].(bool)

	e := cidenc.Default()

	//if base != "" {
	//	var err error
	//	e.Base, err = mbase.EncoderByName(base)
	//	if err != nil {
	//		return e, err
	//	}
	//	if autoUpgrade {
	//		e.Upgrade = true
	//	}
	//}
	//
	//if upgradeDefined {
	//	e.Upgrade = upgrade
	//}

	return e, nil
}
