package controller

type query struct {
    Term string `json: "term"`
    Page int `json: "page"`
    PageSize int `json: "pageSize"`
}
