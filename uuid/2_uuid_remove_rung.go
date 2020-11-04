package main

import (
	"fmt"
	"strings"
)

func main() {
	// uuid 删除横杠

	const (
		HorizontalBar = "-"
		EmptyString   = ""
	)

	str := "89A26769-0582-447D-AB6B-F98082D2CA46"
	str = strings.ReplaceAll(str, HorizontalBar, EmptyString)
	fmt.Println(str)
}
