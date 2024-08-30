# learning-golang

## Defer
A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
```go
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

```



## Stacking defers
Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

```go
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

```


## Pointers
Go has pointers. A pointer holds the memory address of a value.

The type *T is a pointer to a T value. Its zero value is nil.

var p *int
The & operator generates a pointer to its operand.

i := 42
p = &i
The * operator denotes the pointer's underlying value.

fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
This is known as "dereferencing" or "indirecting".

Unlike C, Go has no pointer arithmetic.




## Structs
A struct is a collection of fields.

```go


type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
```

## Struct Fields
Struct fields are accessed using a dot.

```go
ype Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

```


## Pointers to structs
Struct fields can be accessed through a struct pointer.

To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.

```go
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

```


## Struct Literals
A struct literal denotes a newly allocated struct value by listing the values of its fields.

You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)

The special prefix & returns a pointer to the struct value.

```go
type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
```


Arrays
The type [n]T is an array of n values of type T.

The expression

var a [10]int
declares a variable a as an array of ten integers.

An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.

```go
func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
```


## Slices
An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type []T is a slice with elements of type T.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:
```go
a[low : high]
```
This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of a:

a[1:4]

## Slices are like references to arrays
A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes.

```go
func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
```

## Slice literals
A slice literal is like an array literal without the length.

This is an array literal:

[3]bool{true, true, false}
And this creates the same array as above, then builds a slice that references it:

[]bool{true, true, false}

```go
func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
```

## Slice defaults
When slicing, you may omit the high or low bounds to use their defaults instead.

The default is zero for the low bound and the length of the slice for the high bound.

For the array

var a [10]int
these slice expressions are equivalent:

a[0:10]
a[:10]
a[0:]
a[:]

- Arrays are value type slice are ref type 


## Slice length and capacity
A slice has both a length and a capacity.

The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

You can extend a slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.

## Nil slices
The zero value of a slice is nil.

A nil slice has a length and capacity of 0 and has no underlying array.
```go
func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
```


## Creating a slice with make
Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

The make function allocates a zeroed array and returns a slice that refers to that array:

a := make([]int, 5)  // len(a)=5
To specify a capacity, pass a third argument to make:

b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4


## Slices of slices (2D slice)
Slices can contain any type, including other slices.
```go
board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
```

## Appending to a slice
It is common to append new elements to a slice, and so Go provides a built-in append function. The documentation of the built-in package describes append.
```go
func append(s []T, vs ...T) []T
```
The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.

The resulting value of append is a slice containing all the elements of the original slice plus the provided values.

If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

(To learn more about slices, read the Slices: usage and internals article.)


## Range
The range form of the for loop iterates over a slice or map.

When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.

## Range continued
You can skip the index or value by assigning to _.

for i, _ := range pow
for _, value := range pow
If you only want the index, you can omit the second variable.

for i := range pow


## Maps
A map maps keys to values.

The zero value of a map is nil. A nil map has no keys, nor can keys be added.

The make function returns a map of the given type, initialized and ready for use.

## Map literals
Map literals are like struct literals, but the keys are required.

## Mutating Maps
Insert or update an element in map m:

m[key] = elem
Retrieve an element:

elem = m[key]
Delete an element:

delete(m, key)
Test that a key is present with a two-value assignment:

elem, ok = m[key]
If key is in m, ok is true. If not, ok is false.

If key is not in the map, then elem is the zero value for the map's element type.

Note: If elem or ok have not yet been declared you could use a short declaration form:

elem, ok := m[key]



## Function values
Functions are values too. They can be passed around just like other values.

Function values may be used as function arguments and return values.

## Function closures
Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

For example, the adder function returns a closure. Each closure is bound to its own sum variable.

## Methods
Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.

In this example, the Abs method has a receiver of type Vertex named v.
```go
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```

## Methods are functions
Remember: a method is just a function with a receiver argument.

Here's Abs written as a regular function with no change in functionality.
```go
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
```

## Methods continued
You can declare a method on non-struct types, too.

In this example we see a numeric type MyFloat with an Abs method.

You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).

## Methods and pointer indirection 

Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:

var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK
while methods with pointer receivers take either a value or a pointer as the receiver when they are called:

var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
For the statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.

The equivalent thing happens in the reverse direction.

Functions that take a value argument must take a value of that specific type:

var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!
while methods with value receivers take either a value or a pointer as the receiver when they are called:

var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
In this case, the method call p.Abs() is interpreted as (*p).Abs().


Summary
- Methods in Go are functions with a receiver argument, which can be either a value or a pointer.
- Methods provide a way to associate functions with types, enhancing Go’s ability to model real-world concepts.
- Receivers can be value or pointer types, and the choice affects whether the method can modify the receiver.
- Interfaces allow Go to define and use types that satisfy a particular set of behaviors, making Go a powerful language for building modular and reusable code.
- Methods are integral to Go’s approach to object-oriented programming, focusing on composition over inheritance.

```go
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
```

## Choosing a value or pointer receiver
There are two reasons to use a pointer receiver.

1. The first is so that the method can modify the value that its receiver points to.

2. The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.


## Interfaces
An interface type is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.

```go

// Defining an interface
type Speaker interface {
	Speak() string
}

type Dog struct{}

// Implementing the Speak method for Dog
func (d Dog) Speak() string {
	return "woof!"
}

func main() {
	var s Speaker
	s = Dog{}
	//d := Dog{}
	fmt.Println(s.Speak())
}
```

## Interfaces are implemented implicitly
A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

## Interface values
Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

(value, type)
An interface value holds a value of a specific underlying concrete type.

Calling a method on an interface value executes the method of the same name on its underlying type.

```go
package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

## Nil interface values
A nil interface value holds neither value nor concrete type.

Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.

## The empty interface
The interface type that specifies zero methods is known as the empty interface:

interface{}
An empty interface may hold values of any type. (Every type implements at least zero methods.)

Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.

## Type assertions
In Go, type assertions are used to extract the underlying value of an interface variable. Type assertions allow you to check if an interface value holds a specific type and to convert the interface value to that type. They provide a way to work with interface values more flexibly and dynamically.

The syntax for a type assertion is:
```
value := interfaceValue.(ConcreteType)
```
Example of Type Assertion
```go
package main

import "fmt"

func main() {
    var i interface{} = "hello"

    // Type assertion without checking
    s := i.(string)
    fmt.Println(s) // Output: hello

    // Type assertion with checking
    s, ok := i.(string)
    if ok {
        fmt.Println("String value:", s) // Output: String value: hello
    } else {
        fmt.Println("Not a string")
    }

    // Type assertion with a wrong type
    f, ok := i.(float64)
    if !ok {
        fmt.Println("Not a float64, value:", f) // Output: Not a float64, value: 0
    }
}
```

## What is a Type Switch?
A type switch is a construct that uses the switch keyword to check the dynamic type of an interface value. It allows you to execute different blocks of code depending on the type of the value stored in the interface. The type switch operates using the .() type assertion syntax within the switch statement.
```go
value, ok := i.(string)
if ok {
    fmt.Println("String value:", value)
}
```

## Stringers/ The string() method
One of the most ubiquitous interfaces is Stringer defined by the fmt package.

type Stringer interface {
    String() string
}

A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
For printing any struct crerate a method with name String() rith stuct as a reciver 

## The error Interface/ Error handling in go
The built-in error interface is the foundation for error handling in Go. It's defined in the builtin package as:
```go
type error interface {
    Error() string
}
```
Here is example:

```go
package main

import (
    "errors"
    "fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err) // Output: Error: cannot divide by zero
        return
    }
    fmt.Println("Result:", result)
}
```


- In Go, errors are treated as first-class values, allowing them to be passed and manipulated like any other value.
- The error interface, with its single Error() string method, is the cornerstone of error handling.
- Go provides simple ways to create, return, wrap, and unwrap errors using the errors package and fmt.Errorf.
- Using custom error types allows you to attach additional context and handle specific error cases effectively.
- Always check and handle errors immediately, use descriptive error messages, and leverage error wrapping and unwrapping for better traceability and robustness.


## The io.Reader Interface
In Go, readers are used to read data from a variety of input sources, such as files, network connections, buffers, or even generated data. The concept of a reader is embodied by the io.Reader interface, which is a fundamental part of Go's io package. The io.Reader interface is simple yet powerful, allowing for consistent and flexible input handling across different data sources.

```go
package io

type Reader interface {
    Read(p []byte) (n int, err error)
}
```
file open and read
```go
func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    buf := make([]byte, 1024)
    n, err := file.Read(buf)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    fmt.Printf("Read %d bytes: %s\n", n, string(buf[:n]))
}
```

Reading JSON, XML, or Other Formats

Go's encoding/json, encoding/xml, and other packages work seamlessly with io.Reader to decode data from various formats.

Example: Reading JSON from an io.Reader:

```go
package main

import (
    "encoding/json"
    "fmt"
    "strings"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    data := `{"name": "Alice", "age": 30}`
    reader := strings.NewReader(data)

    var person Person
    err := json.NewDecoder(reader).Decode(&person)
    if err != nil {
        fmt.Println("Error decoding JSON:", err)
        return
    }

    fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age) // Output: Name: Alice, Age: 30
}
```

Summary
- The io.Reader interface is central to Go's input handling, allowing for consistent reading from various sources.
- It is defined with a single Read method that reads data into a slice of bytes.
- Common implementations of io.Reader include os.File, bytes.Buffer, strings.Reader, and net.Conn.
- Utility functions like io.Copy, io.TeeReader, and packages like bufio enhance the functionality and efficiency of reading operations.
- Custom readers can be created by implementing the Read method, allowing you to define how and from where data is read.
- Readers integrate seamlessly with other standard library packages for handling JSON, XML, and other formats, making Go a powerful language for building data-driven applications.

## Images
In Go, image processing is primarily handled by the image package and its sub-packages in the standard library. These packages provide a rich set of tools for working with images, including creating, manipulating, and encoding/decoding various image formats.

Resizing and Transforming Images

To resize or perform more complex transformations, you may need third-party packages such as golang.org/x/image or github.com/nfnt/resize.

Example: Resizing an Image with github.com/nfnt/resize

```go
package main

import (
    "image"
    "image/jpeg"
    "os"

    "github.com/nfnt/resize"
)

func main() {
    file, err := os.Open("input.jpg")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    img, _, err := image.Decode(file)
    if err != nil {
        panic(err)
    }

    // Resize the image to 100x100 pixels using Lanczos resampling
    resizedImg := resize.Resize(100, 100, img, resize.Lanczos3)

    out, err := os.Create("resized.jpg")
    if err != nil {
        panic(err)
    }
    defer out.Close()

    jpeg.Encode(out, resizedImg, nil)
}
```

- The image package provides core functionality for working with images in Go, including basic types like image.Image and functions to create and manipulate images.
- Use sub-packages like image/png, image/jpeg, and image/gif for encoding and decoding common image formats.
- The image/color package defines color models and types for handling colors.
- For more complex operations like resizing or transformations, third-party libraries are often used.
- The image/draw package provides basic drawing operations, allowing for image manipulation at the pixel level.

## Type parameters
Type parameters, also known as generics, were introduced in Go 1.18 to provide a way to write functions and types that can operate on any data type. 
This feature allows you to create more flexible and reusable code by defining functions, methods, and types that can work with different types while maintaining type safety.

Basic Syntax:
```go
func FunctionName[T any](param T) {
    // Function implementation
}
```
