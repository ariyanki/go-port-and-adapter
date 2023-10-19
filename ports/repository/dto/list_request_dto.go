package dto

type ListRequestDto struct {
	Filter   map[string]string `json:"filter"`
	OrderBy  map[string]string `json:"order_by"`
	PageSize int               `json:"page_size"`
	Page     int               `json:"page"`
}
