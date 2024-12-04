package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/suffer-sami/realtor-scraper/internal/database"
)

func (cfg *config) executeTransaction(ctx context.Context, txFunc func(context.Context, *database.Queries) error) error {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	tx, err := cfg.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	qtx := cfg.dbQueries.WithTx(tx)

	err = txFunc(ctx, qtx)

	if err != nil {
		return fmt.Errorf("transaction failed: %v", err)
	}
	return tx.Commit()
}

func (cfg *config) storeAgents(agents []Agent) {
	defer cfg.wg.Done()
	for _, agent := range agents {
		cfg.wg.Add(1)
		go func() {
			if err := cfg.storeAgent(agent); err != nil {
				cfg.logger.Errorf("error storing agent (ID: %s): %v", agent.ID, err)
			}
		}()
	}
}

func (cfg *config) storeAgent(agent Agent) error {
	defer cfg.wg.Done()
	return cfg.executeTransaction(context.Background(), func(ctx context.Context, qtx *database.Queries) error {
		cfg.logger.Infof("Agent: %s", agent.PersonName)

		dbAgent, err := qtx.GetAgent(ctx, agent.ID)
		if err != nil {
			if err != sql.ErrNoRows {
				return err
			}

			dbAgent, err = qtx.CreateAgent(ctx, database.CreateAgentParams{
				ID:                   agent.ID,
				FirstName:            stringToNullString(agent.FirstName),
				LastName:             stringToNullString(agent.LastName),
				NickName:             stringToNullString(agent.NickName),
				PersonName:           stringToNullString(agent.PersonName),
				Title:                stringToNullString(agent.Title),
				Slogan:               stringToNullString(agent.Slogan),
				Email:                stringToNullString(agent.Email),
				AgentRating:          intToNullInt64(agent.AgentRating),
				Description:          stringToNullString(agent.Description),
				RecommendationsCount: intToNullInt64(agent.RecommendationsCount),
				ReviewCount:          intToNullInt64(agent.ReviewCount),
				LastUpdated:          strToNullTime(agent.LastUpdated, time.RFC1123),
				FirstMonth:           numericToNullInt(agent.FirstMonth),
				FirstYear:            intToNullInt64(agent.AgentRating),
				Photo:                stringToNullString(agent.Photo.Href),
				Video:                stringToNullString(agent.Video),
				ProfileUrl:           stringToNullString(agent.WebURL),
				Website:              stringToNullString(agent.Href),
			})

			if err != nil {
				return err
			}

			if cfg.saveRawAgents {
				cfg.logger.Debugf("- raw agent:")
				if jsonStrAgent, err := anyToJsonString(agent); err == nil {
					cfg.logger.Debugf("	* %s", jsonStrAgent)
					if err := qtx.CreateRawAgent(ctx, database.CreateRawAgentParams{
						AgentID: stringToNullString(dbAgent.ID),
						Data:    stringToNullString(jsonStrAgent),
					}); err != nil {
						cfg.logger.Errorf("error creating raw agent: %v", err)
					}
				}
			}
		}
		agentId := stringToNullString(dbAgent.ID)

		cfg.logger.Debugf("- sales data: %s", agent.RecentlySold.LastSoldDate)
		if err := qtx.CreateSalesData(ctx, database.CreateSalesDataParams{
			Count:        intToNullInt64(agent.RecentlySold.Count),
			Min:          intToNullInt64(agent.RecentlySold.Min),
			Max:          intToNullInt64(agent.RecentlySold.Max),
			LastSoldDate: strToNullTime(agent.RecentlySold.LastSoldDate, time.DateOnly),
			AgentID:      agentId,
		}); err != nil {
			cfg.logger.Errorf("error creating sales data: %v", err)
		}

		cfg.logger.Debugf("- listing data: %s", agent.ForSalePrice.LastListingDate)
		if err := qtx.CreateListingsData(ctx, database.CreateListingsDataParams{
			Count:           intToNullInt64(agent.ForSalePrice.Count),
			Min:             intToNullInt64(agent.ForSalePrice.Min),
			Max:             intToNullInt64(agent.ForSalePrice.Max),
			LastListingDate: timeToNullTime(agent.ForSalePrice.LastListingDate),
			AgentID:         agentId,
		}); err != nil {
			cfg.logger.Errorf("error creating listing data: %v", err)
		}

		cfg.logger.Debugf("- social medias:")
		for _, socialMedia := range agent.SocialMedias {
			cfg.logger.Debugf("	* %s", socialMedia.Type)
			if err := qtx.CreateSocialMedia(ctx, database.CreateSocialMediaParams{
				Type:    stringToNullString(socialMedia.Type),
				Href:    stringToNullString(socialMedia.Href),
				AgentID: agentId,
			}); err != nil {
				cfg.logger.Errorf("error creating social media: %v", err)
			}
		}

		cfg.logger.Debugf("- feed licences:")
		for _, feedLicense := range agent.FeedLicenses {
			cfg.logger.Debugf("	* (%s, %s)", feedLicense.StateCode, feedLicense.Country)
			if err := qtx.CreateFeedLicense(ctx, database.CreateFeedLicenseParams{
				Country:       stringToNullString(feedLicense.Country),
				LicenseNumber: stringToNullString(feedLicense.LicenseNumber),
				StateCode:     stringToNullString(feedLicense.StateCode),
				AgentID:       agentId,
			}); err != nil {
				cfg.logger.Errorf("error creating feed licence: %v", err)
			}
		}

		cfg.logger.Debugf("- mls:")
		for _, mls := range agent.Mls {
			cfg.logger.Debugf("	* %s", mls.Abbreviation)
			dbMls, err := qtx.GetMultipleListingService(ctx, database.GetMultipleListingServiceParams{
				Abbreviation:  stringToNullString(mls.Abbreviation),
				Type:          stringToNullString(mls.Type),
				MemberID:      stringToNullString(mls.MemberID),
				LicenseNumber: stringToNullString(mls.LicenseNumber),
			})

			if err != nil {
				if err != sql.ErrNoRows {
					return err
				}
				dbMls, err = qtx.CreateMultipleListingService(ctx, database.CreateMultipleListingServiceParams{
					Abbreviation:  stringToNullString(mls.Abbreviation),
					LicenseNumber: stringToNullString(mls.LicenseNumber),
					Type:          stringToNullString(mls.Type),
					MemberID:      stringToNullString(mls.MemberID),
					IsPrimary:     boolToNullBool(mls.Primary),
				})

				if err != nil {
					return err
				}
			}

			if err := qtx.CreateAgentMultipleListingService(ctx, database.CreateAgentMultipleListingServiceParams{
				AgentID:                  agentId,
				MultipleListingServiceID: int64ToNullInt64(dbMls.ID),
			}); err != nil {
				cfg.logger.Errorf("error creating agent mls: %v", err)
			}
		}
		cfg.logger.Debugf("- mls history:")
		for _, mls := range agent.MlsHistory {
			cfg.logger.Debugf("	* %s", mls.Abbreviation)
			dbMls, err := qtx.GetMultipleListingService(ctx, database.GetMultipleListingServiceParams{
				Abbreviation:  stringToNullString(mls.Abbreviation),
				Type:          stringToNullString(mls.Type),
				MemberID:      stringToNullString(mls.Member.ID),
				LicenseNumber: stringToNullString(mls.LicenseNumber),
			})

			if err != nil {
				if err != sql.ErrNoRows {
					return err
				}
				dbMls, err = qtx.CreateMultipleListingService(ctx, database.CreateMultipleListingServiceParams{
					Abbreviation:     stringToNullString(mls.Abbreviation),
					InactivationDate: timeToNullTime(mls.InactivationDate),
					LicenseNumber:    stringToNullString(mls.LicenseNumber),
					IsPrimary:        boolToNullBool(mls.Primary),
					Type:             stringToNullString(mls.Type),
					MemberID:         stringToNullString(mls.Member.ID),
				})

				if err != nil {
					return err
				}
			}

			if dbMls.InactivationDate.Time != mls.InactivationDate {
				cfg.logger.Debugf("	* %s (update inactivation_date: %v)", mls.Abbreviation, mls.InactivationDate)
				if err := qtx.UpdateMultipleListingServiceInactivationDate(ctx, database.UpdateMultipleListingServiceInactivationDateParams{
					InactivationDate: dbMls.InactivationDate,
					ID:               dbMls.ID,
				}); err != nil {
					cfg.logger.Errorf("error (update mls inactivation_date: %v)", err)
				}
			}

			if err = qtx.CreateAgentMultipleListingService(ctx, database.CreateAgentMultipleListingServiceParams{
				AgentID:                  agentId,
				MultipleListingServiceID: int64ToNullInt64(dbMls.ID),
			}); err != nil {
				cfg.logger.Errorf("error creating agent mls: %v", err)
			}
		}

		cfg.logger.Debugf("- languages:")
		for _, lang := range agent.Languages {
			cfg.logger.Debugf("	* %s", lang)
			dbLangID, err := qtx.GetLanguageID(ctx, stringToNullString(lang))
			if err != nil {
				if err != sql.ErrNoRows {
					return err
				}

				dbLangID, err = qtx.CreateLanguage(ctx, stringToNullString(lang))
				if err != nil {
					return err
				}
			}

			if err := qtx.CreateAgentLanguage(ctx, database.CreateAgentLanguageParams{
				LanguageID: int64ToNullInt64(dbLangID),
				AgentID:    agentId,
			}); err != nil {
				cfg.logger.Errorf("error creating agent language: %v", err)
			}
		}
		cfg.logger.Debugf("- user languages:")
		for _, lang := range agent.UserLanguages {
			cfg.logger.Debugf("	* %s", lang)
			dbLangID, err := qtx.GetLanguageID(ctx, stringToNullString(lang))
			if err != nil {
				if err != sql.ErrNoRows {
					return err
				}

				dbLangID, err = qtx.CreateLanguage(ctx, stringToNullString(lang))
				if err != nil {
					return err
				}
			}

			if err := qtx.CreateAgentUserLanguage(ctx, database.CreateAgentUserLanguageParams{
				LanguageID: int64ToNullInt64(dbLangID),
				AgentID:    agentId,
			}); err != nil {
				cfg.logger.Errorf("error creating agent language: %v", err)
			}
		}

		cfg.logger.Debugf("- zips:")
		for _, zip := range agent.Zips {
			cfg.logger.Debugf("	* %s", zip)
			dbZipID, err := qtx.GetZipID(ctx, stringToNullString(zip))
			if err != nil {
				if err != sql.ErrNoRows {
					return err
				}

				dbZipID, err = qtx.CreateZip(ctx, stringToNullString(zip))
				if err != nil {
					return err
				}
			}

			if err := qtx.CreateAgentZip(ctx, database.CreateAgentZipParams{
				ZipID:   int64ToNullInt64(dbZipID),
				AgentID: agentId,
			}); err != nil {
				cfg.logger.Errorf("error creating agent zip: %v", err)
			}
		}

		cfg.logger.Debugf("- areas:")
		for _, area := range agent.ServedAreas {
			cfg.logger.Debugf("	* (%s, %s)", area.Name, area.StateCode)
			dbAreaID, err := qtx.GetAreaID(ctx, database.GetAreaIDParams{
				Name:      stringToNullString(area.Name),
				StateCode: stringToNullString(area.StateCode),
			})

			if err != nil {
				if err != sql.ErrNoRows {
					return err
				}

				dbAreaID, err = qtx.CreateArea(ctx, database.CreateAreaParams{
					Name:      stringToNullString(area.Name),
					StateCode: stringToNullString(area.StateCode),
				})
				if err != nil {
					return err
				}
			}

			if err := qtx.CreateAgentServedArea(ctx, database.CreateAgentServedAreaParams{
				AreaID:  int64ToNullInt64(dbAreaID),
				AgentID: agentId,
			}); err != nil {
				cfg.logger.Errorf("error creating agent served area: %v", err)
			}
		}
		cfg.logger.Debugf("- marketing areas:")
		for _, area := range agent.MarketingAreaCities {
			cfg.logger.Debugf("	* (%s, %s)", area.Name, area.StateCode)
			dbAreaID, err := qtx.GetAreaID(ctx, database.GetAreaIDParams{
				Name:      stringToNullString(area.Name),
				StateCode: stringToNullString(area.StateCode),
			})

			if err != nil {
				if err != sql.ErrNoRows {
					return err
				}

				dbAreaID, err = qtx.CreateArea(ctx, database.CreateAreaParams{
					Name:      stringToNullString(area.Name),
					StateCode: stringToNullString(area.StateCode),
				})
				if err != nil {
					return err
				}
			}

			if err := qtx.CreateAgentMarketingArea(ctx, database.CreateAgentMarketingAreaParams{
				AreaID:  int64ToNullInt64(dbAreaID),
				AgentID: agentId,
			}); err != nil {
				cfg.logger.Errorf("error creating agent marketing area: %v", err)
			}
		}

		cfg.logger.Debugf("- designations:")
		for _, designation := range agent.Designations {
			cfg.logger.Debugf("	* %s", designation.Name)
			dbDesignationID, err := qtx.GetDesignationID(ctx, stringToNullString(designation.Name))
			if err != nil {
				if err != sql.ErrNoRows {
					return err
				}

				dbDesignationID, err = qtx.CreateDesignation(ctx, stringToNullString(designation.Name))
				if err != nil {
					return err
				}
			}

			if err := qtx.CreateAgentDesignation(ctx, database.CreateAgentDesignationParams{
				DesignationID: int64ToNullInt64(dbDesignationID),
				AgentID:       agentId,
			}); err != nil {
				cfg.logger.Errorf("error creating agent designation: %v", err)
			}
		}
		cfg.logger.Debugf("- specializations:")
		for _, specialization := range agent.Specializations {
			cfg.logger.Debugf("	* %s", specialization.Name)
			dbSpecializationID, err := qtx.GetSpecializationID(ctx, stringToNullString(specialization.Name))
			if err != nil {
				if err != sql.ErrNoRows {
					return err
				}

				dbSpecializationID, err = qtx.CreateSpecialization(ctx, stringToNullString(specialization.Name))
				if err != nil {
					return err
				}
			}

			if err := qtx.CreateAgentSpecialization(ctx, database.CreateAgentSpecializationParams{
				SpecializationID: int64ToNullInt64(dbSpecializationID),
				AgentID:          agentId,
			}); err != nil {
				cfg.logger.Errorf("error creating agent specialization: %v", err)
			}
		}

		cfg.logger.Debugf("- broker:")
		cfg.logger.Debugf("	* %s", agent.Broker.Name)
		dbBrokerID, err := qtx.GetBrokerID(ctx, intToNullInt64(agent.Broker.FulfillmentID))
		if err != nil {
			if err != sql.ErrNoRows {
				return err
			}

			dbBrokerID, err = qtx.CreateBroker(ctx, database.CreateBrokerParams{
				FulfillmentID: intToNullInt64(agent.Broker.FulfillmentID),
				Name:          stringToNullString(agent.Broker.Name),
				Photo:         stringToNullString(agent.Broker.Photo.Href),
				Video:         stringToNullString(agent.Broker.Video),
			})

			if err != nil {
				return err
			}
		}

		cfg.logger.Debugf("- address:")
		cfg.logger.Debugf("	* %+v", agent.Address)
		dbAddressID, err := qtx.GetAddressID(ctx, database.GetAddressIDParams{
			Line:       stringToNullString(agent.Address.Line),
			Line2:      stringToNullString(agent.Address.Line2),
			City:       stringToNullString(agent.Address.City),
			StateCode:  stringToNullString(agent.Address.StateCode),
			PostalCode: stringToNullString(agent.Address.PostalCode),
		})
		if err != nil {
			if err != sql.ErrNoRows {
				return err
			}

			dbAddressID, err = qtx.CreateAddress(ctx, database.CreateAddressParams{
				Line:       stringToNullString(agent.Address.Line),
				Line2:      stringToNullString(agent.Address.Line2),
				City:       stringToNullString(agent.Address.City),
				StateCode:  stringToNullString(agent.Address.StateCode),
				State:      stringToNullString(agent.Address.State),
				PostalCode: stringToNullString(agent.Address.PostalCode),
				Country:    stringToNullString(agent.Address.Country),
			})

			if err != nil {
				return err
			}
		}

		cfg.logger.Debugf("- office address:")
		cfg.logger.Debugf("	* %+v", agent.Office.Address)
		dbOfficeAddressID, err := qtx.GetAddressID(ctx, database.GetAddressIDParams{
			Line:       stringToNullString(agent.Address.Line),
			Line2:      stringToNullString(agent.Address.Line2),
			City:       stringToNullString(agent.Address.City),
			StateCode:  stringToNullString(agent.Address.StateCode),
			PostalCode: stringToNullString(agent.Address.PostalCode),
		})
		if err != nil {
			if err != sql.ErrNoRows {
				return err
			}

			dbOfficeAddressID, err = qtx.CreateAddress(ctx, database.CreateAddressParams{
				Line:       stringToNullString(agent.Address.Line),
				Line2:      stringToNullString(agent.Address.Line2),
				City:       stringToNullString(agent.Address.City),
				StateCode:  stringToNullString(agent.Address.StateCode),
				State:      stringToNullString(agent.Address.State),
				PostalCode: stringToNullString(agent.Address.PostalCode),
				Country:    stringToNullString(agent.Address.Country),
			})

			if err != nil {
				return err
			}
		}

		cfg.logger.Debugf("- office:")
		cfg.logger.Debugf("	* %v", agent.Office.Name)
		dbOfficeID, err := qtx.GetOfficeID(ctx, intToNullInt64(agent.Office.FulfillmentID))
		if err != nil {
			if err != sql.ErrNoRows {
				return err
			}

			dbOfficeID, err = qtx.CreateOffice(ctx, database.CreateOfficeParams{
				Name:          stringToNullString(agent.Office.Name),
				Email:         stringToNullString(agent.Office.Email),
				Photo:         stringToNullString(agent.Office.Photo.Href),
				Website:       stringToNullString(agent.Office.Website),
				Slogan:        stringToNullString(agent.Office.Slogan),
				Video:         stringToNullString(agent.Office.Video),
				FulfillmentID: intToNullInt64(agent.Office.FulfillmentID),
				AddressID:     int64ToNullInt64(dbOfficeAddressID),
			})

			if err != nil {
				return err
			}
		}

		cfg.logger.Debugf("- phones:")
		for _, phone := range agent.Phones {
			cfg.logger.Debugf("	* %s", phone.Number)

			dbPhoneID, err := qtx.GetPhoneID(ctx, database.GetPhoneIDParams{
				Ext:    stringToNullString(phone.Ext),
				Number: stringToNullString(phone.Number),
				Type:   stringToNullString(phone.Type),
			})

			if err != nil {
				if err != sql.ErrNoRows {
					return err
				}

				dbPhoneID, err = qtx.CreatePhone(ctx, database.CreatePhoneParams{
					Ext:     stringToNullString(phone.Ext),
					Number:  stringToNullString(phone.Number),
					Type:    stringToNullString(phone.Type),
					IsValid: boolToNullBool(phone.IsValid),
				})

				if err != nil {
					return err
				}
			}

			if err := qtx.CreateAgentPhone(ctx, database.CreateAgentPhoneParams{
				PhonesID: int64ToNullInt64(dbPhoneID),
				AgentID:  agentId,
			}); err != nil {
				cfg.logger.Errorf("error creating agent phone: %v", err)
			}
		}

		cfg.logger.Debugf("- office phones:")
		officePhones := make([]Phone, 0, len(agent.Office.Phones)+len(agent.Office.PhoneList))

		officePhones = append(officePhones, agent.Phones...)

		for _, officePh := range agent.Office.PhoneList {
			officePhones = append(officePhones, officePh)
		}

		for _, phone := range officePhones {
			cfg.logger.Debugf("	* %+v", phone)

			dbPhoneID, err := qtx.GetPhoneID(ctx, database.GetPhoneIDParams{
				Ext:    stringToNullString(phone.Ext),
				Number: stringToNullString(phone.Number),
				Type:   stringToNullString(phone.Type),
			})

			if err != nil {
				if err != sql.ErrNoRows {
					return err
				}

				dbPhoneID, err = qtx.CreatePhone(ctx, database.CreatePhoneParams{
					Ext:     stringToNullString(phone.Ext),
					Number:  stringToNullString(phone.Number),
					Type:    stringToNullString(phone.Type),
					IsValid: boolToNullBool(phone.IsValid),
				})

				if err != nil {
					return err
				}
			}

			if err := qtx.CreateOfficePhone(ctx, database.CreateOfficePhoneParams{
				PhonesID: int64ToNullInt64(dbPhoneID),
				OfficeID: int64ToNullInt64(dbOfficeID),
			}); err != nil {
				cfg.logger.Errorf("error creating office phone: %v", err)
			}
		}

		cfg.logger.Debugf("- agent foreign keys:")
		cfg.logger.Debugf("	* Agent: (AddressID: %d, BrokerID: %d, OfficeID: %d)", dbAddressID, dbBrokerID, dbOfficeID)
		if err := qtx.UpdateAgentForeignKeys(ctx, database.UpdateAgentForeignKeysParams{
			AddressID: int64ToNullInt64(dbAddressID),
			BrokerID:  int64ToNullInt64(dbBrokerID),
			OfficeID:  int64ToNullInt64(dbOfficeID),
			ID:        agent.ID,
		}); err != nil {
			cfg.logger.Errorf("error updating agent foreign keys: %v", err)
		}
		return nil
	})
}
