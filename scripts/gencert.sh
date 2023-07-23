#!/bin/bash

if [ "$#" -ne 4 ]; then
    echo "Usage: $0 <rootCA.key> <rootCA.crt> <signer.csr> <signer.crt>"
    exit 1
fi

ROOT_CA_KEY=$1
ROOT_CA_CRT=$2
SIGNER_CSR=$3
SIGNER_CRT=$4

COUNTRY="JP"
STATE="Tokyo"
LOCALITY="Minato-ku"
ORGANIZATION="Only Test Inc."
ORGANIZATION_UNIT="Only Test Inc."
COMMON_NAME="only test inc"
EMAIL="onlytestinc@exapmle.com"

SUBJ="/C=$COUNTRY/ST=$STATE/L=$LOCALITY/O=$ORGANIZATION/OU=$ORGANIZATION_UNIT/CN=$COMMON_NAME/emailAddress=$EMAIL"

# generate root CA private key
openssl genrsa -out $ROOT_CA_KEY 2048

# generate root certificate
openssl req -x509 -new -nodes -key $ROOT_CA_KEY -sha256 -days 1825 -subj "$SUBJ" -out $ROOT_CA_CRT

# generate CSR
openssl req -new -sha256 -nodes -out $SIGNER_CSR -newkey rsa:2048 -keyout signer.key -subj "$SUBJ"

# generate Certificate
openssl x509 -req -in $SIGNER_CSR -CA $ROOT_CA_CRT -CAkey $ROOT_CA_KEY -CAcreateserial -out $SIGNER_CRT -days 1825 -sha256
