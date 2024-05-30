package models

import "time"

type CommentSubjectModels struct {
	ID        int64     `gorm:"primarykey" json:"id"`                                                //主键 8字节对齐
	ObjType   int8      `gorm:"index:idx_member" json:"obj_type,omitempty"`                          //对象ID和对象类型组成一个键  0为视频 1为文章
	ObjID     int64     `gorm:"index:idx_member" json:"obj_id,omitempty"`                            //对象ID和对象类型组成一个唯一键，如对象是视频or专栏，查找时可以分开
	MemberID  int64     `json:"member_id,omitempty"`                                                 //作者用户id
	CreatedAt time.Time `gorm:"column:createtime;type:datetime(0);autoUpdateTime" json:"created_at"` //创建时间
	UpdatedAt time.Time `gorm:"column:updatetime;type:datetime(0);autoUpdateTime" json:"updated_at"` //修改时间
	Count     int32     `json:"count,omitempty"`                                                     //评论总数，记录楼层号 4字节对齐
	RootCount int32     `json:"root_count,omitempty"`                                                //根评论总数
	AllCount  int32     `json:"all_count,omitempty"`                                                 //评论加回复总数
}
