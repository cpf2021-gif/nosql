package extract

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func extractFields(raw string) (Document, bool) {
	res := Document{}
	strs := strings.Split(raw, ",") // 按 "," 切分字段
	length := len(strs)
	// 如果字段数小于 7，说明数据不完整
	if length < 7 {
		return res, false
	}

	res.Id = strs[0]   // id
	res.Cate = strs[1] // cate
	// content中可能使用了逗号 `,`，需要特殊处理
	// ctime - 1684948299 10位
	// 从倒数第3个字段开始，找到第一个长度为10，全为数字的字段，即为ctime
	// ctime 之后的字段为 keywords
	// ctime 之前的字段为 content

	idx := length - 3
	for ; idx >= 3; idx-- {
		if len(strs[idx]) == 10 {
			_, err := strconv.ParseInt(strs[idx], 10, 64)
			if err == nil {
				break
			}
		}
	}

	res.Content = strings.Join(strs[2:idx], ",") // content
	res.Ctime = strs[idx]                        // ctime

	// key_word
	var key_word []string
	for i := idx + 1; i < length-2; i++ {
		if len(strs[i]) == 0 {
			continue
		}
		s := strings.Trim(strs[i], "\"")
		key_word = append(key_word, s)
	}
	res.RawKeyWords = key_word

	res.Title = strs[length-2] // title
	res.Url = strs[length-1]   // urt
	return res, true
}

func GetDocuments(filename string) ([]any, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var res []any
	reader := bufio.NewReader(file)

	// 跳过第一行 字段名
	reader.ReadString('\n')

	for {
		content, err := reader.ReadString('\n') // 读取一行
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}
		content = strings.TrimSpace(content)   // 去除每行末尾的换行符
		document, ok := extractFields(content) // 提取字段
		if ok {
			// 如果数据完整，添加到结果集
			res = append(res, document)
		}

	}
	return res, nil
}
