package transform

import (
	"fizz/internal/datastore"
	"fizz/models"
)

// TopRequest transforms internal model to its related api model
func TopRequest(topRequest *datastore.TopRequest) *models.MostUsedRequest {
	return &models.MostUsedRequest{
		Hits:  &topRequest.Hits,
		Int1:  &topRequest.Int1,
		Int2:  &topRequest.Int2,
		Limit: &topRequest.Limit,
		Str1:  &topRequest.Str1,
		Str2:  &topRequest.Str2,
	}
}
