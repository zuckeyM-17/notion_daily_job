package main

import (
	"os"

	"setup_daily_jobs/notion"
)

func main() {
	var (
		notionToken = os.Getenv("NOTION_API_TOKEN")
		databaseId  = os.Getenv("DAILY_TASK_DATABASE_ID")
	)

	tasks, _ := notion.GetTasks(notionToken, databaseId)

	for _, task := range tasks {
		notion.UpdateTask(notionToken, databaseId, task.Id)
	}
}
