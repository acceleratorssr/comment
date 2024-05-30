package models

import "time"

type CommentContentModels struct {
	CommentID   int64     `gorm:"primaryKey;foreignKey:CommentID" json:"comment_id,omitempty"`         //主键
	CreatedAt   time.Time `gorm:"column:createtime;type:datetime(0);autoUpdateTime" json:"created_at"` //创建时间
	UpdatedAt   time.Time `gorm:"column:updatetime;type:datetime(0);autoUpdateTime" json:"updated_at"` //修改时间
	AtMemberIds string    `json:"at_member_ids,omitempty"`                                             //发表者用户id
	Message     string    `json:"message,omitempty"`                                                   //评论内容
	//Device      string    `json:"device,omitempty"`                                            //设备信息
	//Meta        string    `json:"meta,omitempty"`                                              //评论元数据：背景、字体
	//Platform    int8      `json:"platform,omitempty"`                                          //发表平台
}
