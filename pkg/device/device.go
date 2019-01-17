package device

import (
	"context"
	"fmt"

	"github.com/logicmonitor/k8s-argus/pkg/config"
	"github.com/logicmonitor/k8s-argus/pkg/constants"
	"github.com/logicmonitor/k8s-argus/pkg/device/builder"
	"github.com/logicmonitor/k8s-argus/pkg/types"
	"github.com/logicmonitor/k8s-argus/pkg/utilities"
	"github.com/logicmonitor/k8s-collectorset-controller/api"
	lm "github.com/logicmonitor/lm-sdk-go"
	log "github.com/sirupsen/logrus"
)

// Manager implements types.DeviceManager
type Manager struct {
	*types.Base
	*builder.Builder
	ControllerClient api.CollectorSetControllerClient
}

func buildDevice(c *config.Config, client api.CollectorSetControllerClient, options ...types.DeviceOption) *lm.RestDevice {
	device := &lm.RestDevice{
		CustomProperties: []lm.NameAndValue{
			{
				Name:  constants.K8sClusterNamePropertyKey,
				Value: c.ClusterName,
			},
		},
		DisableAlerting: c.DisableAlerting,
		HostGroupIds:    "1",
		DeviceType:      constants.K8sDeviceType,
	}

	for _, option := range options {
		option(device)
	}

	reply, err := client.CollectorID(context.Background(), &api.CollectorIDRequest{})
	if err != nil {
		log.Errorf("Failed to get collector ID: %v", err)
	} else {
		log.Infof("Using collector ID %d for %q", reply.Id, device.DisplayName)
		device.PreferredCollectorId = reply.Id
	}

	return device
}

// checkAndUpdateExistDevice tries to find and update the devices which needs to be changed
func (m *Manager) checkAndUpdateExistingDevice(device *lm.RestDevice) (*lm.RestDevice, error) {
	oldDevice, err := m.FindByDisplayName(device.DisplayName)
	if err != nil {
		return nil, err
	}
	if oldDevice == nil {
		return nil, fmt.Errorf("can not find the device: %s", device.DisplayName)
	}

	// the device which is not changed will be ignored
	if device.Name == oldDevice.Name {
		log.Infof("No changes to device (%s). Ignoring update", device.DisplayName)
		return device, nil
	}

	// the device of the other cluster will be ignored
	oldClusterName := ""
	if oldDevice.CustomProperties != nil && len(oldDevice.CustomProperties) > 0 {
		for _, cp := range oldDevice.CustomProperties {
			if cp.Name == constants.K8sClusterNamePropertyKey {
				oldClusterName = cp.Value
			}
		}
	}
	if oldClusterName != m.Config().ClusterName {
		log.Infof("Device (%s) belongs to a different cluster (%s). Ignoring update", device.DisplayName, oldClusterName)
		return device, nil
	}

	newDevice, err := m.updateAndReplace(oldDevice.Id, device)
	if err != nil {
		return nil, err
	}
	log.Infof("Finished updating the device: %s", newDevice.DisplayName)
	return newDevice, nil
}

func (m *Manager) updateAndReplace(id int32, device *lm.RestDevice) (*lm.RestDevice, error) {

	restResponse, apiResponse, err := m.LMClient.UpdateDevice(*device, id, "replace")
	if _err := utilities.CheckAllErrors(restResponse, apiResponse, err); _err != nil {
		return nil, _err
	}
	log.Debugf("%#v", restResponse)

	return &restResponse.Data, nil
}

// FindByDisplayName implements types.DeviceManager.
func (m *Manager) FindByDisplayName(name string) (*lm.RestDevice, error) {
	filter := fmt.Sprintf("displayName:%s", name)
	restResponse, apiResponse, err := m.LMClient.GetDeviceList("", -1, 0, filter)
	if _err := utilities.CheckAllErrors(restResponse, apiResponse, err); _err != nil {
		return nil, _err
	}
	log.Debugf("%#v", restResponse)
	if restResponse.Data.Total == 1 {
		return &restResponse.Data.Items[0], nil
	}

	return nil, nil
}

// Add implements types.DeviceManager.
func (m *Manager) Add(options ...types.DeviceOption) (*lm.RestDevice, error) {
	device := buildDevice(m.Config(), m.ControllerClient, options...)
	log.Debugf("%#v", device)

	restResponse, apiResponse, err := m.LMClient.AddDevice(*device, false)
	if _err := utilities.CheckAllErrors(restResponse, apiResponse, err); _err != nil {
		if restResponse != nil && restResponse.Status == 600 {
			log.Infof("Check and Update the existing device: %s", device.DisplayName)
			newDevice, err := m.checkAndUpdateExistingDevice(device)
			if err != nil {
				return nil, err
			}
			return newDevice, nil
		}

		return nil, _err
	}
	log.Debugf("%#v", restResponse)

	return &restResponse.Data, nil
}

// UpdateAndReplaceByID implements types.DeviceManager.
func (m *Manager) UpdateAndReplaceByID(id int32, options ...types.DeviceOption) (*lm.RestDevice, error) {
	device := buildDevice(m.Config(), m.ControllerClient, options...)
	log.Debugf("%#v", device)

	return m.updateAndReplace(id, device)
}

// UpdateAndReplaceByDisplayName implements types.DeviceManager.
func (m *Manager) UpdateAndReplaceByDisplayName(name string, options ...types.DeviceOption) (*lm.RestDevice, error) {
	d, err := m.FindByDisplayName(name)
	if err != nil {
		return nil, err
	}

	if d == nil {
		log.Warnf("Could not find device %q", name)
		return nil, nil
	}

	// Update the device.
	device, err := m.UpdateAndReplaceByID(d.Id, options...)
	if err != nil {
		return nil, err
	}

	return device, nil
}

// UpdateAndReplaceFieldByID implements types.DeviceManager.
func (m *Manager) UpdateAndReplaceFieldByID(id int32, field string, options ...types.DeviceOption) (*lm.RestDevice, error) {
	device := buildDevice(m.Config(), m.ControllerClient, options...)
	log.Debugf("%#v", device)

	restResponse, apiResponse, err := m.LMClient.PatchDeviceById(*device, id, "replace", field)
	if _err := utilities.CheckAllErrors(restResponse, apiResponse, err); _err != nil {
		return nil, _err
	}
	log.Debugf("%#v", restResponse)

	return &restResponse.Data, nil
}

// UpdateAndReplaceFieldByDisplayName implements types.DeviceManager.
func (m *Manager) UpdateAndReplaceFieldByDisplayName(name string, field string, options ...types.DeviceOption) (*lm.RestDevice, error) {
	d, err := m.FindByDisplayName(name)
	if err != nil {
		return nil, err
	}

	if d == nil {
		log.Infof("Could not find device %q", name)
		return nil, nil
	}

	// Update the device.
	device, err := m.UpdateAndReplaceFieldByID(d.Id, field, options...)
	if err != nil {
		return nil, err
	}

	return device, nil
}

// DeleteByID implements types.DeviceManager.
func (m *Manager) DeleteByID(id int32) error {
	restResponse, apiResponse, err := m.LMClient.DeleteDevice(id)
	return utilities.CheckAllErrors(restResponse, apiResponse, err)
}

// DeleteByDisplayName implements types.DeviceManager.
func (m *Manager) DeleteByDisplayName(name string) error {
	d, err := m.FindByDisplayName(name)
	if err != nil {
		return err
	}

	// TODO: Should this return an error?
	if d == nil {
		log.Infof("Could not find device %q", name)
		return nil
	}

	restResponse, apiResponse, err := m.LMClient.DeleteDevice(d.Id)
	return utilities.CheckAllErrors(restResponse, apiResponse, err)
}

// Config implements types.DeviceManager.
func (m *Manager) Config() *config.Config {
	return m.Base.Config
}

// GetListByGroupID implements getting all the devices belongs to the group directly
func (m *Manager) GetListByGroupID(groupID int32) ([]lm.RestDevice, error) {
	restResponse, apiResponse, err := m.LMClient.GetImmediateDeviceListByDeviceGroupId(groupID, "id,name,displayName,customProperties", -1, 0, "")
	if _err := utilities.CheckAllErrors(restResponse, apiResponse, err); _err != nil {
		return nil, _err
	}
	log.Debugf("%#v", restResponse)
	return restResponse.Data.Items, nil
}
