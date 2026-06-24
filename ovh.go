// Package ovh deploys a togo app to OVHcloud (Public Cloud / OpenStack) by driving the `openstack` CLI.
// Select with deploy.provider=ovh; requires the openstack CLI authenticated.
package ovh

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/togo-framework/deploy"
	"github.com/togo-framework/togo"
)

func init() { deploy.RegisterDriver("ovh", New) }

// New checks the openstack CLI is present.
func New(_ *togo.Kernel) (deploy.Deployer, error) {
	if _, err := exec.LookPath("openstack"); err != nil {
		return nil, fmt.Errorf("deploy-ovh: the %q CLI is required (install + authenticate it)", "openstack")
	}
	return &driver{}, nil
}

type driver struct{}

func run(ctx context.Context, name string, args ...string) (string, error) {
	out, err := exec.CommandContext(ctx, name, args...).CombinedOutput()
	return string(out), err
}

func region(s deploy.Spec, def string) string {
	if s.Region != "" {
		return s.Region
	}
	return def
}

func (d *driver) Provision(ctx context.Context, spec deploy.Spec) (*deploy.Result, error) {
	return d.Deploy(ctx, spec)
}

func (d *driver) Deploy(ctx context.Context, spec deploy.Spec) (*deploy.Result, error) {
	flavor := "b2-7"
	if v, ok := spec.Options["flavor"].(string); ok && v != "" {
		flavor = v
	}
	ud := fmt.Sprintf("#cloud-config\nruncmd:\n  - curl -fsSL https://get.docker.com | sh\n  - docker run -d --restart always -p 80:8080 %s\n", spec.Image)
	out, err := run(ctx, "openstack", "server", "create", "--image", "Ubuntu 22.04", "--flavor", flavor, "--user-data", ud, spec.App)
	if err != nil {
		return nil, fmt.Errorf("openstack server create: %v: %s", err, out)
	}
	return &deploy.Result{Message: "OVH/OpenStack instance creating; app via cloud-init", Raw: map[string]any{"out": out}}, nil
}

func (d *driver) Destroy(ctx context.Context, spec deploy.Spec) error {
	_, err := run(ctx, "openstack", "server", "delete", spec.App)
	return err
}

func (d *driver) Status(ctx context.Context, spec deploy.Spec) (*deploy.Status, error) {
	out, err := run(ctx, "openstack", "server", "show", spec.App, "-f", "value", "-c", "status")
	return &deploy.Status{Healthy: err == nil, Detail: out}, err
}
