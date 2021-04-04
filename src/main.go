package main

import (
	"log"
	"net/http"

	"github.com/a-cordier/httpt/api"
	"github.com/a-cordier/httpt/util"
)

const banner = `

@@@  @@@ @@@@@@@ @@@@@@@ @@@@@@@  @@@@@@@
@@!  @@@   @@!     @@!   @@!  @@@   @@!  
@!@!@!@!   @!!     @!!   @!@@!@!    @!!  
!!:  !!!   !!:     !!:   !!:        !!:  
 :   : :    :       :     :          :  

https://github.com/a-cordier/httpt
ghcr.io/a-cordier/httpt

`

func main() {
	log.Print(banner)
	server := http.NewServeMux()
	port := util.GetEnvOrDefault("PORT", "8080")

	api.RegisterStatusHandlers(server)
	api.RegisterHeadersHandler(server)
	api.RegisterJohnDoesHandler(server)
	api.RegisterWaitHandler(server)
	api.RegisterInfoHandler(server)
	api.RegisterProxyHandler(server)
	log.Printf("%s %s\n\n", "Starting httpt on port", port)
	log.Fatal(http.ListenAndServe(":"+port, server))
}
