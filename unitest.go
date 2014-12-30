package unitest

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

var (
	passRegexp, _     = regexp.Compile(`^\s*unitest\.Pass\s*\(\s*[^,]+\s*,\s*(.+)\s*\)\s*$`)
	notErrorRegexp, _ = regexp.Compile(`^\s*unitest\.NotError\s*\(\s*[^,]+\s*,\s*(.+)\s*\)\s*$`)
)

func Pass(t *testing.T, condition bool) bool {
	if !condition {
		log("[NOT PSSS]", passRegexp, "")
		t.FailNow()
	}
	return condition
}

func NotError(t *testing.T, err error) bool {
	if err != nil {
		log("[ERROR]", notErrorRegexp, err.Error())
		t.FailNow()
	}
	return err != nil
}

func log(title string, regex *regexp.Regexp, val string) {
	if _, file, line, ok := runtime.Caller(2); ok {
		if data, err := ioutil.ReadFile(file); err == nil {
			// Truncate file name at last file name separator.
			if index := strings.LastIndex(file, "/"); index >= 0 {
				file = file[index+1:]
			} else if index = strings.LastIndex(file, "\\"); index >= 0 {
				file = file[index+1:]
			}
			lines := bytes.Split(data, []byte{'\n'})
			cond := regex.FindAllSubmatch(lines[line-1], 1)
			if len(cond) > 0 && len(cond[0]) > 1 {
				if val == "" {
					fmt.Fprintf(os.Stderr, "\t%s %s:%d: %s\n", title, file, line, cond[0][1])
				} else {
					fmt.Fprintf(os.Stderr, "\t%s %s:%d: %s: %s\n", title, file, line, cond[0][1], val)
				}
			}
		}
	}
}

// Command handler
var CommandHandler func(string) bool

func init() {
	go func() {
		for {
			if input, err := ioutil.ReadFile("unitest.cmd"); err == nil && len(input) > 0 {
				ioutil.WriteFile("unitest.cmd", []byte(""), 0744)

				cmd := strings.Trim(string(input), " \n\r\t")

				var (
					profile  *pprof.Profile
					filename string
				)

				switch cmd {
				case "lookup goroutine":
					profile = pprof.Lookup("goroutine")
					filename = "unitest.goroutine"
				case "lookup heap":
					profile = pprof.Lookup("heap")
					filename = "unitest.heap"
				case "lookup threadcreate":
					profile = pprof.Lookup("threadcreate")
					filename = "unitest.thread"
				default:
					if CommandHandler == nil || !CommandHandler(cmd) {
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
