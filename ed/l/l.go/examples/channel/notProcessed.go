package main

func main() {
	ch := make(chan string)

	go func() {
		for m := range ch {
			println("processed:", m)
		}
	}()

	ch <- "cmd.1"
	ch <- "cmd.2" //won't be processed
}

// @link: https://play.golang.org/p/De7XyxEKHt6
