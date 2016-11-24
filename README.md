# Ngicon

## Summary
Ngicon is a dynamic reverse-proxy, using NGINX, that provides routes to services registered to a Consul service discovery system. It gives you the ability to register ephemeral ports, eliminating the need to explicitly register host ports to services. In doing so, one can register multiple instances of the same service on the same host.

Ngicon utilizes consul-template to synchronize with Consul's service directory. Ngicon offers the ability to update a cluster of ngicon's NGINX configs just by updating a consul key/value pair.

## HTTP(S)
Each service responds to a KV namespace matching the service name. For example, to modify a service's NGINX service block, create a `service_block_config` key for the service's namespace and add configuration that you need as that key's value.

Ngicon also supports `location_block_config` which modifies configuration at the root `/` location for a service.

## TCP
Ngicon supports TCP proxying. To enable TCP proxying, register the tag `tcp` to the Consul service requiring TCP proxying. Then add a key `tcp_listen` to the namespace of the respective service with the value of the port you would like the service to be accessed through.

## Run
```bash
docker run -d \
--net host \
--name ngicon \
djenriquez/ngicon
```

## Environment Variables
* CONSUL_ADDRESS: The address of Consul. It is recommended to point Ngicon at the local Consul service. Defaults to `localhost`
* CONSUL_PORT: The port to access Consul. Defaults to `8500`