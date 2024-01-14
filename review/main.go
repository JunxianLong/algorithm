package main

import (
	"fmt"
)

func main() {
	//bm := NewBitMap(8)
	//fmt.Printf("%08b\n", bm.bits)
	//bm.Set(8)
	//fmt.Printf("%08b\n", bm.bits)
	//has := bm.Check(8)
	//fmt.Println(has)
	//bm.UnSet(8)
	//fmt.Printf("%08b\n", bm.bits)

	//nums := []int{2, 2, 1, 1, 1, 2, 2}
	//num := majorityElement(nums)
	//fmt.Println(num)

	//nums := []int{3, 38, 5, 44, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	//nums := []int{9, 9, 9}
	//BubbleSort(nums)
	//SelectSort(nums)
	//QuickSort(nums, 0, len(nums)-1)

	//result := plusOne(nums)

	//nums1 := []int{6, 6, 6, 6}
	//nums2 := []int{2, 2}
	//result := intersect(nums1, nums2)
	//fmt.Println(result)
	//fmt.Println((8 + 5) % 8)

	//n := distributeCandies(nums1)
	//fmt.Println(n)

	//nums1 = []int{2, 1, 3}
	//result := findLengthOfLCIS(nums1)
	//fmt.Println(result)
	//students := []Student{Student{"A"}, Student{"B"}, Student{"C"}}
	//var students2 []*Student
	//for _, s := range students {
	//	students2 = append(students2, &s)
	//}
	//
	//for _, s := range students2 {
	//	fmt.Println(s)
	//}
	//nums := []int{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 0, 6, 3}
	//k := 516
	//result := addToArrayForm(nums, k)
	//fmt.Println(result)
	//words := []string{"bella", "label", "roller"}
	//fmt.Println(commonChars(words))

	nums1 := []int{2, 2}
	nums2 := []int{2, 2}
	fmt.Println(relativeSortArray(nums1, nums2))

}

type Student struct {
	Name string
}
type BitMap struct {
	bits []byte
	vmax uint
}

func NewBitMap(max_val ...uint) *BitMap {
	var max uint = 8192
	if len(max_val) > 0 && max_val[0] > 0 {
		max = max_val[0]
	}

	bm := &BitMap{}
	bm.vmax = max
	sz := (max + 8) / 8
	bm.bits = make([]byte, sz, sz)
	return bm
}

func (bm *BitMap) Set(num uint) {
	if num > bm.vmax {
		bm.vmax += 1024
		if bm.vmax < num {
			bm.vmax = num
		}

		dd := int(num+8)/8 - len(bm.bits)
		if dd > 0 {
			tmp_arr := make([]byte, dd, dd)
			fmt.Printf("tmp:%08b\n", tmp_arr)
			bm.bits = append(bm.bits, tmp_arr...)
		}
	}

	//将1左移num%8后，然后和以前的数据做|，这样就替换成1了
	bm.bits[num/8] |= 1 << (num % 8)
}

func (bm *BitMap) UnSet(num uint) {
	if num > bm.vmax {
		return
	}
	//&^:将1左移num%8后，然后进行与非运算，将运算符左边数据相异的位保留，相同位清零
	bm.bits[num/8] &^= 1 << (num % 8)
}

func (bm *BitMap) Check(num uint) bool {
	if num > bm.vmax {
		return false
	}
	//&:与运算符，两个都是1，结果为1
	return bm.bits[num/8]&(1<<(num%8)) != 0
}
