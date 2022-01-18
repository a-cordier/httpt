package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-cordier/httpt/util"
)

type Status struct {
	code int
	text string
}

func (status *Status) write(w http.ResponseWriter) {
	w.WriteHeader(status.code)
	fmt.Fprintf(w, status.text + "\n")
}

func (status *Status) route() string {
	return fmt.Sprintf("/status/%d", status.code)
}

func (status *Status) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status.write(w)
	}
}

func newStatus(code int) *Status {
	return &Status{
		code,
		http.StatusText(code),
	}
}

var statusList = []*Status{
	newStatus(http.StatusContinue),
	newStatus(http.StatusSwitchingProtocols),
	newStatus(http.StatusProcessing),
	newStatus(http.StatusEarlyHints),
	newStatus(http.StatusOK),
	newStatus(http.StatusCreated),
	newStatus(http.StatusAccepted),
	newStatus(http.StatusNonAuthoritativeInfo),
	newStatus(http.StatusNoContent),
	newStatus(http.StatusResetContent),
	newStatus(http.StatusPartialContent),
	newStatus(http.StatusMultiStatus),
	newStatus(http.StatusAlreadyReported),
	newStatus(http.StatusIMUsed),
	newStatus(http.StatusMultipleChoices),
	newStatus(http.StatusFound),
	newStatus(http.StatusSeeOther),
	newStatus(http.StatusNotModified),
	newStatus(http.StatusUseProxy),
	newStatus(http.StatusTemporaryRedirect),
	newStatus(http.StatusPermanentRedirect),
	newStatus(http.StatusBadRequest),
	newStatus(http.StatusUnauthorized),
	newStatus(http.StatusPaymentRequired),
	newStatus(http.StatusForbidden),
	newStatus(http.StatusNotFound),
	newStatus(http.StatusMethodNotAllowed),
	newStatus(http.StatusNotAcceptable),
	newStatus(http.StatusProxyAuthRequired),
	newStatus(http.StatusRequestTimeout),
	newStatus(http.StatusConflict),
	newStatus(http.StatusGone),
	newStatus(http.StatusLengthRequired),
	newStatus(http.StatusPreconditionFailed),
	newStatus(http.StatusRequestEntityTooLarge),
	newStatus(http.StatusRequestURITooLong),
	newStatus(http.StatusUnsupportedMediaType),
	newStatus(http.StatusRequestedRangeNotSatisfiable),
	newStatus(http.StatusExpectationFailed),
	newStatus(http.StatusTeapot),
	newStatus(http.StatusMisdirectedRequest),
	newStatus(http.StatusUnprocessableEntity),
	newStatus(http.StatusLocked),
	newStatus(http.StatusFailedDependency),
	newStatus(http.StatusTooEarly),
	newStatus(http.StatusUpgradeRequired),
	newStatus(http.StatusPreconditionRequired),
	newStatus(http.StatusTooManyRequests),
	newStatus(http.StatusRequestHeaderFieldsTooLarge),
	newStatus(http.StatusUnavailableForLegalReasons),
	newStatus(http.StatusInternalServerError),
	newStatus(http.StatusNotImplemented),
	newStatus(http.StatusBadGateway),
	newStatus(http.StatusServiceUnavailable),
	newStatus(http.StatusGatewayTimeout),
	newStatus(http.StatusHTTPVersionNotSupported),
	newStatus(http.StatusVariantAlsoNegotiates),
	newStatus(http.StatusInsufficientStorage),
	newStatus(http.StatusLoopDetected),
	newStatus(http.StatusNotExtended),
	newStatus(http.StatusNetworkAuthenticationRequired),
}

func RegisterStatusHandlers(server *http.ServeMux) {
	for _, status := range statusList {
		log.Printf("%s %s (%s %d: %s)", "Registering [*]", status.route(), "Returns HTTP status", status.code, status.text)
		server.HandleFunc(status.route(), util.LogTime(status.handler()))
	}
}
