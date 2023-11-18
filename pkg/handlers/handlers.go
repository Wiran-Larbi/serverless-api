package handlers

import (
	"net/http"

	"github.com/Wiran-Larbi/serverless-api/pkg/user"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var ErrorMethodNotAllowed = "Method not Allowed !"

type ErrorBody struct {
	ErrorMsg *string `json:"error, omit empty"`
}

func GetUser(req events.APIGatewayProxyRequest, tableName string, dynamoClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {

	email := req.QueryStringParameters["email"]
	if len(email) > 0 {
		result, err := user.FetchUser(email, tableName, dynamoClient)
		if err != nil {
			return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}

		return apiResponse(http.StatusOK, result)
	}
	result, err := user.FetchUsers(tableName, dynamoClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusOK, result)
}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynamoClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	response, error := user.CreateUser(req, tableName, dynamoClient)
	if error != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(error.Error()),
		})
	}

	return apiResponse(http.StatusCreated, response)
}

func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynamoClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	response, error := user.UpdateUser(req, tableName, dynamoClient)
	if error != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(error.Error()),
		})
	}
	return apiResponse(http.StatusOK, response)
}

func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynamoClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	error := user.DeleteUser(req, tableName, dynamoClient)
	if error != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(error.Error()),
		})
	}
	return apiResponse(http.StatusOK, nil)
}

func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
