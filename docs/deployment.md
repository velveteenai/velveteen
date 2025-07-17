# Deployment

## Release process

### Staging

* Merging will automatically trigger staging release
* [Circle CI runs on main](https://circleci.enova.com/gh/devex/velveteen/tree/main) and notifies Jenkins
* It is possible to monitor Jenkins at [Enova's Pipeline](https://pipeline.enova.com/view/Container%20Services/view/devex%20velveteen)

!!! warning "Important"
    Wait for the staging process to fully complete before starting the production release process.

### Production

* Create a new tag using `taggit` tool:
  * `taggit tick devex/velveteen main --minor`
  * `taggit tick devex/velveteen_overlays main --patch` (if required)
* Create a new ticket using `pipeline` tool
  * `./pipeline devex/velveteen -s velveteen -t container -d -a <APP_TAG> -o <APP_OVERLAYS_TAG>`
  * It is possible to monitor Jenkins at [Enova's Pipeline](https://pipeline.enova.com/view/Container%20Services/view/devex%20velveteen)
  * It is possible to monitor the generated ticket in the [Software Release Agile Board](https://agile.enova.com/secure/RapidBoard.jspa?rapidView=2)

!!! note "Versioning"
    The usage of `--minor`, `--patch`, or `--major` depends on the changes made to the project. Please refer to [Semantic Versioning](https://semver.org) for more details about version schema.

### Post-release checks

* Visit the [Wharf Service Dashboard](https://app.datadoghq.com/dashboard/n6h-pdy-p3q?tpl_var_environment=production&tpl_var_service%5B0%5D=velveteen)
* Filter for the environment you just released (`staging` or `production`)
* Keep an eye for `Total Containers Restarts` and the charts for `Memory Usage` and `CPU Usage`.

!!! warning "Important"
    If the containers keep constantly restarting, this means something is wrong with the app. You should consider reverting the release.
