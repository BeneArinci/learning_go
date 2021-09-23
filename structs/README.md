## Structs

Struct is probably the most important Go data type. It is something somehow similar to a JS object. <br>

You can create a struct type this way:

```
    type person struct {
        name    string
        surname string
        age     int
    }
```

Assigning values to a struct, a.k.a. creating new instances of a struct types can be done in multiple ways. <br>

A first one relies on the order of the elements in the struct (it's not an optimal option though):

```
    mySelf := person{"Bene", "Arinci", 32}
```

Another and far better way to create an instance of the type struct is to specify the key. This mode is better because we can use the order we want. It looks like this:

```
    mySelf := person{name: "Bene", surname: "Arinci", age: 32}
```

A third and last way of declaring a struct is by initializing the variable first and adding values to it later. Important to consider is that when a struct type variable get initialized, everyone of its values is given a value of "ZERO" based on the its specific type (i.e. a string would be "", int = 0, bool=false etc.)

```
    var gianlu person
```

### Embedded struct

Interesting Go feature is the possibility to create and assign a struct into a struct:

```
    type contactInfo struct {
        email   string
        zipCode int
    }
    type person struct {
        name        string
        surname     string
        age         int
        contactInfo
    }

    mySelf := person{
		name:    "Bene",
		surname: "Arinci",
		age:     32,
		contactInfo: contactInfo{
			email:   "aaa@gmail.com",
			zipCode: 10101,
		},
	}
```

### Pointers in Go

Before speaking about pointers we should mention some important info about RAM and how it works. <br>

RAM is made of many slots (value containers) that store info/variables for us. Each slot has an address that points to it.<br>

When we create a variable with Go, it will go into our RAM and will look for a space in our RAM.<br>

Go is know as a **Pass by Value** language. This means that, every time we create a variable, Go will take that value and copy that in the memory. When we pass a value in a function as a receiver, Go is creating a copy of that value and stores it at another address. That copy is made availabe to the function only.<br>

```
|------------------------------------------------|
|0000|                                           |
|0001| person{name:"Bene", surname:"Arinci", ...}| <-- bene
|0002|                                           |
|0003| person{name:"Bene", surname:"Arinci", ...}| <-- p
|------------------------------------------------|
```

Since the function would be working on p, rather than bene (in the example above), should we use the function to make changes to bene (i.e. name property), we wouldn't be able to do it if it wasn't for pointers.<br>

Functions that aim to modify the value/values of a variable, should take the pointer to that variable as a receiver rather than the variable per se. How to?

```
    mySelf := person{
            name:    "Bene",
            surname: "Arinci",
            age:     32,
            contactInfo: contactInfo{
                email:   "aaa@gmail.com",
                zipCode: 10101,
            },
        }

    // this creates a pointer to the variable mySelf
    mySelfPointer := &mySelf

    // function takes in a pointer as receiver
    func (pointerToPerson *person) updateName(newFirstName string) {
	    (*pointerToPerson).name = newFirstName
    }

    // function is called on the pointer to the variable and not on the variable
    mySelfPointer.updateName("Whatever name")

```

#### Pointers operators

- "&variable": this is an operator that gives us access to the memory address my variable is sitting at (a.k.a. variable pointer);
- "\*variable": this is an operator that take an address and gives back the value stored at that address.

IMPORTANT: When you see a \* sign before a type is very different from seeing it before a pointer! <br>

If you set a \* just before a pointer you're aiming to get the value sitting at that address (i.e. inside the updateName func). <br>

If you use the \* before a type (as it happens when declaring the receiver of a func) you're actually just describing the type of the receiver that, in this case, would be a pointer to a type.

#### Pointers shortcut

When you define a func that takes in a pointer to a type, it can be called also on a variable that is of the original type.
