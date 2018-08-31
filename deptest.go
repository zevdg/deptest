package deptest

import (
	bystander "github.com/zevdg/deptest-bystander"
	sauce "github.com/zevdg/deptest-oss-sauce"
)

func Experiment() {
	bystander.StatusReport()
	sauce.Dip()
}
