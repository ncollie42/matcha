OUTDIR=../backend/generated
PROTO=matcha.proto
all:
	protoc $(PROTO) --go_out=plugins=grpc:$(OUTDIR)
