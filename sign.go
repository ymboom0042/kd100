// Package src /**
package mian

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

func (e *Express) Sign(p Param) string {
	b, err := json.Marshal(p)
	if err != nil {
		return NULL
	}

	s := fmt.Sprintf("%s%s%s", string(b), e.Key, e.Customer)
	return strings.ToUpper(Md5(s))
}

func Md5(s string) string {
	ctx := md5.New()
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}
