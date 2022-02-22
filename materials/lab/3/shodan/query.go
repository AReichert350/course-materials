// Adeline Reichert
// 2/17/22
// Black Hat Go Lab 3b
// Made based on the format of host.go, but formatted for
//   calls to https://api.shodan.io/shodan/query

package shodan

// Do not use this file directly, do not attemp to compile this source file directly
// Go To lab/3/shodan/main/main.go

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Query struct {
	Matches []QueryDescrip `json:"matches"`
	Total	int				`json:"total"`
}

type QueryDescrip struct {
	Votes		int			`json:"votes"`
	Description	string		`json:"description"`
	Tags		[]string	`json:"tags"`
	Timestamp	string		`json:"timestamp"`
	Title		string		`json:"title"`
	Query		string		`json:"query"`
}

func (s *Client) Query(pageNumber int) (*Query, error) {
	res, err := http.Get(
		fmt.Sprintf("%s/shodan/query?key=%s&page=%d", BaseURL, s.apiKey, pageNumber),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret Query
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}