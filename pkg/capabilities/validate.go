package capabilities

import (
	"encoding/json"

	"github.com/invopop/jsonschema"

	"github.com/smartcontractkit/chainlink-common/pkg/values"
)

type Validator struct {
	reflector *jsonschema.Reflector
	requestConfig any
	requestConfigSchema *values.String
}

var _ Validatable = (*Validator)(nil)

func NewValidator(requestConfig any, reflector *jsonschema.Reflector) *Validator {
	return &Validator{
		requestConfig: requestConfig,
		reflector: reflector,
	}
}

func (v *Validator) GetRequestConfigJSONSchema()  (*CapabilityResponse) {
	if v.requestConfigSchema != nil {
		return &CapabilityResponse{
			Value: v.requestConfigSchema,
		} 
	}
	schema := jsonschema.Reflect(v.requestConfig)
	if v.reflector != nil {
		schema = v.reflector.Reflect(v.requestConfig)
	}

	schemaBytes, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return &CapabilityResponse{
			Err: err,
		}
	}

	wrapped := values.NewString(string(schemaBytes))
	v.requestConfigSchema = wrapped
	return &CapabilityResponse{
		Value: wrapped,
	}
}
