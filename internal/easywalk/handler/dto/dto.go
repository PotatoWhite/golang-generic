package dto

type DTO interface {
	Validate() error
	SetModelType(modelType any)
}
