package main
 
import (
    "fmt"
    "reflect"
)
 
type Flyer interface{
    fly() string
}
 
type Walker interface{
    walk() string
}
 
type Bird struct{
    Name string
}
 
func (b *Bird)fly() string{
    return "Flying..."
}
 
func (b *Bird)walk() string{
    return "Walking..."
}
 
func main() {
    var b = Bird{"Chirper"}
     
    fmt.Println(b.fly())     // Flying...
    fmt.Println(b.walk())    // Walking...  
}
