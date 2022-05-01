package handler

import (
	"context"
	"io"
	"mainproject/models"
	"mainproject/utils"
	"time"

	log "go-micro.dev/v4/logger"

	pb "postAvatar/proto"
)

type PostAvatar struct{}

func (e *PostAvatar) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received PostAvatar.Call request: %v", req)
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	log.Infof(req.FileExt)
	fileId, err := models.UploadFDfs(req.File, req.FileExt)
	if err != nil {
		log.Infof("upload function error")
		rsp.Errno = utils.RECODE_IOERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_IOERR)
		return nil
	}
	log.Infof("fileId:", fileId)
	var path pb.ImgPath
	path.AvatarUrl = fileId
	rsp.Path = &path

	db, err := models.InitDb()
	defer db.Close()
	var usrdb models.User
	if err != nil {
		log.Infof("DB ERROR")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}
	err = db.Where("name=?", req.Name).First(&usrdb).Error
	if err != nil {
		log.Infof("Query Error")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}
	usrdb.Avatar_url = fileId
	db.Save(usrdb)
	return nil
}

func (e *PostAvatar) ClientStream(ctx context.Context, stream pb.PostAvatar_ClientStreamStream) error {
	var count int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Infof("Got %v pings total", count)
			return stream.SendMsg(&pb.ClientStreamResponse{Count: count})
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		count++
	}
}

func (e *PostAvatar) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.PostAvatar_ServerStreamStream) error {
	log.Infof("Received PostAvatar.ServerStream request: %v", req)
	for i := 0; i < int(req.Count); i++ {
		log.Infof("Sending %d", i)
		if err := stream.Send(&pb.ServerStreamResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 250)
	}
	return nil
}

func (e *PostAvatar) BidiStream(ctx context.Context, stream pb.PostAvatar_BidiStreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.BidiStreamResponse{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
