package main

import fmt "fmt"

var LETTERS = []string{"A", "B", "C", "D", "E", "F", "G",
                       "H", "I", "J", "K", "L", "M", "N",
                       "O", "P", "Q", "R", "S", "T", "U",
                       "V", "W", "X", "Y", "Z"}


func to_i(number int) string {
    if number < 26 {
        return LETTERS[number]
    } else {
        result := ""
        var x int = number
        var y int

        // while x > 0 {  // なんと、goにはwhile文が無い！forでやれとのこと
        for x > 0 {
            y = x % 26
            result = LETTERS[y - 1] + result
            x = (x - y) / 26
        }
        return result
    }
}


func main() {
    fmt.Println(to_i(1))
    fmt.Println(to_i(4878821185187))
}
