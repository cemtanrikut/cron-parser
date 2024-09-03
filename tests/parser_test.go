package tests

import (
	"testing"

	"github.com/cemtanrikut/cron-parser/pkg/cron"
)

// TestParseCronExpression tests the ParseCronExpression function with a valid cron expression.
func TestParseCronExpression(t *testing.T) {
	expression := "*/15 0 1,15 * 1-5 /usr/bin/find"
	expected := map[string]string{
		"minute":       "0 15 30 45",
		"hour":         "0",
		"day of month": "1 15",
		"month":        "1 2 3 4 5 6 7 8 9 10 11 12",
		"day of week":  "1 2 3 4 5",
		"command":      "/usr/bin/find",
	}

	parsedFields, err := cron.ParseCronExpression(expression)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	for key, val := range expected {
		if parsedFields[key] != val {
			t.Errorf("Expected %s to be %s, got %s", key, val, parsedFields[key])
		}
	}
}

// TestExpandCronField tests the ExpandCronField function for various cron field scenarios.
func TestExpandCronField(t *testing.T) {
	tests := []struct {
		field     string
		fieldType string
		expected  []string
	}{
		{"*/15", "minute", []string{"0", "15", "30", "45"}},
		{"0", "hour", []string{"0"}},
		{"1,15", "day of month", []string{"1", "15"}},
		{"*", "month", []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}},
		{"1-5", "day of week", []string{"1", "2", "3", "4", "5"}},
	}

	for _, tt := range tests {
		result := cron.ExpandCronField(tt.field, tt.fieldType)
		if len(result) != len(tt.expected) {
			t.Errorf("Expected length %d, got %d", len(tt.expected), len(result))
		}
		for i := range result {
			if result[i] != tt.expected[i] {
				t.Errorf("Expected %s, got %s", tt.expected[i], result[i])
			}
		}
	}
}

// TestInvalidCronExpression tests the ParseCronExpression function with an invalid cron expression.
func TestInvalidCronExpression(t *testing.T) {
	_, err := cron.ParseCronExpression("invalid cron expression")
	if err == nil {
		t.Error("Expected an error for invalid cron expression, but got none")
	}
}
