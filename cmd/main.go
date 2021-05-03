package main

import (
	"log"
	"os"

	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/controllers"
	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/db"
	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/models"
	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/utils"
	"github.com/joho/godotenv"
	"github.com/microgolang/logs"
	"github.com/urfave/cli"
)

var commands = []cli.Command{
	startAPI(),
}

func InitRegistry() {
	newR := models.NewRegistry().Init()
	newR.Hash = utils.StrToPtr(`0`)
	newR.Key = utils.StrToPtr(`161373187550089867448191830760110801114155294027693593477164529548269146668`)
	newR.Description = utils.StrToPtr(`age >= 12 & age <= 24`)
	logs.Pretty(newR.Post())

	newR2 := models.NewRegistry().Init()
	newR2.Hash = utils.StrToPtr(`0`)
	newR2.Key = utils.StrToPtr(`418791004851046193537070596848530790547129451305514433175127304050849890764`)
	newR2.Description = utils.StrToPtr(`age >= 60`)
	logs.Pretty(newR2.Post())

	newR3 := models.NewRegistry().Init()
	newR3.Hash = utils.StrToPtr(`0`)
	newR3.Key = utils.StrToPtr(`374546399808851745807054416014379391823657543778127138954064098322040293325`)
	newR3.Description = utils.StrToPtr(`disabled == true`)
	logs.Pretty(newR3.Post())
}

func InitRegistryMapping() {
	newRM := models.NewRegistryMapping().Init()
	newRM.RegistryKey = utils.StrToPtr(`161373187550089867448191830760110801114155294027693593477164529548269146668`)
	newRM.Identity = utils.StrToPtr(`0x3133b8a0cf50d294c308df4de79d1427af2f639d852b702de6ba0569a630aac`)
	newRM.IdentityIndex = utils.Uint64ToPtr(0)
	newRM.Post()

	newRM2 := models.NewRegistryMapping().Init()
	newRM2.RegistryKey = utils.StrToPtr(`161373187550089867448191830760110801114155294027693593477164529548269146668`)
	newRM2.Identity = utils.StrToPtr(`0x6dbf4904695c73b69167a1cdd363c5ea30db5cae3184ee30fc88ccf403e1455`)
	newRM2.IdentityIndex = utils.Uint64ToPtr(1)
	newRM2.Post()

	newRM3 := models.NewRegistryMapping().Init()
	newRM3.RegistryKey = utils.StrToPtr(`161373187550089867448191830760110801114155294027693593477164529548269146668`)
	newRM3.Identity = utils.StrToPtr(`0x751597366e9a9b9f19d6eaa7b6dad23c0a93978b1760e3f7834e49b7c0f9337`)
	newRM3.IdentityIndex = utils.Uint64ToPtr(2)
	newRM3.Post()

	newRM4 := models.NewRegistryMapping().Init()
	newRM4.RegistryKey = utils.StrToPtr(`161373187550089867448191830760110801114155294027693593477164529548269146668`)
	newRM4.Identity = utils.StrToPtr(`0x5273fe46a790b8ba038b2905ec226163a07812f50640c07bed9436ce044ee51`)
	newRM4.IdentityIndex = utils.Uint64ToPtr(3)
	newRM4.Post()

	newRM5 := models.NewRegistryMapping().Init()
	newRM5.RegistryKey = utils.StrToPtr(`161373187550089867448191830760110801114155294027693593477164529548269146668`)
	newRM5.Identity = utils.StrToPtr(`0xd90ee37c279811968ef5807f7ad084f0bb5932b824da0640df88b43eac3325`)
	newRM5.IdentityIndex = utils.Uint64ToPtr(4)
	newRM5.Post()
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
			utils.InitPusher()

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
