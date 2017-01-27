package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	logrus "github.com/Sirupsen/logrus"
)

func main() {
	err := powergrid()
	if err != nil {
		logrus.Fatal(err)
	}
}

func powergrid() error {
	go cli()
	return server()
}

func cli() error {
	defer fmt.Println("CLI terminated.")
	fmt.Println("Welcome to PowerGrid.")
	var input string
	fmt.Scanf("%s", &input)

	fmt.Printf("You entered %s.\n", input)
	return nil
}

func server() error {
	server := &http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rendered, err := json.Marshal(struct {
			PowerGrid string
		}{
			PowerGrid: "things",
		})
		if err != nil {
			panic(err)
		}

		type Template struct {
			Data string
		}
		data := Template{
			Data: string(rendered),
		}
		tmpl, err := template.New("").Parse(`<html><script> data={{.Data}};
		console.log(data);</script></html>`)
		if err != nil {
			panic(err)
		}
		tmpl.Execute(w, data)
	})
	return server.ListenAndServe()
}
