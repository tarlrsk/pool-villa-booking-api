package verifyidtoken

type Profile struct {
	UserID      string `json:"userId"`
	DisplayName string `json:"displayName"`
	PictureURL  string `json:"pictureUrl,omitempty"`
}

// Service verifies a LINE ID token and returns the associated profile.
type Service interface {
	Execute(idToken string) (*Profile, error)
}
