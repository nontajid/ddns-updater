package ddns

import (
	"log"

	"nontajid.com/cloudflareclient"
)

type createCloudFlareUpdater struct {
	connectionInfo cloudflareclient.CloudFlareConnectionInfo
}

func (connector createCloudFlareUpdater) updateIp(recordInfo cloudflareclient.NewRecordInfo) {
	log.Println("Updating Ip from Clodflare Connector")

	connectionInfo := connector.connectionInfo
	zone := cloudflareclient.FetchZoneInfo(connectionInfo)
	zoneId := getTargetZoneId(recordInfo.Tld, zone.Result)
	log.Printf("target zone Id for %s is %s", recordInfo.Tld, zoneId)

	record := cloudflareclient.FetchRecordInfo(connectionInfo, zoneId)
	recordId := getTargetRecordId(recordInfo.Domain, "A", record.Result)

	log.Printf("target record Id for %s is %s", recordInfo.Domain, recordId)

	cloudflareclient.UpdateRemoteRecord(connectionInfo, recordInfo, zoneId, recordId)
}

func getTargetZoneId(tld string, zoneResult []cloudflareclient.ZoneResult) string {
	for _, result := range zoneResult {
		if result.Name == tld {
			return result.Id
		}
	}

	log.Fatalln("target tld is not listed")
	return ""
}

func getTargetRecordId(domain string, recordType string, recordResult []cloudflareclient.RecordResult) string {
	for _, result := range recordResult {
		if result.Name == domain && result.Type == recordType {
			return result.Id
		}
	}

	log.Fatalln("target domain is not listed")
	return ""
}

func CreateCloudFlareUpdater(connectionInfo cloudflareclient.CloudFlareConnectionInfo) *createCloudFlareUpdater {
	// Return Ip Updater
	return &createCloudFlareUpdater{
		connectionInfo: connectionInfo,
	}
}
