package structs

import "strings"

type tagOptions []string

func (t tagOptions) Has(opt string) bool {
	for _, option := range t {
		if option == opt {
			return true
		}
	}
	return false
}
func parseTag(tag string) (string, tagOptions) {
	res := strings.Split(tag,",")
	return res[0],res[1:]
}