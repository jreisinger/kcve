package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"text/tabwriter"
	"time"

	"github.com/jreisinger/kcve/feed"
	"github.com/jreisinger/kcve/gh"
)

func main() {
	cves, err := feed.FetchCVEs()
	if err != nil {
		log.Fatal(err)
	}

	const format = "%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "CREATED", "CLOSED", "SUMMARY", "URL")

	for _, cve := range cves {
		issueNumber := path.Base(cve.Url)
		issue, err := gh.FetchIssue(issueNumber)
		if err != nil {
			log.Print(err)
		}

		created, err := time.Parse("2006-01-02T15:04:05Z", issue.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		closed, err := time.Parse("2006-01-02T15:04:05Z", issue.ClosedAt)
		if err != nil {
			log.Fatal(err)
		}

		u, err := url.Parse(cve.Url)
		if err != nil {
			log.Fatal(err)
		}

		summary := cve.Summary
		if len(summary) > 30 {
			summary = summary[:30] + "..."
		}

		fmt.Fprintf(tw, format, created.Format("2006-01-02"), closed.Format("2006-01-02"), summary, u.Host+u.Path)
	}

	tw.Flush()
}
