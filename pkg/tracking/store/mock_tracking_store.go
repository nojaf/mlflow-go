// Code generated by mockery v2.43.2. DO NOT EDIT.

package store

import (
	context "context"

	contract "github.com/mlflow/mlflow-go/pkg/contract"
	entities "github.com/mlflow/mlflow-go/pkg/entities"

	mock "github.com/stretchr/testify/mock"

	protos "github.com/mlflow/mlflow-go/pkg/protos"
)

// MockTrackingStore is an autogenerated mock type for the TrackingStore type
type MockTrackingStore struct {
	mock.Mock
}

type MockTrackingStore_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTrackingStore) EXPECT() *MockTrackingStore_Expecter {
	return &MockTrackingStore_Expecter{mock: &_m.Mock}
}

// CreateExperiment provides a mock function with given fields: ctx, name, artifactLocation, tags
func (_m *MockTrackingStore) CreateExperiment(ctx context.Context, name string, artifactLocation string, tags []*entities.ExperimentTag) (string, *contract.Error) {
	ret := _m.Called(ctx, name, artifactLocation, tags)

	if len(ret) == 0 {
		panic("no return value specified for CreateExperiment")
	}

	var r0 string
	var r1 *contract.Error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []*entities.ExperimentTag) (string, *contract.Error)); ok {
		return rf(ctx, name, artifactLocation, tags)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []*entities.ExperimentTag) string); ok {
		r0 = rf(ctx, name, artifactLocation, tags)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, []*entities.ExperimentTag) *contract.Error); ok {
		r1 = rf(ctx, name, artifactLocation, tags)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*contract.Error)
		}
	}

	return r0, r1
}

// MockTrackingStore_CreateExperiment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateExperiment'
type MockTrackingStore_CreateExperiment_Call struct {
	*mock.Call
}

// CreateExperiment is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - artifactLocation string
//   - tags []*entities.ExperimentTag
func (_e *MockTrackingStore_Expecter) CreateExperiment(ctx interface{}, name interface{}, artifactLocation interface{}, tags interface{}) *MockTrackingStore_CreateExperiment_Call {
	return &MockTrackingStore_CreateExperiment_Call{Call: _e.mock.On("CreateExperiment", ctx, name, artifactLocation, tags)}
}

func (_c *MockTrackingStore_CreateExperiment_Call) Run(run func(ctx context.Context, name string, artifactLocation string, tags []*entities.ExperimentTag)) *MockTrackingStore_CreateExperiment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].([]*entities.ExperimentTag))
	})
	return _c
}

func (_c *MockTrackingStore_CreateExperiment_Call) Return(_a0 string, _a1 *contract.Error) *MockTrackingStore_CreateExperiment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTrackingStore_CreateExperiment_Call) RunAndReturn(run func(context.Context, string, string, []*entities.ExperimentTag) (string, *contract.Error)) *MockTrackingStore_CreateExperiment_Call {
	_c.Call.Return(run)
	return _c
}

// CreateRun provides a mock function with given fields: ctx, experimentID, userID, startTime, tags, runName
func (_m *MockTrackingStore) CreateRun(ctx context.Context, experimentID string, userID string, startTime int64, tags []*entities.RunTag, runName string) (*entities.Run, *contract.Error) {
	ret := _m.Called(ctx, experimentID, userID, startTime, tags, runName)

	if len(ret) == 0 {
		panic("no return value specified for CreateRun")
	}

	var r0 *entities.Run
	var r1 *contract.Error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64, []*entities.RunTag, string) (*entities.Run, *contract.Error)); ok {
		return rf(ctx, experimentID, userID, startTime, tags, runName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64, []*entities.RunTag, string) *entities.Run); ok {
		r0 = rf(ctx, experimentID, userID, startTime, tags, runName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Run)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, int64, []*entities.RunTag, string) *contract.Error); ok {
		r1 = rf(ctx, experimentID, userID, startTime, tags, runName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*contract.Error)
		}
	}

	return r0, r1
}

// MockTrackingStore_CreateRun_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateRun'
type MockTrackingStore_CreateRun_Call struct {
	*mock.Call
}

// CreateRun is a helper method to define mock.On call
//   - ctx context.Context
//   - experimentID string
//   - userID string
//   - startTime int64
//   - tags []*entities.RunTag
//   - runName string
func (_e *MockTrackingStore_Expecter) CreateRun(ctx interface{}, experimentID interface{}, userID interface{}, startTime interface{}, tags interface{}, runName interface{}) *MockTrackingStore_CreateRun_Call {
	return &MockTrackingStore_CreateRun_Call{Call: _e.mock.On("CreateRun", ctx, experimentID, userID, startTime, tags, runName)}
}

func (_c *MockTrackingStore_CreateRun_Call) Run(run func(ctx context.Context, experimentID string, userID string, startTime int64, tags []*entities.RunTag, runName string)) *MockTrackingStore_CreateRun_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(int64), args[4].([]*entities.RunTag), args[5].(string))
	})
	return _c
}

func (_c *MockTrackingStore_CreateRun_Call) Return(_a0 *entities.Run, _a1 *contract.Error) *MockTrackingStore_CreateRun_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTrackingStore_CreateRun_Call) RunAndReturn(run func(context.Context, string, string, int64, []*entities.RunTag, string) (*entities.Run, *contract.Error)) *MockTrackingStore_CreateRun_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteExperiment provides a mock function with given fields: ctx, id
func (_m *MockTrackingStore) DeleteExperiment(ctx context.Context, id string) *contract.Error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteExperiment")
	}

	var r0 *contract.Error
	if rf, ok := ret.Get(0).(func(context.Context, string) *contract.Error); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contract.Error)
		}
	}

	return r0
}

// MockTrackingStore_DeleteExperiment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteExperiment'
type MockTrackingStore_DeleteExperiment_Call struct {
	*mock.Call
}

// DeleteExperiment is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockTrackingStore_Expecter) DeleteExperiment(ctx interface{}, id interface{}) *MockTrackingStore_DeleteExperiment_Call {
	return &MockTrackingStore_DeleteExperiment_Call{Call: _e.mock.On("DeleteExperiment", ctx, id)}
}

func (_c *MockTrackingStore_DeleteExperiment_Call) Run(run func(ctx context.Context, id string)) *MockTrackingStore_DeleteExperiment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockTrackingStore_DeleteExperiment_Call) Return(_a0 *contract.Error) *MockTrackingStore_DeleteExperiment_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTrackingStore_DeleteExperiment_Call) RunAndReturn(run func(context.Context, string) *contract.Error) *MockTrackingStore_DeleteExperiment_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteRun provides a mock function with given fields: ctx, runID
func (_m *MockTrackingStore) DeleteRun(ctx context.Context, runID string) *contract.Error {
	ret := _m.Called(ctx, runID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteRun")
	}

	var r0 *contract.Error
	if rf, ok := ret.Get(0).(func(context.Context, string) *contract.Error); ok {
		r0 = rf(ctx, runID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contract.Error)
		}
	}

	return r0
}

// MockTrackingStore_DeleteRun_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteRun'
type MockTrackingStore_DeleteRun_Call struct {
	*mock.Call
}

// DeleteRun is a helper method to define mock.On call
//   - ctx context.Context
//   - runID string
func (_e *MockTrackingStore_Expecter) DeleteRun(ctx interface{}, runID interface{}) *MockTrackingStore_DeleteRun_Call {
	return &MockTrackingStore_DeleteRun_Call{Call: _e.mock.On("DeleteRun", ctx, runID)}
}

func (_c *MockTrackingStore_DeleteRun_Call) Run(run func(ctx context.Context, runID string)) *MockTrackingStore_DeleteRun_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockTrackingStore_DeleteRun_Call) Return(_a0 *contract.Error) *MockTrackingStore_DeleteRun_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTrackingStore_DeleteRun_Call) RunAndReturn(run func(context.Context, string) *contract.Error) *MockTrackingStore_DeleteRun_Call {
	_c.Call.Return(run)
	return _c
}

// GetExperiment provides a mock function with given fields: ctx, id.
func (_m *MockTrackingStore) GetExperiment(ctx context.Context, id string) (*entities.Experiment, *contract.Error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetExperiment")
	}

	var r0 *entities.Experiment
	var r1 *contract.Error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entities.Experiment, *contract.Error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entities.Experiment); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Experiment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) *contract.Error); ok {
		r1 = rf(ctx, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*contract.Error)
		}
	}

	return r0, r1
}

// MockTrackingStore_GetExperiment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExperiment'
type MockTrackingStore_GetExperiment_Call struct {
	*mock.Call
}

// GetExperiment is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockTrackingStore_Expecter) GetExperiment(ctx interface{}, id interface{}) *MockTrackingStore_GetExperiment_Call {
	return &MockTrackingStore_GetExperiment_Call{Call: _e.mock.On("GetExperiment", ctx, id)}
}

func (_c *MockTrackingStore_GetExperiment_Call) Run(run func(ctx context.Context, id string)) *MockTrackingStore_GetExperiment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockTrackingStore_GetExperiment_Call) Return(_a0 *entities.Experiment, _a1 *contract.Error) *MockTrackingStore_GetExperiment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTrackingStore_GetExperiment_Call) RunAndReturn(run func(context.Context, string) (*entities.Experiment, *contract.Error)) *MockTrackingStore_GetExperiment_Call {
	_c.Call.Return(run)
	return _c
}

// GetExperimentByName provides a mock function with given fields: ctx, name
func (_m *MockTrackingStore) GetExperimentByName(ctx context.Context, name string) (*entities.Experiment, *contract.Error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetExperimentByName")
	}

	var r0 *entities.Experiment
	var r1 *contract.Error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entities.Experiment, *contract.Error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entities.Experiment); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Experiment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) *contract.Error); ok {
		r1 = rf(ctx, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*contract.Error)
		}
	}

	return r0, r1
}

// MockTrackingStore_GetExperimentByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExperimentByName'
type MockTrackingStore_GetExperimentByName_Call struct {
	*mock.Call
}

// GetExperimentByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockTrackingStore_Expecter) GetExperimentByName(ctx interface{}, name interface{}) *MockTrackingStore_GetExperimentByName_Call {
	return &MockTrackingStore_GetExperimentByName_Call{Call: _e.mock.On("GetExperimentByName", ctx, name)}
}

func (_c *MockTrackingStore_GetExperimentByName_Call) Run(run func(ctx context.Context, name string)) *MockTrackingStore_GetExperimentByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockTrackingStore_GetExperimentByName_Call) Return(_a0 *entities.Experiment, _a1 *contract.Error) *MockTrackingStore_GetExperimentByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTrackingStore_GetExperimentByName_Call) RunAndReturn(run func(context.Context, string) (*entities.Experiment, *contract.Error)) *MockTrackingStore_GetExperimentByName_Call {
	_c.Call.Return(run)
	return _c
}

// GetRun provides a mock function with given fields: ctx, runID
func (_m *MockTrackingStore) GetRun(ctx context.Context, runID string) (*entities.Run, *contract.Error) {
	ret := _m.Called(ctx, runID)

	if len(ret) == 0 {
		panic("no return value specified for GetRun")
	}

	var r0 *entities.Run
	var r1 *contract.Error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entities.Run, *contract.Error)); ok {
		return rf(ctx, runID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entities.Run); ok {
		r0 = rf(ctx, runID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Run)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) *contract.Error); ok {
		r1 = rf(ctx, runID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*contract.Error)
		}
	}

	return r0, r1
}

// MockTrackingStore_GetRun_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRun'
type MockTrackingStore_GetRun_Call struct {
	*mock.Call
}

// GetRun is a helper method to define mock.On call
//   - ctx context.Context
//   - runID string
func (_e *MockTrackingStore_Expecter) GetRun(ctx interface{}, runID interface{}) *MockTrackingStore_GetRun_Call {
	return &MockTrackingStore_GetRun_Call{Call: _e.mock.On("GetRun", ctx, runID)}
}

func (_c *MockTrackingStore_GetRun_Call) Run(run func(ctx context.Context, runID string)) *MockTrackingStore_GetRun_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockTrackingStore_GetRun_Call) Return(_a0 *entities.Run, _a1 *contract.Error) *MockTrackingStore_GetRun_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTrackingStore_GetRun_Call) RunAndReturn(run func(context.Context, string) (*entities.Run, *contract.Error)) *MockTrackingStore_GetRun_Call {
	_c.Call.Return(run)
	return _c
}

// GetRunTag provides a mock function with given fields: ctx, runID, tagKey
func (_m *MockTrackingStore) GetRunTag(ctx context.Context, runID string, tagKey string) (*entities.RunTag, *contract.Error) {
	ret := _m.Called(ctx, runID, tagKey)

	if len(ret) == 0 {
		panic("no return value specified for GetRunTag")
	}

	var r0 *entities.RunTag
	var r1 *contract.Error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*entities.RunTag, *contract.Error)); ok {
		return rf(ctx, runID, tagKey)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *entities.RunTag); ok {
		r0 = rf(ctx, runID, tagKey)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.RunTag)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) *contract.Error); ok {
		r1 = rf(ctx, runID, tagKey)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*contract.Error)
		}
	}

	return r0, r1
}

// MockTrackingStore_GetRunTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRunTag'
type MockTrackingStore_GetRunTag_Call struct {
	*mock.Call
}

// GetRunTag is a helper method to define mock.On call
//   - ctx context.Context
//   - runID string
//   - tagKey string
func (_e *MockTrackingStore_Expecter) GetRunTag(ctx interface{}, runID interface{}, tagKey interface{}) *MockTrackingStore_GetRunTag_Call {
	return &MockTrackingStore_GetRunTag_Call{Call: _e.mock.On("GetRunTag", ctx, runID, tagKey)}
}

func (_c *MockTrackingStore_GetRunTag_Call) Run(run func(ctx context.Context, runID string, tagKey string)) *MockTrackingStore_GetRunTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockTrackingStore_GetRunTag_Call) Return(_a0 *entities.RunTag, _a1 *contract.Error) *MockTrackingStore_GetRunTag_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTrackingStore_GetRunTag_Call) RunAndReturn(run func(context.Context, string, string) (*entities.RunTag, *contract.Error)) *MockTrackingStore_GetRunTag_Call {
	_c.Call.Return(run)
	return _c
}

// LogBatch provides a mock function with given fields: ctx, runID, metrics, params, tags
func (_m *MockTrackingStore) LogBatch(ctx context.Context, runID string, metrics []*entities.Metric, params []*entities.Param, tags []*entities.RunTag) *contract.Error {
	ret := _m.Called(ctx, runID, metrics, params, tags)

	if len(ret) == 0 {
		panic("no return value specified for LogBatch")
	}

	var r0 *contract.Error
	if rf, ok := ret.Get(0).(func(context.Context, string, []*entities.Metric, []*entities.Param, []*entities.RunTag) *contract.Error); ok {
		r0 = rf(ctx, runID, metrics, params, tags)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contract.Error)
		}
	}

	return r0
}

// MockTrackingStore_LogBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LogBatch'
type MockTrackingStore_LogBatch_Call struct {
	*mock.Call
}

// LogBatch is a helper method to define mock.On call
//   - ctx context.Context
//   - runID string
//   - metrics []*entities.Metric
//   - params []*entities.Param
//   - tags []*entities.RunTag
func (_e *MockTrackingStore_Expecter) LogBatch(ctx interface{}, runID interface{}, metrics interface{}, params interface{}, tags interface{}) *MockTrackingStore_LogBatch_Call {
	return &MockTrackingStore_LogBatch_Call{Call: _e.mock.On("LogBatch", ctx, runID, metrics, params, tags)}
}

func (_c *MockTrackingStore_LogBatch_Call) Run(run func(ctx context.Context, runID string, metrics []*entities.Metric, params []*entities.Param, tags []*entities.RunTag)) *MockTrackingStore_LogBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]*entities.Metric), args[3].([]*entities.Param), args[4].([]*entities.RunTag))
	})
	return _c
}

func (_c *MockTrackingStore_LogBatch_Call) Return(_a0 *contract.Error) *MockTrackingStore_LogBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTrackingStore_LogBatch_Call) RunAndReturn(run func(context.Context, string, []*entities.Metric, []*entities.Param, []*entities.RunTag) *contract.Error) *MockTrackingStore_LogBatch_Call {
	_c.Call.Return(run)
	return _c
}

// LogMetric provides a mock function with given fields: ctx, runID, metric
func (_m *MockTrackingStore) LogMetric(ctx context.Context, runID string, metric *entities.Metric) *contract.Error {
	ret := _m.Called(ctx, runID, metric)

	if len(ret) == 0 {
		panic("no return value specified for LogMetric")
	}

	var r0 *contract.Error
	if rf, ok := ret.Get(0).(func(context.Context, string, *entities.Metric) *contract.Error); ok {
		r0 = rf(ctx, runID, metric)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contract.Error)
		}
	}

	return r0
}

// MockTrackingStore_LogMetric_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LogMetric'
type MockTrackingStore_LogMetric_Call struct {
	*mock.Call
}

// LogMetric is a helper method to define mock.On call
//   - ctx context.Context
//   - runID string
//   - metric *entities.Metric
func (_e *MockTrackingStore_Expecter) LogMetric(ctx interface{}, runID interface{}, metric interface{}) *MockTrackingStore_LogMetric_Call {
	return &MockTrackingStore_LogMetric_Call{Call: _e.mock.On("LogMetric", ctx, runID, metric)}
}

func (_c *MockTrackingStore_LogMetric_Call) Run(run func(ctx context.Context, runID string, metric *entities.Metric)) *MockTrackingStore_LogMetric_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*entities.Metric))
	})
	return _c
}

func (_c *MockTrackingStore_LogMetric_Call) Return(_a0 *contract.Error) *MockTrackingStore_LogMetric_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTrackingStore_LogMetric_Call) RunAndReturn(run func(context.Context, string, *entities.Metric) *contract.Error) *MockTrackingStore_LogMetric_Call {
	_c.Call.Return(run)
	return _c
}

// RenameExperiment provides a mock function with given fields: ctx, experimentID, name.
func (_m *MockTrackingStore) RenameExperiment(ctx context.Context, experimentID, name string) *contract.Error {
	ret := _m.Called(ctx, experimentID, name)

	if len(ret) == 0 {
		panic("no return value specified for RenameExperiment")
	}

	var r0 *contract.Error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *contract.Error); ok {
		r0 = rf(ctx, experimentID, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contract.Error)
		}
	}

	return r0
}

// MockTrackingStore_RenameExperiment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RenameExperiment'
type MockTrackingStore_RenameExperiment_Call struct {
	*mock.Call
}

// RenameExperiment is a helper method to define mock.On call
//   - ctx context.Context
//   - experimentID string
//   - name string
func (_e *MockTrackingStore_Expecter) RenameExperiment(ctx interface{}, experimentID interface{}, name interface{}) *MockTrackingStore_RenameExperiment_Call {
	return &MockTrackingStore_RenameExperiment_Call{Call: _e.mock.On("RenameExperiment", ctx, experimentID, name)}
}

func (_c *MockTrackingStore_RenameExperiment_Call) Run(run func(ctx context.Context, experimentID string, name string)) *MockTrackingStore_RenameExperiment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockTrackingStore_RenameExperiment_Call) Return(_a0 *contract.Error) *MockTrackingStore_RenameExperiment_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTrackingStore_RenameExperiment_Call) RunAndReturn(run func(context.Context, string, string) *contract.Error) *MockTrackingStore_RenameExperiment_Call {
	_c.Call.Return(run)
	return _c
}

// RestoreExperiment provides a mock function with given fields: ctx, id
func (_m *MockTrackingStore) RestoreExperiment(ctx context.Context, id string) *contract.Error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for RestoreExperiment")
	}

	var r0 *contract.Error
	if rf, ok := ret.Get(0).(func(context.Context, string) *contract.Error); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contract.Error)
		}
	}

	return r0
}

// MockTrackingStore_RestoreExperiment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RestoreExperiment'
type MockTrackingStore_RestoreExperiment_Call struct {
	*mock.Call
}

// RestoreExperiment is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockTrackingStore_Expecter) RestoreExperiment(ctx interface{}, id interface{}) *MockTrackingStore_RestoreExperiment_Call {
	return &MockTrackingStore_RestoreExperiment_Call{Call: _e.mock.On("RestoreExperiment", ctx, id)}
}

func (_c *MockTrackingStore_RestoreExperiment_Call) Run(run func(ctx context.Context, id string)) *MockTrackingStore_RestoreExperiment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockTrackingStore_RestoreExperiment_Call) Return(_a0 *contract.Error) *MockTrackingStore_RestoreExperiment_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTrackingStore_RestoreExperiment_Call) RunAndReturn(run func(context.Context, string) *contract.Error) *MockTrackingStore_RestoreExperiment_Call {
	_c.Call.Return(run)
	return _c
}

// RestoreRun provides a mock function with given fields: ctx, runID
func (_m *MockTrackingStore) RestoreRun(ctx context.Context, runID string) *contract.Error {
	ret := _m.Called(ctx, runID)

	if len(ret) == 0 {
		panic("no return value specified for RestoreRun")
	}

	var r0 *contract.Error
	if rf, ok := ret.Get(0).(func(context.Context, string) *contract.Error); ok {
		r0 = rf(ctx, runID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contract.Error)
		}
	}

	return r0
}

// MockTrackingStore_RestoreRun_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RestoreRun'
type MockTrackingStore_RestoreRun_Call struct {
	*mock.Call
}

// RestoreRun is a helper method to define mock.On call
//   - ctx context.Context
//   - runID string
func (_e *MockTrackingStore_Expecter) RestoreRun(ctx interface{}, runID interface{}) *MockTrackingStore_RestoreRun_Call {
	return &MockTrackingStore_RestoreRun_Call{Call: _e.mock.On("RestoreRun", ctx, runID)}
}

func (_c *MockTrackingStore_RestoreRun_Call) Run(run func(ctx context.Context, runID string)) *MockTrackingStore_RestoreRun_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockTrackingStore_RestoreRun_Call) Return(_a0 *contract.Error) *MockTrackingStore_RestoreRun_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTrackingStore_RestoreRun_Call) RunAndReturn(run func(context.Context, string) *contract.Error) *MockTrackingStore_RestoreRun_Call {
	_c.Call.Return(run)
	return _c
}

// SearchRuns provides a mock function with given fields: ctx, experimentIDs, filter, runViewType, maxResults, orderBy, pageToken
func (_m *MockTrackingStore) SearchRuns(ctx context.Context, experimentIDs []string, filter string, runViewType protos.ViewType, maxResults int, orderBy []string, pageToken string) ([]*entities.Run, string, *contract.Error) {
	ret := _m.Called(ctx, experimentIDs, filter, runViewType, maxResults, orderBy, pageToken)

	if len(ret) == 0 {
		panic("no return value specified for SearchRuns")
	}

	var r0 []*entities.Run
	var r1 string
	var r2 *contract.Error
	if rf, ok := ret.Get(0).(func(context.Context, []string, string, protos.ViewType, int, []string, string) ([]*entities.Run, string, *contract.Error)); ok {
		return rf(ctx, experimentIDs, filter, runViewType, maxResults, orderBy, pageToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string, string, protos.ViewType, int, []string, string) []*entities.Run); ok {
		r0 = rf(ctx, experimentIDs, filter, runViewType, maxResults, orderBy, pageToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.Run)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string, string, protos.ViewType, int, []string, string) string); ok {
		r1 = rf(ctx, experimentIDs, filter, runViewType, maxResults, orderBy, pageToken)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, []string, string, protos.ViewType, int, []string, string) *contract.Error); ok {
		r2 = rf(ctx, experimentIDs, filter, runViewType, maxResults, orderBy, pageToken)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*contract.Error)
		}
	}

	return r0, r1, r2
}

// MockTrackingStore_SearchRuns_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SearchRuns'
type MockTrackingStore_SearchRuns_Call struct {
	*mock.Call
}

// SearchRuns is a helper method to define mock.On call
//   - ctx context.Context
//   - experimentIDs []string
//   - filter string
//   - runViewType protos.ViewType
//   - maxResults int
//   - orderBy []string
//   - pageToken string
func (_e *MockTrackingStore_Expecter) SearchRuns(ctx interface{}, experimentIDs interface{}, filter interface{}, runViewType interface{}, maxResults interface{}, orderBy interface{}, pageToken interface{}) *MockTrackingStore_SearchRuns_Call {
	return &MockTrackingStore_SearchRuns_Call{Call: _e.mock.On("SearchRuns", ctx, experimentIDs, filter, runViewType, maxResults, orderBy, pageToken)}
}

func (_c *MockTrackingStore_SearchRuns_Call) Run(run func(ctx context.Context, experimentIDs []string, filter string, runViewType protos.ViewType, maxResults int, orderBy []string, pageToken string)) *MockTrackingStore_SearchRuns_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string), args[2].(string), args[3].(protos.ViewType), args[4].(int), args[5].([]string), args[6].(string))
	})
	return _c
}

func (_c *MockTrackingStore_SearchRuns_Call) Return(_a0 []*entities.Run, _a1 string, _a2 *contract.Error) *MockTrackingStore_SearchRuns_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockTrackingStore_SearchRuns_Call) RunAndReturn(run func(context.Context, []string, string, protos.ViewType, int, []string, string) ([]*entities.Run, string, *contract.Error)) *MockTrackingStore_SearchRuns_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateRun provides a mock function with given fields: ctx, runID, runStatus, endTime, runName
func (_m *MockTrackingStore) UpdateRun(ctx context.Context, runID string, runStatus string, endTime *int64, runName string) *contract.Error {
	ret := _m.Called(ctx, runID, runStatus, endTime, runName)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRun")
	}

	var r0 *contract.Error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *int64, string) *contract.Error); ok {
		r0 = rf(ctx, runID, runStatus, endTime, runName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contract.Error)
		}
	}

	return r0
}

// MockTrackingStore_UpdateRun_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateRun'
type MockTrackingStore_UpdateRun_Call struct {
	*mock.Call
}

// UpdateRun is a helper method to define mock.On call
//   - ctx context.Context
//   - runID string
//   - runStatus string
//   - endTime *int64
//   - runName string
func (_e *MockTrackingStore_Expecter) UpdateRun(ctx interface{}, runID interface{}, runStatus interface{}, endTime interface{}, runName interface{}) *MockTrackingStore_UpdateRun_Call {
	return &MockTrackingStore_UpdateRun_Call{Call: _e.mock.On("UpdateRun", ctx, runID, runStatus, endTime, runName)}
}

func (_c *MockTrackingStore_UpdateRun_Call) Run(run func(ctx context.Context, runID string, runStatus string, endTime *int64, runName string)) *MockTrackingStore_UpdateRun_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*int64), args[4].(string))
	})
	return _c
}

func (_c *MockTrackingStore_UpdateRun_Call) Return(_a0 *contract.Error) *MockTrackingStore_UpdateRun_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTrackingStore_UpdateRun_Call) RunAndReturn(run func(context.Context, string, string, *int64, string) *contract.Error) *MockTrackingStore_UpdateRun_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTrackingStore creates a new instance of MockTrackingStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTrackingStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTrackingStore {
	mock := &MockTrackingStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
