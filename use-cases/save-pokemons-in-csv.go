package usecases

import (
	"encoding/csv"
	"github.com/georgexx009-wizeline/ondemand-go-bootcamp/model"
	logger "github.com/georgexx009-wizeline/ondemand-go-bootcamp/utils"
	"os"
	"reflect"
	"strconv"
)

func (*UseCases) SavePokemonsInCSV(fileName string, pokemons []model.Pokemon) error {
	csvFile, err := os.Create(fileName)
	defer func() {
		err = csvFile.Close()
		if err != nil {
			logger.Log(err.Error())
		}
	}()
	if err != nil {
		logger.Log(err.Error())
		return err
	}

	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	fields := getFieldsFromStruct(pokemons[0])
	data := [][]string{fields}

	for _, pokemon := range pokemons {
		row := []string{strconv.Itoa(pokemon.Id), pokemon.Name, strconv.Itoa(pokemon.Weight)}
		data = append(data, row)
	}
	err = csvWriter.WriteAll(data)
	if err != nil {
		logger.Log(err.Error())
		return err
	}
	return nil
}

func getFieldsFromStruct(s model.Pokemon) []string {
	var fields []string
	e := reflect.ValueOf(&s).Elem()

	for i := 0; i < e.NumField(); i++ {
		fields = append(fields, e.Type().Field(i).Name)
	}

	return fields
}
