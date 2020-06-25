// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eks

type ClusterStatus string

// Enum values for ClusterStatus
const (
	ClusterStatusCreating ClusterStatus = "CREATING"
	ClusterStatusActive   ClusterStatus = "ACTIVE"
	ClusterStatusDeleting ClusterStatus = "DELETING"
	ClusterStatusFailed   ClusterStatus = "FAILED"
)

func (enum ClusterStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ClusterStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ErrorCode string

// Enum values for ErrorCode
const (
	ErrorCodeSubnetNotFound        ErrorCode = "SubnetNotFound"
	ErrorCodeSecurityGroupNotFound ErrorCode = "SecurityGroupNotFound"
	ErrorCodeEniLimitReached       ErrorCode = "EniLimitReached"
	ErrorCodeIpNotAvailable        ErrorCode = "IpNotAvailable"
	ErrorCodeAccessDenied          ErrorCode = "AccessDenied"
	ErrorCodeOperationNotPermitted ErrorCode = "OperationNotPermitted"
	ErrorCodeVpcIdNotFound         ErrorCode = "VpcIdNotFound"
	ErrorCodeUnknown               ErrorCode = "Unknown"
)

func (enum ErrorCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ErrorCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type LogType string

// Enum values for LogType
const (
	LogTypeApi               LogType = "api"
	LogTypeAudit             LogType = "audit"
	LogTypeAuthenticator     LogType = "authenticator"
	LogTypeControllerManager LogType = "controllerManager"
	LogTypeScheduler         LogType = "scheduler"
)

func (enum LogType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum LogType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type UpdateParamType string

// Enum values for UpdateParamType
const (
	UpdateParamTypeVersion               UpdateParamType = "Version"
	UpdateParamTypePlatformVersion       UpdateParamType = "PlatformVersion"
	UpdateParamTypeEndpointPrivateAccess UpdateParamType = "EndpointPrivateAccess"
	UpdateParamTypeEndpointPublicAccess  UpdateParamType = "EndpointPublicAccess"
	UpdateParamTypeClusterLogging        UpdateParamType = "ClusterLogging"
)

func (enum UpdateParamType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum UpdateParamType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type UpdateStatus string

// Enum values for UpdateStatus
const (
	UpdateStatusInProgress UpdateStatus = "InProgress"
	UpdateStatusFailed     UpdateStatus = "Failed"
	UpdateStatusCancelled  UpdateStatus = "Cancelled"
	UpdateStatusSuccessful UpdateStatus = "Successful"
)

func (enum UpdateStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum UpdateStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type UpdateType string

// Enum values for UpdateType
const (
	UpdateTypeVersionUpdate        UpdateType = "VersionUpdate"
	UpdateTypeEndpointAccessUpdate UpdateType = "EndpointAccessUpdate"
	UpdateTypeLoggingUpdate        UpdateType = "LoggingUpdate"
)

func (enum UpdateType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum UpdateType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}