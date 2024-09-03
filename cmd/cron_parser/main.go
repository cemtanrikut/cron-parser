package main

import (
	"fmt"
	"os"

	"github.com/cemtanrikut/cron-parser/cmd/cron_parser/main.go/pkg/cron"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./cron_parser \"<cron_expression>\"")
		return
	}

	cronExpression := os.Args[1]

	parsedFields, err := cron.ParseCronExpression(cronExpression)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	cron.FormatOutput(parsedFields)
}
