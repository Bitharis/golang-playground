package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World")

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Printf("Byttes received: %d\n", n)
	})

	http.ListenAndServe(":8080", nil)
}
