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

	query := `{
    "filter": {
			"and": [
				{ "property": "finish", "checkbox": { "equals": true } },
				{ "property": "done_date", "date": { "after": "START_DATE" } },
				{ "property": "done_date", "date": { "before": "END_DATE" } }
			]
    }
	}`

	startDate := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	endDate := time.Now().Format("2006-01-02")
	query = strings.Replace(query, "START_DATE", startDate, 1)
	query = strings.Replace(query, "END_DATE", endDate, 1)

	tasks, _ := notion.GetTasks(notionToken, databaseId, query)

	for _, task := range tasks {
		fmt.Println(task.Properties.Name.Title[0].PlainText)
	}

	properties := `{
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
				"select": {"name": "今日の作業"}
			},
			"category": {
				"select": { "name": "Skill" }
			}
		}
	}`

	properties = strings.Replace(properties, "DATABASE_ID", databaseId, 1)
	today := time.Now().Format("2006-01-02")
	properties = strings.Replace(properties, "START_DATE", today, 1)
	properties = strings.Replace(properties, "TITLE", "["+startDate+" - "+endDate+"]週次振り返り", 1)

	notion.CreateTask(notionToken, properties)
}
