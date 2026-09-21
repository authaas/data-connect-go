// Package data is a data service's behavior on connect: how a statement
// that did not run to completion is reported.
package data

import (
	"context"

	"git.sonicoriginal.software/logger"
	errors "github.com/pbrpc/connect-errors"
)

// ReasonStoreFailed is the error reason a failed statement carries.
const ReasonStoreFailed = "STORE_FAILED"

// StoreFailed logs the failure and reports it as internal under domain, with
// action naming what was attempted, as "read the credential".
func StoreFailed(ctx context.Context, domain, action string, err error) error {
	logger.FromContext(ctx).ErrorContext(ctx, "Failed to "+action, "error", err)

	return errors.Internal(ctx, "failed to "+action, ReasonStoreFailed, domain)
}
