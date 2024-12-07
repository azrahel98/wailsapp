package models

import "encoding/xml"

type Persona struct {
	XMLName   xml.Name `xml:"persona"`
	Nombre    string   `xml:"nombre"`
	Apellido  string   `xml:"apellido"`
	Edad      int      `xml:"edad"`
	Direccion struct {
		Calle  string `xml:"calle"`
		Ciudad string `xml:"ciudad"`
		Pais   string `xml:"pais"`
	} `xml:"direccion"`
}
