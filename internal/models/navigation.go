package models

import "strings"

type NavItem struct {
	Name      string
	Url       string
	Weight    int
	ClassList []string
	Children  []*NavItem
}

type NavList struct {
	NavItems []*NavItem
}

func NewNavItem(name, url string, weight int) *NavItem {
	return &NavItem{
		Name:      name,
		Url:       url,
		Weight:    weight,
		ClassList: []string{},
		Children:  []*NavItem{},
	}
}

func NewNavList() *NavList {
	return &NavList{
		NavItems: []*NavItem{},
	}
}

func (nl *NavList) New(navItem *NavItem) {
	nl.NavItems = append(nl.NavItems, navItem)
}

func (nl *NavList) AddChildrenToItem(name string, navItem ...*NavItem) {
	name = strings.TrimSpace(strings.ToLower(name))
	for _, v := range nl.NavItems {
		n := strings.TrimSpace(strings.ToLower(v.Name))
		if n == name {
			v.Children = append(v.Children, navItem...)
		}
	}
}
