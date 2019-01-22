package util

import (
	"fmt"
	"math"
	"net/url"
)

type Pagination struct {
	Uri       string
	Params    map[string]string
	Page      int
	Size      int
	Total     int
	MaxPage   int
	Pages     []int
	Prev      int
	Next      int
	FirstPage int
	LastPage  int
}

func NewPagination(uri string, page, size, total int, params map[string]string) *Pagination {
	pagination := &Pagination{
		Uri:     uri,
		Params:  params,
		Page:    page,
		Size:    size,
		Total:   total,
		MaxPage: int(math.Ceil(float64(total) / float64(size))),
		Prev:    page - 1,
		Next:    page + 1,
	}
	pagination.FirstPage = 1
	pagination.LastPage = pagination.MaxPage
	switch {
	case pagination.Page > pagination.MaxPage-3 && pagination.MaxPage > 5:
		start := pagination.MaxPage - 4
		pagination.Pages = make([]int, 5)
		for i, _ := range pagination.Pages {
			pagination.Pages[i] = start + i
		}
	case page >= 3 && pagination.MaxPage > 5:
		start := page - 3 + 1
		pagination.Pages = make([]int, 5)
		pagination.FirstPage = page - 3
		for i, _ := range pagination.Pages {
			pagination.Pages[i] = start + i
		}
	default:
		pagination.Pages = make([]int, int(math.Min(5, float64(pagination.MaxPage))))
		for i, _ := range pagination.Pages {
			pagination.Pages[i] = i + 1
		}
		pagination.FirstPage = int(math.Max(float64(1), float64(page-1)))
		pagination.LastPage = page + 1
	}
	return pagination
}

func Pager(pagination Pagination) string {
	html := fmt.Sprintf(`
	<label>
	共%d条记录
	</label>
	<div class="btn-group">
	`, pagination.Total)
	uri := pagination.Uri
	params := url.Values{}
	for k, v := range pagination.Params {
		if k == "page" {
			continue
		}
		params.Add(k, v)
	}
	uri += "?" + params.Encode()
	if pagination.Page == 1 {
		html += `
		<a class="btn btn-white disabled"><i class="fa fa-chevron-left"></i></a>
		`
	} else {
		html += fmt.Sprintf(`<a class="btn btn-white" href="%s&page=%d"><i class="fa fa-chevron-left"></i></a>`, uri, pagination.Prev)
	}
	for _, page := range pagination.Pages {
		if page == pagination.Page {
			html += fmt.Sprintf(`<a class="btn btn-white active">%d</a>`, page)
		} else {
			html += fmt.Sprintf(`<a href="%s&page=%d" class="btn btn-white">%d</a>`, uri, page, page)
		}
	}
	if pagination.Page == pagination.MaxPage || pagination.MaxPage == 0 {
		html += `<a class="btn btn-white disabled"><i class="fa fa-chevron-right"></i></a>`
	} else {
		html += fmt.Sprintf(`<a class="btn btn-white" href="%s&page=%d"><i class="fa fa-chevron-right"></i></a>`, uri, pagination.Next)
	}
	html += "</div>"
	return html
}
