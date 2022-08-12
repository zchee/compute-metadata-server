// Copyright 2022 The compute-metadata-server Authors
// SPDX-License-Identifier: BSD-3-Clause

package fakemetadata

import (
	"fmt"
	"io/fs"
	"os"
	pathpkg "path"
	"strings"
	"time"
	"unsafe"

	"github.com/google/go-safeweb/safehttp"
	"github.com/google/safehtml"
	cpuid "github.com/klauspost/cpuid/v2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

// InstanceHandler holds instance metadata handlers.
//
// See: https://cloud.google.com/compute/docs/metadata/default-metadata-values#vm_instance_metadata
type InstanceHandler struct{}

// RegisterHandlers registers instance handlers to mux.
func (h *InstanceHandler) RegisterHandlers(mux *safehttp.ServeMux) {
	mux.Handle("/computeMetadata/v1/instance/attributes", safehttp.MethodGet, redirectHandler("computeMetadata/v1/instance/attributes/"))
	mux.Handle("/computeMetadata/v1/instance/attributes/", safehttp.MethodGet, h.Attributes(InstanceAttributeMap))
	mux.Handle("/computeMetadata/v1/instance/cpu-platform", safehttp.MethodGet, h.CPUPlatform())
	mux.Handle("/computeMetadata/v1/instance/description", safehttp.MethodGet, h.Description())
	mux.Handle("/computeMetadata/v1/instance/disks", safehttp.MethodGet, redirectHandler("computeMetadata/v1/instance/disks/"))
	mux.Handle("/computeMetadata/v1/instance/disks/", safehttp.MethodGet, h.Disks())
	mux.Handle("/computeMetadata/v1/instance/guest-attributes", safehttp.MethodGet, redirectHandler("computeMetadata/v1/instance/guest-attributes/"))
	mux.Handle("/computeMetadata/v1/instance/guest-attributes/", safehttp.MethodGet, h.GuestAttributes(InstanceGuestAttributeMap))
	mux.Handle("/computeMetadata/v1/instance/hostname", safehttp.MethodGet, h.Hostname())
	mux.Handle("/computeMetadata/v1/instance/id", safehttp.MethodGet, h.ID())
	mux.Handle("/computeMetadata/v1/instance/image", safehttp.MethodGet, h.Image())
	mux.Handle("/computeMetadata/v1/instance/legacy-endpoint-access/", safehttp.MethodGet, h.LegacyEndpointAccess())
	mux.Handle("/computeMetadata/v1/instance/licenses/", safehttp.MethodGet, h.Licenses())
	mux.Handle("/computeMetadata/v1/instance/machine-type", safehttp.MethodGet, h.MachineType())
	mux.Handle("/computeMetadata/v1/instance/maintenance-event", safehttp.MethodGet, h.MaintenanceEvent())
	mux.Handle("/computeMetadata/v1/instance/name", safehttp.MethodGet, h.Name())
	mux.Handle("/computeMetadata/v1/instance/network-interfaces", safehttp.MethodGet, redirectHandler("computeMetadata/v1/instance/network-interfaces/"))
	mux.Handle("/computeMetadata/v1/instance/network-interfaces/", safehttp.MethodGet, h.NetworkInterfaces())
	mux.Handle("/computeMetadata/v1/instance/preempted", safehttp.MethodGet, h.Preempted())
	mux.Handle("/computeMetadata/v1/instance/remaining-cpu-time", safehttp.MethodGet, h.RemainingCPUTime())
	mux.Handle("/computeMetadata/v1/instance/scheduling", safehttp.MethodGet, redirectHandler("computeMetadata/v1/instance/scheduling/"))
	mux.Handle("/computeMetadata/v1/instance/scheduling/", safehttp.MethodGet, h.Scheduling())
	mux.Handle("/computeMetadata/v1/instance/service-accounts", safehttp.MethodGet, redirectHandler("computeMetadata/v1/instance/service-accounts/"))
	mux.Handle("/computeMetadata/v1/instance/service-accounts/", safehttp.MethodGet, h.ServiceAccounts())
	mux.Handle("/computeMetadata/v1/instance/tags", safehttp.MethodGet, h.Tags())
	mux.Handle("/computeMetadata/v1/instance/virtual-clock", safehttp.MethodGet, redirectHandler("computeMetadata/v1/instance/virtual-clock/"))
	mux.Handle("/computeMetadata/v1/instance/virtual-clock/", safehttp.MethodGet, h.VirtualClock())
	mux.Handle("/computeMetadata/v1/instance/zone", safehttp.MethodGet, h.Zone())
}

// InstanceAttributeMap map of instance attributes.
//
// See: https://cloud.google.com/compute/docs/metadata/default-metadata-values#instance-attributes-metadata
var InstanceAttributeMap = map[string]bool{
	// Enables or disables SSH key management on your VM.
	//
	// For more information about OS Login, see Setting up OS Login.
	"enable-oslogin": true,

	// Enable zonal DNS and global DNS for the VM.
	//
	// For more information about internal DNS names, see Configuring DNS names.
	"vmdnssetting": true,

	// If you are managing SSH keys using metadata, this attribute lets you configure public SSH keys that can connect to VMs in this project.
	// If there are multiple SSH keys, each key is separated by a newline character (\n). The value of the ssh-keys attribute is a string.
	//
	// Example: "user1:ssh-rsa mypublickey user1@host.com\nuser2:ssh-rsa mypublickey user2@host.com"
	//
	// SSH keys managed by OS Login aren't visible in metadata.
	"ssh-keys": true,
}

// Attributes a directory of custom metadata values passed to the VM during startup or shutdown.
// These custom values can either be Google Cloud attributes or user-created metadata values.
//
// For a list of instance-level Google Cloud attributes that you can set, see Instance attributes.
//
// For more information about setting custom metadata, see Setting custom metadata.
func (h *InstanceHandler) Attributes(m map[string]bool) safehttp.Handler {
	handler := safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
		if r.URL().Path() == "" {
			attrs := make([]string, len(m))
			i := 0
			for attr := range m {
				attrs[i] = attr
				i++
			}
			return w.Write(safehtml.HTMLEscaped(strings.Join(attrs, "\n")))
		}

		if path := r.URL().Path(); m[path] {
			switch path {
			case "enable-oslogin":
				// TODO(zchee): not implemented
			case "vmdnssetting":
				// TODO(zchee): not implemented
			case "ssh-keys":
				// TODO(zchee): not implemented
			}
		}

		return w.WriteError(safehttp.StatusNotFound)
	})

	return safehttp.StripPrefix("/computeMetadata/v1/instance/attributes/", handler)
}

// CPUPlatform CPU platform of the VM.
//
// For information about CPU platforms, see CPU platforms.
func (h *InstanceHandler) CPUPlatform() safehttp.Handler {
	return safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
		return w.Write(safehtml.HTMLEscaped(detectCPUMicroarchitecture(cpuid.CPU).String()))
	})
}

// Description is the free-text description of an instance that is assigned using the "--description" flag by using the Google Cloud CLI or the API.
func (h *InstanceHandler) Description() safehttp.Handler {
	return safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
		// TODO(zchee): not implemented
		return safehttp.NotWritten()
	})
}

var diskEndpoint = []string{
	"device-name",
	"index",
	"interface",
	"mode",
	"type",
}

// Disks a directory of disks that are attached to the VM.
//
// For each disk, the following information is available:
//
//	device-name
//	index
//	interface
//	mode
//	type
//
// For more information about disks, see Storage options.
func (h *InstanceHandler) Disks() safehttp.Handler {
	handler := safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
		switch r.URL().Path() {
		case "":
			return w.Write(safehtml.HTMLEscaped(strings.Join(diskEndpoint, "\n")))
		case "device-name":
			// TODO(zchee): not implemented
		case "index":
			// TODO(zchee): not implemented
		case "interface":
			// TODO(zchee): not implemented
		case "mode":
			// TODO(zchee): not implemented
		case "type":
			// TODO(zchee): not implemented
		}

		return w.WriteError(safehttp.StatusNotFound)
	})

	return safehttp.StripPrefix("/computeMetadata/v1/instance/disks/", handler)
}

// InstanceGuestAttributeMap map of instance guest attributes.
//
// See: https://cloud.google.com/compute/docs/metadata/default-metadata-values#instance-guest-attributes-metadata
var InstanceGuestAttributeMap = map[string]bool{
	// Stores OS inventory for the VM.
	//
	// Collects and stores OS details information. This includes information such as hostname, kernel version, architecture, and installed packages details.
	//
	// For more information about OS inventory, see Viewing operating system details.
	"guestInventory": true,

	// Stores SSH host keys. Host keys can be used to identify a particular host or machine.
	//
	// For information host keys, see Storing host keys by enabling guest attributes.
	"hostkeys": true,
}

// GuestAttributes sets guest attributes for the VM. These custom values can either be Google Cloud attributes or user-created metadata values.
//
// For a list of instance-level Google Cloud attributes that you can set, see Instance guest attributes.
//
// Note: Any user or process on your VM instance can read and write to the namespaces and keys in guest-attributes metadata.
//
// For more information about guest attributes, see Setting and querying guest attributes.
func (h *InstanceHandler) GuestAttributes(m map[string]bool) safehttp.Handler {
	handler := safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
		if r.URL().Path() == "" {
			attrs := make([]string, len(m))
			i := 0
			for attr := range m {
				attrs[i] = attr
				i++
			}
			return w.Write(safehtml.HTMLEscaped(strings.Join(attrs, "\n")))
		}

		if path := r.URL().Path(); m[path] {
			switch path {
			case "guestInventory":
				// TODO(zchee): not implemented
			case "hostkeys":
				// TODO(zchee): not implemented
			}
		}

		return w.WriteError(safehttp.StatusNotFound)
	})

	return safehttp.StripPrefix("/computeMetadata/v1/instance/guest-attributes/", handler)
}

const EnvInstanceHostname = "GOOGLE_INSTANCE_HOSTNAME"

// Hostname is the hostname of the VM.
func (h *InstanceHandler) Hostname() safehttp.Handler {
	return safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
		if hostname, ok := os.LookupEnv(EnvInstanceHostname); ok {
			return w.Write(safehtml.HTMLEscaped(hostname))
		}

		return w.WriteError(safehttp.StatusNotFound)
	})
}

const EnvInstanceID = "GOOGLE_INSTANCE_ID"

// ID the ID of the VM. This is a unique, numerical ID that is generated by Compute Engine. This is useful for identifying VMs if you don't use VM names.
func (h *InstanceHandler) ID() safehttp.Handler {
	return safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
		if id, ok := os.LookupEnv(EnvInstanceID); ok {
			return w.Write(safehtml.HTMLEscaped(id))
		}

		return w.WriteError(safehttp.StatusNotFound)
	})
}

// Image is the operating system image used by the VM. This value has the following format:
//
//	projects/IMAGE_PROJECT/global/images/IMAGE_NAME
func (h *InstanceHandler) Image() safehttp.Handler {
	return safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
		return safehttp.NotWritten()
	})
}

// LegacyEndpointAccess stores the list of legacy endpoints. Values are 0.1 and v1beta1.
func (h *InstanceHandler) LegacyEndpointAccess() safehttp.Handler {
	return safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
		return safehttp.NotWritten()
	})
}

// Licenses a list of license code IDs that are used to attach the licenses to images, snapshots, and disks.
// directory
func (h *InstanceHandler) Licenses() safehttp.Handler {
	return safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
		return safehttp.NotWritten()
	})
}

// MachineType is the machine type for this VM. This value has the following format: projects/PROJECT_NUM/machineTypes/MACHINE_TYPE
func (h *InstanceHandler) MachineType() safehttp.Handler {
	return safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
		return safehttp.NotWritten()
	})
}

// MaintenanceEvent indicates whether a maintenance event is affecting this VM. For more information, see Live migrate.
func (h *InstanceHandler) MaintenanceEvent() safehttp.Handler {
	return safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
		return safehttp.NotWritten()
	})
}

// Name is the name of the VM.
func (h *InstanceHandler) Name() safehttp.Handler {
	return safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
		return safehttp.NotWritten()
	})
}

// NetworkInterfaces a directory of network interfaces. For each network interface the following information is available:
//
//	access-configs/
//	  external-ip
//	  type
//	dns-servers
//	forwarded-ips/
//	gateway
//	ip
//	ip-aliases/
//	mac
//	mtu
//	network
//	subnetmask
//	target-instance-ips
//
// For more information about network interfaces, see Multiple network interfaces overview.
func (h *InstanceHandler) NetworkInterfaces() safehttp.Handler {
	return safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
		return safehttp.NotWritten()
	})
}

// Preempted a boolean value that indicates whether a VM is about to be preempted.
func (h *InstanceHandler) Preempted() safehttp.Handler {
	return safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
		return safehttp.NotWritten()
	})
}

func (h *InstanceHandler) RemainingCPUTime() safehttp.Handler {
	return safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
		return safehttp.NotWritten()
	})
}

// Scheduling sets the scheduling options for the VM.
//
// Scheduling metadata values include the following:
//
//	on-host-maintenance
//
// indicates whether the VM terminates or live migrates during host maintenance.
//
//	automatic-restart
//
// If this value is TRUE, the VM automatically restarts after a maintenance event or crash.
//
//	preemptible
//
// If this value is TRUE, the VM is preemptible. This value is set when you create a VM, and it can't be changed.
//
// For more information about scheduling options, see Setting instance availability policies.
func (h *InstanceHandler) Scheduling() safehttp.Handler {
	return safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
		return safehttp.NotWritten()
	})
}

const (
	// EnvGoogleApplicationCredentials environment variable name for overrides application default credentials JSON path.
	EnvGoogleApplicationCredentials = "GOOGLE_APPLICATION_CREDENTIALS"

	// EnvGoogleAccountEmail environment variable name for overrides service account email address.
	EnvGoogleAccountEmail = "GOOGLE_ACCOUNT_EMAIL"
)

var serviceAccountsEndpoints = []string{
	"aliases",
	"email",
	"identity",
	"scopes",
	"token",
}

// ServiceAccounts a directory of service accounts associated with the VM. For each service account, the following information is available:
//
//	aliases
//
// The service accounts alias.
//
//	email
//
// The email address for the service account.
//
//	identity
//
// A JSON Web Token that is unique to the VM. You must include the audience parameter in your request for this VM metadata value. For example, "?audience=http://www.example.com".
//
// For information about how to request and verify instance identity tokens, see Verifying the identity of instances.
//
//	scopes
//
// The access scopes assigned to the service account.
//
//	token
//
// The OAuth2 access token that can be used to authenticate applications.
//
// For information about access tokens, see Authenticating applications directly with access tokens.
//
// For more information about service accounts, see Creating and enabling service accounts for instances.
func (h *InstanceHandler) ServiceAccounts() safehttp.Handler {
	handler := safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
		url := r.URL()
		q, err := url.Query()
		if err != nil {
			return w.WriteError(safehttp.StatusInternalServerError)
		}
		queries := strings.Split(q.String("scopes", ""), ",")

		path := url.Path()
		if path == "" {
			saEndpoints := []string{
				"default/",
			}
			saEmail, err := h.findServiceAccountEmail(queries...)
			if err != nil {
				return w.WriteError(NewStatusError(err, safehttp.StatusBadRequest))
			}
			saEndpoints = append(saEndpoints, saEmail+"/")

			return w.Write(safehtml.HTMLEscaped(strings.Join(saEndpoints, "\n")))
		}

		gsa, attr := pathpkg.Split(path)
		switch attr {
		case "":
			return w.Write(safehtml.HTMLEscaped(strings.Join(serviceAccountsEndpoints, "\n")))

		case "aliases":
			// TODO(zchee): not implemented
			return w.WriteError(safehttp.StatusNotImplemented)

		case "email":
			// TODO(zchee): not implemented
			return w.WriteError(safehttp.StatusNotImplemented)

		case "identity":
			// TODO(zchee): not implemented
			return w.WriteError(safehttp.StatusNotImplemented)
			_ = gsa

		case "scopes":
			// TODO(zchee): not implemented
			return w.WriteError(safehttp.StatusNotImplemented)

		case "token":
			return h.serviceAccountsTokenHandler(w, r, queries...)
		}

		return w.WriteError(safehttp.StatusNotFound)
	})

	return safehttp.StripPrefix("/computeMetadata/v1/instance/service-accounts/", handler)
}

func (h *InstanceHandler) findServiceAccountEmail(scopes ...string) (string, error) {
	saEmail, ok := os.LookupEnv(EnvGoogleAccountEmail)
	if !ok {
		// try to find application default credentials JSON path
		filename, ok := os.LookupEnv(EnvGoogleApplicationCredentials)
		if !ok {
			err := fmt.Errorf("both of %q and %q environment variables is empty",
				EnvGoogleAccountEmail,
				EnvGoogleApplicationCredentials,
			)
			return "", err
		}

		jwtCfg, err := h.jwtConfigFromServiceAccount(filename, scopes...)
		if err != nil {
			return "", err
		}
		saEmail = jwtCfg.Email
	}

	return saEmail, nil
}

func (h *InstanceHandler) jwtConfigFromServiceAccount(filenname string, scopes ...string) (*jwt.Config, error) {
	data, err := os.ReadFile(filenname)
	if err != nil {
		if os.IsNotExist(err) {
			err = fs.ErrNotExist // overwrite underlying error to fs.ErrNotExist
			return nil, fmt.Errorf("%s %w", filenname, err)
		}

		return nil, fmt.Errorf("could not read %s file: %w", filenname, err)
	}

	jwtCfg, err := google.JWTConfigFromJSON(data, scopes...)
	if err != nil {
		return nil, fmt.Errorf("could not get jwt configuration: %w", err)
	}

	return jwtCfg, nil
}

// TokenResponse represents a JSON response of service account token.
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func (h *InstanceHandler) serviceAccountsTokenHandler(w safehttp.ResponseWriter, r *safehttp.IncomingRequest, scopes ...string) safehttp.Result {
	now := time.Now().In(time.UTC) // for calculate tokne expires

	creds, err := google.FindDefaultCredentialsWithParams(r.Context(), google.CredentialsParams{
		Scopes: scopes,
	})
	if err != nil {
		return w.Write(NewStatusError(err, safehttp.StatusInternalServerError))
	}

	tok, err := creds.TokenSource.Token()
	if err != nil {
		return w.WriteError(safehttp.StatusInternalServerError)
	}

	expiresIn := tok.Expiry.In(time.UTC).Sub(now).Round(time.Second).Seconds()
	resp := TokenResponse{
		AccessToken: tok.AccessToken,
		ExpiresIn:   *(*int)(unsafe.Pointer(&expiresIn)),
		TokenType:   tok.TokenType,
	}

	return WriteJSON(w, &resp)
}

// Tags lists any network tags associated with the VM.
//
// For more information about network tags, see Configuring network tags.
func (h *InstanceHandler) Tags() safehttp.Handler {
	return safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
		return safehttp.NotWritten()
	})
}

func (h *InstanceHandler) VirtualClock() safehttp.Handler {
	return safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
		return safehttp.NotWritten()
	})
}

// Zone is the zone where this VM is located.
//
// This value has the following format: projects/PROJECT_NUM/zones/ZONE.
func (h *InstanceHandler) Zone() safehttp.Handler {
	return safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
		return safehttp.NotWritten()
	})
}
