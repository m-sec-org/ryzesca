echo "生成 rpc server 代码"
OUT=../server
GOPATH=$(go env GOPATH)
protoc -I ./ \
	-I "${GOPATH}"/src \
	-I "${GOPATH}"/src/google/protobuf \
  --go_out=${OUT} \
  --go-grpc_out=${OUT} \
  --go-grpc_opt=require_unimplemented_servers=false \
  ryzesca.proto

#echo "生成 rpc client 代码"
#
#OUT=../client
#GOPATH=$(go env GOPATH)
#protoc -I ./ \
#	-I "${GOPATH}"/src \
#	-I "${GOPATH}"/src/google/protobuf \
#  --go_out=${OUT} \
#  --go-grpc_out=${OUT} \
#  --go-grpc_opt=require_unimplemented_servers=false \
#  ryzesca.proto
#
## 让他等待不关闭终端
#read -p "Press any key to continue." var