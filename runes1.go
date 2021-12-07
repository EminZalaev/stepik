package main

import (
 "fmt"
  "os"
  "bufio"
)

func main(){
    text,_:= bufio.NewReader(os.Stdin).ReadString('\n')
    bs := []rune(text)
    for i:=0;i<len(bs);i++{
        if i%2!=0{
            fmt.Print(string(bs[i]))
        }
    }
}
