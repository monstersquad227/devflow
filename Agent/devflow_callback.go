package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	var deploymentName string
	var jobId string
	var status string

	flag.StringVar(&deploymentName, "deployment_name", "", "deployment_name")
	flag.StringVar(&jobId, "job_name", "", "job_name")
	flag.StringVar(&status, "status", "", "build status")

	flag.Parse()

	if deploymentName == "" || jobId == "" || status == "" {
		fmt.Println("Usage: devflow_callback OPTIONS")
		fmt.Println("Options:")
		fmt.Println("    --deployment_name string")
		fmt.Println("    --job_name string")
		fmt.Println("    --status string")
		return
	}
	encodeJobId := url.QueryEscape(jobId)

	baseUrl := fmt.Sprintf("http://10.11.11.47:8080/devflow/projects/%s/builds/%s/status/%s", deploymentName, encodeJobId, status)

	req, err := http.NewRequest("PUT", baseUrl, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	token := os.Getenv("devflowToken")
	if token == "" {
		fmt.Println("devflowToken 环境变量为空")
		return
	}

	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", token)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	fmt.Println("response status:", resp.Body)
}
