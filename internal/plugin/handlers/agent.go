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
  "github.com/eclipse-iofog/iofog-go-sdk/pkg/client"
  "github.com/eclipse-iofog/octant-plugin/internal/plugin"
  "github.com/vmware-tanzu/octant/pkg/plugin/service"
  "github.com/vmware-tanzu/octant/pkg/view/component"
  "path"
  "time"
)

func AgentHandler(root bool) func(request *service.Request) (component.ContentResponse, error) {
  return func(request *service.Request) (component.ContentResponse, error) {
    if plugin.ECN == nil {
      return component.ContentResponse{}, nil
    }

    if root {
      contentResponse := component.NewContentResponse(component.TitleFromString("Agents"))
      contentResponse.Add(generateAgentContent())
      return *contentResponse, nil
    }

    _, agentId := path.Split(request.Path)
    agent := plugin.ECN.Agents[agentId]

    contentResponse := component.NewContentResponse(component.TitleFromString(agent.Name))
    contentResponse.Add(generateAgentTab(agent))
    return *contentResponse, nil
  }
}

func generateAgentContent() component.Component {
  tableColumns := []component.TableCol{
    {
      Name:     "Agent Name",
    },
    {
      Name:     "Status",
    },
  }
  var tableRows []component.TableRow
  for uuid, agent := range plugin.ECN.Agents {
    tableRows = append(tableRows, component.TableRow{
      "Agent Name": component.NewLink(agent.Name, agent.Name, "/ioFog/agents/" + uuid),
      "Status": component.NewText(agent.DaemonStatus),
    })
  }
  table := component.NewTableWithRows("Agents", "", tableColumns, tableRows)
  table.Sort("Agent Name", false)

  return table
}

func generateAgentTab(agent client.AgentInfo) component.Component {
  elapsed, _ := plugin.ElapsedRFC(agent.CreatedTimeRFC3339, plugin.NowRFC())
  return plugin.SummaryHelper(agent.Name, []plugin.SectionsContent{
    {
      Label: "Status",
      Value: agent.DaemonStatus,
    },
    {
      Label: "Age",
      Value: elapsed,
    },
    {
      Label: "Uptime",
      Value: plugin.FormatDuration(time.Duration(agent.UptimeMs) * time.Millisecond),
    },
    {
      Label: "IP",
      Value: agent.IPAddressExternal,
    },
    {
      Label: "Version",
      Value: agent.Version,
    },
  })
}
