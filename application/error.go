package application

import "errors"

var ErrInvalidInput = errors.New("invalid input")
var ErrNoteNotFound = errors.New("note not found")

var ErrStorageNotFound = errors.New("requested data not found")
