package dontirunstatemachinesemaphore

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterInterface(
		"@dontirun/state-machine-semaphore.IChainNextable",
		reflect.TypeOf((*IChainNextable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
		},
		func() interface{} {
			j := jsiiProxy_IChainNextable{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsIChainable)
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsINextable)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@dontirun/state-machine-semaphore.Semaphore",
		reflect.TypeOf((*Semaphore)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "prefixStates", GoMethod: "PrefixStates"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberMethod{JsiiMethod: "toSingleState", GoMethod: "ToSingleState"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Semaphore{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsStateMachineFragment)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@dontirun/state-machine-semaphore.SemaphoreProps",
		reflect.TypeOf((*SemaphoreProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@dontirun/state-machine-semaphore.TableReadWriteCapacity",
		reflect.TypeOf((*TableReadWriteCapacity)(nil)).Elem(),
	)
}
