package main
 
import (
    "fmt"
)
 
type Person interface{
    greet() string
}
 
type Human struct{
    Name string
}
 
func (h *Human)greet() string {
    return "Hi, I am " + h.Name 
}
 
func isAPerson(h Person) {
    fmt.Println(h.greet())
}
 
func main() {
    var a = Human{"John"}
     
    fmt.Println(a.greet())    // Hi, I am John
     
    // below function will only work
    // if a is also a person.
    // Here we can see polymorphism in action.
    isAPerson(&a)             // Hi, I am John      
}
