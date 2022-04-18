#!/usr/local/bin/bash
set -e # exit when any command fails

declare -A md5Files
md5Files[".vx/commit/v1"]="testdata/vxcommitv1.md5"
md5Files[".vx/commit/v2"]="testdata/vxcommitv2.md5"
md5Files[".vx/checkout/v1"]="testdata/vxcheckoutv1.md5"
md5Files[".vx/checkout/v2"]="testdata/vxcheckoutv2.md5"

echoerr() { echo "$@" 1>&2; exit 1; }

buildCli() {
  make build-cli
  chmod +x vx
}
init() {
  rm -rf .vx || true
  ./vx init
}
status() {
  local OUTPUT=$(./vx status)
  echo "$OUTPUT"
}
history() {
  local OUTPUT=$(./vx history)
  echo "$OUTPUT"
}
add() {
  ./vx add "$@" --dtime=true
}
commitAs() {
  ./vx commit -m "$1" --dtime=true
}
checkoutTo() {
  ./vx checkout "$1"
}

checkEmptyStatus() {
  emptyStatusOut="No changes on staging area!"
  if [[ "$(status)" != "$emptyStatusOut" ]]; then
    echoerr "ERROR: '$(status)' cannot equal to '$emptyStatusOut'"
  fi
}
checkStatusAfterAdding() {
  if [[ "$(status)" != $(cat "$1") ]]; then
    echoerr "ERROR: '$(status)' cannot match with '$1"
  fi
}
checkFilesMatch() {
  if [[ $(md5deep -rl "$1" | sort) != $(cat ${md5Files["$1"]}) ]]; then
    echoerr "ERROR: ($1).md5 didn't match"
  fi
}
checkHistoryAfterApplyingAllCommits() {
  if [[ "$(history)" != $(cat "testdata/historyAfterAllCommits.txt") ]]; then
    echoerr "ERROR: '$(history)' cannot match with 'testdata/historyAfterAllCommits.txt'"
  fi
}

buildCli
init

add testdata/example README.md
checkStatusAfterAdding "testdata/firstAddStatusOutput.txt"
commitAs "first commit"
checkFilesMatch '.vx/commit/v1'

add testdata/Makefile
checkStatusAfterAdding 'testdata/secondAddStatusOutput.txt'
commitAs "second commit"
checkFilesMatch '.vx/commit/v2'

checkHistoryAfterApplyingAllCommits

checkoutTo v1
checkFilesMatch '.vx/checkout/v1'

checkoutTo v2
checkFilesMatch '.vx/checkout/v2'

echo "Success 🥳"