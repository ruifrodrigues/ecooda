package location

import (
	"github.com/ruifrodrigues/ecooda/config"
	"reflect"
)

const (
	AddCountryCommand      = "AddCountry"
	AddRegionCommand       = "AddRegion"
	AddChallengeCommand    = "AddChallenge"
	RemoveCountryCommand   = "RemoveCountry"
	RemoveRegionCommand    = "RemoveRegion"
	RemoveChallengeCommand = "RemoveChallenge"
)

type Value = reflect.Value

type Args struct {
	value []Value
}

func NewCommandArgs() *Args {
	args := new(Args)
	args.value = []Value{}

	return args
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
	handler := new(CommandHandler)
	handler.repository = NewRepository(conf)

	return handler
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
	command := new(AddCountry)
	command.uuid = uuid
	command.args = args

	return command
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
	command := new(AddRegion)
	command.uuid = uuid
	command.args = args

	return command
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
	command := new(RemoveCountry)
	command.uuid = uuid
	command.args = args

	return command
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
	command := new(RemoveRegion)
	command.uuid = uuid
	command.args = args

	return command
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
	command := new(AddChallenge)
	command.uuid = uuid
	command.args = args

	return command
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

type RemoveChallenge struct {
	uuid string
	args *Args
}

func NewRemoveChallengeCommand(uuid string, args *Args) Command {
	command := new(RemoveChallenge)
	command.uuid = uuid
	command.args = args

	return command
}

func (rc *RemoveChallenge) GetUuid() string {
	return rc.uuid
}

func (rc *RemoveChallenge) GetArgs() []Value {
	return rc.args.Get()
}

func (rc *RemoveChallenge) ToString() string {
	return RemoveChallengeCommand
}
