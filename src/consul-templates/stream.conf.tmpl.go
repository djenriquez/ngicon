{{ keyOrDefault (print "stream_block_config" ) "" }}

{{ range services }}{{ if and (.Tags | contains "tcp") (keyExists (print .Name "/tcp_listen" )) }}
upstream {{ .Name }} {
  {{ range service .Name }}server        {{ .Address }}:{{ .Port }};
  {{ end }}
}
{{ end }}{{end}}

{{ range services }}{{ if and (.Tags | contains "tcp") (keyExists (print .Name "/tcp_listen" )) }}
server {
  listen        {{ keyOrDefault (print .Name "/tcp_listen" ) "" }};
  proxy_pass    {{ .Name }};
}
{{ end }}{{ end }}