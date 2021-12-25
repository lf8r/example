// Copyright (C) Subhajit DasGupta 2021

package persondao

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/lf8r/dbgen/pkg/dbs"
	"github.com/lf8r/dbgen/pkg/ddlgen"
	"github.com/lf8r/dbgen/pkg/reflects"
	"github.com/lf8r/example-data/pkg/data"
	"github.com/stretchr/testify/require"
)

// Warning - This is generated code. It is overwritten on each build.

// LookupRequiredEnvVars returns a map with the values of the named environment
// variables. Returns an error if any of these is not set.
func LookupRequiredEnvVars(names ...string) (map[string]string, error) {
	ret := make(map[string]string)

	for i := range names {
		name := names[i]

		value, ok := os.LookupEnv(name)
		if !ok {
			return nil, fmt.Errorf("environment var not set %s", name)
		}

		ret[name] = value
	}

	return ret, nil
}

// OpenDB creates an sql.DB by reading the following environment variables.
// Returns an errorIf any of these variables is not set.
//  POSTGRES_HOST
//  POSTGRES_PORT
//  POSTGRES_USER
//  POSTGRES_PASSWORD
//  POSTGRES_DB
func OpenDB() (*sql.DB, error) {
	vals, err := LookupRequiredEnvVars("POSTGRES_HOST", "POSTGRES_PORT", "POSTGRES_USER", "POSTGRES_PASSWORD", "POSTGRES_DB")
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	host := vals["POSTGRES_HOST"]

	portStr := vals["POSTGRES_PORT"]

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("open db, could to parse port: %w", err)
	}

	user := vals["POSTGRES_USER"]
	password := vals["POSTGRES_PASSWORD"]
	dbName := vals["POSTGRES_DB"]

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("open db %s: %w", psqlInfo, err)
	}

	return db, nil
}

func setEnv(name, value string) bool {
	if err := os.Setenv(name, value); err != nil {
		// nolint
		fmt.Printf("Could not set env var %s", name)

		return false
	}

	return true
}

func TestMain(m *testing.M) {
	// Setup environment variables dictating a captive Postgres instance for tests.
	if !setEnv(dbs.EnvVarHost, "localhost") {
		return
	}

	if !setEnv(dbs.EnvVarPort, "7000") {
		return
	}

	if !setEnv(dbs.EnvVarUser, "postgres") {
		return
	}

	if !setEnv(dbs.EnvVarPassword, "docker") {
		return
	}

	if !setEnv(dbs.EnvVarDB, "postgres") {
		return
	}

	if !setEnv(dbs.EnvVarSystemSharedBuffers, "3") {
		return
	}

	if !setEnv(dbs.EnvVarMaxConnections, "200") {
		return
	}

	if !setEnv(dbs.EnvVarDataDir, "/tmp/db7000") {
		return
	}

	m.Run()
	os.Exit(0)
}

// newPerson creates a new Person for tests.
func newPerson() *data.Person {
	p := data.Person{}
	if err := reflects.Set(&p); err != nil {
		return nil
	}

	return &p
}

// TestNewPerson generates a new randomly initialized Person.
func TestNewPerson(t *testing.T) {
	assert := require.New(t)
	p := newPerson()
	assert.NotNil(p)

	_, err := jsoniter.MarshalIndent(p, "", "    ")
	assert.Nil(err)
}

// TestGeneratedPersonClientImpl tests the Person Client.
func TestGeneratedPersonClientImpl(t *testing.T) {
	assert := require.New(t)

	localDB, err := dbs.NewLocalDBFromEnv()
	assert.Nil(err)

	// Start the test database.
	assert.Nil(localDB.StartPG())

	// Stop it afterwards.
	defer localDB.StopPG()

	// Obtain an sql.Db instance from the database.
	db, err := localDB.OpenDB()
	assert.Nil(err)

	// Generate the DDL statement(s) to create the Person table.
	createStatements, dropStatements, err := ddlgen.GenerateJSONDDL(reflect.TypeOf(data.Person{}))
	assert.Nil(err)

	// Quick checks on the generated DDL statements.
	assert.Equal(2, len(createStatements))
	assert.Equal(3, len(dropStatements))

	// Remove any leftover data definitions we are about to create.
	for _, drop := range dropStatements {
		localDB.InvokeStatements([]string{drop})
	}

	// Cleanup once the test is done.
	defer func() {
		for _, drop := range dropStatements {
			localDB.InvokeStatements([]string{drop})
		}
	}()

	// Create the data definitions.
	assert.Nil(localDB.InvokeStatements(createStatements))

	// Get a transaction for the remainder of the test.
	tx, err := db.Begin()
	assert.Nil(err)

	// Roll back the transaction afterwards.
	defer tx.Rollback()

	// Create an instance of the Person client.
	c := PersonClientDBImpl{
		Tx: tx,
	}

	// Create a new Person instance.
	value := newPerson()
	PersonName := value.Name

	// Use the client to insert it to the Person table.
	createdValue, err := c.Create(value)
	assert.Nil(err)
	assert.True(createdValue.ID != "")
	assert.Equal(PersonName, createdValue.Name)

	// List Persons. Expect one record (the one created above).
	values, err := c.List(nil)
	assert.Nil(err)
	assert.Equal(1, len(values))

	listedValue := values[0]

	// Check its fields.
	assert.True(listedValue.ID != "")
	assert.Equal(PersonName, listedValue.Name)

	// Get the Person by ID.
	valueByID, err := c.GetByID(listedValue.ID)
	assert.Nil(err)
	assert.Equal(listedValue, *valueByID)

	// Change some attributes of the Person.
	changedName := uuid.NewString()
	valueByID.Name = changedName
	oldModifiedTime := valueByID.Modified

	// Pause for a bit to allow the modified time (below) to differ from "now"
	// on fast computers.
	time.Sleep(250 * time.Millisecond)

	// Update the Person.
	updatedValue, err := c.Update(valueByID)
	assert.Nil(err)
	assert.NotNil(updatedValue)
	assert.Equal(valueByID.ID, updatedValue.ID)
	assert.Equal(changedName, updatedValue.Name)

	// Get the updated Person by ID.
	updatedValueByID, err := c.GetByID(valueByID.ID)
	assert.Nil(err)
	assert.NotNil(updatedValueByID)
	assert.NotEqual(oldModifiedTime, updatedValueByID.Modified)
	assert.True(updatedValueByID.Modified.After(oldModifiedTime))

	// Negative test for GetByID for a non-existent Person.
	valueByID, err = c.GetByID("BadID")
	assert.Nil(err)
	assert.Nil(valueByID)

	// Delete a Person by ID using a bad ID.
	assert.Nil(c.DeleteByID("BadID"))

	// Delete the Person by its ID.
	assert.Nil(c.DeleteByID(listedValue.ID))
	valueByID, err = c.GetByID(listedValue.ID)
	assert.Nil(err)
	assert.Nil(valueByID)

	// Recreate the Person and test Delete Person.
	createdValue, err = c.Create(value)
	assert.Nil(err)

	// Get the Person.
	valueByID, err = c.GetByID(createdValue.ID)
	assert.Nil(err)
	assert.NotNil(valueByID)
	assert.Nil(c.Delete(valueByID))

	// Delete it.
	assert.Nil(c.Delete(valueByID))
	valueByID, err = c.GetByID(listedValue.ID)
	assert.Nil(err)
	assert.Nil(valueByID)
}

// TestBulkAddDeletePerson tests bulk add/delete in the Person Client.
func TestBulkAddDeletePerson(t *testing.T) {
	assert := require.New(t)

	localDB, err := dbs.NewLocalDBFromEnv()
	assert.Nil(err)

	// Start the test database.
	assert.Nil(localDB.StartPG())

	// Stop it afterwards.
	defer localDB.StopPG()

	// Obtain an sql.Db instance from the database.
	db, err := localDB.OpenDB()
	assert.Nil(err)

	// Generate the DDL statement(s) to create the Person table.
	createStatements, dropStatements, err := ddlgen.GenerateJSONDDL(reflect.TypeOf(data.Person{}))
	assert.Nil(err)

	// Quick checks on the generated DDL statements.
	assert.Equal(2, len(createStatements))
	assert.Equal(3, len(dropStatements))

	// Remove any leftover data definitions we are about to create.
	for _, drop := range dropStatements {
		localDB.InvokeStatements([]string{drop})
	}

	// Cleanup once the test is done.
	defer func() {
		for _, drop := range dropStatements {
			localDB.InvokeStatements([]string{drop})
		}
	}()

	// Create the data definitions.
	assert.Nil(localDB.InvokeStatements(createStatements))

	// Get a transaction for the remainder of the test.
	tx, err := db.Begin()
	assert.Nil(err)

	// Roll back the transaction afterwards.
	defer tx.Rollback()

	// Create an instance of the Person client.
	c := PersonClientDBImpl{
		Tx: tx,
	}

	// Add 100 Persons.
	ids := []string{}
	for i := 0; i < 100; i++ {
		p := newPerson()
		c.Create(p)

		assert.True(p.ID != "")
		ids = append(ids, p.ID)
	}

	// Delete each one by ID.
	for _, id := range ids {
		assert.Nil(c.DeleteByID(id))
	}
}

// TestListQueryPerson tests queries for the Person Client.
func TestListQueryPerson(t *testing.T) {
	assert := require.New(t)

	localDB, err := dbs.NewLocalDBFromEnv()
	assert.Nil(err)

	// Start the test database.
	assert.Nil(localDB.StartPG())

	// Stop it afterwards.
	defer localDB.StopPG()

	// Obtain an sql.Db instance from the database.
	db, err := localDB.OpenDB()
	assert.Nil(err)

	// Generate the DDL statement(s) to create the Person table.
	createStatements, dropStatements, err := ddlgen.GenerateJSONDDL(reflect.TypeOf(data.Person{}))
	assert.Nil(err)

	// Quick checks on the generated DDL statements.
	assert.Equal(2, len(createStatements))
	assert.Equal(3, len(dropStatements))

	// Remove any leftover data definitions we are about to create.
	for _, drop := range dropStatements {
		localDB.InvokeStatements([]string{drop})
	}

	// Cleanup once the test is done.
	defer func() {
		for _, drop := range dropStatements {
			localDB.InvokeStatements([]string{drop})
		}
	}()

	// Create the data definitions.
	assert.Nil(localDB.InvokeStatements(createStatements))

	// Get a transaction for the remainder of the test.
	tx, err := db.Begin()
	assert.Nil(err)

	// Roll back the transaction afterwards.
	defer tx.Rollback()

	// Create an instance of the Person client.
	c := PersonClientDBImpl{
		Tx: tx,
	}

	// Add 100 Persons.
	vals := make(map[string]*data.Person)

	for i := 0; i < 100; i++ {
		p := newPerson()
		c.Create(p)

		assert.True(p.ID != "")
		vals[p.ID] = p
	}

	// Query each Person by name.
	for _, v := range vals {
		m := make(map[string]interface{})
		m["Name"] = v.Name

		rows, err := c.List(m)
		assert.Nil(err)
		assert.Equal(1, len(rows))
		assert.Equal(v.Name, rows[0].Name)
	}
}
