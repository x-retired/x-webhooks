package install

import (
	"flag"
	"fmt"
	"github.com/xiexianbin/webhooks/models"
	"os"
	"strings"
)

// Install webhooks, Use method:
// webhooks install -username=admin -password=123456 -email=me@xiexianbin.cn
func Install() {
	if os.Args[1] == "install" && len(os.Args) > 2 {
		username := flag.String("username", "admin", "User Account.")
		password_raw := flag.String("password", "", "User Password.")
		password := strings.TrimSpace(*password_raw)
		email := flag.String("email", "", "User Email.")

		flag.CommandLine.Parse(os.Args[2:])

		if password == "" {
			fmt.Println("Password is Required.")
			os.Exit(2)
		}
		if *email == "" {
			fmt.Println("User email is required.")
			os.Exit(2)
		}

		user := models.NewUser()
		user.UserName = *username
		user.Password = password
		user.Email = *email

		if err := user.Create(); err != nil {
			fmt.Println(err.Error())
			os.Exit(2)
		}
		fmt.Println("Done.")
		os.Exit(2)
	}
}
