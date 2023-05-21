package main

import (
	"os"
	"setup_daily_jobs/notion"
	"strings"
	"time"
)

func init() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	time.Local = jst
}

func main() {
	var (
		notionToken = os.Getenv("NOTION_API_TOKEN")
		databaseId  = os.Getenv("DAILY_TASK_DATABASE_ID")
	)

	query := `{
    "filter": {
			"and": [
				{ "property": "finish", "checkbox": { "equals": true } },
				{ "property": "done_date", "date": { "is_empty": true } }
			]
    }
	}`

	tasks, _ := notion.GetTasks(notionToken, databaseId, query)

	propertiesTemplate := `{
		"properties": {
			"done_date": {
				"date": { "start": "DONE_DATE" }
			}
		}
	}`

	for _, task := range tasks {
		updatedAt, _ := time.Parse(time.RFC3339, task.UpdatedAt)
		properties := strings.Replace(propertiesTemplate, "DONE_DATE", updatedAt.Format("2006-01-02"), 1)
		notion.UpdateTask(notionToken, databaseId, task, properties)
	}
}
