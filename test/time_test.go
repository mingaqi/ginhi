package test

import (
	"fmt"
	"ginhi/util/snowFlake"
	"testing"
)

func TestSnow(t *testing.T) {
	fmt.Println(snowFlake.GetId())
}
