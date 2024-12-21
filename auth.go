package judge0

import (
	"fmt"
)

const (
	authEndpoint      = "/authenticate"
	authorizeEndpoint = "/authorize"
)

// Authenticate verifies if the authentication token is valid
func (client *Client) Authenticate() error {
	url := fmt.Sprintf("https://%s%s", client.baseURL, authEndpoint)

	_, err := client.doRequest(url, "POST", nil)
	if err != nil {
		return fmt.Errorf("authentication failed: %v", err)
	}

	return nil
}

// Authorize verifies if the authorization token is valid
func (client *Client) Authorize() error {
	url := fmt.Sprintf("https://%s%s", client.baseURL, authorizeEndpoint)

	_, err := client.doRequest(url, "POST", nil)
	if err != nil {
		return fmt.Errorf("authorization failed: %v", err)
	}

	return nil
}
