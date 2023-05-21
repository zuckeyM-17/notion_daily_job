package main

import (
	"os"
	"strings"
	"time"

	"setup_daily_jobs/notion"
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
				{ "property": "finish", "checkbox": { "equals": false } },
				{
					"or": [
						{ "property": "start_date", "date": { "equals": "START_DATE" } },
						{ "property": "template", "checkbox": { "equals": true } }
					]
				}
			]
    }
	}`

	query = strings.Replace(query, "START_DATE", time.Now().Format("2006-01-02"), 1)

	tasks, _ := notion.GetTasks(notionToken, databaseId, query)

	for _, task := range tasks {
		notion.UpdateTask(notionToken, databaseId, task)
	}
}
