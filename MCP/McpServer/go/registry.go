package main

import (
	"github.com/input-api/mcp-server/config"
	"github.com/input-api/mcp-server/models"
	tools_connection "github.com/input-api/mcp-server/tools/connection"
	tools_caching "github.com/input-api/mcp-server/tools/caching"
	tools_debug "github.com/input-api/mcp-server/tools/debug"
	tools_encoding "github.com/input-api/mcp-server/tools/encoding"
	tools_text "github.com/input-api/mcp-server/tools/text"
	tools_cors "github.com/input-api/mcp-server/tools/cors"
	tools_security_headers "github.com/input-api/mcp-server/tools/security_headers"
	tools_system "github.com/input-api/mcp-server/tools/system"
	tools_protocol "github.com/input-api/mcp-server/tools/protocol"
	tools_content_negotiation "github.com/input-api/mcp-server/tools/content_negotiation"
	tools_headers "github.com/input-api/mcp-server/tools/headers"
	tools_retry "github.com/input-api/mcp-server/tools/retry"
	tools_json "github.com/input-api/mcp-server/tools/json"
	tools_audit "github.com/input-api/mcp-server/tools/audit"
	tools_auth "github.com/input-api/mcp-server/tools/auth"
	tools_discovery "github.com/input-api/mcp-server/tools/discovery"
	tools_metadata "github.com/input-api/mcp-server/tools/metadata"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_connection.CreateHead_connectionTool(cfg),
		tools_caching.CreateHead_if_none_matchTool(cfg),
		tools_debug.CreateGet_debug_pprofTool(cfg),
		tools_encoding.CreateHead_accept_encodingTool(cfg),
		tools_debug.CreateGet_debug_pprof_cmdlineTool(cfg),
		tools_text.CreateGet_textTool(cfg),
		tools_debug.CreateGet_debug_flagsTool(cfg),
		tools_cors.CreateHead_access_control_allow_methodsTool(cfg),
		tools_security_headers.CreateHead_x_content_type_optionsTool(cfg),
		tools_system.CreateGetTool(cfg),
		tools_protocol.CreateHead_x_app_protocolTool(cfg),
		tools_debug.CreateGet_debug_pprof_symbolTool(cfg),
		tools_cors.CreateHead_access_control_allow_headersTool(cfg),
		tools_content_negotiation.CreateHead_acceptTool(cfg),
		tools_headers.CreateHead_originTool(cfg),
		tools_retry.CreateHead_retry_afterTool(cfg),
		tools_debug.CreateGet_debug_pprof_traceTool(cfg),
		tools_json.CreateGet_jsonTool(cfg),
		tools_cors.CreateHead_access_control_allow_originTool(cfg),
		tools_cors.CreateHead_access_control_expose_headersTool(cfg),
		tools_audit.CreateHead_audit_idTool(cfg),
		tools_debug.CreateGet_debug_flagsTool(cfg),
		tools_auth.CreateHead_authorizationTool(cfg),
		tools_debug.CreateGet_debug_pprofTool(cfg),
		tools_debug.CreateGet_debug_pprof_profileTool(cfg),
		tools_discovery.CreateHead_content_typeTool(cfg),
		tools_metadata.CreateHead_etagTool(cfg),
		tools_cors.CreateHead_access_control_allow_credentialsTool(cfg),
	}
}
