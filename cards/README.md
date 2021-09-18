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

## Functions

When we have a return statement, we need to specify it in the func declaration, just after the brackets.

```
func newCard() string {
	return "Five of Diamonds"
}
```
