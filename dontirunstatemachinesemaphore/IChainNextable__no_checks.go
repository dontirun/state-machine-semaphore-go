//go:build no_runtime_type_checking

package dontirunstatemachinesemaphore

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IChainNextable) validateNextParameters(state awsstepfunctions.IChainable) error {
	return nil
}

