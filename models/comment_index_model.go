package models

import "time"

type CommentIndexModels struct {
	ID             int64                `gorm:"primarykey" json:"id,omitempty"` //主键
	CreateTime     time.Time            `json:"create_time"`                    //创建时间
	UpdateTime     time.Time            `json:"update_time"`                    //修改时间
	Root           int64                `json:"root,omitempty"`                 //根评论iD,不为0是回复评论
	Parent         int64                `json:"parent,omitempty"`               //父评论iD,为0是root评论
	MemberID       int64                `json:"member_id,omitempty"`            //发表者用户id
	ObjID          int64                `json:"obj_id,omitempty"`               //对象id 冗余设计
	Count          int32                `json:"count,omitempty"`                //评论总数
	RootCount      int32                `json:"root_count,omitempty"`           //根评论总数
	Like           int32                `json:"like,omitempty"`                 //点赞数
	Hate           int32                `json:"hate,omitempty"`                 //点踩数
	Attrs          int32                `json:"attrs,omitempty"`                //属性
	State          int8                 `json:"state,omitempty"`                //状态(0-正常、1-隐题)
	ObjType        int8                 `json:"obj_type,omitempty"`             //对象类型 冗余设计
	Floor          int32                `json:"floor,omitempty"`                //评论楼层
	CommentContent CommentContentModels `gorm:"foreignKey:CommentID"`
}
