// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !js

package mono

import _ "unsafe" // for go:linkname

//go:linkname now runtime.nanotime1
func now() int64
