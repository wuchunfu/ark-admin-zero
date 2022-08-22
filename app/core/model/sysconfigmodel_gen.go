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
	sysConfigFieldNames          = builder.RawFieldNames(&SysConfig{})
	sysConfigRows                = strings.Join(sysConfigFieldNames, ",")
	sysConfigRowsExpectAutoSet   = strings.Join(stringx.Remove(sysConfigFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	sysConfigRowsWithPlaceHolder = strings.Join(stringx.Remove(sysConfigFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheArkAdminSysConfigIdPrefix        = "cache:arkAdmin:sysConfig:id:"
	cacheArkAdminSysConfigUniqueKeyPrefix = "cache:arkAdmin:sysConfig:uniqueKey:"
)

type (
	sysConfigModel interface {
		Insert(ctx context.Context, data *SysConfig) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysConfig, error)
		FindOneByUniqueKey(ctx context.Context, uniqueKey string) (*SysConfig, error)
		Update(ctx context.Context, data *SysConfig) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysConfigModel struct {
		sqlc.CachedConn
		table string
	}

	SysConfig struct {
		Id         int64     `db:"id"`          // 编号
		ParentId   int64     `db:"parent_id"`   // 0=配置集 !0=父级id
		Name       string    `db:"name"`        // 名称
		UniqueKey  string    `db:"unique_key"`  // 唯一值
		Value      string    `db:"value"`       // 配置值
		Status     int64     `db:"status"`      // 0=禁用 1=开启
		OrderNum   int64     `db:"order_num"`   // 排序值
		Remark     string    `db:"remark"`      // 备注
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 更新时间
	}
)

func newSysConfigModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultSysConfigModel {
	return &defaultSysConfigModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_config`",
	}
}

func (m *defaultSysConfigModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	arkAdminSysConfigIdKey := fmt.Sprintf("%s%v", cacheArkAdminSysConfigIdPrefix, id)
	arkAdminSysConfigUniqueKeyKey := fmt.Sprintf("%s%v", cacheArkAdminSysConfigUniqueKeyPrefix, data.UniqueKey)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, arkAdminSysConfigIdKey, arkAdminSysConfigUniqueKeyKey)
	return err
}

func (m *defaultSysConfigModel) FindOne(ctx context.Context, id int64) (*SysConfig, error) {
	arkAdminSysConfigIdKey := fmt.Sprintf("%s%v", cacheArkAdminSysConfigIdPrefix, id)
	var resp SysConfig
	err := m.QueryRowCtx(ctx, &resp, arkAdminSysConfigIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysConfigRows, m.table)
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

func (m *defaultSysConfigModel) FindOneByUniqueKey(ctx context.Context, uniqueKey string) (*SysConfig, error) {
	arkAdminSysConfigUniqueKeyKey := fmt.Sprintf("%s%v", cacheArkAdminSysConfigUniqueKeyPrefix, uniqueKey)
	var resp SysConfig
	err := m.QueryRowIndexCtx(ctx, &resp, arkAdminSysConfigUniqueKeyKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `unique_key` = ? limit 1", sysConfigRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, uniqueKey); err != nil {
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

func (m *defaultSysConfigModel) Insert(ctx context.Context, data *SysConfig) (sql.Result, error) {
	arkAdminSysConfigIdKey := fmt.Sprintf("%s%v", cacheArkAdminSysConfigIdPrefix, data.Id)
	arkAdminSysConfigUniqueKeyKey := fmt.Sprintf("%s%v", cacheArkAdminSysConfigUniqueKeyPrefix, data.UniqueKey)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, sysConfigRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ParentId, data.Name, data.UniqueKey, data.Value, data.Status, data.OrderNum, data.Remark)
	}, arkAdminSysConfigIdKey, arkAdminSysConfigUniqueKeyKey)
	return ret, err
}

func (m *defaultSysConfigModel) Update(ctx context.Context, newData *SysConfig) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	arkAdminSysConfigIdKey := fmt.Sprintf("%s%v", cacheArkAdminSysConfigIdPrefix, data.Id)
	arkAdminSysConfigUniqueKeyKey := fmt.Sprintf("%s%v", cacheArkAdminSysConfigUniqueKeyPrefix, data.UniqueKey)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysConfigRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.ParentId, newData.Name, newData.UniqueKey, newData.Value, newData.Status, newData.OrderNum, newData.Remark, newData.Id)
	}, arkAdminSysConfigIdKey, arkAdminSysConfigUniqueKeyKey)
	return err
}

func (m *defaultSysConfigModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheArkAdminSysConfigIdPrefix, primary)
}

func (m *defaultSysConfigModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysConfigRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSysConfigModel) tableName() string {
	return m.table
}
