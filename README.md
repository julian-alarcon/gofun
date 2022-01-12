# gofun

Learning the principles of golang

## Golang Topics managed

## hello-world

* Module -> Package -> Functions
* `main` package and `main` function as special entry points
* Print lines in Go
* Importing packages (`fmt`)
* Run/Build/Install go apps
* Go Extension for VS Code (helps on syntax and import automatically packages as functions are used)

## variables

* Global variables
* Local variables inside functions
* Pointers (best memory management as data is not duplicated)

## envvar

* Environment variables

## functions

* Set functions
* Use conversion from string to integers

## variadicfunc

* Variadic Functions (using `...` ellipsis symbol before the type of the values)
* Use return to return values
* Storage returned values from functions directly in the variable definition

## errorhandling

* Handling errors when returned from a function
* `err` type

## breakingtheloop

* Breaking the loop with reserved word `break`
* Use labels to break specific loops defining it before any for definition and adding `:`.

## slices

* When using slices, the values always change from the original slice as they are reference types
* Definition of slices with reserved function `make`
* `len` and `cap` functions to obtain length and capacity of any slice
* Usage of reserved word `range` to use an slice for a `for` loop
* Usage of reserved character `%d` to add multiple variable values to `Printf` output
* slices from slices
* using a slice with no value before or after `[:5]` ([] or `[3:]`) will respectively define the beginning or the end of a slice
* Adding a new element to a slice beyond the initial size definition will grow the capacity, and always double it (2,4,8,16,32,...)
* Adding a slice to another slice with `append(firstSlice, sliceToAdd...)`. Mind the `...` in the second slice

## maps

* When using maps, the values always change from the original map as they are reference types
* Definition with `make(map[indexType]valueType)`
* Definition adding values
* Printing maps inside loops (using `range`) will not print it in order
* Assign new values (new index and value)
* Deleting elements from a map with the function `delete()`
* Maps Not thread safe

## structs

* Defining structs with `type nameOfStruct struct {}` syntax
* Definition of a variable of Struct type
* Printing modifying specific sections of a variable defined as struct with the `.` operator

## concurrency

* concurrency is different to parallelism
* concurrency is execution multiple tasks but only 1 running
* goroutine encapsulate the different tasks
* anonymous functions (no name) inside other function and setting `()` at the end
* goroutine defined with reserved word `go`
* usage of `sync` module and reserved functions `.Add(#OfGoroutines)`, `.Done()`, `.Wait()`
* channels are used to communication between go routines
  * unbuffered: `myChannel := make(chan int)`
  * buffered: `myChannel := make(chan int, 5)`, just setting the size will create a buffered channel, in this case able to storage 5 integers
