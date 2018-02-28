package model

import (
    "errors"
    "fmt"
)

type Category struct {
    ID int
    Title string
}

var categories []Category = []Category{
    Category{ID: 1, Title: "car"},
}

func GetCategories() []Category {
    return categories
}

func GetCategory(ID int) (*Category, error) {
    for _, c := range categories {
        if c.ID == ID {
            return &c, nil
        }
    }
    return nil, errors.New(fmt.Sprintf("Category with ID %v not found", ID))
}
