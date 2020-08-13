package list

import (
	"context"

	"helm.sh/helm/v3/pkg/release"

	"github.com/gojekfarm/albatross/pkg/helmcli"
	"github.com/gojekfarm/albatross/pkg/helmcli/flags"
)

type service interface {
	List(ctx context.Context, req Request) (Response, error)
}

type Service struct{}

func (s Service) List(ctx context.Context, req Request) (Response, error) {
	listflags := flags.ListFlags{
		GlobalFlags: req.Flags.GlobalFlags,
	}
	lcli := helmcli.NewLister(listflags)
	releases, err := lcli.List(ctx)
	if err != nil {
		return Response{}, err
	}

	respReleases := []Release{}
	for _, release := range releases {
		respReleases = append(respReleases, releaseInfo(release))
	}

	resp := Response{Releases: respReleases}
	return resp, nil
}

func releaseInfo(release *release.Release) Release {
	return Release{
		Name:       release.Name,
		Namespace:  release.Namespace,
		Version:    release.Version,
		Updated:    release.Info.FirstDeployed.Local().Time,
		Status:     release.Info.Status,
		Chart:      release.Chart.ChartFullPath(),
		AppVersion: release.Chart.AppVersion(),
	}
}