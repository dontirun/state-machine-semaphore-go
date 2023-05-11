package dontirunstatemachinesemaphore


// Read and write capacity for a PROVISIONED billing DynamoDB table.
type TableReadWriteCapacity struct {
	ReadCapacity *float64 `field:"required" json:"readCapacity" yaml:"readCapacity"`
	WriteCapacity *float64 `field:"required" json:"writeCapacity" yaml:"writeCapacity"`
}

