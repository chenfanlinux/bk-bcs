/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package store

import (
	"bk-bcs/bcs-common/common/zkclient"
	"time"
)

//dbZk is a struct of the zookeeper client
type dbZk struct {
	ZkHost []string
	ZkCli  *zkclient.ZkClient
}

//NewDbZk create a dbZk object
func NewDbZk(host []string) Dbdrvier {
	zk := dbZk{
		ZkHost: host[:],
		ZkCli:  zkclient.NewZkClient(host),
	}

	return &zk
}

func (zk *dbZk) Connect() error {
	return zk.ZkCli.Connect()
}

func (zk *dbZk) Close() {
	zk.ZkCli.Close()
}

func (zk *dbZk) Insert(path string, value string) error {
	var failed bool
	started := time.Now()

	err := zk.ZkCli.Update(path, value)
	if err != nil {
		failed = true
	}

	reportStorageOperatorMetrics(StoreOperatorCreate, started, failed)
	return err
}

func (zk *dbZk) Fetch(path string) ([]byte, error) {
	var failed bool
	started := time.Now()

	data, err := zk.ZkCli.Get(path)
	if err != nil {
		failed = true
	}

	reportStorageOperatorMetrics(StoreOperatorFetch, started, failed)
	return []byte(data), err
}

func (zk *dbZk) Update(path string, value string) error {
	var failed bool
	started := time.Now()

	err := zk.ZkCli.Update(path, value)
	if err != nil {
		failed = true
	}

	reportStorageOperatorMetrics(StoreOperatorUpdate, started, failed)
	return err

}

func (zk *dbZk) Delete(path string) error {
	var failed bool
	started := time.Now()

	err := zk.ZkCli.Del(path, -1)
	if err != nil {
		failed = true
	}

	reportStorageOperatorMetrics(StoreOperatorDelete, started, failed)
	return err
}

func (zk *dbZk) List(path string) ([]string, error) {
	b, _ := zk.ZkCli.Exist(path)
	if !b {
		return nil, nil
	}

	var failed bool
	started := time.Now()

	childs, err := zk.ZkCli.GetChildren(path)
	if err != nil {
		failed = true
	}

	reportStorageOperatorMetrics(StoreOperatorFetch, started, failed)
	return childs, err
}
