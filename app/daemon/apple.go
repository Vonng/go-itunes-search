package main

import (
	"fmt"
	"time"
	"bytes"
	"strconv"
)

import (
	"github.com/go-pg/pg"
	"github.com/Vonng/go-itunes-search/app"
	log "github.com/Sirupsen/logrus"
)

// ID type indicator
const (
	TypeITunesID = '!'
	TypeBundleID = '@'
	TypeKeywords = '#'
)

// Message hold msg type with one [optional] leading byte and following ID value.
// If no leading letter of `!@#` is provided, Bundle ID is used as default.
type Message struct {
	Type byte
	ID   string
}

// NewMessage will build message from raw string
func NewMessage(msg string) (m Message) {
	if len(msg) < 2 {
		return
	}

	m.Type, m.ID = msg[0], string(msg[1:])
	if m.Type == TypeITunesID || m.Type == TypeBundleID || m.Type == TypeKeywords {
		return
	} else {
		m.Type = TypeBundleID
		m.ID = msg
	}
	return
}

// Message_Valid tells if this message is valid
func (m *Message) Valid() bool {
	return (m.Type == TypeITunesID || m.Type == TypeBundleID || m.Type == TypeKeywords) && len(m.ID) > 0
}

// Global postgreSQL instance
var Pg = pg.Connect(&pg.Options{
	Addr:     ":5432",
	Database: "meta",
	User:     "meta",
	Password: "meta",
})

// SeenID will check whether given iTunesID is already in database
func SeenID(id int64) bool {
	var res int64
	_, err := Pg.Query(&res, `SELECT count(id) FROM apple WHERE id =`+strconv.FormatInt(id, 10))
	if err == nil && res == 1 {
		return true
	}
	return false
}

// HandleAppleByID will fetch and save apple application to database according to iTunesID
func HandleAppleByID(id string) error {
	numID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	if apple, err := app.NewAppByID(numID, "CN"); err != nil {
		return err
	} else {
		return apple.Save(Pg)
	}
}

// HandleAppleByBundleID work just like HandleAppleByID but fetch app using bundleID
func HandleAppleByBundleID(bundleID string) error {
	if apple, err := app.NewAppByBundleID(bundleID, "CN"); err != nil {
		return err
	} else {
		return apple.Save(Pg)
	}
}

// HandleApplesByKeyword find a series of app returned by iTunes Search API
// and put them into queue
func HandleApplesByKeyword(keyword string) error {
	apples, err := app.NewAppsByKeyword(keyword, "CN")
	if err != nil {
		return err
	}

	if len(apples) == 0 {
		return nil
	}

	var sql bytes.Buffer
	cnt := 0
	sql.WriteString("INSERT INTO apple_queue(id) VALUES ")
	for i, apple := range apples {
		if SeenID(apple.ID) {
			continue
		}
		cnt += 1
		if i > 0 {
			sql.WriteByte(',')
		}
		sql.WriteString(`('A`)
		sql.WriteString(strconv.FormatInt(apple.ID, 10))
		sql.WriteString(`')`)
	}
	if cnt == 0 {
		return nil
	}
	sql.WriteString("ON CONFLICT DO NOTHING;")
	res, err := Pg.Exec(sql.String())
	if err != nil {
		return err
	}
	log.Infof("[SEARCH] keyword %s found %d, add %d", keyword, len(apples), res.RowsAffected())
	return nil
}

// Producer will pull task from PostgreSQL table `apple_queue`
func Producer() <-chan Message {
	log.Info("[PROD] initializing...")
	stmt, err := Pg.Prepare(`DELETE FROM apple_queue WHERE id IN (SELECT id FROM apple_queue LIMIT 100) RETURNING id;`)
	if err != nil {
		log.Info("[PROD] prepare job statement failed...check postgres instance")
		return nil
	}
	c := make(chan Message)
	go func(chan<- Message) {
		sleep := time.Second
		for {
			var ids []string
			_, err := stmt.Query(&ids)
			if len(ids) == 0 {
				log.Infof("[PROD] empty queue. sleep %d s", sleep/1e9)
				time.Sleep(sleep)
				if sleep < 30*time.Second {
					sleep *= 2
				}
				continue
			} else {
				// reset sleep counter to 1s
				sleep = time.Second
			}

			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			for _, id := range ids {
				if msg := NewMessage(id); msg.Valid() {
					c <- msg
				}
			}
		}
	}(c)
	log.Info("[PROD] init complete")
	return c
}

// Worker will handle incoming task
func Worker(id int, c <-chan Message) {
	log.Infof("[WORKER:%d] init", id)
	var err error
	for msg := range c {
		switch msg.Type {
		case TypeITunesID:
			log.Infof("[WORKER:%d] handle iTunesID=%s", id, msg.ID)
			if err = HandleAppleByID(msg.ID); err != nil {
				log.Errorf("[WORKER:%d] handle iTunesID=%s failed: %s", id, msg.ID, err.Error())
			}
			log.Infof("[WORKER:%d] done iTunesID=%s", id, msg.ID)
		case TypeBundleID:
			log.Infof("[WORKER:%d] handle BundleID=%s", id, msg.ID)
			if err = HandleAppleByBundleID(msg.ID); err != nil {
				log.Errorf("[WORKER:%d] handle BundleID=%s failed: %s", id, msg.ID, err.Error())
			}
			log.Infof("[WORKER:%d] done BundleID=%s", id, msg.ID)
		case TypeKeywords:
			log.Infof("[WORKER:%d] handle Keyword=%s", id, msg.ID)
			if err = HandleApplesByKeyword(msg.ID); err != nil {
				log.Errorf("[WORKER:%d] handle Keyword=%s failed: %s", id, msg.ID, err.Error())
			}
			log.Infof("[WORKER:%d] done keyword=%s", id, msg.ID)
		}
	}
	log.Infof("[WORK] %d finish", id)
}

// Run will start n worker and one producer.
func Run(n int) {
	log.Infof("[RUN] init with %d worker...", n)
	c := Producer()
	for i := 1; i <= n; i++ {
		go Worker(i, c)
	}
}

func main() {
	log.SetLevel(log.DebugLevel)

	Run(5)
	<-make(chan bool)
}
