package main

import (
	"os"
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

	tasks, _ := notion.GetTasks(notionToken, databaseId)

	for _, task := range tasks {
		notion.UpdateTask(notionToken, databaseId, task)
	}
}
