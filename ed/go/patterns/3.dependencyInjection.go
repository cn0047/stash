package main

type PoemStorage interface {
	Save(name string)
}

type Notebook struct {
}

func (n Notebook) Save(name string) {
	println("[Notebook] save poem into store:", name)
}

type Poem struct {
	Name    string
	Storage PoemStorage
}

func (p Poem) Save() {
	p.Storage.Save(p.Name)
}

func NewPoem(name string, storage PoemStorage) *Poem {
	return &Poem{Name: name, Storage: storage}
}

func main() {
	n := Notebook{}
	p1 := NewPoem("My first poem", n)
	p1.Save()
	p2 := NewPoem("My second poem", n)
	p2.Save()
}
