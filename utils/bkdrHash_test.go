package utils

import (
	"testing"
)

func Test_BKDRHash(t *testing.T) {
	if BKDRHash("https://dinghao.li.github.io") == uint32(431729969) {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}
	if BKDRHash("https://google.fr") == uint32(1545209818) {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}
	if BKDRHash("https://baidu.com") == uint32(1059144619) {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}

	if BKDRHash64("https://dinghao.li.github.io") == uint64(4146912629060513073) {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}
	if BKDRHash64("https://google.fr") == uint64(4380720494318061530) {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}
	if BKDRHash64("https://baidu.com") == uint64(3937770163576849323) {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}

}
