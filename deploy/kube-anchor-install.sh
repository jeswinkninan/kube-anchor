#!/bin/bash

echo "---Generating SSL SAN Certificate for Kube-Anchor Implementations---"

pwd=${pwd}

workspace=$pwd'./kube-anchor/deploy/'


openssl genrsa -out $workspace/ca.key 2048
openssl req -x509 -new -nodes -key $workspace/ca.key -subj "/CN=Kube-Anchor CA" -days 10000 -out $workspace/ca.crt
openssl genrsa -out $workspace/server.key 2048
openssl req -new -key $workspace/server.key -out $workspace/server.csr -config $workspace/csr.conf
openssl x509 -req -in $workspace/server.csr -CA $workspace/ca.crt -CAkey $workspace/ca.key  -CAcreateserial -out $workspace/server.crt -days 10000  -extensions v3_ext -extfile $workspace/csr.conf

echo "---Creating secrets using kubectl---"
kubectl -n kube-system create secret tls kube-anchor-webhook-server-tls --cert "$workspace/server.crt"  --key "$workspace/server.key"

echo "---Converting CA Cert to Base64 Encoded Format and Deploying Kube-Anchor---"
base64ca="$(openssl base64 -A <"$workspace/ca.crt")"
sed -e 's@${CA_Base64}@'"$base64ca"'@g' <"./$workspace/manifests.yaml" | kubectl create -n kube-system -f -



