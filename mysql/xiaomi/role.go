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
	Id          int64     `gorm:"id" json:"id"`
	Title       string    `gorm:"title" json:"title"`
	Description string    `gorm:"description" json:"description"`
	Status      bool      `gorm:"status" json:"status"`
	CreatedAt   time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at" json:"updated_at"`
}

func (Role) TableName() string {
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
    return DB.First(r, r.Id).Error
}

func FindALLRoles() ([]Role, error) {
	var roles []Role
	err := DB.Find(&roles).Error
	return roles, err
}


