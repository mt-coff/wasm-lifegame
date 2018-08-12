// +build js,wasm

package main

import (
	"math"
	"math/rand"
	"syscall/js"
	"time"
)

const (
	cellSize = 8
)

type lifeGame struct {
	canvas js.Value
	ctx    js.Value
	cols   int64
	rows   int64
	cells  [][]int64
}

func main() {
	l := lifeGame{}
	l.canvas = js.Global().Get("document").Call("getElementById", "canvas")
	l.ctx = l.canvas.Call("getContext", "2d")
	l.cols = int64(math.Floor(l.canvas.Get("width").Float() / cellSize))
	l.rows = int64(math.Floor(l.canvas.Get("height").Float() / cellSize))
	l.initCells()

	bs := js.Global().Get("document").Call("getElementById", "btnStart")
	bs.Call("addEventListener", "click", js.NewCallback(func(args []js.Value) {
		l.start()
	}))

	select {}
}

func (l *lifeGame) initCells() {
	l.ctx.Set("fillStyle", "rgb(24, 19, 3)")
	l.ctx.Call("fillRect", 0, 0, l.canvas.Get("width"), l.canvas.Get("height"))
	l.cells = make([][]int64, l.cols)

	for i := int64(0); i < l.cols; i++ {
		l.cells[i] = make([]int64, l.rows)
	}

	l.redraw()
}

func (l *lifeGame) start() {
	rand.Seed(time.Now().UnixNano())

	for i := int64(0); i < l.cols; i++ {
		l.cells[i] = make([]int64, l.rows)
		for j := int64(0); j < l.rows; j++ {
			l.cells[i][j] = int64(math.Round(rand.Float64()))
		}
	}

	for {
		time.Sleep(1 * time.Second)
		l.next()
	}
}

func (l *lifeGame) redraw() {
	for i := int64(0); i < l.cols; i++ {
		for j := int64(0); j < l.rows; j++ {
			l.drawCell(i, j)
		}
	}
}

func (l *lifeGame) drawCell(x, y int64) {
	v := l.cells[x][y]
	if v == 0 {
		l.ctx.Set("fillStyle", "rgb(24, 19, 3)")
	} else {
		l.ctx.Set("fillStyle", "rgb(255, 253, 208)")
	}
	l.ctx.Call("fillRect", x*cellSize, y*cellSize, cellSize-1, cellSize-1)
}

func (l *lifeGame) next() {
	newCells := make([][]int64, l.cols)

	for i := int64(0); i < l.cols; i++ {
		newCells[i] = make([]int64, l.rows)
		for j := int64(0); j < l.rows; j++ {
			cnt := l.countLivingAround(i, j)
			if l.cells[i][j] != 0 {
				if cnt == 2 || cnt == 3 {
					newCells[i][j] = 1
				} else {
					newCells[i][j] = 0
				}
			} else {
				if cnt == 3 {
					newCells[i][j] = 1
				} else {
					newCells[i][j] = 0
				}
			}
		}
	}

	l.cells = newCells
	l.redraw()
}

func (l *lifeGame) countLivingAround(x, y int64) int64 {
	cnt := int64(0)

	for i := int64(-1); i <= 1; i++ {
		for j := int64(-1); j <= 1; j++ {
			if (i != 0 || j != 0) && x+i >= 0 && x+i < l.cols && y+j >= 0 && y+j < l.rows {
				cnt += l.cells[x+i][y+j]
			}
		}
	}

	return cnt
}
