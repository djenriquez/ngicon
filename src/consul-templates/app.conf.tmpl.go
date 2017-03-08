{{ keyOrDefault (print "http_block_config" ) "" }}

{{ range services }}{{ if keyExists (print .Name "/domains" ) }}upstream {{ .Name }}{
        {{ range service .Name }}server {{.Address}}:{{.Port}};
        {{ end }}
}

server {
  listen 80;
  server_name {{ keyOrDefault (print .Name "/domains" ) "" }};
  {{ keyOrDefault (print .Name "/server_block_config" ) "" }}
  location / {
    {{ keyOrDefault (print .Name "/location_block_config" ) "" }}
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;

    proxy_pass  http://{{ .Name}}{{ keyOrDefault (print .Name "/proxy_uri" ) "" }};
  }
}
{{ end }}{{ end }}