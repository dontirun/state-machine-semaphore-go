//go:build !no_runtime_type_checking

// Create distributed semaphores using AWS Step Functions and Amazon DynamoDB to control concurrent invocations of contentious work.
package dontirunstatemachinesemaphore

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

func (i *jsiiProxy_IChainNextable) validateNextParameters(state awsstepfunctions.IChainable) error {
	if state == nil {
		return fmt.Errorf("parameter state is required, but nil was provided")
	}

	return nil
}

