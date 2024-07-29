package main

import (
    "fmt"
)

type Values struct{
    nm float64
}

var dct = map[int]Values{1:{1},2:{2},3:{3}}

func main(){
    // var dct = make(map[int]Values)

    // dct[1] = Values{1}

    fmt.Println(dct)
}