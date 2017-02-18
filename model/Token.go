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
type Token struct {
	AccessToken   string `xorm:"pk VARCHAR(64) UNIQUE"` // 令牌
	Uid           int64  `xorm:"BIGINT(20)"`            // 用户Id
	CreatedAt     string `xorm:"VARCHAR(32)"`           // 创建时间
	ModifiedAt    string `xorm:"VARCHAR(32)"`           // 修改时间
	ExpireAt      string `xorm:"VARCHAR(32)"`           // 到期时间
	RefreshToken  string `xorm:"VARCHAR(64) UNIQUE"`    // 刷新用令牌
	ResourceOwner string `xorm:"VARCHAR(32)"`           // 资源所有者
	ClientId      string `xorm:"VARCHAR(32)"`           // 客户Id
	Scope         string `xorm:"TEXT"`                  // 可访问权限范围

}

func NewToken() *Token {
	return &Token{}
}

func (Token) Meta() storage.TableMeta { return TokenMetaVar }
func (x Token) Key() interface{}      { return x.AccessToken }
func (x *Token) SetKey(value string) error {
	x.AccessToken = value
	return nil
}

func (x Token) GetField(field string) (interface{}, bool) {
	switch field {
	case TokenMetaVar.F_uid:
		return x.Uid, true
	case TokenMetaVar.F_created_at:
		return x.CreatedAt, true
	case TokenMetaVar.F_modified_at:
		return x.ModifiedAt, true
	case TokenMetaVar.F_expire_at:
		return x.ExpireAt, true
	case TokenMetaVar.F_refresh_token:
		return x.RefreshToken, true
	case TokenMetaVar.F_resource_owner:
		return x.ResourceOwner, true
	case TokenMetaVar.F_client_id:
		return x.ClientId, true
	case TokenMetaVar.F_scope:
		return x.Scope, true
	}
	return nil, false
}

func (x *Token) SetField(field, value string) error {
	switch field {
	case TokenMetaVar.F_uid:
		return typeconv.String2Int64(&x.Uid, value)
	case TokenMetaVar.F_created_at:
		x.CreatedAt = value
	case TokenMetaVar.F_modified_at:
		x.ModifiedAt = value
	case TokenMetaVar.F_expire_at:
		x.ExpireAt = value
	case TokenMetaVar.F_refresh_token:
		x.RefreshToken = value
	case TokenMetaVar.F_resource_owner:
		x.ResourceOwner = value
	case TokenMetaVar.F_client_id:
		x.ClientId = value
	case TokenMetaVar.F_scope:
		x.Scope = value
	}
	return nil
}

// Meta
type TokenMeta struct {
	F_uid            string
	F_created_at     string
	F_modified_at    string
	F_expire_at      string
	F_refresh_token  string
	F_resource_owner string
	F_client_id      string
	F_scope          string
}

func (TokenMeta) Name() string     { return "token" }
func (TokenMeta) Key() string      { return "access_token" }
func (TokenMeta) Fields() []string { return _token_fields }

var TokenMetaVar = &TokenMeta{
	F_uid:            "uid",
	F_created_at:     "created_at",
	F_modified_at:    "modified_at",
	F_expire_at:      "expire_at",
	F_refresh_token:  "refresh_token",
	F_resource_owner: "resource_owner",
	F_client_id:      "client_id",
	F_scope:          "scope",
}

var _token_fields = []string{
	TokenMetaVar.F_uid,
	TokenMetaVar.F_created_at,
	TokenMetaVar.F_modified_at,
	TokenMetaVar.F_expire_at,
	TokenMetaVar.F_refresh_token,
	TokenMetaVar.F_resource_owner,
	TokenMetaVar.F_client_id,
	TokenMetaVar.F_scope,
}

// Slice
type TokenSlice []Token

func NewTokenSlice(cap int) *TokenSlice {
	s := TokenSlice(make([]Token, 0, cap))
	return &s
}

func (s TokenSlice) Len() int                                  { return len(s) }
func (s TokenSlice) ReadonlyTable(i int) storage.ReadonlyTable { return s[i] }
func (s *TokenSlice) Slice() []Token                           { return []Token(*s) }

func (s *TokenSlice) New(table string, index int, key string) (storage.FieldSetter, error) {
	for len(*s) <= index {
		*s = append(*s, Token{})
	}
	x := &((*s)[index])
	err := x.SetKey(key)
	return x, err
}

// View
type TokenView struct {
	Token
}

type TokenViewSlice []TokenView

func NewTokenViewSlice(cap int) *TokenViewSlice {
	s := TokenViewSlice(make([]TokenView, 0, cap))
	return &s
}

func (s *TokenViewSlice) Slice() []TokenView {
	return []TokenView(*s)
}

func (s *TokenViewSlice) New(table string, index int, key string) (storage.FieldSetter, error) {
	if table == "token" {
		for len(*s) <= index {
			x := Token{}
			*s = append(*s, TokenView{Token: x})
		}
		x := &((*s)[index].Token)
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
	TokenViewVar  = TokenView{}
	tokenViewRefs = map[string]storage.View{}
)

func (TokenView) Table() string                 { return TokenMetaVar.Name() }
func (TokenView) Fields() storage.FieldList     { return storage.FieldSlice(TokenMetaVar.Fields()) }
func (TokenView) Refs() map[string]storage.View { return tokenViewRefs }
func (view *TokenView) tables() map[string]storage.WriteonlyTable {
	m := make(map[string]storage.WriteonlyTable)
	m["token"] = &view.Token
	return m
}
