package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CookbookMgr struct {
	*_BaseMgr
}

// CookbookMgr open func
func CookbookMgr(db *gorm.DB) *_CookbookMgr {
	if db == nil {
		panic(fmt.Errorf("CookbookMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CookbookMgr{_BaseMgr: &_BaseMgr{DB: db.Table("cookbook"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CookbookMgr) GetTableName() string {
	return "cookbook"
}

// Get 获取
func (obj *_CookbookMgr) Get() (result Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CookbookMgr) Gets() (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CookbookMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithURL url获取 原始链接
func (obj *_CookbookMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// WithCategory1 category1获取 一级分类
func (obj *_CookbookMgr) WithCategory1(category1 int) Option {
	return optionFunc(func(o *options) { o.query["category1"] = category1 })
}

// WithCategory2 category2获取 二级分类
func (obj *_CookbookMgr) WithCategory2(category2 int) Option {
	return optionFunc(func(o *options) { o.query["category2"] = category2 })
}

// WithCategory3 category3获取 三级分类
func (obj *_CookbookMgr) WithCategory3(category3 int) Option {
	return optionFunc(func(o *options) { o.query["category3"] = category3 })
}

// WithCategoryGrade category_grade获取 分类级别:1|2|3
func (obj *_CookbookMgr) WithCategoryGrade(categoryGrade int) Option {
	return optionFunc(func(o *options) { o.query["category_grade"] = categoryGrade })
}

// WithCategory1Name category1_name获取 一级分类名称
func (obj *_CookbookMgr) WithCategory1Name(category1Name string) Option {
	return optionFunc(func(o *options) { o.query["category1_name"] = category1Name })
}

// WithCategory2Name category2_name获取 二级分类名称
func (obj *_CookbookMgr) WithCategory2Name(category2Name string) Option {
	return optionFunc(func(o *options) { o.query["category2_name"] = category2Name })
}

// WithCategory3Name category3_name获取 三级分类名称
func (obj *_CookbookMgr) WithCategory3Name(category3Name string) Option {
	return optionFunc(func(o *options) { o.query["category3_name"] = category3Name })
}

// WithAuthor author获取 作者名
func (obj *_CookbookMgr) WithAuthor(author string) Option {
	return optionFunc(func(o *options) { o.query["author"] = author })
}

// WithTitle title获取 标题
func (obj *_CookbookMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithVideoURL video_url获取 视频地址
func (obj *_CookbookMgr) WithVideoURL(videoURL string) Option {
	return optionFunc(func(o *options) { o.query["video_url"] = videoURL })
}

// WithImg img获取 菜谱图片
func (obj *_CookbookMgr) WithImg(img string) Option {
	return optionFunc(func(o *options) { o.query["img"] = img })
}

// WithTags tags获取 标签
func (obj *_CookbookMgr) WithTags(tags string) Option {
	return optionFunc(func(o *options) { o.query["tags"] = tags })
}

// WithFat fat获取 脂肪
func (obj *_CookbookMgr) WithFat(fat int) Option {
	return optionFunc(func(o *options) { o.query["fat"] = fat })
}

// WithProtein protein获取 蛋白质
func (obj *_CookbookMgr) WithProtein(protein int) Option {
	return optionFunc(func(o *options) { o.query["protein"] = protein })
}

// WithCarbohydrate carbohydrate获取 碳水化合物
func (obj *_CookbookMgr) WithCarbohydrate(carbohydrate int) Option {
	return optionFunc(func(o *options) { o.query["carbohydrate"] = carbohydrate })
}

// WithCalorie calorie获取 卡路里
func (obj *_CookbookMgr) WithCalorie(calorie int) Option {
	return optionFunc(func(o *options) { o.query["calorie"] = calorie })
}

// WithLevel level获取 菜谱级别
func (obj *_CookbookMgr) WithLevel(level string) Option {
	return optionFunc(func(o *options) { o.query["level"] = level })
}

// WithIngredients ingredients获取 食材:名称=量
func (obj *_CookbookMgr) WithIngredients(ingredients string) Option {
	return optionFunc(func(o *options) { o.query["ingredients"] = ingredients })
}

// WithIngredientCount ingredient_count获取 食材数量
func (obj *_CookbookMgr) WithIngredientCount(ingredientCount uint8) Option {
	return optionFunc(func(o *options) { o.query["ingredient_count"] = ingredientCount })
}

// WithCookingSteps cooking_steps获取 烹饪步骤
func (obj *_CookbookMgr) WithCookingSteps(cookingSteps string) Option {
	return optionFunc(func(o *options) { o.query["cooking_steps"] = cookingSteps })
}

// WithCookingStepCount cooking_step_count获取 烹饪步骤数
func (obj *_CookbookMgr) WithCookingStepCount(cookingStepCount uint8) Option {
	return optionFunc(func(o *options) { o.query["cooking_step_count"] = cookingStepCount })
}

// WithTips tips获取 小帖士
func (obj *_CookbookMgr) WithTips(tips string) Option {
	return optionFunc(func(o *options) { o.query["tips"] = tips })
}

// WithCreateTime create_time获取 创建时间
func (obj *_CookbookMgr) WithCreateTime(createTime int) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCookingTimeStr cooking_time_str获取 烹饪时长
func (obj *_CookbookMgr) WithCookingTimeStr(cookingTimeStr string) Option {
	return optionFunc(func(o *options) { o.query["cooking_time_str"] = cookingTimeStr })
}

// WithCookingTime cooking_time获取 烹饪时长
func (obj *_CookbookMgr) WithCookingTime(cookingTime int) Option {
	return optionFunc(func(o *options) { o.query["cooking_time"] = cookingTime })
}

// WithCollectCount collect_count获取 收藏数
func (obj *_CookbookMgr) WithCollectCount(collectCount int) Option {
	return optionFunc(func(o *options) { o.query["collect_count"] = collectCount })
}

// WithViewCount view_count获取 浏览数
func (obj *_CookbookMgr) WithViewCount(viewCount int) Option {
	return optionFunc(func(o *options) { o.query["view_count"] = viewCount })
}

// WithPublishTime publish_time获取 发布时间
func (obj *_CookbookMgr) WithPublishTime(publishTime int) Option {
	return optionFunc(func(o *options) { o.query["publish_time"] = publishTime })
}

// GetByOption 功能选项模式获取
func (obj *_CookbookMgr) GetByOption(opts ...Option) (result Cookbook, err error) {
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
func (obj *_CookbookMgr) GetByOptions(opts ...Option) (results []*Cookbook, err error) {
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
func (obj *_CookbookMgr) GetFromID(id int) (result Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_CookbookMgr) GetBatchFromID(ids []int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromURL 通过url获取内容 原始链接
func (obj *_CookbookMgr) GetFromURL(url string) (result Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url = ?", url).Find(&result).Error

	return
}

// GetBatchFromURL 批量唯一主键查找 原始链接
func (obj *_CookbookMgr) GetBatchFromURL(urls []string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url IN (?)", urls).Find(&results).Error

	return
}

// GetFromCategory1 通过category1获取内容 一级分类
func (obj *_CookbookMgr) GetFromCategory1(category1 int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category1 = ?", category1).Find(&results).Error

	return
}

// GetBatchFromCategory1 批量唯一主键查找 一级分类
func (obj *_CookbookMgr) GetBatchFromCategory1(category1s []int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category1 IN (?)", category1s).Find(&results).Error

	return
}

// GetFromCategory2 通过category2获取内容 二级分类
func (obj *_CookbookMgr) GetFromCategory2(category2 int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category2 = ?", category2).Find(&results).Error

	return
}

// GetBatchFromCategory2 批量唯一主键查找 二级分类
func (obj *_CookbookMgr) GetBatchFromCategory2(category2s []int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category2 IN (?)", category2s).Find(&results).Error

	return
}

// GetFromCategory3 通过category3获取内容 三级分类
func (obj *_CookbookMgr) GetFromCategory3(category3 int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category3 = ?", category3).Find(&results).Error

	return
}

// GetBatchFromCategory3 批量唯一主键查找 三级分类
func (obj *_CookbookMgr) GetBatchFromCategory3(category3s []int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category3 IN (?)", category3s).Find(&results).Error

	return
}

// GetFromCategoryGrade 通过category_grade获取内容 分类级别:1|2|3
func (obj *_CookbookMgr) GetFromCategoryGrade(categoryGrade int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category_grade = ?", categoryGrade).Find(&results).Error

	return
}

// GetBatchFromCategoryGrade 批量唯一主键查找 分类级别:1|2|3
func (obj *_CookbookMgr) GetBatchFromCategoryGrade(categoryGrades []int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category_grade IN (?)", categoryGrades).Find(&results).Error

	return
}

// GetFromCategory1Name 通过category1_name获取内容 一级分类名称
func (obj *_CookbookMgr) GetFromCategory1Name(category1Name string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category1_name = ?", category1Name).Find(&results).Error

	return
}

// GetBatchFromCategory1Name 批量唯一主键查找 一级分类名称
func (obj *_CookbookMgr) GetBatchFromCategory1Name(category1Names []string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category1_name IN (?)", category1Names).Find(&results).Error

	return
}

// GetFromCategory2Name 通过category2_name获取内容 二级分类名称
func (obj *_CookbookMgr) GetFromCategory2Name(category2Name string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category2_name = ?", category2Name).Find(&results).Error

	return
}

// GetBatchFromCategory2Name 批量唯一主键查找 二级分类名称
func (obj *_CookbookMgr) GetBatchFromCategory2Name(category2Names []string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category2_name IN (?)", category2Names).Find(&results).Error

	return
}

// GetFromCategory3Name 通过category3_name获取内容 三级分类名称
func (obj *_CookbookMgr) GetFromCategory3Name(category3Name string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category3_name = ?", category3Name).Find(&results).Error

	return
}

// GetBatchFromCategory3Name 批量唯一主键查找 三级分类名称
func (obj *_CookbookMgr) GetBatchFromCategory3Name(category3Names []string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category3_name IN (?)", category3Names).Find(&results).Error

	return
}

// GetFromAuthor 通过author获取内容 作者名
func (obj *_CookbookMgr) GetFromAuthor(author string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("author = ?", author).Find(&results).Error

	return
}

// GetBatchFromAuthor 批量唯一主键查找 作者名
func (obj *_CookbookMgr) GetBatchFromAuthor(authors []string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("author IN (?)", authors).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 标题
func (obj *_CookbookMgr) GetFromTitle(title string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量唯一主键查找 标题
func (obj *_CookbookMgr) GetBatchFromTitle(titles []string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error

	return
}

// GetFromVideoURL 通过video_url获取内容 视频地址
func (obj *_CookbookMgr) GetFromVideoURL(videoURL string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("video_url = ?", videoURL).Find(&results).Error

	return
}

// GetBatchFromVideoURL 批量唯一主键查找 视频地址
func (obj *_CookbookMgr) GetBatchFromVideoURL(videoURLs []string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("video_url IN (?)", videoURLs).Find(&results).Error

	return
}

// GetFromImg 通过img获取内容 菜谱图片
func (obj *_CookbookMgr) GetFromImg(img string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("img = ?", img).Find(&results).Error

	return
}

// GetBatchFromImg 批量唯一主键查找 菜谱图片
func (obj *_CookbookMgr) GetBatchFromImg(imgs []string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("img IN (?)", imgs).Find(&results).Error

	return
}

// GetFromTags 通过tags获取内容 标签
func (obj *_CookbookMgr) GetFromTags(tags string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tags = ?", tags).Find(&results).Error

	return
}

// GetBatchFromTags 批量唯一主键查找 标签
func (obj *_CookbookMgr) GetBatchFromTags(tagss []string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tags IN (?)", tagss).Find(&results).Error

	return
}

// GetFromFat 通过fat获取内容 脂肪
func (obj *_CookbookMgr) GetFromFat(fat int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("fat = ?", fat).Find(&results).Error

	return
}

// GetBatchFromFat 批量唯一主键查找 脂肪
func (obj *_CookbookMgr) GetBatchFromFat(fats []int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("fat IN (?)", fats).Find(&results).Error

	return
}

// GetFromProtein 通过protein获取内容 蛋白质
func (obj *_CookbookMgr) GetFromProtein(protein int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("protein = ?", protein).Find(&results).Error

	return
}

// GetBatchFromProtein 批量唯一主键查找 蛋白质
func (obj *_CookbookMgr) GetBatchFromProtein(proteins []int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("protein IN (?)", proteins).Find(&results).Error

	return
}

// GetFromCarbohydrate 通过carbohydrate获取内容 碳水化合物
func (obj *_CookbookMgr) GetFromCarbohydrate(carbohydrate int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("carbohydrate = ?", carbohydrate).Find(&results).Error

	return
}

// GetBatchFromCarbohydrate 批量唯一主键查找 碳水化合物
func (obj *_CookbookMgr) GetBatchFromCarbohydrate(carbohydrates []int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("carbohydrate IN (?)", carbohydrates).Find(&results).Error

	return
}

// GetFromCalorie 通过calorie获取内容 卡路里
func (obj *_CookbookMgr) GetFromCalorie(calorie int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("calorie = ?", calorie).Find(&results).Error

	return
}

// GetBatchFromCalorie 批量唯一主键查找 卡路里
func (obj *_CookbookMgr) GetBatchFromCalorie(calories []int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("calorie IN (?)", calories).Find(&results).Error

	return
}

// GetFromLevel 通过level获取内容 菜谱级别
func (obj *_CookbookMgr) GetFromLevel(level string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("level = ?", level).Find(&results).Error

	return
}

// GetBatchFromLevel 批量唯一主键查找 菜谱级别
func (obj *_CookbookMgr) GetBatchFromLevel(levels []string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("level IN (?)", levels).Find(&results).Error

	return
}

// GetFromIngredients 通过ingredients获取内容 食材:名称=量
func (obj *_CookbookMgr) GetFromIngredients(ingredients string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ingredients = ?", ingredients).Find(&results).Error

	return
}

// GetBatchFromIngredients 批量唯一主键查找 食材:名称=量
func (obj *_CookbookMgr) GetBatchFromIngredients(ingredientss []string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ingredients IN (?)", ingredientss).Find(&results).Error

	return
}

// GetFromIngredientCount 通过ingredient_count获取内容 食材数量
func (obj *_CookbookMgr) GetFromIngredientCount(ingredientCount uint8) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ingredient_count = ?", ingredientCount).Find(&results).Error

	return
}

// GetBatchFromIngredientCount 批量唯一主键查找 食材数量
func (obj *_CookbookMgr) GetBatchFromIngredientCount(ingredientCounts []uint8) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ingredient_count IN (?)", ingredientCounts).Find(&results).Error

	return
}

// GetFromCookingSteps 通过cooking_steps获取内容 烹饪步骤
func (obj *_CookbookMgr) GetFromCookingSteps(cookingSteps string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cooking_steps = ?", cookingSteps).Find(&results).Error

	return
}

// GetBatchFromCookingSteps 批量唯一主键查找 烹饪步骤
func (obj *_CookbookMgr) GetBatchFromCookingSteps(cookingStepss []string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cooking_steps IN (?)", cookingStepss).Find(&results).Error

	return
}

// GetFromCookingStepCount 通过cooking_step_count获取内容 烹饪步骤数
func (obj *_CookbookMgr) GetFromCookingStepCount(cookingStepCount uint8) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cooking_step_count = ?", cookingStepCount).Find(&results).Error

	return
}

// GetBatchFromCookingStepCount 批量唯一主键查找 烹饪步骤数
func (obj *_CookbookMgr) GetBatchFromCookingStepCount(cookingStepCounts []uint8) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cooking_step_count IN (?)", cookingStepCounts).Find(&results).Error

	return
}

// GetFromTips 通过tips获取内容 小帖士
func (obj *_CookbookMgr) GetFromTips(tips string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tips = ?", tips).Find(&results).Error

	return
}

// GetBatchFromTips 批量唯一主键查找 小帖士
func (obj *_CookbookMgr) GetBatchFromTips(tipss []string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tips IN (?)", tipss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_CookbookMgr) GetFromCreateTime(createTime int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_CookbookMgr) GetBatchFromCreateTime(createTimes []int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCookingTimeStr 通过cooking_time_str获取内容 烹饪时长
func (obj *_CookbookMgr) GetFromCookingTimeStr(cookingTimeStr string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cooking_time_str = ?", cookingTimeStr).Find(&results).Error

	return
}

// GetBatchFromCookingTimeStr 批量唯一主键查找 烹饪时长
func (obj *_CookbookMgr) GetBatchFromCookingTimeStr(cookingTimeStrs []string) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cooking_time_str IN (?)", cookingTimeStrs).Find(&results).Error

	return
}

// GetFromCookingTime 通过cooking_time获取内容 烹饪时长
func (obj *_CookbookMgr) GetFromCookingTime(cookingTime int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cooking_time = ?", cookingTime).Find(&results).Error

	return
}

// GetBatchFromCookingTime 批量唯一主键查找 烹饪时长
func (obj *_CookbookMgr) GetBatchFromCookingTime(cookingTimes []int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cooking_time IN (?)", cookingTimes).Find(&results).Error

	return
}

// GetFromCollectCount 通过collect_count获取内容 收藏数
func (obj *_CookbookMgr) GetFromCollectCount(collectCount int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("collect_count = ?", collectCount).Find(&results).Error

	return
}

// GetBatchFromCollectCount 批量唯一主键查找 收藏数
func (obj *_CookbookMgr) GetBatchFromCollectCount(collectCounts []int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("collect_count IN (?)", collectCounts).Find(&results).Error

	return
}

// GetFromViewCount 通过view_count获取内容 浏览数
func (obj *_CookbookMgr) GetFromViewCount(viewCount int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("view_count = ?", viewCount).Find(&results).Error

	return
}

// GetBatchFromViewCount 批量唯一主键查找 浏览数
func (obj *_CookbookMgr) GetBatchFromViewCount(viewCounts []int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("view_count IN (?)", viewCounts).Find(&results).Error

	return
}

// GetFromPublishTime 通过publish_time获取内容 发布时间
func (obj *_CookbookMgr) GetFromPublishTime(publishTime int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("publish_time = ?", publishTime).Find(&results).Error

	return
}

// GetBatchFromPublishTime 批量唯一主键查找 发布时间
func (obj *_CookbookMgr) GetBatchFromPublishTime(publishTimes []int) (results []*Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("publish_time IN (?)", publishTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CookbookMgr) FetchByPrimaryKey(id int) (result Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByUniqURL primay or index 获取唯一内容
func (obj *_CookbookMgr) FetchUniqueByUniqURL(url string) (result Cookbook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url = ?", url).Find(&result).Error

	return
}
