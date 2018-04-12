package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

var jst *time.Location

func init() {
	var err error
	jst, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	time.Local = jst
}

var timefmts = []string{
	"2006-01-02 15:04:05",
	"2006-01-02",
	time.RFC3339,
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if len(os.Args) > 1 {
		for _, arg := range os.Args {
			convertAndOutput(arg)
		}
	} else {
		for scanner.Scan() {
			input := scanner.Text()
			convertAndOutput(input)
		}
	}
}

func convertAndOutput(input string) {
	if inputVal, err := strconv.ParseInt(input, 10, 64); err == nil {
		output := time.Unix(inputVal, 0)
		p(input, output)
		return
	} else {
		for _, timefmt := range timefmts {
			if t, err := time.ParseInLocation(timefmt, input, jst); err == nil {
				p(input, t)
			}
		}

	}
}

func p(input string, output time.Time) {
	fmt.Fprintf(os.Stdout, "input %s\n", input)
	fmt.Fprintf(os.Stdout, "output jst %s\n", output.String())
	fmt.Fprintf(os.Stdout, "output utc %s\n", output.UTC().String())
	fmt.Fprintf(os.Stdout, "output unix time %d\n", output.Unix())
}
