package command

type Command interface {
	Name() string
	Execute(args []string) string
}