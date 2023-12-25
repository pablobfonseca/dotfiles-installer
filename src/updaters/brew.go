package dotfiles

import (
	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/vbauerster/mpb/v7"
)

func UpdateBrew(p *mpb.Progress, verbose bool) {
	bar := utils.NewBar("Updating and upgrading brew packages", 2, p)

	if err := utils.ExecuteCommand(verbose, "brew", "update"); err != nil {
		utils.ErrorMessage("Error updating brew packages", err)
	}
	bar.Increment()

	if err := utils.ExecuteCommand(verbose, "brew", "upgrade"); err != nil {
		utils.ErrorMessage("Error upgrading brew packages", err)
	}

	bar.Increment()
}
