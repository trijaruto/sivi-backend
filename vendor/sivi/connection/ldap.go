package connection

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
	"gopkg.in/ldap.v3"
)

func LoginLDAP(username, password string) (code int, errVal error) {
	l, err := ldap.Dial(viper.GetString("server.ldap.network"), fmt.Sprintf("%s:%d", viper.GetString("server.ldap.host"), viper.GetInt("server.ldap.port")))
	if err != nil {
		return http.StatusInternalServerError, nil
	}
	defer l.Close()
	err = l.Bind(fmt.Sprintf("%s\\%s", viper.GetString("server.ldap.domain"), viper.GetString("server.ldap.bindusername")), viper.GetString("server.ldap.bindpassword"))
	if err != nil {
		return http.StatusInternalServerError, nil
	}

	searchRequest := ldap.NewSearchRequest(
		fmt.Sprintf("dc=%s,dc=co,dc=id", viper.GetString("server.ldap.domain")),
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=user)(objectCategory=person)(sAMAccountName=%s))", username),
		[]string{"dn", "memberOf"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		return http.StatusInternalServerError, nil
	}

	if len(sr.Entries) < 1 {
		return http.StatusInternalServerError, nil
	}

	if len(sr.Entries) > 1 {
		return http.StatusInternalServerError, nil
	}

	ldapDN := sr.Entries[0].DN
	if err := l.Bind(ldapDN, password); err != nil {
		return http.StatusUnauthorized, nil
	} else {
		return http.StatusOK, nil
	}

}
