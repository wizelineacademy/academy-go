package translators

import (
	"api-bootcamp/controllers/viewmodels"
	"api-bootcamp/dto"
)

func ToResponseView(input dto.GetCsvDTO) viewmodels.ResponseView {
	return viewmodels.ResponseView{
		ID:   input.ID,
		Name: input.Name,
	}
}
