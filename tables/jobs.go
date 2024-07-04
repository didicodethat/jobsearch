package tables;

import(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Job struct {
	Id             uint
	Name           string
	Url            string
	DescriptionUrl sql.NullString
	Tags           sql.NullString
	Done           bool
}

func GetAllJobs() ([]Job, error) {
	rows, err := db.Query(`select count(*) from jobs`)
	if err != nil {
		return nil, err
	}
	var jobCount int
	rows.Next()
	err = rows.Scan(&jobCount)
	if err != nil {
		return nil, err
	}
	jobs := make([]Job, jobCount)
	rows, err = db.Query(`select id, name, url, description_url, tags, done from jobs limit 100`)
	if err != nil {
		return nil, err
	}
	i := 0
	for rows.Next() {
		job := &jobs[i]
		err = rows.Scan(&job.Id, &job.Name, &job.Url, &job.DescriptionUrl, &job.Tags, &job.Done)
		if err != nil {
			return nil, err
		}
		i++
	}
	return jobs, nil
}
