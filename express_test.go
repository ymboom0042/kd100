package kd100

import (
	"fmt"
	"testing"
)

func TestExpress_SyncQuery(t *testing.T) {
	e, _ := NewExpress("111", "222")

	p := make(Param)
	p.Set("com", ComZT).Set("num", "75871577111107726").Set("resultv2", "4")
	resp, err := e.SyncQuery(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, r := range resp.Data {
		fmt.Println(r.Status)
		fmt.Println(r.StatusCode)
	}
}

func TestExpress_MapTrack(t *testing.T) {
	e, _ := NewExpress("111", "222")

	p := make(Param)
	p.Set("com", ComZT).
		Set("num", "75871577111107726").
		Set("from", "上海").
		Set("to", "安徽")

	resp, err := e.MapTrack(p)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Println(resp.Data)
}
