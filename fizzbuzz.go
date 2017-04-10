package main

import "github.com/DQNEO/go-FizzBuzzEnterpriseEdition/logic"
import "fmt"


func main() {
    var start,end int = 1,30
    for i := start; i <= end ; i++ {
        s := logic.Logic(i)
        fmt.Println(s.String())
    }
}

