package terminal

import (
	"fmt"
	"github.com/liushuochen/gotable/table"
	"github.com/urfave/cli/v2"
	"time"
	"vdns/lib/util/strs"
)

func printTableAndSavaToJSONFile(t *table.Table, ctx *cli.Context) error {
	go spinner()
	fmt.Printf("\r%v", t)
	path := ctx.String("path")
	return toJsonFile(t, path)
}

func printTableAndSavaToCSVFile(t *table.Table, ctx *cli.Context) error {
	go spinner()
	fmt.Printf("\r%v", t)
	path := ctx.String("path")
	return toCSVFile(t, path)
}

func toCSVFile(table *table.Table, path string) error {
	if !strs.IsEmpty(path) {
		err := table.ToCSVFile(path)
		if err != nil {
			return err
		}
		fmt.Printf("sava to: %s\n", path)
	}
	return nil
}

func toJsonFile(table *table.Table, path string) error {
	if !strs.IsEmpty(path) {
		err := table.ToJsonFile(path, 2)
		if err != nil {
			return err
		}
		fmt.Printf("\nsava to: %s\n", path)
	}
	return nil
}

// animation waiting
func spinner() {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
