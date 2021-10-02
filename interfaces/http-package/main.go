package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.thecocktaildb.com/api/json/v1/1/search.php?s=margarita")
	if err != nil {
		println("Error: ", err)
		os.Exit(1)
	}

	// // The empty byte slice we create is quite big because
	// // the Read func doesn't have the ability to resize it
	// // and we do not want to end up with an error
	// --> bs := make([]byte, 99999)

	// // The Read func is applicable to any type of element.
	// // It takes whatever is coming from the receiver into the empty []byte
	// // that is passed as argument.
	// --> resp.Body.Read(bs)

	// --> fmt.Println(string(bs))

	io.Copy(os.Stdout, resp.Body)

}
