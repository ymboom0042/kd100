// Package kd100   /**
package kd100

const (
	NULL    = ""
	gateway = "https://poll.kuaidi100.com"
)

const (
	// 实时快递查询
	methodSyncQuery = "poll/query.do"
	// 快递查询地图轨迹
	methodMapTrack = "poll/maptrack.do"
)

const (
	ComYT = "yuantong"
	ComYD = "yunda"
	ComZT = "zhongtong"
	ComST = "shentong"
	ComYZ = "youzhengguonei"
	ComSF = "shunfeng"
	ComJT = "jtexpress"
)

var ComValue = map[string]string{
	ComYT: "圆通快递",
	ComYD: "韵达快递",
	ComZT: "中通快递",
	ComST: "申通快递",
	ComYZ: "邮政快递",
	ComSF: "顺丰快递",
	ComJT: "极兔速递",
}

const (
	StateOnPassage        = "0"
	StateCollect          = "1"
	StateKnotty           = "2"
	StateChecked          = "3"
	StateChargeBack       = "4"
	StateDeliver          = "5"
	StateCustomsClearance = "8"
	StateRejection        = "14"
)

var StateValue = map[string]string{
	StateOnPassage:        "在途",
	StateCollect:          "揽收",
	StateKnotty:           "疑难",
	StateChecked:          "签收",
	StateChargeBack:       "退签",
	StateDeliver:          "派件",
	StateCustomsClearance: "清关",
	StateRejection:        "拒签",
}
