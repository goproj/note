package models

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func Avatar90(email string) string {
	h := md5.New()
	h.Write([]byte(strings.ToLower(email)))
	avatar := fmt.Sprintf("%x", h.Sum(nil))
	return fmt.Sprintf("http://www.gravatar.com/avatar/%s?s=90&amp;d=mm", avatar)
}
