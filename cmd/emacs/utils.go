package emacs

import (
	"path"

	"github.com/pablobfonseca/dotfiles/src/utils"
)

func emacsInstalled() bool {
	return utils.DirExists(path.Join("/Applications", "Emacs.app"))
}
