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
