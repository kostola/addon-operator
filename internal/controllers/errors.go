package controllers

import "errors"

// This error is returned when a reconciled child object already
// exists and is not owned by the current controller/addon
var ErrNotOwnedByUs = errors.New("object is not owned by us")
