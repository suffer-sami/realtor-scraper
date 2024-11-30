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

func (cfg *config) storeAgent(agent Agent) error {
	return cfg.executeTransaction(context.Background(), func(ctx context.Context, qtx *database.Queries) error {
		dbAgent, err := qtx.GetAgent(ctx, agent.ID)
		if err != nil {
			if err != sql.ErrNoRows {
				return err
			}
			dbAgent, err = qtx.CreateAgent(ctx, database.CreateAgentParams{
				ID:                   agent.ID,
				FirstName:            toNullString(agent.FirstName),
				LastName:             toNullString(agent.LastName),
				NickName:             toNullString(agent.NickName),
				PersonName:           toNullString(agent.PersonName),
				Title:                toNullString(agent.Title),
				Slogan:               toNullString(agent.Slogan),
				Email:                toNullString(agent.Email),
				AgentRating:          toNullInt(agent.AgentRating),
				Description:          toNullString(agent.Description),
				RecommendationsCount: toNullInt(agent.RecommendationsCount),
				ReviewCount:          toNullInt(agent.ReviewCount),
				LastUpdated:          strToNullTime(agent.LastUpdated, time.RFC1123),
				FirstMonth:           numericToNullInt(agent.FirstMonth),
				FirstYear:            toNullInt(agent.AgentRating),
				Video:                toNullString(agent.Video),
				WebUrl:               toNullString(agent.WebURL),
				Href:                 toNullString(agent.Href),
			})

			if err != nil {
				return err
			}
		}
		cfg.logger.Infof("Agent: %s", dbAgent.PersonName.String)

		cfg.logger.Debugf("- sales data: %s", agent.RecentlySold.LastSoldDate)
		qtx.CreateSalesData(ctx, database.CreateSalesDataParams{
			Count:        toNullInt(agent.RecentlySold.Count),
			Min:          toNullInt(agent.RecentlySold.Min),
			Max:          toNullInt(agent.RecentlySold.Max),
			LastSoldDate: strToNullTime(agent.RecentlySold.LastSoldDate, time.DateOnly),
			AgentID:      toNullString(dbAgent.ID),
		})

		cfg.logger.Debugf("- listing data: %s", agent.RecentlySold.LastSoldDate)
		qtx.CreateListingsData(ctx, database.CreateListingsDataParams{
			Count:           toNullInt(agent.ForSalePrice.Count),
			Min:             toNullInt(agent.ForSalePrice.Min),
			Max:             toNullInt(agent.ForSalePrice.Max),
			LastListingDate: timeToNullTime(agent.ForSalePrice.LastListingDate),
			AgentID:         toNullString(dbAgent.ID),
		})

		cfg.logger.Debugf("- social medias:")
		for _, socialMedia := range agent.SocialMedias {
			cfg.logger.Debugf("	* %s", socialMedia.Type)
			qtx.CreateSocialMedia(ctx, database.CreateSocialMediaParams{
				Type:    toNullString(socialMedia.Type),
				Href:    toNullString(socialMedia.Href),
				AgentID: toNullString(dbAgent.ID),
			})
		}

		cfg.logger.Debugf("- feed licences:")
		for _, feedLicense := range agent.FeedLicenses {
			cfg.logger.Debugf("	* (%s, %s)", feedLicense.StateCode, feedLicense.Country)
			qtx.CreateFeedLicense(ctx, database.CreateFeedLicenseParams{
				Country:       toNullString(feedLicense.Country),
				LicenseNumber: toNullString(feedLicense.LicenseNumber),
				StateCode:     toNullString(feedLicense.StateCode),
				AgentID:       toNullString(dbAgent.ID),
			})
		}
		return nil
	})
}
