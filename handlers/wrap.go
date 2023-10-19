package handlers

import (
	"apiwrap/global"
	"apiwrap/misskey"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Wrap(ctx *gin.Context) {
	// Check body
	body := make(map[string]any)

	// Check params
	var err error
	if ctx.Request.Method != http.MethodGet {
		// Auto bind body
		err = ctx.Bind(&body)
	} else {
		// Bind query
		q := ctx.Request.URL.Query()
		for k, v := range q {
			if len(v) > 0 {
				if len(v) == 1 {
					body[k] = v[0]
				} else {
					body[k] = v
				}
			}
		}
	}

	if err != nil {
		global.Logger.Debugf("Failed to parse request body with error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"message": err.Error(),
			},
		})
		return
	}

	// Check Authorization header
	auth := ctx.GetHeader("Authorization")
	if auth != "" {
		global.Logger.Debugf("Authorization header found")
		// Add authorization API key
		authKey := strings.Split(auth, " ")
		if len(authKey) == 1 {
			// Set as auth key
			global.Logger.Debugf("Set Authorization header as raw key")
			body["i"] = auth
		} else if len(authKey) > 2 {
			global.Logger.Debugf("Invalid Authorization header")
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": gin.H{
					"message": "invalid Authorization header",
				},
			})
			return
		} else {
			switch authKey[0] {
			case "Bearer":
				global.Logger.Debugf("Set Authorization header as Bearer key")
				body["i"] = authKey[1]
			default:
				global.Logger.Debugf("Unsupported Authorization scheme")
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"error": gin.H{
						"message": "unsupported Authorization scheme",
					},
				})
				return
			}
		}
	}

	// Send request to Misskey
	res, code, err := misskey.PostAPIRequest(ctx.Request.RequestURI, body)
	if err != nil {
		global.Logger.Debugf("API failure: %v", err)
		ctx.JSON(code, gin.H{
			"error": gin.H{
				"message": err.Error(),
			},
		})
		return
	}

	ctx.JSON(code, res)
}
