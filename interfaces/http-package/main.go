package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

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

	lw := logWriter{}
	io.Copy(lw, resp.Body)

	// // io.Copy func takes in an arg that implements the Writer interface (has a func called write)
	// // and another one that implements the Reader interface (respond to a method read) and prints out
	// // The following func prints out to stdout the response body
	// --> io.Copy(os.Stdout, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
