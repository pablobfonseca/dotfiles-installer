package dotfiles

import (
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/vbauerster/mpb/v7"
)

func UpdateNvim(p *mpb.Progress, verbose bool) {
	updateBar := utils.NewBar("Updating nvim packages", 1, p)

	if err := utils.ExecuteCommand(verbose, "nvim", "+NvChadUpdate", "+qall"); err != nil {
		utils.ErrorMessage("Error updating nvim", err)
	}
	updateBar.Increment()
}
