// Package main.go

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/hokaccha/go-prettyjson"
	"github.com/integrii/flaggy"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	ctv "github.com/sty-holdings/constant-type-vars-go/v2024"
	"github.com/sty-holdings/nats-connect-go-client/src"
	ncs "github.com/sty-holdings/nats-connect-shared/v2024"
	config "github.com/sty-holdings/sty-shared/v2024/configuration"
	pi "github.com/sty-holdings/sty-shared/v2024/programInfo"
)

//goland:noinspection ALL
const (
	SYNADIA_CLOUD_TOKEN    = "uat_HbusMDRgIc7bUK3cJDBP4AQdJLWOM45mj8R8sH4nfpVVFl0YTRQFXOcLcsjAcNow"
	SYNADIA_CLOUD_BASE_URL = "https://cloud.synadia.com"
)

// Add types to the types.go file

var (
	styhClientId   string
	configFileFQN  string
	generateConfig bool
	password       string
	environment    = "production" // this is the default. For development, use 'development' otherwise 'local'.
	programName    = "nats-connect-go-client"
	secretKey      string
	tempDirectory  string
	testingOn      bool
	username       string
	version        = "9999.9999.9999"
)

func init() {

	appDescription := cases.Title(language.English).String(programName) + " is a driver program to test the NATS Connect." +
		"\nVersion: \n" + ctv.SPACES_FOUR + "- " + version + "\n" +
		"\nConstraints: \n" +
		ctv.SPACES_FOUR + "- When using -c you must pass the fully qualified configuration file name.\n" +
		ctv.SPACES_FOUR + "- There is no console available at this time and all log messages are output to your Log_Directory " +
		"\n\tspecified in the config file or command line below.\n" +
		"\nNotes:\n" +
		ctv.SPACES_FOUR + "None\n" +
		"\nFor more info, see link below:\n"

	// Set your program's name and description.  These appear in help output.
	flaggy.SetName("\n" + programName) // "\n" is added to the start of the name to make the output easier to read.
	flaggy.SetDescription(appDescription)

	// You can disable various things by changing bool on the default parser
	// (or your own parser if you have created one).
	flaggy.DefaultParser.ShowHelpOnUnexpected = true

	// You can set a help prepend or append on the default parser.
	flaggy.DefaultParser.AdditionalHelpPrepend = "https://github.com/styh-dev/albert"

	// Add a flag to the main program (this will be available in all subcommands as well).
	flaggy.String(&configFileFQN, "c", "config", "Provides the setup information needed by and is required to start the program.")
	flaggy.Bool(
		&generateConfig,
		"gc",
		"genconfig",
		"This will output a skeleton configuration and note files.\n\t\t\tThis will cause all other options to be ignored.",
	)
	flaggy.String(
		&styhClientId,
		"ci",
		"clientId",
		"The NATS Connect assigned client id. You can find it here: https://production-nc-dashboard.web.app/.",
	)
	flaggy.String(
		&password,
		"p",
		"password",
		"The password you selected when you signed up for NATS connect services. This is encrypted using SSL and only exist in Cognito.",
	)
	flaggy.String(
		&secretKey,
		"sk",
		"secretKey",
		"The NATS Connect assigned secret key. This is encrypted using SSL and a new can be generated at https://production-nc-dashboard."+
			"web.app/.",
	)
	flaggy.String(
		&tempDirectory, "tmp", "tempDir", "The temporary directory where the NATS Client can read and write temporary files.",
	)
	flaggy.Bool(&testingOn, "t", "testingOn", "This puts the program into testing mode.")
	flaggy.String(
		&username,
		"u",
		"username",
		"The username you selected when you signed up for NATS connect services. This is encrypted using SSL and only exist in Cognito.",
	)

	// Set the version and parse all inputs into variables.
	flaggy.SetVersion(version)
	flaggy.Parse()
}

//goland:noinspection ALL
func main() {

	fmt.Println()
	log.Printf("Running %v.\n", programName)

	if generateConfig {
		config.GenerateConfigFileSkeleton(programName, "config/")
		os.Exit(0)
	}

	// This is to prevent the serverName from being empty.
	if programName == ctv.VAL_EMPTY {
		pi.PrintError(pi.ErrProgramNameMissing, fmt.Sprintf("%v %v", ctv.TXT_PROGRAM_NAME, programName))
		os.Exit(1)
	}

	if testingOn == false {
		// This is to prevent the version from being empty or not being set during the build process.
		if version == ctv.VAL_EMPTY || version == "9999.9999.9999" {
			pi.PrintError(pi.ErrVersionInvalid, fmt.Sprintf("%v %v", ctv.TXT_SERVER_VERSION, version))
			flaggy.ShowHelpAndExit("")
		}
		if username == ctv.VAL_EMPTY || password == ctv.VAL_EMPTY || styhClientId == ctv.VAL_EMPTY || secretKey == ctv.VAL_EMPTY || tempDirectory == ctv.VAL_EMPTY {
			// Has the config file location and name been provided, if not, return help.
			if configFileFQN == "" || configFileFQN == "-t" {
				flaggy.ShowHelpAndExit("")
			}
		}
	}

	run(styhClientId, environment, password, secretKey, tempDirectory, username, configFileFQN)

	os.Exit(0)
}

func run(styhClientId, environment, password, secretKey, tempDirectory, username, configFileFQN string) {

	var (
		accountId   string
		clientPtr   src.NCClient
		errorInfo   pi.ErrorInfo
		reply       []byte
		replyData   interface{}
		requestData interface{}
		systemId    string
		teamId      string
		tokenId     string
	)

	// The following is all the code the developer needs to use NATS Connect.

	// Connect to the NATS Connect service.
	if clientPtr, errorInfo = src.NewNCClient(
		styhClientId,
		environment,
		password,
		secretKey,
		tempDirectory,
		username,
		configFileFQN,
	); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		flaggy.ShowHelpAndExit("")
	}

	// Sample call to Synadia Cloud List Teams
	requestData = ncs.ListTeamsRequest{
		SaaSKey: SYNADIA_CLOUD_TOKEN,
		BaseURL: SYNADIA_CLOUD_BASE_URL,
	}
	if replyData, errorInfo = clientPtr.SynaidaListTeams(requestData); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		log.Fatalln()
	} else {
		teamId = replyData.(ncs.ListTeamsReply).Response.Items[0].Id
		fmt.Println("==============================")
		fmt.Println("Synadia List Teams")
		reply, _ = json.Marshal(replyData)
		s, _ := prettyjson.Format(reply)
		fmt.Println(string(s))
	}

	// Sample call to Synadia Cloud Get Team
	requestData = ncs.GetTeamRequest{
		SaaSKey: SYNADIA_CLOUD_TOKEN,
		BaseURL: SYNADIA_CLOUD_BASE_URL,
		TeamId:  teamId,
	}
	if replyData, errorInfo = clientPtr.SynaidaGetTeam(requestData); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		log.Fatalln()
	} else {
		fmt.Println("==============================")
		fmt.Println("Synadia Get Team")
		reply, _ = json.Marshal(replyData)
		s, _ := prettyjson.Format(reply)
		fmt.Println(string(s))
	}

	// Sample call to Synadia Cloud Get Team Limits
	requestData = ncs.GetTeamLimitsRequest{
		SaaSKey: SYNADIA_CLOUD_TOKEN,
		BaseURL: SYNADIA_CLOUD_BASE_URL,
		TeamId:  teamId,
	}
	if replyData, errorInfo = clientPtr.SynaidaGetTeamLimits(requestData); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		log.Fatalln()
	} else {
		fmt.Println("==============================")
		fmt.Println("Synadia Get Team's Limits")
		reply, _ = json.Marshal(replyData)
		s, _ := prettyjson.Format(reply)
		fmt.Println(string(s))
	}

	// Sample call to Synadia Cloud List Information App Users Team
	requestData = ncs.ListInfoAppUserTeamRequest{
		SaaSKey: SYNADIA_CLOUD_TOKEN,
		BaseURL: SYNADIA_CLOUD_BASE_URL,
		TeamId:  teamId,
	}
	if replyData, errorInfo = clientPtr.SynaidaListInfoAppUsersTeam(requestData); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		log.Fatalln()
	} else {
		fmt.Println("==============================")
		fmt.Println("Synadia List Information App Users Team")
		reply, _ = json.Marshal(replyData)
		s, _ := prettyjson.Format(reply)
		fmt.Println(string(s))
	}

	// Sample call to Synadia Cloud List Personal Access Tokens
	requestData = ncs.ListPersonalAccessTokensRequest{
		SaaSKey: SYNADIA_CLOUD_TOKEN,
		BaseURL: SYNADIA_CLOUD_BASE_URL,
		TeamId:  teamId,
	}
	if replyData, errorInfo = clientPtr.SynaidaListPersonalAccessTokens(requestData); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		log.Fatalln()
	} else {
		tokenId = replyData.(ncs.ListPersonalAccessTokensReply).Response.Items[0].Id
		fmt.Println("==============================")
		fmt.Println("Synadia List Personal Access Tokens")
		reply, _ = json.Marshal(replyData)
		s, _ := prettyjson.Format(reply)
		fmt.Println(string(s))
	}

	// Sample call to Synadia Cloud List Team Server Accounts
	requestData = ncs.ListTeamServerAccountsRequest{
		SaaSKey: SYNADIA_CLOUD_TOKEN,
		BaseURL: SYNADIA_CLOUD_BASE_URL,
		TeamId:  teamId,
	}
	if replyData, errorInfo = clientPtr.SynaidaListTeamServerAccounts(requestData); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		log.Fatalln()
	} else {
		fmt.Println("==============================")
		fmt.Println("Synadia List Team Server Accounts")
		reply, _ = json.Marshal(replyData)
		s, _ := prettyjson.Format(reply)
		fmt.Println(string(s))
	}

	// Sample call to Synadia Cloud List Systems
	requestData = ncs.ListSystemsRequest{
		SaaSKey: SYNADIA_CLOUD_TOKEN,
		BaseURL: SYNADIA_CLOUD_BASE_URL,
		TeamId:  teamId,
	}
	if replyData, errorInfo = clientPtr.SynaidaListSystems(requestData); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		log.Fatalln()
	} else {
		systemId = replyData.(ncs.ListSystemsReply).Response.Items[0].Id
		fmt.Println("==============================")
		fmt.Println("Synadia List Systems")
		reply, _ = json.Marshal(replyData)
		s, _ := prettyjson.Format(reply)
		fmt.Println(string(s))
	}

	// Sample call to Synadia Cloud Get System
	requestData = ncs.GetSystemRequest{
		SaaSKey:  SYNADIA_CLOUD_TOKEN,
		BaseURL:  SYNADIA_CLOUD_BASE_URL,
		SystemId: systemId,
	}
	if replyData, errorInfo = clientPtr.SynaidaGetSystem(requestData); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		log.Fatalln()
	} else {
		fmt.Println("==============================")
		fmt.Println("Synadia Get System")
		reply, _ = json.Marshal(replyData)
		s, _ := prettyjson.Format(reply)
		fmt.Println(string(s))
	}

	// Sample call to Synadia Cloud Get System Limits
	requestData = ncs.GetSystemLimitsRequest{
		SaaSKey:  SYNADIA_CLOUD_TOKEN,
		BaseURL:  SYNADIA_CLOUD_BASE_URL,
		SystemId: systemId,
	}
	if replyData, errorInfo = clientPtr.SynaidaGetSystemLimits(requestData); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		log.Fatalln()
	} else {
		fmt.Println("==============================")
		fmt.Println("Synadia Get System Limits")
		reply, _ = json.Marshal(replyData)
		s, _ := prettyjson.Format(reply)
		fmt.Println(string(s))
	}

	// Sample call to Synadia Cloud List Team Server Accounts
	requestData = ncs.ListSystemAccountInfoRequest{
		SaaSKey:  SYNADIA_CLOUD_TOKEN,
		BaseURL:  SYNADIA_CLOUD_BASE_URL,
		SystemId: systemId,
	}
	if replyData, errorInfo = clientPtr.SynaidaListSystemAccountInfo(requestData); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		log.Fatalln()
	} else {
		fmt.Println("==============================")
		fmt.Println("Synadia List System Account Info")
		reply, _ = json.Marshal(replyData)
		s, _ := prettyjson.Format(reply)
		fmt.Println(string(s))
	}

	// Sample call to Synadia Cloud List Accounts
	requestData = ncs.ListAccountsRequest{
		SaaSKey:  SYNADIA_CLOUD_TOKEN,
		BaseURL:  SYNADIA_CLOUD_BASE_URL,
		SystemId: systemId,
	}
	if replyData, errorInfo = clientPtr.SynaidaListAccounts(requestData.(ncs.ListAccountsRequest)); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		log.Fatalln()
	} else {
		accountId = replyData.(ncs.ListAccountsReply).Response.Items[0].Id
		fmt.Println("==============================")
		fmt.Println("Synadia List Accounts")
		reply, _ = json.Marshal(replyData)
		s, _ := prettyjson.Format(reply)
		fmt.Println(string(s))
	}

	// Sample call to Synadia Cloud List System Server Info
	requestData = ncs.ListSystemServerInfoRequest{
		SaaSKey:  SYNADIA_CLOUD_TOKEN,
		BaseURL:  SYNADIA_CLOUD_BASE_URL,
		SystemId: systemId,
	}
	if replyData, errorInfo = clientPtr.SynaidaListSystemServerInfo(requestData); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		log.Fatalln()
	} else {
		fmt.Println("==============================")
		fmt.Println("Synadia List System Server Info")
		reply, _ = json.Marshal(replyData)
		s, _ := prettyjson.Format(reply)
		fmt.Println(string(s))
	}

	// Sample call to Synadia Cloud Get Version
	requestData = ncs.GetVersionRequest{
		SaaSKey: SYNADIA_CLOUD_TOKEN,
		BaseURL: SYNADIA_CLOUD_BASE_URL,
	}
	if replyData, errorInfo = clientPtr.SynaidaGetVersion(requestData); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		log.Fatalln()
	} else {
		fmt.Println("==============================")
		fmt.Println("Synadia Get Version")
		reply, _ = json.Marshal(replyData)
		s, _ := prettyjson.Format(reply)
		fmt.Println(string(s))
	}

	// Sample call to Synadia Cloud List NATS Users
	requestData = ncs.ListNATSUsersRequest{
		SaaSKey:   SYNADIA_CLOUD_TOKEN,
		BaseURL:   SYNADIA_CLOUD_BASE_URL,
		AccountId: accountId,
	}
	if replyData, errorInfo = clientPtr.SynaidaListNATSUsers(requestData); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		log.Fatalln()
	} else {
		fmt.Println("==============================")
		fmt.Println("Synadia List NATS Users")
		reply, _ = json.Marshal(replyData)
		s, _ := prettyjson.Format(reply)
		fmt.Println(string(s))
	}

	// Sample call to Synadia Cloud Get Personal Access Token
	requestData = ncs.GetPersonalAccessTokenRequest{
		SaaSKey: SYNADIA_CLOUD_TOKEN,
		BaseURL: SYNADIA_CLOUD_BASE_URL,
		TokenId: tokenId,
	}
	if replyData, errorInfo = clientPtr.SynaidaGetPersonalAccessToken(requestData); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		log.Fatalln()
	} else {
		fmt.Println("==============================")
		fmt.Println("Synadia Get Personal Access Token")
		reply, _ = json.Marshal(replyData)
		s, _ := prettyjson.Format(reply)
		fmt.Println(string(s))
	}

}
