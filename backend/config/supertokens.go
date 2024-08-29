package config

import (
	"github.com/supertokens/supertokens-golang/recipe/dashboard"
	"github.com/supertokens/supertokens-golang/recipe/dashboard/dashboardmodels"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func ConfigureSuperTokens(cfg *Config) error {
	// https://supertokens.com/docs/emailpassword/pre-built-ui/setup/backend#2-initialise-supertokens

	apiBasePath := "/api/auth"
	websiteBasePath := "/auth"
	return supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: cfg.SuperTokensUrl,
			// APIKey: <API_KEY(if configured)>,
		},
		AppInfo: supertokens.AppInfo{
			AppName:         "fullstackgovitejs",
			APIDomain:       cfg.APIBasePath(),
			WebsiteDomain:   cfg.UiUrl,
			APIBasePath:     &apiBasePath,
			WebsiteBasePath: &websiteBasePath,
		},
		RecipeList: []supertokens.Recipe{
			emailpassword.Init(nil),
			session.Init(nil),
			dashboard.Init(&dashboardmodels.TypeInput{
				Admins: &cfg.SuperTokensAdmins,
			}),
		},
	})
}
