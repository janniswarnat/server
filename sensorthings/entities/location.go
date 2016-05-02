package entities

import (
	"encoding/json"
)

// Location entity locates the Thing or the Things it associated with. A Thing’s Location entity is
// defined as the last known location of the Thing.
// A Thing’s Location may be identical to the Thing’s Observations’ FeatureOfInterest. In the context of the IoT,
// the principle location of interest is usually associated with the location of the Thing, especially for in-situ
// sensing applications. For example, the location of interest of a wifi-connected thermostat should be the building
// or the room in which the smart thermostat is located. And the FeatureOfInterest of the Observations made by the
// thermostat (e.g., room temperature readings) should also be the building or the room. In this case, the content
// of the smart thermostat’s location should be the same as the content of the temperature readings’ feature of interest.
type Location struct {
	ID                     string                `json:"@iot.id"`
	NavSelf                string                `json:"@iot.selfLink"`
	Description            string                `json:"description"`
	EncodingType           string                `json:"encodingtype"`
	Location               string                `json:"location"`
	NavThings              string                `json:"Things@iot.navigationLink,omitempty"`
	NavHistoricalLocations string                `json:"HistoricalLocations@iot.navigationLink,omitempty"`
	Things                 []*Thing              `json:"Things,omitempty"`
	HistoricalLocations    []*HistoricalLocation `json:"HistoricalLocations,omitempty"`
}

// GetEntityType returns the EntityType for Location
func (l *Location) GetEntityType() EntityType {
	return EntityTypeLocation
}

// ParseEntity tries to parse the given json byte array into the current entity
func (l *Location) ParseEntity(data []byte) error {
	location := &l
	err := json.Unmarshal(data, location)
	if err != nil {
		return err
	}

	return nil
}

// ContainsMandatoryParams checks if all mandatory params for Location are available before posting.
func (l *Location) ContainsMandatoryParams() (bool, []error) {
	err := []error{}
	CheckMandatoryParam(&err, l.Description, l.GetEntityType(), "description")
	CheckMandatoryParam(&err, l.EncodingType, l.GetEntityType(), "encodingtype")
	CheckMandatoryParam(&err, l.Location, l.GetEntityType(), "location")

	if len(err) != 0 {
		return false, err
	}

	return true, nil
}

// SetLinks sets the entity specific navigation links if needed
func (l *Location) SetLinks(externalURL string) {
	l.NavSelf = CreateEntitySefLink(externalURL, EntityLinkLocations.ToString(), l.ID)
	l.NavThings = CreateEntityLink(l.Things == nil, EntityLinkLocations.ToString(), EntityLinkThings.ToString(), l.ID)
	l.NavHistoricalLocations = CreateEntityLink(l.HistoricalLocations == nil, EntityLinkLocations.ToString(), EntityLinkHistoricalLocations.ToString(), l.ID)
}