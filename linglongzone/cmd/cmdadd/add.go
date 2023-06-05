package cmdadd

func NewExecutor() *executor {
	return &executor{}
}

type executor struct {
}

func (ex *executor) Execute() error {

	return nil
}
