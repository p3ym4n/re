package re

import "strings"

type Bag struct {
	ops      []Op
	codeInfo CodeInfo
	message  string
	internal error
	kind     Kind
	metaData Meta
}

func (e *Bag) Chain(op Op) Error {
	e.ops = append(e.ops, op)
	return e
}

func (e *Bag) ChainWithMeta(op Op, meta Meta) Error {
	e.ops = append(e.ops, op)
	e.metaData[op.String()] = meta
	return e
}

func (e *Bag) RawMap() map[string]interface{} {
	return map[string]interface{}{
		"operations": e.ops,
		"internal":   e.internal,
		"kind":       e.kind,
		"code_info":  e.codeInfo,
		"meta_data":  e.metaData,
		"message":    e.message,
	}
}

func (e *Bag) ProcessedMap() map[string]string {

	operations := make([]string, len(e.ops))
	for i := len(e.ops) - 1; i >= 0; i-- {
		operations = append(operations, e.ops[i].String())
	}

	return map[string]string{
		"operations": strings.Join(operations, " => "),
		"internal":   e.internal.Error(),
		"kind":       e.kind.String(),
		"code_info":  e.codeInfo.String(),
		"meta_data":  e.metaData.String(),
	}
}

func (e *Bag) Kind() Kind {
	return e.kind
}

func (e *Bag) Message() string {
	return e.message
}
