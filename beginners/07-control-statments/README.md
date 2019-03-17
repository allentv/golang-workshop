# Control statements

Golang supports multiple control statements including loops.

## If - else if - else

This is a conditional statement and follows the syntax as below:

```go
if condition {
    // block to be executed
}
```

Having multiple `if` statements is also possible as shown below:

```go
if condition {
    // body
} else if condition {
    // body
} else {
    // body
}
```

There can be any number of `else if` statements. The `else` block is optional.

**Note:** The `else` should be on the same line as the closing brace (`}`) due to the way Golang automatically inserts semi-colons. More details [here](https://golang.org/ref/spec#Semicolons).

The `if` statement also supports an optional statment before the condition is executed as shown below:

```go
if statement; condition {
    // body
}
```

This is useful when an initialization is required before the condition or following body is executed. Notice the `;` that is used to separate the statement and the condition.

## for loops

There is only one loop statement available in Golang. It is the `for` loop. There is no `while` or `do .. while` loop as present in other languages.

The syntax is as below:

```go
for initialization; condition; post {
    // body
}
```

The initialization is executed only once. After the loop is initialised, the condition will be checked. If the condition evaluates to true, the body of loop inside the `{ }` will be executed followed by the post statement. The post statement will be executed after each successful iteration of the loop. After the post statement is executed, the condition will be rechecked. If it's true, the loop will continue executing, else the for loop terminates.

All the three components namely initialisation, condition and post are optional in Golang.

Below is an example of printing numbers from 1 to 10 using a `for` loop.

```go
for i := 1; i <= 10; i = i + 1 {
    fmt.Println(i)
}
```

The execution starts with the initialization statement where the variable `i` is initialized with the value of `1`. The condition is then checked `1 <= 10` which is statisfied and then the body of the loop is executed which prints `1` to the terminal. As there is nothing else to execute, the post statement is executed which increments the value of the variable `i` by `1`. The condition is again checked to confirm if it is true. If so, then then body of the `for` loop is executed again. This repeats until the condition becomes false.
