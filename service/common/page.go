package common

// Page Binding from JSON
type Page struct {
	Page int `uri:"page" form:"page" json:"page" xml:"page"  binding:"required"`
	Size int `uri:"size" form:"size" json:"size" xml:"size" binding:"required"`
	Order string `uri:"order" form:"order" json:"order" xml:"order"`
}
