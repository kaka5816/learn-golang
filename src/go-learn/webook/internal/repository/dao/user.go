package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 预定义错误
var (
	ErrDuplicateEmail = errors.New("邮箱冲突")
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.Ctime = now
	u.Utime = now
	err := dao.db.WithContext(ctx).Create(&u).Error
	//断定err是*mysql.MySQLError
	if me, ok := err.(*mysql.MySQLError); ok {
		const duplicateError uint16 = 1062
		if me.Number == duplicateError {
			//用户冲突，邮箱冲突
			return ErrDuplicateEmail
		}
	}
	return err
}

func (dao *UserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("email=?", email).First(&u).Error
	return u, err
}

type User struct {
	Id       int64  `gorm:"primary, autoIncrement"` //主键自增
	Email    string `gorm:"unique"`                 //唯一索引，用户email唯一
	Password string
	//考虑到时区问题
	Ctime int64
	Utime int64
}
