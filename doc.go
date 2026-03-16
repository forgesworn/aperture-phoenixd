// Package phoenixd provides a Phoenixd-backed challenger for Aperture's
// L402 authentication system. It implements Aperture's mint.Challenger
// and auth.InvoiceChecker interfaces against Phoenixd's REST API,
// removing the requirement for a full LND node.
//
// This package has no dependency on Aperture or LND. It uses plain Go
// types ([32]byte for payment hashes) that are assignment-compatible
// with Aperture's lntypes.Hash.
package phoenixd
