package location

import (
	"github.com/ruifrodrigues/ecooda/config"
	"reflect"
)

const (
	AddCountryCommand    = "AddCountry"
	AddRegionCommand     = "AddRegion"
	AddChallengeCommand  = "AddChallenge"
	RemoveCountryCommand = "RemoveCountry"
	RemoveRegionCommand  = "RemoveRegion"
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

type AddCountry struct {
	uuid string
	args *Args
}

func NewAddCountryCommand(uuid string, args *Args) Command {
	return &AddCountry{
		uuid: uuid,
		args: args,
	}
}

func (ac *AddCountry) GetUuid() string {
	return ac.uuid
}

func (ac *AddCountry) GetArgs() []Value {
	return ac.args.Get()
}

func (ac *AddCountry) ToString() string {
	return AddCountryCommand
}

type AddRegion struct {
	uuid string
	args *Args
}

func NewAddRegionCommand(uuid string, args *Args) Command {
	return &AddRegion{
		uuid: uuid,
		args: args,
	}
}

func (ar *AddRegion) GetUuid() string {
	return ar.uuid
}

func (ar *AddRegion) GetArgs() []Value {
	return ar.args.Get()
}

func (ar *AddRegion) ToString() string {
	return AddRegionCommand
}

type RemoveCountry struct {
	uuid string
	args *Args
}

func NewRemoveCountryCommand(uuid string, args *Args) Command {
	return &RemoveCountry{
		uuid: uuid,
		args: args,
	}
}

func (rc *RemoveCountry) GetUuid() string {
	return rc.uuid
}

func (rc *RemoveCountry) GetArgs() []Value {
	return rc.args.Get()
}

func (rc *RemoveCountry) ToString() string {
	return RemoveCountryCommand
}

type RemoveRegion struct {
	uuid string
	args *Args
}

func NewRemoveRegionCommand(uuid string, args *Args) Command {
	return &RemoveRegion{
		uuid: uuid,
		args: args,
	}
}

func (rr *RemoveRegion) GetUuid() string {
	return rr.uuid
}

func (rr *RemoveRegion) GetArgs() []Value {
	return rr.args.Get()
}

func (rr *RemoveRegion) ToString() string {
	return RemoveRegionCommand
}

type AddChallenge struct {
	uuid string
	args *Args
}

func NewAddChallengeCommand(uuid string, args *Args) Command {
	return &AddChallenge{
		uuid: uuid,
		args: args,
	}
}

func (ac *AddChallenge) GetUuid() string {
	return ac.uuid
}

func (ac *AddChallenge) GetArgs() []Value {
	return ac.args.Get()
}

func (ac *AddChallenge) ToString() string {
	return AddChallengeCommand
}
