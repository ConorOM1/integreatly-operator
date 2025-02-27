// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package config

import (
	integreatlyv1alpha1 "github.com/integr8ly/integreatly-operator/apis/v1alpha1"
	"sync"
)

// Ensure, that ConfigReadWriterMock does implement ConfigReadWriter.
// If this is not the case, regenerate this file with moq.
var _ ConfigReadWriter = &ConfigReadWriterMock{}

// ConfigReadWriterMock is a mock implementation of ConfigReadWriter.
//
//	func TestSomethingThatUsesConfigReadWriter(t *testing.T) {
//
//		// make and configure a mocked ConfigReadWriter
//		mockedConfigReadWriter := &ConfigReadWriterMock{
//			GetBackupsSecretNameFunc: func() string {
//				panic("mock out the GetBackupsSecretName method")
//			},
//			GetGHOauthClientsSecretNameFunc: func() string {
//				panic("mock out the GetGHOauthClientsSecretName method")
//			},
//			GetOauthClientsSecretNameFunc: func() string {
//				panic("mock out the GetOauthClientsSecretName method")
//			},
//			GetOperatorNamespaceFunc: func() string {
//				panic("mock out the GetOperatorNamespace method")
//			},
//			ReadCloudResourcesFunc: func() (*CloudResources, error) {
//				panic("mock out the ReadCloudResources method")
//			},
//			ReadGrafanaFunc: func() (*Grafana, error) {
//				panic("mock out the ReadGrafana method")
//			},
//			ReadMCGFunc: func() (*MCG, error) {
//				panic("mock out the ReadMCG method")
//			},
//			ReadMarin3rFunc: func() (*Marin3r, error) {
//				panic("mock out the ReadMarin3r method")
//			},
//			ReadProductFunc: func(product integreatlyv1alpha1.ProductName) (ConfigReadable, error) {
//				panic("mock out the ReadProduct method")
//			},
//			ReadRHSSOFunc: func() (*RHSSO, error) {
//				panic("mock out the ReadRHSSO method")
//			},
//			ReadRHSSOUserFunc: func() (*RHSSOUser, error) {
//				panic("mock out the ReadRHSSOUser method")
//			},
//			ReadThreeScaleFunc: func() (*ThreeScale, error) {
//				panic("mock out the ReadThreeScale method")
//			},
//			WriteConfigFunc: func(config ConfigReadable) error {
//				panic("mock out the WriteConfig method")
//			},
//			readConfigForProductFunc: func(product integreatlyv1alpha1.ProductName) (ProductConfig, error) {
//				panic("mock out the readConfigForProduct method")
//			},
//		}
//
//		// use mockedConfigReadWriter in code that requires ConfigReadWriter
//		// and then make assertions.
//
//	}
type ConfigReadWriterMock struct {
	// GetBackupsSecretNameFunc mocks the GetBackupsSecretName method.
	GetBackupsSecretNameFunc func() string

	// GetGHOauthClientsSecretNameFunc mocks the GetGHOauthClientsSecretName method.
	GetGHOauthClientsSecretNameFunc func() string

	// GetOauthClientsSecretNameFunc mocks the GetOauthClientsSecretName method.
	GetOauthClientsSecretNameFunc func() string

	// GetOperatorNamespaceFunc mocks the GetOperatorNamespace method.
	GetOperatorNamespaceFunc func() string

	// ReadCloudResourcesFunc mocks the ReadCloudResources method.
	ReadCloudResourcesFunc func() (*CloudResources, error)

	// ReadGrafanaFunc mocks the ReadGrafana method.
	ReadGrafanaFunc func() (*Grafana, error)

	// ReadMCGFunc mocks the ReadMCG method.
	ReadMCGFunc func() (*MCG, error)

	// ReadMarin3rFunc mocks the ReadMarin3r method.
	ReadMarin3rFunc func() (*Marin3r, error)

	// ReadProductFunc mocks the ReadProduct method.
	ReadProductFunc func(product integreatlyv1alpha1.ProductName) (ConfigReadable, error)

	// ReadRHSSOFunc mocks the ReadRHSSO method.
	ReadRHSSOFunc func() (*RHSSO, error)

	// ReadRHSSOUserFunc mocks the ReadRHSSOUser method.
	ReadRHSSOUserFunc func() (*RHSSOUser, error)

	// ReadThreeScaleFunc mocks the ReadThreeScale method.
	ReadThreeScaleFunc func() (*ThreeScale, error)

	// WriteConfigFunc mocks the WriteConfig method.
	WriteConfigFunc func(config ConfigReadable) error

	// readConfigForProductFunc mocks the readConfigForProduct method.
	readConfigForProductFunc func(product integreatlyv1alpha1.ProductName) (ProductConfig, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetBackupsSecretName holds details about calls to the GetBackupsSecretName method.
		GetBackupsSecretName []struct {
		}
		// GetGHOauthClientsSecretName holds details about calls to the GetGHOauthClientsSecretName method.
		GetGHOauthClientsSecretName []struct {
		}
		// GetOauthClientsSecretName holds details about calls to the GetOauthClientsSecretName method.
		GetOauthClientsSecretName []struct {
		}
		// GetOperatorNamespace holds details about calls to the GetOperatorNamespace method.
		GetOperatorNamespace []struct {
		}
		// ReadCloudResources holds details about calls to the ReadCloudResources method.
		ReadCloudResources []struct {
		}
		// ReadGrafana holds details about calls to the ReadGrafana method.
		ReadGrafana []struct {
		}
		// ReadMCG holds details about calls to the ReadMCG method.
		ReadMCG []struct {
		}
		// ReadMarin3r holds details about calls to the ReadMarin3r method.
		ReadMarin3r []struct {
		}
		// ReadProduct holds details about calls to the ReadProduct method.
		ReadProduct []struct {
			// Product is the product argument value.
			Product integreatlyv1alpha1.ProductName
		}
		// ReadRHSSO holds details about calls to the ReadRHSSO method.
		ReadRHSSO []struct {
		}
		// ReadRHSSOUser holds details about calls to the ReadRHSSOUser method.
		ReadRHSSOUser []struct {
		}
		// ReadThreeScale holds details about calls to the ReadThreeScale method.
		ReadThreeScale []struct {
		}
		// WriteConfig holds details about calls to the WriteConfig method.
		WriteConfig []struct {
			// Config is the config argument value.
			Config ConfigReadable
		}
		// readConfigForProduct holds details about calls to the readConfigForProduct method.
		readConfigForProduct []struct {
			// Product is the product argument value.
			Product integreatlyv1alpha1.ProductName
		}
	}
	lockGetBackupsSecretName        sync.RWMutex
	lockGetGHOauthClientsSecretName sync.RWMutex
	lockGetOauthClientsSecretName   sync.RWMutex
	lockGetOperatorNamespace        sync.RWMutex
	lockReadCloudResources          sync.RWMutex
	lockReadGrafana                 sync.RWMutex
	lockReadMCG                     sync.RWMutex
	lockReadMarin3r                 sync.RWMutex
	lockReadProduct                 sync.RWMutex
	lockReadRHSSO                   sync.RWMutex
	lockReadRHSSOUser               sync.RWMutex
	lockReadThreeScale              sync.RWMutex
	lockWriteConfig                 sync.RWMutex
	lockreadConfigForProduct        sync.RWMutex
}

// GetBackupsSecretName calls GetBackupsSecretNameFunc.
func (mock *ConfigReadWriterMock) GetBackupsSecretName() string {
	if mock.GetBackupsSecretNameFunc == nil {
		panic("ConfigReadWriterMock.GetBackupsSecretNameFunc: method is nil but ConfigReadWriter.GetBackupsSecretName was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetBackupsSecretName.Lock()
	mock.calls.GetBackupsSecretName = append(mock.calls.GetBackupsSecretName, callInfo)
	mock.lockGetBackupsSecretName.Unlock()
	return mock.GetBackupsSecretNameFunc()
}

// GetBackupsSecretNameCalls gets all the calls that were made to GetBackupsSecretName.
// Check the length with:
//
//	len(mockedConfigReadWriter.GetBackupsSecretNameCalls())
func (mock *ConfigReadWriterMock) GetBackupsSecretNameCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetBackupsSecretName.RLock()
	calls = mock.calls.GetBackupsSecretName
	mock.lockGetBackupsSecretName.RUnlock()
	return calls
}

// GetGHOauthClientsSecretName calls GetGHOauthClientsSecretNameFunc.
func (mock *ConfigReadWriterMock) GetGHOauthClientsSecretName() string {
	if mock.GetGHOauthClientsSecretNameFunc == nil {
		panic("ConfigReadWriterMock.GetGHOauthClientsSecretNameFunc: method is nil but ConfigReadWriter.GetGHOauthClientsSecretName was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetGHOauthClientsSecretName.Lock()
	mock.calls.GetGHOauthClientsSecretName = append(mock.calls.GetGHOauthClientsSecretName, callInfo)
	mock.lockGetGHOauthClientsSecretName.Unlock()
	return mock.GetGHOauthClientsSecretNameFunc()
}

// GetGHOauthClientsSecretNameCalls gets all the calls that were made to GetGHOauthClientsSecretName.
// Check the length with:
//
//	len(mockedConfigReadWriter.GetGHOauthClientsSecretNameCalls())
func (mock *ConfigReadWriterMock) GetGHOauthClientsSecretNameCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetGHOauthClientsSecretName.RLock()
	calls = mock.calls.GetGHOauthClientsSecretName
	mock.lockGetGHOauthClientsSecretName.RUnlock()
	return calls
}

// GetOauthClientsSecretName calls GetOauthClientsSecretNameFunc.
func (mock *ConfigReadWriterMock) GetOauthClientsSecretName() string {
	if mock.GetOauthClientsSecretNameFunc == nil {
		panic("ConfigReadWriterMock.GetOauthClientsSecretNameFunc: method is nil but ConfigReadWriter.GetOauthClientsSecretName was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetOauthClientsSecretName.Lock()
	mock.calls.GetOauthClientsSecretName = append(mock.calls.GetOauthClientsSecretName, callInfo)
	mock.lockGetOauthClientsSecretName.Unlock()
	return mock.GetOauthClientsSecretNameFunc()
}

// GetOauthClientsSecretNameCalls gets all the calls that were made to GetOauthClientsSecretName.
// Check the length with:
//
//	len(mockedConfigReadWriter.GetOauthClientsSecretNameCalls())
func (mock *ConfigReadWriterMock) GetOauthClientsSecretNameCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetOauthClientsSecretName.RLock()
	calls = mock.calls.GetOauthClientsSecretName
	mock.lockGetOauthClientsSecretName.RUnlock()
	return calls
}

// GetOperatorNamespace calls GetOperatorNamespaceFunc.
func (mock *ConfigReadWriterMock) GetOperatorNamespace() string {
	if mock.GetOperatorNamespaceFunc == nil {
		panic("ConfigReadWriterMock.GetOperatorNamespaceFunc: method is nil but ConfigReadWriter.GetOperatorNamespace was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetOperatorNamespace.Lock()
	mock.calls.GetOperatorNamespace = append(mock.calls.GetOperatorNamespace, callInfo)
	mock.lockGetOperatorNamespace.Unlock()
	return mock.GetOperatorNamespaceFunc()
}

// GetOperatorNamespaceCalls gets all the calls that were made to GetOperatorNamespace.
// Check the length with:
//
//	len(mockedConfigReadWriter.GetOperatorNamespaceCalls())
func (mock *ConfigReadWriterMock) GetOperatorNamespaceCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetOperatorNamespace.RLock()
	calls = mock.calls.GetOperatorNamespace
	mock.lockGetOperatorNamespace.RUnlock()
	return calls
}

// ReadCloudResources calls ReadCloudResourcesFunc.
func (mock *ConfigReadWriterMock) ReadCloudResources() (*CloudResources, error) {
	if mock.ReadCloudResourcesFunc == nil {
		panic("ConfigReadWriterMock.ReadCloudResourcesFunc: method is nil but ConfigReadWriter.ReadCloudResources was just called")
	}
	callInfo := struct {
	}{}
	mock.lockReadCloudResources.Lock()
	mock.calls.ReadCloudResources = append(mock.calls.ReadCloudResources, callInfo)
	mock.lockReadCloudResources.Unlock()
	return mock.ReadCloudResourcesFunc()
}

// ReadCloudResourcesCalls gets all the calls that were made to ReadCloudResources.
// Check the length with:
//
//	len(mockedConfigReadWriter.ReadCloudResourcesCalls())
func (mock *ConfigReadWriterMock) ReadCloudResourcesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockReadCloudResources.RLock()
	calls = mock.calls.ReadCloudResources
	mock.lockReadCloudResources.RUnlock()
	return calls
}

// ReadGrafana calls ReadGrafanaFunc.
func (mock *ConfigReadWriterMock) ReadGrafana() (*Grafana, error) {
	if mock.ReadGrafanaFunc == nil {
		panic("ConfigReadWriterMock.ReadGrafanaFunc: method is nil but ConfigReadWriter.ReadGrafana was just called")
	}
	callInfo := struct {
	}{}
	mock.lockReadGrafana.Lock()
	mock.calls.ReadGrafana = append(mock.calls.ReadGrafana, callInfo)
	mock.lockReadGrafana.Unlock()
	return mock.ReadGrafanaFunc()
}

// ReadGrafanaCalls gets all the calls that were made to ReadGrafana.
// Check the length with:
//
//	len(mockedConfigReadWriter.ReadGrafanaCalls())
func (mock *ConfigReadWriterMock) ReadGrafanaCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockReadGrafana.RLock()
	calls = mock.calls.ReadGrafana
	mock.lockReadGrafana.RUnlock()
	return calls
}

// ReadMCG calls ReadMCGFunc.
func (mock *ConfigReadWriterMock) ReadMCG() (*MCG, error) {
	if mock.ReadMCGFunc == nil {
		panic("ConfigReadWriterMock.ReadMCGFunc: method is nil but ConfigReadWriter.ReadMCG was just called")
	}
	callInfo := struct {
	}{}
	mock.lockReadMCG.Lock()
	mock.calls.ReadMCG = append(mock.calls.ReadMCG, callInfo)
	mock.lockReadMCG.Unlock()
	return mock.ReadMCGFunc()
}

// ReadMCGCalls gets all the calls that were made to ReadMCG.
// Check the length with:
//
//	len(mockedConfigReadWriter.ReadMCGCalls())
func (mock *ConfigReadWriterMock) ReadMCGCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockReadMCG.RLock()
	calls = mock.calls.ReadMCG
	mock.lockReadMCG.RUnlock()
	return calls
}

// ReadMarin3r calls ReadMarin3rFunc.
func (mock *ConfigReadWriterMock) ReadMarin3r() (*Marin3r, error) {
	if mock.ReadMarin3rFunc == nil {
		panic("ConfigReadWriterMock.ReadMarin3rFunc: method is nil but ConfigReadWriter.ReadMarin3r was just called")
	}
	callInfo := struct {
	}{}
	mock.lockReadMarin3r.Lock()
	mock.calls.ReadMarin3r = append(mock.calls.ReadMarin3r, callInfo)
	mock.lockReadMarin3r.Unlock()
	return mock.ReadMarin3rFunc()
}

// ReadMarin3rCalls gets all the calls that were made to ReadMarin3r.
// Check the length with:
//
//	len(mockedConfigReadWriter.ReadMarin3rCalls())
func (mock *ConfigReadWriterMock) ReadMarin3rCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockReadMarin3r.RLock()
	calls = mock.calls.ReadMarin3r
	mock.lockReadMarin3r.RUnlock()
	return calls
}

// ReadProduct calls ReadProductFunc.
func (mock *ConfigReadWriterMock) ReadProduct(product integreatlyv1alpha1.ProductName) (ConfigReadable, error) {
	if mock.ReadProductFunc == nil {
		panic("ConfigReadWriterMock.ReadProductFunc: method is nil but ConfigReadWriter.ReadProduct was just called")
	}
	callInfo := struct {
		Product integreatlyv1alpha1.ProductName
	}{
		Product: product,
	}
	mock.lockReadProduct.Lock()
	mock.calls.ReadProduct = append(mock.calls.ReadProduct, callInfo)
	mock.lockReadProduct.Unlock()
	return mock.ReadProductFunc(product)
}

// ReadProductCalls gets all the calls that were made to ReadProduct.
// Check the length with:
//
//	len(mockedConfigReadWriter.ReadProductCalls())
func (mock *ConfigReadWriterMock) ReadProductCalls() []struct {
	Product integreatlyv1alpha1.ProductName
} {
	var calls []struct {
		Product integreatlyv1alpha1.ProductName
	}
	mock.lockReadProduct.RLock()
	calls = mock.calls.ReadProduct
	mock.lockReadProduct.RUnlock()
	return calls
}

// ReadRHSSO calls ReadRHSSOFunc.
func (mock *ConfigReadWriterMock) ReadRHSSO() (*RHSSO, error) {
	if mock.ReadRHSSOFunc == nil {
		panic("ConfigReadWriterMock.ReadRHSSOFunc: method is nil but ConfigReadWriter.ReadRHSSO was just called")
	}
	callInfo := struct {
	}{}
	mock.lockReadRHSSO.Lock()
	mock.calls.ReadRHSSO = append(mock.calls.ReadRHSSO, callInfo)
	mock.lockReadRHSSO.Unlock()
	return mock.ReadRHSSOFunc()
}

// ReadRHSSOCalls gets all the calls that were made to ReadRHSSO.
// Check the length with:
//
//	len(mockedConfigReadWriter.ReadRHSSOCalls())
func (mock *ConfigReadWriterMock) ReadRHSSOCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockReadRHSSO.RLock()
	calls = mock.calls.ReadRHSSO
	mock.lockReadRHSSO.RUnlock()
	return calls
}

// ReadRHSSOUser calls ReadRHSSOUserFunc.
func (mock *ConfigReadWriterMock) ReadRHSSOUser() (*RHSSOUser, error) {
	if mock.ReadRHSSOUserFunc == nil {
		panic("ConfigReadWriterMock.ReadRHSSOUserFunc: method is nil but ConfigReadWriter.ReadRHSSOUser was just called")
	}
	callInfo := struct {
	}{}
	mock.lockReadRHSSOUser.Lock()
	mock.calls.ReadRHSSOUser = append(mock.calls.ReadRHSSOUser, callInfo)
	mock.lockReadRHSSOUser.Unlock()
	return mock.ReadRHSSOUserFunc()
}

// ReadRHSSOUserCalls gets all the calls that were made to ReadRHSSOUser.
// Check the length with:
//
//	len(mockedConfigReadWriter.ReadRHSSOUserCalls())
func (mock *ConfigReadWriterMock) ReadRHSSOUserCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockReadRHSSOUser.RLock()
	calls = mock.calls.ReadRHSSOUser
	mock.lockReadRHSSOUser.RUnlock()
	return calls
}

// ReadThreeScale calls ReadThreeScaleFunc.
func (mock *ConfigReadWriterMock) ReadThreeScale() (*ThreeScale, error) {
	if mock.ReadThreeScaleFunc == nil {
		panic("ConfigReadWriterMock.ReadThreeScaleFunc: method is nil but ConfigReadWriter.ReadThreeScale was just called")
	}
	callInfo := struct {
	}{}
	mock.lockReadThreeScale.Lock()
	mock.calls.ReadThreeScale = append(mock.calls.ReadThreeScale, callInfo)
	mock.lockReadThreeScale.Unlock()
	return mock.ReadThreeScaleFunc()
}

// ReadThreeScaleCalls gets all the calls that were made to ReadThreeScale.
// Check the length with:
//
//	len(mockedConfigReadWriter.ReadThreeScaleCalls())
func (mock *ConfigReadWriterMock) ReadThreeScaleCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockReadThreeScale.RLock()
	calls = mock.calls.ReadThreeScale
	mock.lockReadThreeScale.RUnlock()
	return calls
}

// WriteConfig calls WriteConfigFunc.
func (mock *ConfigReadWriterMock) WriteConfig(config ConfigReadable) error {
	if mock.WriteConfigFunc == nil {
		panic("ConfigReadWriterMock.WriteConfigFunc: method is nil but ConfigReadWriter.WriteConfig was just called")
	}
	callInfo := struct {
		Config ConfigReadable
	}{
		Config: config,
	}
	mock.lockWriteConfig.Lock()
	mock.calls.WriteConfig = append(mock.calls.WriteConfig, callInfo)
	mock.lockWriteConfig.Unlock()
	return mock.WriteConfigFunc(config)
}

// WriteConfigCalls gets all the calls that were made to WriteConfig.
// Check the length with:
//
//	len(mockedConfigReadWriter.WriteConfigCalls())
func (mock *ConfigReadWriterMock) WriteConfigCalls() []struct {
	Config ConfigReadable
} {
	var calls []struct {
		Config ConfigReadable
	}
	mock.lockWriteConfig.RLock()
	calls = mock.calls.WriteConfig
	mock.lockWriteConfig.RUnlock()
	return calls
}

// readConfigForProduct calls readConfigForProductFunc.
func (mock *ConfigReadWriterMock) readConfigForProduct(product integreatlyv1alpha1.ProductName) (ProductConfig, error) {
	if mock.readConfigForProductFunc == nil {
		panic("ConfigReadWriterMock.readConfigForProductFunc: method is nil but ConfigReadWriter.readConfigForProduct was just called")
	}
	callInfo := struct {
		Product integreatlyv1alpha1.ProductName
	}{
		Product: product,
	}
	mock.lockreadConfigForProduct.Lock()
	mock.calls.readConfigForProduct = append(mock.calls.readConfigForProduct, callInfo)
	mock.lockreadConfigForProduct.Unlock()
	return mock.readConfigForProductFunc(product)
}

// readConfigForProductCalls gets all the calls that were made to readConfigForProduct.
// Check the length with:
//
//	len(mockedConfigReadWriter.readConfigForProductCalls())
func (mock *ConfigReadWriterMock) readConfigForProductCalls() []struct {
	Product integreatlyv1alpha1.ProductName
} {
	var calls []struct {
		Product integreatlyv1alpha1.ProductName
	}
	mock.lockreadConfigForProduct.RLock()
	calls = mock.calls.readConfigForProduct
	mock.lockreadConfigForProduct.RUnlock()
	return calls
}
