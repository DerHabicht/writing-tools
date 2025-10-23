package documents

import (
	"github.com/fxtlabs/date"

)

type DocMeta interface {
	DocType() DocType
	Title() string
	Author() string
	Date() date.Date
	Data(key string) any	
}

