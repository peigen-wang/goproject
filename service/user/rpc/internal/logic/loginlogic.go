package logic

import (
	"context"

	"mall/common/cryptx"
	"mall/common/xerr"
	"mall/service/user/model"
	"mall/service/user/rpc/internal/svc"
	"mall/service/user/rpc/user"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 登录实现
func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {

	res, err := l.svcCtx.UserModel.FindOneByMobile(in.Mobile)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "根据手机号查询用户信息失败，mobile:%s,err:%v", in.Mobile, err)
		}
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}

	password := cryptx.PasswordEnrypt(l.svcCtx.Config.Salt, in.Password)
	if password != res.Password {
		return nil, errors.Wrapf(xerr.NewErrMsg("账号或密码不正确"), "密码匹配出错")
	}

	return &user.LoginResponse{
		Id:     res.Id,
		Name:   res.Name,
		Mobile: res.Mobile,
		Gender: res.Gender,
	}, nil

}
