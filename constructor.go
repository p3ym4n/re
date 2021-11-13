package re

import (
	"runtime"
	"strings"
)

func New(op Op, internal error, args ...interface{}) *Bag {

	pc, fileName, lineNumber, _ := runtime.Caller(1)
	funcPt := runtime.FuncForPC(pc)
	functionName := "Unknown"
	if funcPt != nil {
		parts := strings.Split(funcPt.Name(), "/")
		secs := strings.SplitN(parts[len(parts)-1], ".", 2)
		functionName = secs[len(secs)-1]
	}

	e := &Bag{
		ops:      []Op{op},
		internal: internal,
		codeInfo: CodeInfo{
			FileName:     fileName,
			FunctionName: functionName,
			LineNumber:   lineNumber,
		},
	}

	for _, arg := range args {
		switch typedArg := arg.(type) {
		case string:
			e.message = typedArg
		case Kind:
			e.kind = typedArg
		case Meta:
			if len(e.metaData) == 0 {
				e.metaData = typedArg
			} else {
				for k, v := range typedArg {
					e.metaData[k] = v
				}
			}
		}
	}

	if e.kind == "" {
		e.kind = KindUnexpected
	}

	return e
}
