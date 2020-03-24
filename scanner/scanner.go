package scanner

import (
	"io/ioutil"
	"path"
)

func ScanAllFile(pathname, ext string) []string {
	files := make([]string, 5)
	count := 0
	scanFiles(pathname, ext, &files, &count)
	return files[:count]
}

func scanFiles(pathname, ext string, files *[]string, count *int) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		panic(err.Error())
	}
	for _, fi := range rd {
		if fi.IsDir() {
			scanFiles(path.Join(pathname, fi.Name()), ext, files, count)
		} else {
			if path.Ext(fi.Name()) == ext {
				(*files)[*count] = path.Join(pathname, fi.Name())
				*count++
				if *count+2 > len(*files) {
					*files = append(*files, make([]string, 100)...)
				}
			}
		}
	}
}
