package usecases

import (
	"GraduationProject/middleware"
	"GraduationProject/models"
	"GraduationProject/repositories"
	"GraduationProject/utils"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

// UserUC 是用户实例层的一个结构 用来实现用户接口 UserInterface
type UserUC struct {
}

// ur 是仓库存储层的一个实例
var ur = repositories.NewUserRepo()

// NewUserUC 会返回实例层用户模块的实例
func NewUserUC() UserInterface {
	return &UserUC{}
}

// Signup 用户注册逻辑
func (this *UserUC) Signup(ctx *gin.Context) (int, *Response) {
	var user = new(models.User)
	//检测注册信息是否解析完成
	if err := ctx.Bind(user); err != nil {
		return StatusClientError, &Response{
			Code:    ErrorParameterParse,
			Message: "解析参数错误",
		}
	}
	//业务逻辑
	{
		//sha256 获得密码摘要
		h := sha256.New()
		h.Write([]byte(user.Password))
		user.Password = hex.EncodeToString(h.Sum(nil))
	}
	//插入数据库
	if id, err := ur.Insert(user); err != nil {
		//插入数据库失败
		return StatusServerError, &Response{
			Code:    ErrorDatabaseInsert,
			Message: "数据库插入出错",
			Data:    err,
		}
	} else {
		//操作成功
		return StatusOK, &Response{
			Code:    StatusOK,
			Message: "用户注册成功!",
			Data:    utils.ParseInt64ToString(id),
		}
	}
}

// Signin 用户登陆逻辑
func (this *UserUC) Signin(ctx *gin.Context) (int, *Response) {
	var user = new(models.User)
	//检测登陆信息是否绑定成功
	if err := ctx.Bind(user); err != nil {
		return StatusClientError, &Response{
			Code:    ErrorParameterParse,
			Message: "解析参数错误",
		}
	}
	//判断数据是否完整
	if user.Username != "" || user.Password != "" {
		//业务逻辑
		{
			//转换password
			h := sha256.New()
			h.Write([]byte(user.Password))
			user.Password = hex.EncodeToString(h.Sum(nil))
			//查询是否存在
			if has, user := ur.RecordGet(user); has {
				//登陆成功
				token := middleware.NewJWT().NewToken(user.Id, user.Username, user.Email, user.Phone)
				//将token在data中返回
				return StatusOK, &Response{
					Code:    StatusOK,
					Message: token,
					Data:    user,
				}
			} else {
				//登陆失败
				return StatusServerError, &Response{
					Code:    ErrorDatabaseQuery,
					Message: "数据库未找到该记录",
				}
			}
		}
	} else {
		return StatusClientError, &Response{
			Code:    ErrorParameterDefect,
			Message: "重要参数缺失",
		}
	}
}

// ForgetSendEmail 用户忘记密码 发送邮件逻辑
func (this *UserUC) Forget2SendEmail(ctx *gin.Context) (int, *Response) {
	var user = new(models.User)
	user.Email = ctx.Query("email")
	if has, _ := ur.RecordGet(user); !has {
		return StatusServerError, &Response{
			Code:    ErrorDatabaseQuery,
			Message: "该邮箱未注册",
		}
	}
	{
		//重置链接URL参数  //path/reset/uuid
		uuid := utils.GenerateUUID()
		//uuid:uid 写入redis 指定一小时过期
		if err := ur.SetResetArgs(uuid, user.Id, utils.GetConfig().User.ResetExpire); err != nil {
			return StatusServerError, &Response{
				Code:    ErrorDatabaseInsert,
				Message: "Redis操作失败",
				Data:    err.Error(),
			}
		}
		//发送邮件
		es := utils.NewEmailServer()
		es.To = []string{user.Email}
		err := es.Send(&utils.EmailContent{
			Subject: "[GraduationProject]重置密码",
			Content: "Hi " + user.Email + ",<br/><br/>我们的系统收到一个请求，说你希望通过电子邮件重新设置你在 GraduationProject 的密码。你可以点击下面的链接开始重设密码：<br/><br/> <a href=http://localhost:8080/reset?uuid=" + uuid + ">http://localhost:8080/reset?uuid=" + uuid + "</a> <br/><br/>如果这个请求不是由你发起的，那没问题，你不用担心，你可以安全地忽略这封邮件。<br/><br/> 如果你有任何疑问，可以回复这封邮件向我们提问。<br/><br/>GraduationProject <br/><br/> p.s. 作为安全备注，这次密码找回请求是由 IP 地址 " + utils.RemoteIp(ctx.Request) + "(" + utils.IpGeolocation(utils.RemoteIp(ctx.Request)) + ") 使用浏览器 " + ctx.Request.Header.Get("User-Agent") + " 在 " + utils.GetNowDateTime() + " 发起的",
		})
		//检测发送状况
		if err != nil {
			return StatusServerError, &Response{
				Code:    ErrorEmailSend,
				Message: "邮件发送失败",
				Data:    err.Error(),
			}
		}
	}
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: "邮件已发送至您的邮箱,请按照邮件提示操作",
	}
}

// ForgetUpdatePassword 用户忘记密码 修改密码逻辑
func (this *UserUC) Forget2ResetPassword(ctx *gin.Context) (int, *Response) {
	key := ctx.Query("uuid")
	password := ctx.Query("password")
	val, err := ur.GetResetArgs(key)
	if err != nil {
		return StatusServerError, &Response{
			Code:    ErrorDatabaseQuery,
			Message: "链接失效",
			Data:    err.Error(),
		}
	}
	var user = new(models.User)
	user.Id = utils.ParseStringToInt64(val)
	// 密码摘要获取
	h := sha256.New()
	h.Write([]byte(password))
	user.Password = hex.EncodeToString(h.Sum(nil))
	if _, err := ur.Update(user); err != nil {
		return StatusServerError, &Response{
			Code:    ErrorDatabaseUpdate,
			Message: "数据库插入失败",
		}
	}
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: "操作成功",
	}
}

// Profile 用户查看个人信息逻辑
func (this *UserUC) Profile(ctx *gin.Context) (int, *Response) {
	//从前端请求头中包含的UID来作用户区别 JWT只做登陆评判依据
	uid := utils.ParseStringToInt64(ctx.Request.Header.Get("uid"))
	if uid == 0 {
		return StatusServerError, &Response{
			Code:    ErrorParseRemote,
			Message: "服务端解析出错",
		}
	}
	user := ur.FindOneById(uid)
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: "查询成功",
		Data:    user,
	}
}

// ModifyInformation 用户修改信息逻辑
func (this *UserUC) ModifyInformation(ctx *gin.Context) (int, *Response) {
	var user = &models.User{}
	err := ctx.Bind(user)
	//检测参数是否正常
	if err != nil || user.Id == 0 {
		return StatusClientError, &Response{
			Code:    ErrorParameterParse,
			Message: "解析参数错误",
			Data:    err,
		}
	}
	if id, err := ur.Update(user); err != nil {
		//插入数据库失败
		return StatusServerError, &Response{
			Code:    ErrorDatabaseUpdate,
			Message: "数据库操作出错",
		}
	} else {
		//操作成功
		return StatusOK, &Response{
			Code:    StatusOK,
			Message: utils.ParseInt64ToString(id),
		}
	}
}

// FindOne 查询某个用户的信息
func (this *UserUC) FindOne(ctx *gin.Context) (int, *Response) {
	var user = &models.User{}
	id := utils.ParseStringToInt64(ctx.Query("id"))
	user = ur.FindOneById(id)
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: utils.ParseInt64ToString(id),
		Data:    user,
	}
}