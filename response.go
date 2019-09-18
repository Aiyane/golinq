package golinq

import (
	"database/sql"
	"database/sql/driver"
	"sync"
)

const (
	// DriverName is the name of the fake driver
	DriverName = "MOCK_FAKE_DRIVER"
)

// Catcher is global instance of Catcher used for attaching all mocks to connection
var Catcher *MockCatcher

// MockCatcher is global entity to save all mocks aka FakeResponses
type MockCatcher struct {
	Mocks                []*FakeResponse // Slice of all mocks
	Logging              bool            // Do we need to log what we catching?
	PanicOnEmptyResponse bool            // If not response matches - do we need to panic?
	mu                   sync.Mutex
}

func (mc *MockCatcher) SetLogging(l bool) {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	mc.Logging = l
}

// Register safely register FakeDriver
func (mc *MockCatcher) Register() {
	driversList := sql.Drivers()
	for _, name := range driversList {
		if name == DriverName {
			return
		}
	}
	sql.Register(DriverName, &FakeDriver{})
}

// 钩子，搞个实现以下可以搞个连接失败
type Exceptions struct {
	HookQueryBadConnection func() bool
	HookExecBadConnection  func() bool
}

// FakeResponse represents mock of response with holding all required values to return mocked response
type FakeResponse struct {
	Pattern      string                            // 保存需要匹配的sql语句
	Strict       bool                              // 是不是必须完全匹配
	Args         []interface{}                     // 参数列表
	Response     []map[string]interface{}          // 响应结果
	Once         bool                              // 就触发一次
	Triggered    bool                              // 触发过就是true
	Callback     func(string, []driver.NamedValue) // Callback to execute when response triggered
	RowsAffected int64                             // 影响的行数，delete和update可以用
	LastInsertID int64                             // 插入完返回的id，insert可以用
	Error        error                             // Any type of error which could happen dur
	mu           sync.Mutex                        // Used to lock concurrent access to variables
	*Exceptions
}

func init() {
	Catcher = &MockCatcher{}
}
