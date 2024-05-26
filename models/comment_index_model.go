package models

import "time"

type CommentIndexModels struct {
	ID         int64     `gorm:"primarykey" json:"id,omitempty"`                                       //主键
	CreateTime time.Time `gorm:"column:createtime;type:datetime(0);autoUpdateTime" json:"create_time"` //创建时间
	UpdateTime time.Time `gorm:"column:updatetime;type:datetime(0);autoUpdateTime" json:"update_time"` //修改时间
	Root       int64     `json:"root,omitempty"`                                                       //根评论ID,不为0是回复评论
	Parent     int64     `json:"parent,omitempty"`                                                     //父评论ID,为0是root评论
	MemberID   int64     `json:"member_id,omitempty"`                                                  //发表者用户id
	ObjType    int8      `gorm:"index:idx_member" json:"obj_type,omitempty"`                           //对象类型 0为视频 1为文章 冗余设计
	ObjID      int64     `gorm:"index:idx_member" json:"obj_id,omitempty"`                             //对象id 即该条评论对应的对象 冗余设计
	Count      int32     `json:"count,omitempty"`                                                      //评论总数
	RootCount  int32     `json:"root_count,omitempty"`                                                 //根评论总数
	Like       int32     `json:"like,omitempty"`                                                       //点赞数
	Hate       int32     `json:"hate,omitempty"`                                                       //点踩数
	State      int8      `json:"state,omitempty"`                                                      //状态(0-正常、1-隐藏)
	//Floor          int32                `json:"floor,omitempty"`                                                      //评论楼层
	CommentContent CommentContentModels `gorm:"foreignKey:CommentID"`
}
