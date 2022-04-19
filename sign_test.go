// Package kd100   /**
package kd100

import (
	"fmt"
	"testing"
)

func TestExpress_Sign(t *testing.T) {
	e, _ := NewExpress("121", "1111")

	p := make(Param)
	p.Set("com", "com_value").Set("num", "num_value")
	s := e.Sign(p)
	fmt.Println(s)
}
