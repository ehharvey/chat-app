package outcome

import "errors"

// This file defines the types of DB errors we will encounter
// Define errors that end-users can understand
// So avoid things like DB permissions, connectivity

var ErrInternal = errors.New("internal service")

var ErrDuplicate = errors.New("already exists")
