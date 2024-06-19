package extract

type Document struct {
	Id          string   `bson:"_id"`
	Cate        string   `bson:"cate"`
	Content     string   `bson:"content"`
	Ctime       string   `bson:"ctime"`
	RawKeyWords []string `bson:"raw_keywords"`
	Title       string   `bson:"title"`
	Url         string   `bson:"url"`
}
