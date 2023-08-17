#!/usr/bin/env bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
ROOT_PKG=github.com/hanjialeok/simple-operator

GET_PKG_LOCATION() {
  pkg_name="${1:-}"

  pkg_location="$(go list -m -f '{{.Dir}}' "${pkg_name}" 2>/dev/null)"
  if [ "${pkg_location}" = "" ]; then
    echo "${pkg_name} is missing. Running 'go mod download'."

    go mod download
    pkg_location=$(go list -m -f '{{.Dir}}' "${pkg_name}")
  fi
  echo "${pkg_location}"
}

# Grab code-generator version from go.sum
CODEGEN_PKG="$(GET_PKG_LOCATION "k8s.io/code-generator")"
echo ">> Using ${CODEGEN_PKG}"

# Ensure we can execute.
chmod +x ${CODEGEN_PKG}/generate-groups.sh
chmod +x ${CODEGEN_PKG}/generate-internal-groups.sh

# code-generator does work with go.mod but makes assumptions about
# the project living in `$GOPATH/src`. To work around this and support
# any location; create a temporary directory, use this as an output
# base, and copy everything back once generated.
TEMP_DIR=$(mktemp -d)
cleanup() {
    echo ">> Removing ${TEMP_DIR}"
    rm -rf ${TEMP_DIR}
}
trap "cleanup" EXIT SIGINT

echo ">> Temporary output directory ${TEMP_DIR}"

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.

${CODEGEN_PKG}/generate-internal-groups.sh "deepcopy,defaulter,client,lister,informer,openapi" \
    github.com/hanjialeok/simple-operator/pkg/client github.com/hanjialeok/simple-operator/pkg/apis github.com/hanjialeok/simple-operator/pkg/apis \
    kubeflow.org:v1 \
    --output-base "${TEMP_DIR}" \
    --go-header-file ${SCRIPT_ROOT}/hack/boilerplate.go.txt

# Copy everything back.
cp -a "${TEMP_DIR}/${ROOT_PKG}/." "${SCRIPT_ROOT}/"
