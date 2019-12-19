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
)

func ConnectorHandler(root bool) func(request *service.Request) (component.ContentResponse, error) {
	return func(request *service.Request) (component.ContentResponse, error) {
		if plugin.ECN == nil {
			return component.ContentResponse{}, nil
		}

		if root {
			contentResponse := component.NewContentResponse(component.TitleFromString("Connectors"))
			contentResponse.Add(generateConnectorContent())
			return *contentResponse, nil
		}

		_, connectorId := path.Split(request.Path)
		connector := plugin.ECN.Connectors[connectorId]

		contentResponse := component.NewContentResponse(component.TitleFromString(connector.Name))
		contentResponse.Add(generateConnectorTab(connector))
		return *contentResponse, nil
	}
}

func generateConnectorContent() component.Component {
	tableColumns := []component.TableCol{
		{
			Name: "Connector Name",
		},
		{
			Name: "IP Address",
		},
	}
	var tableRows []component.TableRow
	for uuid, connector := range plugin.ECN.Connectors {
		tableRows = append(tableRows, component.TableRow{
			"Connector Name": component.NewLink(connector.Name, connector.Name, "/ioFog/connectors/"+uuid),
			"IP Address":     component.NewText(connector.IP),
		})
	}
	table := component.NewTableWithRows("Connectors", "", tableColumns, tableRows)
	table.Sort("Connector Name", false)

	return table
}

func generateConnectorTab(connector client.ConnectorInfo) component.Component {
	return plugin.SummaryHelper(connector.Name, []plugin.SectionsContent{
		{
			Label: "IP",
			Value: connector.IP,
		},
		{
			Label: "Domain",
			Value: connector.Domain,
		},
		{
			Label: "DevMode",
			Value: strconv.FormatBool(connector.DevMode),
		},
	})
}
