// Copyright 2018 Typeform SL. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package appengine

import (
	"context"
	"net/http"

	"google.golang.org/appengine"
)

// CtxMiddleware is a middleware for gokit that creates a new context
// that can be used with other applications
func CtxMiddleware(ctx context.Context, r *http.Request) context.Context {
	return appengine.WithContext(ctx, r)
}
