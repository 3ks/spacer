package scanner

import (
	"io/ioutil"
	"path"
)

func ScanAllFile(pathname string, exts []string) []string {
	files := make([]string, 5)
	count := 0
	e := make(map[string]bool, 0)
	for _, v := range exts {
		e[v] = true
	}
	scanFiles(pathname, e, &files, &count)
	return files[:count]
}

func scanFiles(pathname string, ext map[string]bool, files *[]string, count *int) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		panic(err.Error())
	}
	for _, fi := range rd {
		if fi.IsDir() {
			scanFiles(path.Join(pathname, fi.Name()), ext, files, count)
		} else {
			if ext[path.Ext(fi.Name())] {
				(*files)[*count] = path.Join(pathname, fi.Name())
				*count++
				if *count+2 > len(*files) {
					*files = append(*files, make([]string, 100)...)
				}
			}
		}
	}
}
