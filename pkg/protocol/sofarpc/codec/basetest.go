package codec
//
//import (
//	"fmt"
//	"gitlab.alipay-inc.com/afe/mosn/pkg/network/buffer"
//	"gitlab.alipay-inc.com/afe/mosn/pkg/protocol/sofarpc"
//)
//
//func protocal_test() {
//	//initilize codec engine. TODO: config driven
//	codecImpl := sofarpc.NewProtocols(map[byte]sofarpc.Protocol{
//		sofarpc.PROTOCOL_CODE_V1: BoltV1,
//		sofarpc.PROTOCOL_CODE_V2: BoltV2,
//		//protocol.PROTOCOL_CODE:sofarpc.Tr,
//	})
//
//	fakePipe := make(chan sofarpc.RpcCommand)
//
//	//plug-in tr codec
//	codecImpl.PutProtocol(sofarpc.PROTOCOL_CODE, Tr)
//
//	boltV1EchoBytes := []byte{0x01, 0x01, 0x00, 0x01, 0x01, 0x00, 0x00, 0x00, 0x72, 0x01, 0x00, 0x00, 0x00, 0x64, 0x00, 0x2c, 0x00, 0x45, 0x00, 0x00, 0x01, 0xe0, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x66, 0x61, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x6f, 0x66, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x00, 0x00, 0x00, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x00, 0x00, 0x00, 0x36, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x61, 0x64, 0x65, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x31, 0x2e, 0x30, 0x4f, 0xbc, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x66, 0x61, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x6f, 0x66, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x95, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x17, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x0d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x41, 0x72, 0x67, 0x53, 0x69, 0x67, 0x73, 0x6f, 0x90, 0x07, 0x72, 0x70, 0x63, 0x2d, 0x62, 0x61, 0x72, 0x07, 0x65, 0x63, 0x68, 0x6f, 0x53, 0x74, 0x72, 0x53, 0x00, 0x36, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x61, 0x64, 0x65, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x31, 0x2e, 0x30, 0x4d, 0x11, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x4d, 0x09, 0x73, 0x6f, 0x66, 0x61, 0x52, 0x70, 0x63, 0x49, 0x64, 0x01, 0x30, 0x07, 0x45, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x01, 0x46, 0x0b, 0x73, 0x79, 0x73, 0x50, 0x65, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x73, 0x00, 0x0d, 0x73, 0x6f, 0x66, 0x61, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x63, 0x03, 0x64, 0x65, 0x76, 0x09, 0x7a, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x55, 0x49, 0x44, 0x00, 0x10, 0x7a, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x00, 0x0c, 0x73, 0x6f, 0x66, 0x61, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x70, 0x0d, 0x31, 0x31, 0x2e, 0x31, 0x36, 0x36, 0x2e, 0x32, 0x32, 0x2e, 0x31, 0x36, 0x31, 0x0b, 0x73, 0x6f, 0x66, 0x61, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x1e, 0x30, 0x62, 0x61, 0x36, 0x31, 0x36, 0x61, 0x31, 0x31, 0x35, 0x31, 0x34, 0x34, 0x33, 0x35, 0x33, 0x37, 0x31, 0x39, 0x36, 0x32, 0x31, 0x30, 0x30, 0x34, 0x34, 0x38, 0x30, 0x30, 0x35, 0x0c, 0x73, 0x6f, 0x66, 0x61, 0x50, 0x65, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x73, 0x00, 0x0e, 0x73, 0x6f, 0x66, 0x61, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x5a, 0x6f, 0x6e, 0x65, 0x05, 0x47, 0x5a, 0x30, 0x30, 0x42, 0x0d, 0x73, 0x6f, 0x66, 0x61, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x41, 0x70, 0x70, 0x07, 0x72, 0x70, 0x63, 0x2d, 0x66, 0x6f, 0x6f, 0x0d, 0x7a, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x03, 0x31, 0x30, 0x30, 0x7a, 0x7a, 0x56, 0x74, 0x00, 0x07, 0x5b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x6e, 0x01, 0x10, 0x6a, 0x61, 0x76, 0x61, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x7a, 0x01, 0x61}
//
//	trPostData := buffer.NewIoBufferString("\x0d\x00TaobaoRemoting")
//	boltV1PostData := buffer.NewIoBufferString(string(boltV1EchoBytes))
//	boltV2PostData := buffer.NewIoBufferString("\x02\x00BoltV2")
//
//	//test tr decode branch
//	fmt.Println("-----------> tr test begin")
//	codecImpl.Decode(nil, trPostData, nil)
//	fmt.Println("<----------- tr test end\n")
//
//	//test boltv1 decode branch
//	fmt.Println("-----------> boltv1 test begin")
//	go codecImpl.Decode(nil, boltV1PostData, fakePipe)
//	fmt.Println("<----------- boltv1 test end\n")
//
//	//test boltv2 decode branch
//	fmt.Println("-----------> boltv2 test begin")
//	codecImpl.Decode(nil, boltV2PostData, nil)
//	fmt.Println("<----------- boltv2 test end\n")
//
//	msg := <-fakePipe
//	codecImpl.Handle(1, nil, msg)
//}
