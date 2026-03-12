package cmd

import (
	"fmt"

	"proj1/pkg/checker"

	"github.com/spf13/cobra"
)

var (
	defaultURLs = []string{
		"https://google.com",
		"https://github.com",
		"https://golang.org",
		"https://httpbin.org/delay/2",
	}
)

var checkCmd = &cobra.Command{
	Use:   "check [urls...]",
	Short: "Check one or more URLs concurrently",
	Long:  "Check URLs and print status and latency. If no URLs are given, uses a default list.",
	RunE:  runCheck,
}

func runCheck(cmd *cobra.Command, args []string) error {
	urls := defaultURLs
	if len(args) > 0 {
		urls = args
	}

	fmt.Println("Checking URLs concurrently...")
	results := checker.Run(urls)
	for _, res := range results {
		fmt.Printf("[%s] %s - %v\n", res.Status, res.URL, res.Latency)
	}
	return nil
}
