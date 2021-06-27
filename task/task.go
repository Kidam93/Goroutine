package task

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/disintegration/imaging"
)

type Data struct {
	Dir string
	Dst string
	Files []string
}

func New() *Data {
	d := &Data {
		Dir: "testdata/",
		Dst: "testout/",
		Files: BuildFileList("./testdata"),
	}
	return d
}

func Choice(action string, name string, option string, d *Data) {
	start := time.Now()

	switch action {
	case "blur":
		if name == "all" {
			for _, f := range d.Files {
				filename := filepath.Base(f)
				fmt.Printf("file: %v ", filename)
				blur(filename, d)
				fmt.Println("files are blur")
			}
		} else if name == "all" && option == "speed" {
			//goroutine
			for _, f := range d.Files {
				filename := filepath.Base(f)
				fmt.Printf("file: %v ", filename)
				go blur(filename, d)
				fmt.Println("files are blur")
			}
		} else {
			blur(name, d)
			fmt.Println("file is blur")
		}
	case "grayscale":
		if name == "all" {
			for _, f := range d.Files {
				filename := filepath.Base(f)
				fmt.Printf("file: %v ", filename)
				grayscale(filename, d)
				fmt.Println("files are grey")
			}
		} else {
			grayscale(name, d)
			fmt.Println("file is grey")
		}
	case "invert":
		if name == "all" {
			for _, f := range d.Files {
				filename := filepath.Base(f)
				fmt.Printf("file: %v ", filename)
				invert(filename, d)
				fmt.Println("files are invert")
			}
		} else {
			invert(name, d)
			fmt.Println("file is invert")
		}
	case "delete":
		if name == "all" {
			for _, f := range d.Files {
				filename := filepath.Base(f)
				fmt.Printf("file: %v ", filename)
				delete(filename, d)
				fmt.Println("files are deleted successfuly")
			}
		} else {
			delete(name, d)
			fmt.Println("file is deleted successfuly")
		}		
	default:
		fmt.Println("unknown choice")	
	}

	elapsed := time.Since(start)
	fmt.Printf("duration %v", elapsed)
}

func BuildFileList(rootpath string) []string {
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

func blur(name string, d *Data) {
	
	src, err := imaging.Open(d.Dir + name)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	img := imaging.Blur(src, 0.1)

	dst := imaging.New(400, 400, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, img, image.Pt(0, 0))

	err = imaging.Save(dst, d.Dst + name)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

}

func grayscale(name string, d *Data) {
	
	src, err := imaging.Open(d.Dir + name)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	img := imaging.Grayscale(src)

	dst := imaging.New(400, 400, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, img, image.Pt(0, 0))

	err = imaging.Save(dst, d.Dst + name)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

}

func invert(name string, d *Data) {

	src, err := imaging.Open(d.Dir + name)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	img := imaging.Invert(src)

	dst := imaging.New(400, 400, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, img, image.Pt(0, 0))

	err = imaging.Save(dst, d.Dst + name)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}

func delete(name string, d *Data) {
	e := os.Remove(d.Dst + name)
    if e != nil {
        log.Fatal(e)
    }
}
