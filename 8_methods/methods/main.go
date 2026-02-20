package main


// methods are functions with a receiver argument. 
// The receiver is the type that the method is associated with.
//  Methods can be defined on any type, including built-in types 
// and user-defined types.

// The syntax for defining a method is as follows:
// func (receiver type) methodName(parameters) returnType {
//     // method body
// }

// The receiver is specified in parentheses before the method name. 
// The receiver can be a value or a pointer, depending on 
// whether you want to modify the receiver or not.

// Here is an example of a method defined on a user-defined type:

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

func main() {
	person := Person{Name: "Alice", Age: 30}
	greeting := person.Greet()
	println(greeting) // Output: Hello, my name is Alice
}

// In this example, we defined a method 
// called Greet on the Person type. 
// The Greet method takes a Person receiver and 
// returns a greeting string. 
// We then created an instance of the Person type and 
// called the Greet method to get the greeting message.




// pointer receiver example

type Counter struct {
	Value int
}

func (c *Counter) Increment() {
	c.Value++
}

func pointer() {
	counter := &Counter{Value: 0}
	counter.Increment()
	println(counter.Value) // Output: 1
}



// value receiver example

type Point struct {
	X, Y int
}

func (p Point) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

func value() {
	point := Point{X: 0, Y: 0}
	point.Move(2, 3)
	println(point.X, point.Y) // Output: 0 0
}

// In the pointer receiver example, we defined a method called Increment on the Counter type. 
// The Increment method takes a pointer receiver, which allows us to modify the Value field of the Counter instance. 
// When we call the Increment method on the counter instance, it increments the Value field by 1.

// In the value receiver example, we defined a method called Move on the Point type. 
// The Move method takes a value receiver, which means that it operates on a copy of the Point instance. 
// When we call the Move method on the point instance, it does not modify the original Point instance, and the output remains (0, 0).