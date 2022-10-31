package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Search Issues 函数查询 Github 的 issue 跟踪接口
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q, _ := url.QueryUnescape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// 我们必须在所有的可能分支上面关闭resp.Body
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil { // 流式解码器Decode
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
