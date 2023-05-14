package notion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Task struct {
}

func GetTasks(notionToken, databaseId string) {
	var (
		uri           = "https://api.notion.com/v1/databases/" + databaseId + "/query"
		auth          = "Bearer " + notionToken
		contentType   = "application/json"
		notionVersion = "2022-06-28"
	)

	type SearchData struct{}

	data := SearchData{}

	d, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", uri, bytes.NewBuffer(d))
	// TODO: エラー処理

	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Authorization", auth)
	req.Header.Add("Notion-Version", notionVersion)

	var (
		client = &http.Client{}
	)

	res, _ := client.Do(req)
	// TODO: エラー処理
	defer res.Body.Close()

	r, _ := io.ReadAll(res.Body)
	json := string(r)

	fmt.Println(json)
}
