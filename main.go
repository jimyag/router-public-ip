package main

import (
	"fmt"
	"router-public-ip/report"
)

func main() {
	re := &report.MyReport{}
	fmt.Println(re.Send())
}
