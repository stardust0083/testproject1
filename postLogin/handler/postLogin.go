package handler

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"mainproject/models"
	"mainproject/utils"
	"time"

	"github.com/garyburd/redigo/redis"
	log "go-micro.dev/v4/logger"

	pb "postLogin/proto"
)

type PostLogin struct{}

func (e *PostLogin) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received PostLogin.Call request: %v", req)
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	m5Pwd := fmt.Sprintf("%x", md5.Sum([]byte(req.Password)))
	rdCli := models.InitRedis().Get()
	defer rdCli.Close()
	usrpwd, err1 := rdCli.Do("get", req.Mobile+"_MD5Password")
	usrname, err2 := rdCli.Do("get", req.Mobile+"_name")
	if err1 == nil || err2 == nil {
		log.Infof("Not found in redis")
		db, err := models.InitDb()
		defer db.Close()
		if err != nil {
			rsp.Errno = utils.RECODE_DBERR
			rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
			return nil
		}

		var usr models.User
		err = db.Where("mobile=?", req.Mobile).First(&usr).Error
		if err != nil {
			rsp.Errno = utils.RECODE_DBERR
			rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
			return nil
		}
		if m5Pwd == usr.Password_hash {
			rsp.Name = usr.Name
			rdCli.Do("setex", usr.Mobile+"_MD5Password", 5*60, usr.Password_hash)
			rdCli.Do("setex", usr.Mobile+"_name", 5*60, usr.Name)
		} else {
			rsp.Errno = utils.RECODE_DATAERR
			rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		}
	} else {
		if m5Pwd == usrpwd {
			rsp.Name, _ = redis.String(usrname, redis.ErrNil)
			rsp.Mobile = req.Mobile
		} else {
			rsp.Errno = utils.RECODE_DATAERR
			rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
			return nil
		}
	}

	return nil
}

func (e *PostLogin) ClientStream(ctx context.Context, stream pb.PostLogin_ClientStreamStream) error {
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

func (e *PostLogin) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.PostLogin_ServerStreamStream) error {
	log.Infof("Received PostLogin.ServerStream request: %v", req)
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

func (e *PostLogin) BidiStream(ctx context.Context, stream pb.PostLogin_BidiStreamStream) error {
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
