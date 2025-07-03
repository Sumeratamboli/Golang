// package main

//   import ("fmt")

//  func simpleFunction() {

//       fmt.Println("simple function")

//   }

//   func multiply(a, b int) (int) {
//     return a * b

//   }
//   func main() {
//     fmt.Println("We are learning function in Golang")
//     simpleFunction()

//     ans := multiply(3, 4)
//     fmt.Println("multiply of two number is :",ans)
//   }

// package main

// import "fmt"

// func main() {
    
//     fmt.Println("we are trying to speak english")
// }

// Peano integers are represented by a linked
// list whose nodes contain no data
// (the nodes are the data).
// http://en.wikipedia.org/wiki/Peano_axioms

// This program demonstrates that Go's automatic
// stack management can handle heavily recursive
// computations.

// package main

// import "fmt"

// // Number is a pointer to a Number
// type Number *Number

// // The arithmetic value of a Number is the
// // count of the nodes comprising the list.
// // (See the count function below.)

// // -------------------------------------
// // Peano primitives

// func zero() *Number {
// 	return nil
// }

// func isZero(x *Number) bool {
// 	return x == nil
// }

// func add1(x *Number) *Number {
// 	e := new(Number)
// 	*e = x
// 	return e
// }

// func sub1(x *Number) *Number {
// 	return *x
// }

// func add(x, y *Number) *Number {
// 	if isZero(y) {
// 		return x
// 	}
// 	return add(add1(x), sub1(y))
// }

// func mul(x, y *Number) *Number {
// 	if isZero(x) || isZero(y) {
// 		return zero()
// 	}
// 	return add(mul(x, sub1(y)), x)
// }

// func fact(n *Number) *Number {
// 	if isZero(n) {
// 		return add1(zero())
// 	}
// 	return mul(fact(sub1(n)), n)
// }

// // -------------------------------------
// // Helpers to generate/count Peano integers

// func gen(n int) *Number {
// 	if n > 0 {
// 		return add1(gen(n - 1))
// 	}
// 	return zero()
// }

// func count(x *Number) int {
// 	if isZero(x) {
// 		return 0
// 	}
// 	return count(sub1(x)) + 1
// }

// // -------------------------------------
// // Print i! for i in [0,9]

// func main() {
// 	for i := 0; i <= 9; i++ {
// 		f := count(fact(gen(i)))
// 		fmt.Println(i, "! =", f)
// 	}
// }


// An implementation of Conway's Game of Life.
// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// // Field represents a two-dimensional field of cells.
// type Field struct {
// 	s    [][]bool
// 	w, h int
// }

// // NewField returns an empty field of the specified width and height.
// func NewField(w, h int) *Field {
// 	s := make([][]bool, h)
// 	for i := range s {
// 		s[i] = make([]bool, w)
// 	}
// 	return &Field{s: s, w: w, h: h}
// }

// // Set sets the state of the specified cell to the given value.
// func (f *Field) Set(x, y int, b bool) {
// 	f.s[y][x] = b
// }

// // Alive reports whether the specified cell is alive.
// // If the x or y coordinates are outside the field boundaries they are wrapped
// // toroidally. For instance, an x value of -1 is treated as width-1.
// func (f *Field) Alive(x, y int) bool {
// 	x += f.w
// 	x %= f.w
// 	y += f.h
// 	y %= f.h
// 	return f.s[y][x]
// }

// // Next returns the state of the specified cell at the next time step.
// func (f *Field) Next(x, y int) bool {
// 	// Count the adjacent cells that are alive.
// 	alive := 0
// 	for i := -1; i <= 1; i++ {
// 		for j := -1; j <= 1; j++ {
// 			if (j != 0 || i != 0) && f.Alive(x+i, y+j) {
// 				alive++
// 			}
// 		}
// 	}
// 	// Return next state according to the game rules:
// 	//   exactly 3 neighbors: on,
// 	//   exactly 2 neighbors: maintain current state,
// 	//   otherwise: off.
// 	return alive == 3 || alive == 2 && f.Alive(x, y)
// }

// // Life stores the state of a round of Conway's Game of Life.
// type Life struct {
// 	a, b *Field
// 	w, h int
// }

// // NewLife returns a new Life game state with a random initial state.
// func NewLife(w, h int) *Life {
// 	a := NewField(w, h)
// 	for i := 0; i < (w * h / 4); i++ {
// 		a.Set(rand.Intn(w), rand.Intn(h), true)
// 	}
// 	return &Life{
// 		a: a, b: NewField(w, h),
// 		w: w, h: h,
// 	}
// }

// // Step advances the game by one instant, recomputing and updating all cells.
// func (l *Life) Step() {
// 	// Update the state of the next field (b) from the current field (a).
// 	for y := 0; y < l.h; y++ {
// 		for x := 0; x < l.w; x++ {
// 			l.b.Set(x, y, l.a.Next(x, y))
// 		}
// 	}
// 	// Swap fields a and b.
// 	l.a, l.b = l.b, l.a
// }



// An implementation of Conway's Game of Life.
package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

// Field represents a two-dimensional field of cells.
type Field struct {
	s    [][]bool
	w, h int
}

// NewField returns an empty field of the specified width and height.
func NewField(w, h int) *Field {
	s := make([][]bool, h)
	for i := range s {
		s[i] = make([]bool, w)
	}
	return &Field{s: s, w: w, h: h}
}

// Set sets the state of the specified cell to the given value.
func (f *Field) Set(x, y int, b bool) {
	f.s[y][x] = b
}

// Alive reports whether the specified cell is alive.
// If the x or y coordinates are outside the field boundaries they are wrapped
// toroidally. For instance, an x value of -1 is treated as width-1.
func (f *Field) Alive(x, y int) bool {
	x += f.w
	x %= f.w
	y += f.h
	y %= f.h
	return f.s[y][x]
}

// Next returns the state of the specified cell at the next time step.
func (f *Field) Next(x, y int) bool {
	// Count the adjacent cells that are alive.
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && f.Alive(x+i, y+j) {
				alive++
			}
		}
	}
	// Return next state according to the game rules:
	//   exactly 3 neighbors: on,
	//   exactly 2 neighbors: maintain current state,
	//   otherwise: off.
	return alive == 3 || alive == 2 && f.Alive(x, y)
}

// Life stores the state of a round of Conway's Game of Life.
type Life struct {
	a, b *Field
	w, h int
}

// NewLife returns a new Life game state with a random initial state.
func NewLife(w, h int) *Life {
	a := NewField(w, h)
	for i := 0; i < (w * h / 4); i++ {
		a.Set(rand.Intn(w), rand.Intn(h), true)
	}
	return &Life{
		a: a, b: NewField(w, h),
		w: w, h: h,
	}
}

// Step advances the game by one instant, recomputing and updating all cells.
func (l *Life) Step() {
	// Update the state of the next field (b) from the current field (a).
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			l.b.Set(x, y, l.a.Next(x, y))
		}
	}
	// Swap fields a and b.
	l.a, l.b = l.b, l.a
}
