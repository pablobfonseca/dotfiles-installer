package dotfiles

import (
	"log"

	"github.com/pablobfonseca/dotfiles-cli/src/utils"
	"github.com/vbauerster/mpb/v7"
)

func UpdateNvim(p *mpb.Progress) {
	updateBar := utils.NewBar("Updating nvim packages", 1, p)

	if err := utils.ExecuteCommand("nvim", "+NvChadUpdate", "+qall"); err != nil {
		log.Fatal("Error updating nvim:", err)
	}
	updateBar.Increment()
}
