package outputhelper

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Delta456/box-cli-maker"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/theckman/yacspin"
)

var Red = color.New(color.FgRed).SprintFunc()
var Yellow = color.New(color.FgYellow).SprintFunc()
var Green = color.New(color.FgGreen).SprintFunc()

func SpinnerMessage(message string) *yacspin.Spinner {
	cfg := yacspin.Config{
		Frequency: 100 * time.Millisecond,
		CharSet: []string{
			"[      ]",
			"[=     ]",
			"[==    ]",
			"[===   ]",
			"[====  ]",
			"[===== ]",
			"[======]",
		},
		Suffix:            fmt.Sprintf(" %s", message),
		StopCharacter:     "[  ok  ]",
		StopFailCharacter: "[ fail ]",
		StopColors:        []string{"fgGreen"},
		StopFailColors:    []string{"fgRed"},
	}

	spinner, _ := yacspin.New(cfg)
	// handle the error

	spinner.Start()

	return spinner
}

func Question(question string, opts []string) string {
	fmt.Print(question + ": ")

	if len(opts) > 0 {
		fmt.Println()
		for i, v := range opts {
			fmt.Println(fmt.Sprintf("  %d: %s", i+1, v))
		}
		fmt.Println("(please select an option above, using the numbered index):")
	}

	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')

	return strings.TrimSpace(answer)
}

func Fail(message string, exit bool) {
	fmt.Printf("%s%s.\n", Red("[ fail ] "), message)
	if exit {
		os.Exit(0)
	}

}

func Success(message string, exit bool) {
	fmt.Printf("%s%s.\n", Green("[  ok  ] "), message)
	if exit {
		os.Exit(0)
	}
}

func Boxed(exit bool, title string, messages ...string) {
	payload := ""
	for _, m := range messages {
		payload += fmt.Sprintln(m)
	}

	payload = strings.TrimSpace(payload)
	Box := box.New(box.Config{Px: 2, Py: 2, ContentAlign: "Left", Type: "Double", Color: "Cyan", TitlePos: "Top"})
	Box.Print(title, payload)

	if exit {
		os.Exit(0)
	}
}

func Table(cols []string, rows [][]string) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(cols)

	for _, v := range rows {
		table.Append(v)
	}
	table.Render() // Send output

}
