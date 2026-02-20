package main

// interfaces are a powerful feature in Go that allow you to define a set of methods that a type must implement. 
// An interface is a collection of method signatures that a type can implement. 
// A type that implements all the methods of an interface is said to satisfy that interface.

// The syntax for defining an interface is as follows:
// type InterfaceName interface {
//     Method1(parameters) returnType
//     Method2(parameters) returnType
//     // ...
// }

// Here is an example of an interface and a type that implements it:

type Greeter interface {
	Greet() string
}

type Person struct {
	Name string
}

func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

func main() {
	var greeter Greeter = Person{Name: "Alice"}
	greeting := greeter.Greet()
	println(greeting) // Output: Hello, my name is Alice
}

// In this example, we defined an interface called Greeter with a single method Greet. 
// We then defined a type called Person that implements the Greet method. 
// Finally, we created a variable of type Greeter and assigned it an instance of Person. 
// We called the Greet method on the greeter variable, which returned the greeting message.


// empty interface example

type Any interface{}

func emptyInterface() {
	var a Any = "Hello, World!"
	println(a.(string)) // Output: Hello, World!
}

// In this example, we defined an empty interface called Any. 
// An empty interface can hold values of any type. 
// We assigned a string value to the variable a of type Any and then used type assertion to retrieve the string value.





// embedding interfaces example

type Speaker interface {
	Speak() string
}

type Human struct {
	Name string
}

func (h Human) Speak() string {
	return "Hello, my name is " + h.Name
}

type Employee struct {
	Human // embedding the Human struct
	Position string
}



// type assertion example

func typeAssertion() {
	var greeter Greeter = Person{Name: "Bob"}
	if p, ok := greeter.(Person); ok {
		println("The person's name is " + p.Name) // Output: The person's name is Bob
	} else {
		println("greeter does not implement Person")
	}
}

// In this example, we used type assertion to check if the greeter variable of type Greeter is actually of type Person. 
// If the assertion is successful, we can access the Name field of the Person struct. 
// If the assertion fails, we handle the case where greeter does not implement Person.







// type switch example

func typeSwitch(i interface{}) {
	switch v := i.(type) {
	case string:
		println("i is a string: " + v)
	case int:
		println("i is an int: ", v)
	default:
		println("i is of unknown type")
	}
}

// In this example, we defined a function called typeSwitch that takes an empty interface as a parameter. 
// We used a type switch to determine the actual type of the value passed to the function and handle it accordingly. 
// The type switch allows us to perform different actions based on the type of the value.