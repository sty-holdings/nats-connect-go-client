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

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"

	"github.com/nats-io/nats.go"

	ctv "github.com/sty-holdings/constant-type-vars-go/v2024"
	ncs "github.com/sty-holdings/nats-connect-shared"
	jwts "github.com/sty-holdings/sty-shared/v2024/jwtServices"
	ns "github.com/sty-holdings/sty-shared/v2024/natsSerices"
	pi "github.com/sty-holdings/sty-shared/v2024/programInfo"
)

// ListTeams - returns Synadia Cloud teams information
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func listTeams(
	clientId, secretKey, username, instanceName string,
	request ncs.ListTeamsRequest,
	connPtr *nats.Conn,
) (reply *nats.Msg, errorInfo pi.ErrorInfo) {

	var (
		tEncryptedRequest  string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tJSONRequest       []byte
		tNATSHeader        = make(map[string][]string)
		tRequestMsg        nats.Msg
	)

	if tJSONRequest, errorInfo.Error = json.Marshal(request); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v - %v%v", ctv.TXT_FUNCTION_NAME, tFunctionName, ctv.TXT_SUBJECT, ctv.SUB_SYNADIA_LIST_TEAMS))
		return
	}
	if tEncryptedRequest, errorInfo = jwts.Encrypt(clientId, secretKey, string(tJSONRequest)); errorInfo.Error != nil {
		return
	}

	tNATSHeader[ctv.FN_STYH_CLIENT_ID] = []string{clientId}
	tNATSHeader[ctv.FN_USERNAME] = []string{username}
	tRequestMsg = nats.Msg{
		Subject: ctv.SUB_SYNADIA_LIST_TEAMS,
		Data:    []byte(tEncryptedRequest),
		Header:  tNATSHeader,
	}

	reply, errorInfo = ns.RequestWithHeader(connPtr, instanceName, &tRequestMsg, 2*time.Second)

	return
}
