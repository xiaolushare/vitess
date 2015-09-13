// Code generated by protoc-gen-go.
// source: binlogdata.proto
// DO NOT EDIT!

/*
Package binlogdata is a generated protocol buffer package.

It is generated from these files:
	binlogdata.proto

It has these top-level messages:
	Charset
	BinlogTransaction
	StreamEvent
	StreamUpdateRequest
	StreamUpdateResponse
	StreamKeyRangeRequest
	StreamKeyRangeResponse
	StreamTablesRequest
	StreamTablesResponse
*/
package binlogdata

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import query "github.com/youtube/vitess/go/vt/proto/query"
import topodata "github.com/youtube/vitess/go/vt/proto/topodata"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BinlogTransaction_Statement_Category int32

const (
	BinlogTransaction_Statement_BL_UNRECOGNIZED BinlogTransaction_Statement_Category = 0
	BinlogTransaction_Statement_BL_BEGIN        BinlogTransaction_Statement_Category = 1
	BinlogTransaction_Statement_BL_COMMIT       BinlogTransaction_Statement_Category = 2
	BinlogTransaction_Statement_BL_ROLLBACK     BinlogTransaction_Statement_Category = 3
	BinlogTransaction_Statement_BL_DML          BinlogTransaction_Statement_Category = 4
	BinlogTransaction_Statement_BL_DDL          BinlogTransaction_Statement_Category = 5
	BinlogTransaction_Statement_BL_SET          BinlogTransaction_Statement_Category = 6
)

var BinlogTransaction_Statement_Category_name = map[int32]string{
	0: "BL_UNRECOGNIZED",
	1: "BL_BEGIN",
	2: "BL_COMMIT",
	3: "BL_ROLLBACK",
	4: "BL_DML",
	5: "BL_DDL",
	6: "BL_SET",
}
var BinlogTransaction_Statement_Category_value = map[string]int32{
	"BL_UNRECOGNIZED": 0,
	"BL_BEGIN":        1,
	"BL_COMMIT":       2,
	"BL_ROLLBACK":     3,
	"BL_DML":          4,
	"BL_DDL":          5,
	"BL_SET":          6,
}

func (x BinlogTransaction_Statement_Category) String() string {
	return proto.EnumName(BinlogTransaction_Statement_Category_name, int32(x))
}

// the category of this event
type StreamEvent_Category int32

const (
	StreamEvent_SE_ERR StreamEvent_Category = 0
	StreamEvent_SE_DML StreamEvent_Category = 1
	StreamEvent_SE_DDL StreamEvent_Category = 2
	StreamEvent_SE_POS StreamEvent_Category = 3
)

var StreamEvent_Category_name = map[int32]string{
	0: "SE_ERR",
	1: "SE_DML",
	2: "SE_DDL",
	3: "SE_POS",
}
var StreamEvent_Category_value = map[string]int32{
	"SE_ERR": 0,
	"SE_DML": 1,
	"SE_DDL": 2,
	"SE_POS": 3,
}

func (x StreamEvent_Category) String() string {
	return proto.EnumName(StreamEvent_Category_name, int32(x))
}

// Charset is the per-statement charset info from a QUERY_EVENT binlog entry.
type Charset struct {
	// @@session.character_set_client
	Client int32 `protobuf:"varint,1,opt,name=client" json:"client,omitempty"`
	// @@session.collation_connection
	Conn int32 `protobuf:"varint,2,opt,name=conn" json:"conn,omitempty"`
	// @@session.collation_server
	Server int32 `protobuf:"varint,3,opt,name=server" json:"server,omitempty"`
}

func (m *Charset) Reset()         { *m = Charset{} }
func (m *Charset) String() string { return proto.CompactTextString(m) }
func (*Charset) ProtoMessage()    {}

// BinlogTransaction describes a transaction inside the binlogs.
type BinlogTransaction struct {
	// the statements in this transaction
	Statements []*BinlogTransaction_Statement `protobuf:"bytes,1,rep,name=statements" json:"statements,omitempty"`
	// the timestamp of the statements
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	// the Transaction ID after this statement was applied
	TransactionId string `protobuf:"bytes,3,opt,name=transaction_id" json:"transaction_id,omitempty"`
}

func (m *BinlogTransaction) Reset()         { *m = BinlogTransaction{} }
func (m *BinlogTransaction) String() string { return proto.CompactTextString(m) }
func (*BinlogTransaction) ProtoMessage()    {}

func (m *BinlogTransaction) GetStatements() []*BinlogTransaction_Statement {
	if m != nil {
		return m.Statements
	}
	return nil
}

type BinlogTransaction_Statement struct {
	// what type of statement is this?
	Category BinlogTransaction_Statement_Category `protobuf:"varint,1,opt,name=category,enum=binlogdata.BinlogTransaction_Statement_Category" json:"category,omitempty"`
	// charset of this statement, if different from pre-negotiated default.
	Charset *Charset `protobuf:"bytes,2,opt,name=charset" json:"charset,omitempty"`
	// the sql
	Sql string `protobuf:"bytes,3,opt,name=sql" json:"sql,omitempty"`
}

func (m *BinlogTransaction_Statement) Reset()         { *m = BinlogTransaction_Statement{} }
func (m *BinlogTransaction_Statement) String() string { return proto.CompactTextString(m) }
func (*BinlogTransaction_Statement) ProtoMessage()    {}

func (m *BinlogTransaction_Statement) GetCharset() *Charset {
	if m != nil {
		return m.Charset
	}
	return nil
}

// StreamEvent describes an update stream event inside the binlogs.
type StreamEvent struct {
	Category StreamEvent_Category `protobuf:"varint,1,opt,name=category,enum=binlogdata.StreamEvent_Category" json:"category,omitempty"`
	// table_name, primary_key_fields and primary_key_values are set for SE_DML
	TableName        string         `protobuf:"bytes,2,opt,name=table_name" json:"table_name,omitempty"`
	PrimaryKeyFields []*query.Field `protobuf:"bytes,3,rep,name=primary_key_fields" json:"primary_key_fields,omitempty"`
	PrimaryKeyValues []*query.Row   `protobuf:"bytes,4,rep,name=primary_key_values" json:"primary_key_values,omitempty"`
	// sql is set for SE_DDL or SE_ERR
	Sql string `protobuf:"bytes,5,opt,name=sql" json:"sql,omitempty"`
	// timestamp is set for SE_DML, SE_DDL or SE_ERR
	Timestamp int64 `protobuf:"varint,6,opt,name=timestamp" json:"timestamp,omitempty"`
	// the Transaction ID after this statement was applied
	TransactionId string `protobuf:"bytes,7,opt,name=transaction_id" json:"transaction_id,omitempty"`
}

func (m *StreamEvent) Reset()         { *m = StreamEvent{} }
func (m *StreamEvent) String() string { return proto.CompactTextString(m) }
func (*StreamEvent) ProtoMessage()    {}

func (m *StreamEvent) GetPrimaryKeyFields() []*query.Field {
	if m != nil {
		return m.PrimaryKeyFields
	}
	return nil
}

func (m *StreamEvent) GetPrimaryKeyValues() []*query.Row {
	if m != nil {
		return m.PrimaryKeyValues
	}
	return nil
}

// StreamUpdateRequest is the payload to StreamUpdate
type StreamUpdateRequest struct {
	// where to start
	Position string `protobuf:"bytes,1,opt,name=position" json:"position,omitempty"`
}

func (m *StreamUpdateRequest) Reset()         { *m = StreamUpdateRequest{} }
func (m *StreamUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*StreamUpdateRequest) ProtoMessage()    {}

// StreamUpdateResponse is the response from StreamUpdate
type StreamUpdateResponse struct {
	StreamEvent *StreamEvent `protobuf:"bytes,1,opt,name=stream_event" json:"stream_event,omitempty"`
}

func (m *StreamUpdateResponse) Reset()         { *m = StreamUpdateResponse{} }
func (m *StreamUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*StreamUpdateResponse) ProtoMessage()    {}

func (m *StreamUpdateResponse) GetStreamEvent() *StreamEvent {
	if m != nil {
		return m.StreamEvent
	}
	return nil
}

// StreamKeyRangeRequest is the payload to StreamKeyRange
type StreamKeyRangeRequest struct {
	// where to start
	Position string `protobuf:"bytes,1,opt,name=position" json:"position,omitempty"`
	// type to get
	KeyspaceIdType topodata.KeyspaceIdType `protobuf:"varint,2,opt,name=keyspace_id_type,enum=topodata.KeyspaceIdType" json:"keyspace_id_type,omitempty"`
	// what to get
	KeyRange *topodata.KeyRange `protobuf:"bytes,3,opt,name=key_range" json:"key_range,omitempty"`
	// default charset on the player side
	Charset *Charset `protobuf:"bytes,4,opt,name=charset" json:"charset,omitempty"`
}

func (m *StreamKeyRangeRequest) Reset()         { *m = StreamKeyRangeRequest{} }
func (m *StreamKeyRangeRequest) String() string { return proto.CompactTextString(m) }
func (*StreamKeyRangeRequest) ProtoMessage()    {}

func (m *StreamKeyRangeRequest) GetKeyRange() *topodata.KeyRange {
	if m != nil {
		return m.KeyRange
	}
	return nil
}

func (m *StreamKeyRangeRequest) GetCharset() *Charset {
	if m != nil {
		return m.Charset
	}
	return nil
}

// StreamKeyRangeResponse is the response from StreamKeyRange
type StreamKeyRangeResponse struct {
	BinlogTransaction *BinlogTransaction `protobuf:"bytes,1,opt,name=binlog_transaction" json:"binlog_transaction,omitempty"`
}

func (m *StreamKeyRangeResponse) Reset()         { *m = StreamKeyRangeResponse{} }
func (m *StreamKeyRangeResponse) String() string { return proto.CompactTextString(m) }
func (*StreamKeyRangeResponse) ProtoMessage()    {}

func (m *StreamKeyRangeResponse) GetBinlogTransaction() *BinlogTransaction {
	if m != nil {
		return m.BinlogTransaction
	}
	return nil
}

// StreamTablesRequest is the payload to StreamTables
type StreamTablesRequest struct {
	// where to start
	Position string `protobuf:"bytes,1,opt,name=position" json:"position,omitempty"`
	// what to get
	Tables []string `protobuf:"bytes,2,rep,name=tables" json:"tables,omitempty"`
	// default charset on the player side
	Charset *Charset `protobuf:"bytes,3,opt,name=charset" json:"charset,omitempty"`
}

func (m *StreamTablesRequest) Reset()         { *m = StreamTablesRequest{} }
func (m *StreamTablesRequest) String() string { return proto.CompactTextString(m) }
func (*StreamTablesRequest) ProtoMessage()    {}

func (m *StreamTablesRequest) GetCharset() *Charset {
	if m != nil {
		return m.Charset
	}
	return nil
}

// StreamTablesResponse is the response from StreamTables
type StreamTablesResponse struct {
	BinlogTransaction *BinlogTransaction `protobuf:"bytes,1,opt,name=binlog_transaction" json:"binlog_transaction,omitempty"`
}

func (m *StreamTablesResponse) Reset()         { *m = StreamTablesResponse{} }
func (m *StreamTablesResponse) String() string { return proto.CompactTextString(m) }
func (*StreamTablesResponse) ProtoMessage()    {}

func (m *StreamTablesResponse) GetBinlogTransaction() *BinlogTransaction {
	if m != nil {
		return m.BinlogTransaction
	}
	return nil
}

func init() {
	proto.RegisterEnum("binlogdata.BinlogTransaction_Statement_Category", BinlogTransaction_Statement_Category_name, BinlogTransaction_Statement_Category_value)
	proto.RegisterEnum("binlogdata.StreamEvent_Category", StreamEvent_Category_name, StreamEvent_Category_value)
}
