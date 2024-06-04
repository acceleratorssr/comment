package comment_job

import (
	"comment/comment_service/service"
	"comment/global"
	"comment/models"
	"context"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

func addSubjectAndCommentCache(ctx context.Context, csm *models.CommentSubjectModels,
	cim *[]models.CommentIndexModels,
	comment []string) {
	cs := service.CommentSubject{
		Id:        csm.ID,
		ObjType:   service.ObjType(csm.ObjType),
		ObjId:     csm.ObjID,
		MemberId:  csm.MemberID,
		CreatedAt: timestamppb.New(csm.CreatedAt),
		UpdatedAt: timestamppb.New(csm.UpdatedAt),
		Count:     csm.Count,
		RootCount: csm.RootCount,
		AllCount:  csm.AllCount,
	}
	csMarshal, err := proto.Marshal(&cs)
	if err != nil {
		return
	}

	oidType := strconv.FormatInt(csm.ObjID, 10) + "_" + strconv.FormatInt(int64(csm.ObjType), 10)

	pipe := global.Redis.Pipeline()
	pipe.Set(ctx, oidType, csMarshal, 0)

	oidTypeSort := strconv.FormatInt(csm.ObjID, 10) + "_" + strconv.FormatInt(int64(csm.ObjType), 10) + "sortByDESC"
	for i := 0; i < len(*cim); i++ {
		score := float64((*cim)[i].Like + (*cim)[i].RootCount)
		commentIndexAndMemberID := strconv.Itoa(int((*cim)[i].ID)) + "_" + strconv.Itoa(int((*cim)[i].MemberID))
		pipe.ZAdd(ctx, oidTypeSort, []redis.Z{{Score: score, Member: commentIndexAndMemberID}}...)
	}

	for i := 0; i < len(*cim); i++ {
		content := service.Content{Content: comment[i]}
		contentMarshal, err := proto.Marshal(&content)
		if err != nil {
			return
		}
		pipe.Set(ctx, strconv.FormatInt((*cim)[i].ID, 10), contentMarshal, 0)
	}

	//-------------
	_, err = pipe.Exec(ctx)
	if err != nil {
		global.Log.Error("addSubjectAndCommentCache", err)
		return
	}

	global.Log.Info("message apply in redis -> subject")
	global.Log.Info("message apply in redis -> comment_index")
	global.Log.Info("message apply in redis -> comment_comment")
}

//func addSubjectCache(ctx context.Context, csm *models.CommentSubjectModels) {
//	cs := service.CommentSubject{
//		Id:        csm.ID,
//		ObjType:   service.ObjType(csm.ObjType),
//		ObjId:     csm.ObjID,
//		MemberId:  csm.MemberID,
//		CreatedAt: timestamppb.New(csm.CreatedAt),
//		UpdatedAt: timestamppb.New(csm.UpdatedAt),
//		Count:     csm.Count,
//		RootCount: csm.RootCount,
//		AllCount:  csm.AllCount,
//	}
//	csMarshal, err := proto.Marshal(&cs)
//	if err != nil {
//		return
//	}
//
//	oidType := strconv.FormatInt(csm.ObjID, 10) + "_" + strconv.FormatInt(int64(csm.ObjType), 10)
//	global.Redis.Set(ctx, oidType, csMarshal, 0)
//	global.Log.Info("message apply in redis -> subject")
//}
//
//func addCommentIndexCache(ctx context.Context, objID, objType int64, commentIndexModel *[]models.CommentIndexModels) {
//	oidTypeSort := strconv.FormatInt(objID, 10) + "_" + strconv.FormatInt(objType, 10) + "sortByDESC"
//	for i := 0; i < len(*commentIndexModel); i++ {
//		score := float64((*commentIndexModel)[i].Like + (*commentIndexModel)[i].RootCount)
//		commentIndexAndMemberID := strconv.Itoa(int((*commentIndexModel)[i].ID)) + "_" + strconv.Itoa(int((*commentIndexModel)[i].MemberID))
//		global.Redis.ZAdd(ctx, oidTypeSort, []redis.Z{{Score: score, Member: commentIndexAndMemberID}}...)
//		global.Log.Info("message apply in redis -> comment_index")
//	}
//}
//
//func addCommentCommentCache(ctx context.Context, comment string, id int64) {
//	content := service.Content{Content: comment}
//	contentMarshal, err := proto.Marshal(&content)
//	if err != nil {
//		return
//	}
//
//	global.Redis.Set(ctx, strconv.FormatInt(id, 10), contentMarshal, 0)
//	global.Log.Info("message apply in redis -> comment_comment")
//}
