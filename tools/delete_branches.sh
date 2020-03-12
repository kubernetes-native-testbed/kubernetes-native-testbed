#!/bin/bash
git pull

for i in `git branch -a | grep remotes | grep -v develop | grep -v master`; do
  git push --delete origin  $(echo $i | perl -p -e 's|remotes/origin/||g') &;
done


