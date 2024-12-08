package obsidian

import (
	"testing"
)

func TestIngoreFiles(t *testing.T) {

	ignoreBlobs := getIgnoredFiles("/Users/stevenobelia/Desktop/archive/archive-hugo")

	if ignoreBlobs != nil {

	}
}

func TestWriteLinks(t *testing.T) {

	in := "/Users/stevenobelia/Desktop/archive/archive-hugo/content"
	out := "./assets/indices"
	root := "/Users/stevenobelia/Desktop/archive/archive-hugo"
	//index := true
	baseUrl := getBaseUrl(root)

	ignoreBlobs := getIgnoredFiles(root)
	l, i := walk(in, ".md", true, ignoreBlobs)
	f := filter(l)
	err := write(f, i, true, out, root, baseUrl)
	if err != nil {
		panic(err)
	}
}
