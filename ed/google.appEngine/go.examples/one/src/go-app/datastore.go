package go_app

import (
	"errors"
	"fmt"
	"github.com/thepkg/gcd"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/search"
	"net/http"
	"time"
)

const (
	htmlExample = `
<!doctype html>
<html lang="en" prefix="op: http://media.facebook.com/op#">
    <head>
        <meta charset="utf-8">

        <link rel="canonical" href="https://www.wefsdfsdfsdf.co.uk/what-colours-to-wear-together">
        <meta property="op:markup_version" content="v1.0">

        <meta property="fb:use_automatic_ad_placement" content="false">
        </head>
    <body>

<article>
    <header>


        <h1>6 New Colour Pairings That Feel Fresh for 2018</h1>


        <h2>The brighter the better.</h2>
        <h3 class="op-kicker">Trends</h3>
        <address>
            {657 Emma Spedding Emma Spedding emma-spedding https://cdn.scrtcdnscr.com/cache/users/657/emma-spedding-657-1509101076-main_image.500x500uc.jpg espedding@cmginc.com}
        </address>

        <time class="op-published" datetime="2018-03-31T05:00:00&#43;0000">March 31st, 2018, 05:00:00 am</time>

        <time class="op-modified" datetime="2018-01-30T12:03:52&#43;0000">January 30th, 2018, 12:03:52 pm</time>
    </header>

    <p>Experimenting with bright colours can be a little daunting, making it all to easy to stick to the usual grey-and-navy and camel-and-black combos. If you aim to wear more colour in 2018, there are a number of ways to wear it that you might not have thought of that will look fresh and unexpected, and are bound to brighten your mood too.</p><p>The <a href="http://www.sldfjsldkfjlskdjflskdjfklsd.co.uk/spring-summer-2017-fashion-trends/slide2" target="_blank">S/S 18 collections</a> are brighter than ever, with powerful pinks, Kermit the Frog greens, sunshine yellows and zesty oranges. The only styling rule when it comes to colour for 2018 is wear lots of it. Orange and bright blue? Sure! Pink and purple? Go for it! Sky blue and scarlet? If it&#39;s good enough for Céline <strong>Scroll below to see six of the key colour combinations for 2018.</strong></p>
    <p><strong>Want to know what else you&#39;ll be wearing this year? Here are the seven key looks for <a href="http://www.sldfjsldkfjlskdjflskdjfklsd.co.uk/spring-summer-2017-fashion-trends" target="_blank">S/S 17</a>…</strong></p>
    <footer>
        <aside>
            Follow Us on <a href="https://www.facebook.com/sldfjsldkfjlskdjflskdjfklsdUK">Facebook</a>
            and <a href="https://twitter.com/sldfjsldkfjlskdjflskdjfklsd">Twitter</a>, and don't miss out on the
            latest by signing up for our <a href="http://www.sldfjsldkfjlskdjflskdjfklsd.co.uk/email/subscribe/">newsletter</a>.
        </aside>
    </footer>


<figure class='op-tracker'>
    <iframe>
        <script>
            var _comscore = _comscore || [];

            if (!window.COMSCORE) {
                window.COMSCORE = {};
                window.COMSCORE.beacon = function (obj) {
                    return false;
                }
            }

            _comscore.push({ c1: '2', c2: '10004700', comscorekw: 'fbia' });
            (function() {
                var s = document.createElement('script'), el = document.getElementsByTagName('script')[0]; s.async = true;
                s.src = 'https://sb.scorecardresearch.com/beacon.js';
                el.parentNode.insertBefore(s, el);
            })();
        </script>
    </iframe>
</figure>

<figure class="op-tracker">
    <iframe>
        <script>
            window.ga=window.ga||function(){(ga.q=ga.q||[]).push(arguments)};ga.l=+new Date;

            ga('create', 'UA-41*****-1', 'auto');
            ga('require', 'displayfeatures');
            ga('send', 'pageview', {
                'dimension1': 'what-colours-to-wear-together',
                'title': '6 New Colour Pairings That Feel Fresh for 2018'


            });
        </script>
        <script async src="https://www.google-analytics.com/analytics.js"></script>
    </iframe>
</figure>

<figure class="op-tracker">
    <iframe>
        <script>
            PARSELY = { autotrack: false, onload: function() { PARSELY.beacon.trackPageView({ urlref: 'http://facebook.com/instantarticles' }); return true; } }
        </script>
        <div id="parsely-root" style="display: none">
            <div id="parsely-cfg" data-parsely-site="sldfjsldkfjlskdjflskdjfklsd.co.uk"></div>
        </div>
        <script>
            (function(s, p, d) {
                var h=d.location.protocol, i=p+'-'+s,
                    e=d.getElementById(i), r=d.getElementById(p+'-root'),
                    u=h==='https:'?'d1z2jf7jlzjs58.cloudfront.net'
                        :'static.'+p+'.com';
                if (e) return;
                e = d.createElement(s); e.id = i; e.async = true;
                e.src = h+'//'+u+'/p.js'; r.appendChild(e);
            })('script', 'parsely', document);
        </script>
    </iframe>
</figure>

<figure class="op-tracker">
    <iframe>
INFO     2018-07-13 13:26:25,371 module.py:846] cws: "GET /articles/cron/syndicate HTTP/1.1" 500 35
        <script>(function(l,d) {
            if (l.search.length){
                var m, u = {}, s = /([^&=]+)=?([^&]*)/g, q = l.search.substring(1);
                while (m = s.exec(q)) u[m[1]] = m[2];
                if (("pefbs" in u) && ("pefba" in u) && ("pefbt" in u)) {
                    var pe = d.createElement("script"); pe.type = "text/javascript"; pe.async = true;
                    pe.src = "http://traffic.pubexchange.com/click/" + u.pefbt + "/" + u.pefbs + "/" + u.pefba;
                    var t = d.getElementsByTagName("script")[0]; t.parentNode.insertBefore(pe, t);
                }
            }
        }(window.location, document));</script>
    </iframe>
</figure>

</article>
</body>
</html>
	`
)

type Visit struct {
	TimeStamp             time.Time               `datastore:"TimeStamp,noindex"`
	Path                  string                  `datastore:"Path,noindex"`
	AdditionalInformation []AdditionalInformation `datastore:"AdditionalInformation"`
}

type AdditionalInformation struct {
	Code  string `datastore:"Code"`
	Value string `datastore:"Value"`
}

type HTML struct {
	//
	Body string `datastore:"Body,noindex"`
}

func datastoreHandler(w http.ResponseWriter, r *http.Request) {
	saveVisit(w, r)
	saveVisit2(w, r)

	datastorePut1(w, r)
	datastorePut2(w, r)
	datastorePut3(w, r)
	transactionCommit(w, r)
	transactionRollBack(w, r)
	transactionPanic(w, r)
	datastoreGetByKey(w, r)
	datastoreGet1(w, r)
	datastoreGet2(w, r)
	indexPut1(w, r)
	indexGet1(w, r)

	indexPutAncestor1(w, r)
	indexGetAncestor1(w, r)

	//datastoreDropKind1(w, r)
	datastoreGetKeys(w, r)
	datastoreDropKind2(w, r)
	datastoreCursor(w, r)
}

func indexGet1(w http.ResponseWriter, r *http.Request) {
	index, err := search.Open("users")
	if err != nil {
		fmt.Fprintf(w, "<br>Get index 1 fail 1, error: %+v", err)
		return
	}

	ctx := appengine.NewContext(r)
	for t := index.Search(ctx, "<br>SearchUser: Id = user8", nil); ; {
		var doc SearchUser
		id, err := t.Next(&doc)
		if err == search.Done {
			break
		}
		if err != nil {
			fmt.Fprintf(w, "<br>Search error: %v\n", err)
			break
		}
		fmt.Fprintf(w, "<br>Get index 1: %s -> %#v\n", id, doc)
	}
}

func indexPut1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<hr>")

	index, err := search.Open("users")
	if err != nil {
		fmt.Fprintf(w, "<br>Put index 1 fail 1, error: %+v", err)
		return
	}

	u := SearchUser{Id: "usr8", Name: "User 8", Comment: "1 more this is <em>marked up</em> text"}

	ctx := appengine.NewContext(r)
	res, err := index.Put(ctx, "usr8", &u)
	if err != nil {
		fmt.Fprintf(w, "<br>Put index 1 fail 2, error: %+v", err)
		return
	}

	fmt.Fprintf(w, "<br>Put index 1 OK: %+v", res)
}

func saveVisit(w http.ResponseWriter, r *http.Request) {
	v := Visit{TimeStamp: time.Now(), Path: "/datastore"}
	ctx := appengine.NewContext(r)
	key := datastore.NewIncompleteKey(ctx, "Visit", nil)
	key, err := datastore.Put(ctx, key, &v)
	if err != nil {
		fmt.Fprintf(w, "<br>Failed to store visit, error: %+v", err)
	}
	fmt.Fprintf(w, "<br>Visit saved with key 🔑: %+v", key)
}

func saveVisit2(w http.ResponseWriter, r *http.Request) {
	ai := []AdditionalInformation{{Code: "c1", Value: "Just a test."}}
	v := Visit{TimeStamp: time.Now(), Path: "/datastore/withAdddInfo", AdditionalInformation: ai}
	ctx := appengine.NewContext(r)
	key := datastore.NewIncompleteKey(ctx, "Visit", nil)
	key, err := datastore.Put(ctx, key, &v)
	if err != nil {
		fmt.Fprintf(w, "<br>Failed to store visit, error: %+v", err)
	}
	fmt.Fprintf(w, "<br>Visit saved with key 🔑: %+v", key)
}

func createUser(ctx context.Context, id string, name string, tag string) (datastore.Key, User, error) {
	user := User{Id: id, Name: name, Tag: tag}
	//key := datastore.NewIncompleteKey(ctx, "User", nil)
	key := datastore.NewKey(ctx, "User", id, 0, nil)
	key, err := datastore.Put(ctx, key, &user)
	return *key, user, err
}

func datastorePut1(w http.ResponseWriter, r *http.Request) (User, User) {
	u := User{Id: "usr4", Name: "User 4", Tag: "cli", Comment: "this is <em>marked up</em> text"}
	ctx := appengine.NewContext(r)

	//key := datastore.NewIncompleteKey(ctx, "User", nil)
	key := datastore.NewKey(ctx, "User", "user4", 0, nil)
	k, err := datastore.Put(ctx, key, &u)
	if err != nil {
		fmt.Fprintf(w, "<br>Error to PUT: %+v", err)
		return User{}, User{}
	}
	fmt.Fprintf(w, "<br>PUT 1 - OK, key: %+v, 🔑: %+v || %+v", key, k, u)

	u2 := User{}
	err = datastore.Get(ctx, key, &u2)
	if err != nil {
		fmt.Fprintf(w, "<br>Error to GET: %+v", err)
		return User{}, User{}
	}
	fmt.Fprintf(w, "<br>GET * - OK: %+v", u2)

	fmt.Fprintf(w, "<hr>")

	return u, u2
}

func datastorePut2(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	key, user, err := createUser(ctx, "x-usr-1", "x1", "x")
	if err != nil {
		fmt.Fprintf(w, "<br>Error to PUT x-usr-1: %+v", err)
		return
	}

	fmt.Fprintf(w, "<br>PUT x-usr-1 - OK: [%+v] %+v", key, user)
}

func datastorePut3(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	d := HTML{Body: htmlExample}
	key := datastore.NewKey(ctx, "HTML", "html1", 0, nil)

	k, err := datastore.Put(ctx, key, &d)
	if err != nil {
		fmt.Fprintf(w, "<br>Error to perform datastorePut3: %+v", err)
		return
	}

	fmt.Fprintf(w, "<br>PUT datastorePut3 - OK: [%+v] %+v", k, key)
}

func transactionCommit(w http.ResponseWriter, r *http.Request) {
	u := User{Id: "usr5", Name: "User 5", Tag: "cli"}
	ctx := appengine.NewContext(r)
	//key := datastore.NewIncompleteKey(ctx, "User", nil)
	key := datastore.NewKey(ctx, "User", "user5", 0, nil)

	err := datastore.RunInTransaction(ctx, func(transactionCTX context.Context) error {
		_, err := datastore.Put(transactionCTX, key, &u)
		// to commit transaction - return nil
		// to rollback transaction - return error
		return err
	}, nil)
	if err == nil {
		fmt.Fprintf(w, "<br>TRANSACTION 1 - ✅ OK, key: %+v | %+v", key, u)
	} else {
		fmt.Fprintf(w, "<br>Transaction 1 🔴 failed, Error: %+v", err)
	}
}

func transactionRollBack(w http.ResponseWriter, r *http.Request) {
	u := User{Id: "usr6", Name: "User 6 {ROLL BACK CASE}", Tag: "cli"}
	ctx := appengine.NewContext(r)
	key := datastore.NewIncompleteKey(ctx, "User", nil)

	err := datastore.RunInTransaction(ctx, func(transactionCTX context.Context) error {
		datastore.Put(transactionCTX, key, &u)
		err := errors.New("error: test ROLL BACK case")
		return err
	}, nil)
	fmt.Fprintf(w, "<br>TRANSACTION 2: %+v", err)
}

func transactionPanic(w http.ResponseWriter, r *http.Request) {
	defer func() {
		r := recover()
		fmt.Fprintf(w, "<br>Transaction 3 RECOVERY 🚑: %+v", r)
	}()
	ctx := appengine.NewContext(r)
	err := datastore.RunInTransaction(ctx, func(transactionCTX context.Context) error {
		u := User{Id: "usr7", Name: "User 7", Tag: "cli"}
		key := datastore.NewIncompleteKey(ctx, "User", nil)
		_, err := datastore.Put(transactionCTX, key, &u)
		panic("panic in transaction")
		return err
	}, nil)
	if err != nil {
		fmt.Fprintf(w, "<br>Transaction 3 failed, Error: %+v", err)
	}
}

func datastoreGetByKey(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	k := datastore.NewKey(ctx, "User", "user5", 0, nil)
	u := User{}
	err := datastore.Get(ctx, k, &u)
	if err != nil {
		fmt.Fprintf(w, "<br>Failed get by Key, error: %+v", err)
		return
	}
	fmt.Fprintf(w, "<hr>Get user by key - OK, user: %+v", u)
}

// field tags contains: test & go
func datastoreGet1(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	q := datastore.
		NewQuery("User").
		Filter("Tags =", "test").
		Filter("Tags =", "go")
	u := make([]User, 0)
	_, err := q.GetAll(ctx, &u)
	if err != nil {
		fmt.Fprintf(w, "<br>Error: %+v", err)
		return
	}
	fmt.Fprintf(w, "<hr>SELECT 1 - OK: %+v", u)
}

func datastoreGet2(w http.ResponseWriter, r *http.Request) {

	//entities := make([]User, 0)
	//ctx := appengine.NewContext(r)
	//err = datastore.GetMulti(ctx, []*datastore.Key{"usr1", "usr2"}, entities)
	//q := datastore.GetMulti()
	//
	//u := []User{{Id: "usr1"}, {Id: "usr2"}}
	//ctx := appengine.NewContext(r)
	//key := datastore.NewIncompleteKey(ctx, "User", nil)
	//k, err := datastore.Put(ctx, key, &u)
}

func datastoreGetKeys(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	keys, err := GetKeys(ctx, "User")
	if err != nil {
		fmt.Fprintf(w, "failed to get keys for kind, error: %v", err)
		return
	}

	fmt.Fprintf(w, "<br>datastoreGetKeys: %+v", keys)
}

func datastoreDropKind2(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	err := gcd.DropKind(ctx, "HTML")
	fmt.Fprintf(w, "<hr>datastoreDropKind, error: %+v", err)
}

func datastoreDropKind1(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	err := DropKind(ctx, "HTML")
	fmt.Fprintf(w, "<hr>datastoreDropKind, error: %+v", err)
}

func DropKind(ctx context.Context, kind string) error {
	keys, err := GetKeys(ctx, kind)
	if err != nil {
		return fmt.Errorf("failed to get keys for kind, error: %v", err)
	}

	err = datastore.DeleteMulti(ctx, keys)
	if err != nil {
		return fmt.Errorf("failed to delete all from kind, error: %v", err)
	}

	return nil
}

func GetKeys(ctx context.Context, kind string) ([]*datastore.Key, error) {
	q := datastore.NewQuery(kind).KeysOnly()
	t := q.Run(ctx)

	keys := make([]*datastore.Key, 0)
	for {
		var d interface{}

		key, err := t.Next(&d)
		if err == datastore.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to fetch next key, error: %v", err)
		}

		keys = append(keys, key)
	}

	return keys, nil
}

func datastoreCursor(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<hr>")
	ctx := appengine.NewContext(r)
	q := datastore.NewQuery("User")

  fmt.Fprintf(w, "datastoreCursor: 🔴 %+v", q)
	cursor := q.Run(ctx)
	for {
		var d User

		_, err := cursor.Next(&d)
		if err == datastore.Done {
			fmt.Fprintf(w, "datastoreCursor: %v", "done")
			return
		}
		if err != nil {
			fmt.Fprintf(w, "datastoreCursor: failed to fetch next key, error: %v", err)
			return
		}

		fmt.Fprintf(w, "<br>datastoreCursor ok: %+v", d)
	}
}
