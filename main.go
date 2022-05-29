package main

import (
	"github.com/biuaxia/fastart/code/core"
	"github.com/biuaxia/fastart/code/support"
	_ "gorm.io/driver/mysql"
)

func main() {
	core.APPLICATION = &support.FartApplication{}
	core.APPLICATION.Start()
}
