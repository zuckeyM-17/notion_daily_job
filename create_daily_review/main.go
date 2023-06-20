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

	today := time.Now().Format("2006-01-02")

	query := `{
    "filter": {
			"and": [
				{ "property": "finish", "checkbox": { "equals": true } },
				{ "property": "done_date", "date": { "equals": "DONE_DATE" } }
			]
    }
	}`

	query = strings.Replace(query, "DONE_DATE", today, 1)
	tasks, _ := notion.GetTasks(notionToken, databaseId, query)

	propertiesTemplate := `{
		"parent": { "database_id": "DATABASE_ID" },
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
				"select": { "name": "今日の作業" }
			},
			"category": {
				"select": { "name": "Skill" }
			}
		},
		"children":	[
			CHILDREN
		]
	}`

	content := `{
		"object": "block",
		"type": "paragraph",
		"paragraph": {
			"rich_text": [
				{
					"type": "text",
					"text": {
						"content": "TEXT"
					}
				}
			]
		}
	}`

	var contents string

	for i, task := range tasks {
		c := strings.Replace(content, "TEXT", task.Properties.Name.Title[0].PlainText, 1)
		contents = contents + c
		if i != len(tasks)-1 {
			contents = contents + ","
		}
	}

	properties := strings.Replace(propertiesTemplate, "DATABASE_ID", databaseId, 1)
	properties = strings.Replace(properties, "START_DATE", today, 1)
	properties = strings.Replace(properties, "TITLE", "["+today+"] 日次振り返り", 1)
	properties = strings.Replace(properties, "CHILDREN", contents, 1)

	notion.CreateTask(notionToken, properties)
}
