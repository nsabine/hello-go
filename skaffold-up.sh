#!/bin/bash

git clone https://github.com/tariq-islam/hello-go && cd hello-go
curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/latest/skaffold-linux-amd64 && chmod +x skaffold && sudo mv skaffold /usr/local/bin