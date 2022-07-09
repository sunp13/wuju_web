package utils

import "strings"

func UtilTranslate(handcliap string) string {
	homeHandicap := strings.Replace(handcliap, "0.0", "平手", -1)
	homeHandicap = strings.Replace(homeHandicap, "-", "", -1)
	homeHandicap = strings.Replace(homeHandicap, "+", "受", -1)
	homeHandicap = strings.Replace(homeHandicap, "0.5", "半球", -1)
	homeHandicap = strings.Replace(homeHandicap, "1.0", "一球", -1)
	homeHandicap = strings.Replace(homeHandicap, "1.5", "一球半", -1)
	homeHandicap = strings.Replace(homeHandicap, "2.0", "两球", -1)
	homeHandicap = strings.Replace(homeHandicap, "2.5", "两球半", -1)
	homeHandicap = strings.Replace(homeHandicap, "3.0", "三球", -1)
	homeHandicap = strings.Replace(homeHandicap, "3.5", "三球半", -1)
	homeHandicap = strings.Replace(homeHandicap, "4.0", "四球", -1)
	homeHandicap = strings.Replace(homeHandicap, "4.5", "四球半", -1)
	homeHandicap = strings.Replace(homeHandicap, "5.0", "五球", -1)
	homeHandicap = strings.Replace(homeHandicap, "5.5", "五球半", -1)
	homeHandicap = strings.Replace(homeHandicap, "6.0", "六球", -1)
	homeHandicap = strings.Replace(homeHandicap, "6.5", "六球半", -1)
	homeHandicap = strings.Replace(homeHandicap, "7.0", "七球", -1)
	homeHandicap = strings.Replace(homeHandicap, "7.5", "七球半", -1)
	homeHandicap = strings.Replace(homeHandicap, "8.0", "八球", -1)
	homeHandicap = strings.Replace(homeHandicap, "8.5", "八球半", -1)
	homeHandicap = strings.Replace(homeHandicap, "9.0", "九球", -1)
	homeHandicap = strings.Replace(homeHandicap, "9.5", "九球半", -1)
	homeHandicap = strings.Replace(homeHandicap, "10.0", "十球", -1)
	homeHandicap = strings.Replace(homeHandicap, "10.5", "十球半", -1)
	homeHandicap = strings.Replace(homeHandicap, ",", "/", -1)
	homeHandicap = strings.Replace(homeHandicap, " ", "", -1)
	return homeHandicap

}
