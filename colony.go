package biomebound

import (
	"fmt"
)

// TODO: Buildings and Agents should have a number of ticks that they've been working/turned on for when we switch to production
// going over multiple ticks (a cycle).

var beginnerResourceCounts = [Resource_Max]uint{
	_resource(_landResource_Woods) | (_resource(TreeType_Oak) << 8): 0,
}

type ColonyId int

type Colony struct {
	// TODO: agents are array-of-structs atm. Potentially turn into struct of arrays later. Each agent is a state machine?
	context        *Context
	Id             ColonyId
	tileLocation   TileLocation
	name           string
	agents         []Agent
	resourceCounts [Resource_Max]uint // Current resources in storage
	landResources  [10]ResourceZone   // Available resource zones from land
	//landResources  [LandResource_Max]uint // Available resources from land
	buildings []Building

	// Production and consumption for current tick, the whole integer committed to storage at the start of the next tick.
	currentProduction     [Resource_Max]float64
	currentConsumption    [Resource_Max]float64
	currentTickProduction [Resource_Max]float64

	// resourceConsumers [Resource_Max]*Node
	// landResourceProducers [LandResource_Max]*Node
	// resourceProducers [Resource_Max]*Node
}

func FindBeginnerTileLocation() (tileLocation TileLocation) {
	// Find a suitable tile with a warm or temperate biome for the colony
	foundSuitableTile := false

	// Iterate through every tile in the map instead of random attempts
	for y := 0; y < MapHeight && !foundSuitableTile; y++ {
		for x := 0; x < MapWidth && !foundSuitableTile; x++ {
			// Get the tile at this location
			tile := &Map[y][x]
			if tile.occupied {
				continue
			}

			// Skip water tiles and extreme environments
			if tile.altitude <= 0 || tile.biome == Biome_IceSheet ||
				tile.biome == Biome_SeaIce || tile.biome == Biome_ExtremeDesert {
				continue
			}

			// Mid-temperate or above
			if tile.climate.avgTemp > .4 && tile.hasCoal {
				tileLocation = TileLocation{X: x, Y: y}
				foundSuitableTile = true
			}
		}
	}

	// If we didn't find a suitable tile after checking the entire map, just pick a non-water tile
	if !foundSuitableTile {
		for y := range MapHeight {
			for x := range MapWidth {
				if Map[y][x].occupied {
					continue
				}
				if Map[y][x].altitude > 0 {
					tileLocation = TileLocation{X: x, Y: y}
					foundSuitableTile = true
					break
				}
			}
			if foundSuitableTile {
				break
			}
		}
	}

	return tileLocation
}

func NewColony(context *Context, id ColonyId, name string, initialPopulationSize uint, first bool) *Colony {
	colony := new(Colony)
	colony.Id = id
	colony.context = context
	if first {
		colony.tileLocation = TileLocation{X: 5, Y: 0}
	} else {
		colony.tileLocation = FindBeginnerTileLocation()
	}

	tile := &Map[colony.tileLocation.Y][colony.tileLocation.X]
	tile.occupied = true

	colony.name = name
	colony.agents = make([]Agent, initialPopulationSize)
	colony.resourceCounts = beginnerResourceCounts
	colony.buildings = make([]Building, 0)
	for i := range colony.agents {
		a := &colony.agents[i]

		// TODO: randomize name, age, gender, and sexual orientation
		a.name = fmt.Sprintf("Unknown%2d Unknown", i)
		a.age = 20
		if i < len(colony.agents)/2 {
			a.gender = AgentGender_Male
			a.sexualAttraction = [AgentGender_Max]bool{
				AgentGender_Female: true,
			}
		} else {
			a.gender = AgentGender_Female
			a.sexualAttraction = [AgentGender_Max]bool{
				AgentGender_Male: true,
			}
		}

		a.food = 100
		a.health = 100
		a.state = AgentState_Idle
		a.stress = 20 // 20% stress from starting a new colony off with nothing // TODO: Change this value also based on the chosen tile biome.
		a.familyID = i
		a.assignedZone = -1
	}

	// Basic forest size calculations
	// denseForestTreeSpace := 6                    // square meters, in a dense forest
	// trunkArea := 3                               // square meters
	// treeArea := trunkArea + denseForestTreeSpace // square meters
	// smallForestArea := 50000                     // square meters
	// numberOfTrees := smallForestArea / treeArea

	// colony.landResources[0] = NewResourceZone(0, LandResource_Woods(TreeType_Oak), uint(numberOfTrees))
	colony.landResources[0] = NewResourceZone(0, LandResource(LandResource_Berries), 40000)
	currentI := 1
	for i, treeTypePop := range tile.treeTypes {
		if i >= 9-1 {
			break
		}

		colony.landResources[1+i] = NewResourceZone(ResourceZoneId(1+i), LandResource_Woods(treeTypePop.treeType), uint(treeTypePop.count))
		currentI++
	}
	if tile.hasCoal {
		colony.landResources[currentI] = NewResourceZone(ResourceZoneId(currentI), LandResource(LandResource_Coal), 40000)
		currentI++
	}
	if tile.hasGranite {
		colony.landResources[currentI] = NewResourceZone(ResourceZoneId(currentI), LandResource(LandResource_Granite), 20000)
		currentI++
	}

	return colony
}

func (colony *Colony) Tick() {
	// Commit previous tick's resource production and consumption numbers to storage.
	// This should always be the very first thing done in a tick.
	colony.CommitProductionAndConsumption()

	// The next thing is to update the work/idle/sleep state of each agent based on the current time, day, etc. and their assigned workplace.
	if colony.context.IsWorkTime() {
		for id := range colony.agents {
			a := &colony.agents[id]
			if a.assignedZone != -1 {
				a.state = AgentState_Work
			} else {
				a.state = AgentState_Idle
			}
		}
	} else if colony.context.IsSleepTime() {
		for id := range colony.agents {
			a := &colony.agents[id]
			a.state = AgentState_Sleep
		}
	} else if colony.context.IsFreeTime() {
		for id := range colony.agents {
			a := &colony.agents[id]
			a.state = AgentState_Idle
		}
	}

	// Design TODO: The current tick's consumption always takes from storage and never from the current tick's production.

	// Reset current tick production
	for i := range colony.currentTickProduction {
		colony.currentTickProduction[i] = 0
	}

	// Go through all resource zones (and their buildings/technologies) to set next tick's initial resource production and consumption numbers
	for i := range colony.landResources {
		zone := &colony.landResources[i]
		if zone.landResource == LandResource(LandResource_Unknown) {
			continue
		}

		// Add to the fractional parts this tick's production and consumption values.
		// This might make the production go over 1, which is a good thing, because the commit on next tick
		// will have a chance to see the whole integer and commit it to storage and resource zones before
		// the whole integer part is removed.
		resource := zone.landResource.ToResource()
		if zone.amount > 0 {
			productionThisTick := colony.productionFromZone(zone)
			colony.currentProduction[resource] += productionThisTick
			colony.currentTickProduction[resource] += productionThisTick
		}
	}

	// Design Note TODO:
	// Go through rest of buildings to add to the resource production and consumption numbers.
	// Consumption of buildings should be bounded by the storage plus the production of other buildings.
	// So, we should iterate over all the buildings to get the ones that can use just the storage. Then, we iterate again
	// for the buildings that can consume the production of the storage-using-only buildings along with the rest of the storage.
	// OR RATHER we should probably have some dependency graph here so that we can traverse over it starting from the leaves.

	// TODO: Agents also consume things (like food, water, and clothes, at the bare minimum). These need to be factored in to the consumption numbers.
	foodConsumption := float64(len(colony.agents)) * float64(.000555)
	colony.currentConsumption[Resource_Berries] = min(foodConsumption, float64(colony.resourceCounts[Resource_Berries]))
}

func (colony *Colony) GetCurrentTickProduction(resource _resource) float64 {
	return colony.currentTickProduction[resource]
}

// Commits the previous tick's production and consumption to storage at the start of each tick.
func (colony *Colony) CommitProductionAndConsumption() {
	// Track production per zone to ensure correct subtraction
	zoneProduction := make(map[_resource]map[int]uint)

	// First pass: calculate whole number production for each zone
	for i := range colony.landResources {
		zone := &colony.landResources[i]
		if zone.landResource == LandResource(LandResource_Unknown) {
			continue
		}

		resource := zone.landResource.ToResource()
		if zoneProduction[resource] == nil {
			zoneProduction[resource] = make(map[int]uint)
		}

		zoneWholeProduction := uint(colony.currentProduction[zone.landResource.ToResource()])

		// Ensure we don't over-harvest from this specific zone
		if zoneWholeProduction > zone.amount {
			zoneWholeProduction = zone.amount
		}

		zoneProduction[resource][i] = zoneWholeProduction
	}

	// Update zone amounts and commit the production and consumption to storage
	for resource := range Resource_Max {
		var totalProduction uint = 0

		// Sum up production from all zones that produce this resource
		for zoneIndex := range colony.landResources {
			if productions, exists := zoneProduction[resource]; exists {
				if zoneAmount, hasZone := productions[zoneIndex]; hasZone {
					totalProduction += zoneAmount
				}
			}
		}

		// Handle consumption
		consumptionWhole := uint(colony.currentConsumption[resource])

		// Update Storage
		if consumptionWhole >= colony.resourceCounts[resource]+totalProduction {
			colony.resourceCounts[resource] = 0
		} else {
			colony.resourceCounts[resource] += totalProduction - consumptionWhole
		}

		// Update individual zone amounts
		for i := range colony.landResources {
			if productions, exists := zoneProduction[resource]; exists {
				if zoneAmount, hasZone := productions[i]; hasZone {
					// Subtract only this zone's production from its amount
					if zoneAmount >= colony.landResources[i].amount {
						colony.landResources[i].amount = 0
					} else {
						colony.landResources[i].amount -= zoneAmount
					}
				}
			}
		}

		// Keep fractional parts for next tick
		colony.currentProduction[resource] -= float64(totalProduction)
		colony.currentConsumption[resource] -= float64(consumptionWhole)
	}
}

// Calculates the current tick's production from a zone, taking into account the zone's amount, the number of workers, *and* the state and stats of each worker.
func (colony *Colony) productionFromZone(zone *ResourceZone) float64 {
	cap := float64(zone.amount)

	var numberOfActiveWorkers float64 = 0
	for _, workerId := range zone.workers {
		worker := &colony.agents[workerId]
		if worker.state == AgentState_Work {
			numberOfActiveWorkers += 1
		}
	}

	return min(productionPerDayToPerTicks(zone.landResource.PerDayProductionPerAgent())*numberOfActiveWorkers, cap)
}

func productionPerDayToPerTicks(perDay float64) float64 {
	return perDay / 24 / 60 / 60 * float64(InGameSecondsPerTick)
}
