/* ЗАДАНИЕ:
 * Напишите функцию sumInt, получающую переменное число аргументов типа int,
 * и возвращающую количество переданных аргументов и их сумму.
 */

func sumInt  (a ...int) (int, int){
    var sum int
    for i:=0;i<len(a);i++{
        sum = a[i]+ sum

    }
    return len(a), sum
}
