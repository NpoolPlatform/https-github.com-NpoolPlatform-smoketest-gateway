package detail

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/module"
)

func trace(span trace1.Span, in *npool.ModuleReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("Name.%v", index), in.GetName()),
		attribute.String(fmt.Sprintf("Description.%v", index), in.GetDescription()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.ModuleReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("Name.Op", in.GetName().GetOp()),
		attribute.String("Name.Value", in.GetName().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.ModuleReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
