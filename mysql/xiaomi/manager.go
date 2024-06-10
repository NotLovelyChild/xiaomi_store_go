package xiaomi

import (
	"time"
)

// Manager 定义对应于数据库表的结构体
type Manager struct {
	ID        int64     `gorm:"primaryKey;autoIncrement;comment:主键ID" json:"id"`                                                     // 主键ID
	Username  string    `gorm:"type:varchar(255);comment:用户名" json:"username"`                                                       // 用户名
	Password  string    `gorm:"type:varchar(32);comment:密码" json:"password"`                                                         // 密码
	Mobile    string    `gorm:"type:varchar(11);comment:手机号" json:"mobile"`                                                          // 手机号
	Email     string    `gorm:"type:varchar(255);comment:电子邮件" json:"email"`                                                         // 电子邮件
	Status    bool      `gorm:"type:tinyint(1);default:NULL;comment:状态，1表示启用，0表示禁用" json:"status"`                                   // 状态，1表示启用，0表示禁用
	RoleID    int64     `gorm:"type:int;default:NULL;comment:角色ID" json:"role_id"`                                                   // 角色ID
	AddTime   int64     `gorm:"type:int;default:NULL;comment:添加时间，UNIX时间戳格式" json:"add_time"`                                        // 添加时间，UNIX时间戳格式
	IsSuper   bool      `gorm:"type:tinyint(1);default:0;comment:是否超级管理员，1表示是，0表示否" json:"is_super"`                                 // 是否超级管理员，1表示是，0表示否
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`                             // 创建时间
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
}

func (m *Manager) TableName() string {
	return "manager"
}

func (m *Manager) FindWithUserNameAndPassWord() error {
	return DB.Where("username = ? AND password = ?", m.Username, m.Password).First(m).Error
}
