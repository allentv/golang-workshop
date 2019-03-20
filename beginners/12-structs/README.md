# Structs

A struct is a user defined data type which represents a collection of fields very similar to the struct data structure in C. It can be used in places where it makes sense to group different data types together.

```go
type Employee struct {
    firstName string
    lastName string
    age int
}
```

The above code snippet defines a `struct` with the name `Employee` which defines attributes of an employee like first name, last name and age with the respective types of each attribute. The `type` keyword is used to define a `struct` with a name so that it can be re-used anywhere else in the program with the name `Employee` to refer to an employee record.

When a struct is defined and not explicitly initialised, then the default values for the fields are assigned to it.

Individual fields of a struct can be accessed using the dot notation.

```go
var emp := Employee {
    firstName: "Something"
}

fmt.Println(emp.firstName)

> Something
```

The `firstName` field is accessed using the dot notation. Trying to access other uninitialised fields will result in default values.

## Nested structs

It is possible to nest a struct within another struct and usually done for code re-use. This is also a design pattern called `composition` and is widely used these days.

```go
type Address struct {
    city, state string
}

type Employee struct {
    firstName, lastName string
    age int
    address Address
}

var emp Employee
emp.firstName = "SomeThing"
emp.address = Address {
    city  : "AA",
    state : "CO",
}

fmt.Println(emp)

> {SomeThing  {AA CO}}
```

The address of the employee is initialised separately from the other fields of the struct. When the employee record is printed, all values are displayed.

The fields in nested structs can be accessed directed by following multiple levels of the dot notation.

```go
fmt.Println(emp.address.city)

> AA
```

## Visibility

A struct and its fields are visible outside of a package depending on the case of the first letter in the name. If the employee struct is named `employee` instead of `Employee`, it will not be visible outside of the package it is declared in. This is called an **exportable** type. The same condition also applies to the fields in a struct. If the field names start with a capital letter, then they can be accessed from other packages.

## Equality

A struct is a value type. 2 struct variables are considered equal if all the values of their fields are equal. This is also means that if the structs have fields that cannot be compared like a `map`, then the equality operation (`==`) will fail.

## Methods/Receivers

Methods are functions associated with a specific type and is similar to class methods in the OOP world. The syntax is given below:

```go
func (t Type) MethodName(parameter list) {
    // body
}
```

The methods for a type are usually defined in the same file where the type is defined to keep related code together.

A `print` method can be added for the `Employee` type that was defined earlier which will print the details of an employee record in a formatted fashion.

```go
func (e Employee) Print() {
    fmt.Println("Employee Record:")
    fmt.Println("Name:", e.firstName, e.lastName)
    fmt.Println("Address:", e.address)
}

var emp Employee

emp.Print()

> Employee Record:
  Name: Allen Varghese
  Address: {AA CO}
```

The method `Print` will print out the details of an employee record in a formatted manner.

The main advantage of using methods over just pure functions is that we can define methods with the same name for multiple types but a function can only be defined once with a name.

## Using pointers

Golang support pointers to some extent mainly for updating values in place but does not support pointer arithmetic like C. The `*` is prefixed to the type name while defining a pointer receiver.

```go
func (t *Type) MethodName(parameter list) {
    // body
}
```

The point to note is that, a pointer receiver can also accept a value receiver but not the other way around. Pointers are used when it is usually expensive to copy data around and it is much easier to make a change in place using the pointer.

```go
func (e *Employee) updateAge(newAge int) {
    e.age = newAge
}
emp := Employee{
    age: 33,
}

fmt.Println("Before:", emp.age)
emp.updateAge(34)
fmt.Println("After:", emp.age)

> Before: 33
  After: 34
```

Here the age of the employee has been updated in place using a pointer receiver. Notice the `*` before the type `Employee` when the `updateAge` method was defined.