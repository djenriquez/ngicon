{{ if key (print "http_block_config" ) }}
{{ key (print "http_block_config" ) }}
{{ end }}

{{range services}}upstream {{ .Name }}{
        {{ range service .Name }}server {{.Address}}:{{.Port}};
        {{ end }}
}
{{end}}

{{ range services }}{{ if key (print .Name "/domains" ) }}server {
  listen 80;
  server_name {{ key (print .Name "/domains" ) }};
  {{ if key (print .Name "/server_block_config" ) }}{{ key (print .Name "/server_block_config" ) }}{{ end }}
  location / {
    {{ if key (print .Name "/location_block_config" ) }}{{ key (print .Name "/location_block_config" ) }}{{ end }}
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;

    proxy_pass  http://{{ .Name}}{{if key (print .Name "/proxy_uri" )}}{{key (print .Name "/proxy_uri" )}}{{end}};
  }
}
{{ end }}{{ end }}