package dotfiles

import (
	"fmt"

	utils "github.com/pablobfonseca/dotfiles-cli/src"
	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
)

func UpdateNvim(p *mpb.Progress) {
	updateBar := p.AddBar(100,
		mpb.PrependDecorators(
			decor.Name("Updating nvim", decor.WC{W: len("Updating nvim") + 1, C: decor.DidentRight}),
		),
		mpb.AppendDecorators(
			decor.OnComplete(decor.EwmaETA(decor.ET_STYLE_GO, 60), "done"),
		),
	)

	if err := utils.ExecuteCommand("nvim", "+NvChadUpdate", "+qall"); err != nil {
		fmt.Println("Error updating nvim:", err)
		return
	}
	updateBar.Increment()
}
