// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysUserFieldNames          = builder.RawFieldNames(&SysUser{})
	sysUserRows                = strings.Join(sysUserFieldNames, ",")
	sysUserRowsExpectAutoSet   = strings.Join(stringx.Remove(sysUserFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	sysUserRowsWithPlaceHolder = strings.Join(stringx.Remove(sysUserFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheArkAdminZeroSysUserIdPrefix      = "cache:arkAdminZero:sysUser:id:"
	cacheArkAdminZeroSysUserAccountPrefix = "cache:arkAdminZero:sysUser:account:"
)

type (
	sysUserModel interface {
		Insert(ctx context.Context, data *SysUser) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysUser, error)
		FindOneByAccount(ctx context.Context, account string) (*SysUser, error)
		Update(ctx context.Context, data *SysUser) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysUserModel struct {
		sqlc.CachedConn
		table string
	}

	SysUser struct {
		Id           int64     `db:"id"`            // 编号
		Account      string    `db:"account"`       // 账号
		Password     string    `db:"password"`      // 密码
		Username     string    `db:"username"`      // 姓名
		Nickname     string    `db:"nickname"`      // 昵称
		Avatar       string    `db:"avatar"`        // 头像
		Gender       int64     `db:"gender"`        // 0=保密 1=女 2=男
		Birthday     time.Time `db:"birthday"`      // 生日
		Email        string    `db:"email"`         // 邮件
		Mobile       string    `db:"mobile"`        // 手机号
		ProfessionId int64     `db:"profession_id"` // 职称
		JobId        int64     `db:"job_id"`        // 岗位
		DeptId       int64     `db:"dept_id"`       // 部门
		RoleIds      string    `db:"role_ids"`      // 角色集
		Status       int64     `db:"status"`        // 0=禁用 1=开启
		OrderNum     int64     `db:"order_num"`     // 排序值
		Remark       string    `db:"remark"`        // 备注
		CreateTime   time.Time `db:"create_time"`   // 创建时间
		UpdateTime   time.Time `db:"update_time"`   // 更新时间
	}
)

func newSysUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultSysUserModel {
	return &defaultSysUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_user`",
	}
}

func (m *defaultSysUserModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	arkAdminZeroSysUserAccountKey := fmt.Sprintf("%s%v", cacheArkAdminZeroSysUserAccountPrefix, data.Account)
	arkAdminZeroSysUserIdKey := fmt.Sprintf("%s%v", cacheArkAdminZeroSysUserIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, arkAdminZeroSysUserAccountKey, arkAdminZeroSysUserIdKey)
	return err
}

func (m *defaultSysUserModel) FindOne(ctx context.Context, id int64) (*SysUser, error) {
	arkAdminZeroSysUserIdKey := fmt.Sprintf("%s%v", cacheArkAdminZeroSysUserIdPrefix, id)
	var resp SysUser
	err := m.QueryRowCtx(ctx, &resp, arkAdminZeroSysUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysUserRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysUserModel) FindOneByAccount(ctx context.Context, account string) (*SysUser, error) {
	arkAdminZeroSysUserAccountKey := fmt.Sprintf("%s%v", cacheArkAdminZeroSysUserAccountPrefix, account)
	var resp SysUser
	err := m.QueryRowIndexCtx(ctx, &resp, arkAdminZeroSysUserAccountKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `account` = ? limit 1", sysUserRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, account); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysUserModel) Insert(ctx context.Context, data *SysUser) (sql.Result, error) {
	arkAdminZeroSysUserAccountKey := fmt.Sprintf("%s%v", cacheArkAdminZeroSysUserAccountPrefix, data.Account)
	arkAdminZeroSysUserIdKey := fmt.Sprintf("%s%v", cacheArkAdminZeroSysUserIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysUserRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Account, data.Password, data.Username, data.Nickname, data.Avatar, data.Gender, data.Birthday, data.Email, data.Mobile, data.ProfessionId, data.JobId, data.DeptId, data.RoleIds, data.Status, data.OrderNum, data.Remark)
	}, arkAdminZeroSysUserAccountKey, arkAdminZeroSysUserIdKey)
	return ret, err
}

func (m *defaultSysUserModel) Update(ctx context.Context, newData *SysUser) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	arkAdminZeroSysUserAccountKey := fmt.Sprintf("%s%v", cacheArkAdminZeroSysUserAccountPrefix, data.Account)
	arkAdminZeroSysUserIdKey := fmt.Sprintf("%s%v", cacheArkAdminZeroSysUserIdPrefix, data.Id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysUserRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Account, newData.Password, newData.Username, newData.Nickname, newData.Avatar, newData.Gender, newData.Birthday, newData.Email, newData.Mobile, newData.ProfessionId, newData.JobId, newData.DeptId, newData.RoleIds, newData.Status, newData.OrderNum, newData.Remark, newData.Id)
	}, arkAdminZeroSysUserAccountKey, arkAdminZeroSysUserIdKey)
	return err
}

func (m *defaultSysUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheArkAdminZeroSysUserIdPrefix, primary)
}

func (m *defaultSysUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysUserRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSysUserModel) tableName() string {
	return m.table
}