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

	api.GetStudentHandler = operations.GetStudentHandlerFunc(
		func(params operations.GetStudentParams) middleware.Responder {
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
				return operations.NewGetStudentOK().WithPayload(greeting)
			} else {
				greeting := fmt.Sprintf("This is Get Call" )
				greeting += fmt.Sprintf("\nName, %s!", name)
				greeting += fmt.Sprintf("\nDepartment, %s!", Department)
				greeting += fmt.Sprintf("\nAge, %d!", age)
				greeting += fmt.Sprintf("\nEnrollment No , %d!", EnrollmentNo)
				return operations.NewGetStudentOK().WithPayload(greeting)
			}

		})

	api.PostStudentHandler = operations.PostStudentHandlerFunc(
		func(params operations.PostStudentParams) middleware.Responder {
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
				return operations.NewPostStudentOK().WithPayload(greeting)
			} else {
				greeting := fmt.Sprintf("This is Post Call" )
				greeting += fmt.Sprintf("\nName, %s!", name)
				greeting += fmt.Sprintf("\nDepartment, %s!", Department)
				greeting += fmt.Sprintf("\nAge, %d!", age)
				greeting += fmt.Sprintf("\nEnrollment No , %d!", EnrollmentNo)
				return operations.NewPostStudentOK().WithPayload(greeting)
			}

		})

	api.PutStudentHandler = operations.PutStudentHandlerFunc(
		func(params operations.PutStudentParams) middleware.Responder {
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
				return operations.NewPutStudentOK().WithPayload(greeting)
			} else {
				greeting := fmt.Sprintf("This is Put Call" )
				greeting += fmt.Sprintf("\nName, %s!", name)
				greeting += fmt.Sprintf("\nDepartment, %s!", Department)
				greeting += fmt.Sprintf("\nAge, %d!", age)
				greeting += fmt.Sprintf("\nEnrollment No , %d!", EnrollmentNo)
				return operations.NewPutStudentOK().WithPayload(greeting)
			}

		})

	api.DeleteStudentHandler= operations.DeleteStudentHandlerFunc(
		func(params operations.DeleteStudentParams)middleware.Responder{
			id := swag.StringValue(params.ID)
			greeting := fmt.Sprintf("This is Delete Call " )
			greeting += fmt.Sprintf("ID to delete, %s ",id)
			return operations.NewDeleteStudentOK().WithPayload(greeting)
		})
	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}