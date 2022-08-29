package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-competitor-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

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

func (c *SAPAPICaller) CompetitorCollection(objectID, competitorID string) {
	competitorCollectionData, err := c.callCompetitorSrvAPIRequirementCompetitorCollection("CompetitorCollection", objectID, competitorID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(competitorCollectionData)
}

func (c *SAPAPICaller) callCompetitorSrvAPIRequirementCompetitorCollection(api, objectID, competitorID string) ([]sap_api_output_formatter.CompetitorCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithCompetitorCollection(req, objectID, competitorID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCompetitorCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithCompetitorCollection(req *http.Request, objectID, competitorID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ObjectID eq '%s' and CompetitorID eq '%s'", objectID, competitorID))
	req.URL.RawQuery = params.Encode()
}
