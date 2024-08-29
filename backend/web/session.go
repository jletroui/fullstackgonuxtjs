package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
)

var sessionRequired = true

type SessionVerifier interface {
	VerifySession(c *gin.Context)
	GetUserID(c *gin.Context) string
}

type sessionVerifier struct{}

func NewSuperTokenSessionVerifier() SessionVerifier {
	return &sessionVerifier{}
}

func (*sessionVerifier) VerifySession(c *gin.Context) {
	session.VerifySession(&sessmodels.VerifySessionOptions{SessionRequired: &sessionRequired}, func(rw http.ResponseWriter, r *http.Request) {
		c.Request = c.Request.WithContext(r.Context())
		c.Next()
	})(c.Writer, c.Request)
	// we call Abort so that the next handler in the chain is not called, unless we call Next explicitly
	c.Abort()
}

func (*sessionVerifier) GetUserID(c *gin.Context) string {
	sessionContainer := session.GetSessionFromRequestContext(c.Request.Context())
	return sessionContainer.GetUserID()
}
