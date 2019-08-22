package main

import (
    "fmt"
    "io/ioutil"
)

func grade(score int) string {
    g := ""
    switch {
    case score < 60:
        g = "F"
    case score < 80:
        g = "C"
    case score < 100:
        g = "A"
    default:
        panic(fmt.Sprintf("Wrong score: %d",score))
    }
    return g
}

func main() {
    const  filename  = "abc.txt"
    //地址 GOPATH=/home/lemon/gowork
    /*
    contents, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("%s\n",contents)
    }
    */

    //简化 注意 此contents err 变量只在 此 if代码块内有效
    if contents, err := ioutil.ReadFile(filename); err != nil{
        fmt.Println(err)
    } else {
        fmt.Printf("%s\n",contents)
    }
}
