// @User CPR
package test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSliceAppend(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5, 6}

	arr2 := arr1

	arr2 = append(arr2[:2], arr2[3:]...)

	fmt.Println(arr1)
	fmt.Println(arr2)
}

func TestSlice2(t *testing.T) {
	e := []int32{1, 2, 3}
	fmt.Println("cap of e before:", cap(e), len(e))
	//e = append(e, 4)
	//fmt.Println("cap of e after:", cap(e), len(e))
	e = append(e, 4, 5, 6, 7, 4, 5, 6)
	fmt.Println("cap of e after:", e, cap(e), len(e))
	e = append(e, 4, 5, 6, 7)
	fmt.Println("cap of e after:", e, cap(e), len(e))
}

func TestProcs(t *testing.T) {
	numCPU := runtime.NumCPU()
	fmt.Println(numCPU)
	runtime.GOMAXPROCS(1)
	//runtime.SetMaxThreads(4)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
