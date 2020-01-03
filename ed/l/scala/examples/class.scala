class Foo {
    def log(msg: String) {
        println(msg)
    }
    def bar() {
        println("bar")
        this.log("OUT: bar")
    }
}

val f = new Foo()
f.bar()
