// Package src
// /*
// Copyright 1/2024 STY Holdings Inc
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the “Software”), to deal in
// the Software without restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the
// Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.
// */
package src

//
//import (
//	"fmt"
//	"strconv"
//
//	awsSSM "github.com/aws/aws-sdk-go-v2/service/ssm"
//
//	ctv "github.com/sty-holdings/constant-type-vars-go/v2024"
//	awss "github.com/sty-holdings/sty-shared/v2024/awsServices"
//	hv "github.com/sty-holdings/sty-shared/v2024/helpersValidators"
//	ns "github.com/sty-holdings/sty-shared/v2024/natsSerices"
//	pi "github.com/sty-holdings/sty-shared/v2024/programInfo"
//)
//
//// listTeam - returns Synadia Cloud team information
////
////	Customer Messages: None
////	Errors: None
////	Verifications: None
//func listTeam(
//	awsSettings awss.AWSSettings,
//	idToken string,
//	environment string,
//	natsConfigPtr *ns.NATSConfiguration,
//) (errorInfo pi.ErrorInfo) {
//
//	var (
//		tParameterName    string
//		tParametersOutput awsSSM.GetParametersOutput
//		tParameterValue   string
//	)
//
//	if tParametersOutput, errorInfo = awss.GetParameters(
//		awsSettings.STYHCognitoIdentityInfo,
//		awsSettings.BaseConfig,
//		idToken,
//		ctv.GetParameterName(NC_SSM_PARAMETER_PREFIX, environment, ctv.PARAMETER_NATS_TOKEN),
//		ctv.GetParameterName(NC_SSM_PARAMETER_PREFIX, environment, ctv.PARAMETER_NATS_PORT),
//		ctv.GetParameterName(NC_SSM_PARAMETER_PREFIX, environment, ctv.PARAMETER_NATS_URL),
//		ctv.GetParameterName(NC_SSM_PARAMETER_PREFIX, environment, ctv.PARAMETER_TLS_CERT),
//		ctv.GetParameterName(NC_SSM_PARAMETER_PREFIX, environment, ctv.PARAMETER_TLS_PRIVATE_KEY),
//		ctv.GetParameterName(NC_SSM_PARAMETER_PREFIX, environment, ctv.PARAMETER_TLS_CA_BUNDLE),
//	); errorInfo.Error != nil {
//		pi.PrintErrorInfo(errorInfo)
//		return
//	}
//
//	for _, parameter := range tParametersOutput.Parameters {
//		tParameterName = *parameter.Name
//		tParameterValue = *parameter.Value
//		switch tParameterName {
//		case ctv.GetParameterName(NC_SSM_PARAMETER_PREFIX, environment, ctv.PARAMETER_NATS_TOKEN):
//			natsConfigPtr.NATSToken = tParameterValue
//		case ctv.GetParameterName(NC_SSM_PARAMETER_PREFIX, environment, ctv.PARAMETER_NATS_PORT):
//			natsConfigPtr.NATSPort, _ = strconv.Atoi(tParameterValue)
//		case ctv.GetParameterName(NC_SSM_PARAMETER_PREFIX, environment, ctv.PARAMETER_NATS_URL):
//			natsConfigPtr.NATSURL = tParameterValue
//		case ctv.GetParameterName(NC_SSM_PARAMETER_PREFIX, environment, ctv.PARAMETER_TLS_CERT):
//			natsConfigPtr.NATSTLSInfo.TLSCert = tParameterValue
//		case ctv.GetParameterName(NC_SSM_PARAMETER_PREFIX, environment, ctv.PARAMETER_TLS_PRIVATE_KEY):
//			natsConfigPtr.NATSTLSInfo.TLSPrivateKey = tParameterValue
//		case ctv.GetParameterName(NC_SSM_PARAMETER_PREFIX, environment, ctv.PARAMETER_TLS_CA_BUNDLE):
//			natsConfigPtr.NATSTLSInfo.TLSCABundle = tParameterValue
//		default:
//			// Optional: Handle unknown parameter names (log a warning?)
//		}
//	}
//
//	return
//}
//
//// validateConfiguration - checks the values in the configuration file are valid. ValidateConfiguration doesn't
//// test if the configuration file exists, readable, or parsable.
////
////	Customer Messages: None
////	Errors: ErrEnvironmentInvalid, ErrRequiredArgumentMissing
////	Verifications: None
//func validateConfiguration(
//	styhClientId, environment, secretKey, tempDirectory, username string,
//	passwordPtr *string,
//) (
//	errorInfo pi.ErrorInfo,
//) {
//
//	if styhClientId == ctv.VAL_EMPTY {
//		errorInfo = pi.NewErrorInfo(pi.ErrRequiredArgumentMissing, fmt.Sprintf("%v%v", ctv.TXT_MISSING_PARAMETER, ctv.FN_STYH_CLIENT_ID))
//		return
//	}
//	if hv.IsEnvironmentValid(environment) == false {
//		errorInfo = pi.NewErrorInfo(pi.ErrEnvironmentInvalid, fmt.Sprintf("%v%v", ctv.TXT_EVIRONMENT, ctv.FN_ENVIRONMENT))
//		return
//	}
//	if passwordPtr == nil {
//		errorInfo = pi.NewErrorInfo(pi.ErrRequiredArgumentMissing, fmt.Sprintf("%v%v", ctv.TXT_MISSING_PARAMETER, ctv.FN_PASSWORD))
//		return
//	}
//	if secretKey == ctv.VAL_EMPTY {
//		errorInfo = pi.NewErrorInfo(pi.ErrRequiredArgumentMissing, fmt.Sprintf("%v%v", ctv.TXT_MISSING_PARAMETER, ctv.FN_SECRET_KEY))
//		return
//	}
//	if tempDirectory == ctv.VAL_EMPTY {
//		errorInfo = pi.NewErrorInfo(pi.ErrRequiredArgumentMissing, fmt.Sprintf("%v%v", ctv.TXT_MISSING_PARAMETER, ctv.FN_TEMP_DIRECTORY))
//		return
//	}
//	if username == ctv.VAL_EMPTY {
//		errorInfo = pi.NewErrorInfo(pi.ErrRequiredArgumentMissing, fmt.Sprintf("%v%v", ctv.TXT_MISSING_PARAMETER, ctv.FN_USERNAME))
//		return
//	}
//
//	return
//}
