package models

type PositionOrderByType uint32

const (
	name   PositionOrderByType = iota
	amount                     = iota
	price                      = iota
)

type (
	AddPositionData struct {
		Name   string `json:"name"`
		Amount uint32 `json:"amount"`
		Price  uint32 `json:"price"`
	}

	DeletePositionData struct {
		Id uint32 `json:"id"`
	}

	UpdatePositionData struct {
		Name   string `json:"name"`
		Amount uint32 `json:"amount"`
		Price  uint32 `json:"price"`
	}

	SelectSinglePositionData struct {
		Id uint32 `json:"id"`
	}

	SelectPositionsFilter struct {
		OrderBy PositionOrderByType `json:"orderBy"`
	}

	SelectPaginationData struct {
		Page   uint32 `json:"page"`
		Length uint32 `json:"length"`
		// Probably will need a key.
	}

	SelectPositionsData struct {
		Filter     SelectPositionsFilter `json:"filter"`
		Pagination SelectPaginationData  `json:"pagination"`
	}
)
