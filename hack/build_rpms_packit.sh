#!/usr/bin/env bash
set -ex
# this only works on rpm based systems, for non-rpm this is wrapped by build_rpms.sh
packit 2>/dev/null >/dev/null || (echo "Installing packit" && sudo dnf install -y packit)
rm -f "$(uname -m)"/flightctl-*.rpm 2>/dev/null || true
rm -f bin/rpm/* 2>/dev/null || true
mkdir -p bin/rpm
# save the spec as packit will modify it locally to inject versioning and we don't want that
cp packaging/rpm/flightctl.spec /tmp
packit build locally
cp /tmp/flightctl.spec packaging/rpm
mv noarch/flightctl-*.rpm bin/rpm
mv $(uname -m)/flightctl-*.rpm bin/rpm
