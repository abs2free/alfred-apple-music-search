package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type SearchResult struct {
	TrackName      string `json:"trackName"`
	ArtistName     string `json:"artistName"`
	CollectionName string `json:"collectionName"`
	TrackViewURL   string `json:"trackViewUrl"`
}

type AlfredResult struct {
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	Arg          string `json:"arg"`
	Autocomplete string `json:"autocomplete"`
	Icon         Icon   `json:"icon"`
}

type Icon struct {
	Path string `json:"path"`
}

type Response struct {
	Items []AlfredResult `json:"items"`
}

func main() {
	// 获取搜索关键词
	if len(os.Args) < 2 {
		fmt.Println("Please provide a search term.")
		return
	}
	searchTerm := strings.Join(os.Args[1:], " ")

	// 构建 API URL
	apiURL := fmt.Sprintf("https://itunes.apple.com/search?term=%s&entity=song&limit=20", searchTerm)

	// 发起请求
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 解析 JSON 响应
	var result struct {
		Results []SearchResult `json:"results"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// 构建 Alfred 结果
	var alfredResults []AlfredResult
	for _, item := range result.Results {
		alfredResult := AlfredResult{
			Title:        item.TrackName,
			Subtitle:     fmt.Sprintf("%s - %s", item.ArtistName, item.CollectionName),
			Arg:          strings.Replace(strings.Replace(item.TrackViewURL, "https", "music", 1), "/us/", "/ca/", 1) + "&app=music",
			Autocomplete: fmt.Sprintf("%s %s", item.TrackName, item.ArtistName),
			Icon: Icon{
				Path: "./icon.png",
			},
		}
		alfredResults = append(alfredResults, alfredResult)
	}

	// 构建响应
	response := Response{Items: alfredResults}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// 输出结果
	fmt.Println(string(jsonResponse))
}
