package pilote

/*
 * REST services for manual driving of the robot.
 * Gives access to feedback and entry point for driving
 * A swagger-ui is included => http://localhost:8080/apidocs/?url=http://localhost:8080/apidocs.json
 */

import (
	"log"
	"net/http"

	"robotoby/application"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
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

// PUT http://1.2.3.4:8080/pilote/command
func addCommand(request *restful.Request, response *restful.Response) {
	application.Debug("Adding new command to run")
	command := Command{}
	err := request.ReadEntity(&command)
	if err == nil {
		// Always stop auto pilote if enabled
		stopAutoPilote()
		command.Process()
		response.WriteHeaderAndEntity(http.StatusCreated, command)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

// PUT http://1.2.3.4:8080/pilote/chain
func addChain(request *restful.Request, response *restful.Response) {
	application.Debug("Adding new chain to run")
	chain := ChainedCommand{}
	err := request.ReadEntity(&chain)
	if err == nil {
		// Always stop auto pilote if enabled
		stopAutoPilote()
		chain.Process()
		response.WriteHeaderAndEntity(http.StatusCreated, chain)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

// POST http://1.2.3.4:8080/pilote/auto
func startAuto(request *restful.Request, response *restful.Response) {
	application.Debug("Starting auto-pilote mode")
	startFullAutoPilote()
}

// DELETE http://1.2.3.4:8080/pilote/auto
func stopAuto(request *restful.Request, response *restful.Response) {
	application.Debug("Interrupt auto-pilote mode")
	stopAutoPilote()
}

// Routes for state access
//
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

// Routes for Pilote process
//
func piloteWebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/pilote").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	tags := []string{"pilote"}

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
		Doc("Start auto-pilote mode").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", struct{}{}))

	ws.Route(ws.DELETE("auto").To(stopAuto).
		Doc("Stop any auto pilote mode").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", struct{}{}))
	return ws
}

// StartServer : Start the REST service for pilote
func StartServer() {

	restful.DefaultContainer.
		Add(stateWebService()).
		Add(piloteWebService())

	config := restfulspec.Config{
		WebServices: restful.RegisteredWebServices(), // you control what services are visible
		APIPath:     "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))

	// Swagger-ui
	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("./pilote/swagger-ui/dist-v2"))))

	// Pilote
	http.Handle("/control/", http.StripPrefix("/control/", http.FileServer(http.Dir("./pilote/control-ui"))))

	// Optionally, you may need to enable CORS for the UI to work.
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "PUT"},
		CookiesAllowed: false,
		Container:      restful.DefaultContainer}
	restful.DefaultContainer.Filter(cors.Filter)

	ip := application.State.OutBountIP.String()

	application.Info("Get the API using http://" + ip + ":8080/apidocs.json")
	application.Info("Open Swagger UI using http://" + ip + ":8080/apidocs/?url=http://" + ip + ":8080/apidocs.json")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// enrichSwaggerObject : Add swagger def for service
func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "Robot pilote Rest services",
			Description: "Resources for driving a robot",
			Contact: &spec.ContactInfo{
				Name: "elecomte",
				URL:  "https://www.elecomte.com",
			},
			License: &spec.License{
				Name: "MIT",
				URL:  "http://mit.org",
			},
			Version: "0.0.1",
		},
	}
	swo.Tags = []spec.Tag{
		spec.Tag{TagProps: spec.TagProps{
			Name:        "state",
			Description: "Processing state and feedback on robot"}},
		spec.Tag{TagProps: spec.TagProps{
			Name:        "pilote",
			Description: "Managing input commands and driving"}}}
}
