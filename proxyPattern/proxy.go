package main

import "fmt"

type Database interface {
	GetData() string
}

type Postgresql struct{}

func (p *Postgresql) GetData() string {
	return "some data"
}

type Proxy struct {
	db         Database
	cashedData string
}

func (p *Proxy) GetData() string {
	if p.cashedData == "" {
		p.cashedData = p.db.GetData()
	}
	return p.cashedData
}

type MyProgram struct {
	db Database
}

func (m MyProgram) showData() {
	fmt.Println(m.db.GetData())
}

func main() {
	db := &Postgresql{}
	proxy := &Proxy{db: db}
	program := MyProgram{proxy}

	program.showData()
}
