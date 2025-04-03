#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: $0 <base-branch>"
    exit 1
fi

BASE_BRANCH=$1

for branch in $(git for-each-ref --format="%(refname:short)" refs/remotes/origin/ | grep -v "^refs/remotes/origin/archive/"); do
    echo "Changes in $branch:"
    git diff --name-only origin/$BASE_BRANCH..$branch | grep "^_alp/Agents" || echo "No relevant changes."
done
