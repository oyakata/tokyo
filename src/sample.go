package main

import fmt "fmt"


func main() {
    x := ""
    // stringにappendは使えない
    //   first argument to append must be slice; have string
    // append(x, "a", "b")

    x += "abc"

    fmt.Println(x)
}
