package main

type EngineState interface {
	Start() EngineState
	Stop() EngineState
}

type EngineTurnedOnState struct {
}

func (e EngineTurnedOnState) Start() EngineState {
	println("‼️ Engine already started!")
	return e
}

func (e EngineTurnedOnState) Stop() EngineState {
	println("✅ Engine stopped.")
	return EngineTurnedOffState{}
}

type EngineTurnedOffState struct {
}

func (e EngineTurnedOffState) Start() EngineState {
	println("✅ Engine started.")
	return EngineTurnedOnState{}
}

func (e EngineTurnedOffState) Stop() EngineState {
	println("‼️ Engine already stopped!")
	return e
}

type Car struct {
	engineState EngineState
}

func NewCar() Car {
	return Car{engineState: EngineTurnedOffState{}}
}

func (c *Car) StartEngine() {
	c.engineState = c.engineState.Start()
}

func (c *Car) StopEngine() {
	c.engineState = c.engineState.Stop()
}

func main() {
	c1 := NewCar()
	c1.StopEngine()

	c2 := NewCar()
	c2.StartEngine()
	c2.StopEngine()
}
