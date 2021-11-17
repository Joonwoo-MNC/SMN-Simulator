package context

import (
	"fmt"
	"os"

	"github.com/google/uuid"

	"github.com/free5gc/openapi/Nnrf_NFDiscovery"
	"github.com/free5gc/openapi/Nnrf_NFManagement"
	"github.com/free5gc/openapi/Nudm_SubscriberDataManagement"
	"github.com/free5gc/openapi/models"
	"ssmf.com/factory"
	"ssmf.com/logger"
)

func init() {
	ssmfContext.NfInstanceID = uuid.New().String()
}

var ssmfContext SSMFContext

type SSMFContext struct {
	Name            string
	URIScheme       models.UriScheme
	UriScheme       models.UriScheme
	BindingIPv4     string
	RegisterIPv4    string
	SBIPort         int
	HttpIPv6Address string
	NfInstanceID    string
	NfId            string
	// Key    string
	// PEM    string
	// KeyLog string

	NrfUri string

	// Now only "IPv4" supported
	// TODO: support "IPv6", "IPv4v6", "Ethernet"
	SupportedPDUSessionType string

	//*** For ULCL ** //
	// ULCLSupport    bool
	// ULCLGroups     map[string][]string
	// LocalSEIDCount uint64

	NFManagementClient             *Nnrf_NFManagement.APIClient
	NFDiscoveryClient              *Nnrf_NFDiscovery.APIClient
	SubscriberDataManagementClient *Nudm_SubscriberDataManagement.APIClient
}

// RetrieveDnnInformation gets the corresponding dnn info from S-NSSAI and DNN

// func AllocateLocalSEID() uint64 {
// 	atomic.AddUint64(&ssmfContext.LocalSEIDCount, 1)
// 	return ssmfContext.LocalSEIDCount //if error delete this
// }

func InitSsmfContext(config *factory.Config) {
	if config == nil {
		logger.CtxLog.Error("Config is nil")
		return
	}

	logger.CtxLog.Infof("ssmfconfig Info: Version[%s] Description[%s]", config.Info.Version, config.Info.Description)
	configuration := config.Configuration
	if configuration.SsmfName != "" {
		ssmfContext.Name = configuration.SsmfName
	}

	sbi := configuration.Sbi
	if sbi == nil {
		logger.CtxLog.Errorln("Configuration needs \"sbi\" value")
		return
	} else {
		ssmfContext.URIScheme = models.UriScheme(sbi.Scheme)
		ssmfContext.RegisterIPv4 = factory.SSMF_DEFAULT_IPV4 // default localhost
		ssmfContext.SBIPort = factory.SSMF_DEFAULT_PORT_INT  // default port
		if sbi.RegisterIPv4 != "" {
			ssmfContext.RegisterIPv4 = sbi.RegisterIPv4
		}
		if sbi.Port != 0 {
			ssmfContext.SBIPort = sbi.Port
		}

		ssmfContext.BindingIPv4 = os.Getenv(sbi.BindingIPv4)
		if ssmfContext.BindingIPv4 != "" {
			logger.CtxLog.Info("Parsing ServerIPv4 address from ENV Variable.")
		} else {
			ssmfContext.BindingIPv4 = sbi.BindingIPv4
			if ssmfContext.BindingIPv4 == "" {
				logger.CtxLog.Warn("Error parsing ServerIPv4 address as string. Using the 0.0.0.0 address as default.")
				ssmfContext.BindingIPv4 = "0.0.0.0"
			}
		}
	}

	if configuration.NrfUri != "" {
		ssmfContext.NrfUri = configuration.NrfUri
	} else {
		logger.CtxLog.Warn("NRF Uri is empty! Using localhost as NRF IPv4 address.")
		ssmfContext.NrfUri = fmt.Sprintf("%s://%s:%d", ssmfContext.URIScheme, "127.0.0.1", 29510)
	}

	//SetupNFProfile(config)
}

func SSMF_Self() *SSMFContext {
	return &ssmfContext
}
