package model

import "strconv"

//Bookdetail : book details struct
type Bookdetail struct {
	Author    string
	Publisher string
	Bookpages int
	Price     string
	Score     string
	Info      string
}

func (b Bookdetail) String() string {
	return "作者:" + b.Author + "--出版社:" + b.Publisher + "--页数:" + strconv.Itoa(b.Bookpages) + "--价格:" + b.Price + "--评分:" + b.Score + "--简介:" + b.Info
}
