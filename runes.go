package main

import (
 "fmt"
    "bufio"
    "os"
)

func pal() {

    
}

func main(){
    text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    rs := []rune(text)
    k:=0
    for i:=0; i<len(rs);i++{
        if rs[i]==rs[len(rs)-1-i]{
                k++
            }else{
                break

        }
    }
    if k==len(rs){
    fmt.Print("Палиндром")
    }else{
                fmt.Print("Нет")
    }

}
