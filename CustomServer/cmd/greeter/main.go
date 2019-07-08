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
			EnrollmentNo := int64(swag.Int64Value(params.Enrollment))
			if name == "" {
				name = "Nil"
			}

			if  Department == ""{
				Department = "CS"
			}

			if age < 20 {
				greeting := fmt.Sprintf("Age is less than, %d!", age)
				return operations.NewGetGreetingOK().WithPayload(greeting)
			} else {
				greeting := fmt.Sprintf("This is Get Call" )
				greeting += fmt.Sprintf("\nName, %s!", name)
				greeting += fmt.Sprintf("\nDepartment, %s!", Department)
				greeting += fmt.Sprintf("\nAge, %d!", age)
				greeting += fmt.Sprintf("\nEnrollment No , %d!", EnrollmentNo)
				return operations.NewGetGreetingOK().WithPayload(greeting)
			}

		})

	api.PostGreetingHandler = operations.PostGreetingHandlerFunc(
		func(params operations.PostGreetingParams) middleware.Responder {
			name := swag.StringValue(params.Name)
			Department := swag.StringValue(params.Department)
			age :=  swag.Int32Value(params.Age)
			EnrollmentNo := int64(swag.Int64Value(params.Enrollment))
			if name == "" {
				name = "Nil"
			}

			if  Department == ""{
				Department = "CS"
			}

			if age < 20 {
				greeting := fmt.Sprintf("Age is less than, %d!", age)
				return operations.NewPostGreetingOK().WithPayload(greeting)
			} else {
				greeting := fmt.Sprintf("This is Post Call" )
				greeting += fmt.Sprintf("\nName, %s!", name)
				greeting += fmt.Sprintf("\nDepartment, %s!", Department)
				greeting += fmt.Sprintf("\nAge, %d!", age)
				greeting += fmt.Sprintf("\nEnrollment No , %d!", EnrollmentNo)
				return operations.NewPostGreetingOK().WithPayload(greeting)
			}

		})

	api.PutGreetingHandler = operations.PutGreetingHandlerFunc(
		func(params operations.PutGreetingParams) middleware.Responder {
			name := swag.StringValue(params.Name)
			Department := swag.StringValue(params.Department)
			age :=  swag.Int32Value(params.Age)
			EnrollmentNo := int64(swag.Int64Value(params.Enrollment))
			if name == "" {
				name = "Nil"
			}

			if  Department == ""{
				Department = "CS"
			}

			if age < 20 {
				greeting := fmt.Sprintf("Age is less than, %d!", age)
				return operations.NewPutGreetingOK().WithPayload(greeting)
			} else {
				greeting := fmt.Sprintf("This is Put Call" )
				greeting += fmt.Sprintf("\nName, %s!", name)
				greeting += fmt.Sprintf("\nDepartment, %s!", Department)
				greeting += fmt.Sprintf("\nAge, %d!", age)
				greeting += fmt.Sprintf("\nEnrollment No , %d!", EnrollmentNo)
				return operations.NewPutGreetingOK().WithPayload(greeting)
			}

		})

	api.DeleteGreetingHandler= operations.DeleteGreetingHandlerFunc(
		func(params operations.DeleteGreetingParams)middleware.Responder{
			id := swag.StringValue(params.ID)
			greeting := fmt.Sprintf("This is Delete Call " )
			greeting += fmt.Sprintf("ID to delete, %s ",id)
			return operations.NewDeleteGreetingOK().WithPayload(greeting)
		})
	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}