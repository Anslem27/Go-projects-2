package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

const (
	photoApi = "http://api.pexels.com/v1/"
)

type Client struct {
	Token         string
	hc            http.Client
	RemainingTime int32
}

func NewClient(token string) *Client {
	c := http.Client{}

	return &Client{Token: token, hc: c}
}

type SearchResults struct {
	Page         int32   `json:"page"`
	PerPage      int32   `json:"per_page"`
	TotalResults int32   `json:"total_results"`
	Nextpage     int32   `json:"nextpage"`
	Photos       []Photo `json:"photos"`
}

type Photo struct {
	Id     int32       `json:"id"`
	Width  int32       `json:"width"`
	Height int32       `json:"height"`
	Url    string      `json:"url"`
	Src    PhotoSource `json:"src"`
}

type PhotoSource struct {
	Original  string `json:"original"`
	Landscape string `json:"landscape"`
	Portrait  string `json:"portrait"`
}

func (c *Client) SearchPhotos(search string, perPage, page int) (*SearchResults, error) {
	url := fmt.Sprintf(photoApi+"/search?query=%s&per_page=%d&page=%d", search, perPage, page)

	resp, err := c.requestDoWithAuth("GET", url)
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result SearchResults
	json.Unmarshal(data, &result)
	return &result, err

}

func (c *Client) requestDoWithAuth(method, url string) (*http.Response, err) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", c.Token)
	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}
	times, err := strconv.Atoi(resp.Header.Get("x-RateLimit-Remaining"))

	if err != nil {
		return resp, nil
	} else {
		c.RemainingTime = int32(times)
	}
}
func main() {

	os.Setenv("PexelsToken", "563492ad6f9170000100000117a7423e73b74a488e69f8e650a03157")
	var TOKEN = os.Getenv("PexelsToken")

	var client = NewClient(TOKEN)

	result, err := client.SearchPhotos("scenary", 15, 1)

	//fetch error
	if err != nil {
		fmt.Errorf("Search Error: %v", err)
	}
	//incase of an invalid search
	if result.Page == 0 {
		fmt.Errorf("Invalid Search")
	}
	fmt.Printf(result)

}
