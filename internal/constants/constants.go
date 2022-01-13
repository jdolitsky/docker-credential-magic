package constants

const (
	AnonymousTokenResponse            = "{\"Username\":\"\",\"Secret\":\"\"}\n"
	BinariesSubdir                    = "bin"
	DockerConfigFileBasename          = "config.json"
	DockerConfigFileContents          = "{\"credsStore\":\"magic\"}\n"
	DockerCredentialPrefix            = "docker-credential"
	DockerHomeDir                     = ".docker"
	EmbeddedParentDir                 = "embedded"
	EnvVarDockerConfig                = "DOCKER_CONFIG"
	EnvVarDockerCredentialMagicConfig = "DOCKER_CREDENTIAL_MAGIC_CONFIG"
	EnvVarDockerCredentialMagicChild  = "DOCKER_CREDENTIAL_MAGIC_CHILD"
	EnvVarDockerOrigConfig            = "DOCKER_ORIG_CONFIG"
	EnvVarPath                        = "PATH"
	ExtensionYAML                     = "yml"
	HelperSubcommandGet               = "get"
	MagicCredentialSuffix             = "magic"
	MagicRootDir                      = "/opt/magic"
	MappingsSubdir                    = "etc"
	XDGConfigSubdir                   = "magic"
)
