package main

import "fmt"

func main(){
    var d uint16
    
    fmt.Scan(&d)
    fmt.Printf("It is %d hours %d minutes.", d / 30, 2 * (d % 30))
}
