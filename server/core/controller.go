package core

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

// Controller is the core controller struct for all controllers
type Controller struct {
}

// ErrorResponse is the struct that holds an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// IController is the controller for models
type IController interface {
	LoadRoutes(r *gin.RouterGroup)
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
}

// LoadRoutes loads all the routes of the controller
func (controller *Controller) LoadRoutes(r *gin.RouterGroup) {
}

// Get is a method to get a model upon a request
func (controller *Controller) Get(c *gin.Context) {
	c.AbortWithStatus(http.StatusMethodNotAllowed)
}

// GetAll is a method to get all models upon a request
func (controller *Controller) GetAll(c *gin.Context) {
	c.AbortWithStatus(http.StatusMethodNotAllowed)
}

// Create is a method to create a model upon a request
func (controller *Controller) Create(c *gin.Context) {
	c.AbortWithStatus(http.StatusMethodNotAllowed)
}

// Update is a method to update a model upon a request
func (controller *Controller) Update(c *gin.Context) {
	c.AbortWithStatus(http.StatusMethodNotAllowed)
}

// HandleError is  a method to handle an error upon a request
func (controller *Controller) HandleError(c *gin.Context, object error) {
	switch object {
	// Gorm Errors
	case gorm.ErrRecordNotFound:
		c.AbortWithError(http.StatusNotFound, ErrorNotFound).
			SetType(gin.ErrorTypePublic)
		return
		// Custom Errors
	case ErrorBadRequest:
		c.AbortWithError(http.StatusBadRequest, ErrorBadRequest).
			SetType(gin.ErrorTypePublic)
		return
	case ErrorDuplicated:
		c.AbortWithError(http.StatusBadRequest, ErrorDuplicated).
			SetType(gin.ErrorTypePublic)
		return
	case ErrorInternalServerError:
		c.AbortWithError(http.StatusInternalServerError, ErrorInternalServerError).
			SetType(gin.ErrorTypePublic)
		return
	case ErrorNotFound:
		c.AbortWithError(http.StatusBadRequest, ErrorNotFound).
			SetType(gin.ErrorTypePublic)
		return
	case ErrorBadLogin:
		c.AbortWithError(http.StatusUnauthorized, ErrorBadLogin).
			SetType(gin.ErrorTypePublic)
		return
	case ErrorPasswordsDontMatch:
		c.AbortWithError(http.StatusBadRequest, ErrorPasswordsDontMatch).
			SetType(gin.ErrorTypePublic)
		return
	case ErrorInvalidGender:
		c.AbortWithError(http.StatusBadRequest, ErrorInvalidGender).
			SetType(gin.ErrorTypePublic)
		return
	default:
		c.Error(object).SetType(gin.ErrorTypePrivate)
		c.AbortWithError(http.StatusInternalServerError, ErrorInternalServerError).
			SetType(gin.ErrorTypePublic)
		return
	}
}

// GetRequestID returns the ID value of a request
func (controller *Controller) GetRequestID(c *gin.Context) (uint, error) {
	value, err := controller.GetRequestParam(c, "id")
	if err != nil {
		return 0, ErrorBadRequest
	}

	i64, err := strconv.Atoi(value)
	if err != nil {
		return 0, ErrorBadRequest
	}

	return uint(i64), nil
}

// GetRequestParam returns the ID value of a request
func (controller *Controller) GetRequestParam(c *gin.Context, param string) (string, error) {
	value := c.Param(param)
	if value == "" {
		return "", ErrorBadRequest
	}

	return value, nil
}

// writeError is a method to write the error response upon a request
func (controller *Controller) writeError(c *gin.Context, code int, object error) {
	errorObject := ErrorResponse{
		Error: object.Error(),
	}

	c.JSON(code, errorObject)
}
