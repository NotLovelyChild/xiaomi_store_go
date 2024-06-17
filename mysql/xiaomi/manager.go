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
