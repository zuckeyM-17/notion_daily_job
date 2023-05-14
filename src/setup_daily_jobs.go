package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("NOTION_API_TOKEN: ", os.Getenv("NOTION_API_TOKEN"))
	fmt.Println("DAILY_TASK_DATABASE_ID: ", os.Getenv("DAILY_TASK_DATABASE_ID"))
}
