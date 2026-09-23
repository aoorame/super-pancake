package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	flog, err := os.OpenFile(`server.log`, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer flog.Close()

	mylog := log.New(flog, `serv `, log.LstdFlags|log.Lshortfile)
	mylog.Println(`Start server`)

	srv := server.NewAPIServer(mylog)

	err = srv.Server.ListenAndServe()
	if err != nil {
		mylog.Fatal(err)
	}
	mylog.Println(`Finish server`)
}
