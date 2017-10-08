<!-- vim: ft=markdown:
-->
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


```yml
port: 8080 # Service port
path:
  log: /var/log/searchzin # Log directory
  data: /var/lib/searchzin # Data directory
```


## Development

All the project structure is made in golang, using the
[`gin`](https://github.com/gin-gonic/gin) framework.

Dependencies are managed using [`dep`](https://github.com/golang/dep).

Most of the project toolchain is managed by the
[`Makefile`](https://github.com/mateusduboli/searchzin/tree/master/Makefile),
the important targets are:

* `install`: Install needed dependencies and git hooks
* `readme`: Performs `README.md` inclusion of files
* `lint`: Performs linting and formatting of the code
* `test`: Well, compile and run unit tests
* `build`: Creates a `linux` distributable folder in `dist`
* `run`: Runs a docker container with the `dist` executable
* `release`: Creates a release version on the `dist` folder
* `release-dev`: Creates a release version to be used in `run`
* `publish`: Publishes the docker image in dockerhub using the git sha as
    version
* `publish-latest`: Publishes the docker image in dockerhub with the `latest` tag
*  (*TODO*) `watch`: Performs `lint` and `test` on file modification
*  (*TODO*) `func-test`: Performs functional tests inside the `features` folder

## Architecture

There are 6 main components to this search engine:
* Document database
* Index database
* Indexing service
* Query executor
* Query planner
* Query parser

Each component has a clear responsability in the system, and all of them work
togheter to respond to queries and document indexing requests.

### Document Database

It's responsible to store and give id's to newly created documents. The
constraints are:

* Stores documents and their `id`s
* Enables `id` generation with little to no collisions
* Efective document storing algorithm, being optimized for fast reads and fast
    enought writes
* Aware of the underlying storage unit, being it `ssd` or `hdd`
* Aware of the underlying linux page size and file caching strategy

### Index database

Stores a reverse-index of "terms" and documents

* Stores `terms` to document set relations
* Enable `key` manipulation strategies for queries with keyword approximation
* Optimized for low density keys with lots of documents
* Aware of the underlying linux page size to easily fit and be loaded in-memory

### Indexing service

Given a new document understands it and saves both on the index database and the
document database.

* Knows which fields are indexed and how
* Knows the document structure and can related that to the indexes

### Query parser

Parses the user input and transforms it into a query plan using a tree-like
data structure.

* Parses a string given by the user and turns it into a graph
* The DSL will be similar to [`lucene`'s](https://lucene.apache.org/)

### Query planner

Given a query tree, optimizes it being aware of the restrictions and the
environment in which it will be executed.

* Remove redundant results, making them available to all the steps that need
* Aware of index size to sort which effective retrievals will be done first
* Returns an ordered list of query nodes to be executed

### Query executor

After having a structured plan the query then retrieves effective data from the
`index` database, this step is performed by the executor.

* Knows how to query the index database
* Joins the results given by it in a ordered fashion
* Retrieves the documents
* Stores the query results in a file to be queried later using "cold" storage

## Query language

This query language is heavily based on lucene's, to simplify design and
understand what tradeoffs were made.