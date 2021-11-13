package re

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

func (e *Bag) AsMap() map[string]interface{} {
	return map[string]interface{}{
		"operations": e.ops,
		"internal":   e.internal,
		"kind":       e.kind,
		"code_info":  e.codeInfo,
		"meta_data":  e.metaData,
		"message":    e.message,
	}
}

func (e *Bag) Kind() Kind {
	return e.kind
}

func (e *Bag) Message() string {
	return e.message
}