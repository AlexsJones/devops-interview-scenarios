package main

import (
	"github.com/AlexsJones/blacklist-example/gen/models"
	"github.com/AlexsJones/blacklist-example/gen/restapi"
	"github.com/AlexsJones/blacklist-example/gen/restapi/operations"
	"github.com/AlexsJones/blacklist-example/gen/restapi/operations/blacklist"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	_ "github.com/go-openapi/runtime/middleware"
	"github.com/jessevdk/go-flags"
	"log"
	"os"
)


func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewExampleAPI(swaggerSpec)

	//Example implementation of a blacklist in memory .....
	api.BlacklistGetBlackListHandler = blacklist.GetBlackListHandlerFunc(func (params blacklist.GetBlackListParams) middleware.Responder {

		return blacklist.NewGetBlackListOK().WithPayload(&models.Blacklist{
			Domains: []string {
				"piratebay.org",
				"piratebay.io",
				"piratebay.co.uk",
			},
		})
	})
	// ----------------------------------------------------

	server := restapi.NewServer(api)

	//Set a default port
	server.Port = 8080
	server.Host = "0.0.0.0"
	// ----------------------------------------------------

	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Example API"
	parser.LongDescription = "Example API"

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
