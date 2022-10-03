package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/ssh/terminal"
)

var stack_data [][]byte

func Add_stack_data(v []byte, i int) {
	if i < 0 {
		stack_data = append(stack_data, v)
	} else {
		stack_data = append(stack_data[:i+1], stack_data[i:]...)
		stack_data[i] = v
	}
}

func Get_stack_data(i int) []byte {
	if i < 0 {
		return stack_data[len(stack_data)+i]
	} else {
		return stack_data[i]
	}
}

func Delete_stack_data(i int) {
	if i < 0 {
		stack_data = stack_data[:len(stack_data)-1]
	} else {
		stack_data = append(stack_data[:i], stack_data[i+1:]...)
	}
}

func Load_stack_data(path string) {
	err := errors.New("")
	var Bytes []byte
	for err != nil {
		Bytes, err = ioutil.ReadFile(path)
	}
	_ = json.Unmarshal(Bytes, &stack_data)
}

func Save_stack_data(path string) {
	err := errors.New("")
	var file *os.File
	for err != nil {
		file, err = os.Create(path)
	}
	defer file.Close()
	_ = json.NewEncoder(file).Encode(stack_data)
}

func Check_stack_data(i int) bool {
	if i < 0 {
		return len(stack_data) >= 1
	} else {
		return i < len(stack_data)
	}
}

func File_Exists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

func main() {
	exe_path, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}
	exe_path = filepath.Dir(exe_path)
	if File_Exists(filepath.Join(exe_path, "IO-Stack-data.json")) {
		Load_stack_data(filepath.Join(exe_path, "IO-Stack-data.json"))
	}

	app := &cli.App{
		Name:  "IO-Stack",
		Usage: "make an explosive entrance",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "nd", Value: false},
		},
		Action: func(c *cli.Context) error {
			n := -1
			if c.Args().Len() >= 1 {
				temp := c.Args().Get(0)
				var err error = nil
				n, err = strconv.Atoi(temp)
				if err != nil {
					return err
				}
			}
			if terminal.IsTerminal(int(os.Stdin.Fd())) {
				if !Check_stack_data(n) {
					fmt.Print("")
					return nil
				}
				b := Get_stack_data(n)
				binary.Write(os.Stdout, binary.LittleEndian, b)
				if !c.Bool("nd") {
					Delete_stack_data(n)
				}
				//出力
			} else {
				b, err := ioutil.ReadAll(os.Stdin)
				if err != nil {
					return err
				}
				Add_stack_data(b, n)
				//入力
			}
			Save_stack_data(filepath.Join(exe_path, "IO-Stack-data.json"))
			return nil
		},
	}

	err_CLI := app.Run(os.Args)
	if err_CLI != nil {
		log.Fatal(err_CLI)
	}
}
