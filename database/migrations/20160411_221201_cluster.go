package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Cluster_20160411_221201 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Cluster_20160411_221201{}
	m.Created = "20160411_221201"
	migration.Register("Cluster_20160411_221201", m)
}

// Run the migrations
func (m *Cluster_20160411_221201) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE cluster(`id` int(11) NOT NULL AUTO_INCREMENT,`region` varchar(128) NOT NULL,`name` varchar(128) NOT NULL,`instance_type` varchar(128) NOT NULL,`size` int(11) DEFAULT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Cluster_20160411_221201) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `cluster`")
}
