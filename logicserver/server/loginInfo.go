package server

import (
	"strings"
	"sync"
)

type LoginInfo struct {
	Tokens []string
	UserId string
	//AppId    string
	//DeviceId string
	//Platform string
}

type LoginInfoManager struct {
	tokenLoginInfoMap       map[string]*LoginInfo
	tokenLoginInfoMapMutex  sync.Mutex
	userIdLoginInfoMap      map[string]*LoginInfo
	userIdLoginInfoMapMutex sync.Mutex
}

func NEWLoginInfoManager() *LoginInfoManager {
	s := &LoginInfoManager{
		tokenLoginInfoMap:  make(map[string]*LoginInfo),
		userIdLoginInfoMap: make(map[string]*LoginInfo),
	}
	return s
}

func (s *LoginInfoManager) SafeGetLoginInfoWithToken(token string) *LoginInfo {

	s.tokenLoginInfoMapMutex.Lock()
	loginInfo := s.tokenLoginInfoMap[token]
	s.tokenLoginInfoMapMutex.Unlock()

	return loginInfo
}

func (s *LoginInfoManager) SafeAddLoginInfo(token string, userId string) bool {

	s.tokenLoginInfoMapMutex.Lock()
	loginInfo := s.tokenLoginInfoMap[token]
	if loginInfo == nil {
		s.tokenLoginInfoMap[token] = &LoginInfo{
			UserId: userId,
			Tokens: []string{token},
		}
	} else {
		s.tokenLoginInfoMap[token] = &LoginInfo{
			UserId: userId,
			Tokens: loginInfo.Tokens,
		}
	}
	s.tokenLoginInfoMapMutex.Unlock()

	s.userIdLoginInfoMapMutex.Lock()
	loginInfo = s.userIdLoginInfoMap[userId]
	if loginInfo == nil {
		s.userIdLoginInfoMap[token] = &LoginInfo{
			UserId: userId,
			Tokens: []string{token},
		}
	} else {
		temp := []string{token}
		for _, val := range loginInfo.Tokens {
			if !strings.EqualFold(val, token) {
				temp = append(temp, val)
			}
		}
		loginInfo.Tokens = temp
	}
	s.userIdLoginInfoMapMutex.Unlock()
	return true
}

func (s *LoginInfoManager) SafeRemoveLoginInfo(token string, userId string) bool {

	s.tokenLoginInfoMapMutex.Lock()
	delete(s.tokenLoginInfoMap, token)
	s.tokenLoginInfoMapMutex.Unlock()

	s.userIdLoginInfoMapMutex.Lock()

	loginInfo := s.userIdLoginInfoMap[userId]
	if loginInfo != nil {
		temp := []string{}
		for _, val := range loginInfo.Tokens {
			if !strings.EqualFold(val, token) {
				temp = append(temp, val)
			}
		}
		loginInfo.Tokens = temp
		if len(loginInfo.Tokens) == 0 {
			delete(s.userIdLoginInfoMap, userId)
		}
	}
	s.userIdLoginInfoMapMutex.Unlock()

	return true
}

func (s *LoginInfoManager) SafeGetLoginInfoWithUserId(userId string) *LoginInfo {
	s.userIdLoginInfoMapMutex.Lock()
	loginInfo := s.userIdLoginInfoMap[userId]
	s.userIdLoginInfoMapMutex.Unlock()
	return loginInfo
}
