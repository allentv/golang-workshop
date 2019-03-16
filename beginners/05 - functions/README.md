## Functions

A function is a group of statements that execute a block of work. A function may or may not take input and may or may not return a value back to the caller.

A function is defined using the `func` keyword and has the following format:

```go
func <name>(<arg1> <type of arg1>, ...) <return type> {
    // body of the function spanning one or more lines
}
```

Below is an example of a function that adds two integers.

```go
func AddIntegers(a int, b int) int {
    return a + b
}
```

The `return` keyword is used to signify which value has to be returned from the function.

Invoking the function with 2 integers would return the sum of the values

```go
AddIntegers(10, 20) => 30
```

Unlike most languages, Golang supports returning multiple values from a function. Just use comma separated variables after the `return` statement.

```go
return a, b
```

Below is a function that returns the sum and difference of two integers.

```go
func SumDifference(a int, b int) (int, int) {
    return a + b, a - b
}
```

In the above example, instead of assigning the sum and difference to two variables and then returning them, the steps are simplified to have the calculation happen in one line with the `return` statement.


#### Blank Identifier

The blank identifier is used in place of any value of any type to discard the value especially in cases where one or more of the return values from a function are not being used.

```go
var _, diff = SumDifference(10, 20)

fmt.Println("Difference is ", diff)
```

The above code snippet will only save the difference of two numbers to `diff` variable even though both sum and difference are calculated by the `SumDifference` function.

The order of return values matter. The masking of values happens in the same order as the values are returned from the function.

#### Named return values

When defining a function, the return types can be given a name which can then be referenced in the body of the function.

```go
func Product(a int, b int) prod int {
    prod = a * b
    return
}
```

The above code snippet assigns a name `prod` to the return value of the function. The `prod` can be used as a variable and values assigned to it as a regular variable. Since `prod` is already defined as part of the function signature, it does not have to be included with the `return` statement.
