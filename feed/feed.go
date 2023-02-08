// https://kubernetes.io/docs/reference/issues-security/official-cve-feed/
package feed

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// K8sUrl is the official Kubernetes CVE feed in JSON format.
const K8sUrl = "https://k8s.io/docs/reference/issues-security/official-cve-feed/index.json"

// CVE represents a CVE feed item.
type CVE struct {
	Url     string
	Summary string
}

// FetchCVEs fetches the list of CVEs from official Kubernetes CVE feed (K8sUrl).
func FetchCVEs() ([]CVE, error) {
	resp, err := http.Get(K8sUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("not OK response status: %s", resp.Status)
	}

	var feed struct {
		Items []CVE
	}

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&feed); err != nil {
		return nil, err
	}

	return feed.Items, nil
}
