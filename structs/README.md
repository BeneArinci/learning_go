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
