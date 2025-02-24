package commands

type cliCommand struct {
	name        string
	description string
	callback    func([]string) error
}

func (c cliCommand) Name() string {
	return c.name
}

func (c cliCommand) Description() string {
	return c.description
}

func (c cliCommand) Callback() func([]string) error {
	return c.callback
}

var Commands = map[string]cliCommand{}

func Register(name string, description string, callback func([]string) error) {
	Commands[name] = cliCommand{
		name:        name,
		description: description,
		callback:    callback,
	}
}
