/*
 *  *******************************************************************************
 *  * Copyright (c) 2019 Edgeworx, Inc.
 *  *
 *  * This program and the accompanying materials are made available under the
 *  * terms of the Eclipse Public License v. 2.0 which is available at
 *  * http://www.eclipse.org/legal/epl-2.0
 *  *
 *  * SPDX-License-Identifier: EPL-2.0
 *  *******************************************************************************
 *
 */

package handlers

import (
	"github.com/Edgeworx/octant-iofog-plugin/internal/plugin"
	"github.com/eclipse-iofog/iofog-go-sdk/pkg/client"
	"github.com/vmware-tanzu/octant/pkg/plugin/service"
	"github.com/vmware-tanzu/octant/pkg/view/component"
	"path"
	"strconv"
	"strings"
)

func ApplicationHandler(root bool) func(request *service.Request) (component.ContentResponse, error) {
	return func(request *service.Request) (component.ContentResponse, error) {
		if plugin.ECN == nil {
			return component.ContentResponse{}, nil
		}

		if root {
			contentResponse := component.NewContentResponse(component.TitleFromString("Applications"))
			contentResponse.Add(generateApplicationContent())
			return *contentResponse, nil
		}

		_, appId := path.Split(request.Path)
		app := plugin.ECN.Flows[appId]
		microservices := plugin.ECN.MicroservicesPerFlow[appId]

		contentResponse := component.NewContentResponse(component.TitleFromString(app.Name))
		contentResponse.Add(generateMicroserviceSummary(microservices))

		for _, ms := range microservices {
			contentResponse.Add(generateMicroserviceTab(ms))
		}
		return *contentResponse, nil
	}
}

func generateApplicationContent() component.Component {
	tableColumns := []component.TableCol{
		{
			Name: "Application Name",
		},
		{
			Name: "Is Activated",
		},
	}
	var tableRows []component.TableRow
	for id, app := range plugin.ECN.Flows {
		tableRows = append(tableRows, component.TableRow{
			"Application Name": component.NewLink(app.Name, app.Name, "/ioFog/applications/"+id),
			"Is Activated":     component.NewText(strconv.FormatBool(app.IsActivated)),
		})
	}
	table := component.NewTableWithRows("Applications", "", tableColumns, tableRows)
	table.Sort("Application Name", false)

	return table
}

func generateMicroserviceSummary(microservices map[string]client.MicroserviceInfo) component.Component {
	tableColumns := []component.TableCol{
		{
			Name: "Microservice Name",
		},
		{
			Name: "Status",
		},
	}
	var tableRows []component.TableRow
	for _, ms := range microservices {
		tableRows = append(tableRows, component.TableRow{
			"Microservice Name": component.NewText(ms.Name),
			"Status":            component.NewText(ms.Status.Status),
		})
	}
	table := component.NewTableWithRows("Microservices", "", tableColumns, tableRows)
	table.Sort("Microservice Name", false)

	return table
}

func generateMicroserviceTab(ms client.MicroserviceInfo) component.Component {
	var volumes []string
	for _, volume := range ms.Volumes {
		volumes = append(volumes, volume.HostDestination+":"+volume.ContainerDestination+":"+volume.AccessMode)
	}

	return plugin.SummaryHelper(ms.Name, []plugin.SectionsContent{
		{
			Label: "Status",
			Value: ms.Status.Status,
		},
		{
			Label: "Agent",
			Value: plugin.ECN.Agents[ms.AgentUUID].Name,
		},
		{
			Label: "Config",
			Value: ms.Config,
		},
		{
			Label: "Routes",
			Value: strings.Join(ms.Routes, ", "),
		},
		{
			Label: "Volumes",
			Value: strings.Join(volumes, ", "),
		},
	})
}
