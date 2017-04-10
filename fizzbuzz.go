package main

import "github.com/DQNEO/go-FizzBuzzEnterpriseEdition/logic"
import "fmt"


func getSlice(start int, end int) []int {
    numbers := make([]int, end, end)

    for i := start  - 1; i < end ; i++ {
        numbers[i] = i + 1
    }
    return numbers
}

func main() {
    start := 1
    end := 30

    numbers := getSlice(start, end)
    for _,v := range(numbers) {
        s := logic.Logic(v)
        fmt.Println(s.String())
    }

}

