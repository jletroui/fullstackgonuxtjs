import SuperTokens, { SuperTokensWrapper } from "supertokens-auth-react";
import EmailPassword from "supertokens-auth-react/recipe/emailpassword";
import Session from "supertokens-auth-react/recipe/session";
import { canHandleRoute, getRoutingComponent } from "supertokens-auth-react/ui";
import { EmailPasswordPreBuiltUI } from 'supertokens-auth-react/recipe/emailpassword/prebuiltui';
import TaskCount from './components/tasks/taskcount'
import TaskCount2 from './components/tasks/taskcount2'

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

export function App() {
  if (canHandleRoute([EmailPasswordPreBuiltUI])) {
    // This renders the login UI on the /auth route
    return getRoutingComponent([EmailPasswordPreBuiltUI])
  }

  return (
    <SuperTokensWrapper>
      <main>
        <h1>Hello, Preact!</h1>
        <h2>Task count: <TaskCount/></h2>
        <h2>Task count 2.1: <TaskCount2/></h2>
        <h2>Task count 2.2: <TaskCount2/></h2>
      </main>
    </SuperTokensWrapper>
  )
}
