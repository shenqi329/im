package service

import (
	imerror "im/logicserver/error"
	"log"
	"regexp"
)

func CheckDeviceId(deviceId string) error {
	if len(deviceId) <= 0 {
		return imerror.ErrorIllegalParams
	}

	pattern := `^\w{32}$`
	regexp.MustCompile(pattern)
	match, err := regexp.MatchString(pattern, deviceId)
	if err != nil {
		log.Println(err.Error())
		return imerror.ErrorInternalServerError
	}
	if !match {
		return imerror.ErrorIllegalParams
	}

	return nil
}

func CheckVerifyCode(code string) error {
	if len(code) <= 0 {
		return imerror.ErrorIllegalParams
	}
	return nil
}

func CheckToken(token string) error {
	if len(token) <= 0 {
		return imerror.ErrorIllegalParams
	}

	pattern := `^\w{32}$`
	regexp.MustCompile(pattern)
	match, err := regexp.MatchString(pattern, token)
	if err != nil {
		log.Println(err.Error())
		return imerror.ErrorInternalServerError
	}
	if !match {
		return imerror.ErrorIllegalParams
	}

	return nil
}

// func CheckEmail(email string) error {
// 	if len(email) <= 0 {
// 		return imerror.ErrorIllegalParams
// 	}

// 	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
// 	regexp.MustCompile(pattern)
// 	match, err := regexp.MatchString(pattern, email)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return imerror.ErrorInternalServerError
// 	}
// 	if !match {
// 		return imerror.ErrorRegisterEmailFormat
// 	}

// 	return nil
// }

// /*
// 1.用户名为4-20个字符（英文字母、数字）组成

// 2.密码为6-20个字符（英文字母、数字）组成

// 注：用户名均不能单独以数字命名，且字母与数字组合的用户名不能以数字开头，不便之处，请用户谅解！
// */
// func CheckUserName(userName string) error {
// 	if len(userName) <= 0 {
// 		return ssoerror.ErrorIllegalParams
// 	}
// 	pattern := `^[A-Za-z]\w{3,19}$`
// 	regexp.MustCompile(pattern)
// 	match, err := regexp.MatchString(pattern, userName)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return imerror.ErrorInternalServerError
// 	}
// 	if !match {
// 		return imerror.ErrorUserNameFormatError
// 	}
// 	return nil
// }

// func CheckPassword(password string) error {
// 	if len(password) <= 0 {
// 		return imerror.ErrorIllegalParams
// 	}

// 	//ascii字符集
// 	pattern := `^[\x00-\x7F]{8,20}$`
// 	//pattern := `^[:ascii:]{8,20}$`
// 	regexp.MustCompile(pattern)
// 	match, err := regexp.MatchString(pattern, password)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return imerror.ErrorInternalServerError
// 	}
// 	if !match {
// 		log.Println("密码格式不对")
// 		return imerror.ErrorPasswordFormatError
// 	}

// 	//不能是全数字
// 	pattern = `^\d+$`
// 	regexp.MustCompile(pattern)
// 	match, err = regexp.MatchString(pattern, password)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return imerror.ErrorInternalServerError
// 	}
// 	if match {
// 		log.Println("密码格式不对")
// 		return imerror.ErrorPasswordFormatError
// 	}

// 	return nil
// }
