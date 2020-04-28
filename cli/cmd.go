package cli

import (
	"fmt"
	"github.com/breathbath/erplyapi/db"
	"github.com/breathbath/erplyapi/migrations"
	"github.com/breathbath/erplyapi/migrations/items"
	"github.com/breathbath/erplyapi/server"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

//Version of the application
var Version = "undefined"

//Execute entry point for cli commands
func Execute() error {
	versionStr := fmt.Sprintf("Version: %s", Version)
	rootCliApp := kingpin.New("VisitsCounter", "Visits Counter").Version(versionStr)
	versionCommand := rootCliApp.Command("version", "Gives version output")

	startServerCmd := rootCliApp.Command("server", "Starts REST server")
	migrateCommand := rootCliApp.Command("migrate", "Runs DB migrations")
	parsedCliInput := kingpin.MustParse(rootCliApp.Parse(os.Args[1:]))

	switch parsedCliInput {
	case versionCommand.FullCommand():
		fmt.Println(versionStr)
		return nil
	case startServerCmd.FullCommand():
		err := server.Start()
		if err != nil {
			return err
		}
	case migrateCommand.FullCommand():
		dg, err := db.NewDb()
		if err != nil {
			return err
		}
		registry := migrations.NewRegistry(dg)
		items.AddItems(registry)

		return registry.Execute()
	}

	return nil
}
