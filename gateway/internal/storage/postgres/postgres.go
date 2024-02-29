package postgres

import (
	"errors"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

var (
	queryErr = errors.New("error while creating the query")
	parseErr = errors.New("parsing error")
	dbErr    = errors.New("error while running query")
	dialect  = goqu.Dialect("postgres")
)

var (
	ErrDuplicateKey              = errors.New("duplicate key")
	ErrCheckViolation            = errors.New("check constraint violation")
	ErrForeignKeyViolation       = errors.New("foreign key violation")
	ErrInvalidTextRepresentation = errors.New("invalid input syntax type")
	ErrInvalidID                 = errors.New("invalid id")
)

const (
	TABLE_CONNECTIONS = "connections"
	TABLE_PROMPTS     = "prompts"
)

type Encryptor interface {
	Encrypt(plaintext []byte) ([]byte, error)
	Decrypt(ciphertext []byte) ([]byte, error)
}

func checkPostgresError(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case pgerrcode.UniqueViolation:
			return fmt.Errorf("%w [%s]", ErrDuplicateKey, pgErr.Detail)
		case pgerrcode.CheckViolation:
			return fmt.Errorf("%w [%s]", ErrCheckViolation, pgErr.Detail)
		case pgerrcode.ForeignKeyViolation:
			return fmt.Errorf("%w [%s]", ErrForeignKeyViolation, pgErr.Detail)
		case pgerrcode.InvalidTextRepresentation:
			return fmt.Errorf("%w: [%s %s]", ErrInvalidTextRepresentation, pgErr.Detail, pgErr.Message)
		}
	}
	return err
}
