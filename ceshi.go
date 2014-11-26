package ceshi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"runtime"
	"runtime/pprof"
	"strings"
	"testing"
	"time"
)

var failRegexp, _ = regexp.Compile(`^\s*ceshi.Pass\s*\(\s*[^,]+\s*,\s*(.+)\s*\)\s*$`)

func Pass(t *testing.T, condition bool) {
	if !condition {
		if _, file, line, ok := runtime.Caller(1); ok {
			if data, err := ioutil.ReadFile(file); err == nil {
				// Truncate file name at last file name separator.
				if index := strings.LastIndex(file, "/"); index >= 0 {
					file = file[index+1:]
				} else if index = strings.LastIndex(file, "\\"); index >= 0 {
					file = file[index+1:]
				}
				lines := bytes.Split(data, []byte{'\n'})
				cond := failRegexp.FindAllSubmatch(lines[line-1], 1)
				if len(cond) > 0 && len(cond[0]) > 1 {
					fmt.Fprintf(os.Stderr, "\t[NOT PASS] %s:%d: %s\n", file, line, cond[0][1])
				}
			}
		}
		t.FailNow()
	}
}

func StartMonitor(cmdCallback func(string) bool) {
	go func() {
		for {
			if input, err := ioutil.ReadFile("ceshi.cmd"); err == nil && len(input) > 0 {
				ioutil.WriteFile("ceshi.cmd", []byte(""), 0744)

				cmd := strings.Trim(string(input), " \n\r\t")

				var (
					profile  *pprof.Profile
					filename string
				)

				switch cmd {
				case "lookup goroutine":
					profile = pprof.Lookup("goroutine")
					filename = "ceshi.goroutine"
				case "lookup heap":
					profile = pprof.Lookup("heap")
					filename = "ceshi.heap"
				case "lookup threadcreate":
					profile = pprof.Lookup("threadcreate")
					filename = "ceshi.thread"
				default:
					if !cmdCallback(cmd) {
						println("unknow command: '" + cmd + "'")
					}
				}

				if profile != nil {
					file, err := os.Create(filename)
					if err != nil {
						println("couldn't create " + filename)
					} else {
						profile.WriteTo(file, 2)
					}
				}
			}
			time.Sleep(2 * time.Second)
		}
	}()
}
