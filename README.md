# GoLang101


Go (often referred to as Golang) is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.

Here is an hands-on introduction to Go using example programs.
Run them here: https://play.golang.org/
Run instruction: run program.go

Team: Pranav Kumar Sivakumar, Sowmya Ramakrishnan

##### hello world.go

Our first program will print the classic “hello world” message.

##### var-values-const.go

  - Go has various value types including strings, integers, floats, booleans, etc. 
  - Variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.
  - The := syntax is shorthand for declaring and initializing a variable
  - Go supports constants of character, string, boolean, and numeric values.
##### for.go

for is Go’s only looping construct.

#####  if-else-switch.go

  - Branching with if and else in Go is straight-forward.
  - Note that you don’t need parentheses around conditions in Go, but that the braces are required.
  - There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.
  - Switch statements express conditionals across many branches.

##### arrays.go

In Go, an array is a numbered sequence of elements of a specific length. The builtin len returns the length of an array.
Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.

##### slices.go

  - You’ll see slices much more often than arrays in typical Go. Slices are a key data type, giving a more powerful interface to sequences than array.
  - Unlike arrays, slices are typed only by the elements they contain (not the number of elements). To create an empty slice with non-zero length, use the builtin make.
  - In addition to these basic operations, slices support several more that make them richer than arrays. One is the builtin append, which returns a slice containing one or more new values. Note that we need to accept a return value from append as we may get a new slice value.

##### maps.go

Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).

##### range.go

range iterates over elements in a variety of data structures.

##### functions.go

Functions are central in Go. Go requires explicit returns. Call a function just as you’d expect, with name(args).
Go has built-in support for multiple return values. This feature is used often in idiomatic Go, for example to return both result and error values from a function.

##### variadic_functions.go

Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.

##### closures.go

Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it.

##### pointers.go

Go supports pointers, allowing you to pass references to values and records within your program.

##### structs.go

Go’s structs are typed collections of fields. They’re useful for grouping data together to form records. Structs are mutable. Structs hold only state and no behavior.

##### methods.go

Methods are functions that operate on particular types. They have a receiver clause that mandates what type they operate on. Here are some methods for our example.

##### interfaces.go

Interfaces are types that declare sets of methods. Similarly to interfaces in other languages, they have no implementation. 
But Interfaces in Go are very different from, say, Java interfaces. You don’t explicitly say a data type implements an interface; rather, your data types must implement all of the methods that the interface defines, and the compiler checks to see if assignments to variables of the interface type are valid.

##### struct_embedding.go

The theory behind embedding is pretty straightforward: by including a type as a nameless parameter within another type, the exported parameters and methods defined on the embedded type are accessible through the embedding type. 
So that we don't need to re-implement the functionalities of the embedded types for the embedding type.
