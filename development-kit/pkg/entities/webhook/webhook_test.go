// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package webhook

import (
	"net/http"
	"testing"

	errorsEnum "github.com/ZupIT/horusec/development-kit/pkg/enums/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestWebhook_GetMethod(t *testing.T) {
	t.Run("Should return string by method valid", func(t *testing.T) {
		w := &Webhook{
			Method: "post",
		}
		assert.Equal(t, http.MethodPost, w.GetMethod())
		w = &Webhook{
			Method: "put",
		}
		assert.Equal(t, "", w.GetMethod())
		w = &Webhook{
			Method: "patch",
		}
		assert.Equal(t, "", w.GetMethod())
		w = &Webhook{
			Method: "get",
		}
		assert.Equal(t, "", w.GetMethod())
		w = &Webhook{
			Method: "other",
		}
		assert.Equal(t, "", w.GetMethod())
	})
}

func TestWebhook_GetTable(t *testing.T) {
	t.Run("should return table name", func(t *testing.T) {
		w := &Webhook{}
		assert.Equal(t, "webhooks", w.GetTable())
	})
}

func TestWebhook_GetHeaders(t *testing.T) {
	w := &Webhook{
		Headers: []Headers{
			{
				Key:   "X-Horusec-Authorization",
				Value: "Bearer token",
			},
		},
	}
	headerMap := w.GetHeaders()
	assert.NotEmpty(t, headerMap)
	assert.Equal(t, "Bearer token", headerMap["X-Horusec-Authorization"])
}

func TestWebhook_Validate(t *testing.T) {
	t.Run("Should return error when is url invalid", func(t *testing.T) {
		w := &Webhook{
			URL:          "invalid url",
			Method:       "POST",
			RepositoryID: uuid.New(),
			CompanyID:    uuid.New(),
		}
		err := w.Validate()
		assert.Equal(t, "url: must be a valid URL.", err.Error())
	})
	t.Run("Should return error when is method invalid", func(t *testing.T) {
		w := &Webhook{
			URL:          "http://example.com",
			Method:       "GET",
			RepositoryID: uuid.New(),
			CompanyID:    uuid.New(),
		}
		err := w.Validate()
		assert.Equal(t, "method: must be a valid value.", err.Error())
	})
}

func TestWebhook_SetCompanyIDAndRepositoryID(t *testing.T) {
	t.Run("Should return error when repositoryID is invalid to set in webhook", func(t *testing.T) {
		w := &Webhook{
			URL:    "http://example.com",
			Method: "GET",
		}
		newWebhook, err := w.SetCompanyIDAndRepositoryID(uuid.New().String(), "invalid")
		assert.Equal(t, errorsEnum.ErrorInvalidRepositoryID, err)
		assert.Nil(t, newWebhook)
	})
	t.Run("Should return error when companyID is invalid to set in webhook", func(t *testing.T) {
		w := &Webhook{
			URL:    "http://example.com",
			Method: "GET",
		}
		newWebhook, err := w.SetCompanyIDAndRepositoryID("invalid", uuid.New().String())
		assert.Equal(t, errorsEnum.ErrorInvalidCompanyID, err)
		assert.Nil(t, newWebhook)
	})
	t.Run("Should return error when companyID is invalid to set in webhook", func(t *testing.T) {
		w := &Webhook{
			URL:    "http://example.com",
			Method: "GET",
		}
		newWebhook, err := w.SetCompanyIDAndRepositoryID(uuid.New().String(), uuid.New().String())
		assert.NoError(t, err)
		assert.NotEmpty(t, newWebhook)
	})
}

func TestWebhook_SetWebhookID(t *testing.T) {
	w := &Webhook{
		URL:    "http://example.com",
		Method: "GET",
	}
	assert.Equal(t, uuid.Nil, w.WebhookID)
	w = w.SetWebhookID(uuid.New())
	assert.NotEqual(t, uuid.Nil, w.WebhookID)
}

func TestWebhook_ToBytes(t *testing.T) {
	w := &Webhook{
		URL:    "http://example.com",
		Method: "GET",
	}
	assert.NotEmpty(t, w.ToBytes())
}
