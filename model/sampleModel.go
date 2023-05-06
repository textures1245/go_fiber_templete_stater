package model

type SampleModel struct {
	ID      string `json:"ID"`
	Column1 string `json:"Column1"`
	Column2 string `json:"Column2"`
}

var SQL_simple_add = `INSERT INTO some_table (some_column_1 ,some_column_2) 
	VALUES(@p1, @p2);`

var SQL_simple_get_date = `select GETDATE();`
