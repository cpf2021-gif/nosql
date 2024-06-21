package extract

type Document struct {
	Id          string   `bson:"_id"`
	Cate        string   `bson:"cate"`         // 分类
	Content     string   `bson:"content"`      // 内容
	Ctime       string   `bson:"ctime"`        // 创建时间(时间戳)
	RawKeyWords []string `bson:"raw_keywords"` // 关键词
	Title       string   `bson:"title"`        // 标题
	Url         string   `bson:"url"`          // 新闻链接
}
