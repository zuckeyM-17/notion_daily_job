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

	propertiesTemplate := `{
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
	propertiesTemplate = strings.Replace(propertiesTemplate, "START_DATE", today, 1)

	for _, task := range tasks {
		var taskName string
		if task.Properties.Template.Checkbox {
			taskName = "[" + today + "] " + task.Properties.Name.Title[0].PlainText
		} else {
			taskName = task.Properties.Name.Title[0].PlainText
		}

		properties := strings.Replace(propertiesTemplate, "TITLE", taskName, 1)
		notion.UpdateTask(notionToken, databaseId, task, properties)
	}
}
