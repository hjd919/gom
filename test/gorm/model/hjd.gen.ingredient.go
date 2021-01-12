package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _IngredientMgr struct {
	*_BaseMgr
}

// IngredientMgr open func
func IngredientMgr(db *gorm.DB) *_IngredientMgr {
	if db == nil {
		panic(fmt.Errorf("IngredientMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_IngredientMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ingredient"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_IngredientMgr) GetTableName() string {
	return "ingredient"
}

// Get 获取
func (obj *_IngredientMgr) Get() (result Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_IngredientMgr) Gets() (results []*Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_IngredientMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCategory1 category1获取 一级分类
func (obj *_IngredientMgr) WithCategory1(category1 string) Option {
	return optionFunc(func(o *options) { o.query["category1"] = category1 })
}

// WithCategory2 category2获取 二级分类
func (obj *_IngredientMgr) WithCategory2(category2 string) Option {
	return optionFunc(func(o *options) { o.query["category2"] = category2 })
}

// WithCategories categories获取 分类
func (obj *_IngredientMgr) WithCategories(categories string) Option {
	return optionFunc(func(o *options) { o.query["categories"] = categories })
}

// WithCategoryIDs category_ids获取 分类ids
func (obj *_IngredientMgr) WithCategoryIDs(categoryIDs string) Option {
	return optionFunc(func(o *options) { o.query["category_ids"] = categoryIDs })
}

// WithLink link获取 链接
func (obj *_IngredientMgr) WithLink(link string) Option {
	return optionFunc(func(o *options) { o.query["link"] = link })
}

// WithName name获取 名称
func (obj *_IngredientMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAlias alias获取 别名
func (obj *_IngredientMgr) WithAlias(alias string) Option {
	return optionFunc(func(o *options) { o.query["alias"] = alias })
}

// WithIcon icon获取 图标
func (obj *_IngredientMgr) WithIcon(icon string) Option {
	return optionFunc(func(o *options) { o.query["icon"] = icon })
}

// WithIsCrawl is_crawl获取 是否爬取:0=否,1=1
func (obj *_IngredientMgr) WithIsCrawl(isCrawl int8) Option {
	return optionFunc(func(o *options) { o.query["is_crawl"] = isCrawl })
}

// WithCreateTime create_time获取 创建时间
func (obj *_IngredientMgr) WithCreateTime(createTime int) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithViewCount view_count获取 浏览数
func (obj *_IngredientMgr) WithViewCount(viewCount int) Option {
	return optionFunc(func(o *options) { o.query["view_count"] = viewCount })
}

// GetByOption 功能选项模式获取
func (obj *_IngredientMgr) GetByOption(opts ...Option) (result Ingredient, err error) {
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
func (obj *_IngredientMgr) GetByOptions(opts ...Option) (results []*Ingredient, err error) {
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
func (obj *_IngredientMgr) GetFromID(id int) (result Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_IngredientMgr) GetBatchFromID(ids []int) (results []*Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromCategory1 通过category1获取内容 一级分类
func (obj *_IngredientMgr) GetFromCategory1(category1 string) (results []*Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category1 = ?", category1).Find(&results).Error

	return
}

// GetBatchFromCategory1 批量唯一主键查找 一级分类
func (obj *_IngredientMgr) GetBatchFromCategory1(category1s []string) (results []*Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category1 IN (?)", category1s).Find(&results).Error

	return
}

// GetFromCategory2 通过category2获取内容 二级分类
func (obj *_IngredientMgr) GetFromCategory2(category2 string) (results []*Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category2 = ?", category2).Find(&results).Error

	return
}

// GetBatchFromCategory2 批量唯一主键查找 二级分类
func (obj *_IngredientMgr) GetBatchFromCategory2(category2s []string) (results []*Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category2 IN (?)", category2s).Find(&results).Error

	return
}

// GetFromCategories 通过categories获取内容 分类
func (obj *_IngredientMgr) GetFromCategories(categories string) (results []*Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("categories = ?", categories).Find(&results).Error

	return
}

// GetBatchFromCategories 批量唯一主键查找 分类
func (obj *_IngredientMgr) GetBatchFromCategories(categoriess []string) (results []*Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("categories IN (?)", categoriess).Find(&results).Error

	return
}

// GetFromCategoryIDs 通过category_ids获取内容 分类ids
func (obj *_IngredientMgr) GetFromCategoryIDs(categoryIDs string) (results []*Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category_ids = ?", categoryIDs).Find(&results).Error

	return
}

// GetBatchFromCategoryIDs 批量唯一主键查找 分类ids
func (obj *_IngredientMgr) GetBatchFromCategoryIDs(categoryIDss []string) (results []*Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category_ids IN (?)", categoryIDss).Find(&results).Error

	return
}

// GetFromLink 通过link获取内容 链接
func (obj *_IngredientMgr) GetFromLink(link string) (results []*Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link = ?", link).Find(&results).Error

	return
}

// GetBatchFromLink 批量唯一主键查找 链接
func (obj *_IngredientMgr) GetBatchFromLink(links []string) (results []*Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link IN (?)", links).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_IngredientMgr) GetFromName(name string) (results []*Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 名称
func (obj *_IngredientMgr) GetBatchFromName(names []string) (results []*Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromAlias 通过alias获取内容 别名
func (obj *_IngredientMgr) GetFromAlias(alias string) (results []*Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("alias = ?", alias).Find(&results).Error

	return
}

// GetBatchFromAlias 批量唯一主键查找 别名
func (obj *_IngredientMgr) GetBatchFromAlias(aliass []string) (results []*Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("alias IN (?)", aliass).Find(&results).Error

	return
}

// GetFromIcon 通过icon获取内容 图标
func (obj *_IngredientMgr) GetFromIcon(icon string) (results []*Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon = ?", icon).Find(&results).Error

	return
}

// GetBatchFromIcon 批量唯一主键查找 图标
func (obj *_IngredientMgr) GetBatchFromIcon(icons []string) (results []*Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon IN (?)", icons).Find(&results).Error

	return
}

// GetFromIsCrawl 通过is_crawl获取内容 是否爬取:0=否,1=1
func (obj *_IngredientMgr) GetFromIsCrawl(isCrawl int8) (results []*Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_crawl = ?", isCrawl).Find(&results).Error

	return
}

// GetBatchFromIsCrawl 批量唯一主键查找 是否爬取:0=否,1=1
func (obj *_IngredientMgr) GetBatchFromIsCrawl(isCrawls []int8) (results []*Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_crawl IN (?)", isCrawls).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_IngredientMgr) GetFromCreateTime(createTime int) (results []*Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_IngredientMgr) GetBatchFromCreateTime(createTimes []int) (results []*Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromViewCount 通过view_count获取内容 浏览数
func (obj *_IngredientMgr) GetFromViewCount(viewCount int) (results []*Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("view_count = ?", viewCount).Find(&results).Error

	return
}

// GetBatchFromViewCount 批量唯一主键查找 浏览数
func (obj *_IngredientMgr) GetBatchFromViewCount(viewCounts []int) (results []*Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("view_count IN (?)", viewCounts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_IngredientMgr) FetchByPrimaryKey(id int) (result Ingredient, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
