package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	
	task "github.com/danielaltamirano1993/go-crud-cli/tasks"
)


func main() {
	// read or create the tasks.json file
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var tasks []task.Task

	// get file info
	info, err := file.Stat()

	if err != nil {
		panic(err)
	}
