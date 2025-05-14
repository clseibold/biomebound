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
		return 100 // TODO
	case LandResource_Clay:
		return 0
	case LandResource_Granite:
		return 100 // TODO
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

func (resource _landResource) GetDescription() string {
	// Handle woods separately since they use tree types
	if resource.Type() == _landResource_Woods {
		return GetTreeDescription(resource.Tree())
	}

	switch resource.Type() {
	case LandResource_Unknown:
		return "An unknown resource lies here."

	case LandResource_Dirt:
		return "Rich earth suitable for basic construction and farming. The soil here varies in texture and composition, showing layers of organic matter mixed with minerals built up over countless seasons."

	case LandResource_Pond:
		return "A small body of still water reflects the sky like a mirror. The pond's edges are ringed with vegetation, and the occasional ripple betrays the presence of aquatic life beneath the surface."

	case LandResource_Lake_Vertical:
		return "A long, deep lake stretches north to south, its waters dark with depth. The steep banks suggest this lake may have been carved by ancient glaciers, leaving behind this reservoir of fresh water."

	case LandResource_Lake_Horizontal:
		return "A wide lake extends from east to west, its surface catching the light. The gradually sloping shores provide excellent access to the water, while wetland plants thrive in the shallows."

	case LandResource_Coal:
		return "Dark seams of coal run through the earth here, the preserved remains of ancient forests compressed over millions of years. This valuable fuel source was formed in prehistoric swamps and peat bogs."

	case LandResource_Clay:
		return "Beds of fine clay lie near the surface, their smooth texture perfect for pottery and brickmaking. The clay ranges in color from rich red to pale gray, depending on its mineral content."

	case LandResource_Granite:
		return "Massive granite formations rise from the earth, their crystalline structure sparkling in the light. This ancient igneous rock, formed deep within the earth, provides excellent building material."

	case LandResource_Limestone:
		return "Pale limestone deposits are exposed here, formed from countless ancient marine creatures. The rock is soft enough to work easily but durable enough to last centuries, making it perfect for construction."

	case LandResource_Sandstone:
		return "Layers of sandstone show beautiful ripple patterns from their formation in ancient seas. The stone ranges from pale yellow to deep red, each layer telling a story of different geological eras."

	case LandResource_Marble:
		return "Veins of marble run through the rock here, their crystalline structure refracting light. This metamorphic stone, transformed by heat and pressure, is prized for its beauty and durability."

	case LandResource_Slate:
		return "Sheets of slate cleave cleanly along natural planes, perfect for roofing and paving. This fine-grained stone, born of compressed clay, resists water and weathering admirably."

	case LandResource_Iron:
		return "Rich veins of iron ore streak through the rock, their rusty red color betraying their presence. This essential metal, trapped in mineral form, awaits extraction and smelting."

	case LandResource_Aluminum:
		return "Deposits of bauxite, the primary ore of aluminum, color the soil a distinctive reddish-brown. This lightweight metal's ore requires significant processing to yield its treasure."

	case LandResource_Zinc:
		return "Silvery gray zinc ore shows through the rock face, often mixed with other minerals. This essential metal, crucial for protecting iron and steel, lies waiting to be extracted."

	case LandResource_Copper:
		return "Distinctive green-blue copper deposits stain the surrounding rock, marking rich veins of this conductive metal. Ancient weathering has exposed these valuable ore deposits."

	case LandResource_Nickel:
		return "Streaks of nickel ore run through the rock, their presence marked by subtle color variations. This hardy metal, often found with iron, promises strong alloys when refined."

	case LandResource_Tin:
		return "Cassiterite, the primary ore of tin, shows as dark crystals in the surrounding rock. This essential component of bronze has drawn miners since ancient times."

	case LandResource_Silver:
		return "Glinting veins of silver ore thread through the darker rock, promising precious metal within. The ore typically contains lead and copper as well, requiring careful separation."

	case LandResource_Gold:
		return "Traces of gold glitter in the rock and soil, the precious metal trapped in quartz veins or alluvial deposits. These deposits have drawn prospectors throughout history."

	case LandResource_Haygrass:
		return "Thick stands of haygrass wave in the breeze, ready for harvesting. This versatile grass provides essential feed for livestock and can be stored for long periods."

	case LandResource_RawRice:
		return "Paddies of rice plants stand in shallow water, their green stems swaying gently. The grain-heavy heads bow with the weight of their precious cargo."

	case LandResource_Berries:
		return "Wild berry bushes grow in abundance here, their fruits ranging from deep purple to bright red. These natural sweet treats provide essential nutrients and can be preserved for winter."

	case LandResource_Potatoes:
		return "Hardy potato plants flourish in the soil, their valuable tubers growing unseen below. These reliable crops provide excellent nutrition and store well for long periods."

	case LandResource_Corn:
		return "Tall corn stalks reach for the sky, their leaves rustling in the wind. The golden ears of corn, wrapped in protective husks, promise a bountiful harvest."

	case LandResource_Agave:
		return "Spiky agave plants spread their thick, fleshy leaves in a rosette pattern. These desert-adapted plants store precious water and provide useful fibers."

	case LandResource_Mushrooms:
		return "Colonies of mushrooms sprout from the dark, rich soil, their caps ranging from tiny buttons to broad plates. These fungi thrive in the moist, shaded conditions."

	case LandResource_Strawberries:
		return "Low-growing strawberry plants carpet the ground, their white flowers and red fruits dotting the green leaves. The sweet berries are a valuable source of nutrition."

	default:
		return "A resource deposit of unknown type exists here."
	}
}

func (resource _landResource) IsWoodsType() bool {
	return resource.Type() == _landResource_Woods
}

// These are the resources that have been harvested or crafted.
type _resource uint16

func Resource(Type ResourceType) _resource {
	return _resource(uint16(Type) | (uint16(TreeType_Unknown) << 8))
}

func Resource_Logs(Tree TreeType) _resource {
	return _resource(uint16(_resource_Logs) | (uint16(Tree) << 8))
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

func (resource _resource) IsTreeType() bool {
	return resource.Type() == _resource_Logs
}

// Note: Some foods when eaten raw can give food poisoning

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
