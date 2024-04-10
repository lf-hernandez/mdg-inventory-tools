#!/bin/bash
set -e

git subtree push -P api heroku-server main
git subtree push -P client heroku-client main

echo "Deployment finished!"
