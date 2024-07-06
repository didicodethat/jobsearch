package db;

import(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
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
	err := db.Select(&jobs, `SELECT id, name, url, description_url, tags, done from jobs`) 
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func ToggleJobDone(id uint64) error {
	_, err := db.Exec("UPDATE jobs SET done = not done WHERE id=$1",id)
	return err
}
