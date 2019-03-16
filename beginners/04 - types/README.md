## Types

The following types are available in Golang:
* bool
* Numeric Types:
  * int8, int16, int32, int64, int
  * uint8, uint16, uint32, uint64, uint
  * float32, float64
  * complex64, complex128
  * byte
  * rune
* string

`bool` represents a boolean value with one of `true` or `false`

For numeric types, the numbers alongside the types indicate the number of bits that will be used to represent such values in memory. The number of bits determine how big the number would be.

---

`int8`: represents 8 bit signed integers 

`size`: 8 bits 

`range`: -128 to 127

---

`int16`: represents 16 bit signed integers 

`size`: 16 bits 

`range`: -32768 to 32767

---

`int32`: represents 32 bit signed integers 

`size`: 32 bits 

`range`: -2147483648 to 2147483647

---

`int64`: represents 64 bit signed integers

`size`: 64 bits

`range`: -9223372036854775808 to 9223372036854775807

---

`int`: represents 32 or 64 bit integers depending on the underlying platform. You should generally be using int to represent integers unless there is a need to use a specific sized integer. 

`size`: 32 bits in 32 bit systems and 64 bit in 64 bit systems. 

`range`: -2147483648 to 2147483647 in 32 bit systems and -9223372036854775808 to 9223372036854775807 in 64 bit systems

---

`uint8`: represents 8 bit unsigned integers

`size`: 8 bits

`range`: 0 to 255

---

`uint16`: represents 16 bit unsigned integers

`size`: 16 bits

`range`: 0 to 65535

---

`uint32`: represents 32 bit unsigned integers

`size`: 32 bits

`range`: 0 to 4294967295

---

`uint64`: represents 64 bit unsigned integers

`size`: 64 bits

`range`: 0 to 18446744073709551615

---

`uint` : represents 32 or 64 bit unsigned integersdepending on the underlying platform.

`size` : 32 bits in 32 bit systems and 64 bits in 64 bit systems.

`range` : 0 to 4294967295 in 32 bit systems and 0 to 18446744073709551615 in 64 bit systems

---

`float32`: 32 bit floating point numbers

`float64`: 64 bit floating point numbers

---

`complex64`: complex numbers which have float32 real and imaginary parts

`complex128`: complex numbers with float64 real and imaginary parts

---

The builtin function `complex` is used to construct a complex number with real and imaginary parts. The complex function has the following definition:

```go
func complex(r, i FloatType) ComplexType
```

It takes a real and imaginary part as parameter and returns a complex type. Both the real and imaginary parts should be of the same type. ie either float32 or float64. If both the real and imaginary parts are float32, this function returns a complex value of type complex64. If both the real and imaginary parts are of type float64, this function returns a complex value of type complex128

Complex numbers can also be created using the shorthand syntax:

`c := 6 + 7i`


Below is an example of working with complex numbers:

```go
c1 := complex(5, 7)
c2 := 8 + 27i
cadd := c1 + c2
fmt.Println("sum:", cadd)
cmul := c1 * c2
fmt.Println("product:", cmul)
```

---

`byte` is an alias of `uint8`

`rune` is an alias of `int32`

---

`string` is a collection of characters enclosed in double quotes.

```go
first := "Allen"
last := "Varghese"
name := first +" "+ last
fmt.Println("My name is",name)
```

---

The type of a variable can be printed using `%T` format specifier available in the `Printf` method.

```go
var a = 10
fmt.Printf("a is of type %T, value is ", a, a)
```

---

### Type Conversion

There is no automatic type conversion in Golang. Use the corresponding type functions to do the conversion.

```go
int(<float value>)
float64(<integer value>)
```


### Constants

Constants are values that do not change once assigned to a variable. The keyword `const` is used instead of `var` to declare a constant value along with the type definition.

```go
const a bool = true
const b int32 = 32890
const c string = "Something"
```

