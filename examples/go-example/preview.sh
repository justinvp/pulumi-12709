#!/usr/bin/env bash

set -euo pipefail

pulumi stack select $(pulumi stack ls -j | jq -r '.[0].name')
export HC_SKIP_VPN_CHECK=true
export PATH=$PWD/../../bin:$PATH
export PULUMI_CONFIG_PASSPHRASE=1234567

pulumi destroy --yes

rm -rf "$PWD/log.json"
PULUMI_DEBUG_GRPC="$PWD/log.json" pulumi preview || echo ignoreFAIL
jq . "$PWD/log.json" > "$PWD/log-formatted.json"
