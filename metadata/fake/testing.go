package fakemetadata

import (
	"os"
	"strings"
)

// IsTest reports whether the current state is in test.
func IsTest() bool {
	return strings.HasSuffix(os.Args[0], "_test")
}
