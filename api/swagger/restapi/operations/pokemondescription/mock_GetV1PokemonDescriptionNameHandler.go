// Code generated by mockery v2.3.0. DO NOT EDIT.

package pokemondescription

import (
	middleware "github.com/go-openapi/runtime/middleware"
	mock "github.com/stretchr/testify/mock"
)

// MockGetV1PokemonDescriptionNameHandler is an autogenerated mock type for the GetV1PokemonDescriptionNameHandler type
type MockGetV1PokemonDescriptionNameHandler struct {
	mock.Mock
}

// Handle provides a mock function with given fields: _a0
func (_m *MockGetV1PokemonDescriptionNameHandler) Handle(_a0 GetV1PokemonDescriptionNameParams) middleware.Responder {
	ret := _m.Called(_a0)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(GetV1PokemonDescriptionNameParams) middleware.Responder); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}
