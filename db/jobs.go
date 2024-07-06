package db

import (
	"database/sql"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

type Job struct {
	Id             uint `db:"id"`
	Name           string `db:"name"`
	Url            string `db:"url"`
	DescriptionUrl sql.NullString `db:"description_url"`
	Tags           sql.NullString `db:"tags"`
	Done           bool `db:"done"`
}

func GetAllJobs() ([]Job, error) {
	jobs := []Job{}
	err := db.Select(&jobs, `SELECT id, name, url, description_url, tags, done from jobs ORDER BY done ASC, created_at DESC`) 
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't get jobs")
	}
	return jobs, nil
}


func ToggleJobDone(id uint64) error {
	db.MustExec("UPDATE jobs SET done = not done WHERE id=$1", id)
	return nil
}

func JobFromRequest(req *http.Request) (*Job, error) {
	job := &Job {}
	name := req.PostFormValue("name")
	if name == "" {
		return nil, errors.New("Name must not be empty")
	}
	url := req.PostFormValue("url")
	if url == "" {
		return nil, errors.New("URL must not be empty")
	}
	job.Name = name
	job.Url = url
	job.Done = false
	tags := req.PostFormValue("tags")
	job.Tags = sql.NullString{String: tags, Valid: tags != ""}
	descriptionURL := req.PostFormValue("description_url")
	job.DescriptionUrl = sql.NullString{String: descriptionURL, Valid: descriptionURL != ""}
	return job, nil
}

func (job *Job) Insert() error {
	db := GetDB()
	// Create
	_, err := db.Exec(`
		INSERT INTO jobs (name, url, description_url, tags) 
		VALUES ($1, $2, $3, $4);
	`, job.Name, job.Url, job.DescriptionUrl, job.Tags)
	if err != nil {
		return errors.Wrap(err, "Couldn't create new job")
	}
	query, err := db.Query("SELECT currval('jobs_id_seq')")
	if err != nil {
		return errors.Wrap(err, "Couldn't query for current job Id")
	}
	query.Next()
	var lastId uint
	query.Scan(&lastId)
	job.Id = uint(lastId)
	return nil
}

func (job *Job) Save() error {
	if job.Id == 0 {
		return job.Insert()
	} 
	// Update
	return nil
}

func DeleteJob(id uint) error {
	db := GetDB()
	_, err := db.Exec("DELETE FROM jobs WHERE id=$1", id)
	return err
}
