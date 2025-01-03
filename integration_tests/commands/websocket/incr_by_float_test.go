// This file is part of DiceDB.
// Copyright (C) 2024 DiceDB (dicedb.io).
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package websocket

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestINCRBYFLOAT(t *testing.T) {
	exec := NewWebsocketCommandExecutor()
	conn := exec.ConnectToServer()
	invalidArgMessage := "ERR wrong number of arguments for 'incrbyfloat' command"
	invalidIncrTypeMessage := "ERR value is not a valid float"
	valueOutOfRangeMessage := "ERR value is out of range"

	defer func() {
		exec.FireCommandAndReadResponse(conn, "DEL foo")
	}()

	testCases := []struct {
		name   string
		cmds   []string
		expect []interface{}
		delays []time.Duration
	}{
		{
			name:   "Invalid number of arguments",
			cmds:   []string{"INCRBYFLOAT", "INCRBYFLOAT foo"},
			expect: []interface{}{invalidArgMessage, invalidArgMessage},
			delays: []time.Duration{0, 0},
		},
		{
			name:   "Increment a non existing key",
			cmds:   []string{"INCRBYFLOAT foo 0.1", "GET foo"},
			expect: []interface{}{"0.1", "0.1"},
			delays: []time.Duration{0, 0},
		},
		{
			name:   "Increment a key with an integer value",
			cmds:   []string{"SET foo 1", "INCRBYFLOAT foo 0.1", "GET foo"},
			expect: []interface{}{"OK", "1.1", "1.1"},
			delays: []time.Duration{0, 0, 0},
		},
		{
			name:   "Increment and then decrement a key with the same value",
			cmds:   []string{"SET foo 1", "INCRBYFLOAT foo 0.1", "GET foo", "INCRBYFLOAT foo -0.1", "GET foo"},
			expect: []interface{}{"OK", "1.1", "1.1", "1", "1"},
			delays: []time.Duration{0, 0, 0, 0, 0},
		},
		{
			name:   "Increment a non numeric value",
			cmds:   []string{"SET foo bar", "INCRBYFLOAT foo 0.1"},
			expect: []interface{}{"OK", invalidIncrTypeMessage},
			delays: []time.Duration{0, 0},
		},
		{
			name:   "Increment by a non numeric value",
			cmds:   []string{"SET foo 1", "INCRBYFLOAT foo bar"},
			expect: []interface{}{"OK", invalidIncrTypeMessage},
			delays: []time.Duration{0, 0},
		},
		{
			name:   "Increment by both integer and float",
			cmds:   []string{"SET foo 1", "INCRBYFLOAT foo 1", "INCRBYFLOAT foo 0.1"},
			expect: []interface{}{"OK", "2", "2.1"},
			delays: []time.Duration{0, 0, 0},
		},
		{
			name:   "Increment that would make the value Inf",
			cmds:   []string{"SET foo 1e308", "INCRBYFLOAT foo 1e308", "INCRBYFLOAT foo -1e308"},
			expect: []interface{}{"OK", valueOutOfRangeMessage, "0"},
			delays: []time.Duration{0, 0, 0},
		},
		{
			name:   "Increment that would make the value -Inf",
			cmds:   []string{"SET foo -1e308", "INCRBYFLOAT foo -1e308", "INCRBYFLOAT foo 1e308"},
			expect: []interface{}{"OK", valueOutOfRangeMessage, "0"},
			delays: []time.Duration{0, 0, 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer exec.FireCommandAndReadResponse(conn, "DEL key unsetKey stringkey")
			for i, cmd := range tc.cmds {
				if tc.delays[i] > 0 {
					time.Sleep(tc.delays[i])
				}
				result, err := exec.FireCommandAndReadResponse(conn, cmd)
				assert.Nil(t, err)
				assert.Equal(t, tc.expect[i], result, "Value mismatch for cmd %s", cmd)
			}
		})
	}
}
