///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gorm_gen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package cloud_config_repo

import (
	"fmt"
	"time"

	"github.com/haierkeys/obsidian-image-api-gateway/global"
	"github.com/haierkeys/obsidian-image-api-gateway/internal/model"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/timef"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Connection() *gorm.DB {
	db_driver := global.DBEngine
	db_driver.Config.NamingStrategy = schema.NamingStrategy{
		TablePrefix:   "pre_", // 表名前缀
		SingularTable: true,   // 使用单数表名
	}
	return db_driver
}

func NewModel() *CloudConfig {
	return new(CloudConfig)
}

type cloudConfigRepoQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	whereRaw []struct {
		query  string
		values []interface{}
	}
	limit  int
	offset int
}

func NewQueryBuilder() *cloudConfigRepoQueryBuilder {
	return new(cloudConfigRepoQueryBuilder)
}

func (qb *cloudConfigRepoQueryBuilder) buildQuery() *gorm.DB {
	ret := Connection()
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, where2 := range qb.whereRaw {
		ret = ret.Where(where2.query, where2.values...)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (t *CloudConfig) Create() (id int64, err error) {
	t.CreatedAt = timef.Now()
	db_driver := Connection()
	if err = db_driver.Model(t).Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

func (t *CloudConfig) Save() (err error) {
	t.UpdatedAt = timef.Now()

	db_driver := Connection()
	if err = db_driver.Model(t).Save(t).Error; err != nil {
		return errors.Wrap(err, "update err")
	}
	return nil
}

func (qb *cloudConfigRepoQueryBuilder) Updates(m map[string]interface{}) (err error) {

	db_driver := Connection()
	db_driver = db_driver.Model(&CloudConfig{})

	for _, where := range qb.where {
		db_driver.Where(where.prefix, where.value)
	}

	if err = db_driver.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

// 自减
func (qb *cloudConfigRepoQueryBuilder) Increment(column string, value int64) (err error) {

	db_driver := Connection()
	db_driver = db_driver.Model(&CloudConfig{})

	for _, where := range qb.where {
		db_driver.Where(where.prefix, where.value)
	}

	if err = db_driver.Update(column, gorm.Expr(column+" + ?", value)).Error; err != nil {
		return errors.Wrap(err, "increment err")
	}
	return nil
}

// 自增
func (qb *cloudConfigRepoQueryBuilder) Decrement(column string, value int64) (err error) {

	db_driver := Connection()
	db_driver = db_driver.Model(&CloudConfig{})

	for _, where := range qb.where {
		db_driver.Where(where.prefix, where.value)
	}

	if err = db_driver.Update(column, gorm.Expr(column+" - ?", value)).Error; err != nil {
		return errors.Wrap(err, "decrement err")
	}
	return nil
}

func (qb *cloudConfigRepoQueryBuilder) Delete() (err error) {

	db_driver := Connection()
	for _, where := range qb.where {
		db_driver = db_driver.Where(where.prefix, where.value)
	}

	if err = db_driver.Delete(&CloudConfig{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *cloudConfigRepoQueryBuilder) Count() (int64, error) {
	var c int64
	res := qb.buildQuery().Model(&CloudConfig{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *cloudConfigRepoQueryBuilder) First() (*CloudConfig, error) {
	ret := &CloudConfig{}
	res := qb.buildQuery().First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *cloudConfigRepoQueryBuilder) Get() ([]*CloudConfig, error) {
	return qb.QueryAll()
}

func (qb *cloudConfigRepoQueryBuilder) QueryOne() (*CloudConfig, error) {
	qb.limit = 1
	ret, err := qb.QueryAll()
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *cloudConfigRepoQueryBuilder) QueryAll() ([]*CloudConfig, error) {
	var ret []*CloudConfig
	err := qb.buildQuery().Find(&ret).Error
	return ret, err
}

func (qb *cloudConfigRepoQueryBuilder) Limit(limit int) *cloudConfigRepoQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) Offset(offset int) *cloudConfigRepoQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereRaw(query string, values ...interface{}) *cloudConfigRepoQueryBuilder {
	vals := make([]interface{}, len(values))
	for i, v := range values {
		vals[i] = v
	}
	qb.whereRaw = append(qb.whereRaw, struct {
		query  string
		values []interface{}
	}{
		query,
		vals,
	})
	return qb
}

// ----------

func (qb *cloudConfigRepoQueryBuilder) WhereId(p model.Predicate, value int64) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereIdIn(value []int64) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereIdNotIn(value []int64) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) OrderById(asc bool) *cloudConfigRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "`id` "+order)
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereUid(p model.Predicate, value int64) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "uid", p),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereUidIn(value []int64) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "uid", "IN"),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereUidNotIn(value []int64) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "uid", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) OrderByUid(asc bool) *cloudConfigRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "`uid` "+order)
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereType(p model.Predicate, value string) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", p),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereTypeIn(value []string) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "IN"),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereTypeNotIn(value []string) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) OrderByType(asc bool) *cloudConfigRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "`type` "+order)
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereBucketName(p model.Predicate, value string) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "bucket_name", p),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereBucketNameIn(value []string) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "bucket_name", "IN"),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereBucketNameNotIn(value []string) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "bucket_name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) OrderByBucketName(asc bool) *cloudConfigRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "`bucket_name` "+order)
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereAccountId(p model.Predicate, value string) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "account_id", p),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereAccountIdIn(value []string) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "account_id", "IN"),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereAccountIdNotIn(value []string) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "account_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) OrderByAccountId(asc bool) *cloudConfigRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "`account_id` "+order)
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereAccessKeyId(p model.Predicate, value string) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "access_key_id", p),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereAccessKeyIdIn(value []string) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "access_key_id", "IN"),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereAccessKeyIdNotIn(value []string) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "access_key_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) OrderByAccessKeyId(asc bool) *cloudConfigRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "`access_key_id` "+order)
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereAccessKeySecret(p model.Predicate, value string) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "access_key_secret", p),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereAccessKeySecretIn(value []string) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "access_key_secret", "IN"),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereAccessKeySecretNotIn(value []string) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "access_key_secret", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) OrderByAccessKeySecret(asc bool) *cloudConfigRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "`access_key_secret` "+order)
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereCustomPath(p model.Predicate, value string) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "custom_path", p),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereCustomPathIn(value []string) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "custom_path", "IN"),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereCustomPathNotIn(value []string) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "custom_path", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) OrderByCustomPath(asc bool) *cloudConfigRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "`custom_path` "+order)
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereIsDeleted(p model.Predicate, value int64) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", p),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereIsDeletedIn(value []int64) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "IN"),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereIsDeletedNotIn(value []int64) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) OrderByIsDeleted(asc bool) *cloudConfigRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "`is_deleted` "+order)
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereUpdatedAt(p model.Predicate, value time.Time) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereUpdatedAtIn(value []time.Time) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) OrderByUpdatedAt(asc bool) *cloudConfigRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "`updated_at` "+order)
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereCreatedAt(p model.Predicate, value time.Time) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereCreatedAtIn(value []time.Time) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) OrderByCreatedAt(asc bool) *cloudConfigRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "`created_at` "+order)
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereDeletedAt(p model.Predicate, value time.Time) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "deleted_at", p),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereDeletedAtIn(value []time.Time) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "deleted_at", "IN"),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) WhereDeletedAtNotIn(value []time.Time) *cloudConfigRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "deleted_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *cloudConfigRepoQueryBuilder) OrderByDeletedAt(asc bool) *cloudConfigRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "`deleted_at` "+order)
	return qb
}