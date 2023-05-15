package main

import (
	"fmt"
	"os"

	"setup_daily_jobs/notion"
)

func main() {
	// fmt.Println("Hello World")
	// fmt.Println("NOTION_API_TOKEN: ", os.Getenv("NOTION_API_TOKEN"))
	// fmt.Println("DAILY_TASK_DATABASE_ID: ", os.Getenv("DAILY_TASK_DATABASE_ID"))

	var (
		notionToken = os.Getenv("NOTION_API_TOKEN")
		databaseId  = os.Getenv("DAILY_TASK_DATABASE_ID")
	)

	tasks, _ := notion.GetTasks(notionToken, databaseId)

	for _, task := range tasks {
		fmt.Println(task.Name.Title)
	}
}
