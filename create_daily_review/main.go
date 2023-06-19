package main

import (
	"fmt"
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

	for _, task := range tasks {
		fmt.Println(task.Properties.Name.Title[0].PlainText)
	}

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
			{
				"object": "block",
				"type": "paragraph",
				"paragraph": {
					"rich_text": [
						{
							"type": "text",
							"text": {
								"content": "DONE_TASKS"
							}
						}
					]
				}
			}
		]
	}`

	properties := strings.Replace(propertiesTemplate, "DATABASE_ID", databaseId, 1)
	properties = strings.Replace(properties, "START_DATE", today, 1)
	properties = strings.Replace(properties, "TITLE", "["+today+"] 日次振り返り", 1)

	notion.CreateTask(notionToken, properties)
}
