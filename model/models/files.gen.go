// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"
)

const TableNameFile = "files"

// File mapped from table <files>
type File struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	FileID     int64     `gorm:"column:file_id;primaryKey" json:"file_id"`
	UserID     int64     `gorm:"column:user_id;not null" json:"user_id"`
	UploadType string    `gorm:"column:upload_type;not null;comment:上传类型[七牛云，本地..]" json:"upload_type"`          // 上传类型[七牛云，本地..]
	SceneType  string    `gorm:"column:scene_type;not null;comment:场景类别[风景，人物...]" json:"scene_type"`            // 场景类别[风景，人物...]
	BaseURL    string    `gorm:"column:base_url;not null;comment:url(本地路径就是/，非本地就存远程URL)" json:"base_url"`       // url(本地路径就是/，非本地就存远程URL)
	Path       string    `gorm:"column:path;not null;comment:本地路径 /home/xxx 或者是远程的域名" json:"path"`               // 本地路径 /home/xxx 或者是远程的域名
	Name       string    `gorm:"column:name;not null;comment:文件原始名" json:"name"`                                 // 文件原始名
	Extension  string    `gorm:"column:extension;not null;comment:扩展名" json:"extension"`                         // 扩展名
	Size       int32     `gorm:"column:size;not null;comment:长度" json:"size"`                                    // 长度
	Year       int32     `gorm:"column:year;not null;comment:年份" json:"year"`                                    // 年份
	Month      int32     `gorm:"column:month;not null;comment:月份" json:"month"`                                  // 月份
	Day        int32     `gorm:"column:day;not null;comment:日" json:"day"`                                       // 日
	UploadIP   string    `gorm:"column:upload_ip;not null;comment:上传者ip" json:"upload_ip"`                       // 上传者ip
	Status     int32     `gorm:"column:status;not null;default:1;comment:状态[-1:删除;0:禁用;1启用]" json:"status"`      // 状态[-1:删除;0:禁用;1启用]
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:数据创建时间" json:"create_time"` // 数据创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;comment:数据更新时间" json:"update_time"` // 数据更新时间
}

// TableName File's table name
func (*File) TableName() string {
	return TableNameFile
}