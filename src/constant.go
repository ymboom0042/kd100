// Package src /**
package src

const (
	NULL        = ""
	gateway     = "https://poll.kuaidi100.com"
	contentType = "application/x-www-form-urlencoded"
)

const (
	// 实时快递查询
	methodSyncQuery = "poll/query.do"
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
