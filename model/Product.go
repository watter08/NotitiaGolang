package model

type Product struct {
	ProductCode					int 						`json:"ProductCode"`
	Desciption          		string                   	`json:"Desciption"`
	Title         				string                    	`json:"Title"`
	Name         				string                    	`json:"Name"`
	CategoryCode        		int                   		`json:"CategoryCode"`
	SubCategoryCode     		int                   		`json:"SubCategoryCode"`
	ModelCode           		int                   		`json:"ModelCode"`
	WeightCode          		int                   		`json:"WeightCode"`
	ImageCode           		int                   		`json:"ImageCode"`
	Year           				int                    		`json:"Year"`
	Status              		int                      	`json:"Status"`
	Price		        		float64                     `json:"Price"`
	User						string 						`json:"User"`
	Terminal					string 						`json:"Terminal"`
	CreatedAt   				string                   	`json:"CreatedAt"`
	UpdatedAd   				string                   	`json:"UpdatedAd"`
	Qrcode              		string                   	`json:"Qrcode"`
}

type Count struct {
	CountCode					int 						`json:"CountCode"`
	Stock						float64						`json:"Stock"`
	Status						int 						`json:"Status"`
	User						string 						`json:"User"`
	Terminal					string 						`json:"Terminal"`
	CreatedAt   		    	string                		`json:"CreatedAt"`
	UpdatedAd   		    	string                		`json:"UpdatedAd"`
}


type ProductType struct {
	ProductTypeCode				int 						`json:"ProductTypeCode"`
	Name						string 						`json:"Name"`
	Description					string 						`json:"Description"`
	Status						int 						`json:"Status"`
	User						string 						`json:"User"`
	Terminal					string 						`json:"Terminal"`
	CreatedAt   		    	string                		`json:"CreatedAt"`
	UpdatedAd   		    	string                		`json:"UpdatedAd"`
}


type Color struct { 
	ColorCode 				int							`json:"ColorCode"`
	Name					string 						`json:"Name"`
	Description				string 						`json:"Description"`
	Status					int 						`json:"Status"`
	User					string 						`json:"User"`
	Terminal				string 						`json:"Terminal"`
	CreatedAt   		    string                		`json:"CreatedAt"`
	UpdatedAd   		    string                		`json:"UpdatedAd"`
}


type Category struct { 
	CategoryCode			int							`json:"CategoryCode"`
	Name					string 						`json:"Name"`
	Description				string 						`json:"Description"`
	Status					int 						`json:"Status"`
	User					string 						`json:"User"`
	Terminal				string 						`json:"Terminal"`
	CreatedAt   		    string                		`json:"CreatedAt"`
	UpdatedAd   		    string                		`json:"UpdatedAd"`
}


type SubCategory struct { 
	SubCategoryCode			int							`json:"SubCategoryCode"`
	CategoryCode			int							`json:"CategoryCode"`
	Name					string 						`json:"Name"`
	Description				string 						`json:"Description"`
	Status					int 						`json:"Status"`
	User					string 						`json:"User"`
	Terminal				string 						`json:"Terminal"`
	CreatedAt   		    string                		`json:"CreatedAt"`
	UpdatedAd   		    string                		`json:"UpdatedAd"`
}


type Price struct {
	PriceCode					int						`json:"PriceCode"`
	ProductCode					int						`json:"ProductCode"`
	Description					string					`json:"Description"`
	Price						float64					`json:"Price"`
	MinPrice					float64					`json:"MinPrice"`
	MaxPrice					float64					`json:"MaxPrice"`
	OfferCode					int						`json:"OfferCode"`
	Status						int 					`json:"Status"`
	User						string					`json:"User"`
	Terminal					string					`json:"Terminal"`
	CreatedAt   		    	string	              	`json:"CreatedAt"`
	UpdatedAd   		    	string	              	`json:"UpdatedAd"`
}


type Offer struct {
	OfferCode					int						`json:"OfferCode"`
	ProductCode					int						`json:"ProductCode"`
	Mount						float64					`json:"Mount"`
	Status						int 					`json:"Status"`
	User						string					`json:"User"`
	Terminal					string					`json:"Terminal"`
	CreatedAt   		    	string	              	`json:"CreatedAt"`
	UpdatedAd   		    	string	              	`json:"UpdatedAd"`
}


type Image struct {
	ImageCode 					int						`json:"ImageCode"`
	ProductCode					int						`json:"ProductCode"`
	Url							string					`json:"Url"`
	Status						int 					`json:"Status"`
	User						string					`json:"User"`
	Terminal					string					`json:"Terminal"`
	CreatedAt   		    	string	              	`json:"CreatedAt"`
	UpdatedAd   		    	string	              	`json:"UpdatedAd"`
}