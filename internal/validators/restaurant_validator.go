package validators

type CreateRestaurantRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Code    string `json:"code" binding:"required"`
}
