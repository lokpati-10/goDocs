package main

import "fmt"

/**
Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.

In this example, the getName method has a receiver of type Vertex named v.

Remember: a method is just a function with a receiver argument.

You can only declare a method with a receiver whose type is defined in the same package as the method.
You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
 */

type IStudent interface {
	getStudentName() string
	getStandard() int
	getStudentSection() string
	updateStudentLastName(lastName string)
    printStudentDetails()
}

type Student struct {
	firstName string
	lastName string
	standard int
	section string
}

func (s Student) getStudentName() string {
	return s.firstName + " " + s.lastName
}

func (s Student) getStandard() int {
	return s.standard
}

func (s Student) getStudentSection() string {
	return s.section
}

// pointer receivers is always used when you need to update the receivers value
func (s *Student) updateStudentLastName(lastName string) {
	if s == nil {
		fmt.Println("student can not be nil")
		return
	}
	s.lastName = lastName
}

func (s Student) printStudentDetails() {
	fmt.Println("studentName", s.getStudentName(), " standard:", s.getStandard(), " section:", s.getStudentSection())
}


func describe(i IStudent) {
	if i == nil {
		// nil interface values neither hold value nor concrete type
		return
	}
	fmt.Printf("type = %T value = %v\n", i, i)
}

func main(){

	// nil concept
	var xyz IStudent
	var t *Student

	xyz = t
	xyz.updateStudentLastName("mishra")

	// interface concept
	var student IStudent = &Student{
		firstName: "madhav",
		lastName: "dwivedi",
		standard: 12,
		section: "A",
	}

	student.printStudentDetails()

	student.updateStudentLastName("Dhar Dwivedi")

	student.printStudentDetails()
	describe(student)

}
