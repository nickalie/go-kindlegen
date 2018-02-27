package kindlegen

import (
	"testing"
	"os"
	"fmt"
	"net/http"
	"io"
	"github.com/stretchr/testify/assert"
	"archive/zip"
	"path/filepath"
)

func init() {
	downloadFile("https://royallib.com/get/epub/tolstoy_lev/voyna_i_mir_tom_1.zip", "source.zip")
}

func downloadFile(url, target string) {
	_, err := os.Stat(target)

	if err == nil {
		extractEpub(target)
		return
	}

	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error while downloading test file: %v\n", err)
		panic(err)
	}

	defer resp.Body.Close()

	f, err := os.Create(target)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err = io.Copy(f, resp.Body)

	if err != nil {
		panic(err)
	}

	f.Close()
	extractEpub(target)
}

func extractEpub(value string) {
	zipReader, err := zip.OpenReader(value)

	if err != nil {
		panic(err)
	}

	for _, v := range zipReader.File {
		if filepath.Ext(v.Name) != ".epub" {
			continue
		}

		zipToFile(v, "source.epub")
	}
}

func zipToFile(zipFile *zip.File, target string) {
	f, err := os.Create(target)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	zipReader, err := zipFile.Open()

	if err != nil {
		panic(err)
	}

	defer zipReader.Close()

	_, err = io.Copy(f, zipReader)

	if err != nil {
		panic(err)
	}
}

func TestConvert(t *testing.T) {
	err := Convert("source.epub", "target.mobi")
	assert.Nil(t, err)
}

func TestConvertError(t *testing.T) {
	err := Convert("source2.epub", "target3.mobi")
	assert.NotNil(t, err)
}
