package handler

import (
	"context"
	"io"
	"mainproject/models"
	"mainproject/utils"
	"time"

	log "go-micro.dev/v4/logger"

	pb "putUserName/proto"
)

type PutUserName struct{}

func (e *PutUserName) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received PutUserName.Call request: %v", req)
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)

	redCli := models.InitRedis().Get()
	defer redCli.Close()
	_, err := redCli.Do("get", req.Userid+"_name")
	if err != nil {
		_, err = redCli.Do("del", req.Userid+"_name")
		if err != nil {
			log.Infof("Redis key deleted")
		}

		defer func() {
			go func() {
				time.Sleep(time.Second * 1)
				redCli.Do("del", req.Userid+"_name")
			}()
		}()
	}

	db, err := models.InitDb()
	defer db.Close()
	var usrdb models.User
	if err != nil {
		log.Infof("DB ERROR")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}
	err = db.Where("mobile=?", req.Userid).First(&usrdb).Error
	if err != nil {
		log.Infof("Query Error" + req.Userid)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}
	usrdb.Name = req.Username
	rsp.Username = req.Username
	err = db.Save(&usrdb).Error
	if err != nil {
		log.Infof("Modify Error" + req.Userid)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}
	return nil
}

func (e *PutUserName) ClientStream(ctx context.Context, stream pb.PutUserName_ClientStreamStream) error {
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

func (e *PutUserName) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.PutUserName_ServerStreamStream) error {
	log.Infof("Received PutUserName.ServerStream request: %v", req)
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

func (e *PutUserName) BidiStream(ctx context.Context, stream pb.PutUserName_BidiStreamStream) error {
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
