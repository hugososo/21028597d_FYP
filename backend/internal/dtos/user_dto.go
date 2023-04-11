package dtos

type UserReqDTO struct {
	Address string `json:"address" query:"address" param:"address"`
}
