package main

import (
	"fmt"
)

type Info struct {
	Text string
}

func RedCliDecorator(s string) string {
	return fmt.Sprintf("\033[34m %s \033[0m", s)
}

func BracketsDecorator(s string) string {
	return fmt.Sprintf("[%s]", s)
}

func EmojiDecorator(s string) string {
	return fmt.Sprintf("🔰 %s 🔰", s)
}

func Decorate(f func(s string) string) func(str string) string {
	return func(str string) string {
		return f(str)
	}
}

func main() {
	i := Info{Text: "it works"}

	s := BracketsDecorator(RedCliDecorator(i.Text))
	// or
	r := Decorate(EmojiDecorator)(s)

	fmt.Println(r)
}
