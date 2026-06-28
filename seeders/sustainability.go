package seeders

import (
	"log"
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
)

func SeedPulpProcessSections() {
	var count int64
	config.DB.Model(&models.PulpProcessSection{}).Count(&count)
	if count > 0 {
		return
	}
	sections := []models.PulpProcessSection{
		{Title: "Wood and Chip Handling", Description: "Raw logs from sustainable industrial plantations (Acacia Mangium & Eucalyptus Pellita) are transported to the mill, stored in wet ponds, then processed through debarking and chipping.", ImageURL: "/pulp proces/Chipping-Proses.jpg", SortOrder: 1, IsActive: true},
		{Title: "Fiber Line", Description: "Wood chips are cooked under high pressure with white liquor to separate cellulose fibers from lignin, then washed and delignified.", ImageURL: "/pulp proces/FiberLine.jpg", SortOrder: 2, IsActive: true},
		{Title: "Pulp Machine", Description: "The bleached pulp slurry is formed into sheets on a Fourdrinier wire, pressed, and steam-dried into market pulp.", ImageURL: "/pulp proces/pulp-machine.jpg", SortOrder: 3, IsActive: true},
		{Title: "Warehouse", Description: "Finished pulp sheets are cut, stacked, and compressed into 250 kg bales wrapped in kraft liner for storage and export.", ImageURL: "/pulp proces/IMG_6725.jpg", SortOrder: 4, IsActive: true},
	}
	for _, s := range sections {
		config.DB.Create(&s)
	}
	log.Println("Pulp process sections seeded")
}

func SeedPulpProcessRecoveries() {
	var count int64
	config.DB.Model(&models.PulpProcessRecovery{}).Count(&count)
	if count > 0 {
		return
	}
	recoveries := []models.PulpProcessRecovery{
		{Title: "Evaporator Plant (7-Effect)", Description: "Multiple-effect evaporator concentrates weak black liquor from 15% to 72% dissolved solids for firing in the Recovery Boiler.", ImageURL: "/pulp proces/Epavorator.jpg", Specs: `[{"label":"Inlet Concentration","value":"15% Dissolved Solids"},{"label":"Outlet Concentration","value":"72% Dissolved Solids"},{"label":"Effect Stages","value":"7-Effect System"},{"label":"Heat Source","value":"Exhaust Steam from Boilers"}]`, SortOrder: 1, IsActive: true},
		{Title: "Recovery Boiler", Description: "Burns concentrated black liquor (70% DS) to recover inorganic smelt and generate high-pressure steam.", ImageURL: "/pulp proces/Recovery-Boiler.jpg", Specs: `[{"label":"Design Capacity","value":"2,400 TDS/day"},{"label":"Fuel","value":"Concentrated Black Liquor"},{"label":"Products","value":"Smelt + High-Pressure Steam"},{"label":"Steam Usage","value":"Power Generation + Process"}]`, SortOrder: 2, IsActive: true},
		{Title: "Recausticizing Plant", Description: "Green liquor (Na2CO3 + Na2S) is reacted with Ca(OH)2 to regenerate White Liquor (NaOH + Na2S) for the cooking cycle.", ImageURL: "/pulp proces/Recautisizing.jpg", Specs: `[{"label":"Input","value":"Green Liquor (Na2CO3 + Na2S)"},{"label":"Reagent","value":"Ca(OH)2 (Slaked Lime)"},{"label":"Output","value":"White Liquor (NaOH + Na2S)"},{"label":"By-product","value":"Lime Mud (CaCO3) → Kiln"}]`, SortOrder: 3, IsActive: true},
		{Title: "Lime Kiln", Description: "Rotary kiln (~150m) reburns CaCO3 lime mud at 1000°C to regenerate CaO quicklime for the recausticizing cycle.", ImageURL: "/pulp proces/Lime-Kiln.jpg", Specs: `[{"label":"Length","value":"~150 meters"},{"label":"Temperature","value":"1,000°C"},{"label":"Input","value":"Lime Mud (CaCO3)"},{"label":"Output","value":"Quicklime (CaO) → Recausticizing"}]`, SortOrder: 4, IsActive: true},
	}
	for _, r := range recoveries {
		config.DB.Create(&r)
	}
	log.Println("Pulp process recoveries seeded")
}

func SeedSafetyPolicies() {
	var count int64
	config.DB.Model(&models.SafetyPolicy{}).Count(&count)
	if count > 0 {
		return
	}
	policies := []models.SafetyPolicy{
		{Description: "Certified management system ensuring continuous improvement and full compliance with government safety regulations.", SortOrder: 1, IsActive: true},
	}
	for _, p := range policies {
		config.DB.Create(&p)
	}
	log.Println("Safety policies seeded")
}

func SeedSafetyK3Targets() {
	var count int64
	config.DB.Model(&models.SafetyK3Target{}).Count(&count)
	if count > 0 {
		return
	}
	targets := []models.SafetyK3Target{
		{Description: "Update and adopt government regulations related to OSH by 100% in 2019.", SortOrder: 1, IsActive: true},
		{Description: "Renewed operating license deeds for all machines and OSH-related licenses for all equipment (lifts, lifting equipment, boilers and electricity) by 96% in 2019 issued by the Ministry of Manpower of the Republic of Indonesia.", SortOrder: 2, IsActive: true},
		{Description: "Ensure that companies, contractors, suppliers and vendors comply with regulations by 100% in 2019.", SortOrder: 3, IsActive: true},
		{Description: "Increase awareness of safety and safety in OHS risk assessments by providing training at least 10 times a year.", SortOrder: 4, IsActive: true},
		{Description: "Implement safety patrols a minimum of 10 times in one year.", SortOrder: 5, IsActive: true},
		{Description: "Conduct a hazard assessment in the work environment at least once a year.", SortOrder: 6, IsActive: true},
	}
	for _, t := range targets {
		config.DB.Create(&t)
	}
	log.Println("Safety K3 targets seeded")
}

func SeedSafetyK3Programs() {
	var count int64
	config.DB.Model(&models.SafetyK3Program{}).Count(&count)
	if count > 0 {
		return
	}
	programs := []models.SafetyK3Program{
		{Description: "Comply with government regulations related to occupational safety and health, both new and amended old regulations.", SortOrder: 1, IsActive: true},
		{Description: "Equip all work equipment and operational machinery with an operational permit.", SortOrder: 2, IsActive: true},
		{Description: "Ensuring companies, supplier contractors and vendors meet SMK3 requirements.", SortOrder: 3, IsActive: true},
		{Description: "Perform medical checkups.", SortOrder: 4, IsActive: true},
		{Description: "Conduct Training Need Analysis (TNA) related to K3 and conduct training for employees.", SortOrder: 5, IsActive: true},
		{Description: "K3 Promotion.", SortOrder: 6, IsActive: true},
		{Description: "Patrol and follow up findings / finding / K3 abnormality.", SortOrder: 7, IsActive: true},
		{Description: "Measure the Work Environment once a year and follow up abnormal parameters.", SortOrder: 8, IsActive: true},
		{Description: "Perform preventive maintenance for fire equipment.", SortOrder: 9, IsActive: true},
		{Description: "Implement work permits and fire work permits in the work environment.", SortOrder: 10, IsActive: true},
		{Description: "Investigating occupational accidents, Occupational Diseases (PAK) and fires.", SortOrder: 11, IsActive: true},
	}
	for _, p := range programs {
		config.DB.Create(&p)
	}
	log.Println("Safety K3 programs seeded")
}

func SeedSafetySliders() {
	var count int64
	config.DB.Model(&models.SafetySlider{}).Count(&count)
	if count > 0 {
		return
	}
	sliders := []models.SafetySlider{
		{ImageURL: "/safety/2.png", SortOrder: 1, IsActive: true},
		{ImageURL: "/safety/3.png", SortOrder: 2, IsActive: true},
		{ImageURL: "/safety/4.png", SortOrder: 3, IsActive: true},
		{ImageURL: "/safety/5.png", SortOrder: 4, IsActive: true},
		{ImageURL: "/safety/6.png", SortOrder: 5, IsActive: true},
	}
	for _, s := range sliders {
		config.DB.Create(&s)
	}
	log.Println("Safety sliders seeded")
}

func SeedSupplyChainStrategies() {
	items := []models.SupplyChainStrategy{
		{Title: "Strategy and Strong Management Plan", Description: "To keep stable pulp production and its good quality which consist of:", SortOrder: 1, IsActive: true},
		{Title: "Strategy and Strong Management Plan", Description: "Log Management", SortOrder: 2, IsActive: true},
		{Title: "Strategy and Strong Management Plan", Description: "Procurement Management", SortOrder: 3, IsActive: true},
		{Title: "Strategy and Strong Management Plan", Description: "Spare Part Management", SortOrder: 4, IsActive: true},
		{Title: "Strategy and Strong Management Plan", Description: "Contract Management", SortOrder: 5, IsActive: true},
		{Title: "Strategy and Strong Management Plan", Description: "To implement sound Industrial Relation, Corporate Social Responsibility.", SortOrder: 6, IsActive: true},
		{Title: "Strategy and Strong Management Plan", Description: "To develop competent and motivated People.", SortOrder: 7, IsActive: true},
	}
	for _, i := range items {
		var existing models.SupplyChainStrategy
		if config.DB.Where("description = ?", i.Description).First(&existing).Error != nil {
			config.DB.Create(&i)
		}
	}
	log.Println("Supply chain strategies seeded")
}

func SeedSupplyChainSustainabilityItems() {
	var count int64
	config.DB.Model(&models.SupplyChainSustainabilityItem{}).Count(&count)
	if count > 0 {
		return
	}
	items := []models.SupplyChainSustainabilityItem{
		{Description: "PT.Tanjungenim Lestari Pulp and Paper has a principle stated in a motto \"Fairness, Innovation & Harmony\" which underlines all of its activities and transactions that occur to ensure the sustainability of business activities in harmony.", SortOrder: 1, IsActive: true},
		{Description: "PT.Tanjungenim Lestari Pulp and Paper commits to uphold Good Corporate Governance (GCG) principles in its operation. To achieve this, the participation of all stakeholders including suppliers is very essential.", SortOrder: 2, IsActive: true},
		{Description: "PT.Tanjungenim Lestari Pulp and Paper is part of Marubeni Group who already set its supply chain sustainability guidelines, supporting an environmentally friendly, healthy and sustainable society.", SortOrder: 3, IsActive: true},
		{Description: "Cascading the Marubeni Group supply chain sustainability guidelines, PT.Tanjungenim Lestari Pulp and Paper asks for the understanding and cooperation of its business partners in observing the Guidelines.", SortOrder: 4, IsActive: true},
	}
	for _, i := range items {
		config.DB.Create(&i)
	}
	log.Println("Supply chain sustainability items seeded")
}

func SeedSupplyChainPolicies() {
	var count int64
	config.DB.Model(&models.SupplyChainPolicy{}).Count(&count)
	if count > 0 {
		return
	}
	policies := []models.SupplyChainPolicy{
		{Title: "Observance of Laws", Points: `["Observe the laws of the countries where business is conducted and laws relating to business transactions."]`, SortOrder: 1, IsActive: true},
		{Title: "Respect for Human Rights", Points: `["Respect human rights without discrimination, harassment of any kind, abuse or other inhumane treatment.","No child labor or forced labor.","Proper management of employees' work hours, breaks and holidays.","Payment of the legally mandated minimum wage.","Respect for employees' right to unionize."]`, SortOrder: 2, IsActive: true},
		{Title: "Conservation of the Environment", Points: `["Recognize that climate change issues are important and respond appropriately.","Protect the natural environment.","Reduce environmental negative impact; prevent pollution."]`, SortOrder: 3, IsActive: true},
		{Title: "Fair Transactions", Points: `["Conduct fair transactions and do not inhibit free competition.","Prevent corruption; offer no bribes or illegal contributions."]`, SortOrder: 4, IsActive: true},
		{Title: "Safety and Health", Points: `["Ensure safe and healthy workplaces and maintain a good working environment."]`, SortOrder: 5, IsActive: true},
		{Title: "Quality Control", Points: `["Maintain the quality and safety of products and services."]`, SortOrder: 6, IsActive: true},
		{Title: "Information Disclosure", Points: `["Timely and appropriate disclosure of information.","The Marubeni Group has set out procedures for dealing with vendors that do not meet labor standards:"]`, Procedures: `["When it comes to light that a vendor has failed to meet labor standards, we will ask the vendor to ascertain the facts and prepare a report on the background of the issue and improvement measures.","If we determine that improvement measures are insufficient, we will request that further measures be taken.","If, despite implementing steps (i) and (ii) above, the situation does not improve, we will examine whether to continue our relationship with the vendor."]`, SortOrder: 7, IsActive: true},
	}
	for _, p := range policies {
		config.DB.Create(&p)
	}
	log.Println("Supply chain policies seeded")
}

func SeedCsrVisionStrategies() {
	var count int64
	config.DB.Model(&models.CsrVisionStrategy{}).Count(&count)
	if count > 0 {
		return
	}
	items := []models.CsrVisionStrategy{
		{Description: "To implement sound industrial relations Corporate Social Responsibility (CSR) and Good Corporate Governance", SortOrder: 1, IsActive: true},
		{Description: "To ensure a harmony industrial relations with stakeholders both internal and external", SortOrder: 2, IsActive: true},
		{Description: "To strengthen the implementation of Good Corporate Governance (GCG) in accordance with prevailing regulations and laws", SortOrder: 3, IsActive: true},
		{Description: "To do continuous Improvement Program (CIP) in each department", SortOrder: 4, IsActive: true},
	}
	for _, i := range items {
		config.DB.Create(&i)
	}
	log.Println("CSR vision strategies seeded")
}

func SeedCsrReports() {
	var count int64
	config.DB.Model(&models.CsrReport{}).Count(&count)
	if count > 0 {
		return
	}
	reports := []models.CsrReport{
		{Year: 2018, Quarter: "1st Quarter", Period: "", FileURL: "", SortOrder: 1, IsActive: true},
		{Year: 2018, Quarter: "2nd Quarter", Period: "", FileURL: "", SortOrder: 2, IsActive: true},
		{Year: 2018, Quarter: "3rd Quarter", Period: "", FileURL: "", SortOrder: 3, IsActive: true},
		{Year: 2018, Quarter: "4th Quarter", Period: "", FileURL: "", SortOrder: 4, IsActive: true},
		{Year: 2019, Quarter: "1st Quarter", Period: "Jan – Mar 2019", FileURL: "CSRReport2019_1st_Quarter_Apr_Ju.pdf", SortOrder: 1, IsActive: true},
		{Year: 2019, Quarter: "2nd Quarter", Period: "", FileURL: "", SortOrder: 2, IsActive: true},
		{Year: 2019, Quarter: "3rd Quarter", Period: "", FileURL: "", SortOrder: 3, IsActive: true},
		{Year: 2019, Quarter: "4th Quarter", Period: "jan – mar 2019", FileURL: "CSR-Report-2019-4th-Quarter.pdf", SortOrder: 4, IsActive: true},
		{Year: 2020, Quarter: "1st Quarter", Period: "Apr – jun 2020", FileURL: "CSR-Report-2020-1st-Quarter-Apr-Jun-2020-W.pdf", SortOrder: 1, IsActive: true},
		{Year: 2020, Quarter: "2nd Quarter", Period: "Jul – Sep 2020", FileURL: "CSR-Report-2020-2nd-Quarter-Jul-Sep-2020-f.pdf", SortOrder: 2, IsActive: true},
		{Year: 2020, Quarter: "3rd Quarter", Period: "", FileURL: "", SortOrder: 3, IsActive: true},
		{Year: 2020, Quarter: "4th Quarter", Period: "", FileURL: "", SortOrder: 4, IsActive: true},
	}
	for _, r := range reports {
		config.DB.Create(&r)
	}
	log.Println("CSR reports seeded")
}
