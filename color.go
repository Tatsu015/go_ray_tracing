package main

import "fmt"

func WriteColor(pixcelColor Color) string {
	buf := fmt.Sprintf("%d %d %d\n", int(255.999*pixcelColor.x), int(255.999*pixcelColor.y), int(255.999*pixcelColor.z))
	return buf
}
