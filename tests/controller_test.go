package tests

import (
	"net/http"
	"testing"

	"github.com/leonardo-otero390/race-war/controllers"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	var req = Request{http.MethodGet, "/_health_check", nil}

	res := MockServer(&req, controllers.HealthCheck)
	assert.Equal(t, http.StatusOK, res.Code)
}
