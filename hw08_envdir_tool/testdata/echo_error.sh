#!/usr/bin/env bash

result=$(./go-envdir "$(pwd)/testdata/env" "/bin/bash" "$(pwd)/testdata/error_code.sh" arg1=1 arg2=2)
status=$?
echo ${status}
