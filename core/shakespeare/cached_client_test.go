package shakespeare

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"ozum.safaoglu/pokemon-api/cache"
)

func Test_cachedSPClient_Translate_ChecksCacheFirst(t *testing.T) {
	text := "pikachu"
	translation := &Translation{
		Contents: Content{
			Translated: "translated",
		},
	}
	marshalledTranslation, err := json.Marshal(translation)
	assert.Nil(t, err)
	ctx := context.Background()

	cache := &cache.MockCache{}

	spClient := &MockSPClient{}
	cachedPokeAPI := NewCachedSPClient(spClient, cache)
	textHash, err := cachedPokeAPI.(*cachedSPClient).generateHash(text)
	assert.Nil(t, err)
	cache.On("Get", ctx, translationPrefix+textHash).Return(string(marshalledTranslation), nil)

	translationResp, err := cachedPokeAPI.Translate(ctx, text)
	assert.Nil(t, err)
	spClient.AssertNotCalled(t, "Translate")
	assert.Equal(t, *translation, *translationResp)

	cache.AssertNotCalled(t, "Set")
}

func Test_cachedSPClient_Translate_CallsAPI_WithoutCache(t *testing.T) {
	text := "Pikachu is an electric pokemon"
	translation := &Translation{
		Contents: Content{
			Translated: "translated",
		},
	}

	ctx := context.Background()

	cache := &cache.MockCache{}
	spClient := &MockSPClient{}
	spClientCached := NewCachedSPClient(spClient, cache)
	textHash, err := spClientCached.(*cachedSPClient).generateHash(text)
	assert.Nil(t, err)

	cache.On("Get", ctx, translationPrefix+textHash).Return("", nil)

	marshalled, err := spClientCached.(*cachedSPClient).marshalForCache(translation)
	assert.Nil(t, err)
	cache.On("Set", ctx, translationPrefix+textHash, marshalled, mock.Anything).Return(nil).Times(1)

	spClient.On("Translate", ctx, text).Return(translation, nil)

	translationResp, err := spClientCached.Translate(ctx, text)
	assert.Nil(t, err)
	spClient.AssertNumberOfCalls(t, "Translate", 1)
	assert.Equal(t, *translation, *translationResp)

	cache.AssertExpectations(t)
}
