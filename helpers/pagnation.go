package helpers

import (
	"math"

	"github.com/robbyklein/pages/initializers"
)

type PaginationData struct {
	NextPage     int
	PreviousPage int
	CurrentPage  int
	TotalPages   int
	TwoAfter     int
	TwoBelow     int
	ThreeAfter   int
	Offset       int
	BaseURL      string
}

func GetPaginationData(page int, perPage int, model interface{}, baseUrl string) (PaginationData, error) {
	// Calculate total pages
	var totalRows int64
	result := initializers.DB.Model(model).Count(&totalRows)
	if result.Error != nil {
		return PaginationData{}, result.Error
	}
	totalPages := int(math.Ceil(float64(totalRows) / float64(perPage)))

	// Calculate offset
	offset := (page - 1) * perPage

	return PaginationData{
		NextPage:     page + 1,
		PreviousPage: page - 1,
		CurrentPage:  page,
		TotalPages:   totalPages,
		TwoAfter:     page + 2,
		TwoBelow:     page - 2,
		ThreeAfter:   page + 3,
		Offset:       offset,
		BaseURL:      baseUrl,
	}, nil
}
