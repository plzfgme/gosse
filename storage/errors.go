package storage

import "errors"

// ErrKeyNotFound should be returned when key isn't found in Retriever.Get.
var ErrKeyNotFound = errors.New("Key not found")
