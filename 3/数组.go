package main

import "fmt"

//传递参数 需要5个元素的 int数组
func printArray(arr [5]int)  {
   for i, v := range arr {
       fmt.Println(i, v)
   }
}

//指针
func printArray2(arr *[5]int)  {

    arr[0] = 99
    arr[2] = 101

    for i, v := range arr {
        fmt.Println(i, v)
    }
}


func main() {
    var arr1 [5]int
    arr2 := [3]int{1, 3, 5}
    arr3 := [...]int{2, 3, 4}

    //二位数组
    var grid [4][5]int

    fmt.Println(arr1, arr2, arr3)
    fmt.Println(grid)

    for i, v := range arr3 {
        fmt.Println(i, v)
    }

    //不需要下标
    for _, v := range arr3 {
        fmt.Println(v)
    }

    fmt.Println("arr1")
    printArray(arr1)

    //传递指针
    fmt.Println("arr1")
    printArray2(&arr1)

    fmt.Println("外arr1")
    fmt.Println(arr1)

}
