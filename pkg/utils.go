package pkg

import (
	"fmt"
)

func CreateTableQuery(tableName, tableDefinition string) string {
	return fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
            %s
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
        )
    `, tableName, tableDefinition)
}

