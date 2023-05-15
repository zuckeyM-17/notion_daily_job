package notion

import (
	"bytes"
	"net/http"
	"setup_daily_jobs/util"
	"strings"
	"time"
)

func UpdateTask(notionToken, databaseId, pageId string) {
	var (
		uri           = "https://api.notion.com/v1/pages/" + pageId
		auth          = "Bearer " + notionToken
		contentType   = "application/json"
		notionVersion = "2022-06-28"
	)

	data := `{
		"properties": {
			"start_date": {
				"date": { "start": "START_DATE" }
			},
			"status": {
				"select": {"name": "今日の作業"}
			}
		}
	}`

	replaced := strings.Replace(data, "START_DATE", time.Now().Format("2006-01-02"), 1)

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
