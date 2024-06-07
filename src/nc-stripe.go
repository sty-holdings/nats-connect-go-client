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

// processAWSClientParameters - handles getting and storing the shared AWS SSM Parameters.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
//func processSynadia(
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

// validateConfiguration - checks the values in the configuration file are valid. ValidateConfiguration doesn't
// test if the configuration file exists, readable, or parsable.
//
//	Customer Messages: None
//	Errors: ErrEnvironmentInvalid, ErrRequiredArgumentMissing
//	Verifications: None
