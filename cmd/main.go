package main

import (
	"log"
	"os"

	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/controllers"
	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/db"
	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/ethereum"
	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/utils"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

var commands = []cli.Command{
	startAPI(),
}

func startAPI() cli.Command {
	var env string

	return cli.Command{
		Name: "start-api",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "env",
				Value:       ".env",
				Usage:       "which env file do you want to use",
				Destination: &env,
			},
		},
		Action: func(c *cli.Context) error {
			_ = godotenv.Load(env)
			utils.InitEnvironment()
			db.Init()
			ethereum.Init()

			//INIT REGISTRIES
			// newR := models.NewRegistry().Init()
			// newR.Hash = utils.StrToPtr(`0`)
			// newR.Key = utils.StrToPtr(`161373187550089867448191830760110801114155294027693593477164529548269146668`)
			// newR.Description = utils.StrToPtr(`age >= 12 & age <= 24`)
			// logs.Pretty(newR.Post())

			// newR2 := models.NewRegistry().Init()
			// newR2.Hash = utils.StrToPtr(`0`)
			// newR2.Key = utils.StrToPtr(`418791004851046193537070596848530790547129451305514433175127304050849890764`)
			// newR2.Description = utils.StrToPtr(`age >= 60`)
			// logs.Pretty(newR2.Post())

			// newR3 := models.NewRegistry().Init()
			// newR3.Hash = utils.StrToPtr(`0`)
			// newR3.Key = utils.StrToPtr(`374546399808851745807054416014379391823657543778127138954064098322040293325`)
			// newR3.Description = utils.StrToPtr(`disabled == true`)
			// logs.Pretty(newR3.Post())

			return controllers.NewRouter().Run()
		},
	}
}

func main() {
	api := cli.NewApp()
	api.Commands = commands
	if err := api.Run(os.Args); err != nil {
		log.Fatalf("failed to run the command: %v\n", err)
	}
}
