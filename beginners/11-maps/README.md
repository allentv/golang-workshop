# Maps

A built-in type in Golang that associates a value to a key. The value can be retrieved by specifying the key. A map is created by the `make` function with the below syntax

```go
make(map[type of key]type of value)

eg: make(map[string]int)
```

The above example will create a map with the key of type `string` that would save/retrieve an integer. An empty map can be created without using the `make` function and will have the default value of `nil`. An error will be raised if a key is referenced from an empty map.

## Adding Values

Values in a map are referenced the same way as index based referencing works in arrays. A value is added by assigning a value to a new key. If the specified key already exists, then its value will be overwritten with the new value.

```go
m := make(map[string]int)
m["a"] = 1
m["b"] = 2

fmt.Println(m)

> map[a:1 b:2]
```

A new map is created with the `make` function. Two keys - `a` and `b` are assigned values of `1` and `2` respectively. Note that the types of the key and values should make match with what has been specified during the map declaration else Golang compiler will complain.

A map can also be initialized with values during declaration just like arrays.

```go
m := map[string]int {
    "a": 1,
    "b": 2,
}
```

The above snippet defines 2 (key, value) pairs during declaration.

## Retrieving Values

Values can be retrieved from a map similar to the way values are assigned. The square bracket notation with the name of the key is used to get the value associated with a key. An error is raised if the value is not present when the value is retrieved this way.

A handy way to check if a key is present in a map is below:

```go
data, ok := m["some key"]
```

If `ok` is `true`, then the key is present in the map and `data` will contain the value else the key will be absent.

## Iteration

The `for` loop can be used to iterate over a map. Instead of index and value that is returned when iterating over an array, here the key and value is returned.

```go
m := map[string]int {
    "a": 1,
    "b": 2,
    "c": 3,
}
for key, value := range m {
    fmt.Println("Key:", key, " Value:", value)
}

> Key: a  Value 1
  Key: b  Value 2
  Key: c  Value 3
```

The `for` loop will iterate over the map until all (key, value) pairs are exhausted.

```text
The order of retrieval from the map is not guaranteed each time the for loop iterates over the map
```

## Delete

Values from a map can be deleted using the `delete` function. The function does not return anything. The syntax is given below:

```go
delete(m, key)
```

## Length

The length of the map is returned using the `len` function. This would return the count of keys in a map.

Syntax is

```go
len(m)
```

## Type

A map is a reference type like a slice. So if the same map is assigned to another variable, both point to the same underlying variable unlike an array where a copy of the original array is assigned to the new variable. If a map is used as an argument into a called function, then any updates made to the calling function will reflect in the caller function.
