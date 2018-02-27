package kindlegen

import (
	"github.com/nickalie/go-binwrapper"
	"strings"
	"github.com/pkg/errors"
)

func Convert(source, target string) error {
	b := createBinWrapper()
	err := b.Run(source, "-o", target)

	if err != nil {
		if strings.Contains(err.Error(), "status 1") {
			return nil
		}

		return errors.Wrap(err, string(b.CombinedOutput()))
	}

	return nil
}

func createBinWrapper() *binwrapper.BinWrapper {
	b := binwrapper.NewBinWrapper()
	b.ExecPath("kindlegen")
	return b
}