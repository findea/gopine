package gomicrotrace

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/server"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"goweb/pkg/lighttracer"
	"goweb/pkg/lighttracer/tags"
	"strings"
)

// https://github.com/micro/go-plugins/blob/master/wrapper/trace/opentracing/opentracing.go

type clientWrapper struct {
	tracer *lighttracer.Tracer
	client.Client
}

func (o *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	name := fmt.Sprintf("%s.%s", req.Service(), req.Endpoint())
	ctx, span, err := StartSpanFromContext(ctx, o.tracer, name, ext.SpanKindRPCClient)
	if err != nil {
		return err
	}
	defer span.Finish()

	if bs, err := json.Marshal(req.Body()); err == nil {
		tags.RPCRequestBody.Set(span, string(bs))
	}

	if err = o.Client.Call(ctx, req, rsp, opts...); err != nil {
		tags.Error.Set(span, tags.DefaultErrorNo, err.Error())
	} else {
		if bs, err := json.Marshal(rsp); err == nil {
			tags.RPCResponseBody.Set(span, string(bs))
		}
	}
	return err
}

func (o *clientWrapper) Stream(ctx context.Context, req client.Request, opts ...client.CallOption) (client.Stream, error) {
	name := fmt.Sprintf("%s.%s", req.Service(), req.Endpoint())
	ctx, span, err := StartSpanFromContext(ctx, o.tracer, name, ext.SpanKindRPCClient)
	if err != nil {
		return nil, err
	}
	defer span.Finish()
	stream, err := o.Client.Stream(ctx, req, opts...)
	if err != nil {
		tags.Error.Set(span, tags.DefaultErrorNo, err.Error())
	}
	return stream, err
}

func (o *clientWrapper) Publish(ctx context.Context, p client.Message, opts ...client.PublishOption) error {
	name := fmt.Sprintf("Pub to %s", p.Topic())
	ctx, span, err := StartSpanFromContext(ctx, o.tracer, name, ext.SpanKindRPCClient)
	if err != nil {
		return err
	}
	defer span.Finish()
	if err = o.Client.Publish(ctx, p, opts...); err != nil {
		tags.Error.Set(span, tags.DefaultErrorNo, err.Error())
	}
	return err
}

// NewClientWrapper accepts an open tracing Trace and returns a Client Wrapper
func NewClientWrapper(tracer *lighttracer.Tracer) client.Wrapper {
	return func(c client.Client) client.Client {
		return &clientWrapper{tracer, c}
	}
}

// NewCallWrapper accepts an opentracing Tracer and returns a Call Wrapper
func NewCallWrapper(tracer *lighttracer.Tracer) client.CallWrapper {
	return func(cf client.CallFunc) client.CallFunc {
		return func(ctx context.Context, node *registry.Node, req client.Request, rsp interface{}, opts client.CallOptions) error {
			name := fmt.Sprintf("%s.%s", req.Service(), req.Endpoint())
			ctx, span, err := StartSpanFromContext(ctx, tracer, name, ext.SpanKindRPCServer)
			if err != nil {
				return err
			}
			defer span.Finish()
			if err = cf(ctx, node, req, rsp, opts); err != nil {
				tags.Error.Set(span, tags.DefaultErrorNo, err.Error())
			}
			return err
		}
	}
}

// NewSubscriberWrapper accepts an opentracing Tracer and returns a Subscriber Wrapper
func NewSubscriberWrapper(tracer *lighttracer.Tracer) server.SubscriberWrapper {
	return func(next server.SubscriberFunc) server.SubscriberFunc {
		return func(ctx context.Context, msg server.Message) error {
			name := "Sub from " + msg.Topic()
			ctx, span, err := StartSpanFromContext(ctx, tracer, name, ext.SpanKindRPCServer)
			if err != nil {
				return err
			}
			defer span.Finish()
			if err = next(ctx, msg); err != nil {
				tags.Error.Set(span, tags.DefaultErrorNo, err.Error())
			}
			return err
		}
	}
}

func NewHandlerWrapper(tracer *lighttracer.Tracer) server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			name := fmt.Sprintf("%s.%s", req.Service(), req.Endpoint())
			_, span, err := StartSpanFromContext(ctx, tracer, name, ext.SpanKindRPCServer)
			if err != nil {
				return err
			}
			defer span.Finish()

			if bs, err := json.Marshal(req.Body()); err == nil {
				tags.RPCRequestBody.Set(span, string(bs))
			}

			err = h(ctx, req, rsp)

			if err != nil {
				tags.Error.Set(span, tags.DefaultErrorNo, err.Error())
			} else {
				if bs, err := json.Marshal(rsp); err == nil {
					tags.RPCResponseBody.Set(span, string(bs))
				}
			}
			return err
		}
	}
}

// StartSpanFromContext returns a new span with the given operation name and options. If a span
// is found in the context, it will be used as the parent of the resulting span.
func StartSpanFromContext(ctx context.Context, tracer *lighttracer.Tracer, name string, opts ...opentracing.StartSpanOption) (context.Context, opentracing.Span, error) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(metadata.Metadata)
	}

	// Find parent span.
	// First try to get span within current service boundary.
	// If there doesn't exist, try to get it from go-micro metadata(which is cross boundary)
	if parentSpan := lighttracer.SpanFromContext(ctx); parentSpan != nil {
		opts = append(opts, opentracing.ChildOf(parentSpan.Context()))
	} else if spanCtx, err := tracer.Extract(opentracing.TextMap, opentracing.TextMapCarrier(md)); err == nil {
		opts = append(opts, opentracing.ChildOf(spanCtx))
	}

	// allocate new map with only one element
	nmd := make(metadata.Metadata, 1)

	sp := tracer.StartSpanWithType(name, lighttracer.OperationTypeRPC, opts...)
	ext.Component.Set(sp, "gomicro")

	if err := sp.Tracer().Inject(sp.Context(), opentracing.TextMap, opentracing.TextMapCarrier(nmd)); err != nil {
		return nil, nil, err
	}

	for k, v := range nmd {
		md.Set(strings.Title(k), v)
	}

	ctx = lighttracer.ContextWithSpan(ctx, sp)
	ctx = metadata.NewContext(ctx, md)
	return ctx, sp, nil
}
