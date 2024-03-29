package wallet

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/urfave/cli/v2"
)

var RemoveCmd = &cli.Command{
	Name:      "remove",
	Usage:     "Remove a wallet",
	ArgsUsage: "<address>",
	Flags: []cli.Flag{
		cliutil.ReallyDotItFlag,
	},
	Action: func(c *cli.Context) error {
		if err := cliutil.HandleReallyDoIt(c); err != nil {
			return err
		}
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		defer closer.Close()
		return wallet.RemoveHandler(db, c.Args().Get(0))
	},
}
