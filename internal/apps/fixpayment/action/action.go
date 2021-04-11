package action

import (
	"os"

	"github.com/masalennon/golang-batch-structures-sample/internal/apps/fixpayment/service"
	"github.com/urfave/cli/v2"
)

func Action() {
	app := cli.App{
		Name:    "Fixpayment",
		Usage:   "決済確定バッチ",
		Version: "0.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "mode",
				Aliases: []string{"m"},
				Usage:   "'F'決済確定処理（初め）から処理を開始, 'B':〇〇処理から処理を開始",
			},
		},
		Action: doAction,
	}
	app.Run(os.Args)

	return
}

func doAction(ctx *cli.Context) error {
	fixpaymentService := service.NewFixpayment()
	switch ctx.String("mode") {
	case "F", "f":
		fixpaymentService.Fix()
		fallthrough
	case "B", "b":
		//do something
	default:
		fixpaymentService.Fix()
		//do something
	}
	return nil
}
