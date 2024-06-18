package xiaomi

import "time"

/*
CREATE TABLE `role` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `title` varchar(255) NOT NULL COMMENT '角色名称',
  `description` varchar(255) DEFAULT NULL COMMENT '角色描述',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态，1表示启用，0表示禁用',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) COMMENT '主键索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色表'
*/

type Role struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"` // 主键ID
	Title       string    `gorm:"column:title" json:"title"`                    // 角色名称
	Description string    `gorm:"column:description" json:"description"`        // 角色描述
	Status      bool      `gorm:"column:status" json:"status"`                  // 状态，1表示启用，0表示禁用
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`          // 创建时间
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`          // 更新时间
}

func (r *Role) TableName() string {
	return "role"
}

func (r *Role) Create() error {
	return DB.Create(r).Error
}

func (r *Role) Update() error {
	return DB.Save(r).Error
}

func (r *Role) Delete() error {
	return DB.Delete(r).Error
}

func (r *Role) FindByID() error {
	return DB.First(r, r.ID).Error
}

func FindALLRoles() ([]Role, error) {
	var roles []Role
	err := DB.Find(&roles).Error
	return roles, err
}
