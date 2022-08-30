package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("ERROR:", err.Error())
		return
	}

	url := "https://snyk.io/api/v1/org/" + os.Getenv("ORGANIZATION_ID") + "/project/" + os.Getenv("PROJECT_ID") + "/aggregated-issues"
	method := "POST"

	payload := strings.NewReader(`{
       "filters": {
        "severities": [ "high", "medium", "low", "critical" ],
        "types": [ "vuln", "license" ],
        "ignored": false,
        "patched": false
      }
    }`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "token "+os.Getenv("SNYK_TOKEN"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var resp SnykResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		log.Println(err.Error())
	}
	val, _ := strconv.Atoi(os.Getenv("ALLOWED_ISSUES"))
	if val < len(resp.Issues) {
		log.Println("Operation failed due to maximum limit of issues,", "found", len(resp.Issues), "issues")
		os.Exit(1)
	} else {
		log.Println("Successful")
	}
}
