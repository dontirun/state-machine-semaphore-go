//go:build no_runtime_type_checking

package dontirunstatemachinesemaphore

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Semaphore) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (s *jsiiProxy_Semaphore) validateToSingleStateParameters(options *awsstepfunctions.SingleStateOptions) error {
	return nil
}

func validateSemaphore_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewSemaphoreParameters(scope constructs.Construct, id *string, props *SemaphoreProps) error {
	return nil
}

