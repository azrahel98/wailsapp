package repositories

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type BoletaRepository interface {
	Upload_file(path string) (bool, error)
}
type boletaRepository struct {
}

func (b *boletaRepository) Upload_file(path string) (bool, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	sheetNames := f.GetSheetList()
	fmt.Println("Hojas del Excel:")
	for w, nombre := range sheetNames {
		fmt.Println("- ", nombre)
		fmt.Println(w)
	}

	return true, nil

}

func CreateboletaRepository() BoletaRepository {
	return &boletaRepository{}
}
