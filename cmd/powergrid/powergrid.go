package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/hinshun/powergrid"
	"github.com/hinshun/powergrid/stdgame"

	logrus "github.com/Sirupsen/logrus"
)

func main() {
	err := startPowergrid()
	if err != nil {
		logrus.Fatal(err)
	}
}

func startPowergrid() error {
	fmt.Println("Welcome to PowerGrid. Choose a number of players:")
	var input string
	fmt.Scanf("%s", &input)
	i, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		return fmt.Errorf("failed to parse input: %s", err)
	}
	game := stdgame.New(uint(i))
	go cli(game)
	return server(game)
}

func cli(game powergrid.Game) {
	game.Run()
}

func server(game powergrid.Game) error {
	server := &http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		type Template struct {
			Data string
		}
		data := Template{
			Data: game.GridInfo(),
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
