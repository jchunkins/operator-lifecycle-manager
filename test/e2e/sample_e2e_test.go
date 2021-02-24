package e2e

import (
	"github.com/blang/semver/v4"
	. "github.com/onsi/ginkgo"
	cu "github.com/operator-framework/operator-lifecycle-manager/test/e2e/catalogutil"
)

var _ = Describe("Sample", func() {
	It("SimpleTest", func() {

		operatorImage := "quay.io/cdjohnson/ibm-sample-panamax-operator@sha256:47c9fcef261f3f26dbb6c05f904d11563aae43f61d72e0e4556399f3003c97d0"
		operatorCommand := []string{"ibm-sample-panamax-operator"}
		catalogEntry := []cu.CatalogEntry{
			{Version: semver.MustParse("1.0.0"), ReplacesVersion: "", SkipRange: "", DefaultChannel: "alpha", Channels: []string{"alpha"}, NewIndex: true, PackageName: "testoperatora", OwnedGVKs: cu.A1v1CRDDescription, DependencyGVKs: nil, DependencyPackages: nil, Addmode: cu.SemverSkipPatch, ConfigMap: nil, Secret: nil, CrdVersions: cu.V1CRDVersionV1beta1, OperatorImage: operatorImage, OperatorCommand: operatorCommand},
			{Version: semver.MustParse("1.0.1"), ReplacesVersion: "", SkipRange: "<1.0.1", DefaultChannel: "alpha", Channels: []string{"alpha"}, NewIndex: false, PackageName: "testoperatora", OwnedGVKs: cu.A1v1CRDDescription, DependencyGVKs: nil, DependencyPackages: nil, Addmode: cu.SemverSkipPatch, ConfigMap: nil, Secret: nil, CrdVersions: cu.V1CRDVersionV1beta1, OperatorImage: operatorImage, OperatorCommand: operatorCommand},
			{Version: semver.MustParse("1.0.0"), ReplacesVersion: "", SkipRange: "", DefaultChannel: "alpha", Channels: []string{"alpha"}, NewIndex: false, PackageName: "testoperatorb", OwnedGVKs: cu.B1v1CRDDescription, DependencyGVKs: cu.A1v1CRDDescription, DependencyPackages: nil, Addmode: cu.SemverSkipPatch, ConfigMap: nil, Secret: nil, CrdVersions: cu.V1CRDVersionV1beta1, OperatorImage: operatorImage, OperatorCommand: operatorCommand},
		}
		// TEST FOR DOWNSTREAM REGISTRY

		// stack :=  cu.Stack{
		// 	OpmBinarySourceImage: cu.Downstream4_6,
		// 	CatalogFromImage:     cu.Ubi8,
		// 	CatalogName:          "panamax",
		// 	CatalogTag:           "latest",
		// 	Oc:                   cu.Ocv4_5_0,
		// 	Opmdown:           cu.Opmdownv1_14_3,
		// 	OpmDebug:          true,
		// 	TargetRegistry:    "localhost:5000",
		// 	ContainerCLI:      cu.Docker,
		// 	TargetCatalogType: cu.Registry,
		// }

		// TEST FOR UPSTREAM REGISTRY

		// stack :=  cu.Stack{
		// 	OpmBinarySourceImage: cu.Upstream1_15,
		// 	CatalogFromImage:     cu.Ubi8,
		// 	CatalogName:          "panamax1",
		// 	CatalogTag:           "latest",
		// 	Oc:                   cu.Ocv4_5_0,
		// 	// Opmdown:              cu.Opmdownv1_14_3,
		// 	Opmup:             cu.Opmupv1_15_2,
		// 	OpmDebug:          true,
		// 	TargetRegistry:    "localhost:5000",
		// 	ContainerCLI:      Docker,
		// 	TargetCatalogType: Registry,
		// }

		// TEST FOR DOWNSTREAM IMAGE --- NOTE THIS DOES NOT WORK
		// Hit Error: error copying container directory open /private/var/folders/f_/8tjmq3h93jq1yk1z4k3zrx0w0000gn/T/catalog769363132/index_tmp_766319566/root/.bash_logout: permission denied

		// stack :=  cu.Stack{
		// 	OpmBinarySourceImage: cu.Downstream4_6,
		// 	CatalogFromImage:     cu.Ubi8,
		// 	CatalogName:          "panamax2",
		// 	CatalogTag:           "latest",
		// 	Oc:                   cu.Ocv4_5_0,
		// 	Opmdown:              cu.Opmdownv1_14_3,
		// 	OpmDebug:             true,
		// 	TargetRegistry:       "localhost:5000",
		// 	ContainerCLI:         cu.Docker,
		// 	TargetCatalogType:    cu.Image,
		// }

		// TEST FOR UPSTREAM IMAGE
		stack := cu.Stack{
			OpmBinarySourceImage: cu.Upstream1_15,
			CatalogFromImage:     cu.Ubi8,
			CatalogName:          "panamax3",
			CatalogTag:           "latest",
			Oc:                   cu.Ocv4_5_0,
			// Opmdown:              cu.Opmdownv1_14_3,
			Opmup:             cu.Opmupv1_15_2,
			OpmDebug:          true,
			TargetRegistry:    "localhost:5000",
			ContainerCLI:      cu.Docker,
			TargetCatalogType: cu.Image,
		}

		err := cu.CreateTemporaryCatalog(*toolsBin, catalogEntry, stack)
		_ = err
	})
})
