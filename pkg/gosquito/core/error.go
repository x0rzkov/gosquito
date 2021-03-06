package core

import "errors"

var (
	ERROR_DATA_FIELD_KEY               = errors.New("data field key must be string: %v")
	ERROR_DATA_FIELD_NOT_SLICE         = errors.New("data field not slice: %s")
	ERROR_DATA_FIELD_TYPE_MISMATCH     = errors.New("data field type mismatch: %s")
	ERROR_DATA_FIELD_UNKNOWN           = errors.New("data field unknown: %s")
	ERROR_EXPORTER_LISTEN              = errors.New("exporter error")
	ERROR_FLOW_EXPIRE                  = errors.New("flow expire")
	ERROR_FLOW_NAME_COMPAT             = errors.New("flow name must be compatible: %s")
	ERROR_FLOW_NAME_UNIQUE             = errors.New("flow name must be unique: %s")
	ERROR_FLOW_PARSE                   = errors.New("flow parse error")
	ERROR_FILE_YAML                    = errors.New("only yml/yaml file extensions are accepted")
	ERROR_FLOW_DISABLED                = errors.New("flow disabled")
	ERROR_INTERVAL_FORMAT_UNKNOWN      = errors.New("interval format unknown")
	ERROR_FLOW_ENABLE_DISABLE_CONFLICT = errors.New(
		"default.flow_disable & default.flow_enable are mutual exclusive!")
	ERROR_NO_NEW_DATA           = errors.New("no new data")
	ERROR_NO_VALID_FLOW         = errors.New("no valid flow")
	ERROR_PARAM_ERROR           = errors.New("parameter error")
	ERROR_PARAM_KEY_MUST_STRING = errors.New("parameter key must be string")
	ERROR_PARAM_UNKNOWN         = errors.New("unknown parameter: %s")
	ERROR_PLUGIN_DATA_READ      = errors.New("plugin data read error: %s")
	ERROR_PLUGIN_DATA_WRITE     = errors.New("plugin data write error: %s")
	ERROR_PLUGIN_PROCESS_ORDER  = errors.New("plugin id must be ordered")
	ERROR_PLUGIN_PROCESS_PARAMS = errors.New("plugin must have: [id, plugin, params]")
	ERROR_PLUGIN_REQUIRED_PARAM = errors.New("required parameter wrong or not set: %s")
	ERROR_PLUGIN_UNKNOWN        = errors.New("plugin unknown")
	ERROR_SIZE_MISMATCH         = errors.New("size mismatch: %s, %s")
	ERROR_SIZE_FORMAT_UNKNOWN   = errors.New("size format unknown")
)
