#!/bin/bash

function init {
  cd ~
  mkdir reviews-directory
  cd reviews-directory
  echo 'Ready for reviews!'
}

function clear_projects {
  cd ~
  rm -rf reviews-directory
  echo 'Successfully removed review_directory!'
}

function review_project {
  cd ~
  cd reviews-directory
  git clone ${2}
  cd ${3}
  git fetch --all
  git checkout ${4}
  ${5} .
  npm install
  npm run test
}

function help {
  echo "Usage:"
  echo "./go init ...... Use before reviewer function!! Will setup review directory"
  echo "./go clear_projects ...... Deletes directory created by init command"
  echo "./go review_project [repository_url] [repository_name] [commit] [editorSymLink]"
  echo "the command above installs a project and runs tests for a particular SHA"
}

if [ -z "${1}"] ; then
  echo "The commands you can run are -"
  help
  exit 0
fi

case $1 in
  init) init $@
  ;;
  clear_projects) clear_projects $@
  ;;
  review_project) review_project $@
  ;;
  *) help
esac
