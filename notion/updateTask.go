package notion

import (
	"bytes"
	"net/http"
	"setup_daily_jobs/util"
	"strings"
	"time"
)

func UpdateTask(notionToken, databaseId string, task Task) {
	var (
		uri           = "https://api.notion.com/v1/pages/" + task.Id
		auth          = "Bearer " + notionToken
		contentType   = "application/json"
		notionVersion = "2022-06-28"
	)

	data := `{
		"properties": {
			"name": {
				"title": [
					{ "text": { "content": "TITLE" } }
				]
			},
			"start_date": {
				"date": { "start": "START_DATE" }
			},
			"status": {
				"select": {"name": "今日の作業"}
			}
		}
	}`

	today := time.Now().Format("2006-01-02")
	replaced := strings.Replace(data, "START_DATE", today, 1)

	var taskName string
	if task.Properties.Template.Checkbox {
		taskName = "[" + today + "] " + task.Properties.Name.Title[0].PlainText
	} else {
		taskName = task.Properties.Name.Title[0].PlainText
	}

	replaced = strings.Replace(replaced, "TITLE", taskName, 1)

	req, err := http.NewRequest("PATCH", uri, bytes.NewBuffer([]byte(replaced)))

	if err != nil {
		util.ErrLog(err)
	}

	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Authorization", auth)
	req.Header.Add("Notion-Version", notionVersion)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		util.ErrLog(err)
	}

	defer resp.Body.Close()
}
