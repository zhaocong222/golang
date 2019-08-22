package main

import "fmt"

func eval(a, b int, op string) int {
    switch op {
    case "+":
        return a + b
    case "-":
        return a -b
    default:
        panic("unsupported operation:" + op)
    }
}

func sumArgs(values ...int) int  {
    sum := 0
    for i := range values {
        sum += values[i]
    }
    return sum
}

func div(a, b int) (int, int)  {
    return  a / b, a % b
}

func main() {
    fmt.Println(sumArgs(1,2,3,4))
}
