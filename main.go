package main

import (
	"fmt"
	"flag"
	"time"
	task "goroutine/task"
	async "goroutine/async"
)

func main() {
	submain_2()

}

func submain_1() {
	d := task.New()

	var img = flag.String("img", "", "Entrer le nom de l'image")
	var action = flag.String("action", "", "Entrer le nom du filtre ou de l'action")
	var option = flag.String("option", "", "Entrer le nom de l'option")
	flag.Parse()

	task.Choice(*action, *img, *option, d)
}

func submain_2() {
	w := async.NewWaitGroup()

	start := time.Now()
	w.Process()
	elapsed := time.Since(start)

	//grayscale
	//166.0381ms (with waitgroup)
	//215.127ms (without waitgroup)
	//blur
	//214.421ms (with waitgroup)
	//280.0573ms (without waitgroup)

	fmt.Printf("duration %v", elapsed)
}

func submain_3() {
	
}