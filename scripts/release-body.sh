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
if [ "$#" -ne "1" ]; then
  echo $USAGE
  exit 1
fi

VERSION="$1"

if [[ ! "$VERSION" =~ v[0-9]+\.[0-9]+\.[0-9]+ ]]; then
  echo "Version doesn't matches semver $VERSION"
  exit 1
fi

PREVIOUS_TAG=$(git tag --sort=v:refname | grep -B 1 $VERSION | head -n1)

IFS=$'\n' # It will use n instead of newlines without the $

echo "## Changes"
for COMMIT in $(git log --pretty=format:%s $PREVIOUS_TAG..$VERSION); do
  echo "* $COMMIT"
done

echo "## Contributors"

UNIQUE_AUTHORS=$(git log --pretty=format:"%an <%ae>" $PREVIOUS_TAG..$VERSION )
for AUTHOR in $UNIQUE_AUTHORS; do
  echo "* $AUTHOR"
done
