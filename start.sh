#!/bin/bash
export > /etc/envvars

# Initialize files
touch /etc/nginx/nginx.conf
touch /etc/nginx/conf.d/app.conf
touch /etc/nginx/conf.d/stream.conf

/usr/bin/runsvdir /etc/service