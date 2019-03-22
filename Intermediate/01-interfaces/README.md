# Interfaces

The definition of interfaces from the object oriented programming (OOP) world is that it defines what an object will do, as supposed to how it will be done. The implementation detail for what an object will do, is upto the implementation class.

In Golang, an interface is a set of method signatures. Any type that has a definition for those methods, are said to implement the interface. Unlike other languages, the interface association is implicit. An explicit reference of implementing the interface is not required.

The main advantage of using interface is code reusability. It is common to have a type implement more than one interface.

Consider an example for understanding interfaces better. A company has employees under 2 categories : permanent and contract. They both get a salary at the end of the month but the salary calculation is different. This situation can be easily represented in Go using interfaces.

```go
type Permanent struct {
    empID int
    basicPay int
    pension int
}

type Contract struct {
    empID int
    basicPay int
}

func (p Permanent) calculateSalary() int {
    return p.basicPay + p.pension
}

func (p Contract) calculateSalary() int {
    return p.basicPay
}

type SalaryCalculator interface {
    calculateSalary() int
}

// In main
sc := [...]SalaryCalculator{Permanent{1, 5000, 50}, Contract{3, 7000}}
totalPayout := 0
for _, v := range sc {
    totalPayout += v.calculateSalary
}
fmt.Println(totalPayout)
```

The above code example is a bit long but it is pretty straightforward. Two Go types are defined to map to the employment type in the company - `Permanent` and `Contract` representing the permanent and contract employees respectively. A `calculateSalary` method is defined for both types which calculates the salary for an employee. The calculation is different for each type. The permanent employee has pension while the contract employee doesn't.

To calculate the total payout for all employees, an array of 3 employees are considered. The salary for each employee is calculated using a `for` loop but calling the same method. The type of the array is `SalaryCalculator` which is the name of the interface that has the method declaration. Hence the interface can refer to each of the concrete types and invoke their method.

Extending the above paradigm is quite easy for say a new category of employee like those who do remote work.