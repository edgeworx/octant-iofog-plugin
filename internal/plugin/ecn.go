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

package plugin

import (
	iofogk8s "github.com/Edgeworx/octant-iofog-plugin/internal/k8s"
	iofogclient "github.com/eclipse-iofog/iofog-go-sdk/pkg/client"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"strconv"
)

const lastSystemCatalogItemID int = 3

var ECN *iofogk8s.Deployment

func GetKogDeployment(namespace string) (deployment *iofogk8s.Deployment, err error) {
	kubeConfig := ""
	if home := homedir.HomeDir(); home != "" {
		kubeConfig = filepath.Join(home, ".kube", "config")
	}

	k8s, err := iofogk8s.NewKubernetes(kubeConfig, namespace)
	if err != nil {
		return
	}

	if err = k8s.NamespaceExists(namespace); err != nil {
		return
	}

	kogDeployment, err := k8s.GetKogDeployment()
	if err != nil {
		return
	}
	if kogDeployment == nil {
		return
	}

	controllerIp, err := k8s.GetControllerIp()
	if err != nil {
		return
	}
	kogDeployment.Spec.ControlPlane.LoadBalancerIP = controllerIp

	endpoint, err := k8s.GetControllerEndpoint(controllerIp)
	if err != nil {
		return
	}

	ctrl := iofogclient.New(endpoint)

	endpoint = ctrl.GetEndpoint()

	loginRequest := iofogclient.LoginRequest{
		Email:    kogDeployment.Spec.ControlPlane.IofogUser.Email,
		Password: kogDeployment.Spec.ControlPlane.IofogUser.Password,
	}
	if err = ctrl.Login(loginRequest); err != nil {
		return
	}

	connectorsList, err := ctrl.ListConnectors()
	if err != nil {
		return
	}

	agentsList, err := ctrl.ListAgents()
	if err != nil {
		return
	}

	flowsList, err := ctrl.GetAllFlows()
	if err != nil {
		return
	}

	flows := make(map[string]iofogclient.FlowInfo)
	microservices := make(map[string]map[string]iofogclient.MicroserviceInfo)
	for _, flow := range flowsList.Flows {
		flows[strconv.Itoa(flow.ID)] = flow

		listMsvcs, err := ctrl.GetMicroservicesPerFlow(flow.ID)
		if err != nil {
			return nil, err
		}

		microservices[strconv.Itoa(flow.ID)] = make(map[string]iofogclient.MicroserviceInfo)
		for _, ms := range listMsvcs.Microservices {
			if isSystemMsvc(ms) {
				continue
			}
			microservices[strconv.Itoa(flow.ID)][ms.UUID] = ms
		}
	}

	agents := make(map[string]iofogclient.AgentInfo)
	for _, agent := range agentsList.Agents {
		agents[agent.UUID] = agent
	}

	connectors := make(map[string]iofogclient.ConnectorInfo)
	for _, connector := range connectorsList.Connectors {
		connectors[connector.Name] = connector
	}

	ECN = &iofogk8s.Deployment{
		Agents:               agents,
		Connectors:           connectors,
		Flows:                flows,
		MicroservicesPerFlow: microservices,
		Kog:                  *kogDeployment,
	}

	return ECN, nil
}

func isSystemMsvc(msvc iofogclient.MicroserviceInfo) bool {
	return msvc.CatalogItemID != 0 && msvc.CatalogItemID <= lastSystemCatalogItemID
}
