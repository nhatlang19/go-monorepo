package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/nhatlang19/go-monorepo/services/user/service"

	"github.com/go-playground/validator/v10"
)

type authHeader struct {
	IDToken string `header:"Authorization"`
}
type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

func AuthUser(s service.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}

		// bind Authorization Header to h and check for validation errors
		if err := c.ShouldBindHeader(&h); err != nil {
			if errs, ok := err.(validator.ValidationErrors); ok {
				// we used this type in bind_data to extract desired fields from errs
				// you might consider extracting it
				var invalidArgs []invalidArgument

				for _, err := range errs {
					invalidArgs = append(invalidArgs, invalidArgument{
						err.Field(),
						err.Value().(string),
						err.Tag(),
						err.Param(),
					})
				}

				err := fmt.Errorf("Invalid request parameters. See invalidArgs")

				c.JSON(http.StatusBadRequest, gin.H{
					"error":       err.Error(),
					"invalidArgs": invalidArgs,
				})
				c.Abort()
				return
			}

			// otherwise error type is unknown
			err := fmt.Errorf("Error type is unknown")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		idTokenHeader := strings.Split(h.IDToken, "Bearer ")

		if len(idTokenHeader) < 2 {
			err := fmt.Errorf("Must provide Authorization header with format `Bearer {token}`")

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		// validate ID token here
		user, err := s.ValidateIDToken(idTokenHeader[1])

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		fmt.Println("user", user)
		c.Set("user", user)

		c.Next()
	}
}
