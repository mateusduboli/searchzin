#!/bin/bash
USAGE="$0 <version tag>
Generates the messages betweeen the version tag and it's previous version.
Examples:
  Present tags:
    v1.0.0
    v1.1.0
  $0 v1.1.1 # Generates a message with the commits between v1.1.0..v1.1.1

  Present tags:
    v1.0.0
    v1.1.0
  $0 v1.0.1 # Generates a message with the commits between v1.0.0..v1.0.1
"
if [ "$#"-ne "1" ]; then
  echo $USAGE
  exit 1
fi

VERSION="$1"

if [[ ! "$VERSION" =~ v[0-9]+\.[0-9]+\.[0.9]+ ]]; then
  echo "Version must match semver $VERSION"
  exit 1
fi

PREVIOUS_TAG=$(git tag --sort=v:refname | grep -B 1 $VERSION | head -n1)

for C_MSG in $(git log $PREVIOUS_TAG..$VERSION); do
  echo "* $C_MSG"
done
