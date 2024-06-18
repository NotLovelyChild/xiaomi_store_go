package xiaomi

import (
	"time"
)

/*
CREATE TABLE `manager` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `username` varchar(255) DEFAULT NULL COMMENT '用户名',
  `password` varchar(32) NOT NULL COMMENT '密码',
  `mobile` varchar(11) DEFAULT NULL COMMENT '手机号',
  `email` varchar(255) DEFAULT NULL COMMENT '电子邮件',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态，1表示启用，0表示禁用',
  `role_id` int DEFAULT NULL COMMENT '角色ID',
  `add_time` int DEFAULT NULL COMMENT '添加时间，UNIX时间戳格式',
  `is_super` tinyint(1) DEFAULT '0' COMMENT '是否超级管理员，1表示是，0表示否',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) COMMENT '主键索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='管理员表'
*/

// Manager 定义对应于数据库表的结构体
type Manager struct {
	ID        int64     `gorm:"column:id;primaryKey" json:"id"`              // 主键ID
	Username  string    `gorm:"column:username" json:"username"`             // 用户名
	Password  string    `gorm:"column:password" json:"password"`             // 密码
	Mobile    string    `gorm:"column:mobile" json:"mobile"`                 // 手机号
	Email     string    `gorm:"column:email" json:"email"`                   // 电子邮件
	Status    bool      `gorm:"column:status" json:"status"`                 // 状态，1表示启用，0表示禁用
	RoleID    int64     `gorm:"column:role_id" json:"role_id"`               // 角色ID
	AddTime   int64     `gorm:"column:add_time" json:"add_time"`             // 添加时间，UNIX时间戳格式
	IsSuper   bool      `gorm:"column:is_super" json:"is_super"`             // 是否超级管理员，1表示是，0表示否
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`         // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`         // 更新时间
	Role      *Role     `gorm:"foreignKey:RoleID;references:ID" json:"role"` // 角色
}

func (m *Manager) TableName() string {
	return "manager"
}

func (m *Manager) Create() error {
    return DB.Create(m).Error
}

func (m *Manager) FindWithUserNameAndPassWord() error {
	return DB.Preload("Role").Where("username = ? AND password = ?", m.Username, m.Password).First(m).Error
}

func (m *Manager) FindWithUserName() error {
	return DB.Preload("Role").Where("username = ?", m.Username).First(m).Error
}

func (m *Manager) FindWithID() error {
	return DB.Preload("Role").Where("id = ?", m.ID).First(m).Error
}

func FindAllManager() ([]Manager, error) {
	var managers []Manager
	err := DB.Preload("Role").Find(&managers).Error
	return managers, err
}

func (m *Manager) Update() error {
    return DB.Save(m).Error
}

func (m *Manager) Delete() error {
    return DB.Delete(m).Error
}
