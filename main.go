// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package appstats

import (
	"expvar"
	"sync"
)

var (
	stats  = expvar.NewMap("stats")
	maps   = make(map[string]*expvar.Map)
	mapsmu sync.Mutex
)

// NewString creates a new string stat
func NewString(name string) error {
	if stats.Get(name) != nil {
		return nil
	}

	stats.Set(name, new(expvar.String))

	return nil
}

// SetString sets string stat to value
func SetString(name string, value string) error {
	if stats.Get(name) == nil {
		if err := NewString(name); err != nil {
			return err
		}
	}

	stats.Get(name).(*expvar.String).Set(value)

	return nil
}

// NewInt creates a new int stat
func NewInt(name string) error {
	if stats.Get(name) != nil {
		return nil
	}

	stats.Set(name, new(expvar.Int))

	return nil
}

// SetInt sets int stat to value
func SetInt(name string, value int64) error {
	if stats.Get(name) == nil {
		if err := NewInt(name); err != nil {
			return err
		}
	}

	stats.Get(name).(*expvar.Int).Set(value)

	return nil
}

// IncrementInt increment int stat
func IncrementInt(name string) error {
	if stats.Get(name) == nil {
		if err := NewInt(name); err != nil {
			return err
		}
	}

	stats.Add(name, 1)

	return nil
}

// DecrementInt decrement int stat
func DecrementInt(name string) error {
	if stats.Get(name) == nil {
		if err := NewInt(name); err != nil {
			return err
		}
	}

	stats.Add(name, -1)

	return nil
}

// AddInt adds delta to int
func AddInt(name string, delta int64) error {
	if stats.Get(name) == nil {
		if err := NewInt(name); err != nil {
			return err
		}
	}

	stats.Get(name).(*expvar.Int).Add(delta)

	return nil
}

// NewFloat creates a new float stat
func NewFloat(name string) error {
	if stats.Get(name) != nil {
		return nil
	}

	stats.Set(name, new(expvar.Float))

	return nil
}

// SetFloat sets a float stat to value
func SetFloat(name string, value float64) error {
	if stats.Get(name) == nil {
		if err := NewFloat(name); err != nil {
			return err
		}
	}

	stats.Get(name).(*expvar.Float).Set(value)

	return nil
}

// AddFloat adds value to existing float
func AddFloat(name string, value float64) error {
	if stats.Get(name) == nil {
		if err := NewFloat(name); err != nil {
			return err
		}
	}

	stats.Get(name).(*expvar.Float).Add(value)

	return nil
}

// NewMap creates a new map
func NewMap(name string) error {
	mapsmu.Lock()
	defer mapsmu.Unlock()
	if _, ok := maps[name]; ok {
		return nil
	}

	maps[name] = expvar.NewMap(name)

	return nil
}

// MapAddInt adds delta to int in a map
func MapAddInt(mapName, statName string, delta int64) error {
	mapsmu.Lock()
	defer mapsmu.Unlock()
	if _, ok := maps[mapName]; !ok {
		maps[mapName] = expvar.NewMap(mapName)
	}

	maps[mapName].Add(statName, delta)

	return nil
}

// MapAddFloat adds delta to a float in a map
func MapAddFloat(mapName, statName string, delta float64) error {
	mapsmu.Lock()
	defer mapsmu.Unlock()
	if _, ok := maps[mapName]; !ok {
		maps[mapName] = expvar.NewMap(mapName)
	}

	maps[mapName].AddFloat(statName, delta)

	return nil
}
