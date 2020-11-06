package provider

import (
	"errors"
)

// Provider to define providers
type Provider struct {
	Name    string `json:"name"`
	ID      string `json:"id"`
	BaseURL string `json:"baseURL"`
}

// Providers to keep data of different added providers.
type Providers []Provider

var providers = Providers{}

// AddProvider to add new Provider to the Providers slice
func AddProvider(name, id, baseurl string) error {
	// Check if the provider has already been added.
	for _, prov := range providers {
		switch {
		case prov.Name == name, prov.ID == id, prov.BaseURL == baseurl:
			return errors.New("This provider has already been added")
		}
	}
	// Add the new Provider to the Slice of Providers
	providers = append(providers, Provider{name, id, baseurl})
	return nil
}

// GetProviderByName to return specific Provider by Its name
func GetProviderByName(name string) (Provider, error) {
	for _, pro := range providers {
		if pro.Name == name {
			return pro, nil
		}
	}
	return Provider{}, errors.New("A Provider with this name has not been added yet")
}

// GetProviderByID to return specific Provider by Its ID
func GetProviderByID(id string) (Provider, error) {
	for _, pro := range providers {
		if pro.ID == id {
			return pro, nil
		}
	}
	return Provider{}, errors.New("A Provider with this ID has not been added yet")
}

// GetProviderByBaseURL to return specific Provider by Its Base URL
func GetProviderByBaseURL(base string) (Provider, error) {
	for _, pro := range providers {
		if pro.BaseURL == base {
			return pro, nil
		}
	}
	return Provider{}, errors.New("A Provider with this Base URL has not been added yet")
}

//ToDo: func RemoveProviderByName() error {}
//ToDo: func RemoveProviderByID() error {}
//ToDo: func RemoveProviderByBase() error {}
