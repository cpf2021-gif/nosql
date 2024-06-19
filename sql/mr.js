// Province mapping
var provinceMapping = {
    "广东": ["广东省", "广东"],
    "河北": ["河北省", "河北"],
    "山西": ["山西省", "山西"],
    "辽宁": ["辽宁省", "辽宁"],
    "吉林": ["吉林省", "吉林"],
    "黑龙江": ["黑龙江省", "黑龙江"],
    "江苏": ["江苏省", "江苏"],
    "浙江": ["浙江省", "浙江"],
    "安徽": ["安徽省", "安徽"],
    "福建": ["福建省", "福建"],
    "江西": ["江西省", "江西"],
    "山东": ["山东省", "山东"],
    "河南": ["河南省", "河南"],
    "湖北": ["湖北省", "湖北"],
    "湖南": ["湖南省", "湖南"],
    "海南": ["海南省", "海南"],
    "四川": ["四川省", "四川"],
    "贵州": ["贵州省", "贵州"],
    "云南": ["云南省", "云南"],
    "陕西": ["陕西省", "陕西"],
    "甘肃": ["甘肃省", "甘肃"],
    "青海": ["青海省", "青海"],
    "台湾": ["台湾省", "台湾"],
    "北京": ["北京市", "北京"],
    "上海": ["上海市", "上海"],
    "天津": ["天津市", "天津"],
    "重庆": ["重庆市", "重庆"],
    "内蒙古": ["内蒙古自治区", "内蒙古"],
    "广西": ["广西壮族自治区", "广西"],
    "西藏": ["西藏自治区", "西藏"],
    "宁夏": ["宁夏回族自治区", "宁夏"],
    "新疆": ["新疆维吾尔自治区", "新疆"],
    "香港": ["香港特别行政区", "香港"],
    "澳门": ["澳门特别行政区", "澳门"]
};

// Map function
var mapFunction = function() {
    if (this.raw_keywords == null) {
        emit("其他", 1);
        return;
    }

    for (var i = 0; i < this.raw_keywords.length; i++) {
        var keyword = this.raw_keywords[i];
        for (var key in provinceMapping) {
            if (provinceMapping[key].indexOf(keyword) !== -1) {
                emit(key, 1);
                return; // 只需要匹配到一个省份即可
            }
        }
    }
    emit("其他", 1); // 如果省份信息不在映射中，则计为“其他”
};

// Reduce function
var reduceFunction = function(key, values) {
    return Array.sum(values);
};

// Run MapReduce
db.news.mapReduce(
    mapFunction,
    reduceFunction,
    {
        out: "province_news_count", // 输出到一个新的集合
        scope: { provinceMapping: provinceMapping } // 将省份映射传递到Map和Reduce函数中
    }
);

// 查看结果
db.province_news_count.find();
