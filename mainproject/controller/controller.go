package controller

import (
	"context"
	"fmt"
	GETAREAPB "getArea/proto"
	GETIMAGECODEPB "getImageCode/proto"
	GETSMSCODEPB "getSmsCode/proto"
	GETUSERINFOPB "getUserInfo/proto"
	"path"
	POSTAVATARPB "postAvatar/proto"
	POSTLOGINPB "postLogin/proto"
	POSTRETPB "postRet/proto"

	"image"
	"image/png"
	"mainproject/models"
	"mainproject/utils"
	"net/http"
	"regexp"

	"github.com/afocus/captcha"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/asim/go-micro/plugins/transport/grpc/v4"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
)

func GetArea(ctx *gin.Context) {
	reg := consul.NewRegistry()
	ser := grpc.NewTransport()
	microService := micro.NewService(
		micro.Registry(reg),
		micro.Transport(ser),
	)
	microService.Init()
	client := GETAREAPB.NewGetAreaService("go.micro.srv.GetArea", microService.Client())
	rsp, err := client.Call(context.Background(), &GETAREAPB.CallRequest{})
	if err != nil {
		print("fuck call", rsp)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		// ctx.JSON(http.StatusNotAcceptable, gin.H{"status": http.StatusOK, "message": "fuck!"})
		// return
	}
	// print(rsp)
	// ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Im here!"})
	ctx.JSON(200, rsp)
}

func Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Im here!"})
}

func GetSession(ctx *gin.Context) {
	resp := make(map[string]interface{})

	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

	dataTmp := make(map[string]string)

	s := sessions.Default(ctx)
	userName := s.Get("userName")

	if userName == nil {
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
	} else {
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
		dataTmp["name"] = userName.(string)
		resp["data"] = dataTmp
	}

	ctx.JSON(200, resp)
}

func GetImageCode(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	reg := consul.NewRegistry()
	ser := grpc.NewTransport()
	microService := micro.NewService(
		micro.Registry(reg),
		micro.Transport(ser),
	)
	microService.Init()
	client := GETIMAGECODEPB.NewGetImageCodeService("go.micro.srv.GetImageCode", microService.Client())
	rsp, err := client.Call(context.Background(), &GETIMAGECODEPB.CallRequest{Uuid: uuid})

	if err != nil {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		ctx.JSON(500, rsp)
		return
	}

	//解析返回数据为image,传回前端
	var img image.RGBA
	img.Stride = int(rsp.Stride)

	img.Rect.Min.X = int(rsp.Min.X)
	img.Rect.Min.Y = int(rsp.Min.Y)
	img.Rect.Max.X = int(rsp.Max.X)
	img.Rect.Max.Y = int(rsp.Max.Y)

	img.Pix = []uint8(rsp.Pix)

	var captchaImage captcha.Image
	captchaImage.RGBA = &img

	// 将图片发送给前端
	png.Encode(ctx.Writer, captchaImage)
}

func GetSmscd(ctx *gin.Context) {
	phone := ctx.Param("mobile")
	text := ctx.Query("text")
	uuid := ctx.Query("id")
	fmt.Print(phone, text, uuid)

	reg := consul.NewRegistry()
	ser := grpc.NewTransport()
	microService := micro.NewService(
		micro.Registry(reg),
		micro.Transport(ser),
	)
	microService.Init()
	client := GETSMSCODEPB.NewGetSmsCodeService("go.micro.srv.GetSmsCode", microService.Client())
	_, err := client.Call(context.Background(), &GETSMSCODEPB.CallRequest{Uuid: uuid, Phone: phone, Text: text})
	rsp := make(map[string]interface{})
	rsp["errno"] = utils.RECODE_OK
	rsp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	if err != nil {
		rsp["errno"] = utils.RECODE_MOBILEERR
		rsp["errmsg"] = utils.RecodeText(utils.RECODE_MOBILEERR)
		ctx.JSON(500, rsp)
		return
	}
	regstr, _ := regexp.Compile(`^1(3\d|4[5-9]|5[0-35-9]|6[567]|7[0-8]|8\d|9[0-35-9])\d{8}$`)
	result := regstr.MatchString(phone)
	if !result {
		rsp["errno"] = utils.RECODE_MOBILEERR
		rsp["errmsg"] = utils.RecodeText(utils.RECODE_MOBILEERR)
		ctx.JSON(500, rsp)
		return
	}

	if text == "" || uuid == "" {
		rsp["errno"] = utils.RECODE_DATAERR
		rsp["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR)
		ctx.JSON(500, rsp)
		return
	}
	ctx.JSON(200, rsp)
}

func PostRet(ctx *gin.Context) {

	resp := make(map[string]interface{})
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

	//绑定数据
	var regUser models.RegisterUser
	err := ctx.Bind(&regUser)
	if err != nil {
		resp["errno"] = utils.RECODE_DATAERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR) + "bind error"
		ctx.JSON(200, resp)
		return
	}
	if regUser.Mobile == "" || regUser.Password == "" || regUser.SmsCode == "" {
		resp["errno"] = utils.RECODE_DATAERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR) + "empty error"
		ctx.JSON(200, resp)
		return
	}

	regstr, _ := regexp.Compile(`^1(3\d|4[5-9]|5[0-35-9]|6[567]|7[0-8]|8\d|9[0-35-9])\d{8}$`)
	if !regstr.MatchString(regUser.Mobile) {
		resp["errno"] = utils.RECODE_DATAERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR) + "reg"
		ctx.JSON(200, resp)
		return
	}
	reg := consul.NewRegistry()
	ser := grpc.NewTransport()
	microService := micro.NewService(
		micro.Registry(reg),
		micro.Transport(ser),
	)
	microService.Init()
	client := POSTRETPB.NewPostRetService("go.micro.srv.PostRet", microService.Client())
	rsp, err := client.Call(context.Background(), &POSTRETPB.CallRequest{
		Mobile:   regUser.Mobile,
		Password: regUser.Password,
		SmsCode:  regUser.SmsCode,
	})
	if err != nil {
		resp["errno"] = utils.RECODE_DATAERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR) + "micro" + err.Error()
		ctx.JSON(200, resp)
		return
	}
	if rsp.Errno == utils.RECODE_OK {
		s := sessions.Default(ctx)
		s.Set("userName", regUser.Mobile)
		s.Save()
	}
	ctx.JSON(200, resp)
}

func GetUserInfo(ctx *gin.Context) {
	resp := make(map[string]interface{})
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

	s := sessions.Default(ctx)
	usrid := s.Get("userName")
	if usrid == nil {
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
		ctx.JSON(200, resp)
		return
	}

	reg := consul.NewRegistry()
	ser := grpc.NewTransport()
	microService := micro.NewService(
		micro.Registry(reg),
		micro.Transport(ser),
	)
	microService.Init()
	client := GETUSERINFOPB.NewGetUserInfoService("go.micro.srv.GetUserInfo", microService.Client())
	rsp, err := client.Call(context.Background(), &GETUSERINFOPB.CallRequest{Name: usrid.(string)})

	if err != nil {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		ctx.JSON(200, rsp)
		return
	}
	ctx.JSON(200, rsp)
}

func DeleteSession(ctx *gin.Context) {
	resp := make(map[string]interface{})
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

	s := sessions.Default(ctx)
	s.Delete("userName")
	err := s.Save()
	if err != nil {
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
	}
	ctx.JSON(200, resp)
}

func PostLogin(ctx *gin.Context) {
	resp := make(map[string]interface{})
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

	var usrlogin models.UserLogin
	err := ctx.Bind(&usrlogin)
	if err != nil {
		resp["errno"] = utils.RECODE_REQERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_REQERR)
		ctx.JSON(200, resp)
		return
	}

	reg := consul.NewRegistry()
	ser := grpc.NewTransport()
	microService := micro.NewService(
		micro.Registry(reg),
		micro.Transport(ser),
	)
	microService.Init()
	client := POSTLOGINPB.NewPostLoginService("go.micro.srv.PostLogin", microService.Client())
	rsp, err := client.Call(context.Background(), &POSTLOGINPB.CallRequest{Mobile: usrlogin.Mobile, Password: usrlogin.Password})
	if err != nil {
		resp["errno"] = utils.RECODE_LOGINERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_LOGINERR)
		ctx.JSON(200, resp)
		return
	}
	if rsp.Errno == utils.RECODE_OK {
		s := sessions.Default(ctx)
		s.Set("userName", rsp.Name)
		err = s.Save()
		if err != nil {
			resp["errno"] = utils.RECODE_SESSIONERR
			resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
			ctx.JSON(200, resp)
			return
		}
	} else {
		resp["errno"] = utils.RECODE_LOGINERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_LOGINERR)
		ctx.JSON(200, resp)
		return
	}
	ctx.JSON(200, resp)
}
func PostAvatar(ctx *gin.Context) {
	resp := make(map[string]interface{})
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

	filePost, err := ctx.FormFile("avatar")
	if err != nil {
		resp["errno"] = utils.RECODE_REQERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_REQERR)
		ctx.JSON(200, resp)
		return
	}
	fileExt := path.Ext(filePost.Filename)
	if fileExt != ".jpg" && fileExt != ".png" && fileExt != ".jpeg" {
		resp["errno"] = utils.RECODE_DATAERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR)
		ctx.JSON(200, resp)
		return
	}
	file, _ := filePost.Open()
	buffer := make([]byte, filePost.Size)
	file.Read(buffer)

	s := sessions.Default(ctx)
	usrname := s.Get("userName").(string)
	reg := consul.NewRegistry()
	ser := grpc.NewTransport()
	microService := micro.NewService(
		micro.Registry(reg),
		micro.Transport(ser),
	)
	microService.Init()
	client := POSTAVATARPB.NewPostAvatarService("go.micro.srv.PostAvatar", microService.Client())
	rsp, err := client.Call(context.Background(), &POSTAVATARPB.CallRequest{File: buffer, FileExt: fileExt[1:], Name: usrname})
	if err != nil {
		resp["errno"] = utils.RECODE_IOERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_IOERR)
		ctx.JSON(200, resp)
		return
	}
	ctx.JSON(200, rsp)
}

func PutUserName(ctx *gin.Context){
	
}