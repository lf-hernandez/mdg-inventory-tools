#!/bin/bash

git subtree push --prefix api heroku-backend main
git subtree push --prefix client heroku-frontend main

echo "Deployment finished!"
