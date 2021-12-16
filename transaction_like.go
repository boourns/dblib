package dblib

type TransactionLike interface {
	Rollback() error
	Commit() error
}
