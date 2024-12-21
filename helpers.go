package judge0

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func (client *Client) doRequest(url string, method string, payload []byte) ([]byte, error) {
	judgereq, err := http.NewRequest(method, url, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("error creating the HTTP request: %v", err)
	}
	judgereq.Header.Add("Content-Type", "application/json")
	client.authProvider.SetAuthHeaders(judgereq)

	resp, err := client.httpClient.Do(judgereq)
	if err != nil {
		return nil, fmt.Errorf("error sending request to Judge0: %v", err)
	}

	defer resp.Body.Close()
	statusCode := resp.StatusCode

	if statusCode != 200 {
		return nil, fmt.Errorf("error performing request. Status code: %v", statusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)

	}
	return body, nil
}
