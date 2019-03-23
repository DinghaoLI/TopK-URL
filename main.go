package main

import (
	. "./utils"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	//"strings"
)

var NUM_FILE int = 100
var NUM_TOP int = 100 

func ReadFile(filePath string, handle func(string)) error {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, _, err := buf.ReadLine()
		handle(string(line))
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}


func setPartition(str string) {
	if str == "" {
		return
	}
	temp_dir := "./tmp"
	_, err := os.Stat(temp_dir)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(temp_dir, os.ModePerm)
			if err != nil {
				fmt.Printf("mkdir failed![%v]\n", err)
				return
			}
		}
	}

	partition := strconv.Itoa(int(BKDRHash64(str)) % NUM_FILE)
	f, err := os.OpenFile("./tmp/"+partition+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = f.Write([]byte(str + "\n"))
	}
}

func Print(str string) {
	fmt.Println(str)
}

func MinHeapFile(filePath string) (*MinHeap, error) {
	FreqMap := make(map[string]int64)
	var addToHashmap func(string)
	addToHashmap = func(key string){
		if _, ok := FreqMap[key]; ok {
			FreqMap[key]++
		} else {
			if key != "" {
				FreqMap[key] = 1
			} 
		}
	}

	err := ReadFile(filePath, addToHashmap)
	if err != nil {
		return nil, err
	}

	heap := NewMinHeap()
	for k, v := range FreqMap {
		if heap.Length() < NUM_TOP {
			heap.Insert(&Url{v, k})
			continue
		}
		min, _ := heap.Min()
		if min.Freq <= v {
			heap.DeleteMin()
			heap.Insert(&Url{v, k})
		}
	}

	return heap, nil
}

func mergeTwoHeap(oldH, newH *MinHeap) *MinHeap {
	if newH == nil || newH.Length() == 0 {
		return oldH
	}
	for newH.Length() != 0 {
		value, _ := newH.DeleteMin()
		if oldH.Length() < NUM_TOP {
			oldH.Insert(value)
			continue
		}
		min, _ := oldH.Min()
		if min.Freq <= value.Freq {
			oldH.DeleteMin()
			oldH.Insert(value)
		}
	}
	return oldH
}

func reduce() *MinHeap {
	heap := NewMinHeap()
	for i:=0; i<NUM_FILE; i++{
		NextHeap, err := MinHeapFile("./tmp/"+strconv.Itoa(i)+".txt")
		if err != nil {
			continue
		}
		heap = mergeTwoHeap(heap, NextHeap)
	}
	return heap
}

func heapToFile(heap *MinHeap) error {
	f, err := os.OpenFile("./output.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
		return err
	} else {
		for heap.Length() != 0 {
			item, _ := heap.DeleteMin()
			_, err = f.Write([]byte("Frequence: "+ strconv.FormatInt(item.Freq, 10) + " | Url: " + item.Addr + "\n"))
			if err != nil {
				fmt.Println(err.Error())
				return err
			}		
		}
	return nil
	}
}

func main() {
	// err := GenerateUrlData("./Dataset.txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	err := ReadFile("./Dataset.txt", setPartition)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	heap := reduce()
	
	err = heapToFile(heap)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
