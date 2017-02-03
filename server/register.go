package server

import (
	"net/http"

	"github.com/mkideal/pkg/netutil/httputil"
)

func (svr *Server) registerAllHandlers(mux *httputil.ServeMux) {
	svr.registerHandler(mux, "/v1/access_token", "POST", svr.handleAccessToken)
	svr.registerHandler(mux, "/v1/authorize_check", "GET", svr.handleAuthorizeCheck)
	svr.registerHandler(mux, "/v1/authorize", "POST", svr.handleAuthorize)
	svr.registerHandler(mux, "/v1/help", "GET", svr.handleHelp)
	svr.registerHandler(mux, "/v1/logout", "POST", svr.handleLogout)
	svr.registerHandler(mux, "/v1/signin", "POST", svr.handleSignin)
	svr.registerHandler(mux, "/v1/signup", "POST", svr.handleSignup)
	svr.registerHandler(mux, "/v1/token_info", "POST", svr.handleTokenInfo)
	svr.registerHandler(mux, "/v1/user", "GET", svr.handleUser)
}

func (svr *Server) registerHandler(mux *httputil.ServeMux, pattern, method string, h http.HandlerFunc) {
	mux.Handle(pattern, httputil.NewHandlerFunc(method, h))
}

// 查看帮助
type HelpReq struct {
	Version string
	Cmd     string
}

func parseHelp(r *http.Request) (*HelpReq, error) {
	var (
		argv = new(HelpReq)
		err  error
	)
	argv.Version, err = httputil.ParseString(r, "version")
	if err != nil {
		return nil, err
	}
	argv.Cmd, err = httputil.ParseString(r, "cmd")
	if err != nil {
		return nil, err
	}

	return argv, err
}

type HelpRes struct {
}

// oauth2.0 接口 access_token
type AccessTokenReq struct {
	GrantType    string // 授权类型
	Code         string // 授权码
	RedirectURI  string // 重定向URI
	Username     string // 账号
	Password     string // 密码
	RefreshToken string // 刷新用token

}

func parseAccessToken(r *http.Request) (*AccessTokenReq, error) {
	var (
		argv = new(AccessTokenReq)
		err  error
	)
	argv.GrantType, err = httputil.ParseString(r, "grant_type")
	if err != nil {
		return nil, err
	}
	argv.Code, err = httputil.ParseString(r, "code")
	if err != nil {
		return nil, err
	}
	argv.RedirectURI, err = httputil.ParseString(r, "redirect_uri")
	if err != nil {
		return nil, err
	}
	argv.Username, err = httputil.ParseString(r, "username")
	if err != nil {
		return nil, err
	}
	argv.Password, err = httputil.ParseString(r, "password")
	if err != nil {
		return nil, err
	}
	argv.RefreshToken, err = httputil.ParseString(r, "refresh_token")
	if err != nil {
		return nil, err
	}

	return argv, err
}

type AccessTokenRes struct {
}

// oauth2.0 接口 authorize
type AuthorizeReq struct {
	ClientId     string
	ResponseType string
	Scope        string
	State        string
}

func parseAuthorize(r *http.Request) (*AuthorizeReq, error) {
	var (
		argv = new(AuthorizeReq)
		err  error
	)
	argv.ClientId, err = httputil.ParseString(r, "client_id")
	if err != nil {
		return nil, err
	}
	argv.ResponseType, err = httputil.ParseString(r, "response_type")
	if err != nil {
		return nil, err
	}
	argv.Scope, err = httputil.ParseString(r, "scope")
	if err != nil {
		return nil, err
	}
	argv.State, err = httputil.ParseString(r, "state")
	if err != nil {
		return nil, err
	}

	return argv, err
}

type AuthorizeRes struct {
}

// client 检查
type AuthorizeCheckReq struct {
	ClientId string
}

func parseAuthorizeCheck(r *http.Request) (*AuthorizeCheckReq, error) {
	var (
		argv = new(AuthorizeCheckReq)
		err  error
	)
	argv.ClientId, err = httputil.ParseString(r, "client_id")
	if err != nil {
		return nil, err
	}

	return argv, err
}

type AuthorizeCheckRes struct {
}

// 获取 token 信息
type TokenInfoReq struct {
	AccessToken string
}

func parseTokenInfo(r *http.Request) (*TokenInfoReq, error) {
	var (
		argv = new(TokenInfoReq)
		err  error
	)
	argv.AccessToken, err = httputil.ParseString(r, "access_token")
	if err != nil {
		return nil, err
	}

	return argv, err
}

type TokenInfoRes struct {
}

// 注册
type SignupReq struct {
}

func parseSignup(r *http.Request) (*SignupReq, error) {
	var (
		argv = new(SignupReq)
		err  error
	)

	return argv, err
}

// 登陆
type SigninReq struct {
}

func parseSignin(r *http.Request) (*SigninReq, error) {
	var (
		argv = new(SigninReq)
		err  error
	)

	return argv, err
}

// 登出
type LogoutReq struct {
}

func parseLogout(r *http.Request) (*LogoutReq, error) {
	var (
		argv = new(LogoutReq)
		err  error
	)

	return argv, err
}

// 用户信息
type UserReq struct {
}

func parseUser(r *http.Request) (*UserReq, error) {
	var (
		argv = new(UserReq)
		err  error
	)

	return argv, err
}
