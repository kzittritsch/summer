package main

import (
	ldap "github.com/jtblin/go-ldap-client"
	"github.com/spf13/viper"
	"log"
)

func NewLDAP() *ldap.LDAPClient {
	client := &ldap.LDAPClient{
		Base:         viper.GetString("ldapBaseDN"),
		Host:         viper.GetString("ldapServer"),
		Port:         viper.GetInt("ldapPort"),
		UseSSL:       viper.GetBool("ldapSSL"),
		BindDN:       viper.GetString("ldapBindDN"),
		BindPassword: viper.GetString("ldapBindPasswd"),
		UserFilter:   viper.GetString("ldapUserFilter"),
		GroupFilter:  viper.GetString("ldapGroupFilter"),
		Attributes:   viper.GetStringSlice("ldapAttributes"),
	}

	return client
}

func Authenticate(username, passwd string) bool {
	conn := NewLDAP()
	defer conn.Close()

	success, user, err  := conn.Authenticate(username, passwd)
	log.Println(user)
	if success {
		log.Printf("Auth succeeded for user: %s", username)
		return true
	}
	if err != nil {
		log.Printf("Auth error: %s", err.Error())
	}
	if !success {
		log.Printf("Auth failed for user: %s", username)
	}

	return false
}
