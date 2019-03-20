package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	/*"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"*/
	"io/ioutil"
	"strings"
	"os"
	"fmt"
	"strconv"
)

func main(){
	filename := os.Args[1]
	if filename == ""{
		fmt.Errorf("%s","enter filename as arg to binary")
		return
	}
	fmt.Println("filename:",filename)
	file , err := ioutil.ReadFile(filename)
	if err != nil {
		panic("failed to read file..error: "+err.Error())
	}
	f := string(file)
	var xys plotter.XYs
	p, _ := plot.New()
	for _, v := range strings.Split(f,"\n"){
		d := strings.Split(v,",")
		if len(d) != 0 {
			x, _ := strconv.Atoi(d[1])
			y, _ := strconv.Atoi(d[6])
			xys = append(xys, struct{ X, Y float64 }{X: float64(x), Y: float64(y)})
		}
	}

	/*xys = plotter.XYs{
		{X: 1,Y: 2},{X: 2, Y: 3},
	}*/
	_,scatter, _ := plotter.NewLinePoints(xys)
	p.Title.Text = "line-plot"
	p.X.Max = 10000
	p.Add(scatter)
	w, _ := p.WriterTo(10*vg.Inch, 10*vg.Inch, "line-plot.png")
	w.WriteTo(os.Stdout)
	if err := p.Save(10*vg.Inch, 10*vg.Inch, "line-plot.png"); err != nil {
		panic(err)
	}
}
