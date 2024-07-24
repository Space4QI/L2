Что выведет программа? Объяснить вывод программы.

package main
 
type customError struct {
     msg string
}
 
func (e *customError) Error() string {
    return e.msg
}
 
func test() *customError {
     {
         // do something
     }
     return nil
}
 
func main() {
    var err error
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
}

Ответ: программа выведет "error", потому что, несмотря на то, что функция test возвращает nil типа *customError, 
переменная err типа error содержит тип *customError, и проверка err != nil возвращает true.