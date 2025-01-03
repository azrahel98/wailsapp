package main

import (
	"embed"
	"os"
	"vesgoapp/src/db"
	"vesgoapp/src/repositories"
	"vesgoapp/src/services"

	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

var assets embed.FS

func main() {

	godotenv.Load()

	user := os.Getenv("DBUSER")
	dbname := os.Getenv("DBNAME")
	hostdb := os.Getenv("DBHOST")
	passuser := os.Getenv("DBPASS")

	config := db.Config{
		Host:     hostdb,
		Port:     3306,
		User:     user,
		Password: passuser,
		Database: dbname,
	}

	dbt := db.GetConnection(config)
	defer db.CloseConnection()

	urep := repositories.CreateLoginRepo(dbt)
	lservice := services.NewLoginService(urep)
	dashservice := services.NewDashboardService(repositories.CreatedashboardRepository(dbt))
	persoService := services.NewPersonalService(repositories.CreatePersonalRepository(dbt))

	boletaRepo := repositories.CreateboletaRepository()
	boletaService := services.NewBoletaService(boletaRepo, repositories.CreatePersonalRepository(dbt))

	err := wails.Run(&options.App{
		Title:    "vesgoapp",
		Width:    1024,
		Height:   768,
		MinWidth: 500,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},

		Linux: &linux.Options{
			Icon:                []byte{},
			WindowIsTranslucent: true,
			Messages:            &linux.Messages{},
			WebviewGpuPolicy:    linux.WebviewGpuPolicyOnDemand,
			ProgramName:         "",
		},
		Windows: &windows.Options{
			DisableWindowIcon: false,
			Theme:             windows.Theme(windows.Dark),
			CustomTheme: &windows.ThemeSettings{
				DarkModeTitleBar:  windows.RGB(26, 34, 52),
				DarkModeTitleText: windows.RGB(220, 220, 255),
				DarkModeBorder:    windows.RGB(0, 0, 0),
			},
		},

		Bind: []interface{}{
			lservice,
			dashservice,
			persoService,
			boletaService,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
