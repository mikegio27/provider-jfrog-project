#!/usr/bin/env bash
# Patches upjet to handle single-word Terraform resource names (e.g. "project").
# The DefaultResource function assumes resource names have at least 2 underscore-
# separated words, which panics for providers like jfrog/project.
# Upstream issue: https://github.com/crossplane/upjet/issues

set -euo pipefail

UPJET_VERSION="v2.2.1-0.20260414070754-c6d5213346ac"
COMMON_GO="$(go env GOMODCACHE)/github.com/crossplane/upjet/v2@${UPJET_VERSION}/pkg/config/common.go"

if [ ! -f "${COMMON_GO}" ]; then
    echo "upjet common.go not found, skipping patch (module not yet downloaded)"
    exit 0
fi

if grep -q 'if len(words) < 2' "${COMMON_GO}"; then
    echo "upjet already patched, skipping"
    exit 0
fi

chmod +w "${COMMON_GO}"

python3 - "${COMMON_GO}" << 'PYEOF'
import sys

with open(sys.argv[1], 'r') as f:
    content = f.read()

old = (
    '\twords := strings.Split(name, "_")\n'
    '\t// As group name we default to the second element if resource name\n'
    '\t// has at least 3 elements, otherwise, we took the first element as\n'
    '\t// default group name, examples:\n'
    '\t// - aws_rds_cluster => rds\n'
    '\t// - aws_rds_cluster_parameter_group => rds\n'
    '\t// - kafka_topic => kafka\n'
    '\tgroup := words[1]\n'
    '\t// As kind, we default to camel case version of what is left after dropping\n'
    '\t// elements before what is selected as group:\n'
    '\t// - aws_rds_cluster => Cluster\n'
    '\t// - aws_rds_cluster_parameter_group => ClusterParameterGroup\n'
    '\t// - kafka_topic => Topic\n'
    '\tkind := tjname.NewFromSnake(strings.Join(words[2:], "_")).Camel\n'
    '\tif len(words) < 3 {\n'
    '\t\tgroup = words[0]\n'
    '\t\tkind = tjname.NewFromSnake(words[1]).Camel\n'
    '\t}'
)

new = (
    '\twords := strings.Split(name, "_")\n'
    '\t// As group name we default to the second element if resource name\n'
    '\t// has at least 3 elements, otherwise, we took the first element as\n'
    '\t// default group name, examples:\n'
    '\t// - aws_rds_cluster => rds\n'
    '\t// - aws_rds_cluster_parameter_group => rds\n'
    '\t// - kafka_topic => kafka\n'
    '\tvar group, kind string\n'
    '\tif len(words) < 2 {\n'
    '\t\tgroup = words[0]\n'
    '\t\tkind = tjname.NewFromSnake(words[0]).Camel\n'
    '\t} else {\n'
    '\t\tgroup = words[1]\n'
    '\t\t// As kind, we default to camel case version of what is left after dropping\n'
    '\t\t// elements before what is selected as group:\n'
    '\t\t// - aws_rds_cluster => Cluster\n'
    '\t\t// - aws_rds_cluster_parameter_group => ClusterParameterGroup\n'
    '\t\t// - kafka_topic => Topic\n'
    '\t\tkind = tjname.NewFromSnake(strings.Join(words[2:], "_")).Camel\n'
    '\t\tif len(words) < 3 {\n'
    '\t\t\tgroup = words[0]\n'
    '\t\t\tkind = tjname.NewFromSnake(words[1]).Camel\n'
    '\t\t}\n'
    '\t}'
)

if old not in content:
    print("WARNING: patch target not found - upjet may have already been fixed upstream")
    sys.exit(0)

with open(sys.argv[1], 'w') as f:
    f.write(content.replace(old, new))
print("upjet patched successfully")
PYEOF
