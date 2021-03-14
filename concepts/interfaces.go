package  main

import (
	"fmt"
	"time"
)

/**
The empty interface
The interface type that specifies zero methods is known as the empty interface
An empty interface may hold values of any type. (Every type implements at least zero methods.)
 */


func describeInterface(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func typeAssertion(){

	fmt.Println(" ==========type assertions=========")
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s , ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // panic
	// fmt.Println(f)
}

type Person struct {
	Name string
	surname string
}

func (p *Person) String() string {
	return fmt.Sprintf("%v %v", p.Name, p.surname)
}

type MyError struct {
	when time.Time
	what string
}

func (error *MyError) Error() string {
	return fmt.Sprintf("%v %v", error.when, error.what)
}

func runError() {
	fmt.Println("=======my Error implementation=======")
	var error = &MyError{
		time.Now(),
		"does not work man",
	}
	fmt.Println(error)
}

func stringer(){
	fmt.Println("==========stringer implementation=========")
	a := &Person{"Lokpati" , "mishra"}
	fmt.Println(a)
}

func main() {
	var emptyInterface interface{}
	describeInterface(emptyInterface)

	emptyInterface = "hello"
	describeInterface(emptyInterface)

	typeAssertion()

	stringer()

	runError()
}

