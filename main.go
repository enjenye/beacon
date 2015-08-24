package main

import (
    "fmt"
    "bitbucket.org/binet/go-gnuplot/pkg/gnuplot"
)

func main() {
    fmt.Println("Hello World")

	// start plot
	plt, err := gnuplot.NewPlotter("", false, false)
	if err != nil { return }
	defer plt.Close()

	// sample plot
	err = plt.PlotX([]float64{10, 20, 30}, "my title")


}
