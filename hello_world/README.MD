#### How do we run a Go program?

```
    go run <scriptname.go>
```

#### What other Go commands do we have available?

- **go build**: Accepts one or more Go source code files as arguments and compiles them
- **go run**: Accepts one or more Go source code files as arguments. It compiles AND executes them
- **go fmt**: Formats all the code inside all the different files given as arguments
- **go install**: Compiles and install a package
- **go get**: Downloads the raw source code from someone else's package
- **go test**: is used to run tests associated with the current project

#### What is a package and how does it work?

A Go package is a project or a workspace. A collection of different source codes all related to each other. <br>
Every file inside a project must have a "package name_of_the_package>" statement on top.<br>
There are 2 types of packages in Go:

- Executable: they generate a (complied) file that we can run
- Reusable: they are libraries/code dependencies with helper functions and reusable code
  How do we know if we are creating an executable or reusable package? <br>
  The name of the package you use for the package. "package main" is used to create executable package. Any other name makes a reusable/helper package. <br>
  Every time we create an executable package, we also need an executable function that must be called "main".

#### What are import statements?

Import statements are used to give access to other packages to our current package. <br>
The fmt library is used to print out and format info. Very much used for debugging. <br>
Important to get used to use the Go standard libray.

#### Func statement and Go script structure

Func is short for function (Structure is similar to the JS one). <br>
The structure of a Go script is always the same:

- Package declaration
- Import statements
- Functions
