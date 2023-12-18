package dotfiles

import (
	"log"

	"github.com/pablobfonseca/dotfiles/src/utils"
	"github.com/vbauerster/mpb/v7"
)

func UpdateBrew(p *mpb.Progress, verbose bool) {
	bar := utils.NewBar("Updating brew packages", 1, p)

	if err := utils.ExecuteCommand(verbose, "brew", "update"); err != nil {
		log.Fatal("Error updating homebrew:", err)
		return
	}
	bar.Increment()
}
