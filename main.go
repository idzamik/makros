package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	NAME := os.Args[1]
	fmt.Printf("Создание репозитория с названием: %s", NAME)

	client := &http.Client{}
	var data = strings.NewReader(fmt.Sprintf(`{"name":"%s","private":false}`, NAME))
	req, err := http.NewRequest("POST", "https://api.github.com/user/repos", data)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth("idzamik", string("TOKEN"))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(bodyText, &jsonData); err != nil {
		log.Fatal("Ошибка парсинга JSON:", err)
	}

	cloneURL, ok := jsonData["clone_url"].(string)
	if !ok {
		log.Fatal("Не удалось найти поле clone_url")
	}

	fmt.Printf("\ngit clone %s\n", cloneURL)
}
