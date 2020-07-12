.DEFAULT: protos

protos: generate

generate:
	sh generate_proto.sh