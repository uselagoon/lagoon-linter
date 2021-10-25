package lagoonyml

import "fmt"

// The ErrLint type represents and error in .lagoon.yml validation.
type ErrLint struct {
	Detail string
}

func (e *ErrLint) Error() string {
	return fmt.Sprintf("failed validation: %v", e.Detail)
}
