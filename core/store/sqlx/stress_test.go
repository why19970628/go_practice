package sqlx

import (
	"fmt"
	"gorm.io/gorm"
	"regexp"
	"strings"
	"testing"
)

func TestWithTimeout(t *testing.T) {
	fmt.Println(replaceWithShadowTableV2(nil, "select * from users as u where id in (select id from t)"))
}

func replaceWithShadowTableV2(db *gorm.DB, sql string) string {
	// Regular expression to match table names
	tableNameRegex := regexp.MustCompile(`(?i)\b(?:from|join|update|insert into|delete from)\s+([\w]+)(?:\s+as\s+[\w]+)?\s*|(?i)in\s*\(\s*select\s+[\w]+(?:\s+as\s+[\w]+)?\s+from\s+([\w]+)(?:\s+as\s+[\w]+)?\s*\)`)

	// Replace all occurrences of table names with shadow table names
	sql = tableNameRegex.ReplaceAllStringFunc(sql, func(match string) string {
		// Extract table name from the match
		tableName := extractTableName(match)
		fmt.Println("tableName", tableName)
		// Get shadow table name
		shadowTableName := getShadowTableName(tableName) // Your logic to get shadow table name
		// Replace table name with shadow table name
		return strings.Replace(match, tableName, shadowTableName, -1)
	})

	return sql
}

// Function to extract table name from match string
func extractTableName(match string) string {
	// Extract table name from match string
	parts := strings.Fields(match)
	return parts[len(parts)-1]
}

// Function to get shadow table name based on original table name
func getShadowTableName(tableName string) string {
	// Your logic to derive shadow table name from original table name
	// For example, you can append "_shadow" to the original table name
	return tableName + "_shadow"
}
