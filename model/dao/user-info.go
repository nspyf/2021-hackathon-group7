package dao

import (
	"gorm.io/gorm"
	"strconv"
)

type UserInfo struct {
	gorm.Model
	UserID       uint   `json:",omitempty" gorm:"index;not null"`
	Nickname     string `json:",omitempty" gorm:"type:varchar(64);index"`
	Avatar       string `json:",omitempty" gorm:"type:varchar(255);"`
	Digest       string `json:",omitempty" gorm:"type:varchar(10000);"`
	Verification string `json:",omitempty" gorm:"type:varchar(255);index"`
}

type UserInfoDigest struct {
	UserID       uint   `json:"user_id,omitempty"`
	Nickname     string `json:"nickname,omitempty"`
	Avatar       string `json:"avatar,omitempty"`
	Verification string `json:"verification,omitempty"`
}

type UserInfoProfile struct {
	UserID       uint   `json:"user_id,omitempty"`
	Nickname     string `json:"nickname,omitempty"`
	Avatar       string `json:"avatar,omitempty"`
	Digest       string `json:",omitempty" gorm:"type:varchar(10000);"`
	Verification string `json:"verification,omitempty"`
}

type UserInfoDao struct {
	Info *UserInfo
	Digest *UserInfoDigest
	Profile *UserInfoProfile
}

func (s *UserInfo) Create() error {
	return DB.Create(s).Error
}

func (s *UserInfo) Retrieve() error {
	return DB.Model(s).Where(s).First(s).Error
}

func (s *UserInfo) Update(change interface{}) error {
	return DB.Model(s).Where(s).Updates(change).Error
}

func (s *UserInfo) GetData(data interface{}) error {
	return DB.Model(s).Where(s).First(data).Error
}

func (s *UserInfoDigest) SetCache() error {
	cacheObj := &JsonCache{
		Data: s,
		ID:   strconv.Itoa(int(s.UserID)),
	}
	return cacheObj.SetData(UserDataCacheTime)
}

func (s *UserInfoDigest) GetCache(id uint) error {
	cacheObj := &JsonCache{
		Data: s,
		ID:   strconv.Itoa(int(id)),
	}
	return cacheObj.GetData()
}

func (s *UserInfoDigest) DelCache() error {
	cacheObj := &JsonCache{
		Data: s,
		ID:   strconv.Itoa(int(s.UserID)),
	}
	return cacheObj.DelData()
}

func (s *UserInfoProfile) SetCache() error {
	cacheObj := &JsonCache{
		Data: s,
		ID:   strconv.Itoa(int(s.UserID)),
	}
	return cacheObj.SetData(UserDataCacheTime)
}

func (s *UserInfoProfile) GetCache(id uint) error {
	cacheObj := &JsonCache{
		Data: s,
		ID:   strconv.Itoa(int(id)),
	}
	return cacheObj.GetData()
}

func (s *UserInfoProfile) DelCache() error {
	cacheObj := &JsonCache{
		Data: s,
		ID:   strconv.Itoa(int(s.UserID)),
	}
	return cacheObj.DelData()
}

func (s *UserInfoDao) GetDigest(id uint) error {
	s.Digest = new(UserInfoDigest)
	if s.Digest.GetCache(id) != nil {
		s.Info = new(UserInfo)
		s.Info.ID = id
		err := s.Info.GetData(s.Digest)
		if err != nil {
			return err
		}
		_ = s.Digest.SetCache()
	}
	return nil
}

func (s *UserInfoDao) GetProfile(id uint) error {
	s.Profile = new(UserInfoProfile)
	if s.Profile.GetCache(id) != nil {
		s.Info = new(UserInfo)
		s.Info.ID = id
		err := s.Info.GetData(s.Profile)
		if err != nil {
			return err
		}
		_ = s.Profile.SetCache()
	}
	return nil
}