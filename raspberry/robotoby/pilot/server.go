package pilot

/*
 * REST services for manual driving of the robot.
 * Gives access to feedback and entry point for driving
 * A swagger-ui is included => http://localhost:8080/apidocs/?url=http://localhost:8080/apidocs.json
 */

import (
	"fmt"
	"github.com/go-openapi/spec"
	"log"
	"net/http"
	"strconv"

	"robotoby/application"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	restful "github.com/emicklei/go-restful/v3"
	"github.com/swaggest/swgui/v5emb"
)

// GET http://1.2.3.4:8080/state/robot
func getCurrentRobotState(request *restful.Request, response *restful.Response) {
	application.Debug("Getting robot state data")
	response.WriteEntity(GetCurrentRobotState())
}

// GET http://1.2.3.4:8080/state/application
func getCurrentApplicationState(request *restful.Request, response *restful.Response) {
	application.Debug("Getting application state data")
	response.WriteEntity(application.State)
}

// PUT http://1.2.3.4:8080/pilot/command
func addCommand(request *restful.Request, response *restful.Response) {
	application.Debug("Adding new command to run")
	command := Command{}
	err := request.ReadEntity(&command)
	if err == nil {
		// Always stop autopilot if enabled
		stopAutoPilot()
		err := command.Process()
		if err != nil {
			log.Printf("Command error : %v\n", err)
			return
		}
		response.WriteHeaderAndEntity(http.StatusCreated, command)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

// PUT http://1.2.3.4:8080/pilot/chain
func addChain(request *restful.Request, response *restful.Response) {
	application.Debug("Adding new chain to run")
	chain := ChainedCommand{}
	err := request.ReadEntity(&chain)
	if err == nil {
		// Always stop autopilot if enabled
		stopAutoPilot()
		err := chain.Process()
		if err != nil {
			log.Printf("Command error : %v\n", err)
			return
		}
		response.WriteHeaderAndEntity(http.StatusCreated, chain)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

// POST http://1.2.3.4:8080/pilot/auto
func startAuto(request *restful.Request, response *restful.Response) {
	application.Debug("Starting auto-pilot mode")
	startFullAutoPilot()
}

// DELETE http://1.2.3.4:8080/pilot/auto
func stopAuto(request *restful.Request, response *restful.Response) {
	application.Debug("Interrupt auto-pilot mode")
	stopAutoPilot()
}

// Routes for state access
func stateWebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/state").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	tags := []string{"state"}

	ws.Route(ws.GET("robot").To(getCurrentRobotState).
		Doc("get state data for robot").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(RobotState{}).
		Returns(200, "OK", RobotState{}))

	ws.Route(ws.GET("application").To(getCurrentApplicationState).
		Doc("get state data for application").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(application.ApplicationState{}).
		Returns(200, "OK", application.ApplicationState{}))

	return ws
}

// Routes for Pilot process
func pilotWebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/pilot").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	tags := []string{"pilot"}

	ws.Route(ws.PUT("command").To(addCommand).
		Doc("add a command to process").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(Command{}). // from the request
		Returns(200, "OK", Command{}))

	ws.Route(ws.PUT("chain").To(addChain).
		Doc("add a chain of commands to process").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(ChainedCommand{}). // from the request
		Returns(200, "OK", ChainedCommand{}))

	ws.Route(ws.POST("auto").To(startAuto).
		Doc("Start auto-pilot mode").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", struct{}{}))

	ws.Route(ws.DELETE("auto").To(stopAuto).
		Doc("Stop any auto pilot mode").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", struct{}{}))
	return ws
}

// StartServer : Start the REST service for pilot
func StartServer() {

	port := strconv.Itoa(application.State.Config.PilotServerPort)

	restful.DefaultContainer.
		Add(stateWebService()).
		Add(pilotWebService())

	ip := application.State.OutBountIP.String()

	config := restfulspec.Config{
		WebServices:                   restful.RegisteredWebServices(), // you control what services are visible
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))

	// Swagger-ui
	http.Handle("/api/docs/", v5emb.New(
		application.State.Config.RobotName,
		fmt.Sprintf("http://%s:%s/apidocs.json", ip, port),
		"/api/docs/",
	))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello World!"))
	})

	// Pilot
	http.Handle("/control/", http.StripPrefix("/control/", http.FileServer(http.Dir("./pilot/control-ui"))))

	// Optionally, you may need to enable CORS for the UI to work.
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "PUT"},
		CookiesAllowed: false,
		Container:      restful.DefaultContainer}
	restful.DefaultContainer.Filter(cors.Filter)

	application.Info(fmt.Sprintf("Get the API using http://%s:%s/apidocs.json", ip, port))
	application.Info(fmt.Sprintf("Open Swagger UI using http://%s:%s/api/docs/", ip, port))
	log.Fatalf("http service terminated : %v", http.ListenAndServe(":"+port, nil))
}

// enrichSwaggerObject : Add swagger def for service
func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "Robot pilot Rest services",
			Description: "Resources for driving a robot",
			Version:     "0.0.1",
		},
	}
	swo.Tags = []spec.Tag{
		{TagProps: spec.TagProps{
			Name:        "state",
			Description: "Processing state and feedback on robot"}},
		{TagProps: spec.TagProps{
			Name:        "pilot",
			Description: "Managing input commands and driving"}}}
}
