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
- frontend linting
- structured logging

TODO:
- frontend unit tests for the example data access layer, controller, and page
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
- smaller pool of programmers with experience on it (although very close to Java, so Java developer would adapt quickly)
- more integration work for authentication solutions like Ory Kratos or SuperTokens

Cons of Go:
- less expressive
- less constructs available to avoid technical bugs (immutability, nullability, etc...)
- no mature build tool as powerful as JVM build tools (Gradle, ...)

## Making it work with VSCode:

1. Install the [Vue official extension](https://marketplace.visualstudio.com/items?itemName=Vue.volar).
2. [Disable the built in Typescript extension](https://stackoverflow.com/questions/54839057/vscode-showing-cannot-find-module-ts-error-for-vue-import-while-compiling-doe/73710755#73710755).
