{{ if keyExists (print "nginx_conf" ) }}{{ keyOrDefault (print "nginx_conf" ) "" }}
{{ else }}user  root;
worker_processes  auto;
worker_rlimit_nofile 100000;

error_log  stderr error;
pid        /var/run/nginx.pid;


events {
    worker_connections  16384;
    multi_accept        on;
    use                 epoll;
}


http {
    include       /etc/nginx/mime.types;
    default_type  application/octet-stream;

    real_ip_header X-Forwarded-For;
    set_real_ip_from 10.0.0.0/8;

    log_format  main  '$host $remote_addr $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent $request_time "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';

    access_log  /dev/stdout main;

    sendfile        on;
    #tcp_nopush     on;

    keepalive_requests 1000;
    keepalive_timeout  100;

    server_names_hash_bucket_size   128;

    #gzip  on;

    include /etc/nginx/conf.d/*.conf;
    
    server {
        listen        80 default_server;
        server_name   @;

        location /status {
            stub_status   on;
            access_log    off;
            allow         127.0.0.1;
            deny          all;
        }

        location / {
            return  404;
        }
    }
}

stream {
    include /etc/nginx/stream.d/*.conf;
}
{{ end }}