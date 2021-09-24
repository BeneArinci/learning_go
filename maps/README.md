## Maps

Maps are similar to JS objects as they are data types that include n elements made of key-value pairs.<br>

First peculiarity of a Map data type is that it is **statically typed**. Keys must be all the same type and values the same. Not mandatory that keys and values should be the same type though. <br>

### Declaring a map

```
    colors := map[string]string{
            "red":  "#ff0000",
            "blue": "#00ff00",
        }
```

```
    var colors map[string]string
```

```
    colors := make(map[string]string)
```

In the last 2 examples, we create an empty map. We often do that in Go. We would then populate it later on.

### Adding elements to a map

```
    colors["white"] = "#ffffff"
```

Adding a key-value pair to a map must be done using the square brakets indicating the new key + its assigned value.

### Delete elements from a map

```
    delete(colors, "white")
```

To delete an element from a map we use the delete built in function. The first value we pass to the delete function is the map and the second is the key that we want to delete.
