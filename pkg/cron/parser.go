package cron

import (
	"errors"
	"fmt"
	"strings"
)

// ParseCronExpression parses the cron expression and returns a map of expanded fields.
func ParseCronExpression(cronExpression string) (map[string]string, error) {
	parts := strings.SplitN(cronExpression, " ", 6)
	if len(parts) != 6 {
		return nil, errors.New("invalid cron expression format")
	}

	// Parse each part of the cron expression using the ExpandCronField function
	cronFields := map[string]string{
		"minute":       strings.Join(ExpandCronField(parts[0], "minute"), " "),
		"hour":         strings.Join(ExpandCronField(parts[1], "hour"), " "),
		"day of month": strings.Join(ExpandCronField(parts[2], "day of month"), " "),
		"month":        strings.Join(ExpandCronField(parts[3], "month"), " "),
		"day of week":  strings.Join(ExpandCronField(parts[4], "day of week"), " "),
		"command":      parts[5],
	}

	return cronFields, nil
}

// FormatOutput formats the parsed cron fields into a table-like output.
func FormatOutput(cronFields map[string]string) {
	for _, field := range []string{"minute", "hour", "day of month", "month", "day of week", "command"} {
		// Use %-14s to ensure that each field name is left-aligned and takes up 14 characters
		// Followed by its respective values
		println(fmt.Sprintf("%-14s %s", field, cronFields[field]))
	}
}
