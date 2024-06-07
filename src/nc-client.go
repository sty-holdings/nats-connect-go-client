// Package src
// /*
// ====> This is a sample usage of NATS Connect. The CLI part is to allow an easy place to start.
// ====> The run function is the code to drop into you program.
//
// Copyright 6/5/24 STY Holdings Inc
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
//
//	awsSSM "github.com/aws/aws-sdk-go-v2/service/ssm"
//	"github.com/nats-io/nats.go"
//
//	ctv "github.com/sty-holdings/constant-type-vars-go/v2024"
//	ncs "github.com/sty-holdings/nats-connect-shared"
//	awss "github.com/sty-holdings/sty-shared/v2024/awsServices"
//	cfgs "github.com/sty-holdings/sty-shared/v2024/configuration"
//	hv "github.com/sty-holdings/sty-shared/v2024/helpersValidators"
//	jwts "github.com/sty-holdings/sty-shared/v2024/jwtServices"
//	ns "github.com/sty-holdings/sty-shared/v2024/natsSerices"
//	pi "github.com/sty-holdings/sty-shared/v2024/programInfo"
//)
//
////goland:noinspection ALL
//const (
//	PROGRAM_NAME            = "NATS-Connect-go-client"
//	NC_SSM_PARAMETER_PREFIX = "NC"
//)
//
//type NCClient struct {
//	awsSettings        awss.AWSSettings
//	environment        string
//	natsService        ns.NATSService
//	natsConfig         ns.NATSConfiguration
//	secretKey          string
//	styhCustomerConfig styhCustomerConfig
//	tempDirectory      string
//}
//
//type NCPaymentInfo struct {
//	Amount                    float64  `json:"amount,omitempty"`
//	UseAutomaticPaymentMethod bool     `json:"use_automatic_payment_method,omitempty"`
//	CancellationReason        string   `json:"cancellation_reason,omitempty"`
//	CaptureFunds              string   `json:"capture_funds,omitempty"`
//	Currency                  string   `json:"currency,omitempty"`
//	CustomerId                string   `json:"customer_id,omitempty"`
//	Description               string   `json:"description,omitempty"`
//	ReturnRecordsLimit        int64    `json:"return_records_limit,omitempty"`
//	PaymentIntentId           string   `json:"id,omitempty"`
//	PaymentMethod             string   `json:"payment_method,omitempty"`
//	ReturnURL                 string   `json:"return_url,omitempty,omitempty"`
//	SenderEmailAddress        string   `json:"sender_email_address,omitempty"`
//	SenderName                string   `json:"sender_name,omitempty"`
//	StartingAfterRecord       string   `json:"starting_after_record,omitempty"`
//	ToEmailAddress            string   `json:"to_email_address,omitempty"`
//	ToEmailName               string   `json:"to_email_name,omitempty"`
//	Keys                      SaaSKeys `json:"keys,omitempty"`
//}
//
//type SaaSKeys struct {
//	StripeKey   string `json:"stripe_key"`
//	SendGridKey string `json:"sendgrid_key"`
//}
//
//type styhCustomerConfig struct {
//	clientId  string
//	secretKey string
//	tokens    awss.CognitoTokens
//	username  string
//}
//
//func NewNCClient(styhClientId, environment, password, secretKey, tempDirectory, username, configFileFQN string) (
//	NCClientPtr NCClient,
//	errorInfo pi.ErrorInfo,
//) {
//
//	var (
//		tEnvironment   string
//		tPassword      string
//		tSecretKey     string
//		tSTYHClientId  string
//		tTempDirectory string
//		tUsername      string
//	)
//
//	var (
//		tConfigMap = make(map[string]interface{})
//	)
//
//	if configFileFQN == ctv.VAL_EMPTY {
//		if styhClientId == ctv.VAL_EMPTY {
//			errorInfo = pi.NewErrorInfo(pi.ErrRequiredArgumentMissing, fmt.Sprintf("%v%v", ctv.TXT_MISSING_PARAMETER, ctv.FN_CLIENT_ID))
//			return
//		} else {
//			tSTYHClientId = styhClientId
//		}
//		if password == ctv.VAL_EMPTY {
//			errorInfo = pi.NewErrorInfo(pi.ErrRequiredArgumentMissing, fmt.Sprintf("%v%v", ctv.TXT_MISSING_PARAMETER, ctv.FN_PASSWORD))
//			return
//		} else {
//			tPassword = password
//		}
//		if secretKey == ctv.VAL_EMPTY {
//			errorInfo = pi.NewErrorInfo(pi.ErrRequiredArgumentMissing, fmt.Sprintf("%v%v", ctv.TXT_MISSING_PARAMETER, ctv.FN_SECRET_KEY))
//			return
//		} else {
//			tSecretKey = secretKey
//		}
//		if tempDirectory == ctv.VAL_EMPTY {
//			errorInfo = pi.NewErrorInfo(pi.ErrRequiredArgumentMissing, fmt.Sprintf("%v%v", ctv.TXT_MISSING_PARAMETER, ctv.FN_TEMP_DIRECTORY))
//			return
//		} else {
//			tTempDirectory = tempDirectory
//		}
//		if username == ctv.VAL_EMPTY {
//			errorInfo = pi.NewErrorInfo(pi.ErrRequiredArgumentMissing, fmt.Sprintf("%v%v", ctv.TXT_MISSING_PARAMETER, ctv.FN_USERNAME))
//			return
//		} else {
//			tUsername = username
//		}
//		// environment is validated in awss.NewAWSConfig
//		tEnvironment = environment
//	} else {
//		if tConfigMap, errorInfo = cfgs.GetConfigFile(configFileFQN); errorInfo.Error != nil {
//			return
//		}
//		tSTYHClientId = tConfigMap[ctv.FN_STYH_CLIENT_ID].(string)
//		tEnvironment = tConfigMap[ctv.FN_ENVIRONMENT].(string)
//		tPassword = tConfigMap[ctv.FN_PASSWORD].(string)
//		tConfigMap[ctv.FN_PASSWORD] = ctv.TXT_PROTECTED // Clear the password from memory.
//		tSecretKey = tConfigMap[ctv.FN_SECRET_KEY].(string)
//		tTempDirectory = tConfigMap[ctv.FN_TEMP_DIRECTORY].(string)
//		tUsername = tConfigMap[ctv.FN_USERNAME].(string)
//	}
//
//	if errorInfo = validateConfiguration(tSTYHClientId, tEnvironment, tSecretKey, tTempDirectory, tUsername, &tPassword); errorInfo.Error != nil {
//		pi.PrintErrorInfo(errorInfo)
//		return
//	}
//
//	if NCClientPtr.awsSettings, errorInfo = awss.LoadAWSCustomerSettings(tEnvironment); errorInfo.Error != nil {
//		pi.PrintErrorInfo(errorInfo)
//		return
//	}
//	NCClientPtr.environment = tEnvironment
//	NCClientPtr.tempDirectory = tTempDirectory
//
//	// This returns information about the STYH Customer
//	if NCClientPtr.styhCustomerConfig.tokens.Access,
//		NCClientPtr.styhCustomerConfig.tokens.ID,
//		NCClientPtr.styhCustomerConfig.tokens.Refresh, errorInfo = awss.Login(
//		ctv.AUTH_USER_SRP, tUsername, &tPassword,
//		NCClientPtr.awsSettings.STYHCognitoIdentityInfo, NCClientPtr.awsSettings.BaseConfig,
//	); errorInfo.Error != nil {
//		pi.PrintErrorInfo(errorInfo)
//		return
//	}
//
//	NCClientPtr.styhCustomerConfig.clientId = tSTYHClientId
//	NCClientPtr.styhCustomerConfig.username = tUsername
//	NCClientPtr.secretKey = tSecretKey
//	tPassword = ctv.TXT_PROTECTED  // Clear the password from memory.
//	secretKey = ctv.TXT_PROTECTED  // Clear the secret key from memory.
//	tSecretKey = ctv.TXT_PROTECTED // Clear the secret key from memory.
//
//	if errorInfo = processAWSClientParameters(
//		NCClientPtr.awsSettings,
//		NCClientPtr.styhCustomerConfig.tokens.ID,
//		tEnvironment,
//		&NCClientPtr.natsConfig,
//	); errorInfo.Error != nil {
//		pi.PrintErrorInfo(errorInfo)
//		return
//	}
//
//	if errorInfo = ns.BuildTemporaryFiles(NCClientPtr.tempDirectory, NCClientPtr.natsConfig); errorInfo.Error != nil {
//		pi.PrintErrorInfo(errorInfo)
//		return
//	}
//	NCClientPtr.natsConfig.NATSCredentialsFilename = fmt.Sprintf("%v/%v", tTempDirectory, ns.CREDENTIAL_FILENAME)
//
//	if errorInfo = jwts.BuildTLSTemporaryFiles(NCClientPtr.tempDirectory, NCClientPtr.natsConfig.NATSTLSInfo); errorInfo.Error != nil {
//		pi.PrintErrorInfo(errorInfo)
//		return
//	}
//	NCClientPtr.natsConfig.NATSTLSInfo.TLSCABundleFQN = fmt.Sprintf("%v/%v", tTempDirectory, jwts.TLS_CA_BUNDLE_FILENAME)
//	NCClientPtr.natsConfig.NATSTLSInfo.TLSCertFQN = fmt.Sprintf("%v/%v", tTempDirectory, jwts.TLS_CERT_FILENAME)
//	NCClientPtr.natsConfig.NATSTLSInfo.TLSPrivateKeyFQN = fmt.Sprintf("%v/%v", tTempDirectory, jwts.TLS_PRIVATE_KEY_FILENAME)
//
//	if NCClientPtr.natsService.InstanceName, errorInfo = ns.BuildInstanceName(ns.METHOD_DASHES, NCClientPtr.styhCustomerConfig.clientId); errorInfo.Error != nil {
//		pi.PrintErrorInfo(errorInfo)
//		return
//	}
//	if NCClientPtr.natsService.ConnPtr, errorInfo = ns.GetConnection(NCClientPtr.natsService.InstanceName, NCClientPtr.natsConfig); errorInfo.Error != nil {
//		pi.PrintErrorInfo(errorInfo)
//		return
//	}
//
//	return
//}
//
//func (clientPtr *NCClient) SynadiaListTeams(listTeamsRequest ncs.ListTeamsRequest) (
//	reply []byte,
//	errorInfo pi.ErrorInfo,
//) {
//	var (
//		tReply *nats.Msg
//	)
//
//	if listTeamsRequest.SaaSKey == ctv.VAL_EMPTY && listTeamsRequest.BaseURL == ctv.VAL_EMPTY {
//		errorInfo = pi.NewErrorInfo(pi.ErrRequiredArgumentMissing, fmt.Sprintf("%v and %v %v", ctv.TXT_PUBLIC_KEY, ctv.TXT_SECRET_KEY, ctv.TXT_ARE_MISSING))
//		return
//	}
//	//ctv.SUB_SYNADIA_LIST_TEAMS
//
//	tReply, errorInfo = processCancelPaymentIntent(
//		ai2cClientPtr.styhCustomerConfig.clientId, ai2cClientPtr.secretKey, ai2cClientPtr.styhCustomerConfig.username, &ai2cClientPtr.natsService,
//		ai2CPaymentInfo,
//
//	return
//}
