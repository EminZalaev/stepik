package main

import (
 "fmt"
 "unicode"
 "bufio"
 "os"
)

func main() {
    text,_:=bufio.NewReader(os.Stdin).ReadString('\n')
    rs:=[]rune(text)
    var k int
    for l:=0;l<len(rs);l++{
        if (unicode.Is(unicode.Latin, rs[l])==true&&unicode.IsLetter(rs[l])==true||unicode.IsDigit(rs[l])==true)&&len(rs)>=5{
            k++
        }else{
            fmt.Print("Wrong password")
            break
        }
    }
    if k==len(rs){
        fmt.Print("Ok")
    }
}
