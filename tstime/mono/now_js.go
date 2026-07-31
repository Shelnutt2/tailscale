// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build js

package mono

import "time"

// The js linker rejects go:linkname references to runtime.nanotime1, so derive
// a monotonic clock from time.Since instead. time.Since already uses the
// monotonic reading embedded in a time.Time, so this stays monotonic.

var start = time.Now()

func now() int64 { return int64(time.Since(start)) }
