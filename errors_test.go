// Copyright 2016 Sam Whited.
// Use of this source code is governed by the BSD 2-clause license that can be
// found in the LICENSE file.

package xmpp

import (
	"fmt"
)

var (
	_ error        = (*StanzaError)(nil)
	_ error        = StanzaError{}
	_ fmt.Stringer = (*errorType)(nil)
	_ fmt.Stringer = Auth
)
