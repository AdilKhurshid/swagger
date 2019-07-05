package main

import (
	"CustomServer/gen/restapi"
	"CustomServer/gen/restapi/operations"
	"flag"
	"fmt"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"log"
)

var portFlag = flag.Int("port", 3000, "Port to run this service on")


func main() {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewGreeterAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	// parse flags
	flag.Parse()
	// set the port this service will be run on
	server.Port = *portFlag

	api.GetGreetingHandler = operations.GetGreetingHandlerFunc(
		func(params operations.GetGreetingParams) middleware.Responder {
			name := swag.StringValue(params.Name)
			Department := swag.StringValue(params.Department)
			age :=  swag.Int32Value(params.Age)
			EnrollmentNo := swag.Int64Value(params.Enrollment)
			if name == "" {
				name = "Nil"
			}

			if  Department == ""{
				Department = "CS"
			}

			if age < 20 {
				greeting := fmt.Sprintf("Age is less than, %s!", age)
				return operations.NewGetGreetingOK().WithPayload(greeting)
			} else {
				greeting := fmt.Sprintf("Hello, %s!", name)
				greeting += fmt.Sprintf("\nDepartment, %s!", Department)
				greeting += fmt.Sprintf("\nAge, %s!", age)
				greeting += fmt.Sprintf("\nEnrollment No , %s!", EnrollmentNo)
				return operations.NewGetGreetingOK().WithPayload(greeting)
			}

		})

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}