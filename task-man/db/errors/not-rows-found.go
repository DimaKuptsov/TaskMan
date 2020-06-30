package errors

const PGNoRowsFoundError = "pg: no rows in result set"

type NoRowsFoundError struct {
	Message string
}

func (err NoRowsFoundError) Error() string {
	return err.Message
}
