package model

type AppVersion struct {
	Id                      int    `json:"id"`
	App_release_no          string `json:"app_release_no"`
	App_release_description string `json:"app_release_description"`
	Ref_id                  string `json:"ref_id"`
	Ref_name                string `json:"ref_name"`
	Created_at              string `json:"created_at"`
	Updated_at              string `json:"updated_at"`
	Deleted_at              string `json:"deleted_at"`
	Major_version_no        string `json:"major_version_no"`
	Minor_version_no        string `json:"minor_version_no"`
	Patch_version_no        string `json:"patch_version_no"`
}

type AppHomeScreen struct {
	AccountStatus      string `json:"account_status"`
	KycStatus          string `json:"kyc_status"`
	Message            string `json:"message"`
	UnreadNotification string `json:"unread_notification"`
	App_versions       AppVersion
	User_details       UserDetails
}

type UserAccountDetails struct {
	Status            bool
	Message           string
	User_details      UserDetails
	Totalorders       int `json:"totalorders"`
	LocalAddress      Address
	PermanenetAddress Address
	Ispick_commission []IspickCommission
}

type Address struct {
	Id                    int    `json:"id"`
	User_id               int    `json:"user_id"`
	House_no_or_apartment string `json:"house_no_or_apartment"`
	Street_locality_area  string `json:"street_locality_area"`
	Pincode               string `json:"pincode"`
	City                  string `json:"city"`
	Lat                   string `json:"lat"`
	Lng                   string `json:"lng"`
	Place_id              string `json:"place_id"`
	Google_address        string `json:"google_address"`
	Address_proof_url     string `json:"address_proof_url"`
	Type_of_address       string `json:"type_of_address"`
	Status                string `json:"status"`
	Created_at            string `json:"created_at"`
	Updated_at            string `json:"updated_at"`
	Deleted_at            string `json:"deleted_at"`
	State                 string `json:"state"`
	Country               string `json:"country"`
}

type IspickCommission struct {
	Id         int    `json:"id"`
	Value      int    `json:"value"`
	Condition  string `json:"condition"`
	Charges    int    `json:"charges"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
