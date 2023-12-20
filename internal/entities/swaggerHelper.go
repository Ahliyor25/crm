package entities

// Основыне элементы для сваггера

type swaggerAuthorizationData struct {
	// Номер телефона
	//
	// in: query
	// required: true
	// default: ahliyor.shodiev@icloud.com
	Email string `json:"email"`
	// Пароль
	//
	// in: query
	// required: true
	// default: !Qwerty123
	Password string `json:"password"`
}

//
//type swaggerMerchantID struct {
//	// ID Мерчанта
//	//
//	// in: query
//	// required: true
//	// default: 3ca4a036-8aa0-4f84-83d4-6212d33501fb
//	MerchantID uuid.UUID `json:"merchant_id"`
//}
//
//type swaggerMerchantData struct {
//	// Название мерчанта
//	//
//	// required: true
//	// default: ЖДММ "Багровый лист осеннего солнцестояния"
//	Name string `json:"name"`
//	// Адрес мерчанта
//	//
//	// required: true
//	// default: ул. Джаллодини Руми, д.16
//	Address string `json:"address"`
//	// Город мерчанта
//	//
//	// required: true
//	// default: Душанбе
//	City string `json:"city"`
//	// Баланс мерчанта
//	// Примечание: Баланс указывается в дирамах
//	//
//	// required: true
//	// default: 8812
//	Balance float64 `json:"balance"`
//}
//
//type swaggerUserID struct {
//	// ID Пользователя
//	//
//	// in: query
//	// required: true
//	// default: 3ca4a036-8aa0-4f84-83d4-6212d33501fb
//	UserID uuid.UUID `json:"user_id"`
//}
//
//type swaggerUserData struct {
//	// Имя пользователя
//	//
//	// in: query
//	// required: true
//	// default: Святозавр
//	Name string `json:"name"`
//	// Фамилия пользователя
//	//
//	// in: query
//	// required: true
//	// default: Курдюмов
//	Surname string `json:"surname"`
//	// Номер телефона
//	//
//	// in: query
//	// required: true
//	// default: 992000331333
//	Phone string `json:"phone"`
//	// Номер телефона
//	//
//	// in: query
//	// required: true
//	// default: !Qwerty123
//	Password string `json:"password" gorm:"-"`
//	// Роль пользователя
//	//
//	// in: query
//	// required: true
//	// default: !Qwerty123
//	RoleID uuid.UUID `json:"role_id"`
//	// Статус активности
//	//
//	// in: query
//	// required: true
//	// default: true
//	Active bool `json:"active" gorm:"-"`
//	// Статус активности
//	//
//	// in: query
//	// required: true
//	// default: !Qwerty123
//	MerchantID *uuid.NullUUID `json:"merchant_id"`
//}
//
//type swaggerClientData struct {
//	// Номер телефона
//	//
//	// in: query
//	// required: true
//	// default: 992919877799
//	Phone string `json:"phone"`
//}
//
//type swaggerClientPhone struct {
//	// Номер телефона
//	//
//	// in: query
//	// required: true
//	// default: 992919877799
//	Phone string `json:"phone"`
//}
