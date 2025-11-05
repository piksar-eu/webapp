package matrix

import (
	"fmt"
	"log"
	"os"

	easyconnect_public "github.com/piksar-eu/webapp/apps/core/pkg/easyconnect/public"
	"github.com/piksar-eu/webapp/apps/core/pkg/events"
)

func (m *Module) SendNotificationOnLeadCreatedListener(e events.Event) {
	if e.Type != easyconnect_public.LeadCreatedEvent {
		log.Println("invalid event type")
		return
	}

	payload := e.Payload.(easyconnect_public.LeadCreatedPayload)

	m.matrix.sendTextMsg(os.Getenv("MATRIX_NOTIFICATION_ROOM_ID"), fmt.Sprintf("Pozyskano nowego leada [%s]", payload.LeadId))
}
