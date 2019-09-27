package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
)

func main() {
    //text := "aaabccddd"
    //text := "abccddd"
    //text := "abddd"
    //text := "ffxyzz"
    //text := "bb"
    //text := "baab"
    text := readInput()

    pr := ""
    r := text
    var loop = true
    for loop {
        isOk, nr := f(r)
        pr = r
        r = nr
        if isOk {
            if pr == r {
                loop = false
                fmt.Println(r)
            }
        } else {
            loop = false
            fmt.Println(r)
        }
    }
}

func readInput() (string) {
    reader := bufio.NewReader(os.Stdin)
    data, _ := reader.ReadString('\n')
    return data
}

func f(s string) (isOk bool, result string) {
    if s == "" {
        return false, "Empty String"
    }
    res := make([]string, 0)
    for i := 0; i < len(s); i++ {
        c := string(s[i])
        var nc string
        if i + 1 == len(s) {
            nc = ""
        } else {
            nc = string(s[i + 1])
        }
        if c == nc {
            i++
        } else {
            res = append(res, c)
        }
    }
    r := strings.Join(res, "")
    if r == "" {
        return false, "Empty String"
    } else {
        return true, r
    }
}
