package dontirunstatemachinesemaphore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Interface for creating a Semaphore.
type SemaphoreProps struct {
	// The job (or chained jobs) to be semaphored.
	Job IChainNextable `field:"required" json:"job" yaml:"job"`
	// The maximum number of concurrent executions for the given lock.
	Limit *float64 `field:"required" json:"limit" yaml:"limit"`
	// The name of the semaphore.
	LockName *string `field:"required" json:"lockName" yaml:"lockName"`
	// The State to go to after the semaphored job completes.
	NextState awsstepfunctions.State `field:"required" json:"nextState" yaml:"nextState"`
	// Add detailed comments to lock related states.
	//
	// Significantly increases CloudFormation template size. Default: false.
	Comments *bool `field:"optional" json:"comments" yaml:"comments"`
	// Explicility allow the reuse of a named lock from a previously generated job.
	//
	// Throws an error if a different `limit` is specified. Default: false.
	ReuseLock *bool `field:"optional" json:"reuseLock" yaml:"reuseLock"`
	// Optionally set the DynamoDB table to have a specific read/write capacity with PROVISIONED billing.
	//
	// Note: This property can only be set on the first instantiation of a `Semaphore` per stack.
	// Default: PAY_PER_REQUEST.
	//
	TableReadWriteCapacity *TableReadWriteCapacity `field:"optional" json:"tableReadWriteCapacity" yaml:"tableReadWriteCapacity"`
}

