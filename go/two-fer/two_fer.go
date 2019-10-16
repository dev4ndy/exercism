package twofer

import (
	"fmt"
	"strings"
)

func ShareWith(name string) string {
	if strings.TrimSpace(name) == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %v, one for me.", name)
}
