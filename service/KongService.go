package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type KongService struct{}

func (k *KongService) FetchUpstreams() (interface{}, error) {
	kongUrl := fmt.Sprintf("http://%s:%s/upstreams", "192.168.1.198", "8001")
	req, err := http.NewRequest("GET", kongUrl, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	upstreamsByte, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	type Upstream struct {
		Name string `json:"name"`
	}

	type UpstreamResponse struct {
		Data []Upstream `json:"data"`
	}
	var result UpstreamResponse
	err = json.Unmarshal(upstreamsByte, &result)
	if err != nil {
		return nil, err
	}

	var names []string
	for _, upstream := range result.Data {
		names = append(names, upstream.Name)
	}

	return names, nil
}
