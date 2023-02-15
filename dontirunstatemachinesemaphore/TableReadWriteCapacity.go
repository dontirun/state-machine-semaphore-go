// Create distributed semaphores using AWS Step Functions and Amazon DynamoDB to control concurrent invocations of contentious work.
package dontirunstatemachinesemaphore


// Read and write capacity for a PROVISIONED billing DynamoDB table.
type TableReadWriteCapacity struct {
	ReadCapacity *float64 `field:"required" json:"readCapacity" yaml:"readCapacity"`
	WriteCapacity *float64 `field:"required" json:"writeCapacity" yaml:"writeCapacity"`
}

