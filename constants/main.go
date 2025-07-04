package constants

const (
	MMAR_VERSION = "0.2.6"

	VERSION_CMD        = "version"
	SERVER_CMD         = "server"
	CLIENT_CMD         = "client"
	CLIENT_LOCAL_PORT  = "8000"
	CLIENT_LOCAL_HOST  = "localhost"
	CLIENT_LOCAL_PROTO = "http"
	SERVER_HTTP_PORT   = "3376"
	SERVER_TCP_PORT    = "6673"
	TUNNEL_HOST        = "mmar.dev"
	TUNNEL_HTTP_PORT   = "443"

	MMAR_ENV_VAR_SERVER_HTTP_PORT = "MMAR__SERVER_HTTP_PORT"
	MMAR_ENV_VAR_SERVER_TCP_PORT  = "MMAR__SERVER_TCP_PORT"
	MMAR_ENV_VAR_LOCAL_PORT       = "MMAR__LOCAL_PORT"
	MMAR_ENV_VAR_LOCAL_HOST       = "MMAR__LOCAL_HOST"
	MMAR_ENV_VAR_LOCAL_PROTO      = "MMAR__LOCAL_PROTO"
	MMAR_ENV_VAR_TUNNEL_HTTP_PORT = "MMAR__TUNNEL_HTTP_PORT"
	MMAR_ENV_VAR_TUNNEL_TCP_PORT  = "MMAR__TUNNEL_TCP_PORT"
	MMAR_ENV_VAR_TUNNEL_HOST      = "MMAR__TUNNEL_HOST"

	SERVER_STATS_DEFAULT_USERNAME = "admin"
	SERVER_STATS_DEFAULT_PASSWORD = "admin"

	SERVER_HTTP_PORT_HELP = "Define port where mmar will bind to and run on server for HTTP requests."
	SERVER_TCP_PORT_HELP  = "Define port where mmar will bind to and run on server for TCP connections."

	CLIENT_LOCAL_PORT_HELP  = "Define the port where your local dev server is running to expose through mmar."
	CLIENT_LOCAL_HOST_HELP  = "Define the hostname where your local dev server is running to expose through mmar."
	CLIENT_LOCAL_PROTO_HELP = "Define the protocol where your local dev server is running to expose (http / https)."
	CLIENT_HTTP_PORT_HELP   = "Define port of mmar HTTP server to make requests through the tunnel."
	CLIENT_TCP_PORT_HELP    = "Define port of mmar TCP server for client to connect to, creating a tunnel."
	TUNNEL_HOST_HELP        = "Define host domain of mmar server for client to connect to."

	TUNNEL_MESSAGE_PROTOCOL_VERSION = 3
	TUNNEL_MESSAGE_DATA_DELIMITER   = '\n'
	ID_CHARSET                      = "abcdefghijklmnopqrstuvwxyz0123456789"
	ID_LENGTH                       = 6

	MAX_TUNNELS_PER_IP            = 5
	TUNNEL_RECONNECT_TIMEOUT      = 3
	GRACEFUL_SHUTDOWN_TIMEOUT     = 3
	TUNNEL_CREATE_TIMEOUT         = 3
	REQ_BODY_READ_CHUNK_TIMEOUT   = 3
	DEST_REQUEST_TIMEOUT          = 30
	HEARTBEAT_FROM_SERVER_TIMEOUT = 5
	HEARTBEAT_FROM_CLIENT_TIMEOUT = 2
	READ_DEADLINE                 = 3
	MAX_REQ_BODY_SIZE             = 10000000 // 10mb

	CLIENT_DISCONNECT_ERR_TEXT                    = "Tunnel is closed, cannot connect to mmar client."
	LOCALHOST_NOT_RUNNING_ERR_TEXT                = "Tunneled successfully, but nothing is running on localhost."
	DEST_REQUEST_TIMEDOUT_ERR_TEXT                = "Destination server took too long to respond"
	READ_BODY_CHUNK_ERR_TEXT                      = "Error reading request body"
	READ_BODY_CHUNK_TIMEOUT_ERR_TEXT              = "Timeout reading request body"
	READ_RESP_BODY_ERR_TEXT                       = "Could not read response from destination server, check your server's logs for any errors."
	MAX_REQ_BODY_SIZE_ERR_TEXT                    = "Request too large"
	FAILED_TO_FORWARD_TO_MMAR_CLIENT_ERR_TEXT     = "Failed to forward request to mmar client"
	FAILED_TO_READ_RESP_FROM_MMAR_CLIENT_ERR_TEXT = "Fail to read response from mmad client"

	// TERMINAL ANSI ESCAPED COLORS
	DEFAULT_COLOR = ""
	RED           = "\033[31m"
	GREEN         = "\033[32m"
	YELLOW        = "\033[33m"
	BLUE          = "\033[34m"
	RESET         = "\033[0m"
)

var (
	MMAR_SUBCOMMANDS = [][]string{
		{"server", "Runs a mmar server. Run this on your publicly reachable server if you're self-hosting mmar."},
		{"client", "Runs a mmar client. Run this on your machine to expose your localhost on a public URL."},
		{"version", "Prints the installed version of mmar."},
	}
)
