package biomebound

type ResourceZoneId int

// A resource zone is a zone on the land of a particular resource that can be harvested
// either manually or with buildings. Each colony region can have up to 10 different resource zones.
type ResourceZone struct {
	id           ResourceZoneId
	landResource _landResource
	amount       uint
	workers      []AgentId
}

func NewResourceZone(id ResourceZoneId, resource _landResource, amount uint) ResourceZone {
	return ResourceZone{landResource: resource, amount: amount, workers: make([]AgentId, 0, 20)}
}

func (zone_id ResourceZoneId) AddWorker(colony *Colony, id AgentId, a *Agent) {
	zone := &colony.landResources[zone_id]

	// If already assigned a zone, return
	if a.assignedZone != -1 {
		return
	}

	a.assignedZone = zone.id
	zone.workers = append(zone.workers, id)
}

func (zone_id ResourceZoneId) RemoveWorker(colony *Colony, id AgentId, a *Agent) {
	zone := &colony.landResources[zone_id]

	if len(zone.workers) == 0 {
		return
	}
	if a.assignedZone == -1 {
		return
	}

	// Reset the agent's assignedZone
	a.assignedZone = -1

	// Remove the worker by replacing it with the last element and slicing off the last element
	for i, workerId := range zone.workers {
		if workerId == id {
			zone.workers[i] = zone.workers[len(zone.workers)-1]
			zone.workers = zone.workers[:len(zone.workers)-1]
			break
		}
	}
}

func (zone_id ResourceZoneId) RemoveLastWorker(colony *Colony) {
	zone := &colony.landResources[zone_id]

	lastWorkerId := zone.workers[len(zone.workers)-1]
	colony.agents[lastWorkerId].assignedZone = -1
	zone.workers = zone.workers[:len(zone.workers)-1]
}
