#pkg/handler

Every pkg/handler/{model_name}.{server_type(grpc|connect)}.go has:
- A struct that implements API. The struct also embeds in itself interface described in pkg/{model_name}/{model_name}_service/serv.go that it can use in order to communicate with database.
- A converter method that converts proto representations of models to actual models, and vice versa.
- A set of methods that outline available API. The also match with respective methods outlined in pkg/{model_name}pb/{model_name}.proto.