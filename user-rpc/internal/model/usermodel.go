package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

// User 用户表数据模型
type User struct {
	Id        int64     `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// UserModel 用户模型接口，定义用户数据的数据库操作
type UserModel interface {
	// Insert 新增用户，返回用户 ID
	Insert(ctx context.Context, data *User) (int64, error)
	// FindOne 根据用户 ID 查询用户
	FindOne(ctx context.Context, id int64) (*User, error)
	// FindOneByUsername 根据用户名查询用户（用于登录校验）
	FindOneByUsername(ctx context.Context, username string) (*User, error)
}

type userModel struct {
	conn sqlx.SqlConn
}

// NewUserModel 创建用户模型实例
func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &userModel{conn: conn}
}

// Insert 新增用户，使用 RETURNING id 获取自增主键
func (m *userModel) Insert(ctx context.Context, data *User) (int64, error) {
	query := `INSERT INTO users (username, password, email, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id int64
	err := m.conn.QueryRowCtx(ctx, &id, query,
		data.Username, data.Password, data.Email, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// FindOne 根据用户 ID 查询单个用户信息
func (m *userModel) FindOne(ctx context.Context, id int64) (*User, error) {
	query := `SELECT id, username, password, email, created_at, updated_at
		FROM users WHERE id = $1`
	var user User
	err := m.conn.QueryRowCtx(ctx, &user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindOneByUsername 根据用户名查询用户，登录时通过用户名查找并校验密码
func (m *userModel) FindOneByUsername(ctx context.Context, username string) (*User, error) {
	query := `SELECT id, username, password, email, created_at, updated_at
		FROM users WHERE username = $1`
	var user User
	err := m.conn.QueryRowCtx(ctx, &user, query, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
