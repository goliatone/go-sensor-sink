package device

//DeviceResponse format a device to return
func DeviceResponse(device Device) map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
		"data":   device,
	}
}

//ListDevicesResponse format a list of devices
func ListDevicesResponse(devices []Device) map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
		"data":   devices,
	}
}
