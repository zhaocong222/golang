package main

import (
    "fmt"
    "math"
)

func consts()  {
    const (
        filename = "abc.txt"
        a, b  = 3, 4
    )
    var c int
    //这里定义的是常量, 所以 a*a + b*b 不用转换未float64
    c = int(math.Sqrt(a*a + b*b))
    fmt.Println(filename, c)
}

//枚举类型
func enums()  {
    //自增枚举, _表示跳过
    const (
        cpp = iota
        _
        python
        golang
        javascript
    )

    fmt.Println(cpp, javascript, python, golang)

}

func main() {
    //consts()
    enums()
}
