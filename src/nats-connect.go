// Package src
/*
====> This is a sample usage of NATS Connect. The CLI part is to allow an easy place to start.
====> The run function is the code to drop into you program.

Copyright 6/5/24 STY Holdings Inc

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the “Software”), to deal in
the Software without restriction, including without limitation the rights to use,
copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the
Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.
*/
package src

import (
	"fmt"
	"strconv"

	awsSSM "github.com/aws/aws-sdk-go-v2/service/ssm"

	ctv "github.com/sty-holdings/constant-type-vars-go/v2024"
	awss "github.com/sty-holdings/sty-shared/v2024/awsServices"
	cfgs "github.com/sty-holdings/sty-shared/v2024/configuration"
	hv "github.com/sty-holdings/sty-shared/v2024/helpersValidators"
	jwts "github.com/sty-holdings/sty-shared/v2024/jwtServices"
	ns "github.com/sty-holdings/sty-shared/v2024/natsSerices"
	pi "github.com/sty-holdings/sty-shared/v2024/programInfo"
)

//goland:noinspection ALL
const (
	PROGRAM_NAME            = "NATS-Connect-go-client"
	NC_SSM_PARAMETER_PREFIX = "nc"
)

type styhCustomerConfig struct {
	clientId  string
	secretKey string
	tokens    awss.CognitoTokens
	username  string
}

// NewNCClient - creates an instance to connect to the NATS Connect server
//
//	Customer Messages: None
//	Errors: ErrRequiredArgumentMissing, returned from validateConfiguration, LoadAWSCustomerSettings, Login, processAWSClientParameters, BuildTemporaryFiles,
//	BuildTLSTemporaryFiles, BuildInstanceName, GetConnection
//	Verifications: styhClientId, environment, password, secretKey, tempDirectory, username, configFileFQN
func NewNCClient(styhClientId, environment, password, secretKey, tempDirectory, username, configFileFQN string) (
	NCClientPtr NCClient,
	errorInfo pi.ErrorInfo,
) {

	var (
		tEnvironment   string
		tPassword      string
		tSecretKey     string
		tSTYHClientId  string
		tTempDirectory string
		tUsername      string
	)

	var (
		tConfigMap = make(map[string]interface{})
	)

	// Load arguments
	if configFileFQN == ctv.VAL_EMPTY {
		tSTYHClientId = styhClientId
		tPassword = password
		tSecretKey = secretKey
		tTempDirectory = tempDirectory
		tUsername = username
		// environment is validated in awss.NewAWSConfig
		tEnvironment = environment
	} else {
		if tConfigMap, errorInfo = cfgs.GetConfigFile(configFileFQN); errorInfo.Error != nil {
			return
		}
		tSTYHClientId = tConfigMap[ctv.FN_STYH_CLIENT_ID].(string)
		tEnvironment = tConfigMap[ctv.FN_ENVIRONMENT].(string)
		tPassword = tConfigMap[ctv.FN_PASSWORD].(string)
		tConfigMap[ctv.FN_PASSWORD] = ctv.TXT_PROTECTED // Clear the password from memory.
		tSecretKey = tConfigMap[ctv.FN_SECRET_KEY].(string)
		tTempDirectory = tConfigMap[ctv.FN_TEMP_DIRECTORY].(string)
		tUsername = tConfigMap[ctv.FN_USERNAME].(string)
	}

	if errorInfo = validateConfiguration(tSTYHClientId, tEnvironment, tSecretKey, tTempDirectory, tUsername, &tPassword); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		return
	}

	if NCClientPtr.awsSettings, errorInfo = awss.LoadAWSCustomerSettings(tEnvironment); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		return
	}
	NCClientPtr.environment = tEnvironment
	NCClientPtr.tempDirectory = tTempDirectory

	// This returns information about the STYH Customer
	if NCClientPtr.styhCustomerConfig.tokens.Access,
		NCClientPtr.styhCustomerConfig.tokens.ID,
		NCClientPtr.styhCustomerConfig.tokens.Refresh, errorInfo = awss.Login(
		ctv.AUTH_USER_SRP, tUsername, &tPassword,
		NCClientPtr.awsSettings.STYHCognitoIdentityInfo, NCClientPtr.awsSettings.BaseConfig,
	); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		return
	}

	NCClientPtr.styhCustomerConfig.clientId = tSTYHClientId
	NCClientPtr.styhCustomerConfig.username = tUsername
	NCClientPtr.styhCustomerConfig.secretKey = tSecretKey
	tPassword = ctv.TXT_PROTECTED  // Clear the password from memory.
	secretKey = ctv.TXT_PROTECTED  // Clear the secret key from memory.
	tSecretKey = ctv.TXT_PROTECTED // Clear the secret key from memory.

	// Gets needed information to make connection
	if errorInfo = processAWSClientParameters(
		NCClientPtr.awsSettings,
		NCClientPtr.styhCustomerConfig.tokens.ID,
		tEnvironment,
		&NCClientPtr.natsConfig,
	); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		return
	}

	// Creates needed file for NATS
	if errorInfo = ns.BuildTemporaryFiles(NCClientPtr.tempDirectory, NCClientPtr.natsConfig); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		return
	}
	NCClientPtr.natsConfig.NATSCredentialsFilename = fmt.Sprintf("%v/%v", tTempDirectory, ns.CREDENTIAL_FILENAME)

	// Creates needed file for NATS
	if errorInfo = jwts.BuildTLSTemporaryFiles(NCClientPtr.tempDirectory, NCClientPtr.natsConfig.NATSTLSInfo); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		return
	}
	NCClientPtr.natsConfig.NATSTLSInfo.TLSCABundleFQN = fmt.Sprintf("%v/%v", tTempDirectory, jwts.TLS_CA_BUNDLE_FILENAME)
	NCClientPtr.natsConfig.NATSTLSInfo.TLSCertFQN = fmt.Sprintf("%v/%v", tTempDirectory, jwts.TLS_CERT_FILENAME)
	NCClientPtr.natsConfig.NATSTLSInfo.TLSPrivateKeyFQN = fmt.Sprintf("%v/%v", tTempDirectory, jwts.TLS_PRIVATE_KEY_FILENAME)

	// Builds name for tracking
	if NCClientPtr.natsService.InstanceName, errorInfo = ns.BuildInstanceName(ns.METHOD_DASHES, NCClientPtr.styhCustomerConfig.clientId); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		return
	}
	// Makes connection to the STYH NATS Server
	if NCClientPtr.natsService.ConnPtr, errorInfo = ns.GetConnection(NCClientPtr.natsService.InstanceName, NCClientPtr.natsConfig); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		return
	}

	return
}

// processAWSClientParameters - handles getting and storing the shared AWS SSM Parameters.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func processAWSClientParameters(
	awsSettings awss.AWSSettings,
	idToken string,
	environment string,
	natsConfigPtr *ns.NATSConfiguration,
) (errorInfo pi.ErrorInfo) {

	var (
		tParameterName    string
		tParametersOutput awsSSM.GetParametersOutput
		tParameterValue   string
	)

	if tParametersOutput, errorInfo = awss.GetParameters(
		awsSettings.STYHCognitoIdentityInfo,
		awsSettings.BaseConfig,
		idToken,
		ctv.GetParameterName(NC_SSM_PARAMETER_PREFIX, environment, ctv.PARAMETER_NATS_TOKEN),
		ctv.GetParameterName(NC_SSM_PARAMETER_PREFIX, environment, ctv.PARAMETER_NATS_PORT),
		ctv.GetParameterName(NC_SSM_PARAMETER_PREFIX, environment, ctv.PARAMETER_NATS_URL),
		ctv.GetParameterName(NC_SSM_PARAMETER_PREFIX, environment, ctv.PARAMETER_TLS_CERT),
		ctv.GetParameterName(NC_SSM_PARAMETER_PREFIX, environment, ctv.PARAMETER_TLS_PRIVATE_KEY),
		ctv.GetParameterName(NC_SSM_PARAMETER_PREFIX, environment, ctv.PARAMETER_TLS_CA_BUNDLE),
	); errorInfo.Error != nil {
		return
	}

	for _, parameter := range tParametersOutput.Parameters {
		tParameterName = *parameter.Name
		tParameterValue = *parameter.Value
		switch tParameterName {
		case ctv.GetParameterName(NC_SSM_PARAMETER_PREFIX, environment, ctv.PARAMETER_NATS_TOKEN):
			natsConfigPtr.NATSToken = tParameterValue
		case ctv.GetParameterName(NC_SSM_PARAMETER_PREFIX, environment, ctv.PARAMETER_NATS_PORT):
			natsConfigPtr.NATSPort, _ = strconv.Atoi(tParameterValue)
		case ctv.GetParameterName(NC_SSM_PARAMETER_PREFIX, environment, ctv.PARAMETER_NATS_URL):
			natsConfigPtr.NATSURL = tParameterValue
		case ctv.GetParameterName(NC_SSM_PARAMETER_PREFIX, environment, ctv.PARAMETER_TLS_CERT):
			natsConfigPtr.NATSTLSInfo.TLSCert = tParameterValue
		case ctv.GetParameterName(NC_SSM_PARAMETER_PREFIX, environment, ctv.PARAMETER_TLS_PRIVATE_KEY):
			natsConfigPtr.NATSTLSInfo.TLSPrivateKey = tParameterValue
		case ctv.GetParameterName(NC_SSM_PARAMETER_PREFIX, environment, ctv.PARAMETER_TLS_CA_BUNDLE):
			natsConfigPtr.NATSTLSInfo.TLSCABundle = tParameterValue
		default:
			// Optional: Handle unknown parameter names (log a warning?)
		}
	}

	return
}

// validateConfiguration - Checks inputs are provided.
//
//	Customer Messages: None
//	Errors: ErrRequiredArgumentMissing, ErrEnvironmentInvalid
//	Verifications: None
func validateConfiguration(
	styhClientId, environment, secretKey, tempDirectory, username string,
	passwordPtr *string,
) (
	errorInfo pi.ErrorInfo,
) {

	if styhClientId == ctv.VAL_EMPTY {
		errorInfo = pi.NewErrorInfo(pi.ErrRequiredArgumentMissing, fmt.Sprintf("%v%v", ctv.TXT_MISSING_PARAMETER, ctv.FN_STYH_CLIENT_ID))
		return
	}
	if hv.IsEnvironmentValid(environment) == false {
		errorInfo = pi.NewErrorInfo(pi.ErrEnvironmentInvalid, fmt.Sprintf("%v%v", ctv.TXT_ENVIRONMENT, ctv.FN_ENVIRONMENT))
		return
	}
	if passwordPtr == nil {
		errorInfo = pi.NewErrorInfo(pi.ErrRequiredArgumentMissing, fmt.Sprintf("%v%v", ctv.TXT_MISSING_PARAMETER, ctv.FN_PASSWORD))
		return
	}
	if secretKey == ctv.VAL_EMPTY {
		errorInfo = pi.NewErrorInfo(pi.ErrRequiredArgumentMissing, fmt.Sprintf("%v%v", ctv.TXT_MISSING_PARAMETER, ctv.FN_SECRET_KEY))
		return
	}
	if tempDirectory == ctv.VAL_EMPTY {
		errorInfo = pi.NewErrorInfo(pi.ErrRequiredArgumentMissing, fmt.Sprintf("%v%v", ctv.TXT_MISSING_PARAMETER, ctv.FN_TEMP_DIRECTORY))
		return
	}
	if username == ctv.VAL_EMPTY {
		errorInfo = pi.NewErrorInfo(pi.ErrRequiredArgumentMissing, fmt.Sprintf("%v%v", ctv.TXT_MISSING_PARAMETER, ctv.FN_USERNAME))
		return
	}

	return
}
