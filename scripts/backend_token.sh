#!/bin/sh
[ ! -f .env ] || export $(grep -v '^#' .env | xargs)

user_id=$(shuf -i 8551057951-9551057951 -n 1)

header=$(echo -n '{"alg":"HS256","typ":"JWT"}' | openssl base64 -e -A | tr '+/' '-_' | tr -d '=')

current_time=$(date +%s)
expiration_time=$(($current_time + 999999999999999999))
payload=$(echo -n '{"role": "admin", "user_id": '$user_id', "exp":'$expiration_time'}' | openssl base64 -e -A | tr '+/' '-_' | tr -d '=')

signature=$(echo -n "$header.$payload" | openssl dgst -sha256 -hmac $JWT_SECRET_KEY -binary | openssl base64 -e -A | tr '+/' '-_' | tr -d '=')

jwt="$header.$payload.$signature"

echo $jwt