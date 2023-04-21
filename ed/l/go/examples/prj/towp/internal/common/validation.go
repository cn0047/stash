package common

import (
	"regexp"
)

var (
	// ResourceNameRegex - represents regex pattern to validate resource name.
	ResourceNameRegex = regexp.MustCompile(`^([\w:-])+$`)
)

func IsResourceNameValid(value string) bool {
	return ResourceNameRegex.MatchString(value)
}
