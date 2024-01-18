package forms

import "strings"

const AllowedTypes = "image/jpeg,image/png,image/gif,image/jpg,image/webp"

func IsImg(data string) bool {
	return strings.Contains(data, AllowedTypes)
}
