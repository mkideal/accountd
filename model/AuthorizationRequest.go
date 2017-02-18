// NOTE: AUTO-GENERATED by midc, DON'T edit!!

package model

import (
	"fmt"

	"github.com/midlang/mid/x/go/storage"
	"github.com/mkideal/pkg/typeconv"
)

var (
	_ = fmt.Printf
	_ = storage.Unused
	_ = typeconv.Unused
)

// Table

// 认证请求
type AuthorizationRequest struct {
	AuthorizationCode string `xorm:"pk VARCHAR(64)"` // 认证码
	CreatedAt         string `xorm:"VARCHAR(32)"`    // 创建时间
	RedirectURI       string `xorm:"VARCHAR(256)"`   // 重定向URI
	ResponseType      string `xorm:"VARCHAR(64)"`    // 返回类型
	State             string `xorm:"VARCHAR(128)"`   // 自定义状态
	ClientId          string `xorm:"VARCHAR(64)"`    // 客户端Id
	Uid               int64  `xorm:"BIGINT(20)"`     // 关联的用户Id
	GrantedScopes     string `xorm:"TEXT"`           // 授权范围
	RequestedScopes   string `xorm:"TEXT"`           // 请求范围

}

func NewAuthorizationRequest() *AuthorizationRequest {
	return &AuthorizationRequest{}
}

func (AuthorizationRequest) Meta() storage.TableMeta { return AuthorizationRequestMetaVar }
func (x AuthorizationRequest) Key() interface{}      { return x.AuthorizationCode }
func (x *AuthorizationRequest) SetKey(value string) error {
	x.AuthorizationCode = value
	return nil
}

func (x AuthorizationRequest) GetField(field string) (interface{}, bool) {
	switch field {
	case AuthorizationRequestMetaVar.F_created_at:
		return x.CreatedAt, true
	case AuthorizationRequestMetaVar.F_redirect_uri:
		return x.RedirectURI, true
	case AuthorizationRequestMetaVar.F_response_type:
		return x.ResponseType, true
	case AuthorizationRequestMetaVar.F_state:
		return x.State, true
	case AuthorizationRequestMetaVar.F_client_id:
		return x.ClientId, true
	case AuthorizationRequestMetaVar.F_uid:
		return x.Uid, true
	case AuthorizationRequestMetaVar.F_granted_scopes:
		return x.GrantedScopes, true
	case AuthorizationRequestMetaVar.F_requested_scopes:
		return x.RequestedScopes, true
	}
	return nil, false
}

func (x *AuthorizationRequest) SetField(field, value string) error {
	switch field {
	case AuthorizationRequestMetaVar.F_created_at:
		x.CreatedAt = value
	case AuthorizationRequestMetaVar.F_redirect_uri:
		x.RedirectURI = value
	case AuthorizationRequestMetaVar.F_response_type:
		x.ResponseType = value
	case AuthorizationRequestMetaVar.F_state:
		x.State = value
	case AuthorizationRequestMetaVar.F_client_id:
		x.ClientId = value
	case AuthorizationRequestMetaVar.F_uid:
		return typeconv.String2Int64(&x.Uid, value)
	case AuthorizationRequestMetaVar.F_granted_scopes:
		x.GrantedScopes = value
	case AuthorizationRequestMetaVar.F_requested_scopes:
		x.RequestedScopes = value
	}
	return nil
}

// Meta
type AuthorizationRequestMeta struct {
	F_created_at       string
	F_redirect_uri     string
	F_response_type    string
	F_state            string
	F_client_id        string
	F_uid              string
	F_granted_scopes   string
	F_requested_scopes string
}

func (AuthorizationRequestMeta) Name() string     { return "authorization_request" }
func (AuthorizationRequestMeta) Key() string      { return "authorization_code" }
func (AuthorizationRequestMeta) Fields() []string { return _authorization_request_fields }

var AuthorizationRequestMetaVar = &AuthorizationRequestMeta{
	F_created_at:       "created_at",
	F_redirect_uri:     "redirect_uri",
	F_response_type:    "response_type",
	F_state:            "state",
	F_client_id:        "client_id",
	F_uid:              "uid",
	F_granted_scopes:   "granted_scopes",
	F_requested_scopes: "requested_scopes",
}

var _authorization_request_fields = []string{
	AuthorizationRequestMetaVar.F_created_at,
	AuthorizationRequestMetaVar.F_redirect_uri,
	AuthorizationRequestMetaVar.F_response_type,
	AuthorizationRequestMetaVar.F_state,
	AuthorizationRequestMetaVar.F_client_id,
	AuthorizationRequestMetaVar.F_uid,
	AuthorizationRequestMetaVar.F_granted_scopes,
	AuthorizationRequestMetaVar.F_requested_scopes,
}

// Slice
type AuthorizationRequestSlice []AuthorizationRequest

func NewAuthorizationRequestSlice(cap int) *AuthorizationRequestSlice {
	s := AuthorizationRequestSlice(make([]AuthorizationRequest, 0, cap))
	return &s
}

func (s AuthorizationRequestSlice) Len() int                                  { return len(s) }
func (s AuthorizationRequestSlice) ReadonlyTable(i int) storage.ReadonlyTable { return s[i] }
func (s *AuthorizationRequestSlice) Slice() []AuthorizationRequest            { return []AuthorizationRequest(*s) }

func (s *AuthorizationRequestSlice) New(table string, index int, key string) (storage.FieldSetter, error) {
	for len(*s) <= index {
		*s = append(*s, AuthorizationRequest{})
	}
	x := &((*s)[index])
	err := x.SetKey(key)
	return x, err
}

// View
type AuthorizationRequestView struct {
	AuthorizationRequest
}

type AuthorizationRequestViewSlice []AuthorizationRequestView

func NewAuthorizationRequestViewSlice(cap int) *AuthorizationRequestViewSlice {
	s := AuthorizationRequestViewSlice(make([]AuthorizationRequestView, 0, cap))
	return &s
}

func (s *AuthorizationRequestViewSlice) Slice() []AuthorizationRequestView {
	return []AuthorizationRequestView(*s)
}

func (s *AuthorizationRequestViewSlice) New(table string, index int, key string) (storage.FieldSetter, error) {
	if table == "authorization_request" {
		for len(*s) <= index {
			x := AuthorizationRequest{}
			*s = append(*s, AuthorizationRequestView{AuthorizationRequest: x})
		}
		x := &((*s)[index].AuthorizationRequest)
		err := x.SetKey(key)
		return x, err
	}
	v := &((*s)[index])
	for t, x := range v.tables() {
		if t == table {
			err := x.SetKey(key)
			return x, err
		}
	}
	return nil, storage.ErrTableNotFoundInView
}

var (
	AuthorizationRequestViewVar  = AuthorizationRequestView{}
	authorizationRequestViewRefs = map[string]storage.View{}
)

func (AuthorizationRequestView) Table() string { return AuthorizationRequestMetaVar.Name() }
func (AuthorizationRequestView) Fields() storage.FieldList {
	return storage.FieldSlice(AuthorizationRequestMetaVar.Fields())
}
func (AuthorizationRequestView) Refs() map[string]storage.View { return authorizationRequestViewRefs }
func (view *AuthorizationRequestView) tables() map[string]storage.WriteonlyTable {
	m := make(map[string]storage.WriteonlyTable)
	m["authorization_request"] = &view.AuthorizationRequest
	return m
}
