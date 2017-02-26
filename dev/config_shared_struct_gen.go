// DO NOT EDIT. Generated by 'gorums' plugin for protoc-gen-go
// Source file to edit is: config_shared_struct_tmpl

package dev

import "fmt"

//TODO Make this a customizable struct that replaces FQRespName together with typedecl option in gogoprotobuf.
//(This file could maybe hold all types of structs for the different call semantics)

// GorumsQCReadQCReply encapsulates the reply from a correctable GorumsQCReadQC quorum call.
// It contains the id of each node of the quorum that replied and a single reply.
type GorumsQCReadQCReply struct {
	NodeIDs []uint32
	*Reply
}

func (r GorumsQCReadQCReply) String() string {
	return fmt.Sprintf("node ids: %v | answer: %v", r.NodeIDs, r.Reply)
}

//TODO Make this a customizable struct that replaces FQRespName together with typedecl option in gogoprotobuf.
//(This file could maybe hold all types of structs for the different call semantics)

// GorumsRPCReadQCReply encapsulates the reply from a correctable GorumsRPCReadQC quorum call.
// It contains the id of each node of the quorum that replied and a single reply.
type GorumsRPCReadQCReply struct {
	NodeIDs []uint32
	*Reply
}

func (r GorumsRPCReadQCReply) String() string {
	return fmt.Sprintf("node ids: %v | answer: %v", r.NodeIDs, r.Reply)
}

//TODO Make this a customizable struct that replaces FQRespName together with typedecl option in gogoprotobuf.
//(This file could maybe hold all types of structs for the different call semantics)

// ReadCorrectableReply encapsulates the reply from a correctable ReadCorrectable quorum call.
// It contains the id of each node of the quorum that replied and a single reply.
type ReadCorrectableReply struct {
	NodeIDs []uint32
	*Reply
}

func (r ReadCorrectableReply) String() string {
	return fmt.Sprintf("node ids: %v | answer: %v", r.NodeIDs, r.Reply)
}

//TODO Make this a customizable struct that replaces FQRespName together with typedecl option in gogoprotobuf.
//(This file could maybe hold all types of structs for the different call semantics)

// ReadCorrectablePrelimReply encapsulates the reply from a correctable ReadCorrectablePrelim quorum call.
// It contains the id of each node of the quorum that replied and a single reply.
type ReadCorrectablePrelimReply struct {
	NodeIDs []uint32
	*Reply
}

func (r ReadCorrectablePrelimReply) String() string {
	return fmt.Sprintf("node ids: %v | answer: %v", r.NodeIDs, r.Reply)
}

//TODO Make this a customizable struct that replaces FQRespName together with typedecl option in gogoprotobuf.
//(This file could maybe hold all types of structs for the different call semantics)

// ReadQCCustomReturnReply encapsulates the reply from a correctable ReadQCCustomReturn quorum call.
// It contains the id of each node of the quorum that replied and a single reply.
type ReadQCCustomReturnReply struct {
	NodeIDs []uint32
	*Reply
}

func (r ReadQCCustomReturnReply) String() string {
	return fmt.Sprintf("node ids: %v | answer: %v", r.NodeIDs, r.Reply)
}

//TODO Make this a customizable struct that replaces FQRespName together with typedecl option in gogoprotobuf.
//(This file could maybe hold all types of structs for the different call semantics)

// ReadQCFutureReply encapsulates the reply from a correctable ReadQCFuture quorum call.
// It contains the id of each node of the quorum that replied and a single reply.
type ReadQCFutureReply struct {
	NodeIDs []uint32
	*Reply
}

func (r ReadQCFutureReply) String() string {
	return fmt.Sprintf("node ids: %v | answer: %v", r.NodeIDs, r.Reply)
}

//TODO Make this a customizable struct that replaces FQRespName together with typedecl option in gogoprotobuf.
//(This file could maybe hold all types of structs for the different call semantics)

// WriteQCPerNodeReply encapsulates the reply from a correctable WriteQCPerNode quorum call.
// It contains the id of each node of the quorum that replied and a single reply.
type WriteQCPerNodeReply struct {
	NodeIDs []uint32
	*WriteResp
}

func (r WriteQCPerNodeReply) String() string {
	return fmt.Sprintf("node ids: %v | answer: %v", r.NodeIDs, r.WriteResp)
}

//TODO Make this a customizable struct that replaces FQRespName together with typedecl option in gogoprotobuf.
//(This file could maybe hold all types of structs for the different call semantics)

// WriteQCWithReqReply encapsulates the reply from a correctable WriteQCWithReq quorum call.
// It contains the id of each node of the quorum that replied and a single reply.
type WriteQCWithReqReply struct {
	NodeIDs []uint32
	*WriteResp
}

func (r WriteQCWithReqReply) String() string {
	return fmt.Sprintf("node ids: %v | answer: %v", r.NodeIDs, r.WriteResp)
}
