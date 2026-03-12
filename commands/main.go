package main

import (
	"fmt"
	"os"

	"proj1/cmd"
)

/**
 * Call the Execute function from the cmd package
 * @author: <author_name>
 * @date: <date>
 * @description: <description>
 */
func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
