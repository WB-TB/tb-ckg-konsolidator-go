package usecase

import "fmt"

// the example of helper function
func (e *Example) GetFieldValueExample(index int) string {
	return fmt.Sprintf("Value from helper of index %d", index)
}
