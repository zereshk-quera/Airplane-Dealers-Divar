package repair

import (
	"Airplane-Divar/models"
	"Airplane-Divar/utils"
	"context"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RepairStorer struct {
	db *gorm.DB
}

func NewRepairStorer(db *gorm.DB) RepairStorer {
	return RepairStorer{db: db}
}

func (e RepairStorer) GetByAd(
	ctx context.Context,
	adID int,
	user models.User,
) (models.RepairRequest, error) {
	var repairRequest models.RepairRequest
	query := e.db.WithContext(ctx).Joins("Ads").Where("repair_request.ads_id = ?", adID)

	if user.Role == 4 { // is airline
		query.Where(&models.Ad{UserID: user.ID})
	}
	result := query.First(&repairRequest)

	return repairRequest, result.Error
}

func (e RepairStorer) RequestToRepairCheck(
	ctx context.Context, adID int, userID int,
) error {
	// get user
	var user models.User
	if err := e.db.WithContext(ctx).First(&user, userID).Error; err != nil {
		return err
	}

	// get ad
	var ad models.Ad
	if err := e.db.WithContext(ctx).First(&ad, adID).Error; err != nil {
		return err
	}
	if ad.RepairCheck {
		return errors.New("this had been already repaired by Matin")
	}

	// get or create repair request
	var repairRequest models.RepairRequest
	err := e.db.WithContext(ctx).
		Where("user_id = ? AND ads_id = ?", userID, adID).
		First(&repairRequest).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		return err
	}

	if repairRequest.ID != 0 {
		if repairRequest.Status != utils.MATIN_PENDING_STATUS {
			return errors.New("you had been requested for repairing")
		}
		return nil
	}

	if err := e.db.WithContext(ctx).Model(&models.RepairRequest{}).
		Create(map[string]interface{}{
			"AdsID":  ad.ID,
			"UserID": user.ID,
			"Status": utils.WAIT_FOR_PAYMENT_STATUS,
		}).Error; err != nil {
		return err
	}

	return nil
}

func (e RepairStorer) GetAllRepairRequests(
	ctx context.Context,
	filterAndCondition clause.AndConditions,
	filterOrCondition []clause.OrConditions,
	filterNotCondtion clause.NotConditions,
	page int,
) ([]models.RepairRequest, error) {
	repairRequests := []models.RepairRequest{}

	query := e.db.WithContext(ctx).Scopes(utils.Paginate(page))
	if len(filterAndCondition.Exprs) > 0 {
		query.Where(filterAndCondition)
	}
	if len(filterOrCondition) > 0 {
		for _, filter := range filterOrCondition {
			query.Where(filter)
		}
	}
	if len(filterNotCondtion.Exprs) > 0 {
		query.Where(filterNotCondtion)
	}

	result := query.Find(&repairRequests)
	return repairRequests, result.Error
}

func (e RepairStorer) Update(
	ctx context.Context, repairRequestID int,
	user models.User, body models.UpdateRepairRequest,
) (models.RepairRequest, error) {
	tmpRepairRequest := models.RepairRequest{}
	updatedMap := make(map[string]interface{})

	if user.Role != 1 {
		return tmpRepairRequest, errors.New("not allowed")
	}

	if body.Status != "" {
		updatedMap["status"] = body.Status
		if body.Status == utils.WAIT_FOR_PAYMENT_STATUS {
			return tmpRepairRequest, errors.New("not allowed")
		}
	}

	result := e.db.WithContext(ctx).
		Clauses(clause.Returning{}).
		Model(&tmpRepairRequest).
		Where(
			"id = ?",
			repairRequestID,
		).
		Updates(updatedMap)

	return tmpRepairRequest, result.Error
}

func (e RepairStorer) Delete(
	ctx context.Context,
	adID int,
	user models.User,
) error {
	result := e.db.WithContext(ctx).
		Where(
			"user_id = ? AND ads_id = ? AND status = ?",
			user.ID, adID, utils.WAIT_FOR_PAYMENT_STATUS,
		).
		Delete(&models.RepairRequest{})
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return errors.New("repair request not found")
	}

	return nil

}
