package models

import "time"

type CommentContentModels struct {
	CommentID   int64     `gorm:"primaryKey;foreignKey:CommentID" json:"comment_id,omitempty"` //主键
	IP          int64     `json:"ip,omitempty"`                                                //IP地址
	CreatedAt   time.Time `json:"created_at"`                                                  //创建时间
	UpdatedAt   time.Time `json:"updated_at"`                                                  //修改时间
	AtMemberIds string    `json:"at_member_ids,omitempty"`                                     //对象ID
	Device      string    `json:"device,omitempty"`                                            //设备信息
	Message     string    `json:"message,omitempty"`                                           //评论内容
	Meta        string    `json:"meta,omitempty"`                                              //评论元数据：背景、字体
	Platform    int8      `json:"platform,omitempty"`                                          //发表平台
}
