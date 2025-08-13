// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package vertexai

import (
	"github.com/crossplane/upjet/pkg/config"
)

func init() {
	// import / identifier = "<endpoint> <role> <member>"
	config.TemplatedStringAsIdentifier(
		"google_vertex_ai_endpoint_iam_member",
		"{{ .forProvider.endpoint }} {{ .forProvider.role }} {{ .forProvider.member }}",
	)
}

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	// TRYING CLI APPROACH, NOT FRAMEWORK
	// // Tell Upjet that *this* resource must use the Terraform Plugin Framework
	// p.TerraformPluginFrameworkIncludeList = append(
	// 	p.TerraformPluginFrameworkIncludeList,
	// 	"google_vertex_ai_endpoint_iam_member",
	// )
	p.AddResourceConfigurator("google_vertex_ai_index", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})
	p.AddResourceConfigurator("google_vertex_ai_tensorboard", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})
	p.AddResourceConfigurator("google_vertex_ai_endpoint_iam_member", func(r *config.Resource) {
		r.References["endpoint"] = config.Reference{
			TerraformName: "google_vertex_ai_endpoint",
		}
	})
	// p.AddResourceConfigurator("google_workbench_instance_iam_member", func(r *config.Resource) {
	// 	r.References["name"] = config.Reference{
	// 		TerraformName: "google_workbench_instance",
	// 	}
	// })
}
