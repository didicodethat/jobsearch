package main;

import (
	"fmt"
	"github.com/a-h/templ"
	"jobsearch/views"
	"jobsearch/tables"
	"log"
	"net/http"
)

func main() {
	defer tables.GetDB().Close()
	fmt.Print("Done \n")

	jobs, err := tables.GetAllJobs()
	if err != nil {
		log.Fatal(err)
	}
	index := views.Index(jobs)
	http.Handle("/", templ.Handler(index))

	fmt.Println("Starting Server on 8942")
	http.ListenAndServe(":8942", nil)
}
