package biomebound

// These are the resources that are part of the land, excluding animals, and are usually
// held within resource zones. They are harvested into regular resources.
// type LandResource uint16

type _landResource uint16

func LandResource(Type LandResourceType) _landResource {
	return _landResource(uint16(Type) | (uint16(TreeType_Unknown) << 8))
}

func LandResource_Woods(Tree TreeType) _landResource {
	return _landResource(uint16(_landResource_Woods) | (uint16(Tree) << 8))
}

func (resource _landResource) Type() LandResourceType {
	return LandResourceType(uint16(resource) & 255)
}

func (resource _landResource) Tree() TreeType {
	return TreeType(uint16(resource) >> 8)
}

const LandResource_Max = _landResource((255 << 8) | 255)

type LandResourceType uint8

const (
	LandResource_Unknown LandResourceType = iota
	LandResource_Dirt

	// Water Types
	LandResource_Pond
	LandResource_Lake_Vertical
	LandResource_Lake_Horizontal

	// Tree Types
	_landResource_Woods

	// Fuel Ore
	LandResource_Coal

	// Sone
	LandResource_Clay

	LandResource_Granite
	LandResource_Limestone
	LandResource_Sandstone
	LandResource_Marble
	LandResource_Slate

	// Naturally-ocurring Metal Ores
	LandResource_Iron
	LandResource_Aluminum
	LandResource_Zinc
	LandResource_Copper
	LandResource_Nickel
	LandResource_Tin
	LandResource_Silver
	LandResource_Gold

	// Plants
	LandResource_Haygrass
	LandResource_RawRice
	LandResource_Berries
	LandResource_Potatoes
	LandResource_Corn
	LandResource_Agave
	LandResource_Mushrooms
	LandResource_Strawberries

	LandResourceType_Max
)

func (resource _landResource) PerDayProductionPerAgent() float64 {
	switch resource.Type() {
	case LandResource_Unknown:
		return 0
	case LandResource_Dirt:
		return 0
	case LandResource_Pond:
		return 20 * 20 // Liters - usually 20 at a time, 20 trips per 24 hours
	case LandResource_Lake_Vertical:
		return 20 * 20 // Liters - usually 20 at a time, 20 trips per 24 hours
	case LandResource_Lake_Horizontal:
		return 20 * 20 // Liters - usually 20 at a time, 20 trips per 24 hours
	case _landResource_Woods:
		return 100 // Trees - 100 trees cut per 24 hours, 10 to 50 trees sawmilled per day
	case LandResource_Coal:
		return 0
	case LandResource_Clay:
		return 0
	case LandResource_Granite:
		return 0
	case LandResource_Limestone:
		return 0
	case LandResource_Sandstone:
		return 0
	case LandResource_Marble:
		return 0
	case LandResource_Slate:
		return 0
	case LandResource_Iron:
		return 0
	case LandResource_Aluminum:
		return 0
	case LandResource_Zinc:
		return 0
	case LandResource_Copper:
		return 0
	case LandResource_Nickel:
		return 0
	case LandResource_Tin:
		return 0
	case LandResource_Silver:
		return 0
	case LandResource_Gold:
		return 0
	case LandResource_Haygrass:
		return 0
	case LandResource_RawRice:
		return 0
	case LandResource_Berries:
		return 50000 // OUTDATED: 100000 individual blueberries (about 50 pounds) per 24 hours
	case LandResource_Potatoes:
		return 0
	case LandResource_Corn:
		return 0
	case LandResource_Agave:
		return 0
	case LandResource_Mushrooms:
		return 0
	case LandResource_Strawberries:
		return 15000 // 15000 individual strawberries (about 100 pounds) per 24 hours
	default:
		return 0
	}
}

func (resource _landResource) ToResource() _resource {
	switch resource.Type() {
	case LandResource_Unknown:
		return Resource(Resource_Unknown)
	case LandResource_Dirt:
		return Resource(Resource_Unknown)
	case LandResource_Pond:
		return Resource(Resource_Water)
	case LandResource_Lake_Vertical:
		return Resource(Resource_Water)
	case LandResource_Lake_Horizontal:
		return Resource(Resource_Water)
	case _landResource_Woods:
		return Resource_Logs(resource.Tree())
	case LandResource_Coal:
		return Resource(Resource_Coal)
	case LandResource_Clay:
		return Resource(Resource_Clay)
	case LandResource_Granite:
		return Resource(Resource_Granite)
	case LandResource_Limestone:
		return Resource(Resource_Limestone)
	case LandResource_Sandstone:
		return Resource(Resource_Sandstone)
	case LandResource_Marble:
		return Resource(Resource_Marble)
	case LandResource_Slate:
		return Resource(Resource_Slate)
	case LandResource_Iron:
		return Resource(Resource_Iron)
	case LandResource_Aluminum:
		return Resource(Resource_Aluminum)
	case LandResource_Zinc:
		return Resource(Resource_Zinc)
	case LandResource_Copper:
		return Resource(Resource_Copper)
	case LandResource_Nickel:
		return Resource(Resource_Nickel)
	case LandResource_Tin:
		return Resource(Resource_Tin)
	case LandResource_Silver:
		return Resource(Resource_Silver)
	case LandResource_Gold:
		return Resource(Resource_Gold)
	case LandResource_Haygrass:
		return Resource(Resource_Hay)
	case LandResource_RawRice:
		return Resource(Resource_RawRice)
	case LandResource_Berries:
		return Resource(Resource_Berries)
	case LandResource_Potatoes:
		return Resource(Resource_Potatoes)
	case LandResource_Corn:
		return Resource(Resource_Corn)
	case LandResource_Agave:
		return Resource(Resource_Agave)
	case LandResource_Mushrooms:
		return Resource(Resource_Mushrooms)
	case LandResource_Strawberries:
		return Resource(Resource_Strawberries)
	default:
		return Resource(Resource_Unknown)
	}
}

func (resource _landResource) ToString() string {
	switch resource.Type() {
	case LandResource_Unknown:
		return "Unknown"
	case LandResource_Dirt:
		return "Dirt"
	case LandResource_Pond:
		return "Pond"
	case LandResource_Lake_Vertical:
		return "Lake (Vertical)"
	case LandResource_Lake_Horizontal:
		return "Lake (Horizontal)"
	case _landResource_Woods:
		return resource.Tree().ToString() + " Woods"
	case LandResource_Coal:
		return "Coal"
	case LandResource_Clay:
		return "Clay"
	case LandResource_Granite:
		return "Granite"
	case LandResource_Limestone:
		return "Limestone"
	case LandResource_Sandstone:
		return "Sandstone"
	case LandResource_Marble:
		return "Marble"
	case LandResource_Slate:
		return "Slate"
	case LandResource_Iron:
		return "Iron"
	case LandResource_Aluminum:
		return "Aluminum"
	case LandResource_Zinc:
		return "Zinc"
	case LandResource_Copper:
		return "Copper"
	case LandResource_Nickel:
		return "Nickel"
	case LandResource_Tin:
		return "Tin"
	case LandResource_Silver:
		return "Silver"
	case LandResource_Gold:
		return "Gold"
	case LandResource_Haygrass:
		return "Haygrass"
	case LandResource_RawRice:
		return "Raw Rice"
	case LandResource_Berries:
		return "Berries"
	case LandResource_Potatoes:
		return "Potatoes"
	case LandResource_Corn:
		return "Corn"
	case LandResource_Agave:
		return "Agave"
	case LandResource_Mushrooms:
		return "Mushrooms"
	case LandResource_Strawberries:
		return "Strawberries"
	default:
		return "Unknown Resource"
	}
}

// These are the resources that have been harvested or crafted.
type _resource uint16

func Resource(Type ResourceType) _resource {
	return _resource(uint16(Type) | (uint16(TreeType_Unknown) << 8))
}

func Resource_Logs(Tree TreeType) _resource {
	return _resource(uint16(_landResource_Woods) | (uint16(Tree) << 8))
}

func (resource _resource) Type() ResourceType {
	return ResourceType(uint16(resource) & 255)
}

func (resource _resource) Tree() TreeType {
	return TreeType(uint16(resource) >> 8)
}

const Resource_Max = _resource((255 << 8) | 255)

type ResourceType uint8 // uint8 max is 255, uint16 max is 65535

const (
	// Basics
	Resource_Unknown ResourceType = iota
	Resource_Water
	Resource_ResearchPoints

	// Wood and Fuel
	_resource_Logs
	Resource_Coal

	// Stone
	Resource_Clay // Used for pottery and bricks

	Resource_Granite   // Most common stone type on the earth. Igneous rock composed of quartz, feldspar, and mica. Durable and strong. Often found in mountain ranges and the continental crust.
	Resource_Limestone // Common in sedimentary environments. Composed of calcium carbonate.
	Resource_Sandstone // Sedimentary rock composed mainly of sand-sized mineral particles or rock fragments. Commonly used in construction.
	Resource_Marble    // Metamorphic rock formed from limestone. Used in sculptures and high-end construction.
	Resource_Slate     // Fine-grained metamorphic rock that originates from Shale (heat and pressure over time turns Shale into Slate). Strong and durable, resistant to weathering and suitable for outdoor applications. Water-resistant.
	// Resource_Conglomerate // Sedimentary rock found in riverbeds, composed of rounded fragments of various sizes cemented together.
	// Resource_Quartzite // Hard metamorphic rock that originates from sandstone. Known for durability, used in construction, and as decoration.
	// Resource_Shale // Sedimentary rock formed from clay and silt. Used for bricks and tiles.
	// Resource_Basalt // Strong volcanic rock, used in construction and as aggregate in concrete.
	// Resrouce_Pumice // Lightweight volcanic rock
	// Resource_Flint
	// Resource_Soapstone
	// Resource_Quartz // Crystal-like stone

	// Naturally-ocurring Metals
	Resource_Iron     // Used in construction, manufacturing, and transportation, especially to make steel.
	Resource_Aluminum // Used in packaging, transportation, construction, and consumer goods.
	Resource_Zinc     // Used to galvanize steel to prevent corrosion, as well as in alloys and batteries.
	Resource_Copper   // Exclusively used in electrical wiring, plumbing, and electronics due to excellent conductivity.
	Resource_Nickel   // Used in stainless steel production, batteries (e.g., nickel-cadmium), and various alloys.
	Resource_Tin      // Combined with Copper to make bronze, coated on steel to prevent corrosion, soldering metal parts for electronics or plumbing, and glassmaking.
	Resource_Silver   // Used in electronics, jewelry, and photograpy.
	Resource_Gold     // Used in jewelry, electronics, and as an investment.

	// Resource_Titanium // Lightweight and strong. Used in video games for armor, high-end weapons, and equipment.
	// Resource_Obsidian // Volcanic glass

	// Fabrics and Wools
	Resource_Cloth
	Resource_SheepWool

	// Raw Plant Foods
	Resource_Hay
	Resource_RawRice
	Resource_Berries
	Resource_Potatoes
	Resource_Corn
	Resource_Agave
	Resource_Mushrooms
	Resource_Strawberries

	// Raw Meat Foods
	Resource_Milk
	Resource_Meat
	Resource_InsectMeat

	// --- Crafted Resources ---
	Resource_Steel  // Crafted from Iron and a small percentage of Carbon. Smelt pig iron combined with carbon (in the form of coke, a carbon derived fro coal) and other elements. Alloy with manganese, nickel, chromium, or vanadium, for different types of steel. Finally, refine and cast into various shapes (sheets, bars, and beams).
	Resource_Bronze // Alloy of Copper and Tin

	ResourceType_Max
)

func (resource _resource) ToString() string {
	switch resource.Type() {
	case Resource_Unknown:
		return "Unknown"
	case Resource_Water:
		return "Water"
	case Resource_ResearchPoints:
		return "Research Points"
	case _resource_Logs:
		return resource.Tree().ToString() + " Logs"
	case Resource_Coal:
		return "Coal"
	case Resource_Clay:
		return "Clay"
	case Resource_Granite:
		return "Granite"
	case Resource_Limestone:
		return "Limestone"
	case Resource_Sandstone:
		return "Sandstone"
	case Resource_Marble:
		return "Marble"
	case Resource_Slate:
		return "Slate"
	case Resource_Iron:
		return "Iron"
	case Resource_Aluminum:
		return "Aluminum"
	case Resource_Zinc:
		return "Zinc"
	case Resource_Copper:
		return "Copper"
	case Resource_Nickel:
		return "Nickel"
	case Resource_Tin:
		return "Tin"
	case Resource_Silver:
		return "Silver"
	case Resource_Gold:
		return "Gold"
	case Resource_Cloth:
		return "Cloth"
	case Resource_SheepWool:
		return "Sheep Wool"
	case Resource_Hay:
		return "Hay"
	case Resource_RawRice:
		return "Raw Rice"
	case Resource_Berries:
		return "Berries"
	case Resource_Potatoes:
		return "Potatoes"
	case Resource_Corn:
		return "Corn"
	case Resource_Agave:
		return "Agave"
	case Resource_Mushrooms:
		return "Mushrooms"
	case Resource_Strawberries:
		return "Strawberries"
	case Resource_Milk:
		return "Milk"
	case Resource_Meat:
		return "Meat"
	case Resource_InsectMeat:
		return "Insect Meat"
	case Resource_Steel:
		return "Steel"
	case Resource_Bronze:
		return "Bronze"
	default:
		return "Unknown Resource"
	}
}

// Note: Some foods when eaten raw can give food poisoning

// There are 73 tree types.
type TreeType uint8

const (
	TreeType_Unknown TreeType = iota
	TreeType_Mahogany
	TreeType_Kapok_Tree
	TreeType_Brazil_Nut_Tree
	TreeType_Rubber_Tree
	TreeType_Strangler_Fig
	TreeType_Teak
	TreeType_Sal_Tree
	TreeType_Indian_Rosewood
	TreeType_Flame_Tree
	TreeType_Tamarind
	TreeType_Podocarpus_Trees
	TreeType_Alder
	TreeType_Tree_Ferns
	TreeType_Magnolia
	TreeType_Bamboo
	TreeType_Ironwood
	TreeType_Ebony
	TreeType_Meranti
	TreeType_Rosewood
	TreeType_Nutmeg_Tree
	TreeType_Mango_Tree
	TreeType_Banyan_Tree
	TreeType_Jackfruit_Tree
	TreeType_Sandalwood
	TreeType_Dipterocarp_Trees
	TreeType_Acacia
	TreeType_Baobab
	TreeType_Marula_Tree
	TreeType_Sausage_Tree
	TreeType_Terminalia
	TreeType_Mangrove_Palm
	TreeType_Water_Tupelo
	TreeType_Swamp_Mahogany
	TreeType_Pond_Cypress
	TreeType_Melaleuca
	TreeType_Rattan_Palm
	TreeType_Screw_Pine
	TreeType_Water_Hickory
	TreeType_Oil_Palm
	TreeType_Red_Mangrove
	TreeType_Black_Mangrove
	TreeType_White_Mangrove
	TreeType_Buttonwood
	TreeType_Sea_Hibiscus
	TreeType_Oak // incl. Swamp Chestnut Oak, Cork Oak, White&Red Oak
	TreeType_Maple
	TreeType_Beech
	TreeType_Birch
	TreeType_Hickory
	TreeType_Eastern_Hemlock
	TreeType_Douglas_Fir
	TreeType_Red_Maple
	TreeType_White_Pine
	TreeType_Chestnut
	TreeType_Sitka_Spruce
	TreeType_Western_Red_Cedar
	TreeType_Bigleaf_Maple
	TreeType_Coast_Redwood
	TreeType_Yellow_Cedar
	TreeType_Scots_Pine
	TreeType_Norway_Spruce
	TreeType_Lodgepole_Pine
	TreeType_Bald_Cypress
	TreeType_Black_Gum
	TreeType_Sweetbay_Magnolia
	TreeType_Willow
	TreeType_Gumbo_Limbo
	TreeType_Ombú_Tree
	TreeType_Wild_Olive
	TreeType_Cottonwood
	TreeType_American_Elm
	TreeType_Box_Elder
	TreeType_Tamarack
	TreeType_Black_Spruce
	TreeType_Red_Cedar
	TreeType_Krummholz_Pines
	TreeType_Mountain_Hemlock
	TreeType_Juniper
	TreeType_Dwarf_Willow
	TreeType_Bristlecone_Pine
	TreeType_Dwarf_Birch
	TreeType_Arctic_Willow
	TreeType_Siberian_Elm
	TreeType_Pinyon_Pine
	TreeType_Olive_Trees
	TreeType_Carob_Tree
	TreeType_Aleppo_Pine
	TreeType_Protea_Trees
	TreeType_Silver_Tree
	TreeType_Mesquite
	TreeType_Palo_Verde
	TreeType_Date_Palm
	TreeType_Tamarisk
	TreeType_Desert_Willow

	TreeType_Max
)

func (treeType TreeType) ToString() string {
	switch treeType {
	case TreeType_Unknown:
		return "Unknown"
	case TreeType_Mahogany:
		return "Mahogany"
	case TreeType_Kapok_Tree:
		return "Kapok Tree"
	case TreeType_Brazil_Nut_Tree:
		return "Brazil Nut Tree"
	case TreeType_Rubber_Tree:
		return "Rubber Tree"
	case TreeType_Strangler_Fig:
		return "Strangler Fig"
	case TreeType_Teak:
		return "Teak"
	case TreeType_Sal_Tree:
		return "Sal Tree"
	case TreeType_Indian_Rosewood:
		return "Indian Rosewood"
	case TreeType_Flame_Tree:
		return "Flame Tree"
	case TreeType_Tamarind:
		return "Tamarind"
	case TreeType_Podocarpus_Trees:
		return "Podocarpus Trees"
	case TreeType_Alder:
		return "Alder"
	case TreeType_Tree_Ferns:
		return "Tree Ferns"
	case TreeType_Magnolia:
		return "Magnolia"
	case TreeType_Bamboo:
		return "Bamboo"
	case TreeType_Ironwood:
		return "Ironwood"
	case TreeType_Ebony:
		return "Ebony"
	case TreeType_Meranti:
		return "Meranti"
	case TreeType_Rosewood:
		return "Rosewood"
	case TreeType_Nutmeg_Tree:
		return "Nutmeg Tree"
	case TreeType_Mango_Tree:
		return "Mango Tree"
	case TreeType_Banyan_Tree:
		return "Banyan Tree"
	case TreeType_Jackfruit_Tree:
		return "Jackfruit Tree"
	case TreeType_Sandalwood:
		return "Sandalwood"
	case TreeType_Dipterocarp_Trees:
		return "Dipterocarp Trees"
	case TreeType_Acacia:
		return "Acacia"
	case TreeType_Baobab:
		return "Baobab"
	case TreeType_Marula_Tree:
		return "Marula Tree"
	case TreeType_Sausage_Tree:
		return "Sausage Tree"
	case TreeType_Terminalia:
		return "Terminalia"
	case TreeType_Mangrove_Palm:
		return "Mangrove Palm"
	case TreeType_Water_Tupelo:
		return "Water Tupelo"
	case TreeType_Swamp_Mahogany:
		return "Swamp Mahogany"
	case TreeType_Pond_Cypress:
		return "Pond Cypress"
	case TreeType_Melaleuca:
		return "Melaleuca"
	case TreeType_Rattan_Palm:
		return "Rattan Palm"
	case TreeType_Screw_Pine:
		return "Screw Pine"
	case TreeType_Water_Hickory:
		return "Water Hickory"
	case TreeType_Oil_Palm:
		return "Oil Palm"
	case TreeType_Red_Mangrove:
		return "Red Mangrove"
	case TreeType_Black_Mangrove:
		return "Black Mangrove"
	case TreeType_White_Mangrove:
		return "White Mangrove"
	case TreeType_Buttonwood:
		return "Buttonwood"
	case TreeType_Sea_Hibiscus:
		return "Sea Hibiscus"
	case TreeType_Oak:
		return "Oak"
	case TreeType_Maple:
		return "Maple"
	case TreeType_Beech:
		return "Beech"
	case TreeType_Birch:
		return "Birch"
	case TreeType_Hickory:
		return "Hickory"
	case TreeType_Eastern_Hemlock:
		return "Eastern Hemlock"
	case TreeType_Douglas_Fir:
		return "Douglas Fir"
	case TreeType_Red_Maple:
		return "Red Maple"
	case TreeType_White_Pine:
		return "White Pine"
	case TreeType_Chestnut:
		return "Chestnut"
	case TreeType_Sitka_Spruce:
		return "Sitka Spruce"
	case TreeType_Western_Red_Cedar:
		return "Western Red Cedar"
	case TreeType_Bigleaf_Maple:
		return "Bigleaf Maple"
	case TreeType_Coast_Redwood:
		return "Coast Redwood"
	case TreeType_Yellow_Cedar:
		return "Yellow Cedar"
	case TreeType_Scots_Pine:
		return "Scots Pine"
	case TreeType_Norway_Spruce:
		return "Norway Spruce"
	case TreeType_Lodgepole_Pine:
		return "Lodgepole Pine"
	case TreeType_Bald_Cypress:
		return "Bald Cypress"
	case TreeType_Black_Gum:
		return "Black Gum"
	case TreeType_Sweetbay_Magnolia:
		return "Sweetbay Magnolia"
	case TreeType_Willow:
		return "Willow"
	case TreeType_Gumbo_Limbo:
		return "Gumbo Limbo"
	case TreeType_Ombú_Tree:
		return "Ombú Tree"
	case TreeType_Wild_Olive:
		return "Wild Olive"
	case TreeType_Cottonwood:
		return "Cottonwood"
	case TreeType_American_Elm:
		return "American Elm"
	case TreeType_Box_Elder:
		return "Box Elder"
	case TreeType_Tamarack:
		return "Tamarack"
	case TreeType_Black_Spruce:
		return "Black Spruce"
	case TreeType_Red_Cedar:
		return "Red Cedar"
	case TreeType_Krummholz_Pines:
		return "Krummholz Pines"
	case TreeType_Mountain_Hemlock:
		return "Mountain Hemlock"
	case TreeType_Juniper:
		return "Juniper"
	case TreeType_Dwarf_Willow:
		return "Dwarf Willow"
	case TreeType_Bristlecone_Pine:
		return "Bristlecone Pine"
	case TreeType_Dwarf_Birch:
		return "Dwarf Birch"
	case TreeType_Arctic_Willow:
		return "Arctic Willow"
	case TreeType_Siberian_Elm:
		return "Siberian Elm"
	case TreeType_Pinyon_Pine:
		return "Pinyon Pine"
	case TreeType_Olive_Trees:
		return "Olive Trees"
	case TreeType_Carob_Tree:
		return "Carob Tree"
	case TreeType_Aleppo_Pine:
		return "Aleppo Pine"
	case TreeType_Protea_Trees:
		return "Protea Trees"
	case TreeType_Silver_Tree:
		return "Silver Tree"
	case TreeType_Mesquite:
		return "Mesquite"
	case TreeType_Palo_Verde:
		return "Palo Verde"
	case TreeType_Date_Palm:
		return "Date Palm"
	case TreeType_Tamarisk:
		return "Tamarisk"
	case TreeType_Desert_Willow:
		return "Desert Willow"
	default:
		return "Unknown Tree"
	}
}

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

func (zone *ResourceZone) AddWorker(id AgentId, a *Agent) {
	// If already assigned a zone, return
	if a.assignedZone != -1 {
		return
	}

	a.assignedZone = zone.id
	zone.workers = append(zone.workers, id)
}

func (zone *ResourceZone) RemoveWorker(id AgentId, a *Agent) {
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

func (zone *ResourceZone) RemoveLastWorker(colony *Colony) {
	lastWorkerId := zone.workers[len(zone.workers)-1]
	colony.agents[lastWorkerId].assignedZone = -1
	zone.workers = zone.workers[:len(zone.workers)-1]
}
