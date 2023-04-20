package model

type Note struct {
	tableName   struct{} `pg:"note,discard_unknown_columns"`
	Id          int      `pg:"id,pk" json:"id"`
	Description string   `pg:"description" json:"description"`
}
