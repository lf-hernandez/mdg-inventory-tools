#!/bin/bash

git subtree push --prefix api heroku-backend master
git subtree push --prefix client heroku-frontend master

echo "Deployment finished!"
