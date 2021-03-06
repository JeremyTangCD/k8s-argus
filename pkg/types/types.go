package types

import (
	"github.com/logicmonitor/k8s-argus/pkg/config"
	lm "github.com/logicmonitor/lm-sdk-go"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
)

// Base is a struct for embedding
type Base struct {
	LMClient  *lm.DefaultApi
	K8sClient *kubernetes.Clientset
	Config    *config.Config
}

// Watcher is the LogicMonitor Watcher interface.
type Watcher interface {
	Resource() string
	ObjType() runtime.Object
	AddFunc() func(obj interface{})
	DeleteFunc() func(obj interface{})
	UpdateFunc() func(oldObj, newObj interface{})
}

// DeviceManager is an interface that describes how resources in Kubernetes
// are mapped into LogicMonitor as devices.
type DeviceManager interface {
	DeviceMapper
	DeviceBuilder
}

// DeviceMapper is the interface responsible for mapping a Kubernetes resource to
// a LogicMonitor device.
type DeviceMapper interface {
	// Config returns the Argus config.
	Config() *config.Config
	// FindByDisplayName searches for a device by it's display name. It will return a device if and only if
	// one device was found, and return nil otherwise.
	FindByDisplayName(string) (*lm.RestDevice, error)
	// Add adds a device to a LogicMonitor account.
	Add(...DeviceOption) (*lm.RestDevice, error)
	// UpdateAndReplaceByID updates a device using the 'replace' OpType.
	UpdateAndReplaceByID(int32, ...DeviceOption) (*lm.RestDevice, error)
	// UpdateAndReplaceByDisplayName updates a device using the 'replace' OpType if and onlt if it does not already exist.
	UpdateAndReplaceByDisplayName(string, ...DeviceOption) (*lm.RestDevice, error)
	// UpdateAndReplaceFieldByID updates a device using the 'replace' OpType for a
	// specific field of a device.
	UpdateAndReplaceFieldByID(int32, string, ...DeviceOption) (*lm.RestDevice, error)
	// UpdateAndReplaceFieldByDisplayName updates a device using the 'replace' OpType for a
	// specific field of a device.
	UpdateAndReplaceFieldByDisplayName(string, string, ...DeviceOption) (*lm.RestDevice, error)
	// DeleteByID deletes a device by device ID.
	DeleteByID(int32) error
	// DeleteByDisplayName deletes a device by device display name.
	DeleteByDisplayName(string) error
}

// DeviceOption is the function definition for the functional options pattern.
type DeviceOption func(*lm.RestDevice)

// DeviceBuilder is the interface responsible for building a device struct.
type DeviceBuilder interface {
	// Name sets the device name.
	Name(string) DeviceOption
	// DisplayName sets the device name.
	DisplayName(string) DeviceOption
	// CollectorID sets the preferred collector ID for the device.
	CollectorID(int32) DeviceOption
	// SystemCategories sets the system.categories property on the device.
	SystemCategories(string) DeviceOption
	// ResourceLabels sets custom properties for the device
	ResourceLabels(map[string]string) DeviceOption
	// Auto adds an auto property to the device.
	Auto(string, string) DeviceOption
	// System adds a system property to the device.
	System(string, string) DeviceOption
	// System adds a custom property to the device.
	Custom(string, string) DeviceOption
}
