// Package response {{replace_ProjectName}}
//
// ## {{replace_TitleProjectSmallDescription}}
//
// {{replace_ProjectDescription}}
//
// Version: {{replace_ProjectVersion}}
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// BasePath: /{{replace_BasePath}}
//
// SecurityDefinitions:
// bearer:
//  type: apiKey
//  name: Authorization
//  in: header
//  example: Bearer
//
// swagger:meta
package response

// swagger:response success
type swaggerSuccessResponse struct {
	// in: body
	Data struct {
		// example: success
		Message string `json:"message"`
		// example: null
		Payload interface{} `json:"payload"`
	}
}

// swagger:response conflict
type swaggerConflictResponse struct {
	// in: body
	Data struct {
		// example: conflict
		Message string `json:"message"`
		// example: null
		Payload interface{} `json:"payload"`
	}
}

// swagger:response badRequest
type swaggerBadRequestResponse struct {
	// in: body
	Data struct {
		// example: bad request
		Message string `json:"message"`
		// example: null
		Payload interface{} `json:"payload"`
	}
}

// swagger:response accessDenied
type swaggerAccessDeniedResponse struct {
	// in: body
	Data struct {
		// example: access is denied
		Message string `json:"message"`
		// example: null
		Payload interface{} `json:"payload"`
	}
}

// swagger:response notFound
type swaggerNotFoundResponse struct {
	// in: body
	Data struct {
		// example: notFound
		Message string `json:"message"`
		// example: null
		Payload interface{} `json:"payload"`
	}
}

// swagger:response retryLimitExceeded
type swaggerRetryLimitExceededResponse struct {
	// in: body
	Data struct {
		// example: retry limit exceeded
		Message string `json:"message"`
		// example: null
		Payload interface{} `json:"payload"`
	}
}

// swagger:response notImplementation
type swaggerNotImplementationResponse struct {
	// in: body
	Data struct {
		// example: not implementation
		Message string `json:"message"`
		// example: null
		Payload interface{} `json:"payload"`
	}
}

// swagger:response unauthorized
type swaggerUnauthorizedResponse struct {
	// in: body
	Data struct {
		// example: unauthorized
		Message string `json:"message"`
		// example: null
		Payload interface{} `json:"payload"`
	}
}

// swagger:response somethingWentWrong
type swaggerSomethingWentWrongResponse struct {
	// in: body
	Data struct {
		// example: something went wrong
		Message string `json:"message"`
		// example: null
		Payload interface{} `json:"payload"`
	}
}

// swagger:response internalServer
type swaggerInternalServerResponse struct {
	// in: body
	Data struct {
		// example: internal server error
		Message string `json:"message"`
		// example: null
		Payload interface{} `json:"payload"`
	}
}
