package main

import "fmt"


func add(x,y int) (int, string) {
    return x + y, "test"
}

func main() {
    fmt.Println(add(100, 200))
}
