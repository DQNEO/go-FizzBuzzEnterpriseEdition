package main

import "github.com/DQNEO/go-FizzBuzzEnterpriseEdition/logic"
import "fmt"


func main() {
    start := 1
    end := 30

    numbers := make([]int, end, end)

    for i := start  - 1; i < end ; i++ {
        numbers[i] = i + 1
    }

    for _,v := range(numbers) {
        s := logic.Logic(v)
        fmt.Println(s.String())
    }

}

