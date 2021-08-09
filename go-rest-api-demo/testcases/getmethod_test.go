package testcases

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/subramanyam/go-rest-api/controller"
	"github.com/stretchr/testify/assert"
)

func TestEmpGet(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	controller.EmpGet(c)
	assert.Equal(t, 200, w.Code) // or what value you need it to be

	var got gin.H
	var want gin.H
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, want, got) // want is a gin.H that contains the wanted map.
}
