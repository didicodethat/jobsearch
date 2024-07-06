package routes;

import (
	"net/http"
	"log"
	"strconv"
	"jobsearch/db"
	"jobsearch/views"
	"html/template"
	"github.com/pkg/errors"
	"io"
)

func ErrorResponse(w io.Writer, err error) {
	tmp, devErr := template.New("tmp").Parse(`<div class="error">{{.}}</div>`)
	if devErr != nil {
		log.Fatal(errors.Wrap(devErr, "Invalid Template, dev, please fix"))
	}
	devErr = tmp.Execute(w, err.Error())
	if devErr != nil {
		log.Fatal(errors.Wrap(devErr, "Invalid Template, dev, please fix"))
	}
}

func ToggleJobDone(res http.ResponseWriter, req *http.Request) {
	log.Println("ToggleJobDone route started")
	defer log.Println("ToggleJobDone route finished")
	id, err := strconv.ParseUint(req.PathValue("id"), 10, 64)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		ErrorResponse(res, errors.Wrap(err, "Invalid id"))
		return
	}
	err = db.ToggleJobDone(id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		ErrorResponse(res, errors.Wrap(err, "Couldn't Toggle Job done status"))
		return
	}
	res.Write([]byte("Toggled Successfully"))
}


func CreateJobPosition(res http.ResponseWriter, req *http.Request) {
	log.Println("CreateJobPosition route started")
	defer log.Println("CreateJobPosition route finished")
	job, err := db.JobFromRequest(req)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		ErrorResponse(res, err)
		return
	}
	err = job.Save()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		ErrorResponse(res, err)
		return
	}
	views.JobItem(*job).Render(req.Context(), res)
}

func DeleteJobPosition(res http.ResponseWriter, req *http.Request) {
	log.Println("DeleteJobPosition route started")
	defer log.Println("DeleteJobPosition route finished")
	id, err := strconv.ParseUint(req.PathValue("id"), 10, 64)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		ErrorResponse(res, errors.Wrap(err, "Invalid id"))
		return
	}
	err = db.DeleteJob(uint(id))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		ErrorResponse(res, errors.Wrap(err, "Couldn't delete Job"))
		return
	}
}

func Init() {
	http.HandleFunc("GET /", func(res http.ResponseWriter, req *http.Request) {
		log.Println("GET /")
		jobs, err := db.GetAllJobs()
		if err != nil {
			log.Fatal(err)
		}
		views.Index(jobs).Render(req.Context(), res)
	})
	http.HandleFunc("PATCH /job/{id}/done", ToggleJobDone)
	http.HandleFunc("DELETE /job/{id}", DeleteJobPosition)
	http.HandleFunc("POST /jobs", CreateJobPosition)
	
	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	log.Println("Starting Server on 8942")
	http.ListenAndServe(":8942", nil)
}
