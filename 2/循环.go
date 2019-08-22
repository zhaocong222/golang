package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

//整数转二进制字符串
func convertToBin(n int) string  {
    result := ""
    for ; n > 0; n /= 2{
        lsb := n % 2
        //整数转字符串strconv.Itoa
        result = strconv.Itoa(lsb) + result
    }
    return result
}

//读取文件
func printFile(filename string)  {
    file, err := os.Open(filename)
    if err != nil {
        //panic 报错终止程序往下执行
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    //按行读取
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

}


func main() {

    printFile("abc.txt")
    /*
    fmt.Println(
        convertToBin(5), // 101
        convertToBin(13),// 1101
        )
     */
}
