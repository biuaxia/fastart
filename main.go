package main

import (
	"github.com/biuaxia/fart/code/core"
	"github.com/biuaxia/fart/code/support"
	_ "gorm.io/driver/mysql"
)

func main() {
	core.APPLICATION = &support.FartApplication{}
	core.APPLICATION.Start()
}
