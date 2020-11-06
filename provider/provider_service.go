package provider

import "errors"

// Service to defne service for providers
type Service struct {
	Provider      `json:"provider"`
	ServiceName   string `json:"serviceName"`
	ServiceID     string `json:"serviceID"`
	ServiceRoute  string `json:"serviceRoute"`
	ServiceType   string `json:"serviceType"` //SOAP, REST, ...
	ServiceMethod string `json:"serviceMethod"`
}

// Services is the list of all Providers' services
type Services []Service

var services = Services{}

// AddServiceByProviderName adds new service by giving the provider's name
func AddServiceByProviderName(proname, sername, serid, serroute, sertype, sermethod string) error {
	// Check if the provider with proname is alreadey defined
	pro, err := GetProviderByName(proname)
	if err != nil {
		return err
	}
	// Check if the same service has already been added
	for _, ser := range services {
		switch {
		case ser.ServiceName == sername,
			ser.ServiceID == serid,
			ser.ServiceRoute == serroute:
			return errors.New("This service has already been defined")
		}
	}
	// Add the new Service to the Services slice
	services = append(services, Service{pro, sername, serid, serroute, sertype, sermethod})
	return nil
}

// AddServiceByProviderID adds new service by giving the provider's ID
func AddServiceByProviderID(proid, sername, serid, serroute, sertype, sermethod string) error {
	// Check if the provider with proname is alreadey defined
	pro, err := GetProviderByID(proid)
	if err != nil {
		return err
	}
	// Check if the same service has already been added
	for _, ser := range services {
		switch {
		case ser.ServiceName == sername,
			ser.ServiceID == serid,
			ser.ServiceRoute == serroute:
			return errors.New("This service has already been defined")
		}
	}
	// Add the new Service to the Services slice
	services = append(services, Service{pro, sername, serid, serroute, sertype, sermethod})
	return nil
}

// AddServiceByProviderBaseURL adds new service by giving the provider's Base URL
func AddServiceByProviderBaseURL(probaseurl, sername, serid, serroute, sertype, sermethod string) error {
	// Check if the provider with proname is alreadey defined
	pro, err := GetProviderByBaseURL(probaseurl)
	if err != nil {
		return err
	}
	// Check if the same service has already been added
	for _, ser := range services {
		switch {
		case ser.ServiceName == sername,
			ser.ServiceID == serid,
			ser.ServiceRoute == serroute:
			return errors.New("This service has already been defined")
		}
	}
	// Add the new Service to the Services slice
	services = append(services, Service{pro, sername, serid, serroute, sertype, sermethod})
	return nil
}
