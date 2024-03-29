# sap-api-integrations-competitor-reads  
sap-api-integrations-competitor-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API 競合データを取得するマイクロサービスです。  
sap-api-integrations-competitor-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-competitor-reads は、オンプレミス版である（＝クラウド版ではない）SAPC4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/competitor/overview  

## 動作環境
sap-api-integrations-competitor-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須） 

## クラウド環境での利用  
sap-api-integrations-competitor-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-competitor-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/competitor/overview 
* APIサービス名(=baseURL): c4codataapi

## 本レポジトリ に 含まれる API名
sap-api-integrations-competitor-reads には、次の API をコールするためのリソースが含まれています。  

* CompetitorCollection（競合 - 競合）  

## API への 値入力条件 の 初期値
sap-api-integrations-competitor-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.CompetitorCollection.ObjectID（対象ID）  
* inoutSDC.CompetitorCollection.CompetitorID（競合ID）  


## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"CompetitorCollection" が指定されています。    
  
```
	"api_schema": "Competitor",
	"accepter": ["CompetitorCollection"],
	"competitor_code": "1000000",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "Competitor",
	"accepter": ["ALL"],
	"competitor_code": "1000000",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetCompetitor(objectID, competitorID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "CompetitorCollection":
			func() {
				c.CompetitorCollection(objectID, competitorID)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 競合 の 競合データ が取得された結果の JSON の例です。  
以下の項目のうち、"ObjectID" ～ "ETag" は、/SAP_API_Output_Formatter/type.go 内 の Type CompetitorCollection {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona5/bitbucket/sap-api-integrations-competitor-reads/SAP_API_Caller/caller.go#L58",
	"function": "sap-api-integrations-competitor-reads/SAP_API_Caller.(*SAPAPICaller).CompetitorCollection",
	"level": "INFO",
	"message": [
		{
			"ObjectID": "00163E03A0701EE288BA0BC721A2B04A",
			"CompetitorID": "1000000",
			"CompetitorUUID": "00163E03-A070-1EE2-88BA-0BC721A2B04A",
			"StatusCode": "2",
			"StatusCodeText": "Active",
			"ClassificationCode": "A",
			"ClassificationCodeText": "Big Threat",
			"BusinessPartnerFormattedName": "Kingston Technologies",
			"Name": "Kingston Technologies",
			"AdditionalName": "",
			"FormattedPostalAddressDescription": "3750 Curlington Circle / New York 10299 / US",
			"CountryCode": "US",
			"CountryCodeText": "United States",
			"RegionCode": "NY",
			"RegionCodeText": "New York",
			"CareOfName": "",
			"AddressLine1": "",
			"AddressLine2": "",
			"HouseNumber": "3750",
			"Street": "Curlington Circle",
			"AddressLine4": "",
			"AddressLine5": "",
			"City": "New York",
			"AdditionalCityName": "",
			"District": "",
			"County": "",
			"CompanyPostalCode": "",
			"StreetPostalCode": "10299",
			"POBoxPostalCode": "",
			"POBox": "",
			"POBoxDeviatingCountryCode": "",
			"POBoxDeviatingCountryCodeText": "",
			"POBoxDeviatingCity": "",
			"TimeZoneCode": "EST",
			"TimeZoneCodeText": "(UTC-05:00) Eastern Time (New York, Quebec)",
			"TaxJurisdictionCode": "",
			"TaxJurisdictionCodeText": "",
			"POBoxDeviatingStateCode": "",
			"POBoxDeviatingStateCodeText": "",
			"Phone": "+1 (212) 345-6788",
			"Fax": "+1 (212) 345-9000",
			"Email": "",
			"WebSite": "www.Kingston.com",
			"LanguageCode": "",
			"LanguageCodeText": "",
			"BestReachedByCode": "",
			"BestReachedByCodeText": "",
			"NormalisedPhone": "+12123456788",
			"CreatedOn": "2012-10-29T22:57:17+09:00",
			"CreatedBy": "SAP WORKER",
			"CreatedByIdentityUUID": "00163E03-A070-1EE2-88B6-F539A6B028F3",
			"ChangedOn": "2012-10-29T22:59:25+09:00",
			"ChangedBy": "SAP WORKER",
			"ChangedByIdentityUUID": "00163E03-A070-1EE2-88B6-F539A6B028F3",
			"EntityLastChangedOn": "2012-10-29T22:59:25+09:00",
			"ETag": "2012-10-29T22:59:25+09:00"
		}
	],
	"time": "2022-06-04T16:01:40+09:00"
}
```