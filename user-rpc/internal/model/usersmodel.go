package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		UpdatePassword(ctx context.Context, id int64, passwordHash string) error
		UpdateProfile(ctx context.Context, data *Users) error
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c, opts...),
	}
}

func (m *customUsersModel) UpdateProfile(ctx context.Context, data *Users) error {
	old, err := m.FindOne(ctx, data.Id)
	if err != nil {
		return err
	}

	publicUsersIdKey := fmt.Sprintf("%s%v", cachePublicUsersIdPrefix, old.Id)
	publicUsersUsernameKey := fmt.Sprintf("%s%v", cachePublicUsersUsernamePrefix, old.Username)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		query := fmt.Sprintf("update %s set display_name = $2, avatar = $3, bio = $4, email = $5, website = $6, github = $7, updated_at = now() where id = $1", m.table)
		return conn.ExecCtx(ctx, query, data.Id, data.DisplayName, data.Avatar, data.Bio, data.Email, data.Website, data.Github)
	}, publicUsersIdKey, publicUsersUsernameKey)
	return err
}

func (m *customUsersModel) UpdatePassword(ctx context.Context, id int64, passwordHash string) error {
	old, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	publicUsersIdKey := fmt.Sprintf("%s%v", cachePublicUsersIdPrefix, old.Id)
	publicUsersUsernameKey := fmt.Sprintf("%s%v", cachePublicUsersUsernamePrefix, old.Username)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		query := fmt.Sprintf("update %s set password_hash = $2, updated_at = now() where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id, passwordHash)
	}, publicUsersIdKey, publicUsersUsernameKey)
	return err
}
