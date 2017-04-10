/*
 * Copyright (c) 2014, Psiphon Inc.
 * All rights reserved.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package common

import (
	"bytes"
	"fmt"
	"math"
	"testing"
	"time"
)

func TestMakeRandomPeriod(t *testing.T) {
	min := 1 * time.Nanosecond
	max := 10000 * time.Nanosecond

	res1, err := MakeRandomPeriod(min, max)

	if err != nil {
		t.Error("MakeRandomPeriod failed: %s", err)
	}

	if res1 < min {
		t.Error("duration should not be less than min")
	}

	if res1 > max {
		t.Error("duration should not be more than max")
	}

	res2, err := MakeRandomPeriod(min, max)

	if err != nil {
		t.Error("MakeRandomPeriod failed: %s", err)
	}

	if res1 == res2 {
		t.Error("duration should have randomness difference between calls")
	}
}

func TestJitter(t *testing.T) {

	testCases := []struct {
		n           int64
		factor      float64
		expectedMin int64
		expectedMax int64
	}{
		{100, 0.1, 90, 110},
		{1000, 0.3, 700, 1300},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("jitter case: %+v", testCase), func(t *testing.T) {

			min := int64(math.MaxInt64)
			max := int64(0)

			for i := 0; i < 100000; i++ {

				x := Jitter(testCase.n, testCase.factor)
				if x < min {
					min = x
				}
				if x > max {
					max = x
				}
			}

			if min != testCase.expectedMin {
				t.Errorf("unexpected minimum jittered value: %d", min)
			}

			if max != testCase.expectedMax {
				t.Errorf("unexpected maximum jittered value: %d", max)
			}
		})
	}
}

func TestCompress(t *testing.T) {

	originalData := []byte("test data")

	compressedData := Compress(originalData)

	decompressedData, err := Decompress(compressedData)
	if err != nil {
		t.Error("Uncompress failed: %s", err)
	}

	if bytes.Compare(originalData, decompressedData) != 0 {
		t.Error("decompressed data doesn't match original data")
	}
}
