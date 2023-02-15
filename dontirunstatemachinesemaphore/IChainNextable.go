// Create distributed semaphores using AWS Step Functions and Amazon DynamoDB to control concurrent invocations of contentious work.
package dontirunstatemachinesemaphore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/dontirun/state-machine-semaphore-go/dontirunstatemachinesemaphore/internal"
)

type IChainNextable interface {
	awsstepfunctions.IChainable
	awsstepfunctions.INextable
}

// The jsii proxy for IChainNextable
type jsiiProxy_IChainNextable struct {
	internal.Type__awsstepfunctionsIChainable
	internal.Type__awsstepfunctionsINextable
}

func (i *jsiiProxy_IChainNextable) Next(state awsstepfunctions.IChainable) awsstepfunctions.Chain {
	if err := i.validateNextParameters(state); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		i,
		"next",
		[]interface{}{state},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IChainNextable) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChainNextable) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChainNextable) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

