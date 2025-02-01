package dberrors

type ConflictError struct{}

func (e ConflictError) Error() string {
	return "attempt to insert a record that already exists"
}
