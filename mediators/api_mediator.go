package mediators

import (
	"os"

	"api-bootcamp/dto"

	"github.com/gocarina/gocsv"
)

type ApiMediator interface {
	GetCSVElementByID(id int) (dto.GetCsvDTO, error)
}

type apiMediator struct {
}

func NewApiMediator() ApiMediator {
	return &apiMediator{}
}

func (ap *apiMediator) GetCSVElementByID(id int) (dto.GetCsvDTO, error) {
	in, err := os.Open("csv/pokemon.csv")
	csvDTo := []*dto.GetCsvDTO{}
	if err != nil {
		return dto.GetCsvDTO{}, err
	}
	defer in.Close()

	if err := gocsv.UnmarshalFile(in, &csvDTo); err != nil {
		return dto.GetCsvDTO{}, err
	}
	for _, client := range csvDTo {
		if client.ID == id {
			return *client, nil
		}
	}

	return dto.GetCsvDTO{}, nil
}
