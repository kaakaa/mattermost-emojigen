package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mattermost/mattermost-server/v5/plugin/plugintest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServerAPI(t *testing.T) {
	assert := assert.New(t)
	api := &plugintest.API{}
	api.On("LogDebug", getMockArgumentsWithType("string", 7)...).Return()
	defer api.AssertExpectations(t)
	p := EmojigenPlugin{}
	p.SetAPI(api)
	p.router = p.InitAPI()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	p.ServeHTTP(nil, w, r)

	result := w.Result()
	assert.NotNil(result)
	defer result.Body.Close()
	bodyBytes, err := ioutil.ReadAll(result.Body)
	assert.Nil(err)
	bodyString := string(bodyBytes)
	assert.Equal("This is Mattermost Emojigen v"+manifest.Version+"\n", bodyString)
}

func getMockArgumentsWithType(typeString string, num int) []interface{} {
	ret := make([]interface{}, num)
	for i := 0; i < len(ret); i++ {
		ret[i] = mock.AnythingOfTypeArgument(typeString)
	}
	return ret
}
