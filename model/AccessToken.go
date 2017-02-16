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

// 访问令牌
type AccessToken struct {
	Id            int64  `xorm:"pk BIGINT(20)"` // 递增唯一Id
	Uid           int64  `xorm:"BIGINT(20)"`    // 用户Id
	CreatedAt     string `xorm:"TEXT"`          // 创建时间
	ModifiedAt    string `xorm:"TEXT"`          // 修改时间
	ExpireAt      string `xorm:"TEXT"`          // 到期时间
	Token         string `xorm:"TEXT"`          // 令牌
	RefreshToken  string `xorm:"TEXT"`          // 刷新用令牌
	ResourceOwner string `xorm:"TEXT"`          // 资源所有者
	ClientId      string `xorm:"TEXT"`          // 客户Id
	Scope         string `xorm:"TEXT"`          // 可访问权限范围

}

func NewAccessToken() *AccessToken {
	return &AccessToken{}
}

func (AccessToken) Meta() storage.TableMeta { return AccessTokenMetaVar }
func (x AccessToken) Key() interface{}      { return x.Id }
func (x *AccessToken) SetKey(value string) error {
	return typeconv.String2Int64(&x.Id, value)
}

func (x AccessToken) GetField(field string) (interface{}, bool) {
	switch field {
	case AccessTokenMetaVar.F_uid:
		return x.Uid, true
	case AccessTokenMetaVar.F_created_at:
		return x.CreatedAt, true
	case AccessTokenMetaVar.F_modified_at:
		return x.ModifiedAt, true
	case AccessTokenMetaVar.F_expire_at:
		return x.ExpireAt, true
	case AccessTokenMetaVar.F_token:
		return x.Token, true
	case AccessTokenMetaVar.F_refresh_token:
		return x.RefreshToken, true
	case AccessTokenMetaVar.F_resource_owner:
		return x.ResourceOwner, true
	case AccessTokenMetaVar.F_client_id:
		return x.ClientId, true
	case AccessTokenMetaVar.F_scope:
		return x.Scope, true
	}
	return nil, false
}

func (x *AccessToken) SetField(field, value string) error {
	switch field {
	case AccessTokenMetaVar.F_uid:
		return typeconv.String2Int64(&x.Uid, value)
	case AccessTokenMetaVar.F_created_at:
		x.CreatedAt = value
	case AccessTokenMetaVar.F_modified_at:
		x.ModifiedAt = value
	case AccessTokenMetaVar.F_expire_at:
		x.ExpireAt = value
	case AccessTokenMetaVar.F_token:
		x.Token = value
	case AccessTokenMetaVar.F_refresh_token:
		x.RefreshToken = value
	case AccessTokenMetaVar.F_resource_owner:
		x.ResourceOwner = value
	case AccessTokenMetaVar.F_client_id:
		x.ClientId = value
	case AccessTokenMetaVar.F_scope:
		x.Scope = value
	}
	return nil
}

// Meta
type AccessTokenMeta struct {
	F_uid            string
	F_created_at     string
	F_modified_at    string
	F_expire_at      string
	F_token          string
	F_refresh_token  string
	F_resource_owner string
	F_client_id      string
	F_scope          string
}

func (AccessTokenMeta) Name() string     { return "access_token" }
func (AccessTokenMeta) Fields() []string { return _access_token_fields }

var AccessTokenMetaVar = &AccessTokenMeta{
	F_uid:            "uid",
	F_created_at:     "created_at",
	F_modified_at:    "modified_at",
	F_expire_at:      "expire_at",
	F_token:          "token",
	F_refresh_token:  "refresh_token",
	F_resource_owner: "resource_owner",
	F_client_id:      "client_id",
	F_scope:          "scope",
}

var _access_token_fields = []string{
	AccessTokenMetaVar.F_uid,
	AccessTokenMetaVar.F_created_at,
	AccessTokenMetaVar.F_modified_at,
	AccessTokenMetaVar.F_expire_at,
	AccessTokenMetaVar.F_token,
	AccessTokenMetaVar.F_refresh_token,
	AccessTokenMetaVar.F_resource_owner,
	AccessTokenMetaVar.F_client_id,
	AccessTokenMetaVar.F_scope,
}

// Slice
type AccessTokenSlice []AccessToken

func NewAccessTokenSlice(cap int) *AccessTokenSlice {
	s := AccessTokenSlice(make([]AccessToken, 0, cap))
	return &s
}

func (s AccessTokenSlice) Len() int                                  { return len(s) }
func (s AccessTokenSlice) ReadonlyTable(i int) storage.ReadonlyTable { return s[i] }
func (s *AccessTokenSlice) Slice() []AccessToken                     { return []AccessToken(*s) }

func (s *AccessTokenSlice) New(table string, index int, key string) (storage.FieldSetter, error) {
	for len(*s) <= index {
		*s = append(*s, AccessToken{})
	}
	x := &((*s)[index])
	err := x.SetKey(key)
	return x, err
}

// View
type AccessTokenView struct {
	AccessToken
}

type AccessTokenViewSlice []AccessTokenView

func NewAccessTokenViewSlice(cap int) *AccessTokenViewSlice {
	s := AccessTokenViewSlice(make([]AccessTokenView, 0, cap))
	return &s
}

func (s *AccessTokenViewSlice) Slice() []AccessTokenView {
	return []AccessTokenView(*s)
}

func (s *AccessTokenViewSlice) New(table string, index int, key string) (storage.FieldSetter, error) {
	if table == "access_token" {
		for len(*s) <= index {
			x := AccessToken{}
			*s = append(*s, AccessTokenView{AccessToken: x})
		}
		x := &((*s)[index].AccessToken)
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
	AccessTokenViewVar  = AccessTokenView{}
	accessTokenViewRefs = map[string]storage.View{}
)

func (AccessTokenView) Table() string { return AccessTokenMetaVar.Name() }
func (AccessTokenView) Fields() storage.FieldList {
	return storage.FieldSlice(AccessTokenMetaVar.Fields())
}
func (AccessTokenView) Refs() map[string]storage.View { return accessTokenViewRefs }
func (view *AccessTokenView) tables() map[string]storage.WriteonlyTable {
	m := make(map[string]storage.WriteonlyTable)
	m["access_token"] = &view.AccessToken
	return m
}
