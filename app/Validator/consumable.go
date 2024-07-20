package validator

type CreateConsumableBody struct {
	Title     string  `json:"title" binding:"required"`
	Remaining float64 `json:"remaining" binding:"required"`
	UOM       string  `json:"uom" binding:"required"`
}
