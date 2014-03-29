package main

import "fmt"


func add(x,y int) int {
    return x + y
}

func main() {
    sum:=0
    sum=add(100, 200)
    if sum==300 {
       fmt.Println("we added 100 and 200 correctly")
    } else {
       fmt.Println("computers suck at math")
    }
}
