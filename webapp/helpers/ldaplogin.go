package helpers

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/eaciit/clit"
	"github.com/eaciit/ldap"
	"github.com/eaciit/toolkit"
)

type LDAPDataList struct {
	LDAPDataFullName    string
	LDAPDataFirstName   string
	LDAPDataLastName    string
	LDAPDataUserCountry string
}

func getLDAPConnection(address string, config toolkit.M) *ldap.Connection {
	l := new(ldap.Connection)

	switch config.GetString("type") {
	case "ssl":
		l = ldap.NewSSLConnection(address, config.Get("tlsconfig", nil).(*tls.Config))
	case "tls":
		l = ldap.NewTLSConnection(address, config.Get("tlsconfig", nil).(*tls.Config))
	default:
		l = ldap.NewConnection(address)
	}

	return l
}

func checkloginldap(username string, password string, loginconf toolkit.M, BindUsernameLDAP string, BindPasswordLDAP string, UserAuthAttrLDAP string, UserDNLDAP string, BindFilterLDAP string, LDAPData LDAPDataList) (cond bool, information LDAPDataList, err error) {
	toolkit.Println("# Connecting to LDAP")

	cond = false
	connectTime := time.Now()
	address := loginconf.GetString("address")

	l := getLDAPConnection(address, loginconf)
	defer l.Close()

	err = l.Connect()
	if err != nil {
		toolkit.Println("#ERROR Connecting to LDAP", err.Error())
		return
	}

	if strings.Trim(BindUsernameLDAP, " ") != "" {
		// Bind Through Config
		// defer l.Unbind(BindUsernameLDAP, BindPasswordLDAP)
		err = l.Bind(BindUsernameLDAP, BindPasswordLDAP)
	} else {
		// defer l.Unbind(username, password)
		err = l.Bind(username, password)
	}

	if strings.Trim(BindUsernameLDAP, " ") != "" && err == nil {
		// from Login FORM
		// defer l.Unbind(username, password)
		err = l.Bind(username, password)
		toolkit.Println("aaaaaaaaaaaa")
		if err == nil {
			cond = true
		} else {
			toolkit.Println("#ERROR Binding to LDAP with username : ", username, " - ", err.Error())
		}
	} else {
		toolkit.Println("bbbbbbbbb")
		if err == nil {
			cond = true
		} else {
			toolkit.Println("#ERROR Binding to LDAP  with username : ", username, " - ", err.Error())
		}
	}

	toolkit.Println("# Closing LDAP Connection")
	toolkit.Println("# Connection Time : ", time.Since(connectTime).Seconds(), "s")

	return

	if cond {
		toolkit.Println("# Getting some Information from LDAP ")
		filter := ""

		if strings.Trim(UserAuthAttrLDAP, " ") == "" {
			filter = "(CN=" + username + ")"
		} else {
			filter = "(dn=" + username + ")"
		}

		attributes := []string{}
		if LDAPData.LDAPDataFullName != "" {
			attributes = append(attributes, LDAPData.LDAPDataFullName)
		}

		if LDAPData.LDAPDataFirstName != "" {
			attributes = append(attributes, LDAPData.LDAPDataFirstName)
		}

		if LDAPData.LDAPDataLastName != "" {
			attributes = append(attributes, LDAPData.LDAPDataLastName)
		}
		// if LDAPData.LDAPDataLastName != "" {
		// 	attributes = append(attributes, LDAPData.LDAPDataLastName)
		// }

		// attributes := []string{"cn", "givenName"}
		toolkit.Println("BaseDN : ", loginconf.GetString("basedn"))
		toolkit.Println("Filter : ", filter)
		toolkit.Println("Attributes : ", strings.Join(attributes, ", "))
		search := ldap.NewSearchRequest(loginconf.GetString("basedn"),
			ldap.ScopeWholeSubtree,
			ldap.DerefAlways,
			0,
			0,
			false,
			filter,
			attributes,
			nil)

		sr, err := l.Search(search)
		if err != nil {
			toolkit.Println("#ERROR Getting information from LDAP : ", err.Error())
		}

		for _, v := range sr.Entries {
			for _, str := range attributes {
				val := ""
				if len(v.GetAttributeValues(str)) > 1 {
					val = strings.Join(v.GetAttributeValues(str), "|")
				} else {
					val = v.GetAttributeValue(str)
				}

				if val != "" {
					switch str {
					case LDAPData.LDAPDataFullName:
						information.LDAPDataFullName = val
						break
					case LDAPData.LDAPDataFirstName:
						information.LDAPDataFirstName = val
						break
					case LDAPData.LDAPDataLastName:
						information.LDAPDataLastName = val
						break
					// case LDAPData.LDAPDataLastName:
					// 	information.LDAPDataLastName = val
					// 	break
					default:
						break
					}
				}
			}
		}

		toolkit.Println("# Information Achieved : ")
		toolkit.Println("# - LDAPDataFullName : ", information.LDAPDataFullName)
		toolkit.Println("# - LDAPDataFirstName : ", information.LDAPDataFirstName)
		toolkit.Println("# - LDAPDataLastName : ", information.LDAPDataLastName)
		toolkit.Println("# - LDAPDataLastName : ", information.LDAPDataLastName)
	}

	return

}

func TryToLoginUsingLDAP(username, password string) (bool, LDAPDataList, error) {
	type ForgetMe struct{}

	ldapConf := clit.Config("default", "LDAP", "").(map[string]interface{})

	configLDAPHost := strings.TrimSpace(ldapConf["Host"].(string))
	configLDAPServerName := strings.TrimSpace(ldapConf["ServerName"].(string))
	configLDAPType := strings.TrimSpace(ldapConf["Type"].(string))
	configLDAPUserAuthAttr := strings.TrimSpace(ldapConf["UserAuthAttr"].(string))
	configLDAPUserDN := strings.TrimSpace(ldapConf["UserDN"].(string))
	configLDAPBaseDN := strings.TrimSpace(ldapConf["BaseDN"].(string))
	configLDAPBindUsername := strings.TrimSpace(ldapConf["BindUsername"].(string))
	configLDAPBindPassword := strings.TrimSpace(ldapConf["BindPassword"].(string))
	configLDAPBindFilter := strings.TrimSpace(ldapConf["BindFilter"].(string))
	configLDAPIsInsecureSkipVerify := strings.ToLower(strings.TrimSpace(ldapConf["IsInsecureSkipVerify"].(string))) == "true"
	configLDAPCertificate := ldapConf["Certificate"].(string)

	// if configLDAPBindPassword != "" {
	// 	decoded, _ := Decode(configLDAPBindPassword)
	// 	configLDAPBindPassword = decoded
	// }

	ldapData := LDAPDataList{}
	ldapData.LDAPDataFullName = ldapConf["DataFullName"].(string)
	ldapData.LDAPDataFirstName = ldapConf["DataFirstName"].(string)
	ldapData.LDAPDataLastName = ldapConf["DataLastName"].(string)
	ldapData.LDAPDataUserCountry = ldapConf["DataUserCountry"].(string)

	ldapConfig := toolkit.M{}
	ldapConfig.Set("host", configLDAPHost)
	ldapConfig.Set("basedn", configLDAPBaseDN)

	loginConf := toolkit.M{}
	loginConf.Set("address", configLDAPHost)
	loginConf.Set("basedn", configLDAPBaseDN)

	if configLDAPType != "" {
		toolkit.Println("#Type :", configLDAPType)
		toolkit.Println("#ServerName : ", configLDAPServerName)
		toolkit.Println("#Address : ", configLDAPHost)
		toolkit.Println("#BaseDN : ", configLDAPBaseDN)

		if configLDAPUserAuthAttr != "" {
			toolkit.Println("#UserAuthAttrLDAP : ", configLDAPUserAuthAttr)
		}

		if configLDAPUserDN != "" {
			toolkit.Println("#UserDNLDAP : ", configLDAPUserDN)
		}

		toolkit.Println("#Username : ", username)
		loginConf.Set("type", configLDAPType)

		tlsconfig := tls.Config{}
		if configLDAPServerName != "" {
			tlsconfig.ServerName = configLDAPServerName
		}

		tlsconfig.InsecureSkipVerify = configLDAPIsInsecureSkipVerify
		tlsconfig.Certificates = []tls.Certificate{}
		caCertPool := x509.NewCertPool()

		for _, c := range strings.Split(configLDAPCertificate, ",") {
			if c == "," {
				continue
			}

			file, err := ioutil.ReadFile(c)
			if err != nil {
				err = errors.New(fmt.Sprintf("Found error : %v", err.Error()))
			}

			_ = caCertPool.AppendCertsFromPEM(file)
			// toolkit.Println("Certificate ", (x + 1), " # Added : ", s)
		}

		tlsconfig.RootCAs = caCertPool
		loginConf.Set("tlsconfig", &tlsconfig)
	}

	rawUsername := username
	UserAuthAttr := "CN"
	if configLDAPUserAuthAttr != "" {
		UserAuthAttr = configLDAPUserAuthAttr
		username = UserAuthAttr + "=" + username
	}

	if configLDAPUserDN != "" {
		if configLDAPUserAuthAttr == "" {
			username = UserAuthAttr + "=" + username
		}
		username = username + "," + configLDAPUserDN
	}

	// ldapUserDNComponent := strings.Split(configLDAPUserDN, ",")
	// if len(ldapUserDNComponent) > 0 {
	// 	if ldapUserDNComponent[0] == rawUsername {
	// 		//username = configLDAPUserDN
	// 		username = UserAuthAttr + "=" + ldapUserDNComponent[0]
	// 	}
	// }

	ldapUserDNComponent := strings.Split(configLDAPBaseDN, ",")
	if len(ldapUserDNComponent) > 0 {
		if strings.Replace(ldapUserDNComponent[0], UserAuthAttr+"=", "", -1) == rawUsername {
			//username = configLDAPUserDN
			username = configLDAPBaseDN
		}
	}

	// if IsAdminLogin {
	// 	username = BaseDNLDAP
	// }

	isSuccess, information, err := checkloginldap(username, password, loginConf, configLDAPBindUsername, configLDAPBindPassword, configLDAPUserAuthAttr, configLDAPUserDN, configLDAPBindFilter, ldapData)
	return isSuccess, information, err
}

func FindDataLdap(addr, basedn, filter string, param toolkit.M) (arrtkm []toolkit.M, err error) {
	arrtkm = make([]toolkit.M, 0, 0)

	l := getLDAPConnection(addr, param)
	err = l.Connect()
	if err != nil {
		return
	}
	defer l.Close()

	if param.Has("username") {
		// defer l.Unbind(toolkit.ToString(param["username"]), toolkit.ToString(param["password"]))
		err = l.Bind(toolkit.ToString(param["username"]), toolkit.ToString(param["password"]))
		if err != nil {
			return
		}
	}

	attributes := make([]string, 0, 0)
	if param.Has("attributes") {
		attributes = param["attributes"].([]string)
	}
	// filter = "(*" + filter + "*)"
	search := ldap.NewSearchRequest(basedn,
		ldap.ScopeWholeSubtree,
		ldap.DerefAlways,
		0,
		0,
		false,
		filter,
		attributes,
		nil)

	sr, err := l.Search(search)

	for _, v := range sr.Entries {
		tkm := toolkit.M{}

		for _, str := range attributes {
			if len(v.GetAttributeValues(str)) > 1 {
				tkm.Set(str, v.GetAttributeValues(str))
			} else {
				tkm.Set(str, v.GetAttributeValue(str))
			}
		}

		if len(tkm) > 0 {
			arrtkm = append(arrtkm, tkm)
		}
	}

	return
}
