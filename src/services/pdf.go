package services

import (
	"fmt"
	"log"
	"os"
	"strings"
	"vesgoapp/src/models"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/breakline"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/repository"
)

func MakePdf(name string, persona models.Boleta) {

	m := GetMaroto(persona)

	d, err := m.Generate()
	if err != nil {
		fmt.Println(err)
	}
	key := os.Getenv("PDFRESULT")
	err = d.Save(fmt.Sprintf("%s/%s.pdf", key, name))
	if err != nil {
		fmt.Println(err)
	}

}

func GetMaroto(persona models.Boleta) core.Maroto {

	bytes, err := os.ReadFile("D:/proyectos/wails/wailsapp/src/asset/mves.jpg")
	if err != nil {
		fmt.Println(err)
	}

	customFont := "Ubuntu"

	customFonts, err := repository.New().
		AddUTF8Font(customFont, fontstyle.Normal, "D:/proyectos/wails/wailsapp/src/asset/fonts/Ubuntu-Regular.ttf").
		AddUTF8Font(customFont, fontstyle.Bold, "D:/proyectos/wails/wailsapp/src/asset/fonts/Ubuntu-Bold.ttf").
		Load()
	if err != nil {
		fmt.Println(err)
	}

	cfg := config.NewBuilder().
		WithLeftMargin(8).
		WithTopMargin(8).
		WithRightMargin(8).
		WithPageSize(pagesize.A4).
		WithCustomFonts(customFonts).
		WithBackgroundImage(bytes, extension.Jpg).
		Build()

	mrt := maroto.New(cfg)

	m := maroto.NewMetricsDecorator(mrt)

	err = m.RegisterHeader(getPageHeader())
	if err != nil {
		log.Fatal(err.Error())
	}

	m.AddRow(10,
		text.NewCol(4, fmt.Sprintf("Periodo: %s", persona.Fecha), props.Text{

			Size: 9,

			Align:  align.Left,
			Family: "Ubuntu",
		}),
		text.NewCol(4, "BOLETA DE PAGO", props.Text{
			Size:   15,
			Style:  fontstyle.Bold,
			Align:  align.Center,
			Family: "Ubuntu",
		}),
		text.NewCol(4, "N° 00000000004", props.Text{

			Size: 9,

			Align:  align.Center,
			Family: "Ubuntu",
		}),
	)

	m.AddRow(5)

	m.AddRow(6, text.NewCol(3, "Doc. de Identidad :", props.Text{
		Align:  align.Left,
		Size:   9,
		Family: "Ubuntu",
	}),
		text.NewCol(3, persona.NumeroDoc, props.Text{
			Align:  align.Left,
			Size:   9,
			Style:  fontstyle.Bold,
			Family: "Ubuntu",
		}),
		text.NewCol(3, "Codigo AIRHSP:", props.Text{
			Align:  align.Left,
			Size:   9,
			Family: "Ubuntu",
		}),
		text.NewCol(3, persona.NumeroDoc, props.Text{
			Align:  align.Left,
			Size:   9,
			Style:  fontstyle.Bold,
			Family: "Ubuntu",
		}),
	)

	m.AddRow(6, text.NewCol(2, "Apellidos y Nombres:", props.Text{
		Align:  align.Left,
		Size:   9,
		Family: "Ubuntu",
	}),
		text.NewCol(4, persona.Nombre, props.Text{
			Align:             align.Center,
			Size:              9,
			Style:             fontstyle.Bold,
			BreakLineStrategy: breakline.EmptySpaceStrategy,
			Family:            "Ubuntu",
		}),
		text.NewCol(3, "Establecimiento:", props.Text{
			Align:  align.Left,
			Size:   9,
			Family: "Ubuntu",
		}),
		text.NewCol(3, "Villa el Salvador", props.Text{
			Align:  align.Left,
			Size:   9,
			Family: "Ubuntu",
		}),
	)
	m.AddRow(6, text.NewCol(3, "Fecha de Ingreso:", props.Text{
		Align:  align.Left,
		Size:   9,
		Family: "Ubuntu",
	}),
		text.NewCol(3, persona.Ingreso, props.Text{
			Align:  align.Left,
			Family: "Ubuntu",
			Size:   9,
		}),
		text.NewCol(3, "Regimen Laboral:", props.Text{
			Align:  align.Left,
			Family: "Ubuntu",
			Size:   9,
		}),
		text.NewCol(3, strings.TrimSpace(persona.TipoTrabajador), props.Text{
			Align:  align.Left,
			Style:  fontstyle.Bold,
			Size:   8,
			Family: "Ubuntu",
		}),
	)

	m.AddRow(6, text.NewCol(3, "Regimen Pensionario:", props.Text{
		Align:  align.Left,
		Size:   9,
		Family: "Ubuntu",
	}),
		text.NewCol(3, persona.RegPensionario, props.Text{
			Align:  align.Left,
			Size:   8,
			Family: "Ubuntu",
		}),
		text.NewCol(3, "Condicion", props.Text{
			Align:  align.Left,
			Size:   9,
			Family: "Ubuntu",
		}),
		text.NewCol(3, persona.Condicion, props.Text{
			Align:  align.Left,
			Family: "Ubuntu",
			Size:   9,
		}),
	)

	m.AddRow(6, text.NewCol(3, "Administrador de Pension:", props.Text{
		Align:  align.Left,
		Size:   9,
		Family: "Ubuntu",
	}),
		text.NewCol(3, persona.RegPensionario, props.Text{
			Align:  align.Left,
			Size:   8,
			Family: "Ubuntu",
		}),
		text.NewCol(3, "Grupo Ocupacional:", props.Text{
			Align:  align.Left,
			Size:   9,
			Family: "Ubuntu",
		}),
		text.NewCol(3, "NO APLICA", props.Text{
			Align:  align.Left,
			Size:   9,
			Family: "Ubuntu",
		}),
	)
	m.AddRow(6, text.NewCol(3, "CUSPP:", props.Text{
		Align:  align.Left,
		Size:   9,
		Family: "Ubuntu",
	}),
		text.NewCol(3, *persona.Cussp, props.Text{
			Align:  align.Left,
			Family: "Ubuntu",
			Size:   9,
		}),
		text.NewCol(3, "Cargo Estructural:", props.Text{
			Align:  align.Left,
			Size:   9,
			Family: "Ubuntu",
		}),
		text.NewCol(3, "NO APLICA", props.Text{
			Align:  align.Left,
			Family: "Ubuntu",
			Size:   9,
		}),
	)
	m.AddRow(6, text.NewCol(6, "Cargo: ", props.Text{
		Align:  align.Left,
		Size:   9,
		Family: "Ubuntu",
	}),
		text.NewCol(3, "Jornada Laboral:", props.Text{
			Align:  align.Left,
			Size:   9,
			Family: "Ubuntu",
		}),
		text.NewCol(3, "48 HORAS", props.Text{
			Align:  align.Left,
			Size:   9,
			Family: "Ubuntu",
		}),
	)

	// err = m.RegisterFooter(getPageFooter())
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	m.AddRow(5)

	m.AddRow(7, text.NewCol(2, "Codigo", props.Text{
		Align:  align.Middle,
		Size:   9,
		Style:  fontstyle.Bold,
		Family: "Ubuntu",
	}).WithStyle(&props.Cell{BorderColor: getGrayColor(), BorderType: border.Full, BorderThickness: 0.5}),
		text.NewCol(4, "Concepto", props.Text{
			Align:  align.Middle,
			Size:   9,
			Style:  fontstyle.Bold,
			Family: "Ubuntu",
		}).WithStyle(&props.Cell{BorderColor: getGrayColor(), BorderType: border.Full, BorderThickness: 0.5}), text.NewCol(2, "Ingresos", props.Text{
			Align:           align.Middle,
			VerticalPadding: cfg.Dimensions.Height,
			Size:            9,
			Style:           fontstyle.Bold,
			Family:          "Ubuntu",
		}).WithStyle(&props.Cell{BorderColor: getGrayColor(), BorderType: border.Full, BorderThickness: 0.5}),
		text.NewCol(2, "Decuento", props.Text{
			Align:  align.Middle,
			Size:   9,
			Style:  fontstyle.Bold,
			Family: "Ubuntu",
		}).WithStyle(&props.Cell{BorderColor: getGrayColor(), BorderType: border.Full, BorderThickness: 0.5}), text.NewCol(2, "Total Neto", props.Text{
			Align:  align.Middle,
			Size:   9,
			Style:  fontstyle.Bold,
			Family: "Ubuntu",
		}).WithStyle(&props.Cell{BorderColor: getGrayColor(), BorderType: border.Full, BorderThickness: 0.5}))

	m.AddRow(5, text.NewCol(12, "Ingresos").WithStyle(&props.Cell{BorderThickness: .4, BorderColor: getGrayColor(), BorderType: border.Bottom}))

	for _, x := range persona.Ingresos {
		m.AddRow(7, text.NewCol(2, x.Codigo, props.Text{
			Align:  align.Middle,
			Size:   9,
			Family: "Ubuntu",
		}), text.NewCol(4, x.Nombre, props.Text{
			Align:  align.Left,
			Size:   8,
			Family: "Ubuntu",
		}), text.NewCol(2, x.Monto, props.Text{
			Align:  align.Center,
			Size:   9,
			Style:  fontstyle.Bold,
			Family: "Ubuntu",
		}))
	}
	m.AddRow(8, text.NewCol(12, "Descuentos").WithStyle(&props.Cell{BorderThickness: .4, BorderColor: getGrayColor(), BorderType: border.Bottom}))

	for _, x := range persona.Descuentos {
		m.AddRow(7, text.NewCol(2, x.Codigo, props.Text{
			Align:  align.Middle,
			Size:   9,
			Family: "Ubuntu",
		}), text.NewCol(4, x.Nombre, props.Text{
			Align:  align.Left,
			Style:  fontstyle.Bold,
			Size:   8,
			Family: "Ubuntu",
		}), text.NewCol(2, "", props.Text{
			Align:  align.Left,
			Size:   9,
			Family: "Ubuntu",
		}), text.NewCol(2, x.Monto, props.Text{
			Align:  align.Center,
			Size:   9,
			Style:  fontstyle.Bold,
			Family: "Ubuntu",
		}))
	}

	m.AddRow(8, text.NewCol(12, "Aporte del Trabajador").WithStyle(&props.Cell{BorderThickness: .4, BorderColor: getGrayColor(), BorderType: border.Bottom}))

	for _, x := range persona.AportesTrabajador {
		m.AddRow(7, text.NewCol(2, x.Codigo, props.Text{
			Align:  align.Middle,
			Size:   9,
			Family: "Ubuntu",
		}), text.NewCol(4, x.Nombre, props.Text{
			Align:  align.Left,
			Size:   8,
			Family: "Ubuntu",
			Style:  fontstyle.Bold,
		}), text.NewCol(2, "", props.Text{
			Align:  align.Left,
			Size:   9,
			Family: "Ubuntu",
		}), text.NewCol(2, x.Monto, props.Text{
			Align:  align.Center,
			Style:  fontstyle.Bold,
			Size:   9,
			Family: "Ubuntu",
		}))
	}
	m.AddRow(5, text.NewCol(6, "Total Neto", props.Text{Color: &props.WhiteColor, Style: fontstyle.Bold}), text.NewCol(6, persona.Neto, props.Text{Align: align.Right, Color: &props.WhiteColor, Style: fontstyle.Bold})).WithStyle(&props.Cell{BackgroundColor: getDarkGrayColor()})

	return m
}

func getPageHeader() core.Row {
	return row.New(30).Add(
		image.NewFromFileCol(4, "D:/proyectos/wails/wailsapp/src/asset/logo.png", props.Rect{
			Center:             true,
			Top:                20,
			JustReferenceWidth: true,
			Percent:            100,
		}),
		col.New(2),
		col.New(6).Add(
			text.New("Municipalidad Distrital de Villa el Salvador", props.Text{
				Align:             align.Right,
				BreakLineStrategy: breakline.DashStrategy,
				Style:             fontstyle.Normal,
				Size:              12,
				Family:            "Ubuntu",
			}),
			text.New("Ruc: 20187346488 ", props.Text{Top: 4.5, Align: align.Right, Family: "Ubuntu"}),
			text.New("Sector 2, Grupo 15, Av. Revolución S/N cruce con Av. César Vallejo", props.Text{Top: 9, Family: "Ubuntu", Align: align.Right, Size: 8}),
		),
	)
}

func getPageFooter() core.Row {
	return row.New(20).Add(
		col.New(12).Add(
			text.New("Tel: 55 024 12345-1234", props.Text{
				Top:   13,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Left,
				Color: getBlueColor(),
			}),
			text.New("www.mycompany.com", props.Text{
				Top:   16,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Left,
				Color: getBlueColor(),
			}),
		),
	)
}

func getDarkGrayColor() *props.Color {
	return &props.Color{
		Red:   55,
		Green: 55,
		Blue:  55,
	}
}

func getGrayColor() *props.Color {
	return &props.Color{
		Red:   200,
		Green: 200,
		Blue:  200,
	}
}

func getBlueColor() *props.Color {
	return &props.Color{
		Red:   10,
		Green: 10,
		Blue:  150,
	}
}
