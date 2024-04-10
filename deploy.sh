#!/bin/bash
set -e

git subtree push --prefix api heroku-server main
git subtree push --prefix client heroku-client main

echo "Deployment finished!"
