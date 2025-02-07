package packet

const (
	CALLBACK_OUTPUT            = 0
	CALLBACK_KEYSTROKES        = 1
	CALLBACK_FILE              = 2
	CALLBACK_SCREENSHOT        = 3
	CALLBACK_CLOSE             = 4
	CALLBACK_READ              = 5
	CALLBACK_CONNECT           = 6
	CALLBACK_PING              = 7
	CALLBACK_FILE_WRITE        = 8
	CALLBACK_FILE_CLOSE        = 9
	CALLBACK_PIPE_OPEN         = 10
	CALLBACK_PIPE_CLOSE        = 11
	CALLBACK_PIPE_READ         = 12
	CALLBACK_POST_ERROR        = 13
	CALLBACK_PIPE_PING         = 14
	CALLBACK_TOKEN_STOLEN      = 15
	CALLBACK_TOKEN_GETUID      = 16
	CALLBACK_PROCESS_LIST      = 17
	CALLBACK_POST_REPLAY_ERROR = 18
	CALLBACK_PWD               = 19
	CALLBACK_JOBS              = 20
	CALLBACK_HASHDUMP          = 21
	CALLBACK_PENDING           = 22
	CALLBACK_ACCEPT            = 23
	CALLBACK_NETVIEW           = 24
	CALLBACK_PORTSCAN          = 25
	CALLBACK_DEAD              = 26
	CALLBACK_SSH_STATUS        = 27
	CALLBACK_CHUNK_ALLOCATE    = 28
	CALLBACK_CHUNK_SEND        = 29
	CALLBACK_OUTPUT_OEM        = 30
	CALLBACK_ERROR             = 31
	CALLBACK_OUTPUT_UTF8       = 32
)

const (
	CMD_TYPE_SLEEP                      = 4
	CMD_TYPE_PAUSE                      = 47
	CMD_TYPE_SHELL                      = 78
	CMD_TYPE_UPLOAD_START               = 10
	CMD_TYPE_UPLOAD_LOOP                = 67
	CMD_TYPE_DOWNLOAD                   = 11
	CMD_TYPE_EXIT                       = 3
	CMD_TYPE_CD                         = 5
	CMD_TYPE_PWD                        = 39
	CMD_TYPE_FILE_BROWSE                = 53
	CMD_TYPE_SPAWN_X64                  = 44
	CMD_TYPE_SPAWN_X86                  = 1
	CMD_TYPE_EXECUTE                    = 12
	CMD_TYPE_GETUID                     = 27
	CMD_TYPE_GET_PRIVS                  = 77
	CMD_TYPE_STEAL_TOKEN                = 31
	CMD_TYPE_PS                         = 32
	CMD_TYPE_KILL                       = 33
	CMD_TYPE_DRIVES                     = 55
	CMD_TYPE_RUNAS                      = 38
	CMD_TYPE_MKDIR                      = 54
	CMD_TYPE_RM                         = 56
	CMD_TYPE_CP                         = 73
	CMD_TYPE_MV                         = 74
	CMD_TYPE_REV2SELF                   = 28
	CMD_TYPE_MAKE_TOKEN                 = 49
	CMD_TYPE_PIPE                       = 40
	CMD_TYPE_PORTSCAN_X86               = 89
	CMD_TYPE_PORTSCAN_X64               = 90
	CMD_TYPE_KEYLOGGER                  = 101
	CMD_TYPE_EXECUTE_ASSEMBLY_X64       = 88
	CMD_TYPE_EXECUTE_ASSEMBLY_X86       = 87
	CMD_TYPE_EXECUTE_ASSEMBLY_TOKEN_X64 = 71
	CMD_TYPE_EXECUTE_ASSEMBLY_TOKEN_X86 = 70
	CMD_TYPE_IMPORT_POWERSHELL          = 37
	CMD_TYPE_POWERSHELL_PORT            = 79
	CMD_TYPE_INJECT_X64                 = 43
	CMD_TYPE_INJECT_X86                 = 9
	CMD_TYPE_BOF                        = 100
	CMD_TYPE_RUNU                       = 76
	CMD_TYPE_ARGUE_QUERY                = 85
	CMD_TYPE_ARGUE_REMOVE               = 84
	CMD_TYPE_ARGUE_ADD                  = 83
)
