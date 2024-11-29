package main

import (
	"context"
	"embed"
	"vesgoapp/src/db"
	"vesgoapp/src/repositories"
	"vesgoapp/src/services"

	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

//go:embed all:frontend/dist
var assets embed.FS

type App struct {
	ctx           context.Context
	boletaService *services.BoletaService
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	if a.boletaService != nil {
		a.boletaService.SetContext(ctx)
	}
}

func main() {

	godotenv.Load()

	config := db.Config{
		Host:     "192.168.18.125",
		Port:     3306,
		User:     "root",
		Password: "pleyades",
		Database: "rrhh",
	}

	dbt := db.GetConnection(config)
	defer db.CloseConnection()

	urep := repositories.CreateLoginRepo(dbt)
	lservice := services.NewLoginService(urep)
	dashservice := services.NewDashboardService(repositories.CreatedashboardRepository(dbt))
	persoService := services.NewPersonalService(repositories.CreatePersonalRepository(dbt))

	boletaRepo := repositories.CreateboletaRepository()
	boletaService := services.NewBoletaService(boletaRepo)

	app := &App{
		boletaService: boletaService,
	}

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
		Bind: []interface{}{
			app,
			lservice,
			dashservice,
			persoService,
			boletaService,
		},
		OnStartup: app.startup,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
