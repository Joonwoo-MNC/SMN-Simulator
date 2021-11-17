package context

import (
	"fmt"
	"os"

	"github.com/google/uuid"

	"github.com/free5gc/openapi/Nnrf_NFDiscovery"
	"github.com/free5gc/openapi/Nnrf_NFManagement"
	"github.com/free5gc/openapi/Nudm_SubscriberDataManagement"
	"github.com/free5gc/openapi/models"
	"samf.com/factory"
	"samf.com/logger"
)

func init() {
	samfContext.NfInstanceID = uuid.New().String()
}

var samfContext SAMFContext

type SAMFContext struct {
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
// 	atomic.AddUint64(&samfContext.LocalSEIDCount, 1)
// 	return samfContext.LocalSEIDCount //if error delete this
// }

func InitSamfContext(config *factory.Config) {
	if config == nil {
		logger.CtxLog.Error("Config is nil")
		return
	}

	logger.CtxLog.Infof("samfconfig Info: Version[%s] Description[%s]", config.Info.Version, config.Info.Description)
	configuration := config.Configuration
	if configuration.SamfName != "" {
		samfContext.Name = configuration.SamfName
	}

	sbi := configuration.Sbi
	if sbi == nil {
		logger.CtxLog.Errorln("Configuration needs \"sbi\" value")
		return
	} else {
		samfContext.URIScheme = models.UriScheme(sbi.Scheme)
		samfContext.RegisterIPv4 = factory.SAMF_DEFAULT_IPV4 // default localhost
		samfContext.SBIPort = factory.SAMF_DEFAULT_PORT_INT  // default port
		if sbi.RegisterIPv4 != "" {
			samfContext.RegisterIPv4 = sbi.RegisterIPv4
		}
		if sbi.Port != 0 {
			samfContext.SBIPort = sbi.Port
		}

		samfContext.BindingIPv4 = os.Getenv(sbi.BindingIPv4)
		if samfContext.BindingIPv4 != "" {
			logger.CtxLog.Info("Parsing ServerIPv4 address from ENV Variable.")
		} else {
			samfContext.BindingIPv4 = sbi.BindingIPv4
			if samfContext.BindingIPv4 == "" {
				logger.CtxLog.Warn("Error parsing ServerIPv4 address as string. Using the 0.0.0.0 address as default.")
				samfContext.BindingIPv4 = "0.0.0.0"
			}
		}
	}

	if configuration.NrfUri != "" {
		samfContext.NrfUri = configuration.NrfUri
	} else {
		logger.CtxLog.Warn("NRF Uri is empty! Using localhost as NRF IPv4 address.")
		samfContext.NrfUri = fmt.Sprintf("%s://%s:%d", samfContext.URIScheme, "127.0.0.1", 29510)
	}

	//SetupNFProfile(config)
}

func SAMF_Self() *SAMFContext {
	return &samfContext
}
