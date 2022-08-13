// Create distributed semaphores using AWS Step Functions and Amazon DynamoDB to control concurrent invocations of contentious work.
package dontirunstatemachinesemaphore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/dontirun/state-machine-semaphore-go/dontirunstatemachinesemaphore/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/dontirun/state-machine-semaphore-go/dontirunstatemachinesemaphore/internal"
)

// Generates a semaphore for a StepFunction job (or chained set of jobs) to limit parallelism across executions.
type Semaphore interface {
	awsstepfunctions.StateMachineFragment
	// The states to chain onto if this fragment is used.
	EndStates() *[]awsstepfunctions.INextable
	// Descriptive identifier for this chainable.
	Id() *string
	// The tree node.
	Node() constructs.Node
	// The start state of this state machine fragment.
	StartState() awsstepfunctions.State
	// Continue normal execution with the given state.
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	// Prefix the IDs of all states in this state machine fragment.
	//
	// Use this to avoid multiple copies of the state machine all having the
	// same state IDs.
	PrefixStates(prefix *string) awsstepfunctions.StateMachineFragment
	// Wrap all states in this state machine fragment up into a single state.
	//
	// This can be used to add retry or error handling onto this state
	// machine fragment.
	//
	// Be aware that this changes the result of the inner state machine
	// to be an array with the result of the state machine in it. Adjust
	// your paths accordingly. For example, change 'outputPath' to
	// '$[0]'.
	ToSingleState(options *awsstepfunctions.SingleStateOptions) awsstepfunctions.Parallel
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Semaphore
type jsiiProxy_Semaphore struct {
	internal.Type__awsstepfunctionsStateMachineFragment
}

func (j *jsiiProxy_Semaphore) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Semaphore) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Semaphore) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Semaphore) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}


func NewSemaphore(scope constructs.Construct, id *string, props *SemaphoreProps) Semaphore {
	_init_.Initialize()

	j := jsiiProxy_Semaphore{}

	_jsii_.Create(
		"@dontirun/state-machine-semaphore.Semaphore",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSemaphore_Override(s Semaphore, scope constructs.Construct, id *string, props *SemaphoreProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@dontirun/state-machine-semaphore.Semaphore",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func Semaphore_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@dontirun/state-machine-semaphore.Semaphore",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Semaphore) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		s,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Semaphore) PrefixStates(prefix *string) awsstepfunctions.StateMachineFragment {
	var returns awsstepfunctions.StateMachineFragment

	_jsii_.Invoke(
		s,
		"prefixStates",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Semaphore) ToSingleState(options *awsstepfunctions.SingleStateOptions) awsstepfunctions.Parallel {
	var returns awsstepfunctions.Parallel

	_jsii_.Invoke(
		s,
		"toSingleState",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Semaphore) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

