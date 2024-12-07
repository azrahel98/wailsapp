package repositories

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type BoletaRepository interface {
	Upload_file(path string) (bool, error)
}
type boletaRepository struct {
}

type Workbook struct {
	Worksheet Worksheet `xml:"ss:Worksheet"`
}

// Worksheet representa el elemento ss:Worksheet del archivo XML
type Worksheet struct {
	Table Table `xml:"ss:Table"`
}

// Table representa el elemento ss:Table del archivo XML
type Table struct {
	Rows []Row `xml:"ss:Row"`
}

// Row representa el elemento ss:Row del archivo XML
type Row struct {
	Index string `xml:"ss:Index,attr"`
	Cells []Cell `xml:"ss:Cell"`
}

// Cell representa el elemento ss:Cell del archivo XML
type Cell struct {
	Index string `xml:"ss:Index,attr"`
	Data  Data   `xml:"ss:Data"`
}

// Data representa el valor de una celda en el archivo XML
type Data struct {
	Value string `xml:",chardata"`
	Type  string `xml:"ss:Type,attr"`
}

func obtenerValorXML(archivo string, rowIndex string, cellIndex string) (string, error) {
	// Abrir el archivo XML
	xmlFile, err := os.Open(archivo)
	if err != nil {
		return "", err
	}
	defer xmlFile.Close()

	// Leer el contenido del archivo XML
	contenido, err := io.ReadAll(xmlFile)
	if err != nil {
		return "", err
	}

	// Deserializar el contenido XML en una estructura Workbook
	var workbook Workbook
	err = xml.Unmarshal(contenido, &workbook)
	if err != nil {
		return "", err
	}

	fmt.Println(workbook)

	// Buscar la fila y celda correspondientes
	for _, row := range workbook.Worksheet.Table.Rows {
		if row.Index == rowIndex {
			for _, cell := range row.Cells {
				if cell.Index == cellIndex {
					fmt.Println(cell.Data.Value)
					return cell.Data.Value, nil
				}
			}
		}
	}

	// Si no se encuentra la fila o celda, devolver un error
	return "", fmt.Errorf("no se encontró la fila o celda correspondiente")
}

func (b *boletaRepository) Upload_file(path string) (bool, error) {

	x, err := obtenerValorXML(path, "2", "9")

	if err != nil {
		fmt.Println("El erroorooooooooooooo esta aquiii")
		fmt.Println(err)
		return false, err
	}

	fmt.Println("Estoy aquiiiiiiiiiiiiiiiii")
	fmt.Println(x)

	return true, nil

}

func CreateboletaRepository() BoletaRepository {
	return &boletaRepository{}
}
