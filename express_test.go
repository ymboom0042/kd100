/**
 * @Author: YMBoom
 * @Description:
 * @File:  express_test
 * @Version: 1.0.0
 * @Date: 2022/04/19 14:07
 */
package mian

import (
	"fmt"
	"testing"
)

func TestExpress_SyncQuery(t *testing.T) {
	e, _ := NewExpress("1231", "1111")

	p := make(Param)
	p.Set("com", ComZT).Set("num", "75871577107726")
	resp, err := e.SyncQuery(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, r := range resp.Data {
		fmt.Println(r)
	}
}
