# Mastercard Match API

### Install Revel

```bash
$ go get github.com/revel/revel
$ go get github.com/revel/cmd/revel
```

### Start the web server:

```bash
# export the required environment variables

# one by one
$ export REVEL_APP_SECRET=value
$ export MASTERCARD_CONSUMER_KEY=value
$ export MASTERCARD_SIGNING_KEY_PATH=value
$ export MASTERCARD_SIGNING_KEY_SECRET=value
$ export MASTERCARD_API_URL=value
# OR copy and fill the values in your env file
$ cp .env.example .env
$ source .env

# run the aplication
$ revel run -a  mastercard-match
```
### Integrating with OpenAPI Generator API Client Libraries

OpenAPi Generator

Client libraries can be generated using the following command:

```bash
$ java -jar openapi-generator-cli.jar generate -i app/services/mastercard/match/api/openapi.yaml -g go --additional-properties=packageName=match -o match
```

See also:

* [OpenAPI Match Docs](app/services/mastercard/match/README.md)
* [OpenAPI Generator](https://mvnrepository.com/artifact/org.openapitools/openapi-generator-cli)
* [CONFIG OPTION for Go](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/go.md)



## Code Layout Revel conventions

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites



