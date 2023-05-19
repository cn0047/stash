package main

import (
	"fmt"
)

type Cursor interface {
	New() Cursor
	GetLimit() uint64
}

type LimitCursor struct {
	Limit    uint64     `json:"-"`
}

func (c *LimitCursor) GetLimit() uint64 {
	return c.Limit
}

func (c *LimitCursor) New() Cursor {
	return &LimitCursor{Limit: 10}
}

type Paginator interface {
	NewCursor() (Cursor, error)
}

type Base64Paginator[T Cursor] struct {
	ExcludedField   string
	DefaultExcluded any
}

func (p *Base64Paginator[T]) NewCursor() (Cursor, error) {
	var c T
	cursor := c.New()

	return cursor, nil
}

func newLimitPaginator(excludedField string, defaultExcluded any) *Base64Paginator[*LimitCursor] {
	return &Base64Paginator[*LimitCursor]{
		ExcludedField:   excludedField,
		DefaultExcluded: defaultExcluded,
	}
}

func main() {
	p := newLimitPaginator("id", 0)
	c, err := p.NewCursor()
	fmt.Printf("err: %v \n", err)
	fmt.Printf("c: %v \n", c)
}
