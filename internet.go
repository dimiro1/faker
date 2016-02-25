package faker

import (
	"strings"

	"github.com/dimiro1/faker/defaults"
)

type EmailOptions struct {
	Name string
}

type UserNameOptions struct {
	Name string
	Sep  string `default:"-"`
}

type PassWordOptions struct {
	Length int  `default:"8"`
	Alpha  bool `default:"false"`
}

type URLOptions struct {
	Protocol string `default:"http"`
	Domain   string `default:"example.com"`
	Path     string `default:"/"`
}

type SlugOptions struct {
	Value string
	Sep   string `default:"-"`
}

func (f Faker) Email() string {
	return "eliza@mann.net"
}

func (f Faker) EmailWithOptions(o EmailOptions) string {
	return "eliza@mann.net"
}

func (f Faker) FreeEmail() string {
	return "eliza@gmail.com"
}

func (f Faker) FreeEmailWithOptions(o EmailOptions) string {
	return "eliza@gmail.com"
}

func (f Faker) SafeEmail() string {
	return "eliza@gmail.com"
}

func (f Faker) SafeEmailWithOptions(o EmailOptions) string {
	return "eliza@gmail.com"
}

// UserName returns a username
func (f Faker) UserName() string {
	return strings.ToLower(f.FirstName())
}

// UserNameWithOptions returns a username with the default options
func (f Faker) UserNameWithOptions(o UserNameOptions) string {
	defaults.Apply(&o)

	var words []string

	for _, word := range strings.Split(o.Name, " ") {
		words = append(words, strings.ToLower(word))
	}

	return strings.Join(words, o.Sep)
}

// Password returns a 8 length password
func (f Faker) Password() string {
	return format("********")
}

// PasswordWithOptions returns password defined by the options
func (f Faker) PasswordWithOptions(o PassWordOptions) string {
	defaults.Apply(&o)

	var pass string

	if o.Alpha {
		pass = fillString("*", o.Length)
	} else {
		pass = fillString("#", o.Length)
	}

	return format(pass)
}

func (f Faker) DomainName() string {
	return ""
}

func (f Faker) DomainSuffix() string {
	return ""
}

func (f Faker) IPV4() string {
	return "127.0.0.1"
}

func (f Faker) PublicIPV4() string {
	return "203.134.76.1"
}

func (f Faker) IPV6() string {
	return ""
}

func (f Faker) MacAddress() string {
	return "00:77:89:34:a4"
}

func (f Faker) URL() string {
	return ""
}

func (f Faker) URLWithOptions(o URLOptions) string {
	return ""
}

func (f Faker) Slug() string {
	return "pariatur_laudantium"
}

func (f Faker) SlugWithOptions(o SlugOptions) string {
	return "hello-world"
}

func (f Faker) UserAgent() string {
	return "Mozilla/5.0 (iPod; U; CPU iPhone OS 3_2 like Mac OS X; en-US) AppleWebKit/531.15.3 (KHTML, like Gecko) Version/4.0.5 Mobile/8B119 Safari/6531.15.3"
}

func (f Faker) IEUserAgent() string {
	return ""
}

func (f Faker) FirefoxUserAgent() string {
	return ""
}

func (f Faker) ChromeUserAgent() string {
	return ""
}

func (f Faker) OperaUserAgent() string {
	return ""
}

func (f Faker) SafariUserAgent() string {
	return ""
}

func (f Faker) Browser() string {
	return ""
}
