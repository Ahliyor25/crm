package entities

// swagger:parameters authLogin
type swaggerAuthorization struct {
	// in:body
	Body struct {
		swaggerAuthorizationData
	}
}

//// swagger:parameters merchantGet merchantDelete
//type swaggerMerchantIDRequest struct {
//	swaggerMerchantID
//}
//
//// swagger:parameters merchantCreate
//type swaggerMerchantPOST struct {
//	// in:body
//	Body struct {
//		swaggerMerchantData
//	}
//}
//
//// swagger:parameters merchantUpdate
//type swaggerMerchantPUT struct {
//	// in:body
//	Body struct {
//		swaggerMerchantID
//		swaggerMerchantData
//	}
//}
//
//// swagger:parameters UserGet UserDelete
//type swaggerUserIDRequest struct {
//	swaggerUserID
//}
//
//// swagger:parameters UserCreate
//type swaggerUserPOST struct {
//	// in:body
//	Body struct {
//		swaggerUserData
//	}
//}
//
//// swagger:parameters UserUpdate
//type swaggerUserPUT struct {
//	// in:body
//	Body struct {
//		swaggerUserID
//		swaggerUserData
//	}
//}
//
//// swagger:parameters clientCreate
//type swaggerClientPOST struct {
//	// in:body
//	Body struct {
//		swaggerClientData
//	}
//}
//
//// swagger:parameters clientGet clientGet
//type swaggerClientGet struct {
//	swaggerClientPhone
//}
