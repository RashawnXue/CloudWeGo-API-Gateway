// Code generated by thriftgo (0.2.12). DO NOT EDIT.

package service

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type Request struct {
	Operand_1 int32 `thrift:"operand_1,1" frugal:"1,default,i32" json:"operand_1"`
	Operand_2 int32 `thrift:"operand_2,2" frugal:"2,default,i32" json:"operand_2"`
}

func NewRequest() *Request {
	return &Request{}
}

func (p *Request) InitDefault() {
	*p = Request{}
}

func (p *Request) GetOperand_1() (v int32) {
	return p.Operand_1
}

func (p *Request) GetOperand_2() (v int32) {
	return p.Operand_2
}
func (p *Request) SetOperand_1(val int32) {
	p.Operand_1 = val
}
func (p *Request) SetOperand_2(val int32) {
	p.Operand_2 = val
}

var fieldIDToName_Request = map[int16]string{
	1: "operand_1",
	2: "operand_2",
}

func (p *Request) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Request[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *Request) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.Operand_1 = v
	}
	return nil
}

func (p *Request) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.Operand_2 = v
	}
	return nil
}

func (p *Request) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Request"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *Request) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("operand_1", thrift.I32, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.Operand_1); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *Request) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("operand_2", thrift.I32, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.Operand_2); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *Request) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Request(%+v)", *p)
}

func (p *Request) DeepEqual(ano *Request) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Operand_1) {
		return false
	}
	if !p.Field2DeepEqual(ano.Operand_2) {
		return false
	}
	return true
}

func (p *Request) Field1DeepEqual(src int32) bool {

	if p.Operand_1 != src {
		return false
	}
	return true
}
func (p *Request) Field2DeepEqual(src int32) bool {

	if p.Operand_2 != src {
		return false
	}
	return true
}

type Response struct {
	Success bool   `thrift:"success,1" frugal:"1,default,bool" json:"success"`
	Message string `thrift:"message,2" frugal:"2,default,string" json:"message"`
	Data    int32  `thrift:"data,3" frugal:"3,default,i32" json:"data"`
}

func NewResponse() *Response {
	return &Response{}
}

func (p *Response) InitDefault() {
	*p = Response{}
}

func (p *Response) GetSuccess() (v bool) {
	return p.Success
}

func (p *Response) GetMessage() (v string) {
	return p.Message
}

func (p *Response) GetData() (v int32) {
	return p.Data
}
func (p *Response) SetSuccess(val bool) {
	p.Success = val
}
func (p *Response) SetMessage(val string) {
	p.Message = val
}
func (p *Response) SetData(val int32) {
	p.Data = val
}

var fieldIDToName_Response = map[int16]string{
	1: "success",
	2: "message",
	3: "data",
}

func (p *Response) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.BOOL {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Response[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *Response) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return err
	} else {
		p.Success = v
	}
	return nil
}

func (p *Response) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Message = v
	}
	return nil
}

func (p *Response) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.Data = v
	}
	return nil
}

func (p *Response) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Response"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *Response) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("success", thrift.BOOL, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteBool(p.Success); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *Response) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("message", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Message); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *Response) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("data", thrift.I32, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.Data); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *Response) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Response(%+v)", *p)
}

func (p *Response) DeepEqual(ano *Response) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Success) {
		return false
	}
	if !p.Field2DeepEqual(ano.Message) {
		return false
	}
	if !p.Field3DeepEqual(ano.Data) {
		return false
	}
	return true
}

func (p *Response) Field1DeepEqual(src bool) bool {

	if p.Success != src {
		return false
	}
	return true
}
func (p *Response) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Message, src) != 0 {
		return false
	}
	return true
}
func (p *Response) Field3DeepEqual(src int32) bool {

	if p.Data != src {
		return false
	}
	return true
}

type FirstLevelCalService interface {
	Add(ctx context.Context, request *Request) (r *Response, err error)

	Sub(ctx context.Context, request *Request) (r *Response, err error)
}

type FirstLevelCalServiceClient struct {
	c thrift.TClient
}

func NewFirstLevelCalServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *FirstLevelCalServiceClient {
	return &FirstLevelCalServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewFirstLevelCalServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *FirstLevelCalServiceClient {
	return &FirstLevelCalServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewFirstLevelCalServiceClient(c thrift.TClient) *FirstLevelCalServiceClient {
	return &FirstLevelCalServiceClient{
		c: c,
	}
}

func (p *FirstLevelCalServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *FirstLevelCalServiceClient) Add(ctx context.Context, request *Request) (r *Response, err error) {
	var _args FirstLevelCalServiceAddArgs
	_args.Request = request
	var _result FirstLevelCalServiceAddResult
	if err = p.Client_().Call(ctx, "Add", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
func (p *FirstLevelCalServiceClient) Sub(ctx context.Context, request *Request) (r *Response, err error) {
	var _args FirstLevelCalServiceSubArgs
	_args.Request = request
	var _result FirstLevelCalServiceSubResult
	if err = p.Client_().Call(ctx, "Sub", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type FirstLevelCalServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      FirstLevelCalService
}

func (p *FirstLevelCalServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *FirstLevelCalServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *FirstLevelCalServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewFirstLevelCalServiceProcessor(handler FirstLevelCalService) *FirstLevelCalServiceProcessor {
	self := &FirstLevelCalServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("Add", &firstLevelCalServiceProcessorAdd{handler: handler})
	self.AddToProcessorMap("Sub", &firstLevelCalServiceProcessorSub{handler: handler})
	return self
}
func (p *FirstLevelCalServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type firstLevelCalServiceProcessorAdd struct {
	handler FirstLevelCalService
}

func (p *firstLevelCalServiceProcessorAdd) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := FirstLevelCalServiceAddArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Add", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := FirstLevelCalServiceAddResult{}
	var retval *Response
	if retval, err2 = p.handler.Add(ctx, args.Request); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Add: "+err2.Error())
		oprot.WriteMessageBegin("Add", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("Add", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type firstLevelCalServiceProcessorSub struct {
	handler FirstLevelCalService
}

func (p *firstLevelCalServiceProcessorSub) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := FirstLevelCalServiceSubArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Sub", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := FirstLevelCalServiceSubResult{}
	var retval *Response
	if retval, err2 = p.handler.Sub(ctx, args.Request); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Sub: "+err2.Error())
		oprot.WriteMessageBegin("Sub", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("Sub", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type FirstLevelCalServiceAddArgs struct {
	Request *Request `thrift:"request,1" frugal:"1,default,Request" json:"request"`
}

func NewFirstLevelCalServiceAddArgs() *FirstLevelCalServiceAddArgs {
	return &FirstLevelCalServiceAddArgs{}
}

func (p *FirstLevelCalServiceAddArgs) InitDefault() {
	*p = FirstLevelCalServiceAddArgs{}
}

var FirstLevelCalServiceAddArgs_Request_DEFAULT *Request

func (p *FirstLevelCalServiceAddArgs) GetRequest() (v *Request) {
	if !p.IsSetRequest() {
		return FirstLevelCalServiceAddArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *FirstLevelCalServiceAddArgs) SetRequest(val *Request) {
	p.Request = val
}

var fieldIDToName_FirstLevelCalServiceAddArgs = map[int16]string{
	1: "request",
}

func (p *FirstLevelCalServiceAddArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *FirstLevelCalServiceAddArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FirstLevelCalServiceAddArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *FirstLevelCalServiceAddArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Request = NewRequest()
	if err := p.Request.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *FirstLevelCalServiceAddArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Add_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *FirstLevelCalServiceAddArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Request.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *FirstLevelCalServiceAddArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FirstLevelCalServiceAddArgs(%+v)", *p)
}

func (p *FirstLevelCalServiceAddArgs) DeepEqual(ano *FirstLevelCalServiceAddArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Request) {
		return false
	}
	return true
}

func (p *FirstLevelCalServiceAddArgs) Field1DeepEqual(src *Request) bool {

	if !p.Request.DeepEqual(src) {
		return false
	}
	return true
}

type FirstLevelCalServiceAddResult struct {
	Success *Response `thrift:"success,0,optional" frugal:"0,optional,Response" json:"success,omitempty"`
}

func NewFirstLevelCalServiceAddResult() *FirstLevelCalServiceAddResult {
	return &FirstLevelCalServiceAddResult{}
}

func (p *FirstLevelCalServiceAddResult) InitDefault() {
	*p = FirstLevelCalServiceAddResult{}
}

var FirstLevelCalServiceAddResult_Success_DEFAULT *Response

func (p *FirstLevelCalServiceAddResult) GetSuccess() (v *Response) {
	if !p.IsSetSuccess() {
		return FirstLevelCalServiceAddResult_Success_DEFAULT
	}
	return p.Success
}
func (p *FirstLevelCalServiceAddResult) SetSuccess(x interface{}) {
	p.Success = x.(*Response)
}

var fieldIDToName_FirstLevelCalServiceAddResult = map[int16]string{
	0: "success",
}

func (p *FirstLevelCalServiceAddResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FirstLevelCalServiceAddResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FirstLevelCalServiceAddResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *FirstLevelCalServiceAddResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewResponse()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *FirstLevelCalServiceAddResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Add_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *FirstLevelCalServiceAddResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *FirstLevelCalServiceAddResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FirstLevelCalServiceAddResult(%+v)", *p)
}

func (p *FirstLevelCalServiceAddResult) DeepEqual(ano *FirstLevelCalServiceAddResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *FirstLevelCalServiceAddResult) Field0DeepEqual(src *Response) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

type FirstLevelCalServiceSubArgs struct {
	Request *Request `thrift:"request,1" frugal:"1,default,Request" json:"request"`
}

func NewFirstLevelCalServiceSubArgs() *FirstLevelCalServiceSubArgs {
	return &FirstLevelCalServiceSubArgs{}
}

func (p *FirstLevelCalServiceSubArgs) InitDefault() {
	*p = FirstLevelCalServiceSubArgs{}
}

var FirstLevelCalServiceSubArgs_Request_DEFAULT *Request

func (p *FirstLevelCalServiceSubArgs) GetRequest() (v *Request) {
	if !p.IsSetRequest() {
		return FirstLevelCalServiceSubArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *FirstLevelCalServiceSubArgs) SetRequest(val *Request) {
	p.Request = val
}

var fieldIDToName_FirstLevelCalServiceSubArgs = map[int16]string{
	1: "request",
}

func (p *FirstLevelCalServiceSubArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *FirstLevelCalServiceSubArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FirstLevelCalServiceSubArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *FirstLevelCalServiceSubArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Request = NewRequest()
	if err := p.Request.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *FirstLevelCalServiceSubArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Sub_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *FirstLevelCalServiceSubArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Request.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *FirstLevelCalServiceSubArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FirstLevelCalServiceSubArgs(%+v)", *p)
}

func (p *FirstLevelCalServiceSubArgs) DeepEqual(ano *FirstLevelCalServiceSubArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Request) {
		return false
	}
	return true
}

func (p *FirstLevelCalServiceSubArgs) Field1DeepEqual(src *Request) bool {

	if !p.Request.DeepEqual(src) {
		return false
	}
	return true
}

type FirstLevelCalServiceSubResult struct {
	Success *Response `thrift:"success,0,optional" frugal:"0,optional,Response" json:"success,omitempty"`
}

func NewFirstLevelCalServiceSubResult() *FirstLevelCalServiceSubResult {
	return &FirstLevelCalServiceSubResult{}
}

func (p *FirstLevelCalServiceSubResult) InitDefault() {
	*p = FirstLevelCalServiceSubResult{}
}

var FirstLevelCalServiceSubResult_Success_DEFAULT *Response

func (p *FirstLevelCalServiceSubResult) GetSuccess() (v *Response) {
	if !p.IsSetSuccess() {
		return FirstLevelCalServiceSubResult_Success_DEFAULT
	}
	return p.Success
}
func (p *FirstLevelCalServiceSubResult) SetSuccess(x interface{}) {
	p.Success = x.(*Response)
}

var fieldIDToName_FirstLevelCalServiceSubResult = map[int16]string{
	0: "success",
}

func (p *FirstLevelCalServiceSubResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FirstLevelCalServiceSubResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FirstLevelCalServiceSubResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *FirstLevelCalServiceSubResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewResponse()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *FirstLevelCalServiceSubResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Sub_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *FirstLevelCalServiceSubResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *FirstLevelCalServiceSubResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FirstLevelCalServiceSubResult(%+v)", *p)
}

func (p *FirstLevelCalServiceSubResult) DeepEqual(ano *FirstLevelCalServiceSubResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *FirstLevelCalServiceSubResult) Field0DeepEqual(src *Response) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}
