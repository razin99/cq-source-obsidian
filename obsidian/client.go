package obsidian

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ObsidianClient struct {
	Token string
	// Optional: defaults to "cloudquery"
	UserAgent string
}

var baseUrl string = "https://api.obsec.io/v1/gql"

type GqlOperation struct {
	OperationName string `json:"operationName"`
	// Something like `variables: { cursor: ID }`
	Variables any    `json:"variables"`
	Query     string `json:"query"`
}

func (o ObsidianClient) NewRequest(ctx context.Context, operation GqlOperation) (*http.Response, error) {
	buf := new(bytes.Buffer)

	if err := json.NewEncoder(buf).Encode(operation); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseUrl, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", o.Token))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if o.UserAgent == "" {
		req.Header.Set("User-Agent", "cloudquery")
	} else {
		req.Header.Set("User-Agent", o.UserAgent)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	type Error struct {
		Message    string `json:"message"`
		Extensions struct {
			Code string `json:"code"`
		} `json:"extensions"`
	}

	if res.StatusCode != 200 {
		var fail struct {
			Errors []Error `json:"errors"`
		}
		if err := json.NewDecoder(res.Body).Decode(&fail); err != nil {
			return nil, fmt.Errorf("[%s] - %s", operation.OperationName, res.Status)
		}
		return nil, fmt.Errorf("[%s] - %s - %v", operation.OperationName, res.Status, fail)
	}

	return res, nil
}
