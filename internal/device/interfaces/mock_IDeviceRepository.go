// Code generated by mockery v2.43.2. DO NOT EDIT.

package interfaces

import (
	model "device-manager-service/internal/device/model"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockIDeviceRepository is an autogenerated mock type for the IDeviceRepository type
type MockIDeviceRepository struct {
	mock.Mock
}

type MockIDeviceRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIDeviceRepository) EXPECT() *MockIDeviceRepository_Expecter {
	return &MockIDeviceRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: device
func (_m *MockIDeviceRepository) Create(device *model.Device) error {
	ret := _m.Called(device)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Device) error); ok {
		r0 = rf(device)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIDeviceRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockIDeviceRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - device *model.Device
func (_e *MockIDeviceRepository_Expecter) Create(device interface{}) *MockIDeviceRepository_Create_Call {
	return &MockIDeviceRepository_Create_Call{Call: _e.mock.On("Create", device)}
}

func (_c *MockIDeviceRepository_Create_Call) Run(run func(device *model.Device)) *MockIDeviceRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.Device))
	})
	return _c
}

func (_c *MockIDeviceRepository_Create_Call) Return(_a0 error) *MockIDeviceRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIDeviceRepository_Create_Call) RunAndReturn(run func(*model.Device) error) *MockIDeviceRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: deviceID
func (_m *MockIDeviceRepository) Delete(deviceID uuid.UUID) error {
	ret := _m.Called(deviceID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(deviceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIDeviceRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockIDeviceRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - deviceID uuid.UUID
func (_e *MockIDeviceRepository_Expecter) Delete(deviceID interface{}) *MockIDeviceRepository_Delete_Call {
	return &MockIDeviceRepository_Delete_Call{Call: _e.mock.On("Delete", deviceID)}
}

func (_c *MockIDeviceRepository_Delete_Call) Run(run func(deviceID uuid.UUID)) *MockIDeviceRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID))
	})
	return _c
}

func (_c *MockIDeviceRepository_Delete_Call) Return(_a0 error) *MockIDeviceRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIDeviceRepository_Delete_Call) RunAndReturn(run func(uuid.UUID) error) *MockIDeviceRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: deviceID
func (_m *MockIDeviceRepository) Get(deviceID uuid.UUID) (*model.Device, error) {
	ret := _m.Called(deviceID)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *model.Device
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*model.Device, error)); ok {
		return rf(deviceID)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *model.Device); ok {
		r0 = rf(deviceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Device)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(deviceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIDeviceRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockIDeviceRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - deviceID uuid.UUID
func (_e *MockIDeviceRepository_Expecter) Get(deviceID interface{}) *MockIDeviceRepository_Get_Call {
	return &MockIDeviceRepository_Get_Call{Call: _e.mock.On("Get", deviceID)}
}

func (_c *MockIDeviceRepository_Get_Call) Run(run func(deviceID uuid.UUID)) *MockIDeviceRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID))
	})
	return _c
}

func (_c *MockIDeviceRepository_Get_Call) Return(_a0 *model.Device, _a1 error) *MockIDeviceRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIDeviceRepository_Get_Call) RunAndReturn(run func(uuid.UUID) (*model.Device, error)) *MockIDeviceRepository_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields: limit, offset, brand
func (_m *MockIDeviceRepository) GetAll(limit int, offset int, brand string) ([]*model.Device, error) {
	ret := _m.Called(limit, offset, brand)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []*model.Device
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, string) ([]*model.Device, error)); ok {
		return rf(limit, offset, brand)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []*model.Device); ok {
		r0 = rf(limit, offset, brand)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Device)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) error); ok {
		r1 = rf(limit, offset, brand)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIDeviceRepository_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type MockIDeviceRepository_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
//   - limit int
//   - offset int
//   - brand string
func (_e *MockIDeviceRepository_Expecter) GetAll(limit interface{}, offset interface{}, brand interface{}) *MockIDeviceRepository_GetAll_Call {
	return &MockIDeviceRepository_GetAll_Call{Call: _e.mock.On("GetAll", limit, offset, brand)}
}

func (_c *MockIDeviceRepository_GetAll_Call) Run(run func(limit int, offset int, brand string)) *MockIDeviceRepository_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(int), args[2].(string))
	})
	return _c
}

func (_c *MockIDeviceRepository_GetAll_Call) Return(_a0 []*model.Device, _a1 error) *MockIDeviceRepository_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIDeviceRepository_GetAll_Call) RunAndReturn(run func(int, int, string) ([]*model.Device, error)) *MockIDeviceRepository_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: device
func (_m *MockIDeviceRepository) Update(device *model.Device) error {
	ret := _m.Called(device)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Device) error); ok {
		r0 = rf(device)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIDeviceRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockIDeviceRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - device *model.Device
func (_e *MockIDeviceRepository_Expecter) Update(device interface{}) *MockIDeviceRepository_Update_Call {
	return &MockIDeviceRepository_Update_Call{Call: _e.mock.On("Update", device)}
}

func (_c *MockIDeviceRepository_Update_Call) Run(run func(device *model.Device)) *MockIDeviceRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.Device))
	})
	return _c
}

func (_c *MockIDeviceRepository_Update_Call) Return(_a0 error) *MockIDeviceRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIDeviceRepository_Update_Call) RunAndReturn(run func(*model.Device) error) *MockIDeviceRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateStatus provides a mock function with given fields: deviceID, status
func (_m *MockIDeviceRepository) UpdateStatus(deviceID uuid.UUID, status bool) error {
	ret := _m.Called(deviceID, status)

	if len(ret) == 0 {
		panic("no return value specified for UpdateStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, bool) error); ok {
		r0 = rf(deviceID, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIDeviceRepository_UpdateStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateStatus'
type MockIDeviceRepository_UpdateStatus_Call struct {
	*mock.Call
}

// UpdateStatus is a helper method to define mock.On call
//   - deviceID uuid.UUID
//   - status bool
func (_e *MockIDeviceRepository_Expecter) UpdateStatus(deviceID interface{}, status interface{}) *MockIDeviceRepository_UpdateStatus_Call {
	return &MockIDeviceRepository_UpdateStatus_Call{Call: _e.mock.On("UpdateStatus", deviceID, status)}
}

func (_c *MockIDeviceRepository_UpdateStatus_Call) Run(run func(deviceID uuid.UUID, status bool)) *MockIDeviceRepository_UpdateStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID), args[1].(bool))
	})
	return _c
}

func (_c *MockIDeviceRepository_UpdateStatus_Call) Return(_a0 error) *MockIDeviceRepository_UpdateStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIDeviceRepository_UpdateStatus_Call) RunAndReturn(run func(uuid.UUID, bool) error) *MockIDeviceRepository_UpdateStatus_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIDeviceRepository creates a new instance of MockIDeviceRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIDeviceRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIDeviceRepository {
	mock := &MockIDeviceRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
