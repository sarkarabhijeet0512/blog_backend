package types

import (
	"math"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Page is pagination structure for database queries.
// It can be used for both request and response.
type Page struct {
	Page     int  `json:"current"`
	Size     int  `json:"size"`
	Total    int  `json:"total_item_count"`
	Pages    int  `json:"total"`
	Next     bool `json:"has_next"`
	Previous bool `json:"has_previous"`

	Low  int `json:"-"`
	High int `json:"-"`
}

// NewPage returns a new page object.
// It can be used for both pagination request and pagination response.
// `total` is number of total docs. Optional as not required while request.
func NewPage(page, size int) *Page {
	return &Page{
		Page: page,
		Size: size,
	}
}

// Res returns a page data to be used in API response
func (p *Page) Res(total int) (res *Page) {
	p.Total = total

	if total <= 0 {
		p.Pages = 0
	}
	p.Pages = int(math.Ceil(float64(float64(total) / float64(p.Size))))

	p.Next = p.Page < p.Pages
	p.Previous = p.Page > 1

	p.low()
	p.high(total)
	return p
}

func (p *Page) low() {
	p.Low = (p.Page - 1) * p.Size
}

func (p *Page) high(total int) {
	p.High = (p.Page * p.Size)
	if total < p.High {
		p.High = total
	}
}

type PageX struct {
	NextPageID *primitive.ObjectID `json:"next_page_id,omitempty"`
	HasNext    bool                `json:"has_next"`
}

func Pagination(pageLimit int64, docsLen int, lastDocID *primitive.ObjectID) (page *PageX) {
	page = &PageX{}
	if int64(docsLen) <= pageLimit {
		return
	}

	page = &PageX{
		NextPageID: lastDocID,
		HasNext:    true,
	}
	return
}
