package notion

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/tidwall/gjson"

	"setup_daily_jobs/util"
)

type Select struct {
	Name string `json:"name"`
}

type Status struct {
	Select Select `json:"select"`
}

type Template struct {
	Checkbox bool `json:"checkbox"`
}

type Date struct {
	Start string `json:"start"`
}

type StartDate struct {
	Date string `json:"date"`
}

type Properties struct {
	Name      Name      `json:"name"`
	Status    Status    `json:"status"`
	Template  Template  `json:"template"`
	StartDate StartDate `json:"start_date"`
}

type Title struct {
	PlainText string `json:"plain_text"`
}
type Name struct {
	Title []Title `json:"title"`
}

type Task struct {
	Id         string     `json:"id"`
	Url        string     `json:"url"`
	Properties Properties `json:"properties"`
	UpdatedAt  string     `json:"last_edited_time"`
}

func GetTasks(notionToken, databaseId, query string) ([]Task, error) {
	var (
		uri           = "https://api.notion.com/v1/databases/" + databaseId + "/query"
		auth          = "Bearer " + notionToken
		contentType   = "application/json"
		notionVersion = "2022-06-28"
	)

	type SearchData struct{}

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer([]byte(query)))
	if err != nil {
		util.ErrLog(err)
		return []Task{}, err
	}

	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Authorization", auth)
	req.Header.Add("Notion-Version", notionVersion)

	client := &http.Client{}

	res, _ := client.Do(req)
	if err != nil {
		util.ErrLog(err)
		return []Task{}, err
	}
	defer res.Body.Close()

	r, _ := io.ReadAll(res.Body)
	jsonStr := string(r)

	var tasks []Task
	results := gjson.Get(jsonStr, "results").String()
	json.Unmarshal([]byte(results), &tasks)

	return tasks, nil
}
