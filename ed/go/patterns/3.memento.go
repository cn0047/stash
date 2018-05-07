package main

type Memento struct {
	State string
}

type GamePlay struct {
	State string
}

func (gp GamePlay) SaveToMemento() Memento {
	return Memento{State: gp.State}
}

func (gp *GamePlay) RestoreFromMemento(m Memento) {
	gp.State = m.State
}

func main() {
	states := make([]Memento, 0)
	gp := GamePlay{}

	println("Init game:")
	gp.State = " ➡ Init normal Super Mario size."
	states = append(states, gp.SaveToMemento())
	println(gp.State)

	println("Ate mushroom:")
	gp.State = " ➡ Init big Super Mario size."
	states = append(states, gp.SaveToMemento())
	println(gp.State)

	println("Collided with turtle:")
	gp.RestoreFromMemento(states[len(states)-2]) // restore penultimate state
	println(gp.State)
}
