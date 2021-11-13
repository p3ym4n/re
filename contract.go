package re

type Error interface {
	AsMap() map[string]interface{}
	Kind() Kind
	Message() string

	Chain(op Op) Error
	ChainWithMeta(op Op, meta Meta) Error
}
