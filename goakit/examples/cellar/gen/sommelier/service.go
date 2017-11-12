// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// sommelier service
//
// Command:
// $ goa gen goa.design/plugins/goakit/examples/cellar/design

package sommelier

import "context"

// The sommelier service retrieves bottles given a set of criteria.
type Service interface {
	// Pick implements pick.
	Pick(context.Context, *Criteria) (StoredBottleCollection, error)
}

// Criteria is the payload type of the sommelier service pick method.
type Criteria struct {
	// Name of bottle to pick
	Name *string
	// Varietals in preference order
	Varietal []string
	// Winery of bottle to pick
	Winery *string
}

// StoredBottleCollection is the result type of the sommelier service pick
// method.
type StoredBottleCollection []*StoredBottle

// A StoredBottle describes a bottle retrieved by the storage service.
type StoredBottle struct {
	// ID is the unique id of the bottle.
	ID string
	// Name of bottle
	Name string
	// Winery that produces wine
	Winery *Winery
	// Vintage of bottle
	Vintage uint32
	// Composition is the list of grape varietals and associated percentage.
	Composition []*Component
	// Description of bottle
	Description *string
	// Rating of bottle from 1 (worst) to 5 (best)
	Rating *uint32
}

type Winery struct {
	// Name of winery
	Name string
	// Region of winery
	Region string
	// Country of winery
	Country string
	// Winery website URL
	URL *string
}

type Component struct {
	// Grape varietal
	Varietal string
	// Percentage of varietal in wine
	Percentage *uint32
}

type NoCriteria struct {
	// Missing criteria
	Value string
}

type NoMatch struct {
	// No bottle matched given criteria
	Value string
}

// Error returns "no_criteria".
func (e *NoCriteria) Error() string {
	return "no_criteria"
}

// Error returns "no_match".
func (e *NoMatch) Error() string {
	return "no_match"
}