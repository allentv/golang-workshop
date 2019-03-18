# Arrays

A collection of elements of the same type is called an Array. Since Golang is a strongly typed language, it is not possible to mix values of different types.

An array is defined as `[size]type`. Both parts are required to be considered an array. Differently sized arrays with the same type are not considered equal as the size is part of the array definition.

```go
var a [3]int
fmt.Println(a)

> [0 0 0]
```

Here we define a 3 element integer array. The output is all zeros because that is default value of an integer type.

Array indexing starts from `0` instead of `1`. Values can be assigned to different indexes of an array by referencing the index and providing a value via the `=` operator.

```go
var a [3]int
a[0] = 1
a[1] = 2
a[2] = 3

fmt.Println(a)

> [1 2 3]
```

An array can also be initialized with values at declaration time.

```go
a := [3]int{1, 2, 3}
```

You can skip the number of elements when initializing if elements are provided

```go
a := [...]int{1, 2, 3}
```

## Value Type

An important thing to note is that an Array is a value type i.e. each time an array is assigned to a new variable, a copy of the original array is made and then assigned to the new variable. If changes are made to the new variable, it will not be reflected in the original array.

```go
a := [...]string{"IRL", "IND", "US", "CAN"}
b := a

b[1] = "CHN"

fmt.Println("Original:", a)
fmt.Println("Copy    :", b)
```

This also means that when an array is passed as an input argument to a function, it is passed by value. So any changes made to the input argument within the function, does not reflect in the array used in the caller function.

The length of an array is found by the `len` function.

```go
a := [...]int{1, 2, 3}
fmt.Println(len(a))

> 3
```

A concise way of iterating an array is to use the `range` keyword. It returns both the index and value at that index. Here is an example of finding the sum of all the values in an array and printing out the value at each index while the sum is calculated

```go
a := [...]int{1, 2, 3, 4, 5}
sum := 0
for i, v := range a {
    fmt.Println("Index:", i, " Value is:", v)
    sum += v
}
fmt.Println("Sum:", sum)

> 15
```

The array of integers from 1 to 5 is assigned to the variable `a`. A new variable `sum` is declared to save the running sum of values. At each iteration, the index and value is printed out. Finally the total sum is printed.

The blank identified (`_`) can be used if only one of the return values from `range` is required.

```go
_, v := range a
```

The index is discarded here and only the value at the current index is available.

## Multi-dimensional arrays

Arrays of more than one dimension can be created by adding more sizes as below:

```go
a := [3][2]int
```

The above line defines a 3 x 2 array i.e. 3 rows and 2 columns.

```go
a := [3][2]string{
    {"lion", "tiger"},
    {"cat", "dog"},
    {"pigeon", "peacock"},
}

fmt.Println(a)

> [[lion tiger] [cat dog] [pigeon peacock]]
```

The above output can be changed to print the values in different lines using multipe for loops as shown below

```go
func printArray(a) {
    for _, v1 := range a {
        for _, v2 := range v1 {
            fmt.Printf("%s ", v2)
        }
        fmt.Println("")
    }
}
```

The first `range` iterates through the rows and the second `range` iterates through the columns. Each of the column value is printed out with a space after it. At the end of printing an entire row, a new line is added.

## Slices

Arrays have fixed lengths and pre-allocated in memory and are value types. A slice is a flexible wrapper on top of arrays. They don't own any data but only point to an underlying array.

An empty slice can be created as below

```go
var sa []int
```

The value of a empty slice is `nil`

A slice can also be created by pointing to a subset of values in an array as below:

```go
a := [...]int{1, 2, 3, 4, 5}
sa := a[1: 4]
fmt.Println(sa)

> [2 3 4]
```

The slice `sa` points to elements in the array `a` from indices `1` to `3`. When specifying a range of indices in an array, the last index is not considered.

Since a slice is a reference type, modifying the value at an index will update the underlying array.

```go
a := [...]int{1, 2, 3, 4, 5}
sa := a[1: 4]

fmt.Println("Before:", a)
sa[0] = 22

fmt.Println("After:", a)

> Before: [1 2 3 4 5]
> After: [1 22 3 4 5]
```

Here the first index of the slice points to index `1` of the underlying array. So updating the value from `2` to `22` in the slice will update the array value too.

A slice can also be created using the `make` function by providing the type and the length with an optional capacity that signifies the maxium size the slice can grow to.

```go
make([]type, length[, capacity])
```

Creating a slice with `make` will initialized with the default values of the type that is used

```go
i := make([]int, 5, 5)
fmt.Println(i)

> [0 0 0 0 0]
```

The size of a slice can be increased by using the `append` function. This will create a new underlying array with the new size and copy over the existing value from the original array and append the new values at the end.

The syntax for the append function is

```go
append(destination, value1, value2, ...)
```

Instead of a values, it can also be another slice. The `...` operator is used to expand the slice into it's values

```go
sa := []int{1, 2, 3}
newSa := append([]int{}, sa...)

fmt.Println(newSa)

> [1, 2, 3]
```

The `append` function is to used to append 3 values from an existing slice `sa` into an empty slice.

A slice can be passed into a function just like an array. Modifying the values of a slice in the called function, will update the underlying array in the caller function unlike arrays which is a value type. This is because a slice is a reference to an underlying array.