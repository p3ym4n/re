## Rich Error

### Why we need better error handling?

the error messages in go are ugly and very hard to read, also they don't give us much data about what exactly has
happened. for having better observability on errors, we need some more data, this package helps to have those additional
info to with less additional effort.

### How does it work? what's the difference?

almost any error in go that happens in deeper layers, must be passed to its parent to finally generate the appropriate
response to the client.

for example, we may have the error on the adapter layer here, but will not have a clue about _how we have got there_,
_what was the type of the error_, and _what was our runtime arguments_ that may have caused that error.

```
delivery
    |
    ├── interactor
        |
        ├── internal
            |
            ├── adapter <- this is where the error happens
```

to fix this, we will generate the RichError when the error occurs and then keep passing it to the higher level until we
reach the top. and in each level we will add the related data to that layer.

### the RichError features

in RichError, each error can have these:

- **op** the name of the method or function
- **error** the real error that has happened
- **message** any additional message if you want to pass
- **meta_data** any additional runtime argument that you may want to pass
- **kind** which indicates the kind of the error, for example: forbidden, not found, or unexpected
- **code_info** which will automatically add the exact file path and line that error has occurred

### How to Install

 ```asciidoc
$ go get github.com/p3ym4n/re 
```

### How to Use

when ever the error happens you can use RichError like this:

on the deepest layer we will make the RichError

```go
package user_repo

import "github.com/p3ym4n/re"

func (repo *UserRepo) FindUserById(userID uint) (*entity.User, re.Error) {
	const op = re.Op("user_repo.FindUserById") // <- this is the op that we make it as a const on top of each func 
	meta := re.Meta{"user_id": userID}         // <- this is the meta that hold the runtime arguments (optional)

	user, err := repo.handler.FindById(userID)
	if err != nil {
		return nil, re.New(op, err, meta, re.KindNotFound) // <- here we make the RichError
	}
	return user, nil
}

```

and on the higher levels we will just chain it to the next layer

```go
package user_interactor

import "github.com/p3ym4n/re"

func (i *UserInteractor) GetListOfUsers() ([]*entity.User, re.Error) {
	const op = re.Op("user_interactor.GetListOfUsers") // <- we will make the op

	userIds := i.userRepo.GetAllAvaiables()
	users := make([]*entity.User, 0)
	for _, userID := range userIds {
		user, err := i.userRepo.FindUserById(userID)
		if err != nil {
			return nil, err.Chain(op) // <- here we just chain the RichError
		}
		users = append(users, user)
	}
	return users, nil
}

```

other examples for constructing:

```go
const op = re.Op("package_name.func_name")

err := CallingAnything()
if err != nil{
    re.New(op, err) // <- in this case the kind will be Unexpected
}

err2 := CallingAnotherFunc()
if err2 != nil{
    re.New(op, err, KindForbidden, "you are not allowed to do this") // <- you can add additional data
}


```

other examples for chaining:

```go
const op = re.Op("higher_package_name.func_name")

err := callingDeepChildWhoReturnsRichError()
return err.Chain(op)

err := callingAnotherDeepChildWhoReturnsRichError()
return err.ChainWithMeta(op, re.Meta{"arg1":"value1"})

```

for making a decision on how to show the error to the client you can have these attributes from the RichError:

```go
package re

type Error interface {
	Kind() Kind // <- you can add kinds based on your business logics
	Message() string // <- if you have add any messages otherwise it will be ""
	Internal() error // <- this will return ths inner error

	RawMap() map[string]interface{} // <- it will return all the collected values in raw format
	ProcessedMap() map[string]string // <- same as RawMap() but with sanitized values

	Chain(op Op) Error
	ChainWithMeta(op Op, meta Meta) Error
}
```

