package helpers

import (
	"errors"
	"flyawayhub-cli/logging"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var _ logging.Logger = (*MockLogger)(nil) // Ensure MockLogger implements logging.Logger

// MockLogger mocks the Logger interface for testing.
type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Info(msg string, keysAndValues ...interface{}) {
	m.Called(msg, keysAndValues)
}

func (m *MockLogger) Error(msg string, keysAndValues ...interface{}) {
	m.Called(msg, keysAndValues)
}

func (m *MockLogger) Debug(msg string, keysAndValues ...interface{}) {
	m.Called(msg, keysAndValues)
}

func (m *MockLogger) Warn(msg string, keysAndValues ...interface{}) {
	m.Called(msg, keysAndValues)
}

func (m *MockLogger) DPanic(msg string, keysAndValues ...interface{}) {
	m.Called(msg, keysAndValues)
}

func (m *MockLogger) Panic(msg string, keysAndValues ...interface{}) {
	m.Called(msg, keysAndValues)
}

func (m *MockLogger) Fatal(msg string, keysAndValues ...interface{}) {
	m.Called(msg, keysAndValues)
}

// MockKeyring mocks the Keyring interface for testing.
type MockKeyring struct {
	mock.Mock
}

func (m *MockKeyring) Set(service, account, secret string) error {
	args := m.Called(service, account, secret)
	return args.Error(0)
}

func (m *MockKeyring) Get(service, account string) (string, error) {
	args := m.Called(service, account)
	return args.String(0), args.Error(1)
}

func (m *MockKeyring) Delete(service, account string) error {
	args := m.Called(service, account)
	return args.Error(0)
}

// Example of testing SetSecret assuming you can inject a mock keyring and logger.
func TestSetSecret(t *testing.T) {
	mockLogger := new(MockLogger)

	// Adjust to match the exact logging messages and key-value pairs used in your SetSecret function
	infoMsg := "Secret set successfully"
	service := "service"
	account := "account"
	secret := "secret"

	// Assuming the error path is not triggered for this test, only Info should be called
	mockLogger.On("Info", infoMsg, []interface{}{"service", service, "account", account}).Once()

	// Set the mockLogger as the global logger
	logging.SetLogger(mockLogger)
	defer logging.ResetLogger()

	// Call your function
	err := SetSecret(service, account, secret)

	assert.NoError(t, err)
	mockLogger.AssertExpectations(t)
}

func TestGetSecret(t *testing.T) {
	mockKeyring := new(MockKeyring)
	mockLogger := new(MockLogger)
	logging.SetLogger(mockLogger)
	defer logging.ResetLogger()

	mockKeyring.On("Get", "service", "account").Return("", errors.New("keyring error"))
	mockLogger.On("Error", mock.Anything, mock.Anything).Once()

	_, err := GetSecret("service", "account")

	assert.Error(t, err)
	mockKeyring.AssertExpectations(t)
	mockLogger.AssertExpectations(t)
}

func TestDeleteSecret(t *testing.T) {
	mockKeyring := new(MockKeyring)
	mockKeyring.On("Delete", "service", "account").Return(nil)

	// Inject mockKeyring into your helpers package or specific function, depending on implementation
	err := DeleteSecret("service", "account") // Adjust based on actual implementation

	assert.NoError(t, err)
	mockKeyring.AssertExpectations(t)
}
