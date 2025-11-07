package shared

import "errors"

var ErrBookNotFound = errors.New("book not found")
var ErrNoBooksFound = errors.New("no books found")
