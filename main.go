package main

import (
	"fmt"
	"log"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/go-ldap/ldap/v3"
)

var (
	ldapaddr         = kingpin.Flag("addr", "ldap addr").Default("127.0.0.1").String()
	ldapport         = kingpin.Flag("port", "ldap connect port").Default("389").Int()
	ldapusername     = kingpin.Flag("username", "ldap connect usernmae").Default("cn=admin,dc=rongfengliang,dc=com").String()
	ldapuserpassword = kingpin.Flag("password", "ldap connect password").Default("12sROjpn*^").String()
	debug            = kingpin.Flag("debug", "run with debug").Default("false").Bool()
)

func main() {
	kingpin.Parse()
	fmt.Printf("%v, %d\n", *ldapaddr, *ldapport)
	con, err := ldap.DialURL(fmt.Sprintf("ldap://%s:%d", *ldapaddr, *ldapport))
	if err != nil {
		log.Fatal("connect err:", err)
	}
	defer con.Close()
	con.Debug.Enable(*debug)
	err = con.Bind(*ldapusername, *ldapuserpassword)
	if err != nil {
		log.Fatal("bind err:", err)
	}
	searchRequest := ldap.NewSearchRequest(
		"dc=rongfengliang,dc=com",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(objectClass=*)",
		[]string{"dn", "cn", "objectClass"},
		nil,
	)
	searchResult, err := con.Search(searchRequest)
	if err != nil {
		log.Println("can't search ", err.Error())
	}
	log.Printf("%d", len(searchResult.Entries))
	for _, item := range searchResult.Entries {
		item.PrettyPrint(4)
	}
}
