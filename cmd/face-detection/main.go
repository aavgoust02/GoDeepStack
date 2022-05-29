package main

import (
	"fmt"
	d "github.com/aavgoust02/GoDeepStack"
)

func main() {
	c := d.NewDeepStackClient("127.0.0.1", 5000, false, "")
	od := c.NewFaceDetector()
	objs := od.Detect("~/image.jpg")
	fmt.Println(objs)
}
