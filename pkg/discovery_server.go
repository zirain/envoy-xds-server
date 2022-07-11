package pkg

import (

	//core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/sirupsen/logrus"
)

// DiscoveryStream is a server interface for XDS.
type DiscoveryStream = discovery.AggregatedDiscoveryService_StreamAggregatedResourcesServer

// DiscoveryClient is a client interface for XDS.
type DiscoveryClient = discovery.AggregatedDiscoveryService_StreamAggregatedResourcesClient

type DiscoveryServer struct {
}

// StreamAggregatedResources implements the ADS interface.
func (s *DiscoveryServer) StreamAggregatedResources(stream discovery.AggregatedDiscoveryService_StreamAggregatedResourcesServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		switch req.TypeUrl {
		default:
			logrus.Infof("receive node_id: %s TypeUrl: %s ResourceNames: %v", req.Node.Id, req.TypeUrl, req.ResourceNames)
		}
	}
}

func (s *DiscoveryServer) DeltaAggregatedResources(stream discovery.AggregatedDiscoveryService_DeltaAggregatedResourcesServer) error {
	return nil
}

func (s *DiscoveryServer) Stream(stream DiscoveryStream) error {

	return nil
}
