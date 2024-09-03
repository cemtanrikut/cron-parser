package cron

import (
	"strconv"
	"strings"
)

// ExpandCronField expands a cron field based on its type (minute, hour, day of month, month, or day of week).
func ExpandCronField(field, fieldType string) []string {
	var min, max int
	switch fieldType {
	case "minute":
		min, max = 0, 59
	case "hour":
		min, max = 0, 23
	case "day of month":
		min, max = 1, 31
	case "month":
		min, max = 1, 12
	case "day of week":
		min, max = 0, 6
	}

	if field == "*" {
		// Expand wildcard (*) to full range
		return rangeToList(min, max)
	} else if strings.Contains(field, "*/") {
		// Expand step values (e.g., "*/15")
		step, _ := strconv.Atoi(strings.TrimPrefix(field, "*/"))
		return stepToList(min, max, step)
	} else if strings.Contains(field, ",") {
		// Expand list values (e.g., "1,15")
		return expandListValues(field)
	} else if strings.Contains(field, "-") {
		// Expand range values (e.g., "1-5")
		parts := strings.Split(field, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		return rangeToList(start, end)
	}
	// Single value
	return []string{field}
}

// rangeToList generates a list of numbers from min to max.
func rangeToList(min, max int) []string {
	var list []string
	for i := min; i <= max; i++ {
		list = append(list, strconv.Itoa(i))
	}
	return list
}

// stepToList generates a list of numbers from min to max with a specific step.
func stepToList(min, max, step int) []string {
	var list []string
	for i := min; i <= max; i += step {
		list = append(list, strconv.Itoa(i))
	}
	return list
}

// expandListValues expands a comma-separated list of values (e.g., "1,15,30").
func expandListValues(field string) []string {
	values := strings.Split(field, ",")
	var expanded []string
	for _, v := range values {
		expanded = append(expanded, v)
	}
	return expanded
}
