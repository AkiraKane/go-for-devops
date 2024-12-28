// Creating Maps
// 1. Using var: Creates a nil map. The map is not initialized and cannot store data until initialized.
var m map[string]int
fmt.Println(m) // map[]

m["key"] = 1 // Panic: assignment to entry in nil map

// 2. Using make: Initializes a map ready for use.
m := make(map[string]int) // Empty map
m["key"] = 42
fmt.Println(m) // map[key:42]

// You can also predefine the map's size for better performance:
m := make(map[string]int, 10) // Pre-allocates space for 10 entries
fmt.Println(m) // map[]

// 3. Using a Composite Literal: Defines and initializes a map inline.
m := map[string]int{
    "one": 1,
    "two": 2,
}
fmt.Println(m) // map[one:1 two:2]
