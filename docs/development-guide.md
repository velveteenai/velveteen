# Development guide

## Project layout

The project was structured based on the (_nonofficial_) standard [Go project layout](https://github.com/golang-standards/project-layout) and taking in consideration some common [good practices](https://rakyll.org/style-packages) for writing packages.

### Dependencies

The project uses [barad-dur](https://git.enova.com/go/barad-dur) for exposing utility endpoints and logging. If logging is required in a function, inject the `barad-dur` application and use the functions available from `app.Log`.

```go
func hello(app *baraddur.Application) {
  app.Log.Print("Executing hello")
}
```

To log errors to Sentry, just log as `Error()` using the `barad-dur` application. Errors will be sent based on the `SENTRY_DSN` environment variable. Here are some examples of errors that will be sent to Sentry:

```go
app.Log.WithError(err).Error()
```

```go
app.Log.Error(err)
```

```go
app.Log.Errorf("some error...")
```

To log to Splunk, just log as `Info()` or `Warn()` using the `barad-dur` application. Logs will be sent collected by wharf and transmitted to Splunk automatically.

```go
app.Log.Info("Message")
```

### Interface compliance

We verify interface compliance at compile time by adding the following statement alongside any structs that implement an interface ([read more](https://github.com/uber-go/guide/blob/master/style.md#verify-interface-compliance)). Among other benefits, it avoids the need for unit testing whether or not a struct _X_ has implemented an interface _Y_.

```go
var _ somepkg.SomeInterface = (*SomeStruct)(nil)
```

## Installation

1. Install Go (see [Getting Started with Go](https://wiki.enova.com/display/EA/Getting+Started+with+Go))
2. Install [Docker](https://www.docker.com)
3. Clone this repository: `git clone https://git.enova.com/devex/velveteen.git`
4. Get Docker images ready: `docker-compose build`
5. Get your local environment ready (not needed if you use Docker): `make setup`

## Running

The *Velveteen* can be run using *Docker* (recommended):

```bash
docker-compose up -d
```

Or you can execute it locally:

```bash
go run main.go
```

## Development

The following commands will run locally by default. To execute them from __velveteen__ container do:

```bash
docker compose exec velveteen make <COMMAND>
```

### Build executable

```bash
make build
```

### Regenerate code (including mocks from mockgen)

```bash
go generate ./...
```

### Run unit tests

*Run all tests:*

```bash
make test
```

*You can modify the test run by adding some of these options: `TEST_FILES`, `TEST_PATTERN`, `TEST_OPTIONS`*

```bash
make test TEST_FILES=./some_folder/... TEST_PATTERN=MyTest TEST_OPTIONS="-count=1 -v"
```

### Format code indentation and imports

```bash
make fmt
```

### Run linter (`golangci-lint`)

```bash
make lint
```

### Remove object files, test results and executables

```bash
make clean
```

## Unit testing

All tests should be created using a [test package](https://tip.golang.org/cmd/go/#hdr-Test_packages), which means appending a `_test` to the tested package name, in order to write black-box tests. The project uses the following testing packages to help design the tests and mocks:

* HTTP requests with [Gock](https://github.com/h2non/gock)
* General mocks with [Gomock](https://github.com/golang/mock)
* General asserts with [Testify](https://github.com/stretchr/testify)

### Writing tests

Please refer to [Enova's best practices](https://git.enova.com/pages/Automation/learning-testing/unit_testing/unit_testing/#best-practices) when writing unit tests, especially when in doubt about [naming](https://git.enova.com/pages/Automation/learning-testing/unit_testing/unit_testing/#naming-your-tests), [arranging](https://git.enova.com/pages/Automation/learning-testing/unit_testing/unit_testing/#arranging-your-tests) and [ordering](https://git.enova.com/pages/Automation/learning-testing/unit_testing/unit_testing/#following-the-flow-of-the-application) them. When naming tests we go one step further and also add the struct __and/or__ the method being tested, something similar to the naming convention to declare [Examples](https://pkg.go.dev/testing#hdr-Examples) in the Go testing library.

So let's say you have the following piece of code:

```go
type MyStruct struct {
  // ...
}

func (s MyStruct) MyFunction() error {
  // ...
}
```

When adding a test to validate that code, the test name should look like this:

```go
func TestMyStruct_MyFunction_ReturnsErrorIfSomeGoesWrong(t *testing.T) {
  // ...
}
```

Please take a look at some existing tests for more examples of this pattern usage.

### Mocking

The command to generate mock files can be found in the beginning of the interface files. In case you need to add a brand new interface, make sure to follow the example below:

```go
//go:generate mockgen -destination=./mock/mock_interface.go -package=mock . MyInterface

// MyInterface ...
type MyInterface interface {
  // ...
}
```

After that, running `go generate ./...` will generate the mock file for the interface.

## Keeping docs up to date

This project uses MkDocs. For full documentation visit [mkdocs.org](https://www.mkdocs.org).

In general, you will just need Python (v3.8 or later is recommended), pip, mkdocs, and mkdocs material theme installed.

Having Python and pip installed, install MkDocs running:

```bash
pip3 install mkdocs # or pip, depending on the installation
```

To install mkdocs-material, run:

```bash
pip3 install mkdocs-material
```

### Commands

* `mkdocs new [dir-name]` - Create a new project.
* `mkdocs serve` - Start the live-reloading docs server.
* `mkdocs build` - Build the documentation site.
* `mkdocs -h` - Print help message and exit.
* `mkdocs gh-deploy` - Deploy documentation.

### Docs layout

```yaml
mkdocs.yml  # The configuration file.
docs/
  index.md  # The documentation homepage.
  ...       # Other markdown pages, images and other files.
```

### Markdown tricks

!!! note
    mkdocs-material has a lot of awesome extensions to turn the experience of writing docs really pleasant. If you want to know more, check [Material for MkDocs](https://squidfunk.github.io/mkdocs-material/getting-started).

To add a block like the one above you may use a `!!!` indicator followed by a type.

Example:

```markdown
!!! note
    Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla et euismod nulla. Curabitur feugiat, tortor non consequat finibus, justo purus auctor massa, nec semper lorem quam in massa.
```

## Release process

See [deployment](deployment.md) on the following page.
