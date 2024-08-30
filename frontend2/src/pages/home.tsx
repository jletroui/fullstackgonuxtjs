import { SessionAuth } from "supertokens-auth-react/recipe/session";

export default function HomePage() {
    return (
        <SessionAuth>
            <main>
                <h1>Welcome Home</h1>
            </main>
        </SessionAuth>
    )
  }
  