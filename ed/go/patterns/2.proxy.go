package main

import (
	"fmt"
)

type RealSubject struct{}

func (s RealSubject) DoAction() {
	fmt.Println("Real action")
}

type Proxy struct {
	Subject RealSubject
}

func (p Proxy) DoAction() {
	fmt.Println("Proxying...")
	p.Subject.DoAction()
}

func NewProxy() Proxy {
	return Proxy{Subject: RealSubject{}}
}

func main() {
	obj := NewProxy()
	obj.DoAction()
}
