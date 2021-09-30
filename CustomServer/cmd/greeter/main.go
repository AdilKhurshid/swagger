package main

import (
	"CustomServer/gen/restapi"
	"CustomServer/gen/restapi/operations"
	"CustomServer/gen/restapi/operations/student"
	"flag"
	"fmt"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"log"
)

var portFlag = flag.Int("port", 3000, "Port to run this service on")


func main() {

	type StudentStruct struct {
		ID int64
		Name string
		Age int32
		Enrollment int64
		Department string
	}

	StudentMap := make(map[int64]StudentStruct)
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

	api.StudentGetStudentHandler = student.GetStudentHandlerFunc(
		func(params student.GetStudentParams) middleware.Responder {
			var StudentObj StudentStruct
			StudentObj = StudentMap[int64(params.ID)]
			StudentInfo := fmt.Sprintf("This is Get Call" )
			StudentInfo += fmt.Sprintf("\nID, %d!", StudentObj.ID)
			StudentInfo += fmt.Sprintf("\nName, %s!", StudentObj.Name)
			StudentInfo += fmt.Sprintf("\nDepartment, %s!", StudentObj.Department)
			StudentInfo += fmt.Sprintf("\nAge, %d!", StudentObj.Age)
			StudentInfo += fmt.Sprintf("\nEnrollment No , %d!", StudentObj.Enrollment)
			return student.NewGetStudentOK().WithPayload(StudentInfo)
		})

	api.StudentPostStudentHandler = student.PostStudentHandlerFunc(
		func(params student.PostStudentParams) middleware.Responder {
			var StudentObj StudentStruct
			StudentObj.ID = int64(params.Body.ID)
			StudentObj.Name = swag.StringValue(params.Body.Name)
			StudentObj.Department = swag.StringValue(params.Body.Department)
			StudentObj.Age =  swag.Int32Value(params.Body.Age)
			StudentObj.Enrollment = int64(swag.Int64Value(params.Body.Enrollment))
			StudentMap[StudentObj.ID]= StudentObj
			StudentInfo := fmt.Sprintf("This is Post Call" )
			StudentInfo += fmt.Sprintf("\nID, %d!", StudentObj.ID)
			StudentInfo += fmt.Sprintf("\nName, %s!", StudentObj.Name)
			StudentInfo += fmt.Sprintf("\nDepartment, %s!", StudentObj.Department)
			StudentInfo += fmt.Sprintf("\nAge, %d!", StudentObj.Age)
			StudentInfo += fmt.Sprintf("\nEnrollment No , %d!", StudentObj.Enrollment)
			return student.NewPostStudentOK().WithPayload(StudentInfo)
		})

	api.StudentPutStudentHandler= student.PutStudentHandlerFunc(
		func(params student.PutStudentParams) middleware.Responder {
			StudentObj := StudentMap[int64(params.Body.ID)]
			name := swag.StringValue(params.Body.Name)
			Department := swag.StringValue(params.Body.Department)
			age :=  swag.Int32Value(params.Body.Age)
			EnrollmentNo := int64(swag.Int64Value(params.Body.Enrollment))
			StudentObj.Name = name
			StudentObj.Department = Department
			StudentObj.Age = age
			StudentObj.Enrollment = EnrollmentNo
			StudentMap[int64(params.Body.ID)] = StudentObj
			StudentInfo := fmt.Sprintf("This is Put Call" )
			StudentInfo += fmt.Sprintf("\nName, %s!", name)
			StudentInfo += fmt.Sprintf("\nDepartment, %s!", Department)
			StudentInfo += fmt.Sprintf("\nAge, %d!", age)
			StudentInfo += fmt.Sprintf("\nEnrollment No , %d!", EnrollmentNo)
			return student.NewPutStudentOK().WithPayload(StudentInfo)

		})

	api.StudentDeleteStudentHandler= student.DeleteStudentHandlerFunc(
		func(params student.DeleteStudentParams)middleware.Responder{
			ID := int64(params.ID)
			StudentInfo := fmt.Sprintf("This is Delete Call " )
			StudentInfo += fmt.Sprintf("ID to delete, %d! ",ID)
			delete(StudentMap,ID)
			return student.NewDeleteStudentNoContent().WithPayload(StudentInfo)
		})

	api.StudentGetStudentListHandler = student.GetStudentListHandlerFunc(
		func(params student.GetStudentListParams) middleware.Responder {
			StudentInfo := ""
			for key , StudentObj := range StudentMap{
				StudentInfo += fmt.Sprintf("\n**This is Get Call**")
				StudentInfo += fmt.Sprintf("\nID, %d!", key)
				StudentInfo += fmt.Sprintf("\nName, %s!", StudentObj.Name)
				StudentInfo += fmt.Sprintf("\nDepartment, %s!", StudentObj.Department)
				StudentInfo += fmt.Sprintf("\nAge, %d!", StudentObj.Age)
				StudentInfo += fmt.Sprintf("\nEnrollment No , %d!", StudentObj.Enrollment)
			}
			return student.NewGetStudentListOK().WithPayload(StudentInfo)
		})
	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}