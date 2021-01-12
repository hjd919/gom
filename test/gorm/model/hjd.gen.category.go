package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CategoryMgr struct {
	*_BaseMgr
}

// CategoryMgr open func
func CategoryMgr(db *gorm.DB) *_CategoryMgr {
	if db == nil {
		panic(fmt.Errorf("CategoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CategoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("category"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CategoryMgr) GetTableName() string {
	return "category"
}

// Get 获取
func (obj *_CategoryMgr) Get() (result Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CategoryMgr) Gets() (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CategoryMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithParentID parent_id获取 分类父id
func (obj *_CategoryMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithGrade grade获取 级别
func (obj *_CategoryMgr) WithGrade(grade int) Option {
	return optionFunc(func(o *options) { o.query["grade"] = grade })
}

// WithName name获取 分类名称
func (obj *_CategoryMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithIcon icon获取 分类图标
func (obj *_CategoryMgr) WithIcon(icon string) Option {
	return optionFunc(func(o *options) { o.query["icon"] = icon })
}

// WithSlug slug获取 简称
func (obj *_CategoryMgr) WithSlug(slug string) Option {
	return optionFunc(func(o *options) { o.query["slug"] = slug })
}

// WithType type获取 类型
func (obj *_CategoryMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithSort sort获取 排序
func (obj *_CategoryMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithRoleID role_id获取 角色
func (obj *_CategoryMgr) WithRoleID(roleID string) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// WithStatus status获取 状态
func (obj *_CategoryMgr) WithStatus(status int16) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateTime create_time获取 创建时间
func (obj *_CategoryMgr) WithCreateTime(createTime int) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// GetByOption 功能选项模式获取
func (obj *_CategoryMgr) GetByOption(opts ...Option) (result Category, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CategoryMgr) GetByOptions(opts ...Option) (results []*Category, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_CategoryMgr) GetFromID(id uint) (result Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_CategoryMgr) GetBatchFromID(ids []uint) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 分类父id
func (obj *_CategoryMgr) GetFromParentID(parentID int) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("parent_id = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量唯一主键查找 分类父id
func (obj *_CategoryMgr) GetBatchFromParentID(parentIDs []int) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("parent_id IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromGrade 通过grade获取内容 级别
func (obj *_CategoryMgr) GetFromGrade(grade int) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("grade = ?", grade).Find(&results).Error

	return
}

// GetBatchFromGrade 批量唯一主键查找 级别
func (obj *_CategoryMgr) GetBatchFromGrade(grades []int) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("grade IN (?)", grades).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 分类名称
func (obj *_CategoryMgr) GetFromName(name string) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 分类名称
func (obj *_CategoryMgr) GetBatchFromName(names []string) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromIcon 通过icon获取内容 分类图标
func (obj *_CategoryMgr) GetFromIcon(icon string) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon = ?", icon).Find(&results).Error

	return
}

// GetBatchFromIcon 批量唯一主键查找 分类图标
func (obj *_CategoryMgr) GetBatchFromIcon(icons []string) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon IN (?)", icons).Find(&results).Error

	return
}

// GetFromSlug 通过slug获取内容 简称
func (obj *_CategoryMgr) GetFromSlug(slug string) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("slug = ?", slug).Find(&results).Error

	return
}

// GetBatchFromSlug 批量唯一主键查找 简称
func (obj *_CategoryMgr) GetBatchFromSlug(slugs []string) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("slug IN (?)", slugs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 类型
func (obj *_CategoryMgr) GetFromType(_type string) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 类型
func (obj *_CategoryMgr) GetBatchFromType(_types []string) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_CategoryMgr) GetFromSort(sort int) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sort = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量唯一主键查找 排序
func (obj *_CategoryMgr) GetBatchFromSort(sorts []int) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sort IN (?)", sorts).Find(&results).Error

	return
}

// GetFromRoleID 通过role_id获取内容 角色
func (obj *_CategoryMgr) GetFromRoleID(roleID string) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("role_id = ?", roleID).Find(&results).Error

	return
}

// GetBatchFromRoleID 批量唯一主键查找 角色
func (obj *_CategoryMgr) GetBatchFromRoleID(roleIDs []string) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("role_id IN (?)", roleIDs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_CategoryMgr) GetFromStatus(status int16) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 状态
func (obj *_CategoryMgr) GetBatchFromStatus(statuss []int16) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_CategoryMgr) GetFromCreateTime(createTime int) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_CategoryMgr) GetBatchFromCreateTime(createTimes []int) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CategoryMgr) FetchByPrimaryKey(id uint) (result Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByIndexName  获取多个内容
func (obj *_CategoryMgr) FetchIndexByIndexName(name string) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// FetchIndexByIndexSlug  获取多个内容
func (obj *_CategoryMgr) FetchIndexByIndexSlug(slug string) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("slug = ?", slug).Find(&results).Error

	return
}
