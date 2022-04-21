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
	ReturnCode string `json:"returnCode"` // 错误会返回
	Message    string `json:"message"`
}

type QueryResp struct {
	BaseResp
	State   State       `json:"state"`
	IsCheck string      `json:"isCheck"`
	Nu      string      `json:"nu"`
	Data    []QueryData `json:"data"`
}

type QueryData struct {
	Context string `json:"context"`
	Time    string `json:"time"`
	Ftime   string `json:"ftime"`
}

// MapTrackResp 快递查询地图轨迹
type MapTrackResp struct {
	BaseResp
	Data        []mapTrackData `json:"data"`
	State       string         `json:"state"`
	Condition   string         `json:"condition"`
	IsLoop      bool           `json:"isLoop"`
	TrailUrl    string         `json:"trailUrl"` // 地图链接
	ArrivalName string         `json:"arrivalName"`
	TotalTime   string         `json:"totalTime"`
	RemainTime  string         `json:"remainTime"`
}

type mapTrackData struct {
	QueryData
	AreaCode   string `json:"areaCode"`
	AreaName   string `json:"areaName"`
	Status     string `json:"status"`
	Location   string `json:"location"`
	AreaCenter string `json:"areaCenter"`
	AreaPinYin string `json:"areaPinYin"`
	StatusCode string `json:"statusCode"`
}
