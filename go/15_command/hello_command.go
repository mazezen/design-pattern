package command

type HelloCommand struct{}

func (h *HelloCommand) Name() string {
	return "hello"
}

func (h *HelloCommand) Execute(args []string) string {
	if len(args) == 0 {
		return "hello world"
	}
	return "hello " + args[0]
}