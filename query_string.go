//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package bleve

import (
	"github.com/blevesearch/bleve/search"
)

type QueryStringQuery struct {
	Query    string  `json:"query"`
	BoostVal float64 `json:"boost,omitempty"`
}

func NewQueryStringQuery(query string) *QueryStringQuery {
	return &QueryStringQuery{
		Query:    query,
		BoostVal: 1.0,
	}
}

func (q *QueryStringQuery) Boost() float64 {
	return q.BoostVal
}

func (q *QueryStringQuery) SetBoost(b float64) *QueryStringQuery {
	q.BoostVal = b
	return q
}

func (q *QueryStringQuery) Searcher(i *indexImpl, explain bool) (search.Searcher, error) {
	newQuery, err := ParseQuerySyntax(q.Query, i.m)
	if err != nil {
		return nil, err
	}
	return newQuery.Searcher(i, explain)
}

func (q *QueryStringQuery) Validate() error {
	return nil
}