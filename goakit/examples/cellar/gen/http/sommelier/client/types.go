// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// sommelier HTTP client types
//
// Command:
// $ goa gen goa.design/plugins/goakit/examples/cellar/design

package client

import (
	"unicode/utf8"

	goa "goa.design/goa"
	sommelier "goa.design/plugins/goakit/examples/cellar/gen/sommelier"
)

// PickRequestBody is the type of the "sommelier" service "pick" endpoint HTTP
// request body.
type PickRequestBody struct {
	// Name of bottle to pick
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Varietals in preference order
	Varietal []string `form:"varietal,omitempty" json:"varietal,omitempty" xml:"varietal,omitempty"`
	// Winery of bottle to pick
	Winery *string `form:"winery,omitempty" json:"winery,omitempty" xml:"winery,omitempty"`
}

// PickResponseBody is the type of the "sommelier" service "pick" endpoint HTTP
// response body.
type PickResponseBody []*StoredBottleResponseBody

// PickNoCriteriaResponseBody is the type of the "sommelier" service "pick"
// endpoint HTTP response body for the "no_criteria" error.
type PickNoCriteriaResponseBody struct {
	// Missing criteria
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// PickNoMatchResponseBody is the type of the "sommelier" service "pick"
// endpoint HTTP response body for the "no_match" error.
type PickNoMatchResponseBody struct {
	// No bottle matched given criteria
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
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

// NewPickRequestBody builds the HTTP request body from the payload of the
// "pick" endpoint of the "sommelier" service.
func NewPickRequestBody(p *sommelier.Criteria) *PickRequestBody {
	body := &PickRequestBody{
		Name:   p.Name,
		Winery: p.Winery,
	}
	if p.Varietal != nil {
		body.Varietal = make([]string, len(p.Varietal))
		for j, val := range p.Varietal {
			body.Varietal[j] = val
		}
	}
	return body
}

// NewPickStoredBottleCollectionOK builds a "sommelier" service "pick" endpoint
// result from a HTTP "OK" response.
func NewPickStoredBottleCollectionOK(body PickResponseBody) sommelier.StoredBottleCollection {
	v := make([]*sommelier.StoredBottle, len(body))
	for i, val := range body {
		v[i] = &sommelier.StoredBottle{
			ID:          *val.ID,
			Name:        *val.Name,
			Vintage:     *val.Vintage,
			Description: val.Description,
			Rating:      val.Rating,
		}
		v[i].Winery = unmarshalWineryResponseBodyToWinery(val.Winery)
		v[i].Composition = make([]*sommelier.Component, len(val.Composition))
		for j, val := range val.Composition {
			v[i].Composition[j] = &sommelier.Component{
				Varietal:   *val.Varietal,
				Percentage: val.Percentage,
			}
		}
	}
	return v
}

// NewPickNoCriteria builds a sommelier service pick endpoint no_criteria error.
func NewPickNoCriteria(body *PickNoCriteriaResponseBody) *sommelier.NoCriteria {
	v := &sommelier.NoCriteria{
		Value: *body.Value,
	}
	return v
}

// NewPickNoMatch builds a sommelier service pick endpoint no_match error.
func NewPickNoMatch(body *PickNoMatchResponseBody) *sommelier.NoMatch {
	v := &sommelier.NoMatch{
		Value: *body.Value,
	}
	return v
}

// Validate runs the validations defined on PickResponseBody
func (body PickResponseBody) Validate() (err error) {
	for _, e := range body {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Validate runs the validations defined on PickNoCriteriaResponseBody
func (body *PickNoCriteriaResponseBody) Validate() (err error) {
	if body.Value == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("value", "body"))
	}
	return
}

// Validate runs the validations defined on PickNoMatchResponseBody
func (body *PickNoMatchResponseBody) Validate() (err error) {
	if body.Value == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("value", "body"))
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