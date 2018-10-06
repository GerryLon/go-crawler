package parser

import (
	"github.com/GerryLon/go-crawler/engine"
	"github.com/GerryLon/go-crawler/model"
	"regexp"
	"strconv"
)

// 正则集合
var reMapping = map[string]*regexp.Regexp{
	"Age":           regexp.MustCompile(`<td><span class="label">年龄：</span>(\d+)岁</td>`),
	"Gender":        regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([男女])</span></td>`),
	"Height":        regexp.MustCompile(`<td><span class="label">身高：</span><span field="">(\d+)[^<]*</span></td>`),
	"Weight":        regexp.MustCompile(`<td><span class="label">体重：</span><span field="">(\d+)[^<]*</span></td>`),
	"Salary":        regexp.MustCompile(`<td><span class="label">月收入：</span>([^>]+)</td>`),
	"Marriage":      regexp.MustCompile(`<td><span class="label">婚况：</span>([^>]+)</td>`),
	"Education":     regexp.MustCompile(`<td><span class="label">学历：</span>([^>]+)</td>`),
	"Occupation":    regexp.MustCompile(`<td><span class="label">职业： </span>(\p{Han}+)</td>`),
	"NativePlace":   regexp.MustCompile(`<td><span class="label">籍贯：</span>(\p{Han}+)</td>`),
	"Workplace":     regexp.MustCompile(`<td><span class="label">工作地：</span>([^>]+)</td>`),
	"Constellation": regexp.MustCompile(`<td><span class="label">星座：</span>([^>]+)</td>`),
	"Zodiac":        regexp.MustCompile(`<td><span class="label">生肖：</span><span field="">([^>]+)</span></td>`),
	"House":         regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">(\p{Han}+)</span></td>`),
	"Car":           regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">(\p{Han}+)</span></td>`),
	"Pic":           regexp.MustCompile(`<img data-big-img="([^"]+)"[^>]*>`),
}

func ParseProfile(contents []byte, name string, homepage string) engine.ParseResult {
	result := engine.ParseResult{}
	profile := model.Profile{}

	profile.Name = name
	profile.Homepage = homepage
	profile.Age = extractInt(contents, reMapping["Age"])
	profile.Gender = extractString(contents, reMapping["Gender"])
	profile.Height = extractInt(contents, reMapping["Height"])
	profile.Weight = extractInt(contents, reMapping["Weight"])
	profile.Salary = extractString(contents, reMapping["Salary"])
	profile.Marriage = extractString(contents, reMapping["Marriage"])
	profile.Education = extractString(contents, reMapping["Education"])
	profile.Occupation = extractString(contents, reMapping["Occupation"])
	profile.NativePlace = extractString(contents, reMapping["NativePlace"])
	profile.Workplace = extractString(contents, reMapping["Workplace"])
	profile.Constellation = extractString(contents, reMapping["Constellation"])
	profile.Zodiac = extractString(contents, reMapping["Zodiac"])
	profile.House = extractString(contents, reMapping["House"])
	profile.Car = extractString(contents, reMapping["Car"])
	profile.Pic = extractString(contents, reMapping["Pic"])

	result.Items = append(result.Items, profile)

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}

	return ""
}

func extractInt(contents []byte, re *regexp.Regexp) int {
	i, err := strconv.Atoi(extractString(contents, re))

	if err != nil {
		return 0
	}

	return i
}
