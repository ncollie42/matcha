package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) > 0 {
		fmt.Println(os.Args[1])
		fs := http.FileServer(http.Dir(os.Args[1]))
		// fs := http.FileServer(http.Dir("../../test2/client"))
		http.Handle("/", fs)

		log.Println("Listening on :3000...")
		err := http.ListenAndServe(":3000", nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}
