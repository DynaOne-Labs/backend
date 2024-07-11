package model

import (
	"github.com/zhufuyi/sponge/pkg/ggorm"
	"time"
)

type User struct {
	ggorm.Model `gorm:"embedded"` // embed id and time

	TwId          string    `gorm:"column:twId;type:varchar(255)" json:"twId"`           // 推特id
	TwName        string    `gorm:"column:twName;type:varchar(255)" json:"twName"`       // 推特名
	AvatarUrl     string    `gorm:"column:avatarUrl;type:varchar(255)" json:"avatarUrl"` // 头像
	Token         string    `gorm:"column:token;type:varchar(255)" json:"token"`
	TokenSecret   string    `gorm:"column:tokenSecret;type:varchar(255)" json:"tokenSecret"`
	WalletAddress string    `gorm:"column:walletAddress;type:varchar(255)" json:"walletAddress"` // 钱包地址
	UpdatedAt     time.Time `gorm:"column:updatedAt;type:datetime;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	CreatedAt     time.Time `gorm:"column:createdAt;type:datetime" json:"createdAt"`
	DeletedAt     time.Time `gorm:"column:deletedAt;type:datetime" json:"deletedAt"`
	LikeCount     int       `gorm:"column:likeCount;type:int(11)" json:"likeCount"` // 点赞数量
}

// TableName table name
func (m *User) TableName() string {
	return "user"
}
