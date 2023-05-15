package main

import (
	"fmt"
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
		fmt.Println(task.Properties.Name.Title[0].PlainText)
		notion.UpdateTask(notionToken, databaseId, task.Id)
	}
}
