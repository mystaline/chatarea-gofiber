package dto

type BaseQuery struct {
	Page      int64  `query:"page,omitempty"`
	Limit     int64  `query:"limit,omitempty"`
	SortBy    string `query:"sortBy,omitempty"`
	SortOrder string `query:"sortOrder,omitempty"`
}
