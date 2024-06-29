package main

import (
	"github.com/knabben/tutorial-istio-sec/magefiles/pkg/kind"
	"github.com/knabben/tutorial-istio-sec/magefiles/writter"
	"github.com/magefile/mage/mg"
	"strings"
)

const (
	CLUSTER_NAME_1 = "watch"
	PATH_1         = "1-watch/specs/"
)

type Demo1 mg.Namespace

// Install create a new Kind Cluster
func (Demo1) Install() error {
	return kind.InstallKind(CLUSTER_NAME_1, PATH_1, false)
}

func (Demo1) Secret() error {
	ExecDocker("apt-get", "update")
	ExecDocker("apt-get", "-y", "install", "curl", "sudo")
	cmd := "curl -L https://storage.googleapis.com/etcd/v3.5.14/etcd-v3.5.14-linux-amd64.tar.gz -o /tmp/etcd-v3.5.14-linux-amd64.tar.gz"
	ExecDocker(strings.Split(cmd, " ")...)
	ExecDocker("tar", "zxvf", "/tmp/etcd-v3.5.14-linux-amd64.tar.gz")
	ExecDocker("sudo", "install", "etcd-v3.5.14-linux-amd64/etcdctl", "/usr/local/bin")
	ExecDocker(
		"etcdctl",
		"--cert=/etc/kubernetes/pki/etcd/server.crt",
		"--key=/etc/kubernetes/pki/etcd/server.key",
		"--cacert=/etc/kubernetes/pki/etcd/ca.crt",
		"--endpoints=https://localhost:2379",
		"get", "/registry/secrets/kube-system/bootstrap-token-abcdef",
		"-w", "fields",
	)
	writter.Kubectl("view-secret", "-n", "kube-system", "bootstrap-token-abcdef", "-a")
	return nil
}
func ExecDocker(args ...string) {
	_ = writter.Docker(append([]string{"exec", "-it", CLUSTER_NAME_1 + "-control-plane"}, args...)...)
}

// Delete cleans up resources from cluster
func (Demo1) Delete() error {
	return kind.DeleteKind(CLUSTER_NAME_1)
}
