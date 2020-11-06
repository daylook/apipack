package provider

import (
	"errors"
)

// Provider to define providers
type Provider struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Base string `json:"base"`
}

// Providers to keep datas of different added providers.
var Providers = []Provider{}

// AddProvider to add new Provider to the Providers slice
func AddProvider(name, id, base string) error {
	// Check if the provider has already been added.
	for _, prov := range Providers {
		if prov.Name == name {
			return errors.New("The Provider with the same name has already been added")
		}
	}
	// Add the new Provider to the Slice of Providers
	Providers = append(Providers, Provider{name, id, base})
	return nil
}
