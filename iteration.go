// Iterating Over a Slice
// 1. Using for with Index: You can loop through the slice by accessing elements via their index.
slice := []int{10, 20, 30, 40, 50}
for i := 0; i < len(slice); i++ {
    fmt.Println("Index:", i, "Value:", slice[i])
}

// 2. Using range: The range keyword provides both the index and value.
slice := []int{10, 20, 30, 40, 50}
for i, v := range slice {
    fmt.Println("Index:", i, "Value:", v)
}

// If you only need the values, ignore the index using _:
for _, v := range slice {
    fmt.Println("Value:", v)
}

// Iterating Over a Map
// 1. Using range: The range keyword iterates over key-value pairs in a map.
m := map[string]int{"one": 1, "two": 2, "three": 3}
for k, v := range m {
    fmt.Println("Key:", k, "Value:", v)
}

// 2. Keys Only: If you only need the keys, ignore the value using _.
for k := range m {
    fmt.Println("Key:", k)
}

// 3. Values Only: If you only need the values, iterate through range and ignore the key.
for _, v := range m {
    fmt.Println("Value:", v)
}
