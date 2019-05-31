package main

import (
	"fmt"
	"runtime"
)

const LongSize = 1000000

type astruct struct {
	iteration int
	value int
}

type bstruct struct {
	value int
	ptr *bstruct
}

type cstruct struct {
	value int
	iface Blahable
}

func (x *cstruct) Blah() int {
	return x.value
}

type Blahable interface {
	Blah() int
}

func main() {
	longslice := []astruct{}
	longmap := map[string]astruct{}
	var buf []byte
	
	for i := 0; i < LongSize; i++ {
		longslice = append(longslice, astruct{
			iteration: i,
			value: (i*i)-1,
		})
		
		longmap[fmt.Sprintf("iteration %d", i)] = astruct{
			iteration: i,
			value: (i*i)-1,
		}
	}
	
	for i := 0; i < LongSize/100; i++ {
		buf = append(buf, []byte(fmt.Sprintf("<iteration %d>", i))...)
		buf = append(buf, []byte("0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789")...)
	}
	longstr := string(buf)
	
	infptr := &bstruct{
		value: 1,
		ptr: &bstruct{
			value: 2,
			ptr: &bstruct{
				value: 3,
				ptr: nil,
			},
		},
	}
	infptr.ptr.ptr.ptr = infptr
	
	infiface := &cstruct{
		value: 1,
		iface: &cstruct{
			value: 2,
			iface: &cstruct{
				value: 3,
			},
		},
	}
	infiface.iface.(*cstruct).iface.(*cstruct).iface = infiface
	
	intmap := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	
	x := 15
	
	runtime.Breakpoint()
	fmt.Println(longslice, longmap, longstr, infptr,infiface, intmap, x)
}
