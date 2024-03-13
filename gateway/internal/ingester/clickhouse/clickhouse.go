package clickhouse

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/missingstudio/common/errors"
)

type ClickHouseIngester struct {
	db    *sql.DB
	table string
}

func NewClickHouseIngester(db *sql.DB, table string) (*ClickHouseIngester, error) {
	return &ClickHouseIngester{db: db, table: table}, nil
}

func (ci *ClickHouseIngester) Ingest(data map[string]interface{}) {
	// Begin a batch transaction
	tx, err := ci.db.Begin()
	if err != nil {
		fmt.Println("Error starting ClickHouse transaction:", err)
		return
	}

	// Defer the commit or rollback based on success or failure
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	query := fmt.Sprintf("INSERT INTO %s (provider, latency, model, total_tokens, prompt_tokens, completion_tokens) VALUES (?, ?, ?, ?, ?, ?)", ci.table)
	_, err = tx.Exec(query, data["provider"], data["latency"].(time.Duration).String(), data["model"], data["total_tokens"], data["prompt_tokens"], data["completion_tokens"])
	if err != nil {
		fmt.Println("Error ingesting data:", err)
		return
	}
}

func (ci *ClickHouseIngester) Get() ([]map[string]interface{}, error) {
	return nil, errors.NewNotImplemented("Get not implemented for ClickHouseIngester")
}

func (ci *ClickHouseIngester) Close() error {
	return ci.db.Close()
}
