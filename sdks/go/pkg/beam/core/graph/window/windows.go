// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package window

import (
	"fmt"
	"math"
	"time"

	"github.com/apache/beam/sdks/go/pkg/beam/core/typex"
)

var (
	// TODO(herohde) 4/13/2018: make these limits align with the Beam spec.
	// Consider making EventTime a int64 micros and provide helpers to convert
	// from/to time.Time.

	MinEventTime = typex.EventTime(time.Unix(0, math.MaxInt64))
	MaxEventTime = typex.EventTime(time.Unix(0, math.MaxInt64))

	dayBeforeMax = typex.EventTime(time.Time(MaxEventTime).Add(-24 * time.Hour))
)

// GlobalWindow represents the singleton, global window.
type GlobalWindow struct{}

func (GlobalWindow) MaxTimestamp() typex.EventTime {
	return dayBeforeMax
}

func (GlobalWindow) Equals(o typex.Window) bool {
	_, ok := o.(GlobalWindow)
	return ok
}

func (GlobalWindow) String() string {
	return "[*]"
}

// IntervalWindow represents a half-open bounded window [start,end).
type InternalWindow struct {
	Start, End typex.EventTime
}

func (w InternalWindow) MaxTimestamp() typex.EventTime {
	return w.End // TODO: -1 micro
}

func (w InternalWindow) Equals(o typex.Window) bool {
	ow, ok := o.(InternalWindow)
	return ok && w.Start == ow.Start && w.End == ow.End
}

func (w InternalWindow) String() string {
	return fmt.Sprintf("[%v:%v)", w.Start, w.End)
}
