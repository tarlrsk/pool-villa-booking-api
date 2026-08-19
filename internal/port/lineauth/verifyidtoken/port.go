package verifyidtoken

// VerifyResult is the raw result of a LINE ID-token verification call.
type VerifyResult struct {
	Sub     string
	Name    string
	Picture string
}

// Verifier is the port LINE ID-token verification depends on to call LINE's API.
type Verifier interface {
	Verify(idToken, channelID string) (*VerifyResult, error)
}
