package dotfiles

import (
	"fmt"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
)

func NewBar(name string, total int64, p *mpb.Progress) *mpb.Bar {
	return p.New(total,
		mpb.BarStyle().Lbound("╢").Filler("▌").Tip("▌").Padding("░").Rbound("╟"),
		mpb.PrependDecorators(
			decor.Name(fmt.Sprintf("%s: ", name), decor.WC{C: decor.DidentRight | decor.DextraSpace}),
			decor.OnComplete(decor.AverageETA(decor.ET_STYLE_GO), "✅"),
		),
		mpb.AppendDecorators(decor.Percentage()),
	)
}
