# GNAPI

This project provides an http web access to [RESTful APIs documentation] of Global Names
projects. These APIs follow OpenAPI standard.

## Editing API docs with swagger-editor

The easiest way to edit documents is by using a docker image of [swagger-editor]

```bash
docker run -d -p 80:8080 swaggerapi/swagger-editor
```

## Usage

Run the service from the command line

```bash
gnapi -p 8888
```

Run from a docker image

```bash
docker run -d -p 80:8888 gnames/gnapi
```

[RESTful APIs documentation]: https://apidoc.globalnames.org
