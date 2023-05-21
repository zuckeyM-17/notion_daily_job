package notion

import (
	"bytes"
	"net/http"
	"setup_daily_jobs/util"
)

func UpdateTask(notionToken, databaseId string, task Task, properties string) {
	var (
		uri           = "https://api.notion.com/v1/pages/" + task.Id
		auth          = "Bearer " + notionToken
		contentType   = "application/json"
		notionVersion = "2022-06-28"
	)

	req, err := http.NewRequest("PATCH", uri, bytes.NewBuffer([]byte(properties)))

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
