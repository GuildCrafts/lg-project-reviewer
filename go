#!/bin/bash

function init {
  cd ~
  mkdir lg-reviews-directory
  cd lg-reviews-directory
  echo 'Ready for reviews!'
}

function clear_projects {
  cd ~
  rm -rf lg-reviews-directory
  echo 'Successfully removed lg-reviews-directory!'
}

function review_project {
  cd ~
  cd lg-reviews-directory
  git clone ${2}
  cd ${3}
  git fetch --all
  git checkout ${4}
  ${5} .
  npm install
  npm run test
  echo 'The branch is in detached HEAD state...'
}

function help {
  echo "Usage:"
  echo "sh go init ...... Use before reviewer function!! Will setup review directory"
  echo "sh go clear_projects ...... Deletes directory created by init command"
  echo "sh go review_project [repository_url] [repository_name] [commit] [editorSymLink]"
  echo "the command above installs a project and runs tests for a particular SHA"
}

if [ -z ${1} ] ; then
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
