package gram

import (
	"net/http"

	"github.com/apex/log"
	"github.com/garyburd/redigo/redis"
	"github.com/ilgooz/gram/telegram"
)

type Gram struct {
	decoder StateDecoder
	encoder StateEncoder

	options *options
	tg      *telegram.Telegram
	rds     redis.Conn
}

type options struct {
	redisAddr string
	redisDB   string
	bots      []string
}

type Option func(*Gram)

type StateDecoder interface {
	Decode(data []byte) error
}

type StateEncoder interface {
	Encode(data []byte) error
}

func New(ops ...Option) (*Gram, error) {
	// we should wait for some time after a network failure before retrying to execute our code.
	// time duration we retry count seç

	// error ve step number state bultiple versiyonda bi yerde tutulmalı, marshall, unmarshall support lazım
	// default in memory olsun. interface sun. memorynin impelementasyonunu yap
	// diğer state verisi için redis örneği hazırla

	// bot seviyesinde ayarlanabilir şekilde send ve promp vs için yap.
	//
	// func do(f func() error) error {
	// 	b := &backoff.Backoff{
	// 		Min:    100 * time.Millisecond,
	// 		Max:    10 * time.Minute,
	// 		Factor: 2,
	// 	}

	// 	return try.Do(func(attempt int) (bool, error) {
	// 		err := f()
	// 		if err != nil {
	// 			time.Sleep(b.Duration())
	// 		}
	// 		return attempt < 5, err
	// 	})
	// }

	g := &Gram{
		options: &options{},
	}

	for _, o := range ops {
		o(g)
	}

	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return g, err
	}
	g.rds = c

	return g, nil

}

func RedisOption(addr, db string) Option {
	return func(g *Gram) {
		g.options.redisAddr = addr
		g.options.redisDB = db
	}
}

func BotOption(key string) Option {
	return func(g *Gram) {
		g.options.bots = append(g.options.bots, key)
	}
}

func EnableDistributedModeOption() Option {
	return func(g *Gram) {

	}
}

func GracefulSendOption() Option {
	return func(g *Gram) {
		// use GracefullSendFuncOption here to implement backoff sleeps
	}
}

// func sleep or chan wait?
func GracefulSendFuncOption(func()) Option {
	return func(g *Gram) {

	}
}

func (g *Gram) Attach(pipeline *Pipeline, otherPipelines ...*Pipeline) {

}

type webhookHandler struct{}

func (wh *webhookHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

func (g *Gram) WebhookHandler() http.Handler {
	return &webhookHandler{}
}

func (g *Gram) StartPooling() error {
	key := g.options.bots[0]
	tg := telegram.New(key)

	data, _ := redis.Int64(g.rds.Do("GET", "nextUpdate"))

	next := int(data)

	updates, err := tg.GetUpdates(next)
	if err != nil {
		log.Error(err.Error())
	}
	if len(updates) > 0 {
		u := updates[len(updates)-1]
		_, err := g.rds.Do("SET", "nextUpdate", u.ID+1)
		if err != nil {
			return err
		}
	}

	g.tg = tg

	return nil
}

func (g *Gram) StopPooling() error {
	return nil
}

func (g *Gram) UpdateCommands() error {
	return nil
}
