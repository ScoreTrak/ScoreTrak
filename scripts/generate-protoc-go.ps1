# https://github.com/grpc/grpc-go/issues/3669#issuecomment-692639536
protoc .\pkg\proto\utilpb\uuid.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  --experimental_allow_proto3_optional
protoc .\pkg\check\checkpb\check.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  --experimental_allow_proto3_optional
protoc .\pkg\property\propertypb\property.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  --experimental_allow_proto3_optional
protoc .\pkg\round\roundpb\round.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  --experimental_allow_proto3_optional
protoc .\pkg\service\servicepb\service.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  --experimental_allow_proto3_optional
protoc .\pkg\service_group\service_grouppb\service_group.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  --experimental_allow_proto3_optional
protoc .\pkg\host\hostpb\host.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  --experimental_allow_proto3_optional
protoc .\pkg\host_group\host_grouppb\host_group.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  --experimental_allow_proto3_optional
protoc .\pkg\team\teampb\team.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  --experimental_allow_proto3_optional
protoc .\pkg\competition\competitionpb\competition.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  --experimental_allow_proto3_optional
protoc .\pkg\report\reportpb\report.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  --experimental_allow_proto3_optional
protoc .\pkg\config\configpb\config.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  --experimental_allow_proto3_optional
protoc .\pkg\user\userpb\user.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  --experimental_allow_proto3_optional
protoc .\pkg\policy\policypb\policy.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  --experimental_allow_proto3_optional
protoc .\pkg\auth\auth.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  --experimental_allow_proto3_optional