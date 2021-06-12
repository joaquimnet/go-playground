package main

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pravj/geopattern"
)

func main() {
	imgOptions := map[string]string{"phrase": strings.Join(os.Args[1:], " ")}
	gp := geopattern.Generate(imgOptions)

	data := []byte(gp)

	os.Mkdir("images", os.ModeAppend)
	os.WriteFile("./images/"+strconv.FormatInt(time.Now().UnixNano(), 10)+".svg", data, os.ModeAppend)
}
