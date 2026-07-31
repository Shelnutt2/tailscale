// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build js

package magicsock

import (
	"tailscale.com/types/logger"
	"tailscale.com/types/nettype"
)

// tryPreventFragmentation is a no-op on js. WebAssembly has no access to
// socket options, and the browser networking stack owns fragmentation.
func tryPreventFragmentation(pconn nettype.PacketConn, logf logger.Logf, network string) {}
