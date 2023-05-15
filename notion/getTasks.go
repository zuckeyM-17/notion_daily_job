package notion

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/tidwall/gjson"

	"setup_daily_jobs/util"
)

type Task struct {
	Id         string
	Url        string
	Properties struct {
		Status struct {
			multi_select []struct {
				Name string
			}
		}
		Template struct {
			checkbox bool
		}
		start_date struct {
			date *string
		}
	}
	Name struct {
		Title []struct {
			PlainText string
		}
	}
}

func GetTasks(notionToken, databaseId string) ([]Task, error) {
	var (
		uri           = "https://api.notion.com/v1/databases/" + databaseId + "/query"
		auth          = "Bearer " + notionToken
		contentType   = "application/json"
		notionVersion = "2022-06-28"
	)

	type SearchData struct{}

	data := `{
    "filter": {
			"or": [
					{
						"property": "template",
						"checkbox": {
							"equals": true
						}
					},
					{
						"property": "start_date",
						"date": {
							"equals": "START_DATE"
						}
					},
					{
						"property": "finish",
						"checkbox": {
							"equals": false
					}
				}
			]
    }
	}`
	replaced := strings.Replace(data, "START_DATE", time.Now().Format("2006-01-02"), 1)

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer([]byte(replaced)))
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
