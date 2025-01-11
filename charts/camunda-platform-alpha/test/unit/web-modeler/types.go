package web_modeler

// REST API ---
type WebModelerRestAPIApplicationYAML struct {
	Camunda CamundaYAML `yaml:"camunda"`
	Spring  SpringYAML  `yaml:"spring"`
}

type SpringYAML struct {
	Mail       MailYAML       `yaml:"mail"`
	Datasource DatasourceYAML `yaml:"datasource"`
}
type DatasourceYAML struct {
	Url      string `yaml:"url"`
	Username string `yaml:"username"`
}

type MailYAML struct {
	Username string `yaml:"username"`
}

type CamundaYAML struct {
	Modeler  ModelerYAML  `yaml:"modeler"`
	Identity IdentityYAML `yaml:"identity"`
}

type IdentityYAML struct {
	BaseURL string `yaml:"base-url"`
	Type    string `yaml:"type"`
}
type ModelerYAML struct {
	Security SecurityYAML  `yaml:"security"`
	Clusters []ClusterYAML `yaml:"clusters"`
}

type SecurityYAML struct {
	JWT JwtYAML `yaml:"jwt"`
}

type JwtYAML struct {
	Audience AudienceYAML `yaml:"audience"`
	Issuer   IssuerYAML   `yaml:"issuer"`
}

type IssuerYAML struct {
	BackendUrl string `yaml:"backend-url"`
}

type AudienceYAML struct {
	InternalAPI string `yaml:"internal-api"`
	PublicAPI   string `yaml:"public-api"`
}

type ClusterYAML struct {
	Id             string    `yaml:"id"`
	Name           string    `yaml:"name"`
	Version        string    `yaml:"version"`
	Authentication string    `yaml:"authentication"`
	Url            UrlYAML   `yaml:"url"`
	Oauth          OAuthYAML `yaml:"oauth"`
}

type UrlYAML struct {
	Zeebe    ZeebeUrlYAML `yaml:"zeebe"`
	Operate  string       `yaml:"operate"`
	Tasklist string       `yaml:"tasklist"`
}

type ZeebeUrlYAML struct {
	Grpc string `yaml:"grpc"`
	Rest string `yaml:"rest"`
}

type OAuthYAML struct {
	Url      string            `yaml:"url"`
	Scope    string            `yaml:"scope"`
	Audience OAuthAudienceYAML `yaml:"audience"`
}

type OAuthAudienceYAML struct {
	Zeebe    string `yaml:"zeebe"`
	Operate  string `yaml:"operate"`
	Tasklist string `yaml:"tasklist"`
}

// Web App ---

type WebModelerWebAppTOML struct {
	OAuth2   OAuth2Config   `toml:"oAuth2"`
	Client   ClientConfig   `toml:"client"`
	Identity IdentityConfig `toml:"identity"`
	Server   ServerConfig   `toml:"server"`
}
type ServerConfig struct {
	HttpsOnly string `toml:"httpsOnly"`
}
type IdentityConfig struct {
	BaseUrl string `toml:"baseUrl"`
}

type ClientConfig struct {
	Pusher PusherConfig `toml:"pusher"`
}

type PusherConfig struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Path     string `toml:"path"`
	ForceTLS string `toml:"forceTLS"`
}

type OAuth2Config struct {
	Token    TokenConfig `toml:"token"`
	ClientId string      `toml:"clientId"`
	Type     string      `toml:"type"`
}

type TokenConfig struct {
	Audience string `toml:"audience"`
	JwksUrl  string `toml:"jwksUrl"`
}