package capabilities

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	p2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/values"
)

// CapabilityType is an enum for the type of capability.
type CapabilityType int

var ErrStopExecution = &errStopExecution{}

type errStopExecution struct{}

const errStopExecutionMsg = "__workflow_stop_execution"

func (e errStopExecution) Error() string {
	return errStopExecutionMsg
}

func (e errStopExecution) Is(err error) bool {
	return err.Error() == errStopExecutionMsg
}

// CapabilityType enum values.
const (
	CapabilityTypeTrigger CapabilityType = iota
	CapabilityTypeAction
	CapabilityTypeConsensus
	CapabilityTypeTarget
)

// String returns a string representation of CapabilityType
func (c CapabilityType) String() string {
	switch c {
	case CapabilityTypeTrigger:
		return "trigger"
	case CapabilityTypeAction:
		return "action"
	case CapabilityTypeConsensus:
		return "consensus"
	case CapabilityTypeTarget:
		return "target"
	}

	// Panic as this should be unreachable.
	panic("unknown capability type")
}

// IsValid checks if the capability type is valid.
func (c CapabilityType) IsValid() error {
	switch c {
	case CapabilityTypeTrigger,
		CapabilityTypeAction,
		CapabilityTypeConsensus,
		CapabilityTypeTarget:
		return nil
	}

	return fmt.Errorf("invalid capability type: %s", c)
}

// CapabilityResponse is a struct for the Execute response of a capability.
type CapabilityResponse struct {
	Value *values.Map
	Err   error
}

type RequestMetadata struct {
	WorkflowID               string
	WorkflowOwner            string
	WorkflowExecutionID      string
	WorkflowName             string
	WorkflowDonID            uint32
	WorkflowDonConfigVersion uint32
}

type RegistrationMetadata struct {
	WorkflowID    string
	WorkflowOwner string
}

// CapabilityRequest is a struct for the Execute request of a capability.
type CapabilityRequest struct {
	Metadata RequestMetadata
	Config   *values.Map
	Inputs   *values.Map
}

type TriggerEvent struct {
	TriggerType string
	ID          string
	Timestamp   string
	// Trigger-specific payload+metadata
	Metadata values.Value
	Payload  values.Value
}

type RegisterToWorkflowRequest struct {
	Metadata RegistrationMetadata
	Config   *values.Map
}

type UnregisterFromWorkflowRequest struct {
	Metadata RegistrationMetadata
	Config   *values.Map
}

// CallbackExecutable is an interface for executing a capability.
type CallbackExecutable interface {
	RegisterToWorkflow(ctx context.Context, request RegisterToWorkflowRequest) error
	UnregisterFromWorkflow(ctx context.Context, request UnregisterFromWorkflowRequest) error
	// Capability must respect context.Done and cleanup any request specific resources
	// when the context is cancelled. When a request has been completed the capability
	// is also expected to close the callback channel.
	// Request specific configuration is passed in via the request parameter.
	// A successful response must always return a value. An error is assumed otherwise.
	// The intent is to make the API explicit.
	Execute(ctx context.Context, request CapabilityRequest) (<-chan CapabilityResponse, error)
}

type Validatable interface {
	// ValidateSchema returns the JSON schema for the capability.
	//
	// This schema includes the configuration, input and output schemas.
	Schema() (string, error)
}

// BaseCapability interface needs to be implemented by all capability types.
// Capability interfaces are intentionally duplicated to allow for an easy change
// or extension in the future.
type BaseCapability interface {
	Info(ctx context.Context) (CapabilityInfo, error)
}

type TriggerExecutable interface {
	RegisterTrigger(ctx context.Context, request CapabilityRequest) (<-chan CapabilityResponse, error)
	UnregisterTrigger(ctx context.Context, request CapabilityRequest) error
}

// TriggerCapability interface needs to be implemented by all trigger capabilities.
type TriggerCapability interface {
	BaseCapability
	TriggerExecutable
}

// CallbackCapability is the interface implemented by action, consensus and target
// capabilities. This interface is useful when trying to capture capabilities of varying types.
type CallbackCapability interface {
	BaseCapability
	CallbackExecutable
}

// ActionCapability interface needs to be implemented by all action capabilities.
type ActionCapability interface {
	CallbackCapability
}

// ConsensusCapability interface needs to be implemented by all consensus capabilities.
type ConsensusCapability interface {
	CallbackCapability
}

// TargetsCapability interface needs to be implemented by all target capabilities.
type TargetCapability interface {
	CallbackCapability
}

// DON represents a network of connected nodes.
//
// For an example of an empty DON check, see the following link:
// https://github.com/smartcontractkit/chainlink/blob/develop/core/capabilities/transmission/local_target_capability.go#L31
type DON struct {
	ID               uint32
	ConfigVersion    uint32
	Members          []p2ptypes.PeerID
	F                uint8
	IsPublic         bool
	AcceptsWorkflows bool
}

// Node contains the node's peer ID and the DONs it is part of.
//
// Note the following relationships between the workflow and capability DONs and this node.
//
// There is a 1:0..1 relationship between this node and a workflow DON.
// This means that this node can be part at most one workflow DON at a time.
// As a side note, a workflow DON can have multiple nodes.
//
// There is a 1:N relationship between this node and capability DONs, where N is the number of capability DONs.
// This means that this node can be part of multiple capability DONs at a time.
//
// Although WorkflowDON is a value rather than a pointer, a node can be part of no workflow DON but 0 or more capability DONs.
// You can assert this by checking for zero values in the WorkflowDON field.
// See https://github.com/smartcontractkit/chainlink/blob/develop/core/capabilities/transmission/local_target_capability.go#L31 for an example.
type Node struct {
	PeerID         *p2ptypes.PeerID
	WorkflowDON    DON
	CapabilityDONs []DON
}

// CapabilityInfo is a struct for the info of a capability.
type CapabilityInfo struct {
	// The capability ID is a fully qualified identifier for the capability.
	//
	// It takes the form of `{name}:{label1_key}_{labe1_value}:{label2_key}_{label2_value}@{version}`
	//
	// The labels within the ID are ordered alphanumerically.
	ID             string
	CapabilityType CapabilityType
	Description    string
	DON            *DON
}

// Parse out the version from the ID.
func (c CapabilityInfo) Version() string {
	return c.ID[strings.Index(c.ID, "@")+1:]
}

// Info returns the info of the capability.
func (c CapabilityInfo) Info(ctx context.Context) (CapabilityInfo, error) {
	return c, nil
}

// This regex allows for the following format:
//
// {name}:{label1_key}_{labe1_value}:{label2_key}_{label2_value}@{version}
//
// The version regex is taken from https://semver.org/, but has been modified to support only major versions.
//
// It is also validated when a workflow is being ingested. See the following link for more details:
// https://github.com/smartcontractkit/chainlink/blob/a0d1ce5e9cddc540bba8eb193865646cf0ebc0a8/core/services/workflows/models_yaml.go#L309
//
// The difference between the regex within the link above and this one is that we do not use double backslashes, since
// we only needed those for JSON schema regex validation.
var idRegex = regexp.MustCompile(`^[a-z0-9_\-:]+@(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

const (
	// TODO: this length was largely picked arbitrarily.
	// Consider what a realistic/desirable value should be.
	// See: https://smartcontract-it.atlassian.net/jira/software/c/projects/KS/boards/182
	idMaxLength = 128
)

// NewCapabilityInfo returns a new CapabilityInfo.
func NewCapabilityInfo(
	id string,
	capabilityType CapabilityType,
	description string,
) (CapabilityInfo, error) {
	return NewRemoteCapabilityInfo(id, capabilityType, description, nil)
}

// NewRemoteCapabilityInfo returns a new CapabilityInfo for remote capabilities.
// This is largely intended for internal use by the registry syncer.
// Capability developers should use `NewCapabilityInfo` instead as this
// omits the requirement to pass in the DON Info.
func NewRemoteCapabilityInfo(
	id string,
	capabilityType CapabilityType,
	description string,
	don *DON,
) (CapabilityInfo, error) {
	if len(id) > idMaxLength {
		return CapabilityInfo{}, fmt.Errorf("invalid id: %s exceeds max length %d", id, idMaxLength)
	}
	if !idRegex.MatchString(id) {
		return CapabilityInfo{}, fmt.Errorf("invalid id: %s. Allowed: %s", id, idRegex)
	}

	if err := capabilityType.IsValid(); err != nil {
		return CapabilityInfo{}, err
	}

	return CapabilityInfo{
		ID:             id,
		CapabilityType: capabilityType,
		Description:    description,
		DON:            don,
	}, nil
}

// MustNewCapabilityInfo returns a new CapabilityInfo,
// `panic`ing if we could not instantiate a CapabilityInfo.
func MustNewCapabilityInfo(
	id string,
	capabilityType CapabilityType,
	description string,
) CapabilityInfo {
	return MustNewRemoteCapabilityInfo(id, capabilityType, description, nil)
}

// MustNewRemoteCapabilityInfo returns a new CapabilityInfo,
// `panic`ing if we could not instantiate a CapabilityInfo.
func MustNewRemoteCapabilityInfo(
	id string,
	capabilityType CapabilityType,
	description string,
	don *DON,
) CapabilityInfo {
	c, err := NewRemoteCapabilityInfo(id, capabilityType, description, don)
	if err != nil {
		panic(err)
	}

	return c
}

// TODO: this timeout was largely picked arbitrarily.
// Consider what a realistic/desirable value should be.
// See: https://smartcontract-it.atlassian.net/jira/software/c/projects/KS/boards/182
var maximumExecuteTimeout = 60 * time.Second

// ExecuteSync executes a capability synchronously.
// We are not handling a case where a capability panics and crashes.
// There is default timeout of 10 seconds. If a capability takes longer than
// that then it should be executed asynchronously.
func ExecuteSync(ctx context.Context, c CallbackExecutable, request CapabilityRequest) (*values.List, error) {
	return ExecuteSyncWithTimeout(ctx, c, request, maximumExecuteTimeout)
}

// ExecuteSyncWithTimeout allows explicitly passing in a timeout to customise the desired duration.
func ExecuteSyncWithTimeout(ctx context.Context, c CallbackExecutable, request CapabilityRequest, timeout time.Duration) (*values.List, error) {
	ctxWithT, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	responseCh, err := c.Execute(ctxWithT, request)
	if err != nil {
		return nil, fmt.Errorf("error executing capability: %w", err)
	}

	vs := make([]values.Value, 0)
outerLoop:
	for {
		select {
		case response, isOpen := <-responseCh:
			if !isOpen {
				break outerLoop
			}
			// An error means execution has been interrupted.
			// We'll return the value discarding values received
			// until now.
			if response.Err != nil {
				return nil, response.Err
			}

			vs = append(vs, response.Value)
		// Timeout when a capability exceeds maximum permitted execution time or the caller cancels the context and does not close the channel.
		case <-ctxWithT.Done():
			return nil, fmt.Errorf("context timed out after %f seconds", maximumExecuteTimeout.Seconds())
		}
	}

	// If the capability did not return any values, we deem it as an error.
	// The intent is for the API to be explicit.
	if len(vs) == 0 {
		return nil, errors.New("capability did not return any values")
	}

	return &values.List{Underlying: vs}, nil
}

const (
	DefaultRegistrationRefreshMs = 30_000
	DefaultRegistrationExpiryMs  = 120_000
	DefaultMessageExpiryMs       = 120_000
)

type RemoteTriggerConfig struct {
	RegistrationRefreshMs   uint32
	RegistrationExpiryMs    uint32
	MinResponsesToAggregate uint32
	MessageExpiryMs         uint32
}

// NOTE: consider splitting this config into values stored in Registry (KS-118)
// and values defined locally by Capability owners.
func (c *RemoteTriggerConfig) ApplyDefaults() {
	if c.RegistrationRefreshMs == 0 {
		c.RegistrationRefreshMs = DefaultRegistrationRefreshMs
	}
	if c.RegistrationExpiryMs == 0 {
		c.RegistrationExpiryMs = DefaultRegistrationExpiryMs
	}
	if c.MessageExpiryMs == 0 {
		c.MessageExpiryMs = DefaultMessageExpiryMs
	}
}

type CapabilityConfiguration struct {
	ExecuteConfig       *values.Map
	RemoteTriggerConfig RemoteTriggerConfig
}
