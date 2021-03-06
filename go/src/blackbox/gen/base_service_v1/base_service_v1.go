// Code generated by Thrift Compiler (0.14.2). DO NOT EDIT.

package base_service_v1

import(
	"bytes"
	"context"
	"fmt"
	"time"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

type BaseService interface {
  GetName(ctx context.Context) (_r string, _err error)
}

type BaseServiceClient struct {
  c thrift.TClient
  meta thrift.ResponseMeta
}

func NewBaseServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *BaseServiceClient {
  return &BaseServiceClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

func NewBaseServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *BaseServiceClient {
  return &BaseServiceClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewBaseServiceClient(c thrift.TClient) *BaseServiceClient {
  return &BaseServiceClient{
    c: c,
  }
}

func (p *BaseServiceClient) Client_() thrift.TClient {
  return p.c
}

func (p *BaseServiceClient) LastResponseMeta_() thrift.ResponseMeta {
  return p.meta
}

func (p *BaseServiceClient) SetLastResponseMeta_(meta thrift.ResponseMeta) {
  p.meta = meta
}

func (p *BaseServiceClient) GetName(ctx context.Context) (_r string, _err error) {
  var _args0 BaseServiceGetNameArgs
  var _result2 BaseServiceGetNameResult
  var _meta1 thrift.ResponseMeta
  _meta1, _err = p.Client_().Call(ctx, "get_name", &_args0, &_result2)
  p.SetLastResponseMeta_(_meta1)
  if _err != nil {
    return
  }
  return _result2.GetSuccess(), nil
}

type BaseServiceProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler BaseService
}

func (p *BaseServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *BaseServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *BaseServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewBaseServiceProcessor(handler BaseService) *BaseServiceProcessor {

  self3 := &BaseServiceProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self3.processorMap["get_name"] = &baseServiceProcessorGetName{handler:handler}
return self3
}

func (p *BaseServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err2 := iprot.ReadMessageBegin(ctx)
  if err2 != nil { return false, thrift.WrapTException(err2) }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(ctx, thrift.STRUCT)
  iprot.ReadMessageEnd(ctx)
  x4 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(ctx, name, thrift.EXCEPTION, seqId)
  x4.Write(ctx, oprot)
  oprot.WriteMessageEnd(ctx)
  oprot.Flush(ctx)
  return false, x4

}

type baseServiceProcessorGetName struct {
  handler BaseService
}

func (p *baseServiceProcessorGetName) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := BaseServiceGetNameArgs{}
  var err2 error
  if err2 = args.Read(ctx, iprot); err2 != nil {
    iprot.ReadMessageEnd(ctx)
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
    oprot.WriteMessageBegin(ctx, "get_name", thrift.EXCEPTION, seqId)
    x.Write(ctx, oprot)
    oprot.WriteMessageEnd(ctx)
    oprot.Flush(ctx)
    return false, thrift.WrapTException(err2)
  }
  iprot.ReadMessageEnd(ctx)

  tickerCancel := func() {}
  // Start a goroutine to do server side connectivity check.
  if thrift.ServerConnectivityCheckInterval > 0 {
    var cancel context.CancelFunc
    ctx, cancel = context.WithCancel(ctx)
    defer cancel()
    var tickerCtx context.Context
    tickerCtx, tickerCancel = context.WithCancel(context.Background())
    defer tickerCancel()
    go func(ctx context.Context, cancel context.CancelFunc) {
      ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
      defer ticker.Stop()
      for {
        select {
        case <-ctx.Done():
          return
        case <-ticker.C:
          if !iprot.Transport().IsOpen() {
            cancel()
            return
          }
        }
      }
    }(tickerCtx, cancel)
  }

  result := BaseServiceGetNameResult{}
  var retval string
  if retval, err2 = p.handler.GetName(ctx); err2 != nil {
    tickerCancel()
    if err2 == thrift.ErrAbandonRequest {
      return false, thrift.WrapTException(err2)
    }
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing get_name: " + err2.Error())
    oprot.WriteMessageBegin(ctx, "get_name", thrift.EXCEPTION, seqId)
    x.Write(ctx, oprot)
    oprot.WriteMessageEnd(ctx)
    oprot.Flush(ctx)
    return true, thrift.WrapTException(err2)
  } else {
    result.Success = &retval
  }
  tickerCancel()
  if err2 = oprot.WriteMessageBegin(ctx, "get_name", thrift.REPLY, seqId); err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err2 = result.Write(ctx, oprot); err == nil && err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err2 = oprot.WriteMessageEnd(ctx); err == nil && err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

type BaseServiceGetNameArgs struct {
}

func NewBaseServiceGetNameArgs() *BaseServiceGetNameArgs {
  return &BaseServiceGetNameArgs{}
}

func (p *BaseServiceGetNameArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(ctx, fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *BaseServiceGetNameArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "get_name_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *BaseServiceGetNameArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("BaseServiceGetNameArgs(%+v)", *p)
}

// Attributes:
//  - Success
type BaseServiceGetNameResult struct {
  Success *string `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewBaseServiceGetNameResult() *BaseServiceGetNameResult {
  return &BaseServiceGetNameResult{}
}

var BaseServiceGetNameResult_Success_DEFAULT string
func (p *BaseServiceGetNameResult) GetSuccess() string {
  if !p.IsSetSuccess() {
    return BaseServiceGetNameResult_Success_DEFAULT
  }
return *p.Success
}
func (p *BaseServiceGetNameResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *BaseServiceGetNameResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField0(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *BaseServiceGetNameResult)  ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *BaseServiceGetNameResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "get_name_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *BaseServiceGetNameResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRING, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteString(ctx, string(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(ctx); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *BaseServiceGetNameResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("BaseServiceGetNameResult(%+v)", *p)
}


