// Defining a Struct
type Person struct {
    Name   string
    Age    int
    Gender string
}

// Using a Struct Literal You can initialize all fields explicitly using a struct literal.
p := Person{Name: "Alice", Age: 30, Gender: "Female"}
fmt.Println(p)

// You can omit field names if you assign values in order (not recommended for readability):
p := Person{"Alice", 30, "Female"}
fmt.Println(p)

// Using the Zero Value Create a struct with fields initialized to their zero values (empty strings, 0, false, etc.).
var p Person
fmt.Println(p) // { 0 }

// Pointer to a Struct Use the & operator to get a pointer to a struct.
p := &Person{Name: "Bob", Age: 25, Gender: "Male"}
fmt.Println(p)       // &{Bob 25 Male}
fmt.Println(*p)      // {Bob 25 Male}
fmt.Println(p.Name)  // Bob (fields can be accessed directly)

// Partial Initialization You can initialize only some fields, leaving others at their zero values.
p := Person{Name: "Charlie"}
fmt.Println(p) // {Charlie 0 }
