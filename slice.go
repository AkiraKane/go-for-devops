// Creating Slices
// 1. Using var: Creates a nil slice. The slice has no elements and no underlying array allocated.
var slice []int
fmt.Println(slice, len(slice), cap(slice)) // [] 0 0

// 2. Using make: Allocates and initializes an empty slice with a specified length and capacity.
slice := make([]int, 5) // Creates a slice of length 5 and capacity 5
fmt.Println(slice)     // [0 0 0 0 0]

slice2 := make([]int, 3, 5) // Creates a slice of length 3 and capacity 5
fmt.Println(slice2)         // [0 0 0]

// 3. Using a Composite Literal: Defines and initializes a slice inline.
slice := []int{1, 2, 3, 4, 5}
fmt.Println(slice) // [1 2 3 4 5]

// 4. Slicing an Array: Slices are views of arrays. Slicing an array creates a new slice.
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4] // Slice from index 1 to 3 (4 is exclusive)
fmt.Println(slice) // [2 3 4]

// 5. Using append: You can dynamically grow slices using append.
slice := []int{} // []
slice = append(slice, 1, 2, 3)
fmt.Println(slice) // [1 2 3]
