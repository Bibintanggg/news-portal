package pagination

import "errors"

var (
	ErrorMaxPage     = errors.New("Choosen page more than total page")
	ErrorPage        = errors.New("Page must greater than 0")
	ErrorPageEmpty   = errors.New("Page cannot be empty")
	ErrorPageInvalid = errors.New("Page invalid, must be number")
)
