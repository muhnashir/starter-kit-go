package base

import "starter-kit-go/app/response"

type ServiceInterface interface {
	Do() response.ServiceResponse
}
