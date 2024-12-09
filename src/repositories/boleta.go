package repositories

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"vesgoapp/src/models"
)

type BoletaRepository interface {
	ReadXmls_folder() (*[]models.ListOfFileDirectory, error)
	Extraer_Datos(path string) (*models.Boleta, error)
}
type boletaRepository struct {
}

func (b *boletaRepository) ReadXmls_folder() (*[]models.ListOfFileDirectory, error) {

	var list []models.ListOfFileDirectory
	files, err := os.ReadDir("D:/xmls")

	if err != nil {
		return nil, err
	}

	for _, file := range files {
		list = append(list, models.ListOfFileDirectory{
			Dni:  strings.Split(file.Name(), "_")[3],
			Mes:  strings.Split(file.Name(), "_")[2],
			Año:  strings.Split(file.Name(), "_")[1],
			Full: file.Name(),
		})
	}
	return &list, nil
}

func (b *boletaRepository) Extraer_Datos(path string) (*models.Boleta, error) {

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()
	decoder := xml.NewDecoder(file)

	var insideWorksheet bool
	var table models.Table

	// Procesar el XML token por token
	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}

		switch elem := token.(type) {
		case xml.StartElement:

			if elem.Name.Local == "Worksheet" && elem.Name.Space == "urn:schemas-microsoft-com:office:spreadsheet" {
				insideWorksheet = true
			}
			if insideWorksheet && elem.Name.Local == "Table" && elem.Name.Space == "urn:schemas-microsoft-com:office:spreadsheet" {
				if err := decoder.DecodeElement(&table, &elem); err != nil {
					fmt.Printf("Error al decodificar la tabla: %v\n", err)
					return nil, err
				}
			}
		case xml.EndElement:
			if elem.Name.Local == "Worksheet" && elem.Name.Space == "urn:schemas-microsoft-com:office:spreadsheet" {
				insideWorksheet = false
			}
		}
	}

	// Imprime los datos decodificados
	var boleta models.Boleta

	var ingresos int
	var descuento int
	var aportes int
	var empleador int
	var neto int

	for index, row := range table.Rows {
		for e, cell := range row.Cells {
			// fmt.Printf(" Row %d Cell %d Dato : %s \n", index, e, cell.Data)
			if cell.Data == "Ingresos" {
				ingresos = index + 1
			}
			if cell.Data == "Descuentos" {
				descuento = index + 1
			}
			if cell.Data == "Aportes del Trabajador" {
				aportes = index + 1
			}
			if cell.Data == "Aportes de Empleador" {
				empleador = index + 1
			}
			if cell.Data == "Neto a Pagar" {
				neto = index
			}
			if index == 4 && e == 0 {
				boleta.Fecha = strings.TrimSpace(strings.Split(cell.Data, ":")[1])
			}
			if index == 2 && e == 0 {
				boleta.Ruc = strings.TrimSpace(strings.Split(cell.Data, ":")[1])
			}
			if index == 3 && e == 0 {
				boleta.Empleador = strings.TrimSpace(strings.Split(cell.Data, ":")[1])
			}
			if index == 8 && e == 0 {
				boleta.TipoDocIdentidad = cell.Data
			}
			if index == 8 && e == 1 {
				boleta.NumeroDoc = cell.Data
			}
			if index == 8 && e == 2 {
				boleta.Nombre = cell.Data
			}
			if index == 8 && e == 3 {
				boleta.Situacion = cell.Data
			}
			if index == 10 && e == 0 {
				boleta.Ingreso = cell.Data
			}
			if index == 10 && e == 1 {
				boleta.TipoTrabajador = cell.Data
			}
			if index == 10 && e == 2 {
				boleta.RegPensionario = cell.Data
			}
			if index == 10 && e == 3 {
				boleta.Cussp = &cell.Data
			}
			if index == 13 && e == 0 {
				boleta.DiasLab = cell.Data
			}
			if index == 13 && e == 1 {
				boleta.DiasnoLab = cell.Data
			}
			if index == 13 && e == 2 {
				boleta.DiasSubs = cell.Data
			}
			if index == 13 && e == 3 {
				boleta.Condicion = cell.Data
			}
			if index == 13 && e == 4 {
				boleta.Holaslab = cell.Data
			}
			if index == 16 && e == 0 {
				boleta.TipoSuspe = cell.Data
			}
			if index == 16 && e == 1 {
				boleta.MoitvoSuspe = cell.Data
			}
			if index == 16 && e == 2 {
				boleta.DiasSuspe = cell.Data
			}
			if index == 16 && e == 0 {
				boleta.TipoSuspe = cell.Data
			}
			if index == 29 && e == 1 {
				boleta.TipoSuspe = cell.Data
			}
		}
	}
	var itemIngre []models.Item
	var itemDesc []models.Item
	var itemAportTra []models.Item
	var itemAportem []models.Item
	for e, row := range table.Rows {
		if e >= ingresos && e < descuento-1 {

			var it models.Item

			it.Codigo = row.Cells[0].Data
			it.Nombre = row.Cells[1].Data
			it.Monto = row.Cells[2].Data
			itemIngre = append(itemIngre, it)
		}
		if e >= descuento && e < aportes-1 {
			var it models.Item

			it.Codigo = row.Cells[0].Data
			it.Nombre = row.Cells[1].Data
			it.Monto = row.Cells[2].Data
			itemDesc = append(itemDesc, it)
		}
		if e >= aportes && e < neto {
			var it models.Item

			it.Codigo = row.Cells[0].Data
			it.Nombre = row.Cells[1].Data
			it.Monto = row.Cells[2].Data
			itemAportTra = append(itemAportTra, it)
		}

		if e >= empleador && e < len(table.Rows)-1 {
			var it models.Item

			it.Codigo = row.Cells[0].Data
			it.Nombre = row.Cells[1].Data
			it.Monto = row.Cells[2].Data
			itemAportem = append(itemAportem, it)
		}
	}
	boleta.Neto = table.Rows[neto].Cells[1].Data
	boleta.Ingresos = itemIngre
	boleta.Descuentos = itemDesc
	boleta.AporteEmpleador = itemAportem
	boleta.AportesTrabajador = itemAportTra

	return &boleta, nil

}

func CreateboletaRepository() BoletaRepository {
	return &boletaRepository{}
}
