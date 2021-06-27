package async

import (
	"os"
	"fmt"
	"log"
	"sync"
	"path"
	"image"
	"strings"
	"image/color"
	"path/filepath"
	"github.com/disintegration/imaging"
)

type WaitGroup struct {
	dir string
	dst string
	files []string
}

func NewWaitGroup() *WaitGroup {
	w := &WaitGroup{
		dir: "testdata/",
		dst: "testout/",
		files: buildFileList("./testdata"),
	}
	return w
}

func (w *WaitGroup) Process() error {
	var wg sync.WaitGroup
	size := len(w.files)
	for i, f := range w.files {
		filename := filepath.Base(f)
		dst := path.Join(w.dst, filename)
		wg.Add(1)

		//go w.applyFilter(f, dst, &wg, i+1, size)
		w.applyFilter(f, dst, &wg, i+1, size)
	}
	wg.Wait()
	fmt.Println("Done processing files !")
	return nil
}

func (w *WaitGroup) applyFilter(src string, dst string, wg *sync.WaitGroup, i int, size int) {
	//grayscale(src, dst)
	blur(src, dst)
	fmt.Printf("Processed [%d/%d] %v => %v\n", i, size, src, dst)
	wg.Done()
}

func buildFileList(rootpath string) []string {
	files := []string{}
	fmt.Println("generating files list..")
	filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || !strings.HasSuffix(path, ".jpg") {
			return nil
		}
		files = append(files, path)
		return nil
	})

	return files
}

func grayscale(srcs string, dsts string) {
	
	src, err := imaging.Open(srcs)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	img := imaging.Grayscale(src)

	dst := imaging.New(400, 400, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, img, image.Pt(0, 0))

	err = imaging.Save(dst, dsts)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

}

func blur(srcs string, dsts string) {
	
	src, err := imaging.Open(srcs)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	img := imaging.Blur(src, 0.1)

	dst := imaging.New(400, 400, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, img, image.Pt(0, 0))

	err = imaging.Save(dst, dsts)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

}