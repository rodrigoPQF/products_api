package products

type CreateProductsInputDto struct {
	Name  string `json:"name" validate:"required" copier:"Name"`
	Description     string `json:"description" validate:"required" copier:"Description"`
	Price   float64 `json:"price" validate:"required" copier:"Price"`
}

type CreateProductsOutputDto struct {
	Id  string `json:"id,omitempty" copier:"Id"`
	Name  string `json:"name,omitempty" copier:"Name"`
	Description     string `json:"description,omitempty" copier:"Description"`
	Price   float64 `json:"price,omitempty" copier:"Price"`
	CreatedAt  string `json:"created_at,omitempty" copier:"CreatedAt"`
	UpdatedAt  string `json:"updated_at,omitempty" copier:"UpdatedAt"`
}

type UpdateProductsInputDto struct {
	Id  string `json:"id,omitempty" copier:"Id"`
	Name  string `json:"name" validate:"required" copier:"Name"`
	Description     string `json:"description" validate:"required" copier:"Description"`
	Price   float64 `json:"price" validate:"required" copier:"Price"`
	CreatedAt  string `json:"created_at,omitempty" copier:"CreatedAt"`
	UpdatedAt  string `json:"updated_at,omitempty" copier:"UpdatedAt"`
}

type UpdateProductsOutputDto struct {
	Id  string `json:"id,omitempty" copier:"Id"`
	Name  string `json:"name,omitempty" copier:"Name"`
	Description     string `json:"description,omitempty" copier:"Description"`
	Price   float64 `json:"price,omitempty" copier:"Price"`
	CreatedAt  string `json:"created_at,omitempty" copier:"CreatedAt"`
	UpdatedAt  string `json:"updated_at,omitempty" copier:"UpdatedAt"`
}

type GetProductInputDto struct {
	Id  string `json:"id,omitempty" copier:"Id"`
}

type GetProductOutputDto struct {
	Id  string `json:"id,omitempty" copier:"Id"`
	Name  string `json:"name,omitempty" copier:"Name"`
	Description     string `json:"description,omitempty" copier:"Description"`
	Price   float64 `json:"price,omitempty" copier:"Price"`
	CreatedAt  string `json:"created_at,omitempty" copier:"CreatedAt"`
	UpdatedAt  string `json:"updated_at,omitempty" copier:"UpdatedAt"`
}


type DeleteProductInputDto struct {
	Id  string `json:"id,omitempty" copier:"Id"`
}

type DeleteProductOutputDto struct {
	Message  string `json:"message,omitempty" copier:"Message"`
}
