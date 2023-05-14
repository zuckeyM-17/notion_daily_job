package main

import (
	"os"

	"github.com/zuckeyM-17/setup_daily_jobs/notion"
)

func main() {
	// fmt.Println("Hello World")
	// fmt.Println("NOTION_API_TOKEN: ", os.Getenv("NOTION_API_TOKEN"))
	// fmt.Println("DAILY_TASK_DATABASE_ID: ", os.Getenv("DAILY_TASK_DATABASE_ID"))

	var (
		notionToken = os.Getenv("NOTION_API_TOKEN")
		databaseId  = os.Getenv("DAILY_TASK_DATABASE_ID")
	)

	notion.GetTasks(notionToken, databaseId)
}
