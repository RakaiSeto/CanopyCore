package main

import (
	"CanopyCore/modules"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"github.com/robfig/cron"
	"github.com/streadway/amqp"
)

var db *sql.DB
var rc *redis.Client
var cx context.Context
var connRabbit *amqp.Connection
var chIncoming *amqp.Channel

var process = 1

const THEURL1 = "https://dapo.kemdikbud.go.id/rekap/progresSP?id_level_wilayah=3&kode_wilayah="
const THEURL2 = "&semester_id=20232&bentuk_pendidikan_id="

type Sekolah struct {
	Nama                      string      `json:"nama"`
	SekolahID                 string      `json:"sekolah_id"`
	Npsn                      interface{} `json:"npsn,omitempty"`
	JumlahKirim               interface{} `json:"jumlah_kirim"`
	Ptk                       interface{} `json:"ptk"`
	Pegawai                   interface{} `json:"pegawai"`
	Pd                        interface{} `json:"pd"`
	Rombel                    interface{} `json:"rombel"`
	JmlRk                     interface{} `json:"jml_rk"`
	JmlLab                    interface{} `json:"jml_lab"`
	JmlPerpus                 interface{} `json:"jml_perpus"`
	IndukKecamatan            string      `json:"induk_kecamatan"`
	KodeWilayahIndukKecamatan string      `json:"kode_wilayah_induk_kecamatan"`
	IndukKabupaten            string      `json:"induk_kabupaten"`
	KodeWilayahIndukKabupaten string      `json:"kode_wilayah_induk_kabupaten"`
	IndukProvinsi             string      `json:"induk_provinsi"`
	KodeWilayahIndukProvinsi  string      `json:"kode_wilayah_induk_provinsi"`
	BentukPendidikan          string      `json:"bentuk_pendidikan"`
	StatusSekolah             string      `json:"status_sekolah"`
	SinkronTerakhir           string      `json:"sinkron_terakhir"`
	SekolahIDEnkrip           string      `json:"sekolah_id_enkrip"`
}

func main() {
	// Load configuration file
	modules.InitiateGlobalVariables()
	runtime.GOMAXPROCS(4)

	start := time.Now()

	// Initiate Database
	var errDB error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		modules.MapConfig["pgsqlHost"], modules.MapConfig["pgsqlPort"], modules.MapConfig["pgsqlUser"],
		modules.MapConfig["pgsqlPassword"], modules.MapConfig["pgsqlName"])
	db, errDB = sql.Open("postgres", psqlInfo) // db udah di defined diatas, jadi harus pake = bukan :=

	if errDB != nil {
		modules.Logging(modules.Resource(), "STARTING UP", "START SERVICE", "SERVER IP", "Postgres unconnected", errDB)
		panic(errDB)
	} else {
		modules.Logging(modules.Resource(), "STARTING UP", "START SERVICE", "SERVER IP", "Postgres connected", nil)
	}

	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(50)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			modules.Logging(modules.Resource(), "STARTING UP", "START SERVICE", "SERVER IP", "Failed to close Postgres", err)
		} else {
			modules.Logging(modules.Resource(), "STARTING UP", "START SERVICE", "SERVER IP", "Success to close Postgres", nil)
		}
	}(db)

	errDB = db.Ping()
	if errDB != nil {
		panic(errDB)
	}

	c := cron.New()
	defer c.Stop()

	c.AddFunc("@weekly", func() {
		truncateQuery := `TRUNCATE TABLE public.global_school_data`

		_, errTruncate := db.Exec(truncateQuery)
		if errTruncate != nil {
			modules.Logging(modules.Resource(), "", "SERVER", "", "err when truncate table", errTruncate)
			return
		} else {
			workers := 1000
			guard := make(chan struct{}, workers)
			for i := 1; i < 40; i++ {
				for j := 1; j < 100; j++ {
					for k := 1; k < 100; k++ {
						guard <- struct{}{}
						go func(i int, j int, k int) {
							getData(i, j, k)
							<-guard
						}(i, j, k)
					}
				}
			}
		}

		elapsed := time.Since(start)
		fmt.Printf("Binomial took %s", elapsed)
	})

	go c.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

func getData(provinsi int, kota int, kecamatan int) {
	defer catch()

	tracecode := modules.GenerateUUID()

	mapHeader := make(map[string]interface{})
	mapHeader["Content-Type"] = "application/json; charset=UTF-8"

	theprovinsi := ""
	thekota := ""
	thekecamatan := ""

	theprovinsi = fmt.Sprintf("%02d", provinsi)
	thekota = fmt.Sprintf("%02d", kota)
	thekecamatan = fmt.Sprintf("%02d", kecamatan)

	fmt.Println("PROCESS : ", theprovinsi, thekota, thekecamatan)

	url := THEURL1 + theprovinsi + thekota + thekecamatan + THEURL2

	var client = &http.Client{}
	var data []Sekolah

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		modules.Logging(modules.Resource(), tracecode, "SERVER", "SERVER IP", "error creating request", err)
	}

	response, err := client.Do(request)
	if err != nil {
		modules.Logging(modules.Resource(), tracecode, "SERVER", "SERVER IP", "error do request", err)
	}
	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		modules.Logging(modules.Resource(), tracecode, "SERVER", "SERVER IP", "error read response body", err)
	}

	body := string(b)

	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		modules.Logging(modules.Resource(), tracecode, "SERVER", "SERVER IP", "error unmarshal request", err)
	}

	if response.Status == "200 OK" && len(data) > 0 {
		modules.Logging(modules.Resource(), tracecode, "SERVER", "SERVER IP", "done get school for "+data[0].IndukProvinsi+", "+data[0].IndukKabupaten+", "+data[0].IndukKecamatan+": "+fmt.Sprint(len(data)), nil)
	}

	if len(data) > 0 {

		insertIdQuery := `insert into public.global_school_data(nama, sekolah_id, npsn, jumlahkirim, ptk, pegawai, pd, rombel, jumlah_rk, jumlah_lab, jumlah_perpus, induk_kecamatan, induk_kabupaten, kode_wilayah_induk_kabupaten, kode_wilayah_induk_kecamatan, induk_provinsi, kode_wilayah_induk_provinsi, bentuk_pendidikan, status_sekolah, sinkron_terakhir, sekolah_id_enkrip) values `
		vals := []interface{}{}

		for x := 0; x < len(data); x++ {
			insertIdQuery += "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?),"
			innerInput := data[x]
			vals = append(vals, innerInput.Nama, innerInput.SekolahID, innerInput.Npsn, innerInput.JumlahKirim, innerInput.Ptk, innerInput.Pegawai, innerInput.Pd, innerInput.Rombel, innerInput.JmlRk, innerInput.JmlLab, innerInput.JmlPerpus, innerInput.IndukKecamatan, innerInput.IndukKabupaten, innerInput.KodeWilayahIndukKabupaten, innerInput.KodeWilayahIndukKecamatan, innerInput.IndukProvinsi, innerInput.KodeWilayahIndukProvinsi, innerInput.BentukPendidikan, innerInput.StatusSekolah, innerInput.SinkronTerakhir, innerInput.SekolahIDEnkrip)
		}
		insertIdQuery = strings.TrimSuffix(insertIdQuery, ",")

		//Replacing ? with $n for postgres
		insertIdQuery = modules.ReplaceSQL(insertIdQuery, "?")
		stmt, _ := db.Prepare(insertIdQuery)

		//format all vals at once
		_, errUpdated := stmt.Exec(vals...)
		if errUpdated != nil {
			modules.Logging(modules.Resource(), "", "SERVER", "", "err when inserting data : "+data[0].IndukProvinsi+", "+data[0].IndukKabupaten+", "+data[0].IndukProvinsi, errUpdated)
		}
	}
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
