module sigs.k8s.io/cluster-api-bootstrap-provider-kubeadm

go 1.12

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/go-logr/logr v0.1.0
	github.com/golang/groupcache v0.0.0-20190129154638-5b532d6fd5ef // indirect
	github.com/google/go-cmp v0.3.1 // indirect
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/onsi/ginkgo v1.10.1
	github.com/onsi/gomega v1.7.0
	github.com/pkg/errors v0.8.1
	github.com/prometheus/procfs v0.0.5 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.4.0 // indirect
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/zap v1.10.0 // indirect
	golang.org/x/crypto v0.0.0-20190911031432-227b76d455e7 // indirect
	golang.org/x/net v0.0.0-20190909003024-a7b16738d86b
	golang.org/x/sys v0.0.0-20190911201528-7ad0cfa0b7b5 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.2.4 // indirect
	k8s.io/api v0.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/cluster-bootstrap v0.0.0-20190516232516-d7d78ab2cfe7
	k8s.io/klog v0.4.0
	sigs.k8s.io/cluster-api v0.2.5
	sigs.k8s.io/controller-runtime v0.3.0
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20191003000013-35e20aa79eb8
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20191003002041-49e3d608220c
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190913080033-27d36303b655
	k8s.io/apiserver => k8s.io/apiserver v0.0.0-20191003001037-3c8b233e046c
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20191003002408-6e42c232ac7d
	k8s.io/client-go => k8s.io/client-go v0.0.0-20191003000419-f68efa97b39e
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20191003003426-b4b1f434fead
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.0.0-20191003003255-c493acd9e2ff
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20190927045949-f81bca4f5e85
	k8s.io/component-base => k8s.io/component-base v0.0.0-20191003000551-f573d376509c
	k8s.io/cri-api => k8s.io/cri-api v0.0.0-20190828162817-608eb1dad4ac
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.0.0-20191003003551-0eecdcdcc049
	k8s.io/helm => github.com/keleustes/helm v1.16.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20191003001317-a019a9d85a86
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.0.0-20191003003129-09316795c0dd
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.0.0-20191003002707-f6b7b0f55cc0
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.0.0-20191003003001-314f0beee0a9
	k8s.io/kubectl => k8s.io/kubectl v0.0.0-20191003004222-1f3c0cd90ca9
	k8s.io/kubelet => k8s.io/kubelet v0.0.0-20191003002833-e367e4712542
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.0.0-20191003003732-7d49cdad1c12
	k8s.io/metrics => k8s.io/metrics v0.0.0-20191003002233-837aead57baf
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.0.0-20191003001538-80f33ca02582
	sigs.k8s.io/cluster-api => github.com/keleustes/cluster-api v1.16.2
	sigs.k8s.io/controller-runtime => github.com/keleustes/controller-runtime v1.16.2
// sigs.k8s.io/kustomize/v3 => github.com/keleustes/kustomize/v3 v3.16.8
)
