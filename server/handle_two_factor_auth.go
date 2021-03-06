package server

import (
	"net/http"

	"github.com/mkideal/log"
	"github.com/mkideal/pkg/netutil/httputil"

	"github.com/mkideal/accountd/api"
	"github.com/mkideal/accountd/model"
)

func (svr *Server) handleTwoFactorAuth(w http.ResponseWriter, r *http.Request) {
	ip := httputil.IP(r)
	argv := new(api.TwoFactorAuthReq)
	err := argv.Parse(r)
	if err != nil {
		log.Info("TwoFactorAuth parse arguments error: %v, IP=%v", err, ip)
		svr.errorResponse(w, r, api.ErrorCode_BadArgument.NewError(err.Error()))
		return
	}
	log.WithJSON(argv).Debug("TwoFactorAuth request, IP=%v", ip)

	var user *model.User
	switch argv.AuthType {
	case api.TwoFaType_Telno:
		user = svr.telno2faAuth(argv, w, r, argv.AuthId, argv.AuthCode)
	case api.TwoFaType_Email:
		user = svr.email2faAuth(w, r, argv.AuthId, argv.AuthCode)
	default:
		svr.errorResponse(w, r, api.ErrorCode_Unsupported2FaType.NewError(argv.AuthType))
	}
	if user != nil {
		_, err = svr.setSession(w, r, user.Id)
		if err != nil {
			log.Error("%s: set session error: %v", argv.CommandName(), err)
			svr.errorResponse(w, r, err)
			return
		}
		token, err := svr.tokenRepo.NewToken(user, "", "")
		if err != nil {
			log.Error("%s: new token error: %v", argv.CommandName(), err)
			svr.errorResponse(w, r, err)
			return
		}
		svr.response(w, r, api.TwoFactorAuthRes{
			User:  makeUserInfo(user),
			Token: makeTokenInfo(token),
		})
	}
}

func (svr *Server) telno2faAuth(argv *api.TwoFactorAuthReq, w http.ResponseWriter, r *http.Request, telno, code string) *model.User {
	if !model.IsTelno(telno) {
		log.Info("illegal telno `%s`", telno)
		svr.errorResponse(w, r, api.ErrorCode_IllegalTelno)
		return nil
	}
	if !svr.config.IsWhiteTelno(telno) {
		if argv.UseThirdVerifyService {
			// TODO: use third party verify service
		} else {
			vcode, err := svr.telnoVerifyCodeRepo.FindTelnoCode(telno)
			if err != nil {
				log.Error("find telno verify code %s error: %v", telno, err)
				svr.errorResponse(w, r, err)
				return nil
			}
			if vcode == nil {
				log.Info("telno verify code for %s not found", telno)
				svr.errorResponse(w, r, api.ErrorCode_VerifyCodeNotFound)
				return nil
			}
			if model.IsExpired(vcode.ExpireAt) {
				log.Info("telno verify code for %s expired", telno)
				svr.errorResponse(w, r, api.ErrorCode_TelnoVerifyCodeExpired)
				return nil
			}
		}
	}
	user, err := svr.userRepo.GetUserByAccount(telno)
	if err != nil {
		log.Error("get user by account %s error: %v", telno, err)
		svr.errorResponse(w, r, err)
		return nil
	} else if user == nil {
		log.Info("TwoFactorAuth: new user for telno %s", telno)
		user = new(model.User)
		user.AccountType = model.AccountType_Telno
		user.CreatedIp = httputil.IP(r)
		user.Account = telno
		user.Nickname = telno
		if err := svr.userRepo.AddUser(user, ""); err != nil {
			log.Error("add user %s error: %v", telno, err)
			svr.errorResponse(w, r, err)
			return nil
		}
	}
	return user
}

func (svr *Server) email2faAuth(w http.ResponseWriter, r *http.Request, email, code string) *model.User {
	// TODO
	return nil
}
