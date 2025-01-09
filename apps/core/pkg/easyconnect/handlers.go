package easyconnect

import "github.com/piksar-eu/webapp/apps/core/pkg/shared"

func SubscribeFn(leadRepo LeadRepository) func(email string) error {
	return func(email string) error {
		email, err := shared.SanitizeEmail(email)

		if err != nil {
			return err
		}

		lead, err := leadRepo.Get(email)

		if lead != nil && lead.MarketingConsent {
			return nil
		}

		if lead == nil {
			lead, err = createLead(email, "newsletter")
		}

		if err != nil {
			return err
		}

		lead.changeMarketingConsent(true)

		err = leadRepo.Save(lead)
		if err != nil {
			return err
		}

		return nil
	}
}
