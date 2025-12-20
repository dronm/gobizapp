package models

// This file containes structures for unmarshaling
// values from client queries.
// It is an abstraction over sql structures
// for using in different transport protocols: json/xml etc.

type SgnParam string
type SortParam string
type FilterJoinParam string

// client query values
const (
	SGN_PAR_E       SgnParam = "e"       //equal
	SGN_PAR_L       SgnParam = "l"       //less
	SGN_PAR_G       SgnParam = "g"       //greater
	SGN_PAR_LE      SgnParam = "le"      //less and equal
	SGN_PAR_GE      SgnParam = "ge"      //greater and equal
	SGN_PAR_LK      SgnParam = "lk"      //like
	SGN_PAR_NE      SgnParam = "ne"      //not equal
	SGN_PAR_I       SgnParam = "i"       // IS
	SGN_PAR_IN      SgnParam = "in"      // in
	SGN_PAR_INCL    SgnParam = "incl"    //include
	SGN_PAR_ANY     SgnParam = "any"     //Any
	SGN_PAR_OVERLAP SgnParam = "overlap" //overlap
)

const (
	SORT_PAR_ASC  SortParam = "a" // asc
	SORT_PAR_DESC SortParam = "d" // desc
)

const (
	FILTER_PAR_JOIN_AND FilterJoinParam = "and"
	FILTER_PAR_JOIN_OR  FilterJoinParam = "or"
)

type ListSort struct {
	Field  string    `json:"f"`
	Direct SortParam `json:"d"`
}

type ListFilter struct {
	Field string          `json:"f" require:"true" maxLength:"50"`
	Val   string          `json:"v" require:"true"`
	Sign  SgnParam        `json:"s" require:"true" maxLength:"10"`
	Join  FilterJoinParam `json:"j" maxLength:"1"`
}

type ListSorts []ListSort
type ListFilters []ListFilter
type ListFrom int
type ListCount int
