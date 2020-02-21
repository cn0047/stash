package main

func main() {
	ch := make(chan string)

	go func() {
		for m := range ch {
			println("processed:", m)
		}
	}()

	ch <- "cmd.1"
	ch <- "cmd.2"
	ch <- "cmd.3"
}

// @link: https://play.golang.org/p/Ut_zSB6zvPi
