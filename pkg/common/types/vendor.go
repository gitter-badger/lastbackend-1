//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2017] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package types

import (
	"encoding/json"
	"golang.org/x/oauth2"
)

type Vendor struct {
	ServiceID string        `json:"service_id"`
	Username  string        `json:"username"`
	Vendor    string        `json:"vendor"`
	Host      string        `json:"host"`
	Token     *oauth2.Token `json:"token"`
}

func (v *Vendor) ToJson() ([]byte, error) {
	buf, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

type Vendors map[string]*Vendor

func (v *Vendors) ToJson() ([]byte, error) {
	buf, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
