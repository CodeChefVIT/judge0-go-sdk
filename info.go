package judge0

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Language struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Status struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

// returns a slice of all supported languages by judge0
func (c *Client) GetLanguages() ([]Language, error) {
	data, err := c.doRequest(c.baseURL+"/languages", "GET", nil)

	if err != nil {
		return nil, fmt.Errorf("failed to get languages: %w", err)
	}

	var languages []Language
	err = json.Unmarshal(data, &languages)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize/Unmarshal languages: %w", err)
	}

	return languages, nil
}

// returns the Judge0 language ID for the given language name.
func (c *Client) GetLanguageIDByName(name string) (int, error) {
	languages, err := c.GetLanguages()
	if err != nil {
		return 0, fmt.Errorf("failed to get languages for lookup: %w", err)
	}

	for _, lang := range languages {
		if strings.EqualFold(lang.Name, name) {
			return lang.ID, nil
		}
	}

	return 0, fmt.Errorf("language %q not found", name)
}

// returns a slice of all the judge0 status codes
func (c *Client) GetStatuses() ([]Status, error) {

	data, err := c.doRequest(c.baseURL+"/statuses", "GET", nil)

	if err != nil {
		return nil, fmt.Errorf("failed to get statuses: %w", err)
	}

	var statuses []Status
	err = json.Unmarshal(data, &statuses)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize/Unmarshal statuses: %w", err)
	}

	return statuses, nil

}

const (
	IN_QUEUE              = 1
	PROCESSING            = 2
	ACCEPTED              = 3
	WRONG_ANSWER          = 4
	TIME_LIMIT_EXCEEDED   = 5
	COMPILATION_ERROR     = 6
	RUNTIME_ERROR_SIGSEGV = 7
	RUNTIME_ERROR_SIGXFSZ = 8
	RUNTIME_ERROR_SIGFPE  = 9
	RUNTIME_ERROR_SIGABRT = 10
	RUNTIME_ERROR_NZEC    = 11
	RUNTIME_ERROR_OTHER   = 12
	INTERNAL_ERROR        = 13
	EXEC_FORMAT_ERROR     = 14
)

var statusMap = map[int]string{
	IN_QUEUE:              "IN_QUEUE",
	PROCESSING:            "PROCESSING",
	ACCEPTED:              "ACCEPTED",
	WRONG_ANSWER:          "WRONG_ANSWER",
	TIME_LIMIT_EXCEEDED:   "TIME_LIMIT_EXCEEDED",
	COMPILATION_ERROR:     "COMPILATION_ERROR",
	RUNTIME_ERROR_SIGSEGV: "RUNTIME_ERROR_SIGSEGV",
	RUNTIME_ERROR_SIGXFSZ: "RUNTIME_ERROR_SIGXFSZ",
	RUNTIME_ERROR_SIGFPE:  "RUNTIME_ERROR_SIGFPE",
	RUNTIME_ERROR_SIGABRT: "RUNTIME_ERROR_SIGABRT",
	RUNTIME_ERROR_NZEC:    "RUNTIME_ERROR_NZEC",
	RUNTIME_ERROR_OTHER:   "RUNTIME_ERROR_OTHER",
	INTERNAL_ERROR:        "INTERNAL_ERROR",
	EXEC_FORMAT_ERROR:     "EXEC_FORMAT_ERROR",
}