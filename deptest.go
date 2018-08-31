package main

import (
	bystander "github.com/zevdg/deptest-bystander"
	sauce "github.com/zevdg/deptest-oss-sauce"
)

func main() {
	bystander.StatusReport()
	sauce.Dip()
}
