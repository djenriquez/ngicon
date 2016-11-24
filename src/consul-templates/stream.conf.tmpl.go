{{ if key (print "stream_block_config" ) }}
{{ key (print "stream_block_config" ) }}
{{ end }}

{{ range services }}{{ if and (.Tags.Contains "tcp") (key (print .Name "/tcp_listen" )) }}
upstream {{ .Name }} {
  {{ range service .Name }}server        {{ .Address }}:{{ .Port }};
  {{ end }}
}
{{ end }}{{end}}

{{ range services }}{{ if and (.Tags.Contains "tcp") (key (print .Name "/tcp_listen" )) }}
server {
  listen        {{ key (print .Name "/tcp_listen" ) }};
  proxy_pass    {{ .Name }};
}
{{ end }}{{ end }}