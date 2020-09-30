# gopayroll
Payroll Application implemented with Go!
# Employee Endpoints

### Parameters
**_id** is the unique ID of the user

**name** assigned name of the user (doesn't need to be unique)

**username** assigned username of the user (needs to be unique)

**password** assigned password of the user (minimum of 8 characters)

**birthday** assigned birthday of the user 

**email** assigned birthday of the user 


## `POST /user/signup`

Creates a user in the db with the coressponding **name**, **username**, **password**, **birthday**, **email**


## Request
```
{
	"name": "Kurt"
	"username": "Kurtified"
	"password": "123456789"
	"birthday": "1988-01-20"
	"email": "kurt@google.com"
}	
```


## Response
```
{
	"data": {
		"_id": "5e25bfed830ff6000c7ecb3e",
		"username": "Kurtified",
		"user" : {
			"token": "27acf480-c19f-4a40-ae32-42765ad49ba127acf480-c19f-4a40-ae32-42765ad49ba127acf480-c19f-4a40-ae32-42765ad49ba1"
		}    
	},
	"success": true
}
```
---
## `GET /user/login/:_id`

Retrieve a user and its corresponding token from the service using the coresponding **_id**

## Response
```
{
	"data": {
		"_id": "5e25bfed830ff6000c7ecb3e",
		"username": "Kurtified"
	},
	"success": true
}
```
---
## `GET /user/:_id/overtime?per_page=x&page=y`

Retrieve a user's regular and special overtime from the service using the coresponding **_id** The body parameters contains a startDate and endDate. 

The entries has pagination and the page and per_page should have default values. page is the page offset in the entry list starting from 1 and should have a default value of 1. per_page is the number of entries per page and should have a defualt value of 10
## Request
```
{
	"startDate": "2019-10-15"
	"endDate": "2020-10-31"
	"token": "27acf480-c19f-4a40-ae32-42765ad49ba127acf480-c19f-4a40-ae32-42765ad49ba127acf480-c19f-4a40-ae32-42765ad49ba1"
}	
```
If entries are blank, retrieve for current Payroll Period. 
```
{
	"token": "27acf480-c19f-4a40-ae32-42765ad49ba127acf480-c19f-4a40-ae32-42765ad49ba127acf480-c19f-4a40-ae32-42765ad49ba1"
}	
```
## Response
```
{
	"data": {
		"_id": "5e25bfed830ff6000c7ecb3e",
		"username": "Kurtified",
		"overtime" : {
			"regular": [
				{	
					"date": "8/28/2020",
					"hours": 2
				},
				{	
					"date": "9/4/2020",
					"hours": 1.5
				},
				{	
					"date": "9/8/2020",
					"hours": 2
				}
			],
			"special": [
				{	
					"date": "8/28/2020",
					"hours": 12
				},
				{	
					"date": "9/4/2020",
					"hours": 12
				}
			]
		}
	},
	"success": true
}
```
---
## `GET /user/:_id/leaves?per_page=x&page=y`

Retrieve a user's sick leaves, vacation leaves, and absences from the service using the coresponding **_id** The body parameters contains a startDate and endDate. 

The entries has pagination and the page and per_page should have default values. page is the page offset in the entry list starting from 1 and should have a default value of 1. per_page is the number of entries per page and should have a defualt value of 10
## Request
```
{
	"startDate": "2019-10-15"
	"endDate": "2020-10-31"
	"token": "27acf480-c19f-4a40-ae32-42765ad49ba127acf480-c19f-4a40-ae32-42765ad49ba127acf480-c19f-4a40-ae32-42765ad49ba1"
}	
```
If entries are blank, retrieve for current Payroll Period. 
```
{
	"token": "27acf480-c19f-4a40-ae32-42765ad49ba127acf480-c19f-4a40-ae32-42765ad49ba127acf480-c19f-4a40-ae32-42765ad49ba1"
}	
```
## Response
```
{
	"data": {
		"_id": "5e25bfed830ff6000c7ecb3e",
		"username": "Kurtified",
		"leaves" : {	
			"sick": [
				{	
					"date": "8/28/2020T12:00:00",
					"reason": "I am not feeling well today",
					"filingDate" : "8/30/2020",
					"approvalDate" : "8/31/2020",
					"duration" : 0.5
					"period" : "am"
				},
				{	
					"date": "8/28/2020T08:00:00",
					"reason": "I am not feeling well today",
					"filingDate" : "8/30/2020",
					"approvalDate" : "8/31/2020",
					"duration" : 1
				},
				{	
					"date": "8/28/2020T08:00:00",
					"reason": "I am at the hospital for checkup today",
					"filingDate" : "9/29/2020",
					"approvalDate" : "9/29/2020",
					"duration" : 0.5
					"period" : "pm"
				},
				{	
					"date": "10/8/2020",
					"reason": "I feel unwell today",
					"filingDate" : "10/9/2020"
					"approvalDate" : "10/10/2020",
					"duration" : 1 
					
				}
			],
			"vacation": [
				{	
					"date": "8/28/2020",
					"reason": "Travel Plans",
					"filingDate" : "8/20/2020"
					"approvalDate" : "8/20/2020",
					"duration" : 1
				},
				{	
					"date": "9/4/2020",
					"reason": "Travel Plans",
					"filingDate" : "9/5/2020"
					"approvalDate" : "9/6/2020",
					"duration" : 0.5
					"period" : "pm"
				}
			],
			"absent": [
				{	
					"date": "9/21/2020",
					"reason": "Travel Plans",
					"filingDate" : "9/22/2020"
					"approvalDate" : "9/22/2020",
					"duration" : 0.5
					"period" : "pm"
				},
				{	
					"date": "9/25/2020",
					"reason": "Travel Plans",
					"filingDate" : "9/26/2020"
					"approvalDate" : "9/28/2020",
					"duration" : 0.5
					"period" : "am"
				}
			]
		}
	},
	"success": true
}
```
---
## `GET /user/:_id/loans?per_page=x&page=y`

Retrieve a user's loans and cash advance from the service using the coresponding **_id** The body parameters contains a startDate and endDate. 

The entries has pagination and the page and per_page should have default values. page is the page offset in the entry list starting from 1 and should have a default value of 1. per_page is the number of entries per page and should have a defualt value of 10
## Request
```
{
	"startDate": "2019-10-15"
	"endDate": "2020-10-31"
	"token": "27acf480-c19f-4a40-ae32-42765ad49ba127acf480-c19f-4a40-ae32-42765ad49ba127acf480-c19f-4a40-ae32-42765ad49ba1"
}	
```
If entries are blank, retrieve for current Payroll Period. 
```
{
	"token": "27acf480-c19f-4a40-ae32-42765ad49ba127acf480-c19f-4a40-ae32-42765ad49ba127acf480-c19f-4a40-ae32-42765ad49ba1"
}	
```
## Response
```
{
	"data": {
		"_id": "5e25bfed830ff6000c7ecb3e",
		"username": "Kurtified",
		"loans" : [
			{
				"startDate": "2020-10-15"
				"endDate": "2020-10-15" 
				"name": "SSS",
				"debt": 10920.00,
				"monthly": 496.36
				"remaining": 5460.00,
				"months": 24
				"completion: 11 /* (11/24) payments */
			},
			{
				"startDate": "2020-10-31"
				"endDate": "2022-10-31" 
				"name": "HDMF",
				"debt": 18840.96,
				"monthly": 785.04,
				"remaining": 17270.88,
				"months": 24
				"completion: 2 /* (2/24) payments */
			},
			{
				"startDate": "2019-12-31"
				"endDate": "2022-10-31" 
				"name": "RCBC",
				"debt": 20000.8,
				"monthly": 1000.04
				"remaining": 10000.04,
				"months": 24
				"completion: 10 /* (10/24) payments */
			},
			{
			
				"startDate": "2020-10-1"
				"endDate": "2020-10-15" 
				"name": "CashAdvance",
				"debt": 6384.00,
				"semiMonthly": 6384.00
				"remaining": 6384.00,
				"months": 0.5
				"completion: 1 /* (0/1) payments */
				"sources" : [
					{
						"date": "2020-9-30"
						"name": "RPV/ER FUND"
						"debt": 1000
					},
					{
						"date": "2020-9-28"
						"name": "Office/Vit.C"
						"debt": 500
					},
					{
						"date": "2020-9-31"
						"name": "Trisha/CA"
						"debt": 1500
					},
					{
						"date": "2020-9-26"
						"name": "Trisha/CA"
						"debt": 1500
					},
					{
						"date": "2020-9-26"
						"name": "RICO/A"
						"debt": 500
					},
					{
						"date": "2020-9-26"
						"name": "RICO/B"
						"debt": 500
					},
					{
						"date": "2020-9-26"
						"name": "JV"
						"debt": 884
					}
				]
				
			}
		]
	},
	"success": true
}
```
---
## `GET /user/:_id/earnings?per_page=x&page=y`

Retrieve a user's earnings from the service using the coresponding **_id** The body parameters contains a startDate and endDate. 

The entries has pagination and the page and per_page should have default values. page is the page offset in the entry list starting from 1 and should have a default value of 1. per_page is the number of entries per page and should have a defualt value of 10
## Request
```
{
	"startDate": "2019-10-15"
	"endDate": "2020-10-31"
	"token": "27acf480-c19f-4a40-ae32-42765ad49ba127acf480-c19f-4a40-ae32-42765ad49ba127acf480-c19f-4a40-ae32-42765ad49ba1"
}	
```
If entries are blank, retrieve for current Payroll Period. 
```
{
	"token": "27acf480-c19f-4a40-ae32-42765ad49ba127acf480-c19f-4a40-ae32-42765ad49ba127acf480-c19f-4a40-ae32-42765ad49ba1"
}	
```
## Response
```
{
	"data": {
		"_id": "5e25bfed830ff6000c7ecb3e",
		"username": "Kurtified",
		"earnings" : {
			Rate              float32            `json:"rate"`                               // computed from BasicSalary OR fixed
			BasicSalary       float32            `json:"basicSalary"`                        // arbitrary value
			Deminimis         float32            `json:"deminimis"`                          // arbitrary value
			"allowances": [
				{
					"name": "ECOLA",
					"amount": 500
				},
				{
					"name": "Transpo",
					"amount": 500
				},
				{
					"name": "Meal",
					"amount": 500
				},				{
					"name": "Others",
					"amount": 500
				}
			]
			"holiday": 600.00 
			"birthday: 500.00
			"overtime":  1302.09 
			"SL/VL":  600.00
			"adjustment": 500
			"total": 10,218.76
		}
	},
	"success": true
}
```
---


# Admin Endpoints
To be continued