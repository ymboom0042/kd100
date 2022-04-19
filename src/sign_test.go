/**
 * @Author: YMBoom
 * @Description:
 * @File:  sign_test
 * @Version: 1.0.0
 * @Date: 2022/04/19 13:38
 */
package src

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
