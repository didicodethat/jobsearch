package routes;

import (
	"net/http"
	"log"
	"fmt"
	"strconv"
	"jobsearch/db"
	"jobsearch/views"
)

func Init() {
	http.HandleFunc("GET /", func(res http.ResponseWriter, req *http.Request) {
		log.Println("GET /")
		jobs, err := db.GetAllJobs()
		if err != nil {
			log.Fatal(err)
		}
		views.Index(jobs).Render(req.Context(), res)
	})
	http.HandleFunc("PATCH /job/{id}/done", func(res http.ResponseWriter, req *http.Request) {
		log.Println("PATCH /job/{id}/done started")
		defer log.Println("PATCH /job/{id}/done finished")
		id, err := strconv.ParseUint(req.PathValue("id"), 10, 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("<div class='error'>id must be a valid integer</div>"))
			return
		}
		log.Println("parse done")
		err = db.ToggleJobDone(id)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte(err.Error()))
			return
		}
		res.Write([]byte("Toggled Successfully"))
	})
	//http.Handle("/", templ.Handler(index))
	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("GET /static/", http.StripPrefix("/static/", fileServer))
	fmt.Println("Starting Server on 8942")
	http.ListenAndServe(":8942", nil)
}