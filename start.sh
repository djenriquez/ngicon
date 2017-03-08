#!/bin/bash
export > /etc/envvars

# Initialize files
touch /etc/nginx/nginx.conf

mkdir -p /etc/nginx/stream.d/
touch /etc/nginx/stream.d/stream.conf

touch /etc/nginx/conf.d/app.conf

/usr/bin/runsvdir /etc/service