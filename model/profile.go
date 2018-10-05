package model

type Profile struct {
	Name          string
	Age           int
	Gender        int      // 性别：1: 男， 2： 女， 3：未知
	Height        int      // 身高
	Weight        int      // 体重
	Salary        string   // 收入
	Marriage      string   // 婚姻状况
	Education     string   // 学历
	Occupation    string   // 职业
	NativePlace   string   // 籍贯
	Workplace     string   // 工作地
	Constellation string   // 星座
	Zodiac        string   // 生肖
	House         string   // 购房
	Car           string   // 购车
	Pics          []string // 个人生活照片
}
