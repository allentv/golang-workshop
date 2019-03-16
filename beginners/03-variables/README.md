## Variables

A variable is the name given to a memory location that stores a value for a particular type.

The `var` keyword is used to declare variables. The syntax is `var <name> <type>` followed by optional value assignment.

Below are few examples:
```go
var a int // Default initialization

var b = 10 // Automatic type inference

var c, d = 20, 30 // Multiple value assignment

e, f := 40, 50 // Shorthand assignment
```

There are default values for different types of values that Golang uses if a value is not assigned to a variable during variable declaration.

Printing the value of a variable is done in the following way

`fmt.Println("a = ", a)`

The value of variable `a` is displayed when used in the above manner.

Follow the accompanying code file - `main.go` to try out the various variable usages.
