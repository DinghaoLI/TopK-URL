package utils

import (
	"os"
	"strconv"
)

func GenerateUrlData(str string) error {
	f, err := os.OpenFile(str, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		return err		
	} else {
		count := 0
		for count < 1000 {
			for iter:=count; iter>=0; iter--{
				_, err = f.Write([]byte("https://dinghao.li.github.io/" + strconv.Itoa(count) + "\n"))
				if err != nil {
					return err
				}
			}
			count++
		}
	}
	return nil
}