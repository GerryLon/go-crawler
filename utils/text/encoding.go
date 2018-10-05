package text

import (
	"github.com/axgle/mahonia"
)

func ConvertToString(src string, srcCode string, tagCode string) []byte {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	return cdata
}

func Gbk2Utf8(content []byte) []byte {
	return ConvertToString(string(content), "gbk", "utf-8")
}
