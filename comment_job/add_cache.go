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

func addSubjectCache(ctx context.Context, csm *models.CommentSubjectModels) {
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
	global.Redis.Set(ctx, oidType, csMarshal, 0)
	global.Log.Info("message apply in redis -> subject")
}

func addCommentIndexCache(ctx context.Context, objID, objType int64, commentIndexModel *models.CommentIndexModels) {
	oidTypeSort := strconv.FormatInt(objID, 10) + "_" + strconv.FormatInt(objType, 10) + "sortByDESC"
	score := float64(commentIndexModel.Like + commentIndexModel.RootCount)
	commentIndexAndMemberID := strconv.Itoa(int(commentIndexModel.ID)) + "_" + strconv.Itoa(int(commentIndexModel.MemberID))
	global.Redis.ZAdd(ctx, oidTypeSort, []redis.Z{{Score: score, Member: commentIndexAndMemberID}}...)
	global.Log.Info("message apply in redis -> comment_index")
}

func addCommentCommentCache(ctx context.Context, comment string, id int64) {
	content := service.Content{Content: comment}
	contentMarshal, err := proto.Marshal(&content)
	if err != nil {
		return
	}

	global.Redis.Set(ctx, strconv.FormatInt(id, 10), contentMarshal, 0)
	global.Log.Info("message apply in redis -> comment_comment")
}
