package main

import (
	kog "github.com/Edgeworx/octant-iofog-plugin/internal/apis/iofog/v1"
	iofogplugin "github.com/Edgeworx/octant-iofog-plugin/internal/plugin"
	"github.com/Edgeworx/octant-iofog-plugin/internal/plugin/handlers"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"log"

	"github.com/vmware-tanzu/octant/pkg/navigation"
	"github.com/vmware-tanzu/octant/pkg/plugin"
	"github.com/vmware-tanzu/octant/pkg/plugin/service"
)

var (
	namespace  = "k8s"
	pluginName = "ioFog"
)

func main() {
	log.SetPrefix("")

	iofogplugin.GetKogDeployment(namespace)

	kogGVK := schema.GroupVersionKind{Version: kog.SchemeGroupVersion.Version, Kind: "Kog"}

	capabilities := &plugin.Capabilities{
		SupportsPrinterConfig: []schema.GroupVersionKind{kogGVK},
		SupportsTab:           []schema.GroupVersionKind{kogGVK},
		SupportsObjectStatus:  []schema.GroupVersionKind{kogGVK},
		IsModule:              true,
	}

	options := []service.PluginOption{
		service.WithNavigation(handleNavigation, initRoutes),
	}

	p, err := service.Register(pluginName, "ioFog plugin", capabilities, options...)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("ioFog plugin is starting")
	p.Serve()
}

func handleNavigation(request *service.NavigationRequest) (navigation.Navigation, error) {
	iofog := navigation.Navigation{
		Title:    "ioFog",
		Path:     request.GeneratePath(),
		Children: []navigation.Navigation{},
		IconName: "cloud",
	}

	deployment, err := iofogplugin.GetKogDeployment(namespace)
	if err != nil {
		log.Print("Unable to get ioFog ECN", err)
		return iofog, nil
	}
	iofogplugin.ECN = deployment

	controllers := navigation.Navigation{
		Title:    "Control Plane",
		Path:     request.GeneratePath("controllers"),
		IconName: "folder",
	}

	agents := navigation.Navigation{
		Title:    "Agents",
		Path:     request.GeneratePath("agents"),
		IconName: "folder",
		Children: []navigation.Navigation{},
	}
	for uuid, agent := range deployment.Agents {
		agents.Children = append(agents.Children, navigation.Navigation{
			Title:    agent.Name,
			Path:     request.GeneratePath("agents", uuid),
			IconName: "cloud",
		})
	}

	connectors := navigation.Navigation{
		Title:    "Connectors",
		Path:     request.GeneratePath("connectors"),
		IconName: "folder",
		Children: []navigation.Navigation{},
	}
	for name, connector := range deployment.Connectors {
		connectors.Children = append(connectors.Children, navigation.Navigation{
			Title:    connector.Name,
			Path:     request.GeneratePath("connectors", name),
			IconName: "cloud",
		})
	}

	applications := navigation.Navigation{
		Title:    "Applications",
		Path:     request.GeneratePath("applications"),
		IconName: "folder",
		Children: []navigation.Navigation{},
	}
	for id, flow := range deployment.Flows {
		applications.Children = append(applications.Children, navigation.Navigation{
			Title:    flow.Name,
			Path:     request.GeneratePath("applications", id),
			IconName: "folder",
		})
	}

	iofog.Children = []navigation.Navigation{
		controllers,
		connectors,
		agents,
		applications,
	}

	return iofog, nil
}

func initRoutes(router *service.Router) {
	router.HandleFunc("/agents", handlers.AgentHandler(true))
	router.HandleFunc("/agents/*", handlers.AgentHandler(false))

	router.HandleFunc("/connectors", handlers.ConnectorHandler(true))
	router.HandleFunc("/connectors/*", handlers.ConnectorHandler(false))

	router.HandleFunc("/controllers", handlers.ControllerHandler(true))

	router.HandleFunc("/applications", handlers.ApplicationHandler(true))
	router.HandleFunc("/applications/*", handlers.ApplicationHandler(false))
}
