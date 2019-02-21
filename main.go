package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"time"

	cors "github.com/heppu/simple-cors"
)

var currency Response

type Response struct {
	Source string

	Quotes struct {
		// people *people

		// //A

		USDAED float64 `json:"USDAED"`
		USDAFN float64 `json:"USDAFN"`
		USDALL float64 `json:"USDALL"`
		USDAMD float64 `json:"USDAMD"`
		USDANG float64 `json:"USDANG"`
		USDAOA float64 `json:"USDAOA"`
		USDARS float64 `json:"USDARS"`
		USDAUD float64 `json:"USDAUD"`
		USDAWG float64 `json:"USDAWG"`
		USDAZN float64 `json:"USDAZN"`

		// //B
		USDBAM float64 `json:"USDBAM"`
		USDBBD float64 `json:"USDBBD"`
		USDBDT float64 `json:"USDBDT"`
		USDBGN float64 `json:"USDBGN"`
		USDBHD float64 `json:"USDBHD"`
		USDBIF float64 `json:"USDBIF"`
		USDBMD float64 `json:"USDBMD"`
		USDBND float64 `json:"USDBND"`
		USDBOB float64 `json:"USDBOB"`
		USDBRL float64 `json:"USDBRL"`
		USDBSD float64 `json:"USDBSD"`
		USDBTC float64 `json:"USDBTC"`
		USDBTN float64 `json:"USDBTN"`
		USDBWP float64 `json:"USDBWP"`
		USDBYN float64 `json:"USDBYN"`
		USDBYR float64 `json:"USDBYR"`
		USDBZD float64 `json:"USDBZD"`

		//C
		USDCAD float64 `json:"USDCAD"`
		USDCDF float64 `json:"USDCDF"`
		USDCHF float64 `json:"USDCHF"`
		USDCLF float64 `json:"USDCLF"`
		USDCLP float64 `json:"USDCLP"`
		USDCNY float64 `json:"USDCNY"`
		USDCOP float64 `json:"USDCOP"`
		USDCRC float64 `json:"USDCRC"`
		USDCUC float64 `json:"USDCUC"`
		USDCUP float64 `json:"USDCUP"`
		USDCVE float64 `json:"USDCVE"`
		USDCZK float64 `json:"USDCZK"`

		//D
		USDDJF float64 `json:"USDDJF"`
		USDDKK float64 `json:"USDDKK"`
		USDDOP float64 `json:"USDDOP"`
		USDDZD float64 `json:"USDDZD"`

		// //E
		USDEGP float64 `json:"USDEGP"`
		USDERN float64 `json:"USDERN"`
		USDETB float64 `json:"USDETB"`
		USDEUR float64 `json:"USDEUR"`

		//F
		USDFJD float64 `json:"USDFJD"`
		USDFKP float64 `json:"USDFKP"`

		//G
		USDGBP float64 `json:"USDGBP"`
		USDGEL float64 `json:"USDGEL"`
		USDGGP float64 `json:"USDGGP"`
		USDGHS float64 `json:"USDGHS"`
		USDGIP float64 `json:"USDGIP"`
		USDGMD float64 `json:"USDGMD"`
		USDGNF float64 `json:"USDGNF"`
		USDGTQ float64 `json:"USDGTQ"`
		USDGYD float64 `json:"USDGYD"`

		//H
		USDHKD float64 `json:"USDHKD"`
		USDHNL float64 `json:"USDHNL"`
		USDHRK float64 `json:"USDHRK"`
		USDHTG float64 `json:"USDHTG"`
		USDHUF float64 `json:"USDHUF"`

		//I
		USDIDR float64 `json:"USDIDR"`
		USDILS float64 `json:"USDILS"`
		USDIMP float64 `json:"USDIMP"`
		USDINR float64 `json:"USDINR"`
		USDIQD float64 `json:"USDIQD"`
		USDIRR float64 `json:"USDIRR"`
		USDISK float64 `json:"USDISK"`

		//J
		USDJEP float64 `json:"USDJEP"`
		USDJMD float64 `json:"USDJMD"`
		USDJOD float64 `json:"USDJOD"`
		USDJPY float64 `json:"USDJPY"`

		//K
		USDKES float64 `json:"USDKES"`
		USDKGS float64 `json:"USDKGS"`
		USDKHR float64 `json:"USDKHR"`
		USDKMF float64 `json:"USDKMF"`
		USDKPW float64 `json:"USDKPW"`
		USDKRW float64 `json:"USDKRW"`
		USDKWD float64 `json:"USDKWD"`
		USDKYD float64 `json:"USDKYD"`
		USDKZT float64 `json:"USDKZT"`

		//L
		USDLAK float64 `json:"USDLAK"`
		USDLBP float64 `json:"USDLBP"`
		USDLKR float64 `json:"USDLKR"`
		USDLRD float64 `json:"USDLRD"`
		USDLSL float64 `json:"USDLSL"`
		USDLTL float64 `json:"USDLTL"`
		USDLVL float64 `json:"USDLVL"`
		USDLYD float64 `json:"USDLYD"`

		//M
		USDMAD float64 `json:"USDMAD"`
		USDMDL float64 `json:"USDMDL"`
		USDMGA float64 `json:"USDMGA"`
		USDMKD float64 `json:"USDMKD"`
		USDMMK float64 `json:"USDMMK"`
		USDMNT float64 `json:"USDMNT"`
		USDMOP float64 `json:"USDMOP"`
		USDMRO float64 `json:"USDMRO"`
		USDMUR float64 `json:"USDMUR"`
		USDMVR float64 `json:"USDMVR"`
		USDMWK float64 `json:"USDMWK"`
		USDMXN float64 `json:"USDMXN"`
		USDMYR float64 `json:"USDMYR"`
		USDMZN float64 `json:"USDMZN"`

		//N
		USDNAD float64 `json:"USDNAD"`
		USDNGN float64 `json:"USDNGN"`
		USDNIO float64 `json:"USDNIO"`
		USDNOK float64 `json:"USDNOK"`
		USDNPR float64 `json:"USDNPR"`
		USDNZD float64 `json:"USDNZD"`

		// //O
		USDOMR float64 `json:"USDOMR"`

		//P
		USDPAB float64 `json:"USDPAB"`
		USDPEN float64 `json:"USDPEN"`
		USDPGK float64 `json:"USDPGK"`
		USDPHP float64 `json:"USDPHP"`
		USDPKR float64 `json:"USDPKR"`
		USDPLN float64 `json:"USDPLN"`
		USDPYG float64 `json:"USDPYG"`

		//Q
		USDQAR float64 `json:"USDQAR"`

		//R
		USDRON float64 `json:"USDRON"`
		USDRSD float64 `json:"USDRSD"`
		USDRUB float64 `json:"USDRUB"`
		USDRWF float64 `json:"USDRWF"`

		//S
		USDSAR float64 `json:"USDSAR"`
		USDSBD float64 `json:"USDSBD"`
		USDSCR float64 `json:"USDSCR"`
		USDSDG float64 `json:"USDSDG"`
		USDSEK float64 `json:"USDSEK"`
		USDSGD float64 `json:"USDSGD"`
		USDSHP float64 `json:"USDSHP"`
		USDSLL float64 `json:"USDSLL"`
		USDSOS float64 `json:"USDSOS"`
		USDSRD float64 `json:"USDSRD"`
		USDSTD float64 `json:"USDSTD"`
		USDSVC float64 `json:"USDSVC"`
		USDSYP float64 `json:"USDSYP"`
		USDSZL float64 `json:"USDSZL"`

		//T
		USDTHB float64 `json:"USDTHB"`
		USDTJS float64 `json:"USDTJS"`
		USDTMT float64 `json:"USDTMT"`
		USDTND float64 `json:"USDTND"`
		USDTOP float64 `json:"USDTOP"`
		USDTRY float64 `json:"USDTRY"`
		USDTTD float64 `json:"USDTTD"`
		USDTWD float64 `json:"USDTWD"`
		USDTZS float64 `json:"USDTZS"`

		//U
		USDUAH float64 `json:"USDUAH"`
		USDUGX float64 `json:"USDUGX"`
		USDUSD float64 `json:"USDUSD"`
		USDUYU float64 `json:"USDUYU"`
		USDUZS float64 `json:"USDUZS"`

		// //V
		USDVEF float64 `json:"USDVEF"`
		USDVND float64 `json:"USDVND"`
		USDVUV float64 `json:"USDVUV"`

		//W
		USDWST float64 `json:"USDWST"`

		//X
		USDXAF float64 `json:"USDXAF"`
		USDXAG float64 `json:"USDXAG"`
		USDXAU float64 `json:"USDXAU"`
		USDXCD float64 `json:"USDXCD"`
		USDXDR float64 `json:"USDXDR"`
		USDXOF float64 `json:"USDXOF"`
		USDXPF float64 `json:"USDXPF"`

		//Y
		USDYER float64 `json:"USDYER"`

		//Z
		USDZAR float64 `json:"USDZAR"`
		USDZMK float64 `json:"USDZMK"`
		USDZMW float64 `json:"USDZMW"`
		USDZWL float64 `json:"USDZWL"`
	}
}

//import _ "github.com/joho/godotenv/autoload"

var (
	// WelcomeMessage A constant to hold the welcome message
	WelcomeMessage = "Welcome, what do you want to order?"

	// sessions = {
	//   "uuid1" = Session{...},
	//   ...
	// }
	sessions = map[string]Session{}

	processor = sampleProcessor
)

type (
	// Session Holds info about a session
	Session map[string]interface{}

	// JSON Holds a JSON object
	JSON map[string]interface{}

	// Processor Alias for Process func
	Processor func(session Session, message string) (string, error)
)

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func sampleProcessor(session Session, message string) (string, error) {
	{

		postSplit := strings.Split(message, " ")
		fmt.Println(postSplit[0])
		postSplit[1] = strings.ToUpper(postSplit[1])
		postSplit[2] = strings.ToUpper(postSplit[2])
		//url := "https://openexchangerates.org/api/convert/" + postSplit[0] + "/" + postSplit[1] + "/" + postSplit[2] + "?app_id=f7ae914f94ab440583268529f2701292"
		url := "http://www.apilayer.net/api/live?access_key=3f5cc35110aff4a24f2b825b0ebca08b&format=1"
		spaceClient := http.Client{
			Timeout: time.Second * 200, // Maximum of 2 secs
		}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("User-Agent", "spacecount-tutorial")

		res, getErr := spaceClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}

		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		jsonErr := json.Unmarshal(body, &currency)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		var i float64 = 0

		x, err := strconv.Atoi(postSplit[0])
		if err != nil {

			return "", fmt.Errorf(postSplit[0] + " is not a number")

		}

		f := float64(x)

		from := strings.ToUpper(postSplit[1])

		switch from {

		//USDEGP float64 `json:"USDEGP"`
		// USDERN float64 `json:"USDERN"`
		// USDETB float64 `json:"USDETB"`
		// USDEUR float64 `json:"USDEUR"`

		case "EGP":
			{
				i = f / currency.Quotes.USDEGP
			}
		case "ERN":
			{
				i = f / currency.Quotes.USDERN
			}
		case "ETB":
			{
				i = f / currency.Quotes.USDETB
			}
		case "EUR":
			{
				i = f / currency.Quotes.USDEUR
			}

		// USDZAR float64 `json:"USDZAR"`
		// USDZMK float64 `json:"USDZMK"`
		// USDZMW float64 `json:"USDZMW"`
		// USDZWL float64 `json:"USDZWL"`

		case "ZAR":
			{
				i = f / currency.Quotes.USDZAR
			}
		case "ZMK":
			{
				i = f / currency.Quotes.USDZMK
			}
		case "ZMW":
			{
				i = f / currency.Quotes.USDZMW
			}
		case "ZWL":
			{
				i = f / currency.Quotes.USDZWL
			}

		case "XAF":
			{
				i = f / currency.Quotes.USDXAF
			}
		case "XAG":
			{
				i = f / currency.Quotes.USDXAG
			}
		case "XAU":
			{
				i = f / currency.Quotes.USDXAU
			}
		case "XCD":
			{
				i = f / currency.Quotes.USDXCD
			}
		case "XDR":
			{
				i = f / currency.Quotes.USDXDR
			}
		case "XOF":
			{
				i = f / currency.Quotes.USDXOF
			}
		case "XPF":
			{
				i = f / currency.Quotes.USDXPF
			}
			// USD float64 `json:"USDXAG"`
			// USDfloat64 `json:"USDXAU"`
			// USDXCD float64 `json:"USDXCD"`
			// USD float64 `json:"USDXDR"`
			// USD float64 `json:"USDXOF"`
			// USD float64 `json:"USDXPF"`

		// USD float64 `json:"USDAED"`
		// USD float64 `json:"USDAFN"`
		// USD float64 `json:"USDALL"`
		// USD float64 `json:"USDAMD"`
		// USD float64 `json:"USDANG"`
		// USD float64 `json:"USDAOA"`
		// USDARS float64 `json:"USDARS"`
		// USDAUD float64 `json:"USDAUD"`
		// USDAWG float64 `json:"USDAWG"`
		// USDAZN float64 `json:"USDAZN"`

		case "AED":
			{
				i = f / currency.Quotes.USDAED
			}
		case "AFN":
			{
				i = f / currency.Quotes.USDAFN
			}
		case "ALL":
			{
				i = f / currency.Quotes.USDALL
			}
		case "AMD":
			{
				i = f / currency.Quotes.USDAMD
			}
		case "ANG":
			{
				i = f / currency.Quotes.USDANG
			}
		case "AOA":
			{
				i = f / currency.Quotes.USDAOA
			}
		case "ARS":
			{
				i = f / currency.Quotes.USDARS
			}
		case "AUD":
			{
				i = f / currency.Quotes.USDAUD
			}
		case "AWG":
			{
				i = f / currency.Quotes.USDAWG
			}
		case "AZN":
			{
				i = f / currency.Quotes.USDAZN
			}

		case "BAM":
			{
				i = f / currency.Quotes.USDBAM
			}
		case "BBD":
			{
				i = f / currency.Quotes.USDBBD
			}
		case "BDT":
			{
				i = f / currency.Quotes.USDBDT
			}
		case "BGN":
			{
				i = f / currency.Quotes.USDBGN
			}

		case "BHD":
			{
				i = f / currency.Quotes.USDBHD
			}
		case "BIF":
			{
				i = f / currency.Quotes.USDBIF
			}
		case "BMD":
			{
				i = f / currency.Quotes.USDBMD
			}
		case "BND":
			{
				i = f / currency.Quotes.USDBND
			}
		case "BOB":
			{
				i = f / currency.Quotes.USDBOB
			}
		case "BRL":
			{
				i = f / currency.Quotes.USDBRL
			}
		case "BSD":
			{
				i = f / currency.Quotes.USDBSD
			}
		case "BTC":
			{
				i = f / currency.Quotes.USDBTC
			}
		case "BTN":
			{
				i = f / currency.Quotes.USDBTN
			}
		case "BWP":
			{
				i = f / currency.Quotes.USDBWP
			}
		case "BYN":
			{
				i = f / currency.Quotes.USDBYN
			}
		case "BYR":
			{
				i = f / currency.Quotes.USDBYR
			}
		case "BZD":
			{
				i = f / currency.Quotes.USDBZD
			}

		case "CAD":
			{
				i = f / currency.Quotes.USDCAD
			}
		case "CDF":
			{
				i = f / currency.Quotes.USDCHF
			}
		case "CHF":
			{
				i = f / currency.Quotes.USDXAU
			}
		case "YER":
			{
				i = f / currency.Quotes.USDYER
			}

		case "FJD":
			{
				i = f / currency.Quotes.USDFJD
			}
		case "FKP":
			{
				i = f / currency.Quotes.USDFKP
			}
		case "GBP":
			{
				i = f / currency.Quotes.USDGBP
			}
		case "GEL":
			{
				i = f / currency.Quotes.USDGEL
			}
		case "GGP":
			{
				i = f / currency.Quotes.USDGGP
			}
		case "GHS":
			{
				i = f / currency.Quotes.USDGHS
			}
		case "GIP":
			{
				i = f / currency.Quotes.USDGIP
			}
		case "GMD":
			{
				i = f / currency.Quotes.USDGMD
			}
		case "GNF":
			{
				i = f / currency.Quotes.USDGNF
			}
		case "GTQ":
			{
				i = f / currency.Quotes.USDGTQ
			}
		case "GYD":
			{
				i = f / currency.Quotes.USDGYD
			}
		case "HKD":
			{
				i = f / currency.Quotes.USDHKD
			}
		case "HNL":
			{
				i = f / currency.Quotes.USDHNL
			}
		case "HRK":
			{
				i = f / currency.Quotes.USDHRK
			}
		case "HTG":
			{
				i = f / currency.Quotes.USDHTG
			}
		case "HUF":
			{
				i = f / currency.Quotes.USDHUF
			}
		case "IDR":
			{
				i = f / currency.Quotes.USDIDR
			}
		case "ILS":
			{
				i = f / currency.Quotes.USDILS
			}
		case "IMP":
			{
				i = f / currency.Quotes.USDIMP
			}
		case "INR":
			{
				i = f / currency.Quotes.USDINR
			}
		case "IQD":
			{
				i = f / currency.Quotes.USDIQD
			}
		case "IRR":
			{
				i = f / currency.Quotes.USDIRR
			}
		case "ISK":
			{
				i = f / currency.Quotes.USDISK
			}
		case "JEP":
			{
				i = f / currency.Quotes.USDJEP
			}
		case "JMD":
			{
				i = f / currency.Quotes.USDJMD
			}
		case "JOD":
			{
				i = f / currency.Quotes.USDJOD
			}
		case "JPY":
			{
				i = f / currency.Quotes.USDJPY
			}
		case "KES":
			{
				i = f / currency.Quotes.USDKES
			}
		case "KGS":
			{
				i = f / currency.Quotes.USDKGS
			}
		case "KHR":
			{
				i = f / currency.Quotes.USDKHR
			}
		case "KMF":
			{
				i = f / currency.Quotes.USDKMF
			}
		case "KPW":
			{
				i = f / currency.Quotes.USDKPW
			}
		case "KRW":
			{
				i = f / currency.Quotes.USDKRW
			}
		case "KWD":
			{
				i = f / currency.Quotes.USDKWD
			}
		case "KYD":
			{
				i = f / currency.Quotes.USDKYD
			}

		case "KZT":
			{
				i = f / currency.Quotes.USDKZT
			}
		case "LAK":
			{
				i = f / currency.Quotes.USDLAK
			}
		case "LBP":
			{
				i = f / currency.Quotes.USDLBP
			}
		case "LKR":
			{
				i = f / currency.Quotes.USDLKR
			}
		case "LRD":
			{
				i = f / currency.Quotes.USDLRD
			}
		case "LSL":
			{
				i = f / currency.Quotes.USDLSL
			}
		case "LTL":
			{
				i = f / currency.Quotes.USDLTL
			}
		case "LVL":
			{
				i = f / currency.Quotes.USDLVL
			}
		case "LYD":
			{
				i = f / currency.Quotes.USDLYD
			}
		case "MAD":
			{
				i = f / currency.Quotes.USDMAD
			}
		case "MDL":
			{
				i = f / currency.Quotes.USDMDL
			}
		case "MGA":
			{
				i = f / currency.Quotes.USDMGA
			}
		case "MKD":
			{
				i = f / currency.Quotes.USDMKD
			}
		case "MMK":
			{
				i = f / currency.Quotes.USDMMK
			}
		case "MNT":
			{
				i = f / currency.Quotes.USDMNT
			}
		case "MOP":
			{
				i = f / currency.Quotes.USDMOP
			}
		case "MRO":
			{
				i = f / currency.Quotes.USDMRO
			}
		case "MUR":
			{
				i = f / currency.Quotes.USDMUR
			}
		case "MVR":
			{
				i = f / currency.Quotes.USDMVR
			}
		case "MWK":
			{
				i = f / currency.Quotes.USDMWK
			}
		case "MXN":
			{
				i = f / currency.Quotes.USDMXN
			}
		case "MYR":
			{
				i = f / currency.Quotes.USDMYR
			}
		case "MZN":
			{
				i = f / currency.Quotes.USDMZN
			}
		case "NAD":
			{
				i = f / currency.Quotes.USDNAD
			}
		case "NGN":
			{
				i = f / currency.Quotes.USDNGN
			}
		case "NIO":
			{
				i = f / currency.Quotes.USDNIO
			}
		case "NOK":
			{
				i = f / currency.Quotes.USDNOK
			}
		case "NPR":
			{
				i = f / currency.Quotes.USDNPR
			}
		case "NZD":
			{
				i = f / currency.Quotes.USDNZD
			}
		case "OMR":
			{
				i = f / currency.Quotes.USDOMR
			}

		// P

		case "PAB":
			{
				i = f / currency.Quotes.USDPAB
			}
		// ====================== //

		case "PEN":
			{
				i = f / currency.Quotes.USDPEN
			}
		// ====================== //
		case "PGK":
			{
				i = f / currency.Quotes.USDPGK
			}
		// ====================== //
		case "PHP":
			{
				i = f / currency.Quotes.USDPHP
			}
		// ====================== //
		case "PKR":
			{
				i = f / currency.Quotes.USDPKR
			}
		// ====================== //
		case "PLN":
			{
				i = f / currency.Quotes.USDPLN
			}
		// ====================== //
		case "PYG":
			{
				i = f / currency.Quotes.USDPYG
			}

		// ====================== //

		// ==============================================//

		//Q

		case "QAR":
			{
				i = f / currency.Quotes.USDQAR
			}
		// ====================== //

		// ==============================================//

		//R
		case "RON":
			{
				i = f / currency.Quotes.USDRON
			}
		// ====================== //
		case "RSD":
			{
				i = f / currency.Quotes.USDRSD
			}
		// ====================== //
		case "RUB":
			{
				i = f / currency.Quotes.USDRUB
			}
		// ====================== //
		case "RWF":
			{
				i = f / currency.Quotes.USDRWF
			}
		// ====================== //

		// ====================== //

		// ==============================================//

		//S
		case "SAR":
			{
				i = f / currency.Quotes.USDSAR
			}
		// ====================== //
		case "SBD":
			{
				i = f / currency.Quotes.USDSBD
			}
		// ====================== //
		case "SCR":
			{
				i = f / currency.Quotes.USDSCR
			}
		// ====================== //
		case "SDG":
			{
				i = f / currency.Quotes.USDSDG
			}
		// ====================== //
		case "SEK":
			{
				i = f / currency.Quotes.USDSEK
			}
		// ====================== //
		case "SHP":
			{
				i = f / currency.Quotes.USDSHP
			}
		// ====================== //
		case "SLL":
			{
				i = f / currency.Quotes.USDSLL
			}
		// ====================== //
		case "SOS":
			{
				i = f / currency.Quotes.USDSOS
			}
		// ====================== //
		case "SRD":
			{
				i = f / currency.Quotes.USDSRD
			}
		// ====================== //
		case "STD":
			{
				i = f / currency.Quotes.USDSTD
			}
		// ====================== //
		case "SVC":
			{
				i = f / currency.Quotes.USDSVC
			}
		// ====================== //
		case "SYP":
			{
				i = f / currency.Quotes.USDSYP
			}
		// ====================== //
		case "SZL":
			{
				i = f / currency.Quotes.USDSZL
			}
		// ====================== //
		// ==============================================//

		//T

		case "THB":
			{
				i = f / currency.Quotes.USDTHB
			}
		// ====================== //
		case "TJS":
			{
				i = f / currency.Quotes.USDTJS
			}
		// ====================== //
		case "TMT":
			{
				i = f / currency.Quotes.USDTMT
			}
		// ====================== //
		case "TND":
			{
				i = f / currency.Quotes.USDTND
			}
		// ====================== //
		case "TOP":
			{
				i = f / currency.Quotes.USDTOP
			}
		// ====================== //
		case "TRY":
			{
				i = f / currency.Quotes.USDTRY
			}
		// ====================== //
		case "TTD":
			{
				i = f / currency.Quotes.USDTTD
			}
		// ====================== //
		case "TWD":
			{
				i = f / currency.Quotes.USDTWD
			}
		// ====================== //
		case "TZS":
			{
				i = f / currency.Quotes.USDTZS
			}
		// ====================== //
		// ==============================================//

		//U

		case "UAH":
			{
				i = f / currency.Quotes.USDUAH
			}
		// ====================== //
		case "UGX":
			{
				i = f / currency.Quotes.USDUGX
			}
		// ====================== //
		case "USD":
			{
				i = f / currency.Quotes.USDUSD
			}
		// ====================== //
		case "UYU":
			{
				i = f / currency.Quotes.USDUYU
			}
		// ====================== //
		case "UZS":
			{
				i = f / currency.Quotes.USDUZS
			}
			// ====================== //

			// USDCAD float64 `json:"USDCAD"`
			// USDCDF float64 `json:"USDCDF"`
			// USDCHF float64 `json:"USDCHF"`
			// USDCLF float64 `json:"USDCLF"`
			// USDCLP float64 `json:"USDCLP"`
			// USDCNY float64 `json:"USDCNY"`
			// USDCOP float64 `json:"USDCOP"`
			// USDCRC float64 `json:"USDCRC"`
			// USDCUC float64 `json:"USDCUC"`
			// USDCUP float64 `json:"USDCUP"`
			// USDCVE float64 `json:"USDCVE"`
			// USDCZK float64 `json:"USDCZK"`

		default:
			{
				return "", fmt.Errorf(from + " Currency Not Found Please Re enter")
			}
		}

		to := strings.ToUpper(postSplit[2])

		switch to {
		case "EGP":
			{

				i = i * currency.Quotes.USDEGP
				//s := strconv.Itoa(i)

				fmt.Println(i)
			}

		case "ERN":
			{
				i = i * currency.Quotes.USDERN
			}
		case "ETB":
			{
				i = i * currency.Quotes.USDETB
			}
		case "EUR":
			{
				i = i * currency.Quotes.USDEUR
			}

		// USDZAR float64 `json:"USDZAR"`
		// USDZMK float64 `json:"USDZMK"`
		// USDZMW float64 `json:"USDZMW"`
		// USDZWL float64 `json:"USDZWL"`

		case "ZAR":
			{
				i = i * currency.Quotes.USDZAR
			}
		case "ZMK":
			{
				i = i * currency.Quotes.USDZMK
			}
		case "ZMW":
			{
				i = i * currency.Quotes.USDZMW
			}
		case "ZWL":
			{
				i = i * currency.Quotes.USDZWL
			}

		case "XAF":
			{
				i = i * currency.Quotes.USDXAF
			}
		case "XAG":
			{
				i = i * currency.Quotes.USDXAG
			}
		case "XAU":
			{
				i = i * currency.Quotes.USDXAU
			}
		case "XCD":
			{
				i = i * currency.Quotes.USDXCD
			}
		case "XDR":
			{
				i = i * currency.Quotes.USDXDR
			}
		case "XOF":
			{
				i = i * currency.Quotes.USDXOF
			}
		case "XPF":
			{
				i = i * currency.Quotes.USDXPF
			}
			// USD float64 `json:"USDXAG"`
			// USDfloat64 `json:"USDXAU"`
			// USDXCD float64 `json:"USDXCD"`
			// USD float64 `json:"USDXDR"`
			// USD float64 `json:"USDXOF"`
			// USD float64 `json:"USDXPF"`

		// USD float64 `json:"USDAED"`
		// USD float64 `json:"USDAFN"`
		// USD float64 `json:"USDALL"`
		// USD float64 `json:"USDAMD"`
		// USD float64 `json:"USDANG"`
		// USD float64 `json:"USDAOA"`
		// USDARS float64 `json:"USDARS"`
		// USDAUD float64 `json:"USDAUD"`
		// USDAWG float64 `json:"USDAWG"`
		// USDAZN float64 `json:"USDAZN"`

		case "AED":
			{
				i = i * currency.Quotes.USDAED
			}
		case "AFN":
			{
				i = i * currency.Quotes.USDAFN
			}
		case "ALL":
			{
				i = i * currency.Quotes.USDALL
			}
		case "AMD":
			{
				i = i * currency.Quotes.USDAMD
			}
		case "ANG":
			{
				i = i * currency.Quotes.USDANG
			}
		case "AOA":
			{
				i = i * currency.Quotes.USDAOA
			}
		case "ARS":
			{
				i = i * currency.Quotes.USDARS
			}
		case "AUD":
			{
				i = i * currency.Quotes.USDAUD
			}
		case "AWG":
			{
				i = i * currency.Quotes.USDAWG
			}
		case "AZN":
			{
				i = i * currency.Quotes.USDAZN
			}

		case "BAM":
			{
				i = i * currency.Quotes.USDBAM
			}
		case "BBD":
			{
				i = i * currency.Quotes.USDBBD
			}
		case "BDT":
			{
				i = f / currency.Quotes.USDBDT
			}
		case "BGN":
			{
				i = i * currency.Quotes.USDBGN
			}

		case "BHD":
			{
				i = i * currency.Quotes.USDBHD
			}
		case "BIF":
			{
				i = i * currency.Quotes.USDBIF
			}
		case "BMD":
			{
				i = i * currency.Quotes.USDBMD
			}
		case "BND":
			{
				i = i * currency.Quotes.USDBND
			}
		case "BOB":
			{
				i = i * currency.Quotes.USDBOB
			}
		case "BRL":
			{
				i = i * currency.Quotes.USDBRL
			}
		case "BSD":
			{
				i = i * currency.Quotes.USDBSD
			}
		case "BTC":
			{
				i = i * currency.Quotes.USDBTC
			}
		case "BTN":
			{
				i = i * currency.Quotes.USDBTN
			}
		case "BWP":
			{
				i = i * currency.Quotes.USDBWP
			}
		case "BYN":
			{
				i = i * currency.Quotes.USDBYN
			}
		case "BYR":
			{
				i = i * currency.Quotes.USDBYR
			}
		case "BZD":
			{
				i = i * currency.Quotes.USDBZD
			}

		case "CAD":
			{
				i = i * currency.Quotes.USDCAD
			}
		case "CDF":
			{
				i = i * currency.Quotes.USDCHF
			}
		case "CHF":
			{

				i = i * currency.Quotes.USDXAU
			}
		case "FJD":
			{
				i = i * currency.Quotes.USDFJD
			}
		case "FKP":
			{
				i = i * currency.Quotes.USDFKP
			}
		case "GBP":
			{
				i = i * currency.Quotes.USDGBP
			}
		case "GEL":
			{
				i = i * currency.Quotes.USDGEL
			}
		case "GGP":
			{
				i = i * currency.Quotes.USDGGP
			}
		case "GHS":
			{
				i = i * currency.Quotes.USDGHS
			}
		case "GIP":
			{
				i = i * currency.Quotes.USDGIP
			}
		case "GMD":
			{
				i = i * currency.Quotes.USDGMD
			}
		case "GNF":
			{
				i = i * currency.Quotes.USDGNF
			}
		case "GTQ":
			{
				i = i * currency.Quotes.USDGTQ
			}
		case "GYD":
			{
				i = i * currency.Quotes.USDGYD
			}
		case "HKD":
			{
				i = i * currency.Quotes.USDHKD
			}
		case "HNL":
			{
				i = i * currency.Quotes.USDHNL
			}
		case "HRK":
			{
				i = i * currency.Quotes.USDHRK
			}
		case "HTG":
			{
				i = i * currency.Quotes.USDHTG
			}
		case "HUF":
			{
				i = i * currency.Quotes.USDHUF
			}
		case "IDR":
			{
				i = i * currency.Quotes.USDIDR
			}
		case "ILS":
			{
				i = i * currency.Quotes.USDILS
			}
		case "IMP":
			{
				i = i * currency.Quotes.USDIMP
			}
		case "INR":
			{
				i = i * currency.Quotes.USDINR
			}
		case "IQD":
			{
				i = i * currency.Quotes.USDIQD
			}
		case "IRR":
			{
				i = i * currency.Quotes.USDIRR
			}
		case "ISK":
			{
				i = i * currency.Quotes.USDISK
			}
		case "JEP":
			{
				i = i * currency.Quotes.USDJEP
			}
		case "JMD":
			{
				i = i * currency.Quotes.USDJMD
			}
		case "JOD":
			{
				i = i * currency.Quotes.USDJOD
			}
		case "JPY":
			{
				i = i * currency.Quotes.USDJPY
			}
		case "KES":
			{
				i = i * currency.Quotes.USDKES
			}
		case "KGS":
			{
				i = i * currency.Quotes.USDKGS
			}
		case "KHR":
			{

				i = i * currency.Quotes.USDKHR
			}
		case "KMF":
			{
				i = i * currency.Quotes.USDKMF
			}
		case "KPW":
			{
				i = i * currency.Quotes.USDKPW
			}
		case "KRW":
			{
				i = i * currency.Quotes.USDKRW
			}
		case "KWD":
			{
				i = i * currency.Quotes.USDKWD
			}
		case "KYD":
			{
				i = i * currency.Quotes.USDKYD
			}

		case "KZT":
			{
				i = i * currency.Quotes.USDKZT
			}
		case "LAK":
			{
				i = i * currency.Quotes.USDLAK
			}
		case "LBP":
			{
				i = i * currency.Quotes.USDLBP
			}
		case "LKR":
			{
				i = i * currency.Quotes.USDLKR
			}
		case "LRD":
			{
				i = i * currency.Quotes.USDLRD
			}
		case "LSL":
			{
				i = i * currency.Quotes.USDLSL
			}
		case "LTL":
			{
				i = i * currency.Quotes.USDLTL
			}
		case "LVL":
			{
				i = i * currency.Quotes.USDLVL
			}
		case "LYD":
			{
				i = i * currency.Quotes.USDLYD
			}
		case "MAD":
			{
				i = i * currency.Quotes.USDMAD
			}
		case "MDL":
			{
				i = i * currency.Quotes.USDMDL
			}
		case "MGA":
			{
				i = i * currency.Quotes.USDMGA
			}
		case "MKD":
			{
				i = i * currency.Quotes.USDMKD
			}
		case "MMK":
			{
				i = i * currency.Quotes.USDMMK
			}
		case "MNT":
			{
				i = i * currency.Quotes.USDMNT
			}
		case "MOP":
			{
				i = i * currency.Quotes.USDMOP
			}
		case "MRO":
			{
				i = i * currency.Quotes.USDMRO
			}
		case "MUR":
			{
				i = i * currency.Quotes.USDMUR
			}
		case "MVR":
			{
				i = i * currency.Quotes.USDMVR
			}
		case "MWK":
			{
				i = i * currency.Quotes.USDMWK
			}
		case "MXN":
			{
				i = i * currency.Quotes.USDMXN
			}
		case "MYR":
			{
				i = i * currency.Quotes.USDMYR
			}
		case "MZN":
			{
				i = i * currency.Quotes.USDMZN
			}
		case "NAD":
			{
				i = i * currency.Quotes.USDNAD
			}
		case "NGN":
			{
				i = i * currency.Quotes.USDNGN
			}
		case "NIO":
			{
				i = i * currency.Quotes.USDNIO
			}
		case "NOK":
			{
				i = i * currency.Quotes.USDNOK
			}
		case "NPR":
			{
				i = i * currency.Quotes.USDNPR
			}
		case "NZD":
			{
				i = i * currency.Quotes.USDNZD
			}
		case "OMR":
			{
				i = i * currency.Quotes.USDOMR
			}

		// P

		case "PAB":
			{
				i = i * currency.Quotes.USDPAB
			}
		// ====================== //

		case "PEN":
			{
				i = i * currency.Quotes.USDPEN
			}
		// ====================== //
		case "PGK":
			{
				i = i * currency.Quotes.USDPGK
			}
		// ====================== //
		case "PHP":
			{
				i = i * currency.Quotes.USDPHP
			}
		// ====================== //
		case "PKR":
			{
				i = i * currency.Quotes.USDPKR
			}
		// ====================== //
		case "PLN":
			{
				i = i * currency.Quotes.USDPLN
			}
		// ====================== //
		case "PYG":
			{
				i = i * currency.Quotes.USDPYG
			}

		// ====================== //

		// ==============================================//

		//Q

		case "QAR":
			{
				i = i * currency.Quotes.USDQAR
			}
		// ====================== //

		// ==============================================//

		//R
		case "RON":
			{
				i = i * currency.Quotes.USDRON
			}
		// ====================== //
		case "RSD":
			{
				i = i * currency.Quotes.USDRSD
			}
		// ====================== //
		case "RUB":
			{
				i = i * currency.Quotes.USDRUB
			}
		// ====================== //
		case "RWF":
			{
				i = i * currency.Quotes.USDRWF
			}

		// ==============================================//

		//S
		case "SAR":
			{
				i = i * currency.Quotes.USDSAR
			}
		// ====================== //
		case "SBD":
			{
				i = i * currency.Quotes.USDSBD
			}
		// ====================== //
		case "SCR":
			{
				i = i * currency.Quotes.USDSCR
			}
		// ====================== //
		case "SDG":
			{
				i = i * currency.Quotes.USDSDG
			}
		// ====================== //
		case "SEK":
			{
				i = i * currency.Quotes.USDSEK
			}
		// ====================== //
		case "SHP":
			{
				i = i * currency.Quotes.USDSHP
			}
		// ====================== //
		case "SLL":
			{
				i = i * currency.Quotes.USDSLL
			}
		// ====================== //
		case "SOS":
			{
				i = i * currency.Quotes.USDSOS
			}
		// ====================== //
		case "SRD":
			{
				i = i * currency.Quotes.USDSRD
			}
		// ====================== //
		case "STD":
			{
				i = i * currency.Quotes.USDSTD
			}
		// ====================== //
		case "SVC":
			{
				i = i * currency.Quotes.USDSVC
			}
		// ====================== //
		case "SYP":
			{
				i = i * currency.Quotes.USDSYP
			}
		// ====================== //
		case "SZL":
			{
				i = i * currency.Quotes.USDSZL
			}
		// ====================== //
		// ==============================================//

		//T

		case "THB":
			{
				i = i * currency.Quotes.USDTHB
			}
		// ====================== //
		case "TJS":
			{
				i = i * currency.Quotes.USDTJS
			}
		// ====================== //
		case "TMT":
			{
				i = i * currency.Quotes.USDTMT
			}
		// ====================== //
		case "TND":
			{
				i = i * currency.Quotes.USDTND
			}
		// ====================== //
		case "TOP":
			{
				i = i * currency.Quotes.USDTOP
			}
		// ====================== //
		case "TRY":
			{
				i = i * currency.Quotes.USDTRY
			}
		// ====================== //
		case "TTD":
			{
				i = i * currency.Quotes.USDTTD
			}
		// ====================== //
		case "TWD":
			{
				i = i * currency.Quotes.USDTWD
			}
		// ====================== //
		case "TZS":
			{
				i = i * currency.Quotes.USDTZS
			}
		// ====================== //
		// ==============================================//

		//U

		case "UAH":
			{
				i = i * currency.Quotes.USDUAH
			}
		// ====================== //
		case "UGX":
			{
				i = i * currency.Quotes.USDUGX
			}
		// ====================== //
		case "USD":
			{
				i = i * currency.Quotes.USDUSD
			}
		// ====================== //
		case "UYU":
			{
				i = i * currency.Quotes.USDUYU
			}
		// ====================== //
		case "UZS":
			{
				i = i * currency.Quotes.USDUZS
			}
			// ====================== //

		default:
			{

				return "", fmt.Errorf(to + " Currency Not Found Please Re enter")
			}
		}

		returnMessage := FloatToString(f) + " " + from + " is equivalent to " + FloatToString(i) + " " + to
		return returnMessage, nil

		// } else {
		// 	writeJSON(w, JSON{
		// 		"message": FloatToString(f) + " " + from + " is equivalent to " + FloatToString(i) + " " + to,
		// 	})
		// }

		// fmt.Println(currency)

		//Respond with the array of items

	}
}

// withLog Wraps HandlerFuncs to log requests to Stdout
func withLog(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := httptest.NewRecorder()
		fn(c, r)
		log.Printf("[%d] %-4s %s\n", c.Code, r.Method, r.URL.Path)

		for k, v := range c.HeaderMap {
			w.Header()[k] = v
		}
		w.WriteHeader(c.Code)
		c.Body.WriteTo(w)
	}
}

// writeJSON Writes the JSON equivilant for data into ResponseWriter w
func writeJSON(w http.ResponseWriter, data JSON) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// ProcessFunc Sets the processor of the chatbot
func ProcessFunc(p Processor) {
	processor = p
}

// handleWelcome Handles /welcome and responds with a welcome message and a generated UUID
func handleWelcome(w http.ResponseWriter, r *http.Request) {
	// Generate a UUID.
	hasher := md5.New()
	hasher.Write([]byte(strconv.FormatInt(time.Now().Unix(), 10)))
	uuid := hex.EncodeToString(hasher.Sum(nil))

	// Create a session for this UUID
	sessions[uuid] = Session{}

	// Write a JSON containg the welcome message and the generated UUID
	writeJSON(w, JSON{
		"uuid":    uuid,
		"message": WelcomeMessage,
	})
}

func handleChat(w http.ResponseWriter, r *http.Request) {
	// Make sure only POST requests are handled
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed.", http.StatusMethodNotAllowed)
		return
	}

	// Make sure a UUID exists in the Authorization header
	uuid := r.Header.Get("Authorization")
	if uuid == "" {
		http.Error(w, "Missing or empty Authorization header.", http.StatusUnauthorized)
		return
	}

	// Make sure a session exists for the extracted UUID
	session, sessionFound := sessions[uuid]
	if !sessionFound {
		http.Error(w, fmt.Sprintf("No session found for: %v.", uuid), http.StatusUnauthorized)
		return
	}

	// Parse the JSON string in the body of the request
	data := JSON{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, fmt.Sprintf("Couldn't decode JSON: %v.", err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Make sure a message key is defined in the body of the request
	_, messageFound := data["message"]
	if !messageFound {
		http.Error(w, "Missing message key in body.", http.StatusBadRequest)
		return
	}

	// Process the received message
	message, err := processor(session, data["message"].(string))

	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}

	//Write a JSON containg the processed response
	writeJSON(w, JSON{
		"message": message,
	})

	// Process the received message
	// processor(session, data["message"].(string), w)

	// Write a JSON containg the processed response

}

// handle Handles /
func handle(w http.ResponseWriter, r *http.Request) {
	body :=
		"<!DOCTYPE html><html><head><title>Chatbot</title></head><body><pre style=\"font-family: monospace;\">\n" +
			"Available Routes:\n\n" +
			"  GET  /welcome -> handleWelcome\n" +
			"  POST /chat    -> handleChat\n" +
			"  GET  /        -> handle        (current)\n" +
			"</pre></body></html>"
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintln(w, body)
}

// Engage Gives control to the chatbot
func Engage(addr string) error {
	// HandleFuncs
	mux := http.NewServeMux()
	mux.HandleFunc("/welcome", withLog(handleWelcome))
	mux.HandleFunc("/chat", withLog(handleChat))
	mux.HandleFunc("/", withLog(handle))

	// Start the server
	return http.ListenAndServe(addr, cors.CORS(mux))
}

func main() {
	// Uncomment the following lines to customize the chatbot
	WelcomeMessage = "Currency conversion?"
	ProcessFunc(processor)

	// Use the PORT environment variable
	port := os.Getenv("PORT")
	// Default to 3000 if no PORT environment variable was defined
	if port == "" {
		port = "3000"
	}

	// Start the server
	fmt.Printf("Listening on port %s...\n", port)
	log.Fatalln(Engage(":" + port))
}
