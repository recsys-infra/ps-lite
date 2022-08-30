//go:build tools

package api

//go:generate protoc -I . --go_out=paths=source_relative:. command.proto
//go:generate protoc -I . --go_out=paths=source_relative:. role.proto
//go:generate protoc -I . --go_out=paths=source_relative:. data_type.proto
//go:generate protoc -I . --go_out=paths=source_relative:. node.proto
//go:generate protoc -I . --go_out=paths=source_relative:. meta.proto
//go:generate protoc -I . --go_out=paths=source_relative:. control.proto
