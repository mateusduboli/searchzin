# Searchzin

A simple search engine implementation.

## Motivation

Study purposes, mostly for understanding the implementation details of how
search engines are made, performance trade-offs and structure.

## Description

The idea is to make a isomorphic application from the UI to the database system.

## Usage

The application can be deployed using either docker or the binary released in
github.

```
./searchzin -c <path-to-config>.yml
```

After that you can look into `http://localhost:8080` to see the search page.

## Configuration

The configuration can be made by either the configuration file located by
default in `/etc/searchzin/config.yml`, or providing configuration keys in the
form `-C key=value`, the second form overrides the first.

Configuration defaults:

<!-- include:BEGIN config/config.yml -->
```yml

port: 8080 # Service port
path:
  log: /var/log/searchzin # Log directory
  data: /var/lib/searchzin # Data directory
```
<!-- include:END config/config.yml -->

## Development

All the project structure is made in golang, using the
[`gin`](https://github.com/gin-gonic/gin) framework.

Most of the project toolchain is managed by the
[`Makefile`](https://github.com/mateusduboli/searchzin/tree/master/Makefile),
the important targets are:

* `readme`: Performs `README.md` inclusion of files
* `lint`: Performs linting and formatting of the code
* `test`: Well, compile and run unit tests
* `watch`: Performs `lint` and `test` on file modification
* `func-test`: Performs functional tests inside the `features` folder
* `run`: Runs a local instance
* `release`: Creates a release version on the `dist` folder
* `publish-docker`: Publishes the docker image in dockerhub
* `publish-release`: Publishes the compiled version on github
* `publish-all`: Does all both publishings

## Architecture

On launching the app, 

## Search Engine Algorithm

