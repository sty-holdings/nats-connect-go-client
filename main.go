// Package main.go

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/integrii/flaggy"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	ctv "github.com/sty-holdings/constant-type-vars-go/v2024"
	"github.com/sty-holdings/nats-connect-go-client/src"
	config "github.com/sty-holdings/sty-shared/v2024/configuration"
	pi "github.com/sty-holdings/sty-shared/v2024/programInfo"
)

const (
// Add Constants to the constants.go file
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
		clientPtr src.NCClient
		errorInfo pi.ErrorInfo
		//reply     []byte
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

	fmt.Println(clientPtr)

	//// Sample call to Synadia Cloud List Teams
	//if reply, errorInfo = clientPtr.SynaidaListTeams(data); errorInfo.Error != nil {
	//	pi.PrintErrorInfo(errorInfo)
	//} else {
	//	fmt.Println("==============================")
	//	fmt.Println("List available payment methods")
	//	s, _ := prettyjson.Format(reply)
	//	fmt.Println(string(s))
	//}
}
