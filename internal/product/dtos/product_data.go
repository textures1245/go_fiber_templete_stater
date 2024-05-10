package dtos

type ProductDataRes struct {
	Id            int64   `json:"id" db:"id"`
	Name          string  `json:"name" db:"product_name"`
	Description   string  `json:"description" db:"product_description"`
	Price         float64 `json:"price" db:"product_price"`
	Stock         int64   `json:"stock" db:"stock_qty"`
	CreatedAt     string  `json:"created_at" db:"created_at"`
	UpdatedAt     string  `json:"updated_at" db:"updated_at"`
	FileName      *string `db:"file_name" json:"file_name"`
	FileBase64Url *[]byte `db:"file_data" json:"file_base64_url"`
	FileType      *string `db:"file_type" json:"file_type"`
	FileUrl       *string `json:"file_url"`
}
