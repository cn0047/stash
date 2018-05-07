package main

type Mediator interface {
	Tick(data interface{})
}

type VerboseMediator struct {
}

func (vm VerboseMediator) Tick(data interface{}) {
	v := data.(int)
	// Interact with other objects:
	println("[VerboseMediator] emit into stdout:", v)
	println("[VerboseMediator] emit into log file:", v)
	println("[VerboseMediator] emit into ELK:", v)
	println("[VerboseMediator] emit into slack:", v)
	println("[VerboseMediator] emit into email:", v)
}

type EventEmitter struct {
}

type MediatorCLI struct {
}

func (vm MediatorCLI) Tick(data interface{}) {
	v := data.(int)
	println("[MediatorCLI] emit into stdout: ", v)
}

func (ee EventEmitter) Emit(m Mediator) {
	for i := 0; i < 2; i++ {
		m.Tick(i)
	}
}

func main() {
	ee := EventEmitter{}
	ee.Emit(VerboseMediator{})
	ee.Emit(MediatorCLI{})
}
