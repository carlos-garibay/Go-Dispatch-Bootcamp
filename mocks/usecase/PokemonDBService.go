// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	model "github.com/carlos-garibay/Go-Dispatch-Bootcamp/model"
	mock "github.com/stretchr/testify/mock"
)

// PokemonDBService is an autogenerated mock type for the PokemonDBService type
type PokemonDBService struct {
	mock.Mock
}

// CatchPokemon provides a mock function with given fields: p
func (_m *PokemonDBService) CatchPokemon(p *model.PokemonAPI) (*model.Pokemon, error) {
	ret := _m.Called(p)

	var r0 *model.Pokemon
	if rf, ok := ret.Get(0).(func(*model.PokemonAPI) *model.Pokemon); ok {
		r0 = rf(p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Pokemon)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.PokemonAPI) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePokemon provides a mock function with given fields: p
func (_m *PokemonDBService) CreatePokemon(p *model.Pokemon) error {
	ret := _m.Called(p)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Pokemon) error); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilterPokemons provides a mock function with given fields: typ, items, itemPerWorker
func (_m *PokemonDBService) FilterPokemons(typ string, items int, itemPerWorker int) (model.Pokemons, error) {
	ret := _m.Called(typ, items, itemPerWorker)

	var r0 model.Pokemons
	if rf, ok := ret.Get(0).(func(string, int, int) model.Pokemons); ok {
		r0 = rf(typ, items, itemPerWorker)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Pokemons)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(typ, items, itemPerWorker)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllPokemon provides a mock function with given fields:
func (_m *PokemonDBService) GetAllPokemon() (model.Pokemons, error) {
	ret := _m.Called()

	var r0 model.Pokemons
	if rf, ok := ret.Get(0).(func() model.Pokemons); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Pokemons)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPokemonById provides a mock function with given fields: id
func (_m *PokemonDBService) GetPokemonById(id int) (*model.Pokemon, error) {
	ret := _m.Called(id)

	var r0 *model.Pokemon
	if rf, ok := ret.Get(0).(func(int) *model.Pokemon); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Pokemon)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
