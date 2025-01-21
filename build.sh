#!/bin/bash
set -e

root="$PWD"/dependencies/darwin-arm64
export DEVELOP_ROOT=$root
export PKG_CONFIG_PATH=$root/lib/pkgconfig
export PKGCONFIG_EXE=$root/bin/pkg-config
$root/python/Python.framework/Versions/Current/bin/python3 setup.py "$@"
