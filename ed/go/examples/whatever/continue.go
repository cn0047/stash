package main

func main() {
  for i := 0; i < 10; i++ {
    if i > 0 {
      if i > 1 {
        if i > 2 {
          if i > 3 {
            continue
          }
        }
      }
    }
    println(i)
  }
}

/*

0
1
2
3

*/
