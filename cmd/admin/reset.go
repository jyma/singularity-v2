package admin

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/admin"
	"github.com/urfave/cli/v2"
)

var ResetCmd = &cli.Command{
	Name:  "reset",
	Usage: "Reset the database",
	Flags: []cli.Flag{cliutil.ReallyDotItFlag},
	Action: func(context *cli.Context) error {
		if err := cliutil.HandleReallyDoIt(context); err != nil {
			return err
		}
		db, closer, err := database.OpenFromCLI(context)
		if err != nil {
			return err
		}
		defer closer.Close()
		return admin.ResetHandler(db)
	},
}
