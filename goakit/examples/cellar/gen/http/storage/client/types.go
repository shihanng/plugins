// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// storage HTTP client types
//
// Command:
// $ goa gen goa.design/plugins/goakit/examples/cellar/design

package client

import (
	"unicode/utf8"

	goa "goa.design/goa"
	storage "goa.design/plugins/goakit/examples/cellar/gen/storage"
)

// AddRequestBody is the type of the "storage" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	// Name of bottle
	Name string `form:"name" json:"name" xml:"name"`
	// Winery that produces wine
	Winery *WineryRequestBody `form:"winery" json:"winery" xml:"winery"`
	// Vintage of bottle
	Vintage uint32 `form:"vintage" json:"vintage" xml:"vintage"`
	// Composition is the list of grape varietals and associated percentage.
	Composition []*ComponentRequestBody `form:"composition,omitempty" json:"composition,omitempty" xml:"composition,omitempty"`
	// Description of bottle
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Rating of bottle from 1 (worst) to 5 (best)
	Rating *uint32 `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
}

// ListResponseBody is the type of the "storage" service "list" endpoint HTTP
// response body.
type ListResponseBody []*StoredBottleResponseBody

// ShowResponseBody is the type of the "storage" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// ID is the unique id of the bottle.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of bottle
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Winery that produces wine
	Winery *Winery `form:"winery,omitempty" json:"winery,omitempty" xml:"winery,omitempty"`
	// Vintage of bottle
	Vintage *uint32 `form:"vintage,omitempty" json:"vintage,omitempty" xml:"vintage,omitempty"`
	// Composition is the list of grape varietals and associated percentage.
	Composition []*Component `form:"composition,omitempty" json:"composition,omitempty" xml:"composition,omitempty"`
	// Description of bottle
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Rating of bottle from 1 (worst) to 5 (best)
	Rating *uint32 `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
}

// ShowNotFoundResponseBody is the type of the "storage" service "show"
// endpoint HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// ID of missing bottle
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// StoredBottleResponseBody is used to define fields on response body types.
type StoredBottleResponseBody struct {
	// ID is the unique id of the bottle.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of bottle
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Winery that produces wine
	Winery *WineryResponseBody `form:"winery,omitempty" json:"winery,omitempty" xml:"winery,omitempty"`
	// Vintage of bottle
	Vintage *uint32 `form:"vintage,omitempty" json:"vintage,omitempty" xml:"vintage,omitempty"`
	// Composition is the list of grape varietals and associated percentage.
	Composition []*ComponentResponseBody `form:"composition,omitempty" json:"composition,omitempty" xml:"composition,omitempty"`
	// Description of bottle
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Rating of bottle from 1 (worst) to 5 (best)
	Rating *uint32 `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
}

// WineryResponseBody is used to define fields on response body types.
type WineryResponseBody struct {
	// Name of winery
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Region of winery
	Region *string `form:"region,omitempty" json:"region,omitempty" xml:"region,omitempty"`
	// Country of winery
	Country *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	// Winery website URL
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}

// ComponentResponseBody is used to define fields on response body types.
type ComponentResponseBody struct {
	// Grape varietal
	Varietal *string `form:"varietal,omitempty" json:"varietal,omitempty" xml:"varietal,omitempty"`
	// Percentage of varietal in wine
	Percentage *uint32 `form:"percentage,omitempty" json:"percentage,omitempty" xml:"percentage,omitempty"`
}

// Winery is used to define fields on response body types.
type Winery struct {
	// Name of winery
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Region of winery
	Region *string `form:"region,omitempty" json:"region,omitempty" xml:"region,omitempty"`
	// Country of winery
	Country *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	// Winery website URL
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}

// Component is used to define fields on response body types.
type Component struct {
	// Grape varietal
	Varietal *string `form:"varietal,omitempty" json:"varietal,omitempty" xml:"varietal,omitempty"`
	// Percentage of varietal in wine
	Percentage *uint32 `form:"percentage,omitempty" json:"percentage,omitempty" xml:"percentage,omitempty"`
}

// WineryRequestBody is used to define fields on request body types.
type WineryRequestBody struct {
	// Name of winery
	Name string `form:"name" json:"name" xml:"name"`
	// Region of winery
	Region string `form:"region" json:"region" xml:"region"`
	// Country of winery
	Country string `form:"country" json:"country" xml:"country"`
	// Winery website URL
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}

// ComponentRequestBody is used to define fields on request body types.
type ComponentRequestBody struct {
	// Grape varietal
	Varietal string `form:"varietal" json:"varietal" xml:"varietal"`
	// Percentage of varietal in wine
	Percentage *uint32 `form:"percentage,omitempty" json:"percentage,omitempty" xml:"percentage,omitempty"`
}

// NewAddRequestBody builds the HTTP request body from the payload of the "add"
// endpoint of the "storage" service.
func NewAddRequestBody(p *storage.Bottle) *AddRequestBody {
	body := &AddRequestBody{
		Name:        p.Name,
		Vintage:     p.Vintage,
		Description: p.Description,
		Rating:      p.Rating,
	}
	if p.Winery != nil {
		body.Winery = marshalWineryToWineryRequestBody(p.Winery)
	}
	if p.Composition != nil {
		body.Composition = make([]*ComponentRequestBody, len(p.Composition))
		for j, val := range p.Composition {
			body.Composition[j] = &ComponentRequestBody{
				Varietal:   val.Varietal,
				Percentage: val.Percentage,
			}
		}
	}
	return body
}

// NewListStoredBottleCollectionOK builds a "storage" service "list" endpoint
// result from a HTTP "OK" response.
func NewListStoredBottleCollectionOK(body ListResponseBody) storage.StoredBottleCollection {
	v := make([]*storage.StoredBottle, len(body))
	for i, val := range body {
		v[i] = &storage.StoredBottle{
			ID:          *val.ID,
			Name:        *val.Name,
			Vintage:     *val.Vintage,
			Description: val.Description,
			Rating:      val.Rating,
		}
		v[i].Winery = unmarshalWineryResponseBodyToWinery(val.Winery)
		v[i].Composition = make([]*storage.Component, len(val.Composition))
		for j, val := range val.Composition {
			v[i].Composition[j] = &storage.Component{
				Varietal:   *val.Varietal,
				Percentage: val.Percentage,
			}
		}
	}
	return v
}

// NewShowStoredBottleOK builds a "storage" service "show" endpoint result from
// a HTTP "OK" response.
func NewShowStoredBottleOK(body *ShowResponseBody) *storage.StoredBottle {
	v := &storage.StoredBottle{
		ID:          *body.ID,
		Name:        *body.Name,
		Vintage:     *body.Vintage,
		Description: body.Description,
		Rating:      body.Rating,
	}
	v.Winery = unmarshalWineryToWinery(body.Winery)
	v.Composition = make([]*storage.Component, len(body.Composition))
	for j, val := range body.Composition {
		v.Composition[j] = &storage.Component{
			Varietal:   *val.Varietal,
			Percentage: val.Percentage,
		}
	}
	return v
}

// NewShowNotFound builds a storage service show endpoint not_found error.
func NewShowNotFound(body *ShowNotFoundResponseBody) *storage.NotFound {
	v := &storage.NotFound{
		Message: *body.Message,
		ID:      *body.ID,
	}
	return v
}

// Validate runs the validations defined on ListResponseBody
func (body ListResponseBody) Validate() (err error) {
	for _, e := range body {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Validate runs the validations defined on ShowResponseBody
func (body *ShowResponseBody) Validate() (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Winery == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("winery", "body"))
	}
	if body.Vintage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("vintage", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	if body.Winery != nil {
		if err2 := body.Winery.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Vintage != nil {
		if *body.Vintage < 1900 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.vintage", *body.Vintage, 1900, true))
		}
	}
	if body.Vintage != nil {
		if *body.Vintage > 2020 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.vintage", *body.Vintage, 2020, false))
		}
	}
	for _, e := range body.Composition {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if body.Description != nil {
		if utf8.RuneCountInString(*body.Description) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", *body.Description, utf8.RuneCountInString(*body.Description), 2000, false))
		}
	}
	if body.Rating != nil {
		if *body.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.rating", *body.Rating, 1, true))
		}
	}
	if body.Rating != nil {
		if *body.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.rating", *body.Rating, 5, false))
		}
	}
	return
}

// Validate runs the validations defined on ShowNotFoundResponseBody
func (body *ShowNotFoundResponseBody) Validate() (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	return
}

// Validate runs the validations defined on StoredBottleResponseBody
func (body *StoredBottleResponseBody) Validate() (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Winery == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("winery", "body"))
	}
	if body.Vintage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("vintage", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	if body.Winery != nil {
		if err2 := body.Winery.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Vintage != nil {
		if *body.Vintage < 1900 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.vintage", *body.Vintage, 1900, true))
		}
	}
	if body.Vintage != nil {
		if *body.Vintage > 2020 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.vintage", *body.Vintage, 2020, false))
		}
	}
	for _, e := range body.Composition {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if body.Description != nil {
		if utf8.RuneCountInString(*body.Description) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", *body.Description, utf8.RuneCountInString(*body.Description), 2000, false))
		}
	}
	if body.Rating != nil {
		if *body.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.rating", *body.Rating, 1, true))
		}
	}
	if body.Rating != nil {
		if *body.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.rating", *body.Rating, 5, false))
		}
	}
	return
}

// Validate runs the validations defined on WineryResponseBody
func (body *WineryResponseBody) Validate() (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Region == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("region", "body"))
	}
	if body.Country == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("country", "body"))
	}
	if body.Region != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.region", *body.Region, "(?i)[a-z '\\.]+"))
	}
	if body.Country != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.country", *body.Country, "(?i)[a-z '\\.]+"))
	}
	if body.URL != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.url", *body.URL, "(?i)^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	}
	return
}

// Validate runs the validations defined on ComponentResponseBody
func (body *ComponentResponseBody) Validate() (err error) {
	if body.Varietal == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("varietal", "body"))
	}
	if body.Varietal != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.varietal", *body.Varietal, "[A-Za-z' ]+"))
	}
	if body.Varietal != nil {
		if utf8.RuneCountInString(*body.Varietal) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.varietal", *body.Varietal, utf8.RuneCountInString(*body.Varietal), 100, false))
		}
	}
	if body.Percentage != nil {
		if *body.Percentage < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.percentage", *body.Percentage, 1, true))
		}
	}
	if body.Percentage != nil {
		if *body.Percentage > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.percentage", *body.Percentage, 100, false))
		}
	}
	return
}

// Validate runs the validations defined on Winery
func (body *Winery) Validate() (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Region == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("region", "body"))
	}
	if body.Country == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("country", "body"))
	}
	if body.Region != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.region", *body.Region, "(?i)[a-z '\\.]+"))
	}
	if body.Country != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.country", *body.Country, "(?i)[a-z '\\.]+"))
	}
	if body.URL != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.url", *body.URL, "(?i)^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	}
	return
}

// Validate runs the validations defined on Component
func (body *Component) Validate() (err error) {
	if body.Varietal == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("varietal", "body"))
	}
	if body.Varietal != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.varietal", *body.Varietal, "[A-Za-z' ]+"))
	}
	if body.Varietal != nil {
		if utf8.RuneCountInString(*body.Varietal) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.varietal", *body.Varietal, utf8.RuneCountInString(*body.Varietal), 100, false))
		}
	}
	if body.Percentage != nil {
		if *body.Percentage < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.percentage", *body.Percentage, 1, true))
		}
	}
	if body.Percentage != nil {
		if *body.Percentage > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.percentage", *body.Percentage, 100, false))
		}
	}
	return
}

// Validate runs the validations defined on WineryRequestBody
func (body *WineryRequestBody) Validate() (err error) {
	err = goa.MergeErrors(err, goa.ValidatePattern("body.region", body.Region, "(?i)[a-z '\\.]+"))
	err = goa.MergeErrors(err, goa.ValidatePattern("body.country", body.Country, "(?i)[a-z '\\.]+"))
	if body.URL != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.url", *body.URL, "(?i)^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	}
	return
}

// Validate runs the validations defined on ComponentRequestBody
func (body *ComponentRequestBody) Validate() (err error) {
	err = goa.MergeErrors(err, goa.ValidatePattern("body.varietal", body.Varietal, "[A-Za-z' ]+"))
	if utf8.RuneCountInString(body.Varietal) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.varietal", body.Varietal, utf8.RuneCountInString(body.Varietal), 100, false))
	}
	if body.Percentage != nil {
		if *body.Percentage < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.percentage", *body.Percentage, 1, true))
		}
	}
	if body.Percentage != nil {
		if *body.Percentage > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.percentage", *body.Percentage, 100, false))
		}
	}
	return
}