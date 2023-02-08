package gh

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Issue struct {
	CreatedAt string `json:"created_at"`
	ClosedAt  string `json:"closed_at"`
	State     string
	Body      string
}

// FetchIssue gets issue from GitHub API. It caches the response for 1h. This is
// because GitHub API allows only 60 unauthenticated requests per minute.
func FetchIssue(issueNumber string) (*Issue, error) {
	cachedir := filepath.Join("/tmp", "kcve")
	cachefile := filepath.Join(cachedir, issueNumber)

	err := os.MkdirAll(cachedir, 0750)
	if err != nil && !os.IsExist(err) {
		return nil, err
	}

	// Create cache file if it does not exist or is older then 1h.
	cache, err := os.Stat(cachefile)
	if errors.Is(err, os.ErrNotExist) || cache.ModTime().Before(time.Now().Add(-time.Hour)) {
		url := "https://api.github.com/repos/kubernetes/kubernetes/issues/" + issueNumber
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("not OK response status: %s", resp.Status)
		}

		body, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, err
		}

		body = removeResponseHeaders(body)

		if err := os.WriteFile(cachefile, body, 0640); err != nil {
			return nil, err
		}
	}

	file, err := os.Open(cachefile)
	if err != nil {
		return nil, err
	}

	var issue Issue
	dec := json.NewDecoder(file)
	if err := dec.Decode(&issue); err != nil {
		return nil, fmt.Errorf("decoding JSON: %v", err)
	}

	// lines := strings.Split(issue.Body, "\n")
	// versionRE := regexp.MustCompile(`v\d+\.\d+(.\d+)?`)
	// for _, l := range lines {
	// 	if versionRE.MatchString(l) {
	// 		fmt.Println(l)
	// 	}
	// }

	return &issue, nil
}

func removeResponseHeaders(b []byte) []byte {
	if len(b) == 0 {
		return b
	}
	parts := strings.Split(string(b), "\r\n\r\n")
	if len(parts) < 2 {
		return []byte{}
	}
	return []byte(parts[1])
}
