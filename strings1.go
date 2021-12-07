/*Дана строка, содержащая только десятичные цифры. Найти и вывести наибольшую цифру.

Входные данные

Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков и строка содержит только десятичные цифры.

Выходные данные

Выведите максимальную цифру, которая встречается во введенной строке.

Sample Input:

1112221112
Sample Output:

2
*/

package main

import (
    "fmt"
    "strings"
    "strconv"
)

func main(){
    var a string
    fmt.Scan(&a)
    d := strings.Split(a, "")
     i, _ := strconv.Atoi(d[0])
    sum:=i
   for j:=0;j<len(d);j++{
     i, _ := strconv.Atoi(d[j])
         if i>sum{
         sum=i

     }
    }
    fmt.Print(sum)
}
