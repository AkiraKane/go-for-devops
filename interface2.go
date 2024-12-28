// type assertion
package main
 
import (
    "fmt"
)
 
type B struct{
    s string
}
 
func main() {
    var i interface{} = B{"a sample string"}  
     
    fmt.Println(i.(B)) // {a sample string}
}
