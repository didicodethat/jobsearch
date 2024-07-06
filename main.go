package main

import (
	"fmt"
	"strconv"
	//"github.com/a-h/templ"
	"jobsearch/db"
	"jobsearch/views"
	"log"
	"net/http"
)

func main() {
	defer db.GetDB().Close()
	fmt.Print("Done \n")

	http.HandleFunc("GET /", func(res http.ResponseWriter, req *http.Request) {
		jobs, err := db.GetAllJobs()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(jobs)
		views.Index(jobs).Render(req.Context(), res)
	})
	http.HandleFunc("PATCH /job/{id}", func(res http.ResponseWriter, req *http.Request) {
		log.Println("/job/{id}")
		id, err := strconv.ParseUint(req.PathValue("id"), 10, 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("<div class='error'>id must be a valid integer</div>"))
			return
		}
		req.ParseForm()
		form := req.Form
		if form.Has("done") {
			err = db.ToggleJobDone(id)
			if db.ToggleJobDone(id) != nil {
				res.WriteHeader(http.StatusInternalServerError)
				res.Write([]byte(err.Error()))
				return
			}
			res.Write([]byte("Toggled Successfully"))
		}
	})
	//http.Handle("/", templ.Handler(index))
	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("GET /static/", http.StripPrefix("/static/", fileServer))
	fmt.Println("Starting Server on 8942")
	http.ListenAndServe(":8942", nil)
}
