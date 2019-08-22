package main

import (
    "fmt"
    "math"
    "math/cmplx"
)

func euler()  {

    fmt.Println(
            cmplx.Pow(math.E, 1i * math.Pi) + 1)
}

//强制类型转转
func triangle()  {
    var a, b int = 3,4
    var c int
    //必须要强制转义类型 math.Sqrt 接收float64类型参数
    c = int(math.Sqrt(float64(a * a + b * b)))
    fmt.Println(c)
}



func main() {
    //euler()
    triangle()
}
