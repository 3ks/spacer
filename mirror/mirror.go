package mirror

import (
	"fmt"
	"io/ioutil"
	"os"
	"spacer/scanner"
	"strings"
	"sync"
	"time"
)

// 将某个目录下的某些文件按照对应的层次结构，复制至另外一个目录
func Start(pathname, new string) {
	fmt.Println("start mirror...")
	now := time.Now()
	files := scanner.ScanAllFile(pathname, []string{".jpg", ".jpeg", ".png", ".svg"})
	wg := sync.WaitGroup{}
	for _, v := range files {
		wg.Add(1)
		go move(v, new, &wg)
	}
	wg.Wait()
	cost := time.Now().Sub(now).Milliseconds()
	fmt.Printf("done! affect %v files, cost %v ms\n", len(files), cost)
}

func move(pathname, new string, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	newPathname := strings.ReplaceAll(pathname, "en/", "zh/")
	f, err := os.OpenFile(pathname, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("open source file fail. path: %v, err: %v\n", pathname, err.Error())
		return
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("read source file fail. path: %v, err: %v\n", pathname, err.Error())
		return
	}

	_ = os.Remove(newPathname)
	f2, err := os.OpenFile(newPathname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("open target file fail. path: %v, err: %v\n", newPathname, err.Error())
		return
	}

	_, err = f2.Write(data)
	if err != nil {
		fmt.Printf("write target file fail. path: %v, err: %v\n", newPathname, err.Error())
		return
	}
}
