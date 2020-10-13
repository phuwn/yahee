package google

import (
	"github.com/phuwn/tools/errors"
	"github.com/phuwn/tools/util"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/people/v1"

	"github.com/phuwn/yahee/src/model"
)

// Service - Google service implementation
type Service interface {
	GetOauth2Token(code, redirectURL string) (*oauth2.Token, error)
	GetUserGoogleInfo(token *oauth2.Token) (*model.User, error)
}

type googleService struct {
	config *oauth2.Config
}

// NewService - create new google service
func NewService() Service {
	return &googleService{
		config: &oauth2.Config{
			ClientID:     util.Getenv("GOOGLE_CLIENT_ID", ""),
			ClientSecret: util.Getenv("GOOGLE_CLIENT_SECRET", ""),
			Endpoint:     google.Endpoint,
			Scopes:       []string{"profile", "email"},
		},
	}
}

// GetAccessToken - exchange user's code for access_token
func (g *googleService) GetOauth2Token(code, redirectURL string) (*oauth2.Token, error) {
	g.config.RedirectURL = redirectURL
	token, err := g.config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, errors.Customize(err, 400, "invalid auth code")
	}
	return token, nil
}

// GetUserGoogleInfo - get user's google info
func (g *googleService) GetUserGoogleInfo(token *oauth2.Token) (*model.User, error) {
	srv, err := people.New(g.config.Client(oauth2.NoContext, token))
	if err != nil {
		return nil, errors.Customize(err, 500, "failed to create new google people service")
	}

	r, err := srv.People.Get("people/me").
		PersonFields("names,emailAddresses,photos").Do()
	if err != nil {
		return nil, errors.Customize(err, 500, "failed to get user's google info")
	}

	var name, email, avatar string
	if len(r.Names) == 0 {
		return nil, errors.New("missing userinfo.profile scope", 400)
	}
	name = r.Names[0].DisplayName

	if len(r.EmailAddresses) == 0 {
		return nil, errors.New("missing userinfo.email scope", 400)
	}
	email = r.EmailAddresses[0].Value

	if len(r.Photos) == 0 {
		return nil, errors.New("missing userinfo.profile scope", 400)
	}
	avatar = r.Photos[0].Url

	return &model.User{Name: name, Email: email, Avatar: avatar}, nil
}
