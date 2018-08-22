package main

func main() {
  defer println("top")
  defer println("bottom")

  println("main")
}

/*

main
bottom
top

*/
