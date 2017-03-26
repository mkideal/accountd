//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// NOTE: auto generated by midc, DON'T edit
//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
package api

import (
	"net/http"
	"net/url"
)

type Request interface {
	CommandName() string
	CommandMethod() string
	Parse(*http.Request) error
	Values() url.Values
}

var commands = make(map[string]Request)
var commandList = make([]Request, 0)

func registerCommand(cmd Request) bool {
	commands[cmd.CommandName()] = cmd
	commandList = append(commandList, cmd)
	return true
}

func Commands() []Request            { return commandList }
func GetCommand(name string) Request { return commands[name] }

type ErrorCode int

const (
	ErrorCode_InternalServerError     ErrorCode = 500001
	ErrorCode_ClientUnauthorized      ErrorCode = 401002
	ErrorCode_MissingArgument         ErrorCode = 400101
	ErrorCode_BadArgument             ErrorCode = 400102
	ErrorCode_IllegalUsername         ErrorCode = 400103
	ErrorCode_IllegalAccountType      ErrorCode = 400104
	ErrorCode_IllegalPassword         ErrorCode = 400105
	ErrorCode_IllegalEmail            ErrorCode = 400106
	ErrorCode_IllegalTelno            ErrorCode = 400107
	ErrorCode_UserNotFound            ErrorCode = 417201
	ErrorCode_TokenNotFound           ErrorCode = 417202
	ErrorCode_ClientNotFound          ErrorCode = 417203
	ErrorCode_SessionNotFound         ErrorCode = 417204
	ErrorCode_CommandNotFound         ErrorCode = 417205
	ErrorCode_VerifyCodeNotFound      ErrorCode = 417206
	ErrorCode_IncorrectPassword       ErrorCode = 417301
	ErrorCode_IncorrectClientSecret   ErrorCode = 417302
	ErrorCode_AccountDuplicated       ErrorCode = 417401
	ErrorCode_TelnoVerifyCodeTooOften ErrorCode = 417402
	ErrorCode_TelnoVerifyCodeExpired  ErrorCode = 417403
	ErrorCode_Unsupported2FaType      ErrorCode = 417404
	ErrorCode_FailedToSendSMSCode     ErrorCode = 417405
	ErrorCode_ThirdPartyError         ErrorCode = 417406
	ErrorCode_TokenExpired            ErrorCode = 417407
)

func (x ErrorCode) Status() int {
	return int(x) / 1000
}

func (x ErrorCode) NewError(description string) Error {
	err := NewError(x.Error(), description)
	err.SetStatus(x.Status())
	return err
}

func (x ErrorCode) Error() string {
	switch x {
	case ErrorCode_InternalServerError:
		return "e_internal_server_error"
	case ErrorCode_ClientUnauthorized:
		return "e_client_unauthorized"
	case ErrorCode_MissingArgument:
		return "e_missing_argument"
	case ErrorCode_BadArgument:
		return "e_bad_argument"
	case ErrorCode_IllegalUsername:
		return "e_illegal_username"
	case ErrorCode_IllegalAccountType:
		return "e_illegal_account_type"
	case ErrorCode_IllegalPassword:
		return "e_illegal_password"
	case ErrorCode_IllegalEmail:
		return "e_illegal_email"
	case ErrorCode_IllegalTelno:
		return "e_illegal_telno"
	case ErrorCode_UserNotFound:
		return "e_user_not_found"
	case ErrorCode_TokenNotFound:
		return "e_token_not_found"
	case ErrorCode_ClientNotFound:
		return "e_client_not_found"
	case ErrorCode_SessionNotFound:
		return "e_session_not_found"
	case ErrorCode_CommandNotFound:
		return "e_command_not_found"
	case ErrorCode_VerifyCodeNotFound:
		return "e_verify_code_not_found"
	case ErrorCode_IncorrectPassword:
		return "e_incorrect_password"
	case ErrorCode_IncorrectClientSecret:
		return "e_incorrect_client_secret"
	case ErrorCode_AccountDuplicated:
		return "e_account_duplicated"
	case ErrorCode_TelnoVerifyCodeTooOften:
		return "e_telno_verify_code_too_often"
	case ErrorCode_TelnoVerifyCodeExpired:
		return "e_telno_verify_code_expired"
	case ErrorCode_Unsupported2FaType:
		return "e_Unsupported2FaType"
	case ErrorCode_FailedToSendSMSCode:
		return "e_failed_to_send_smscode"
	case ErrorCode_ThirdPartyError:
		return "e_third_party_error"
	case ErrorCode_TokenExpired:
		return "e_token_expired"

	}
	return "e_unknown_error"
}

type UserInfo struct {
	Id          int64  `json:"id"`
	Account     string `json:"account"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Gender      int    `json:"gender"`
	Birthday    string `json:"birthday"`
	LastLoginAt string `json:"last_login_at"`
	LastLoginIp string `json:"last_login_ip"`
}

type TokenInfo struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	ExpireAt     string `json:"expire_at"`
}

// 查看帮助
type HelpReq struct {
	Version string `json:"version" cli:"v"`
	Cmd     string `json:"cmd" cli:"c"`
}

func (HelpReq) CommandName() string   { return "Help" }
func (HelpReq) CommandMethod() string { return "GET" }

var _ = registerCommand(&HelpReq{})

type HelpRes struct {
	Commands []string `json:"commands,omitempty"`
	Routers  []string `json:"routers,omitempty"`
	Command  string   `json:"command,omitempty"`
	Method   string   `json:"method,omitempty"`
	Router   string   `json:"router,omitempty"`
}

// oauth2.0 接口 token
type TokenReq struct {
	GrantType    string `json:"grant_type" cli:"g"`     // 授权类型
	Scope        string `json:"scope" cli:"s"`          // token使用范围
	Code         string `json:"code" cli:"c"`           // 授权码
	RedirectURI  string `json:"redirect_uri" cli:"uri"` // 重定向URI
	Username     string `json:"username" cli:"u"`       // 账号
	Password     string `json:"password" cli:"p"`       // 密码
	RefreshToken string `json:"refresh_token" cli:"r"`  // 刷新用token

}

func (TokenReq) CommandName() string   { return "Token" }
func (TokenReq) CommandMethod() string { return "POST" }

var _ = registerCommand(&TokenReq{})

type TokenRes struct {
	TokenType string    `json:"token_type"`
	Token     TokenInfo `json:"token"`
}

// oauth2.0 接口 authorize
type AuthorizeReq struct {
	ClientId     string `json:"client_id" cli:"cid"`
	ResponseType string `json:"response_type" cli:"t"`
	Uid          int64  `json:"uid" cli:"uid"`
	Scope        string `json:"scope" cli:"s,scope"`
	State        string `json:"state" cli:"state"`
}

func (AuthorizeReq) CommandName() string   { return "Authorize" }
func (AuthorizeReq) CommandMethod() string { return "POST" }

var _ = registerCommand(&AuthorizeReq{})

// client 检查
type AuthorizeCheckReq struct {
	ClientId string `json:"client_id" cli:"cid"`
}

func (AuthorizeCheckReq) CommandName() string   { return "AuthorizeCheck" }
func (AuthorizeCheckReq) CommandMethod() string { return "GET" }

var _ = registerCommand(&AuthorizeCheckReq{})

type AuthorizeCheckRes struct {
	Application string `json:"application"`
	Username    string `json:"username"`
}

// token 认证
type TokenAuthReq struct {
	AccessToken string `json:"access_token" cli:"t,token"`
}

func (TokenAuthReq) CommandName() string   { return "TokenAuth" }
func (TokenAuthReq) CommandMethod() string { return "POST" }

var _ = registerCommand(&TokenAuthReq{})

type TokenAuthRes struct {
	User  UserInfo  `json:"user"`
	Token TokenInfo `json:"token"`
}

// 注册
type SignupReq struct {
	AccountType       int    `json:"account_type" cli:"t"`        // 账号类型:参见 model.AccountType 枚举
	Account           string `json:"account" cli:"a"`             // 账号:当 accountType 为第三方账号时为 autorization code
	Password          string `json:"password" cli:"p"`            // 密码:当 accountType 为第三方账号时不需要
	Nickname          string `json:"nickname" cli:"n"`            // 昵称,可选
	ThirdClientId     string `json:"third_client_id" cli:"-"`     // 第三方应用Id
	ThirdClientSecret string `json:"third_client_secret" cli:"-"` // 第三方应用密钥
	ThirdOpenId       string `json:"third_open_id" cli:"-"`       // 第三方用户的openId
	ThirdAccessToken  string `json:"third_access_token" cli:"-"`  // 第三方用户的accessToken

}

func (SignupReq) CommandName() string   { return "Signup" }
func (SignupReq) CommandMethod() string { return "POST" }

var _ = registerCommand(&SignupReq{})

type SignupRes struct {
	User  UserInfo  `json:"user"`
	Token TokenInfo `json:"token"`
}

// 快捷注册
type AutoSignupReq struct {
}

func (AutoSignupReq) CommandName() string   { return "AutoSignup" }
func (AutoSignupReq) CommandMethod() string { return "POST" }

var _ = registerCommand(&AutoSignupReq{})

type AutoSignupRes struct {
	Uid   int64     `json:"uid"`
	Token TokenInfo `json:"token"`
}

// 检查账号是否存在
type AccountExistReq struct {
	Username string `json:"username" cli:"u"`
}

func (AccountExistReq) CommandName() string   { return "AccountExist" }
func (AccountExistReq) CommandMethod() string { return "GET" }

var _ = registerCommand(&AccountExistReq{})

type AccountExistRes struct {
	Existed bool `json:"existed"`
}

// 登陆
type SigninReq struct {
	AccountType int    `json:"account_type" cli:"t"` // 账号类型:参见 model.AccountType 枚举
	Account     string `json:"account" cli:"a"`      // 账号:当 accountType 为第三方账号时为openId
	Password    string `json:"password" cli:"p"`     // 密码:当 accountType 为第三方账号时不需要

}

func (SigninReq) CommandName() string   { return "Signin" }
func (SigninReq) CommandMethod() string { return "POST" }

var _ = registerCommand(&SigninReq{})

type SigninRes struct {
	User  UserInfo  `json:"user"`
	Token TokenInfo `json:"token"`
}

// 登出
type SignoutReq struct {
	Uid int64 `json:"uid" cli:"uid"`
}

func (SignoutReq) CommandName() string   { return "Signout" }
func (SignoutReq) CommandMethod() string { return "POST" }

var _ = registerCommand(&SignoutReq{})

type SignoutRes struct {
}

// 获取手机验证码
type SMSCodeReq struct {
	Telno string `json:"telno" cli:"t"`
}

func (SMSCodeReq) CommandName() string   { return "SMSCode" }
func (SMSCodeReq) CommandMethod() string { return "POST" }

var _ = registerCommand(&SMSCodeReq{})

type SMSCodeRes struct {
}

// 两阶段认证
type TwoFactorAuthReq struct {
	AuthType              string `json:"auth_type" cli:"t"` // telno or email
	AuthId                string `json:"auth_id" cli:"i"`
	AuthCode              string `json:"auth_code" cli:"c"`
	UseThirdVerifyService bool   `json:"use_third_verify_service" cli:"third"`
}

func (TwoFactorAuthReq) CommandName() string   { return "TwoFactorAuth" }
func (TwoFactorAuthReq) CommandMethod() string { return "POST" }

var _ = registerCommand(&TwoFactorAuthReq{})

type TwoFactorAuthRes struct {
	User  UserInfo  `json:"user"`
	Token TokenInfo `json:"token"`
}

// 用户信息
type UserReq struct {
	Uid     int64  `json:"uid" cli:"uid"`
	Account string `json:"account" cli:"a"`
}

func (UserReq) CommandName() string   { return "User" }
func (UserReq) CommandMethod() string { return "GET" }

var _ = registerCommand(&UserReq{})

type UserRes struct {
	User UserInfo `json:"user"`
}
