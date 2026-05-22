package discogsapi

import "strings"

func normalize(input string) string {
	fields := strings.Fields(strings.ToLower(input))

	return strings.Join(fields, " ")
}
