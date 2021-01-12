package model

// Category 分类
type Category struct {
	ID         uint   `gorm:"primaryKey;column:id;type:int unsigned;not null"`
	ParentID   int    `gorm:"column:parent_id;type:int;not null"`                      // 分类父id
	Grade      int    `gorm:"column:grade;type:int;not null"`                          // 级别
	Name       string `gorm:"index:index_name;column:name;type:varchar(255);not null"` // 分类名称
	Icon       string `gorm:"column:icon;type:varchar(255);not null"`                  // 分类图标
	Slug       string `gorm:"index:index_slug;column:slug;type:varchar(128);not null"` // 简称
	Type       string `gorm:"column:type;type:varchar(64);not null"`                   // 类型
	Sort       int    `gorm:"column:sort;type:int;not null"`                           // 排序
	RoleID     string `gorm:"column:role_id;type:varchar(64);not null"`                // 角色
	Status     int16  `gorm:"column:status;type:smallint;not null"`                    // 状态
	CreateTime int    `gorm:"column:create_time;type:int"`                             // 创建时间
}

// Cookbook 菜谱
type Cookbook struct {
	ID               int    `gorm:"primaryKey;column:id;type:int;not null"`
	URL              string `gorm:"unique;column:url;type:varchar(255);not null"`             // 原始链接
	Category1        int    `gorm:"column:category1;type:int;not null"`                       // 一级分类
	Category2        int    `gorm:"column:category2;type:int;not null"`                       // 二级分类
	Category3        int    `gorm:"column:category3;type:int;not null"`                       // 三级分类
	CategoryGrade    int    `gorm:"column:category_grade;type:int;not null"`                  // 分类级别:1|2|3
	Category1Name    string `gorm:"column:category1_name;type:varchar(50);not null"`          // 一级分类名称
	Category2Name    string `gorm:"column:category2_name;type:varchar(50);not null"`          // 二级分类名称
	Category3Name    string `gorm:"column:category3_name;type:varchar(50);not null"`          // 三级分类名称
	Author           string `gorm:"column:author;type:varchar(255);not null"`                 // 作者名
	Title            string `gorm:"column:title;type:varchar(255);not null"`                  // 标题
	VideoURL         string `gorm:"column:video_url;type:varchar(255);not null"`              // 视频地址
	Img              string `gorm:"column:img;type:varchar(255);not null"`                    // 菜谱图片
	Tags             string `gorm:"column:tags;type:varchar(255);not null"`                   // 标签
	Fat              int    `gorm:"column:fat;type:int;not null"`                             // 脂肪
	Protein          int    `gorm:"column:protein;type:int;not null"`                         // 蛋白质
	Carbohydrate     int    `gorm:"column:carbohydrate;type:int;not null"`                    // 碳水化合物
	Calorie          int    `gorm:"column:calorie;type:int;not null"`                         // 卡路里
	Level            string `gorm:"column:level;type:varchar(32);not null"`                   // 菜谱级别
	Ingredients      string `gorm:"column:ingredients;type:json;not null"`                    // 食材:名称=量
	IngredientCount  uint8  `gorm:"column:ingredient_count;type:tinyint unsigned;not null"`   // 食材数量
	CookingSteps     string `gorm:"column:cooking_steps;type:json;not null"`                  // 烹饪步骤
	CookingStepCount uint8  `gorm:"column:cooking_step_count;type:tinyint unsigned;not null"` // 烹饪步骤数
	Tips             string `gorm:"column:tips;type:text"`                                    // 小帖士
	CreateTime       int    `gorm:"column:create_time;type:int;not null"`                     // 创建时间
	CookingTimeStr   string `gorm:"column:cooking_time_str;type:varchar(50);not null"`        // 烹饪时长
	CookingTime      int    `gorm:"column:cooking_time;type:int;not null"`                    // 烹饪时长
	CollectCount     int    `gorm:"column:collect_count;type:int;not null"`                   // 收藏数
	ViewCount        int    `gorm:"column:view_count;type:int;not null"`                      // 浏览数
	PublishTime      int    `gorm:"column:publish_time;type:int;not null"`                    // 发布时间
}

// Ingredient 食材
type Ingredient struct {
	ID          int    `gorm:"primaryKey;column:id;type:int;not null"`
	Category1   string `gorm:"column:category1;type:varchar(50);not null"`     // 一级分类
	Category2   string `gorm:"column:category2;type:varchar(50)"`              // 二级分类
	Categories  string `gorm:"column:categories;type:varchar(255);not null"`   // 分类
	CategoryIDs string `gorm:"column:category_ids;type:varchar(100);not null"` // 分类ids
	Link        string `gorm:"column:link;type:varchar(255)"`                  // 链接
	Name        string `gorm:"column:name;type:varchar(30);not null"`          // 名称
	Alias       string `gorm:"column:alias;type:varchar(100);not null"`        // 别名
	Icon        string `gorm:"column:icon;type:varchar(255);not null"`         // 图标
	IsCrawl     int8   `gorm:"column:is_crawl;type:tinyint;not null"`          // 是否爬取:0=否,1=1
	CreateTime  int    `gorm:"column:create_time;type:int;not null"`           // 创建时间
	ViewCount   int    `gorm:"column:view_count;type:int;not null"`            // 浏览数
}
