/*
Package systemctl provides a wrapper to control systemd services.

example:

	nm, err := systemctl.Unit("NetworkManager.service")

	if err != nil {
	    // Handle error
	}

	if nm.IsActive() {
		// Handle active service
	}

	err = nm.Start()
	if err != nil {
	    // Handle error
	}

*/
package systemctl

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

/*
Service provides a struct to store the unit's properties.
WILL BE EXPANDED LATER!
*/
type Service struct {
	Name    string
	Active  bool
	Enabled bool
}

/*
unitExist check whether the given service is exist.
*/
func unitExist(name string) (bool, error) {

	unitPaths := []string{
		"/usr/lib/systemd/system/",
		"/etc/systemd/system/",
		"/usr/local/lib/systemd/system/",
		"/etc/systemd/user/",
		"/etc/systemd/system.control/",
		"/run/systemd/system.control/",
		"/run/systemd/transient/",
		"/run/systemd/generator.early/",
		"/etc/systemd/systemd.attached/",
		"/run/systemd/system/",
		"/run/systemd/systemd.attached/",
		"/run/systemd/generator/",
		"/lib/systemd/system/",
		"/run/systemd/generator.late/",
		"/usr/lib/systemd/user/"}

	for _, unitPath := range unitPaths {

		if _, err := os.Stat(unitPath); os.IsNotExist(err) {
			continue
		}

		files, err := ioutil.ReadDir(unitPath)

		if err != nil {
			return false, err
		}

		for _, file := range files {
			if file.Name() == name {
				return true, nil
			}
		}
	}

	return false, nil
}

/*
Unit gives back Service.
It checks whether a service with name exist.
The struct saves the actual state when the function is called.
*/
func Unit(name string) (Service, error) {

	if exist, err := unitExist(name); err != nil {
		return Service{}, err
	} else if exist != true {
		return Service{}, fmt.Errorf("unit not exist: %s", name)
	}

	active, err := IsActive(name)

	if err != nil {
		return Service{}, err
	}

	enabled, err := IsEnabled(name)

	if err != nil {
		return Service{}, err
	}

	return Service{Name: name, Active: active, Enabled: enabled}, nil
}

/*
IsActive checks if the given service is running.
Returns true if the the given service is active, returns false otherwise.
*/
func IsActive(name string) (bool, error) {

	//if exist, err := unitExist(name); err

	output, err := exec.Command("/usr/bin/systemctl", "is-active", name).CombinedOutput()

	if err != nil {
		return false, fmt.Errorf("failed to run systemctl: %s %s", output, err)
	}

	switch string(output) {
	case "active\n":
		return true, nil
	case "inactive\n":
		return false, nil
	default:
		return false, fmt.Errorf("invalid response: %s", string(output))
	}
}

/*
IsEnabled check if the given service is enabled in systemd.
Returns true if the the given service is enabled.
*/
func IsEnabled(name string) (bool, error) {

	output, err := exec.Command("/usr/bin/systemctl", "is-enabled", name).CombinedOutput()

	if err != nil {
		return false, fmt.Errorf("failed to run systemctl: %s %s", output, err)
	}

	switch string(output) {
	case "enabled\n":
		return true, nil
	case "disabled\n":
		return false, nil
	default:
		return false, fmt.Errorf("invalid response: %s", string(output))
	}
}

/*
Enable function enables the given service in systemd.
*/
func (s *Service) Enable() error {

	output, err := exec.Command("/usr/bin/systemctl", "enable", s.Name).CombinedOutput()

	if err != nil {
		return fmt.Errorf("%s %s", output, err)
	}

	return nil
}

/*
Disable function disable the given service in systemd.
*/
func (s *Service) Disable() error {

	output, err := exec.Command("/usr/bin/systemctl", "disable", s.Name).CombinedOutput()

	if err != nil {
		return fmt.Errorf("%s %s", output, err)
	}

	return nil
}

/*
Start function start the given service with systemctl.
*/
func (s *Service) Start() error {

	output, err := exec.Command("/usr/bin/systemctl", "start", s.Name).CombinedOutput()

	if err != nil {
		return fmt.Errorf("%s %s", output, err)
	}

	return nil
}

/*
Stop function is stop the given service with systemctl.
*/
func (s *Service) Stop() error {

	output, err := exec.Command("/usr/bin/systemctl", "stop", s.Name).CombinedOutput()

	if err != nil {
		return fmt.Errorf("%s %s", output, err)
	}

	return nil
}

/*
Restart function restart the given service with systemctl.
*/
func (s *Service) Restart() error {

	output, err := exec.Command("/usr/bin/systemctl", "restart", s.Name).CombinedOutput()

	if err != nil {
		return fmt.Errorf("%s %s", output, err)
	}

	return nil
}

/*
Reload function reload the given service with systemctl.
*/
func (s *Service) Reload() error {

	output, err := exec.Command("/usr/bin/systemctl", "reload", s.Name).CombinedOutput()

	if err != nil {
		return fmt.Errorf("%s %s", output, err)
	}

	return nil
}

/*
EnableService function enables the given service in systemd.
*/
func EnableService(name string) error {

	if exist, err := unitExist(name); err != nil {
		return err
	} else if exist != true {
		return fmt.Errorf("unit not exist: %s", name)
	}

	output, err := exec.Command("/usr/bin/systemctl", "enable", name).CombinedOutput()

	if err != nil {
		return fmt.Errorf("%s %s", output, err)
	}

	return nil
}

/*
DisableService function disable the given service in systemd.
*/
func DisableService(name string) error {

	if exist, err := unitExist(name); err != nil {
		return err
	} else if exist != true {
		return fmt.Errorf("unit not exist: %s", name)
	}

	output, err := exec.Command("/usr/bin/systemctl", "disable", name).CombinedOutput()

	if err != nil {
		return fmt.Errorf("%s %s", output, err)
	}

	return nil
}

/*
StartService function start the given service with systemctl.
*/
func StartService(name string) error {

	if exist, err := unitExist(name); err != nil {
		return err
	} else if exist != true {
		return fmt.Errorf("unit not exist: %s", name)
	}

	output, err := exec.Command("/usr/bin/systemctl", "start", name).CombinedOutput()

	if err != nil {
		return fmt.Errorf("%s %s", output, err)
	}

	return nil
}

/*
StopService function is stop the given service with systemctl.
*/
func StopService(name string) error {

	if exist, err := unitExist(name); err != nil {
		return err
	} else if exist != true {
		return fmt.Errorf("unit not exist: %s", name)
	}

	output, err := exec.Command("/usr/bin/systemctl", "stop", name).CombinedOutput()

	if err != nil {
		return fmt.Errorf("%s %s", output, err)
	}

	return nil
}

/*
RestartService function restart the given service with systemctl.
*/
func RestartService(name string) error {

	if exist, err := unitExist(name); err != nil {
		return err
	} else if exist != true {
		return fmt.Errorf("unit not exist: %s", name)
	}

	output, err := exec.Command("/usr/bin/systemctl", "restart", name).CombinedOutput()

	if err != nil {
		return fmt.Errorf("%s %s", output, err)
	}

	return nil
}

/*
ReloadService function reload the given service with systemctl.
*/
func ReloadService(name string) error {

	if exist, err := unitExist(name); err != nil {
		return err
	} else if exist != true {
		return fmt.Errorf("unit not exist: %s", name)
	}

	output, err := exec.Command("/usr/bin/systemctl", "reload", name).CombinedOutput()

	if err != nil {
		return fmt.Errorf("%s %s", output, err)
	}

	return nil
}
