package integration

import (
	"log"
	"os"
	"testing"

	"github.com/SyaibanAhmadRamadhan/go-clean-arch/infrastructures/config"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/tests/integration/utils"
	"github.com/ory/dockertest/v3"
)

func TestMain(t *testing.M) {
	pool := utils.InitDocker()
	var resources []*dockertest.Resource

	pgResource, dbPg, url := utils.PostgresStart(pool)
	resources = append(resources, pgResource)
	db = dbPg
	if db == nil {
		panic("db nil")
	}

	minioResourece, endpoint := utils.MinioStart(utils.InitDocker())
	resources = append(resources, minioResourece)
	minioConn, err := config.NewMinioConn(endpoint, "MYACCESSKEY", "MYSECRETKEY", false)
	if err != nil {
		panic(err)
	}
	minioClient = minioConn
	utils.StartMigration(url, db)
	// Run tests
	code := t.Run()
	// You can't defer this because os.Exit doesn't care for defer
	for _, resource := range resources {
		if err := pool.Purge(resource); err != nil {
			log.Fatal(err)
		}
	}
	os.Exit(code)
}
