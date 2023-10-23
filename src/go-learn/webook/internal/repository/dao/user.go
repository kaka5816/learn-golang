package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type UserDAO struct {
	db *gorm.DB
}

func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.Ctime = now
	u.Utime = now
	return dao.db.WithContext(ctx).Create(&u).Error
}
func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}

}

type User struct {
	Id       int64  `gorm:"primary, autoIncrement"` //主键自增
	Email    string `gorm:"unique"`                 //唯一索引，用户email唯一
	Password string
	//考虑到时区问题
	Ctime int64
	Utime int64
}
