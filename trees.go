package biomebound

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

// GetTreeDescription returns a flavor text description of the forest/woods for each tree type
func GetTreeDescription(treeType TreeType) string {
	switch treeType {
	case TreeType_Mahogany:
		return "Towering mahogany trees dominate this area, their massive buttressed trunks rising like ancient pillars. The canopy above is a dense tapestry of broad leaves, and the bark of these prized hardwoods shows distinctive gray crackled patterns. The wood here is famously strong and beautiful, sought after by craftsmen worldwide."

	case TreeType_Kapok_Tree:
		return "Giant kapok trees pierce the canopy, their enormous trunks supported by massive buttress roots. Young trees bristle with sharp spines, while mature specimens tower 50-60 meters overhead. The air is often filled with floating kapok fiber from their seed pods, giving the forest an ethereal quality."

	case TreeType_Brazil_Nut_Tree:
		return "These majestic giants of the Amazon rise from the forest floor with straight, massive trunks. The ground is littered with their large, heavy fruit capsules - each containing precious Brazil nuts. The trees form a towering canopy, their thick, oil-rich wood having stood the test of time."

	case TreeType_Rubber_Tree:
		return "A well-ordered stand of rubber trees stretches before you, their smooth gray bark marked by diagonal tapping cuts. These commercially valuable trees drop their leaves seasonally, and their white flowers give way to tripart seed pods. The wood is surprisingly hard for such a utilitarian species."

	case TreeType_Strangler_Fig:
		return "An eerie forest of strangler figs dominates the area, their complex network of aerial roots having long since enveloped their host trees. White latex oozes from any bark damage, and the ground is littered with small, bird-favored figs. The wood is soft but the trees' structure is impressively complex."

	case TreeType_Teak:
		return "Noble teak trees stand tall here, their large leaves casting dappled shade. The straight, thick trunks wear flaking bark, and the wood beneath is legendary for its durability. Small white flowers bloom in terminal clusters, a delicate contrast to the robust timber that makes this wood so prized."

	case TreeType_Podocarpus_Trees:
		return "These elegant conifers rise with dense, dark-green foliage. Their leathery leaves cast deep shade, while small fleshy cones add spots of color. The wood is prized for its fine, even grain - a favorite of master craftsmen in tropical highlands."

	case TreeType_Alder:
		return "A grove of alders stands near the water, their smooth gray bark marked with horizontal lenticels. The wood is pale and soft when fresh-cut, but takes on a rich reddish hue when exposed to air. Though not the most durable, it has traditionally been prized for its water-resistance."

	case TreeType_Tree_Ferns:
		return "Ancient-looking tree ferns spread their dramatic fronds overhead, creating a prehistoric atmosphere. Their fibrous trunks are actually made of tightly packed roots, rising like natural columns. While not true wood, the trunks have been used traditionally for walls and garden elements."

	case TreeType_Magnolia:
		return "These stately trees feature thick, glossy leaves and massive fragrant flowers. The wood is fine-grained and moderately hard, historically favored for furniture and interior finishing. The pale cream-colored timber occasionally shows subtle hints of green or brown."

	case TreeType_Bamboo:
		return "Towering culms of bamboo rise in dense stands, their hollow stems creating a distinctive rustle in the breeze. Though technically a grass, the woody stems are remarkably strong and versatile, used worldwide for construction and crafting. The rapid growth can create groves of impressive height in just a few years."

	case TreeType_Ironwood:
		return "These dense-wooded trees stand like iron sentinels, their incredibly hard timber a testament to their name. The canopy is relatively sparse, with small, tough leaves that rustle in the breeze. The wood is so heavy it sinks in water, and tools often need resharpening when working it."

	case TreeType_Ebony:
		return "A grove of ebony trees rises before you, their dark bark hiding even darker heartwood beneath. The dense canopy of glossy leaves creates deep shade. This wood is among the heaviest and most precious in the world, prized for its striking black color and mirror-like finish when polished."

	case TreeType_Meranti:
		return "Towering meranti trees dominate this area, their massive, straight trunks supporting a high canopy. The reddish-brown wood is highly valued, and resin seeps from any wounds in the bark. These giants can reach impressive heights, their buttress roots sprawling across the forest floor."

	case TreeType_Rosewood:
		return "These valuable trees stand gracefully, their finely-grained wood hiding subtle hints of rose-like fragrance when cut. The canopy is moderately dense, and fallen leaves release a sweet scent. The timber is famous for its rich, dark colors and musical resonance."

	case TreeType_Nutmeg_Tree:
		return "An aromatic grove of nutmeg trees spreads before you, their branches laden with yellow fruits that split to reveal the precious spice within. The wood is pale and straight-grained, while the canopy creates a dense shade. A sweet, spicy scent hangs in the air."

	case TreeType_Mango_Tree:
		return "Wide-spreading mango trees create a generous canopy, their dark leaves contrasting with colorful fruits when in season. The wood is coarse-grained and sturdy, taking on a grayish tinge with age. These trees provide both valuable timber and delicious fruit."

	case TreeType_Banyan_Tree:
		return "Massive banyan trees spread their aerial roots like curtains, creating a forest within a forest. The main trunks are enormous, while hundreds of prop roots descend from the branches to form new supports. The wood is soft but flexible, and the shade beneath is deep and cool."

	case TreeType_Jackfruit_Tree:
		return "Sturdy jackfruit trees rise with dense, dark green canopies, their massive fruits growing directly from the trunk and main branches. The wood has a beautiful golden color when polished, resistant to termites and decay. The air is sweet with the scent of ripe fruit."

	case TreeType_Sandalwood:
		return "These precious trees grow in scattered stands, their unremarkable appearance hiding incredibly fragrant heartwood. The wood is fine-grained and heavy, releasing its distinctive sweet, warm scent when cut. The canopy is relatively sparse, with leathery oval leaves."

	case TreeType_Dipterocarp_Trees:
		return "Majestic dipterocarp trees soar overhead, their massive trunks rising clean and straight. Their distinctive two-winged fruits spiral down like helicopters in season. The wood is reddish-brown and highly valued, while their great height creates a towering, cathedral-like forest."

	case TreeType_Acacia:
		return "Thorny acacia trees spread their flat-topped canopies across the landscape, their small leaves creating dappled shade patterns. The dark, dense heartwood is renowned for its durability and rich color. These hardy trees bear curved thorns and fragrant flowers that attract countless insects."

	case TreeType_Baobab:
		return "Ancient baobab trees stand like giant sentinels, their massive swollen trunks storing thousands of liters of water. The branches, bare much of the year, spread like roots into the sky. The fibrous wood is soft and wet, making these trees living reservoirs of life-giving moisture."

	case TreeType_Marula_Tree:
		return "Spreading marula trees dot the landscape, their round crowns providing welcome shade. The pale, fine-grained wood surrounds a soft core, while the bark is distinctively mottled. Sweet fruits attract wildlife when in season, and the nuts contain precious oils."

	case TreeType_Sausage_Tree:
		return "Strange sausage trees stand with their distinctive fruits hanging like giant sausages from long rope-like stalks. The wood is soft but durable, and the dark bark is deeply fissured. Purple-red flowers bloom at night, attracting bats and nocturnal moths."

	case TreeType_Terminalia:
		return "Elegant terminalia trees rise in graceful tiers, their branches arranged in distinct horizontal layers. The wood is yellowish-brown and moderately hard, while the bark contains valuable tannins. Their tiered shape provides excellent shade and shelter."

	case TreeType_Mangrove_Palm:
		return "These specialized palms thrive in the brackish water, their stilt roots providing stability in the soft mud. The wood is dense and water-resistant, while the fronds arch gracefully overhead. They form an important part of the coastal ecosystem."

	case TreeType_Water_Tupelo:
		return "Stately water tupelos rise from the swamp, their swollen bases adapted to frequent flooding. The wood is light but tough, traditionally used for bee hives. In autumn, their leaves turn brilliant scarlet before falling into the dark water below."

	case TreeType_Swamp_Mahogany:
		return "These water-loving trees stand tall in the wetlands, their reddish wood protected by thick, spongy bark. The canopy is dense and glossy, providing year-round shade. The wood is highly valued for its water resistance and rich color."

	case TreeType_Pond_Cypress:
		return "Graceful pond cypresses rise from the water, their knees breaking the surface like wooden stalagmites. The reddish-brown wood is highly decay-resistant, while the feathery foliage creates delicate patterns against the sky. The bases of the trunks flare out dramatically where they meet the water."

	case TreeType_Melaleuca:
		return "Paperbark melaleuca trees stand with their distinctive white, papery bark peeling in thick layers. The wood is hard and heavy, naturally oily and water-resistant. These trees release a strong, medicinal scent when their leaves are crushed."

	case TreeType_Rattan_Palm:
		return "Slender rattan palms climb through the canopy, their flexible stems reaching impressive lengths. The wood is actually a vine-like stem, incredibly strong yet pliable, perfect for weaving furniture. Sharp spines along the stems and whip-like tendrils help these palms climb skyward."

	case TreeType_Screw_Pine:
		return "Peculiar screw pines stand on their stilt roots like many-legged sentinels. Their spiral arrangement of sword-like leaves gives them their name. The fibrous wood and leaves have traditional uses in weaving and construction, while their pineapple-like fruits are sometimes eaten."

	case TreeType_Water_Hickory:
		return "These flood-tolerant hickories thrive in wet soils, their strong wood protected by thick, plated bark. The canopy is relatively open, allowing dappled sunlight to reach the water below. Their sweet nuts are prized by wildlife, and the tough wood is excellent for tool handles."

	case TreeType_Oil_Palm:
		return "Rows of oil palms stand like columns, their feathery fronds forming a dense crown. The thick trunks are marked by old leaf bases, while heavy clusters of oil-rich fruits hang below the canopy. Though not true wood, their trunks provide useful building material."

	case TreeType_Red_Mangrove:
		return "Red mangroves create an impenetrable maze of arching prop roots along the coastline. Their reddish wood is dense and rot-resistant, while the complex root system provides shelter for countless marine creatures. Young trees drop like spears from the parent before taking root."

	case TreeType_Black_Mangrove:
		return "Black mangroves line the shore, their dark trunks rising from beds of pneumatophores - finger-like breathing roots that pierce the mud. The wood is heavy and hard, excellent for fuel. Salt often crystalizes on the thick, leathery leaves that handle brackish waters with ease."

	case TreeType_White_Mangrove:
		return "These hardy mangroves form the landward edge of the swamp, their pale bark contrasting with darker species. The wood is lighter than other mangroves but still remarkably tough. Their leaves are slightly succulent, helping them cope with salty conditions."

	case TreeType_Buttonwood:
		return "Gnarled buttonwood trees guard the transition between mangrove and upland, their silvery bark reflecting the light. The dense, fine-grained wood takes a beautiful polish, while the button-like fruit clusters give these trees their name. Salt often encrusts their thick leaves."

	case TreeType_Sea_Hibiscus:
		return "Graceful sea hibiscus trees line the coast, their heart-shaped leaves rustling in the sea breeze. Large yellow flowers fade to deep red as they age. The light, soft wood is easily worked, while the fibrous bark has traditionally been used for cordage."

	case TreeType_Oak:
		return "Mighty oaks spread their broad crowns wide, their sturdy branches wearing a thick mantle of leaves. The legendary strength of their wood is matched by its beauty when worked. Acorns litter the ground beneath, supporting diverse wildlife in these ancient guardians' shade."

	case TreeType_Maple:
		return "These handsome maples form a dense canopy of distinctive lobed leaves that turn brilliant colors in autumn. The wood is prized for its beautiful grain patterns and moderate hardness, especially the highly figured 'bird's eye' variety. Sweet sap flows beneath the smooth, gray bark in early spring."

	case TreeType_Beech:
		return "Stately beech trees create a cathedral-like canopy, their smooth, silver-gray bark unmarred by age. The dense wood is pale and fine-grained, while the thin leaves let dappled sunlight create shifting patterns below. Triangular nuts crunch underfoot in autumn."

	case TreeType_Birch:
		return "A grove of birch trees stands with their distinctive white bark peeling in papery layers. The wood is pale and fine-grained, traditionally used for everything from paper to canoes. Their small leaves dance in the slightest breeze, creating a shimmering effect in sunlight."

	case TreeType_Hickory:
		return "Strong hickory trees rise straight and tall, their deeply furrowed bark protecting some of the toughest wood in the forest. The canopy is open but full, and fallen nuts attract wildlife. The wood is famous for its strength and shock resistance, perfect for tool handles."

	case TreeType_Eastern_Hemlock:
		return "Dark hemlock trees create deep shade with their drooping branches and fine needles. The bark is rich in tannins, while the wood, though soft, is durable when kept dry. These long-lived trees often shelter delicate understory plants in their perpetual shade."

	case TreeType_Douglas_Fir:
		return "Towering Douglas firs pierce the sky, their straight trunks hosting thick, deeply furrowed bark. The wood is strong and relatively light, making it perfect for construction. Their distinctive cones with mouse-tail bracts scatter the forest floor."

	case TreeType_Red_Maple:
		return "Red maples stand proud with their smooth gray bark and distinctive leaves that turn blazing scarlet in autumn. The wood is softer than sugar maple but still valuable, with occasional decorative grain patterns. Their early spring flowers provide crucial nectar for awakening insects."

	case TreeType_White_Pine:
		return "Majestic white pines tower overhead, their soft blue-green needles in bundles of five. The wood is light, straight-grained, and easily worked, historically prized for ship masts. Their long horizontal branches create distinct layers in the canopy."

	case TreeType_Chestnut:
		return "These mighty chestnuts wear deeply furrowed bark protecting their rot-resistant wood. Though rare now in many places, their legacy lives on in weathered fence posts and barn beams. The sweet nuts were once a staple food for both wildlife and humans."

	case TreeType_Sitka_Spruce:
		return "Giant Sitka spruce trees dominate the coastal forest, their sharp needles and scaley bark glistening with sea mist. The wood is strong and light, resonating beautifully in musical instruments. Their towering forms create cathedral-like spaces beneath."

	case TreeType_Western_Red_Cedar:
		return "Magnificent western red cedars rise from the misty forest floor, their stringy bark hanging in long strips. The aromatic wood is legendarily rot-resistant, prized by indigenous peoples for centuries. Their drooping branches form graceful layers draped with scale-like foliage."

	case TreeType_Bigleaf_Maple:
		return "Massive bigleaf maples spread their enormous leaves overhead, often draped with curtains of moss and ferns. The wood features beautiful ripple patterns prized by craftsmen. In autumn, their giant golden leaves carpet the forest floor."

	case TreeType_Coast_Redwood:
		return "Ancient coast redwoods soar to staggering heights, their reddish bark soft and fibrous. These giants can live for thousands of years, their uppermost branches lost in the fog. The wood is naturally resistant to decay and surprisingly light for its strength."

	case TreeType_Yellow_Cedar:
		return "Yellow cedars stand straight and tall, their scaly foliage releasing a sharp, spicy scent when crushed. The wood is remarkably dense and durable for a cedar, with a fine, uniform texture. Their slow growth produces exceptionally tight grain patterns."

	case TreeType_Scots_Pine:
		return "Hardy Scots pines show their distinctive orange-red upper bark against the sky. The wood is strong and resinous, while the twisted blue-green needles create an open, airy canopy. These adaptable trees thrive in poor soils where others struggle."

	case TreeType_Norway_Spruce:
		return "Dark Norway spruce trees form perfect pyramids, their drooping branches sweeping the ground. The wood is light but strong, historically prized for violin making. Their long cones hang like decorative ornaments from the tips of branches."

	case TreeType_Lodgepole_Pine:
		return "Straight lodgepole pines stand in dense groves, their slender trunks reaching uniformly skyward. The wood is light and straight-grained, traditionally used for tipi poles. Their sealed cones often wait years for fire's heat to release their seeds."

	case TreeType_Bald_Cypress:
		return "Ancient bald cypresses rise from the swamp, their distinctive 'knees' breaking the water's surface. The reddish, rot-resistant wood has survived centuries of flooding. In autumn, their feathery needles turn copper before falling, unlike most conifers."

	case TreeType_Black_Gum:
		return "Black gum trees stand with their distinctively horizontal branches creating perfect layers. The wood is extremely tough and resistant to splitting, while the glossy leaves turn brilliant scarlet in autumn. Their small blue fruits are favored by wildlife."

	case TreeType_Sweetbay_Magnolia:
		return "Elegant sweetbay magnolias grace the wetlands, their silvery-backed leaves flashing in the breeze. The creamy white flowers release their lemony fragrance throughout summer. The pale wood is soft but durable, taking on a silky luster when polished."

	case TreeType_Willow:
		return "Graceful willows droop their long branches toward the water, creating curtains of slender leaves. The wood is soft and lightweight, traditionally used for cricket bats and prosthetics. Their extensive root systems help stabilize stream banks and prevent erosion."

	case TreeType_Gumbo_Limbo:
		return "Distinctive gumbo limbo trees stand out with their peeling, copper-red bark that earned them the nickname 'tourist tree'. The wood is soft and light but resistant to tropical storms. Their thick, resinous sap has traditional medicinal uses."

	case TreeType_Ombú_Tree:
		return "Massive ombú trees spread their umbrella-like canopies across the grassland, their swollen trunks storing water for dry seasons. Though technically not true wood, their fibrous trunk material has been used traditionally. These living landmarks provide rare shade in the pampas."

	case TreeType_Wild_Olive:
		return "Gnarled wild olive trees twist their ancient trunks skyward, their small silvery leaves shimmering in the breeze. The wood is incredibly dense and beautifully figured, prized for small decorative items. Their small fruits attract birds throughout the season."

	case TreeType_Cottonwood:
		return "Tall cottonwoods line the riverbanks, their broad leaves creating a constant whisper in the breeze. The soft, light wood was traditionally used for dugout canoes. In early summer, their seeds drift through the air like summer snow."

	case TreeType_American_Elm:
		return "Stately American elms arch their branches in classic vase-shapes, their serrated leaves creating dense shade. The wood is hard and cross-grained, resistant to splitting. Once common street trees, these survivors show remarkable resilience."

	case TreeType_Box_Elder:
		return "Fast-growing box elders spread their irregular crowns, their compound leaves bringing variety to the forest edge. The soft, light wood occasionally shows attractive patterns called 'flame' grain. Their winged seeds spin down in abundance each autumn."

	case TreeType_Tamarack:
		return "Delicate tamaracks stand tall in the northern forest, their soft needles turning brilliant gold before falling in autumn. Unlike most conifers, they shed their needles each year. The wood is tough, rot-resistant, and traditionally used in boat-building."

	case TreeType_Black_Spruce:
		return "Hardy black spruce trees form dense stands in the boreal forest, their narrow spires adapted to heavy snow. The wood is light but strong, prized for paper-making. Their shallow root systems create characteristic mounded growth in boggy areas."

	case TreeType_Red_Cedar:
		return "Aromatic red cedars rise from rocky soils, their reddish, stringy bark protecting fragrant wood. The dense, rot-resistant heartwood has long been used for chest-making and closet linings. Their blue, berry-like cones are favored by wildlife."

	case TreeType_Krummholz_Pines:
		return "Windswept krummholz pines grow twisted and stunted by harsh mountain conditions, their branches growing mainly on the sheltered side. These hardy survivors rarely grow taller than a person, their dense, twisted wood telling stories of centuries of alpine endurance."

	case TreeType_Mountain_Hemlock:
		return "Graceful mountain hemlocks dot the subalpine slopes, their drooping branches often bent by snow load. The wood is moderately hard and resistant to decay at high altitudes. Their delicate needles create a feathery silhouette against mountain skies."

	case TreeType_Juniper:
		return "Ancient junipers twist their weathered forms across exposed slopes, their shredding bark and blue berries lasting through winter. The aromatic, rot-resistant wood has been prized for centuries. These slow-growing survivors can live for thousands of years in harsh conditions."

	case TreeType_Dwarf_Willow:
		return "Tiny dwarf willows hug the ground in arctic and alpine areas, their entire height rarely exceeding ankle level. Though small, their wood is surprisingly tough and flexible. These diminutive trees form dense mats that help stabilize the soil in extreme environments."

	case TreeType_Bristlecone_Pine:
		return "Ancient bristlecone pines stand like living sculptures, their weathered wood polished by millennia of wind and ice. These are among the oldest living things on Earth, their tough, resinous wood preserving their forms long after death. Their needles can persist for forty years."

	case TreeType_Dwarf_Birch:
		return "Low-growing dwarf birch forms thickets in arctic and alpine regions, their small leaves turning brilliant gold in autumn. Though barely knee-high, their wood is dense and strong. These hardy shrub-like trees create important shelter for arctic wildlife."

	case TreeType_Arctic_Willow:
		return "Creeping arctic willows spread across the tundra, their stems rarely rising more than a few inches high. Their wood, though tiny in diameter, is remarkably resilient. These ground-hugging trees are vital components of the arctic ecosystem."

	case TreeType_Siberian_Elm:
		return "Tough Siberian elms endure harsh continental climates, their cork-like bark protecting them from extreme temperatures. The wood is hard and durable, resistant to decay. These adaptable trees spread their canopies wide when conditions allow."

	case TreeType_Pinyon_Pine:
		return "Sturdy pinyon pines dot arid highlands, their spreading crowns providing welcome shade. The fragrant wood burns hot and slow, while their nutritious nuts have sustained desert cultures for millennia. These trees define the transition between desert and mountain."

	case TreeType_Olive_Trees:
		return "Gnarled olive trees twist their ancient trunks skyward, their silver-green leaves shimmering in the Mediterranean breeze. The dense, beautifully figured wood is prized for carving. These trees can live for centuries, their productive years spanning multiple human generations."

	case TreeType_Carob_Tree:
		return "Ancient carob trees spread their dense evergreen canopies, their dark pods rich with sweet pulp. The hard, reddish wood takes a beautiful polish, while the deeply furrowed bark speaks of centuries of survival. These drought-resistant trees thrive where others struggle."

	case TreeType_Aleppo_Pine:
		return "Hardy Aleppo pines stand against the Mediterranean sky, their bark ranging from dark gray to reddish-brown. The wood is resinous and durable, historically used in shipbuilding. Their asymmetrical crowns and twisted trunks show their adaptation to coastal winds."

	case TreeType_Protea_Trees:
		return "Distinctive protea trees grace the landscape with their remarkable flowers and leathery leaves. The wood is fine-grained and pinkish, though rarely harvested. These ancient lineage trees create unique habitats for specialized wildlife."

	case TreeType_Silver_Tree:
		return "Shimmering silver trees catch the light with their silky-haired leaves, creating a magical effect in the breeze. The wood is soft but durable, while the silvery bark adds to their otherworldly appearance. These rare trees are living treasures of their native range."

	case TreeType_Mesquite:
		return "Tough mesquite trees spread their thorny branches across the desert, their extremely hard wood hiding beneath twisted trunks. The dense, chocolate-brown heartwood is prized for furniture and smoking meat. Their deep roots tap into underground water sources."

	case TreeType_Palo_Verde:
		return "Green-barked palo verde trees stand out in the desert landscape, their tiny leaves allowing the branches to photosynthesize. The wood is hard but brittle, while the golden spring flowers create stunning displays. These trees are masters of desert survival."

	case TreeType_Date_Palm:
		return "Stately date palms rise from the desert soil, their feathery fronds crowning thick trunks marked by old leaf bases. Though not true wood, their fibrous trunks provide valuable building material. Clusters of sweet dates hang in abundance when in season."

	case TreeType_Tamarisk:
		return "Salt-tolerant tamarisk trees create feathery screens of scale-like leaves, their pink flowers adding color to harsh environments. The wood is strong but irregular, while their deep roots seek out hidden water. These adaptable trees can survive where few others persist."

	case TreeType_Desert_Willow:
		return "Graceful desert willows bloom with orchid-like flowers, their slender branches dancing in desert breezes. The wood is light but strong, traditionally used for bow-making. These drought-adapted trees bring beauty to dry watercourses and canyon bottoms."

	default:
		return "A grove of trees grows here."
	}
}
