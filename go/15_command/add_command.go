package command

import "strconv"

type AddCommand struct{}

func (a *AddCommand) Name() string {
	return "add"
}

func (a *AddCommand) Execute(args []string) string {
	if len(args) < 2 {
		return "need 2 numbers"
	}

	x, err1 := strconv.Atoi(args[0])
	y, err2 := strconv.Atoi(args[1])

	if err1 != nil || err2 != nil {
		return "invalid number"
	}
	
	return strconv.Itoa(x + y)
}