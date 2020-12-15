package test

import (
	"fmt"
	MD5 "ginhi/util/md5"
	"ginhi/util/snowFlake"
	"testing"
)

func TestSnow(t *testing.T) {
	fmt.Println(snowFlake.GetId())
	fmt.Println(MD5.MD5Encode("10057xsrmt"))
}
