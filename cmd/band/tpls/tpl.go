package tpls

import "embed"

//go:embed service
var ServiceFiles embed.FS

//go:embed server
var ServerFiles embed.FS

//go:embed proto
var ProtoFiles embed.FS
