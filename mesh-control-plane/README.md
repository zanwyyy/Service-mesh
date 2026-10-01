# mesh-control-plane

Initial Golang scaffold for a service mesh control plane:

- `cmd/controlplane`: entrypoint
- `internal/discovery`: service discovery
- `internal/config`: config source/watch
- `internal/xds`: xDS server
- `internal/health`: health aggregation
- `internal/store`: in-memory state
- `api/proto`: protobuf API definitions
- `config`: local config files
- `deploy`: local deployment manifests
