package service

/*
NewService()
AddServiceConfigFile() to add in apibox/config/service/serviceName.yaml
AddServiceErrorFile() to add in apibox/error/service/serviceName.yaml
AddServiceLogFile() to add in apibox/log/service/serviceName.yaml
ReadService() to get all the fields structure of each service definition.
EditSerivce() to modify the fields and structure of the service definition.
StopService()
StartService()
DeleteService()

==================
SortService(): by value, by provider, by priority, by time

SearchService(): by value, by provider, by priority, by time

ServiceRoutes(): Define routes for each service API

ServiceProxy(): adapt Service API to fit other services

ServiceProtocol(): Handling different protocols, REST/RPC/SOAP, for each service to communicate with others.

ServiceMap(): mapping provider's api to service APIs

ServiceAggregate(): Aggregating different provider's API to one Service API

ServiceHandler(): Defining handler for each service route.

ServiceDatabase(): Database connection for each service

ServiceFilter(): filtering includes in pass/block of data-type value-range value-format value-order(sort)
				 time-range  provider-selection priority-selection



*/

/*
service map and aggregation

Different services
	(different APIs)
		different routes
			different methods (get, post, put, delete)
				Request Header
				Request Body
				Response Header
					map[string][]string		//key:value
				Response Body
					API structure
						fields: Each fiels maps to several fields of different TPP API

*/

/*
API Proxy
convert the structure of on API to another structured API





*/
