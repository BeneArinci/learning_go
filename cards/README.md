## Variables

Go is a **static type** language. When you create a variable of a specific type, only values of that type can be assigned to that variable. <br>

Assignign a variable can be done in 2 ways:

```
    var card string = "Ace of Spades"
```

```
    card := "Ace of Spades
```

In this second example, the compiler will infer the type of the variable from the value we assign to it. <br>

If we then want to reassign a new value to the same variable we can just type:

```
    card = "Two of Diamonds"
```

We might be able to initialize a variable and then later assign a variable to it:

```
    var deckSize int
    deckSize = 52
```

We cannot assign a value to a variable outside a function and using it inside it but we can initialize the variable outside the function and assign it a value + use it inside the function.

### Basic Go types we are going to use

- bool
- string
- int
- float64
- array
- slice
- map

## Functions

When we have a return statement, we need to specify it in the func declaration, just after the brackets.

```
    func newCard() string {
        return "Five of Diamonds"
    }
```

## Arrays

Arrays and Slices are data types used for storing a list of elements.<br>
The main difference is that Arrays are fixed **length**. Slices are variable in terms of lenght/number of \*elements<br>
Both slices and Arrays must be declared and contain a specific data type. It can be a list of string as an example.

## For loop

```
    for i, card := range cards {
    }
```

The real key word for every for loop in Go is the **range** one. It indicates what is the slice we want to iterate through. Then, for every element, we can assign (:=) index and value to 2 specific variables (i.e.: i and card).

## Custom type declarations

Go is an object oriented language that works differently from the other OO ones.<br>
Common OOP languages allow to create classes that are blue prints containing methods and variables that will be related to any instance of that specific class. <br>
Go allows to declare **new types** but these new types extend from existing base Go types. The declaration works this way:

```
type deck []string
```

In the above case, the type deck extends the type slice of strings. <br>

To link functions to this new type we create, we will use what are called **receiver functions** that have the following structure:

```
    func (receiver newType) funcName() {
        for i, e := range oneLetterVariable {
            fmt.Println(i, e)
        }
    }
```

With the receiver statement ("(receiver newType)") we determine that every variable of type = newType will get access to funcName(). The "receiver" refers to any instance of the newType. It is generally named with one or 2 letters (by convention).
