package main

import "fmt"

//值传递
func swap(a, b int)  {
    b, a = a, b
}

//指针传递
func swap2(a, b *int)  {
    *b, *a = *a, *b
}

//最好的做法
func swap3(a, b int) (int,int)  {
    return b, a
}

func main() {
    a, b := 3, 4
    swap(a, b)
    fmt.Println(a, b)
    swap2(&a, &b)
    fmt.Println(a, b)

    a, b = swap3(a, b)
    fmt.Println(a, b)
}
