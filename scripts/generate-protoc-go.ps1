# https://github.com/grpc/grpc-go/issues/3669#issuecomment-692639536
protoc .\scoretrakapis\pkg\proto\proto\v1\uuid.proto --go_out=.\ --go_opt=paths=source_relative --go-grpc_out=.\ --go-grpc_opt=paths=source_relative --proto_path=.\scoretrakapis
protoc .\scoretrakapis\pkg\proto\check\v1\check.proto --go_out=.\ --go_opt=paths=source_relative --go-grpc_out=.\ --go-grpc_opt=paths=source_relative --proto_path=.\scoretrakapis
protoc .\scoretrakapis\pkg\proto\property\v1\property.proto --go_out=.\ --go_opt=paths=source_relative --go-grpc_out=.\ --go-grpc_opt=paths=source_relative --proto_path=.\scoretrakapis
protoc .\scoretrakapis\pkg\proto\round\v1\round.proto --go_out=.\ --go_opt=paths=source_relative --go-grpc_out=.\ --go-grpc_opt=paths=source_relative --proto_path=.\scoretrakapis
protoc .\scoretrakapis\pkg\proto\service\v1\service.proto --go_out=.\ --go_opt=paths=source_relative --go-grpc_out=.\ --go-grpc_opt=paths=source_relative --proto_path=.\scoretrakapis
protoc .\scoretrakapis\pkg\proto\service_group\v1\service_group.proto --go_out=.\ --go_opt=paths=source_relative --go-grpc_out=.\ --go-grpc_opt=paths=source_relative --proto_path=.\scoretrakapis
protoc .\scoretrakapis\pkg\proto\host\v1\host.proto --go_out=.\ --go_opt=paths=source_relative --go-grpc_out=.\ --go-grpc_opt=paths=source_relative --proto_path=.\scoretrakapis
protoc .\scoretrakapis\pkg\proto\host_group\v1\host_group.proto --go_out=.\ --go_opt=paths=source_relative --go-grpc_out=.\ --go-grpc_opt=paths=source_relative --proto_path=.\scoretrakapis
protoc .\scoretrakapis\pkg\proto\team\v1\team.proto --go_out=.\ --go_opt=paths=source_relative --go-grpc_out=.\ --go-grpc_opt=paths=source_relative --proto_path=.\scoretrakapis
protoc .\scoretrakapis\pkg\proto\competition\v1\competition.proto --go_out=.\ --go_opt=paths=source_relative --go-grpc_out=.\ --go-grpc_opt=paths=source_relative --proto_path=.\scoretrakapis
protoc .\scoretrakapis\pkg\proto\report\v1\report.proto --go_out=.\ --go_opt=paths=source_relative --go-grpc_out=.\ --go-grpc_opt=paths=source_relative --proto_path=.\scoretrakapis
protoc .\scoretrakapis\pkg\proto\config\v1\config.proto --go_out=.\ --go_opt=paths=source_relative --go-grpc_out=.\ --go-grpc_opt=paths=source_relative --proto_path=.\scoretrakapis
protoc .\scoretrakapis\pkg\proto\user\v1\user.proto --go_out=.\ --go_opt=paths=source_relative --go-grpc_out=.\ --go-grpc_opt=paths=source_relative --proto_path=.\scoretrakapis
protoc .\scoretrakapis\pkg\proto\policy\v1\policy.proto --go_out=.\ --go_opt=paths=source_relative --go-grpc_out=.\ --go-grpc_opt=paths=source_relative --proto_path=.\scoretrakapis
protoc .\scoretrakapis\pkg\proto\auth\v1\auth.proto --go_out=.\ --go_opt=paths=source_relative --go-grpc_out=.\ --go-grpc_opt=paths=source_relative --proto_path=.\scoretrakapis