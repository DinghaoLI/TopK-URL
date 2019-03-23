package main

import (
	. "./utils"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
	//"strings"
)

var NUM_FILE int = 100
var NUM_TOP int = 100
var SIZE_BATCH int = 3900000

var memString []string

func ReadFile(filePath string, handle func([]string)) error {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	count := 0
	memString = make([]string, 0)
	for {
		line, _, err := buf.ReadLine()
		if count == SIZE_BATCH {
			handle(memString)
			memString = make([]string, 0)
			count = 0
		}
		memString = append(memString, string(line))
		if err != nil {
			if err == io.EOF {
				if len(memString) > 0 {
					handle(memString)
					memString = make([]string, 0)
					count = 0
				}
				return nil
			}
			return err
		}
		count++
	}
}

func setPartition(strs []string) {
	fileMap := make(map[string][]string)
	for _, str := range strs {
		if str == "" {
			continue
		}
		partition := "./tmp/" + strconv.Itoa(int(BKDRHash64(str))%NUM_FILE) + ".txt"
		if _, ok := fileMap[partition]; ok {
			fileMap[partition] = append(fileMap[partition], str)
		} else {
			fileMap[partition] = []string{str}
		}
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
	for k, vs := range fileMap {
		f, err := os.OpenFile(k, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			for _, v := range vs {
				_, err = f.Write([]byte(v + "\n"))
			}
		}
		f.Close()
	}
}

func Print(str string) {
	fmt.Println(str)
}

func MinHeapFile(filePath string) (*MinHeap, error) {
	FreqMap := make(map[string]int64)
	var addToHashmap func([]string)
	addToHashmap = func(keys []string) {
		for _, key := range keys {
			if _, ok := FreqMap[key]; ok {
				FreqMap[key]++
			} else {
				if key != "" {
					FreqMap[key] = 1
				}
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
	for i := 0; i < NUM_FILE; i++ {
		NextHeap, err := MinHeapFile("./tmp/" + strconv.Itoa(i) + ".txt")
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
			_, err = f.Write([]byte("Frequence: " + strconv.FormatInt(item.Freq, 10) + " | Url: " + item.Addr + "\n"))
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
		}
		return nil
	}
}

func main() {

	t1 := time.Now() // get current time

	// err := GenerateUrlData("./Dataset.txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)

	t2 := time.Now() // get current time

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

	elapsed = time.Since(t2)
	fmt.Println("App elapsed: ", elapsed)

}
