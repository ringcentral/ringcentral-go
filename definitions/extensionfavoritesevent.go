package definitions

// ExtensionFavoritesEvent Extension Favorites Event
type ExtensionFavoritesEvent struct {
	Uuid string `json:"uuid"`
	Event string `json:"event"`
	Timestamp string `json:"timestamp"`
	SubscriptionId string `json:"subscriptionId"`
	Body ExtensionFavoritesEventBody `json:"body"`
}
