package service

type commentMessage struct {
}

var CommentMessage = &commentMessage{}

func (c *commentMessage) CreateCommentMessageTwoStream(stream MessageService_CreateCommentMessageTwoStreamServer) error {
	err := stream.Send(&CreateMessageResponse{
		Success: true,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *commentMessage) GetCommentTwoStream(stream MessageService_GetCommentTwoStreamServer) error {
	return nil
}

func (c *commentMessage) mustEmbedUnimplementedMessageServiceServer() {}
