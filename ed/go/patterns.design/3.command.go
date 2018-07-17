package main

type Lamp struct {
}

type Command interface {
	Execute()
}

type TurnOnCommand struct {
}

func (toc TurnOnCommand) Execute() {
	println("Turn on lamp.")
}

type TurnOffCommand struct {
}

func (toc TurnOffCommand) Execute() {
	println("Turn off lamp.")
}

type CommandRegistry struct {
	registry []Command
}

func (cr *CommandRegistry) Add(c Command) {
	cr.registry = append(cr.registry, c)
}

func (cr CommandRegistry) ExecuteAll() {
	for _, c := range cr.registry {
		c.Execute()
	}
}

func main() {
	r := CommandRegistry{}
	r.Add(TurnOnCommand{})
	r.Add(TurnOffCommand{})
	r.Add(TurnOnCommand{})
	r.ExecuteAll()
}
