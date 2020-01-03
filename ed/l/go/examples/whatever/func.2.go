package main

func main()  {
	f1()

	funcs := map[string]func() {"f1": f1}

	funcs["f1"]()
}

func f1() {
	println("It works.")
}

/*
It works.
It works.
*/
