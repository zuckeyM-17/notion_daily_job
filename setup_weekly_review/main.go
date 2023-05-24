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

	query = strings.Replace(query, "START_DATE", time.Now().AddDate(0, 0, -7).Format("2006-01-02"), 1)
	query = strings.Replace(query, "END_DATE", time.Now().Format("2006-01-02"), 1)

	tasks, _ := notion.GetTasks(notionToken, databaseId, query)

	for _, task := range tasks {
		fmt.Println(task.Properties.Name.Title[0].PlainText)
	}
}
