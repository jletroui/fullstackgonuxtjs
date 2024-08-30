import SuperTokens, { SuperTokensWrapper } from "supertokens-auth-react";
import EmailPassword from "supertokens-auth-react/recipe/emailpassword";
import Session from "supertokens-auth-react/recipe/session";
import { canHandleRoute, getRoutingComponent } from "supertokens-auth-react/ui";
import { EmailPasswordPreBuiltUI } from 'supertokens-auth-react/recipe/emailpassword/prebuiltui';
import { lazy, LocationProvider, Router, Route } from 'preact-iso';

SuperTokens.init({
    appInfo: {
        // learn more about this on https://supertokens.com/docs/emailpassword/appinfo
        appName: "fullstackgovitejs",
        apiDomain: "http://localhost:8080",
        websiteDomain: "http://localhost:5173",
        apiBasePath: "/auth",
        websiteBasePath: "/auth",
    },
    recipeList: [
        EmailPassword.init(),
        Session.init()
    ]
});

const LandingPage = lazy(() => import('./pages/landing'));
const HomePage = lazy(() => import('./pages/home'));
const NotFoundPage = lazy(() => import('./pages/404'));

export function App() {
  if (canHandleRoute([EmailPasswordPreBuiltUI])) {
    // This renders the login UI on the /auth route
    return getRoutingComponent([EmailPasswordPreBuiltUI])
  }

  return (
    <LocationProvider>
      <SuperTokensWrapper>
        <Router>
          <Route path="/" component={LandingPage} />
          <Route path="/home" component={HomePage } />
          <Route default component={NotFoundPage}/>
        </Router>
      </SuperTokensWrapper>
    </LocationProvider>
  )
}
