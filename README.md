# Velveteen

[![CircleCI](https://circleci.enova.com/gh/devex/velveteen/tree/main.svg?style=svg)](https://circleci.enova.com/gh/devex/velveteen/tree/main)
[![codecov](https://enova.codecov.io/ghe/devex/velveteen/graph/badge.svg)](https://enova.codecov.io/ghe/devex/velveteen)

> test app

## Observability

- [Sentry](https://sentry.aws.enova.com/enova/velveteen)
- [Jenkins](https://pipeline.enova.com/view/Container%20Services/view/devex%20velveteen)
- [Splunk](https://splunk.enova.com/en-US/app/cnu_search/search?q=search%20%22parsed.app%22%3D%22velveteen%22)
- [Datadog - Wharf Service Dashboard](https://app.datadoghq.com/dashboard/n6h-pdy-p3q?tpl_var_environment=production&tpl_var_service%5B0%5D=velveteen)

## Environment variables

The project uses the following environment variables:

| Variable  | Definition                                          |
| --------- | --------------------------------------------------- |
| `ENV`     | Set up app.Env() (development, staging, production) |
