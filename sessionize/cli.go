package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

//e.g:
//./sessionize -i 3000,3100,3200,6000,6300,9000,29000,29100,29200,29340,29360,29500,29600,30700,30750,30769,31000 -s 1800
//[{3000 200 200} {3100 200 100} {3200 200 0} {6000 300 300} {6300 300 0} {9000 0 0} {29000 1769 1769} {29100 1769 1669} {29200 1769 1569} {29340 1769 1429} {29360 1769 1409} {29500 1769 1269} {29600 1769 1169} {30700 1769 69} {30750 1769 19} {30769 1769 0} {31000 0 0}]

///sessionize -i 3000,3100,3200,6000,6300,9000,29000 -s 1800
//[{3000 200 200} {3100 200 100} {3200 200 0} {6000 300 300} {6300 300 0} {9000 0 0} {29000 0 0}]
func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "sessionize"
	app.Usage = "sessionize -s time_period"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "sessiontime, s",
			Usage: "The time of a session",
			Value: 1800,
		},
		cli.StringFlag{
			Name:  "inputtimes, i",
			Usage: "Comma-separated list of times",
		},
	}

	app.Action = func(ctx *cli.Context) error {
		times, err := validateInput(ctx)
		if err != nil {
			return err
		}

		output := sessionize(times, ctx.Int("sessiontime"))
		fmt.Println(output)

		return nil
	}

	return app
}

func validateInput(ctx *cli.Context) ([]int, error) {
	times := strings.Split(ctx.String("inputtimes"), ",")
	if len(times) <= 1 {
		return nil, fmt.Errorf("The length of the input should be more than 1. Please provide a valid input (e.g: 3000,3100,3200,6000,6300,9000,29000)")
	}

	intTimes := make([]int, len(times))
	for idx, t := range times {
		intTime, err := strconv.Atoi(t)
		if err != nil {
			return nil, fmt.Errorf("%v is not a valid integer (a valid time). Got %v", t, err)
		}
		intTimes[idx] = intTime
	}
	return intTimes, nil
}
