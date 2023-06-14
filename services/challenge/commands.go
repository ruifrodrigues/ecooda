package challenge

import (
	"github.com/ruifrodrigues/ecooda/config"
	"reflect"
)

const (
	AddCategoryCommand    = "AddCategory"
	RemoveCategoryCommand = "RemoveCategory"
)

type Value = reflect.Value

type Args struct {
	value []Value
}

func NewCommandArgs() *Args {
	return &Args{
		value: []Value{},
	}
}

func (arg *Args) Add(value interface{}) *Args {
	arg.value = append(arg.value, reflect.ValueOf(value))

	return arg
}

func (arg *Args) Get() []Value {
	return arg.value
}

type Handler interface {
	Handle(cmd Command) error
}

type CommandHandler struct {
	repository Behaviour
}

func NewCommandHandler(conf config.Config) Handler {
	return &CommandHandler{
		repository: NewRepository(conf),
	}
}

func (ch *CommandHandler) Handle(cmd Command) error {
	aggregate, err := ch.repository.Get(cmd.GetUuid())
	if err != nil {
		return err
	}

	result := reflect.ValueOf(aggregate).
		MethodByName(cmd.ToString()).
		Call(cmd.GetArgs())

	if err, ok := result[0].Interface().(error); ok {
		return err
	}

	if err := ch.repository.Save(); err != nil {
		return err
	}

	return nil
}

type Command interface {
	GetUuid() string
	ToString() string
	GetArgs() []Value
}

type AddCategory struct {
	uuid string
	args *Args
}

func NewAddCategoryCommand(uuid string, args *Args) Command {
	return &AddCategory{
		uuid: uuid,
		args: args,
	}
}

func (ac *AddCategory) GetUuid() string {
	return ac.uuid
}

func (ac *AddCategory) GetArgs() []Value {
	return ac.args.Get()
}

func (ac *AddCategory) ToString() string {
	return AddCategoryCommand
}

type RemoveCategory struct {
	uuid string
	args *Args
}

func NewRemoveCategoryCommand(uuid string, args *Args) Command {
	return &RemoveCategory{
		uuid: uuid,
		args: args,
	}
}

func (rc *RemoveCategory) GetUuid() string {
	return rc.uuid
}

func (rc *RemoveCategory) GetArgs() []Value {
	return rc.args.Get()
}

func (rc *RemoveCategory) ToString() string {
	return RemoveCategoryCommand
}
