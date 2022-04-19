// Package kd100   /**
package kd100

type State string

func (s State) Value() string {
	if v, exist := StateValue[string(s)]; exist {
		return v
	} else {
		return "未知"
	}

}

type BaseResp struct {
	ReturnCode string `json:"return_code"` // 错误会返回
	Message    string `json:"message"`
}

type QueryResp struct {
	BaseResp
	State   State       `json:"state"`
	IsCheck string      `json:"is_check"`
	Nu      string      `json:"nu"`
	Data    []QueryData `json:"data"`
}

type QueryData struct {
	Context string `json:"context"`
	Time    string `json:"time"`
	Ftime   string `json:"ftime"`
}
