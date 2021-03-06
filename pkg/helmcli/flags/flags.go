package flags

type GlobalFlags struct {
	KubeContext   string `json:"kube_context,omitempty"`
	KubeToken     string `json:"kube_token,omitempty"`
	KubeAPIServer string `json:"kube_apiserver,omitempty"`
	Namespace     string
}

type UpgradeFlags struct {
	DryRun  bool
	Install bool
	Version string
	GlobalFlags
}

type InstallFlags struct {
	DryRun  bool
	Version string
	GlobalFlags
}

type ListFlags struct {
	AllNamespaces bool
	Deployed      bool
	Failed        bool
	Pending       bool
	Uninstalled   bool
	Uninstalling  bool
	GlobalFlags
}
