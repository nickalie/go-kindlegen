package kindlegen

import (
	"github.com/nickalie/go-binwrapper"
	"strings"
	"sync"
	"errors"
)

var pool = sync.Pool{
	New: func() interface{} {
		return createBinWrapper()
	},
}

func Convert(source, target string) error {
	b := pool.Get().(*binwrapper.BinWrapper)
	err := b.Run(source, "-o", target)
	b.Reset()
	pool.Put(b)

	if err != nil {
		if strings.Contains(err.Error(), "status 1") {
			return nil
		}

		return errors.New(string(b.CombinedOutput()) + "\n" + err.Error())
	}

	return nil
}

func createBinWrapper() *binwrapper.BinWrapper {
	b := binwrapper.NewBinWrapper()
	b.ExecPath("kindlegen")
	return b
}