package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	task "github.com/faztweb/go-crud-cli/tasks"
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

	// check the size
	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}
	} else {
		tasks = []task.Task{}
	}

	// check the user input and print options to choose
	if len(os.Args) < 2 {
		printUsage()
		return
	}