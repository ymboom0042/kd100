package kd100

import (
	"fmt"
	"testing"
)

func TestExpress_SyncQuery(t *testing.T) {
	e, _ := NewExpress("1231", "1111")

	p := make(Param)
	p.Set("com", ComZT).Set("num", "75871577111107726")
	resp, err := e.SyncQuery(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, r := range resp.Data {
		fmt.Println(r)
	}
}

func TestExpress_MapTrack(t *testing.T) {
	e, _ := NewExpress("043FEBD7E22E86380780064E0436EEAA", "HicFkjwu2981")

	p := make(Param)
	p.Set("com", ComZT).
		Set("num", "75871577107726").
		Set("from", "上海").
		Set("to", "安徽")

	resp, err := e.MapTrack(p)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Println(resp.Data)
}
