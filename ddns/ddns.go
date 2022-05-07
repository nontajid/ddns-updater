package ddns

import (
	"log"

	"nontajid.com/cloudflareclient"
)

type DDNSIpUpdater interface {
	updateIp(recordInfo cloudflareclient.NewRecordInfo)
}

func UpdateIp(recordInfo cloudflareclient.NewRecordInfo, updater DDNSIpUpdater) {
	log.Printf("Initailing DNS record update with with information %+v\n", recordInfo)
	updater.updateIp(recordInfo)
}
