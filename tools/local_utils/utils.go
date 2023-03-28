package local_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/lfxnxf/emo-frame/utils"
	"time"
)

const (
	SqlstateErDupEntry    = 1062
	SqlstateErUnkownError = 1105

	ExcelDefaultSheet = "Sheet1"
)

func IsDupEntryError(err error) bool {
	if errMysql, ok := err.(*mysql.MySQLError); ok {
		if errMysql.Number == SqlstateErDupEntry {
			return true
		}
	}
	return false
}

func GetMysqlErrCode(err error) uint16 {
	if errMysql, ok := err.(*mysql.MySQLError); ok {
		return errMysql.Number
	}
	return SqlstateErUnkownError
}

func TransErrIfDupEntry(originErr error, customErr error) error {
	if errMysql, ok := originErr.(*mysql.MySQLError); ok {
		if errMysql.Number == SqlstateErDupEntry {
			return customErr
		}
	}
	return originErr
}

/*
*
获取下周周一的日期
*/
func GetNextWeekFirstDate() (weekMonday string) {
	thisWeekMonday := utils.GetNowDateOfWeek()
	weekMonday = GetNextWeekSameDayDate(thisWeekMonday)
	return
}

/*
*
获取七天后的日期
*/
func GetNextWeekSameDayDate(sourceWeekDay string) (nextSameDay string) {
	date, _ := time.Parse("2006-01-02", sourceWeekDay)
	nextSameDay = date.AddDate(0, 0, 7).Format("2006-01-02")
	return
}

/*
*
获取六天后的日期
*/
func GetNextWeekYesterdayDate(sourceWeekDay string) (nextSameDay string) {
	date, _ := time.Parse("2006-01-02", sourceWeekDay)
	nextSameDay = date.AddDate(0, 0, 6).Format("2006-01-02")
	return
}

func DifferenceIds(a, b []int64) []int64 {
	var diffArray []int64
	temp := map[int64]struct{}{}

	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}

	for _, val := range a {
		if _, ok := temp[val]; !ok {
			diffArray = append(diffArray, val)
		}
	}

	return diffArray
}

// 根据生日计算当前年龄
func CalcAge(birthdayStr string) int64 {
	birthday, _ := time.Parse("2006-01-02", birthdayStr)
	birthYear := birthday.Year()
	thisYear := time.Now().Year()
	return int64(thisYear - birthYear)
}
