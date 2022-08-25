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
	sysPermMenuFieldNames          = builder.RawFieldNames(&SysPermMenu{})
	sysPermMenuRows                = strings.Join(sysPermMenuFieldNames, ",")
	sysPermMenuRowsExpectAutoSet   = strings.Join(stringx.Remove(sysPermMenuFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	sysPermMenuRowsWithPlaceHolder = strings.Join(stringx.Remove(sysPermMenuFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheArkAdminSysPermMenuIdPrefix = "cache:arkAdmin:sysPermMenu:id:"
)

type (
	sysPermMenuModel interface {
		Insert(ctx context.Context, data *SysPermMenu) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*SysPermMenu, error)
		Update(ctx context.Context, data *SysPermMenu) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultSysPermMenuModel struct {
		sqlc.CachedConn
		table string
	}

	SysPermMenu struct {
		Id           uint64    `db:"id"`            // 编号
		ParentId     uint64    `db:"parent_id"`     // 父级id
		Name         string    `db:"name"`          // 名称
		Router       string    `db:"router"`        // 路由
		Perms        string    `db:"perms"`         // 权限
		Type         uint64    `db:"type"`          // 0=目录 1=菜单 2=权限
		Icon         string    `db:"icon"`          // 图标
		OrderNum     uint64    `db:"order_num"`     // 排序值
		ViewPath     string    `db:"view_path"`     // 页面路径
		IsShow       uint64    `db:"is_show"`       // 0=隐藏 1=显示
		ActiveRouter string    `db:"active_router"` // 当前激活的菜单
		CreateTime   time.Time `db:"create_time"`   // 创建时间
		UpdateTime   time.Time `db:"update_time"`   // 更新时间
	}
)

func newSysPermMenuModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultSysPermMenuModel {
	return &defaultSysPermMenuModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_perm_menu`",
	}
}

func (m *defaultSysPermMenuModel) Delete(ctx context.Context, id uint64) error {
	arkAdminSysPermMenuIdKey := fmt.Sprintf("%s%v", cacheArkAdminSysPermMenuIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, arkAdminSysPermMenuIdKey)
	return err
}

func (m *defaultSysPermMenuModel) FindOne(ctx context.Context, id uint64) (*SysPermMenu, error) {
	arkAdminSysPermMenuIdKey := fmt.Sprintf("%s%v", cacheArkAdminSysPermMenuIdPrefix, id)
	var resp SysPermMenu
	err := m.QueryRowCtx(ctx, &resp, arkAdminSysPermMenuIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysPermMenuRows, m.table)
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

func (m *defaultSysPermMenuModel) Insert(ctx context.Context, data *SysPermMenu) (sql.Result, error) {
	arkAdminSysPermMenuIdKey := fmt.Sprintf("%s%v", cacheArkAdminSysPermMenuIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysPermMenuRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ParentId, data.Name, data.Router, data.Perms, data.Type, data.Icon, data.OrderNum, data.ViewPath, data.IsShow, data.ActiveRouter)
	}, arkAdminSysPermMenuIdKey)
	return ret, err
}

func (m *defaultSysPermMenuModel) Update(ctx context.Context, data *SysPermMenu) error {
	arkAdminSysPermMenuIdKey := fmt.Sprintf("%s%v", cacheArkAdminSysPermMenuIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysPermMenuRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.ParentId, data.Name, data.Router, data.Perms, data.Type, data.Icon, data.OrderNum, data.ViewPath, data.IsShow, data.ActiveRouter, data.Id)
	}, arkAdminSysPermMenuIdKey)
	return err
}

func (m *defaultSysPermMenuModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheArkAdminSysPermMenuIdPrefix, primary)
}

func (m *defaultSysPermMenuModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysPermMenuRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSysPermMenuModel) tableName() string {
	return m.table
}
