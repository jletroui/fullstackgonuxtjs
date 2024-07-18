## Overview

Application template supporting the common dev and operations for the combination Kotlin (backend) and NextJS (frontend).
The goal was to explore how would look a modern full stack app with speed of maintainability over long period of time as its main driver. 

Supports:
- cross-platform dev operations command line (bash)
- dependencies management (go modules and npm)
- dev database, database migrations (Postgres and go migrate)
- performant web server (go Gin)
- configuration and secrets
- backend unit tests for the example data access layer, controller, and page
- backend linting

TODO:
- frontend linting
- frontend unit tests for the example data access layer, controller, and page
- structured logging
- deployment (production simulated locally)
- frontend and backend CI

Not covered:
- monitoring (other than logging)
- authentication
- SSL termination

## Reasoning behind the tech choices

Kotlin vs Go:

Both languages are performant, have a strong static compiler, are mature, have a massive community.

What we forgo when choosing one or the other:

Cons of Kotlin:
- less simple
- smaller pool of programmers with experience on it (although very close to Java)
- more integration work for auth tools like Ory Kratos or SuperTokens

Cons of Go:
- less expressive
- less constructs available to avoid technical bugs (immutability, nullability, etc...)
- no mature build tool as turn key, simple, and powerful as JVM build tools (Gradle, ...)
