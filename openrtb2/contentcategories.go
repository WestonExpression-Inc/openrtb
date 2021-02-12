package openrtb

// ORTBContentCategory represents the IAB’s contextual taxonomy for categorization. Standard IDs have been
// adopted to easily support the communication of primary and secondary categories for various objects.
// This OpenRTB table has values derived from the IAB Tech Lab Content Taxonomy. Practitioners should
// keep in sync with updates as published on www.iab.com.
// OpenRTB 2.5 Section 5.1
type ORTBContentCategory string

const (
	// ORTBContentCatArtsAndEntertainment is a content category for Arts & Entertainment
	ORTBContentCatArtsAndEntertainment ORTBContentCategory = "IAB1"
	// ORTBContentCatBooksLiterature is a content category for Books Literature
	ORTBContentCatBooksLiterature ORTBContentCategory = "IAB1-1"
	// ORTBContentCatCelebrityFanGossip is a content category for Celebrity Fan/Gossip
	ORTBContentCatCelebrityFanGossip ORTBContentCategory = "IAB1-2"
	// ORTBContentCatFineArt is a content category for Fine Art
	ORTBContentCatFineArt ORTBContentCategory = "IAB1-3"
	// ORTBContentCatHumor is a content category for Humor
	ORTBContentCatHumor ORTBContentCategory = "IAB1-4"
	// ORTBContentCatMovies is a content category for Movies
	ORTBContentCatMovies ORTBContentCategory = "IAB1-5"
	// ORTBContentCatMusic is a content category for Music
	ORTBContentCatMusic ORTBContentCategory = "IAB1-6"
	// ORTBContentCatTelevision is a content category for Television
	ORTBContentCatTelevision ORTBContentCategory = "IAB1-7"
	// ORTBContentCatAutomotive is a content category for Automotive
	ORTBContentCatAutomotive ORTBContentCategory = "IAB2"
	// ORTBContentCatAutoParts is a content category for Auto Parts
	ORTBContentCatAutoParts ORTBContentCategory = "IAB2-1"
	// ORTBContentCatAutoRepair is a content category for Auto Repair
	ORTBContentCatAutoRepair ORTBContentCategory = "IAB2-2"
	// ORTBContentCatBuyingSellingCars is a content category for Buying/Selling Cars
	ORTBContentCatBuyingSellingCars ORTBContentCategory = "IAB2-3"
	// ORTBContentCatCarCulture is a content category for Car Culture
	ORTBContentCatCarCulture ORTBContentCategory = "IAB2-4"
	// ORTBContentCatCertifiedPreOwned is a content category for Certified Pre-Owned
	ORTBContentCatCertifiedPreOwned ORTBContentCategory = "IAB2-5"
	// ORTBContentCatConvertible is a content category for Convertible
	ORTBContentCatConvertible ORTBContentCategory = "IAB2-6"
	// ORTBContentCatCoupe is a content category for Coupe
	ORTBContentCatCoupe ORTBContentCategory = "IAB2-7"
	// ORTBContentCatCrossover is a content category for Crossover
	ORTBContentCatCrossover ORTBContentCategory = "IAB2-8"
	// ORTBContentCatDiesel is a content category for Diesel
	ORTBContentCatDiesel ORTBContentCategory = "IAB2-9"
	// ORTBContentCatElectricVehicle is a content category for Electric Vehicle
	ORTBContentCatElectricVehicle ORTBContentCategory = "IAB2-10"
	// ORTBContentCatHatchback is a content category for Hatchback
	ORTBContentCatHatchback ORTBContentCategory = "IAB2-11"
	// ORTBContentCatHybrid is a content category for Hybrid
	ORTBContentCatHybrid ORTBContentCategory = "IAB2-12"
	// ORTBContentCatLuxury is a content category for Luxury
	ORTBContentCatLuxury ORTBContentCategory = "IAB2-13"
	// ORTBContentCatMinivan is a content category for Minivan
	ORTBContentCatMinivan ORTBContentCategory = "IAB2-14"
	// ORTBContentCatMotorcycles is a content category for Motorcycles
	ORTBContentCatMotorcycles ORTBContentCategory = "IAB2-15"
	// ORTBContentCatOffRoadVehicles is a content category for Off-Road Vehicles
	ORTBContentCatOffRoadVehicles ORTBContentCategory = "IAB2-16"
	// ORTBContentCatPerformanceVehicles is a content category for Performance Vehicles
	ORTBContentCatPerformanceVehicles ORTBContentCategory = "IAB2-17"
	// ORTBContentCatPickup is a content category for Pickup
	ORTBContentCatPickup ORTBContentCategory = "IAB2-18"
	// ORTBContentCatRoadSideAssistance is a content category for Road-Side Assistance
	ORTBContentCatRoadSideAssistance ORTBContentCategory = "IAB2-19"
	// ORTBContentCatSedan is a content category for Sedan
	ORTBContentCatSedan ORTBContentCategory = "IAB2-20"
	// ORTBContentCatTrucksAndAccessories is a content category for Trucks & Accessories
	ORTBContentCatTrucksAndAccessories ORTBContentCategory = "IAB2-21"
	// ORTBContentCatVintageCars is a content category for Vintage Cars
	ORTBContentCatVintageCars ORTBContentCategory = "IAB2-22"
	// ORTBContentCatWagon is a content category for Wagon
	ORTBContentCatWagon ORTBContentCategory = "IAB2-23"
	// ORTBContentCatBusiness is a content category for Business
	ORTBContentCatBusiness ORTBContentCategory = "IAB3"
	// ORTBContentCatAdvertising is a content category for Advertising
	ORTBContentCatAdvertising ORTBContentCategory = "IAB3-1"
	// ORTBContentCatAgriculture is a content category for Agriculture
	ORTBContentCatAgriculture ORTBContentCategory = "IAB3-2"
	// ORTBContentCatBiotechBiomedical is a content category for Biotech/Biomedical
	ORTBContentCatBiotechBiomedical ORTBContentCategory = "IAB3-3"
	// ORTBContentCatBusinessSoftware is a content category for Business Software
	ORTBContentCatBusinessSoftware ORTBContentCategory = "IAB3-4"
	// ORTBContentCatConstruction is a content category for Construction
	ORTBContentCatConstruction ORTBContentCategory = "IAB3-5"
	// ORTBContentCatForestry is a content category for Forestry
	ORTBContentCatForestry ORTBContentCategory = "IAB3-6"
	// ORTBContentCatGovernment is a content category for Government
	ORTBContentCatGovernment ORTBContentCategory = "IAB3-7"
	// ORTBContentCatGreenSolutions is a content category for Green Solutions
	ORTBContentCatGreenSolutions ORTBContentCategory = "IAB3-8"
	// ORTBContentCatHumanResources is a content category for Human Resources
	ORTBContentCatHumanResources ORTBContentCategory = "IAB3-9"
	// ORTBContentCatLogistics is a content category for Logistics
	ORTBContentCatLogistics ORTBContentCategory = "IAB3-10"
	// ORTBContentCatMarketing is a content category for Marketing
	ORTBContentCatMarketing ORTBContentCategory = "IAB3-11"
	// ORTBContentCatMetals is a content category for Metals
	ORTBContentCatMetals ORTBContentCategory = "IAB3-12"
	// ORTBContentCatCareers is a content category for Careers
	ORTBContentCatCareers ORTBContentCategory = "IAB4"
	// ORTBContentCatCareerPlanning is a content category for Career Planning
	ORTBContentCatCareerPlanning ORTBContentCategory = "IAB4-1"
	// ORTBContentCatCollege is a content category for College
	ORTBContentCatCollege ORTBContentCategory = "IAB4-2"
	// ORTBContentCatFinancialAid is a content category for Financial Aid
	ORTBContentCatFinancialAid ORTBContentCategory = "IAB4-3"
	// ORTBContentCatJobFairs is a content category for Job Fairs
	ORTBContentCatJobFairs ORTBContentCategory = "IAB4-4"
	// ORTBContentCatJobSearch is a content category for Job Search
	ORTBContentCatJobSearch ORTBContentCategory = "IAB4-5"
	// ORTBContentCatResumeWritingAdvice is a content category for Resume Writing/Advice
	ORTBContentCatResumeWritingAdvice ORTBContentCategory = "IAB4-6"
	// ORTBContentCatNursing is a content category for Nursing
	ORTBContentCatNursing ORTBContentCategory = "IAB4-7"
	// ORTBContentCatScholarships is a content category for Scholarships
	ORTBContentCatScholarships ORTBContentCategory = "IAB4-8"
	// ORTBContentCatTelecommuting is a content category for Telecommuting
	ORTBContentCatTelecommuting ORTBContentCategory = "IAB4-9"
	// ORTBContentCatUSMilitary is a content category for U.S. Military
	ORTBContentCatUSMilitary ORTBContentCategory = "IAB4-10"
	// ORTBContentCatCareerAdvice is a content category for Career Advice
	ORTBContentCatCareerAdvice ORTBContentCategory = "IAB4-11"
	// ORTBContentCatEducation is a content category for Education
	ORTBContentCatEducation ORTBContentCategory = "IAB5"
	// ORTBContentCat712Education is a content category for 7-12 Education
	ORTBContentCat712Education ORTBContentCategory = "IAB5-1"
	// ORTBContentCatAdultEducation is a content category for Adult Education
	ORTBContentCatAdultEducation ORTBContentCategory = "IAB5-2"
	// ORTBContentCatArtHistory is a content category for Art History
	ORTBContentCatArtHistory ORTBContentCategory = "IAB5-3"
	// ORTBContentCatCollegeAdministration is a content category for College Administration
	ORTBContentCatCollegeAdministration ORTBContentCategory = "IAB5-4"
	// ORTBContentCatCollegeLife is a content category for College Life
	ORTBContentCatCollegeLife ORTBContentCategory = "IAB5-5"
	// ORTBContentCatDistanceLearning is a content category for Distance Learning
	ORTBContentCatDistanceLearning ORTBContentCategory = "IAB5-6"
	// ORTBContentCatEnglishasa2ndLanguage is a content category for English as a 2nd Language
	ORTBContentCatEnglishasa2ndLanguage ORTBContentCategory = "IAB5-7"
	// ORTBContentCatLanguageLearning is a content category for Language Learning
	ORTBContentCatLanguageLearning ORTBContentCategory = "IAB5-8"
	// ORTBContentCatGraduateSchool is a content category for Graduate School
	ORTBContentCatGraduateSchool ORTBContentCategory = "IAB5-9"
	// ORTBContentCatHomeschooling is a content category for Homeschooling
	ORTBContentCatHomeschooling ORTBContentCategory = "IAB5-10"
	// ORTBContentCatHomeworkStudyTips is a content category for Homework/Study Tips
	ORTBContentCatHomeworkStudyTips ORTBContentCategory = "IAB5-11"
	// ORTBContentCatK6Educators is a content category for K-6 Educators
	ORTBContentCatK6Educators ORTBContentCategory = "IAB5-12"
	// ORTBContentCatPrivateSchool is a content category for Private School
	ORTBContentCatPrivateSchool ORTBContentCategory = "IAB5-13"
	// ORTBContentCatSpecialEducation is a content category for Special Education
	ORTBContentCatSpecialEducation ORTBContentCategory = "IAB5-14"
	// ORTBContentCatStudyingBusiness is a content category for Studying Business
	ORTBContentCatStudyingBusiness ORTBContentCategory = "IAB5-15"
	// ORTBContentCatFamilyAndParenting is a content category for Family & Parenting
	ORTBContentCatFamilyAndParenting ORTBContentCategory = "IAB6"
	// ORTBContentCatAdoption is a content category for Adoption
	ORTBContentCatAdoption ORTBContentCategory = "IAB6-1"
	// ORTBContentCatBabiesAndToddlers is a content category for Babies & Toddlers
	ORTBContentCatBabiesAndToddlers ORTBContentCategory = "IAB6-2"
	// ORTBContentCatDaycarePreSchool is a content category for Daycare/Pre School
	ORTBContentCatDaycarePreSchool ORTBContentCategory = "IAB6-3"
	// ORTBContentCatFamilyInternet is a content category for Family Internet
	ORTBContentCatFamilyInternet ORTBContentCategory = "IAB6-4"
	// ORTBContentCatParentingK6Kids is a content category for Parenting - K-6 Kids
	ORTBContentCatParentingK6Kids ORTBContentCategory = "IAB6-5"
	// ORTBContentCatParentingteens is a content category for Parenting teens
	ORTBContentCatParentingteens ORTBContentCategory = "IAB6-6"
	// ORTBContentCatPregnancy is a content category for Pregnancy
	ORTBContentCatPregnancy ORTBContentCategory = "IAB6-7"
	// ORTBContentCatSpecialNeedsKids is a content category for Special Needs Kids
	ORTBContentCatSpecialNeedsKids ORTBContentCategory = "IAB6-8"
	// ORTBContentCatEldercare is a content category for Eldercare
	ORTBContentCatEldercare ORTBContentCategory = "IAB6-9"
	// ORTBContentCatHealthAndFitness is a content category for Health & Fitness
	ORTBContentCatHealthAndFitness ORTBContentCategory = "IAB7"
	// ORTBContentCatExercise is a content category for Exercise
	ORTBContentCatExercise ORTBContentCategory = "IAB7-1"
	// ORTBContentCatADD is a content category for ADD
	ORTBContentCatADD ORTBContentCategory = "IAB7-2"
	// ORTBContentCatAIDSHIV is a content category for AIDS/HIV
	ORTBContentCatAIDSHIV ORTBContentCategory = "IAB7-3"
	// ORTBContentCatAllergies is a content category for Allergies
	ORTBContentCatAllergies ORTBContentCategory = "IAB7-4"
	// ORTBContentCatAlternativeMedicine is a content category for Alternative Medicine
	ORTBContentCatAlternativeMedicine ORTBContentCategory = "IAB7-5"
	// ORTBContentCatArthritis is a content category for Arthritis
	ORTBContentCatArthritis ORTBContentCategory = "IAB7-6"
	// ORTBContentCatAsthma is a content category for Asthma
	ORTBContentCatAsthma ORTBContentCategory = "IAB7-7"
	// ORTBContentCatAutismPDD is a content category for Autism/PDD
	ORTBContentCatAutismPDD ORTBContentCategory = "IAB7-8"
	// ORTBContentCatBipolarDisorder is a content category for Bipolar Disorder
	ORTBContentCatBipolarDisorder ORTBContentCategory = "IAB7-9"
	// ORTBContentCatBrainTumor is a content category for Brain Tumor
	ORTBContentCatBrainTumor ORTBContentCategory = "IAB7-10"
	// ORTBContentCatCancer is a content category for Cancer
	ORTBContentCatCancer ORTBContentCategory = "IAB7-11"
	// ORTBContentCatCholesterol is a content category for Cholesterol
	ORTBContentCatCholesterol ORTBContentCategory = "IAB7-12"
	// ORTBContentCatChronicFatigueSyndrome is a content category for Chronic Fatigue Syndrome
	ORTBContentCatChronicFatigueSyndrome ORTBContentCategory = "IAB7-13"
	// ORTBContentCatChronicPain is a content category for Chronic Pain
	ORTBContentCatChronicPain ORTBContentCategory = "IAB7-14"
	// ORTBContentCatColdAndFlu is a content category for Cold & Flu
	ORTBContentCatColdAndFlu ORTBContentCategory = "IAB7-15"
	// ORTBContentCatDeafness is a content category for Deafness
	ORTBContentCatDeafness ORTBContentCategory = "IAB7-16"
	// ORTBContentCatDentalCare is a content category for Dental Care
	ORTBContentCatDentalCare ORTBContentCategory = "IAB7-17"
	// ORTBContentCatDepression is a content category for Depression
	ORTBContentCatDepression ORTBContentCategory = "IAB7-18"
	// ORTBContentCatDermatology is a content category for Dermatology
	ORTBContentCatDermatology ORTBContentCategory = "IAB7-19"
	// ORTBContentCatDiabetes is a content category for Diabetes
	ORTBContentCatDiabetes ORTBContentCategory = "IAB7-20"
	// ORTBContentCatEpilepsy is a content category for Epilepsy
	ORTBContentCatEpilepsy ORTBContentCategory = "IAB7-21"
	// ORTBContentCatGERDAcidReflux is a content category for GERD/Acid Reflux
	ORTBContentCatGERDAcidReflux ORTBContentCategory = "IAB7-22"
	// ORTBContentCatHeadachesMigraines is a content category for Headaches/Migraines
	ORTBContentCatHeadachesMigraines ORTBContentCategory = "IAB7-23"
	// ORTBContentCatHeartDisease is a content category for Heart Disease
	ORTBContentCatHeartDisease ORTBContentCategory = "IAB7-24"
	// ORTBContentCatHerbsforHealth is a content category for Herbs for Health
	ORTBContentCatHerbsforHealth ORTBContentCategory = "IAB7-25"
	// ORTBContentCatHolisticHealing is a content category for Holistic Healing
	ORTBContentCatHolisticHealing ORTBContentCategory = "IAB7-26"
	// ORTBContentCatIBSCrohnsDisease is a content category for IBS/Crohn’s Disease
	ORTBContentCatIBSCrohnsDisease ORTBContentCategory = "IAB7-27"
	// ORTBContentCatIncestAbuseSupport is a content category for Incest/Abuse Support
	ORTBContentCatIncestAbuseSupport ORTBContentCategory = "IAB7-28"
	// ORTBContentCatIncontinence is a content category for Incontinence
	ORTBContentCatIncontinence ORTBContentCategory = "IAB7-29"
	// ORTBContentCatInfertility is a content category for Infertility
	ORTBContentCatInfertility ORTBContentCategory = "IAB7-30"
	// ORTBContentCatMensHealth is a content category for Men’s Health
	ORTBContentCatMensHealth ORTBContentCategory = "IAB7-31"
	// ORTBContentCatNutrition is a content category for Nutrition
	ORTBContentCatNutrition ORTBContentCategory = "IAB7-32"
	// ORTBContentCatOrthopedics is a content category for Orthopedics
	ORTBContentCatOrthopedics ORTBContentCategory = "IAB7-33"
	// ORTBContentCatPanicAnxietyDisorders is a content category for Panic/Anxiety Disorders
	ORTBContentCatPanicAnxietyDisorders ORTBContentCategory = "IAB7-34"
	// ORTBContentCatPediatrics is a content category for Pediatrics
	ORTBContentCatPediatrics ORTBContentCategory = "IAB7-35"
	// ORTBContentCatPhysicalTherapy is a content category for Physical Therapy
	ORTBContentCatPhysicalTherapy ORTBContentCategory = "IAB7-36"
	// ORTBContentCatPsychologyPsychiatry is a content category for Psychology/Psychiatry
	ORTBContentCatPsychologyPsychiatry ORTBContentCategory = "IAB7-37"
	// ORTBContentCatSeniorHealth is a content category for Senior Health
	ORTBContentCatSeniorHealth ORTBContentCategory = "IAB7-38"
	// ORTBContentCatSexuality is a content category for Sexuality
	ORTBContentCatSexuality ORTBContentCategory = "IAB7-39"
	// ORTBContentCatSleepDisorders is a content category for Sleep Disorders
	ORTBContentCatSleepDisorders ORTBContentCategory = "IAB7-40"
	// ORTBContentCatSmokingCessation is a content category for Smoking Cessation
	ORTBContentCatSmokingCessation ORTBContentCategory = "IAB7-41"
	// ORTBContentCatSubstanceAbuse is a content category for Substance Abuse
	ORTBContentCatSubstanceAbuse ORTBContentCategory = "IAB7-42"
	// ORTBContentCatThyroidDisease is a content category for Thyroid Disease
	ORTBContentCatThyroidDisease ORTBContentCategory = "IAB7-43"
	// ORTBContentCatWeightLoss is a content category for Weight Loss
	ORTBContentCatWeightLoss ORTBContentCategory = "IAB7-44"
	// ORTBContentCatWomensHealth is a content category for Womens Health
	ORTBContentCatWomensHealth ORTBContentCategory = "IAB7-45"
	// ORTBContentCatFoodAndDrink is a content category for Food & Drink
	ORTBContentCatFoodAndDrink ORTBContentCategory = "IAB8"
	// ORTBContentCatAmericanCuisine is a content category for American Cuisine
	ORTBContentCatAmericanCuisine ORTBContentCategory = "IAB8-1"
	// ORTBContentCatBarbecuesAndGrilling is a content category for Barbecues & Grilling
	ORTBContentCatBarbecuesAndGrilling ORTBContentCategory = "IAB8-2"
	// ORTBContentCatCajunCreole is a content category for Cajun/Creole
	ORTBContentCatCajunCreole ORTBContentCategory = "IAB8-3"
	// ORTBContentCatChineseCuisine is a content category for Chinese Cuisine
	ORTBContentCatChineseCuisine ORTBContentCategory = "IAB8-4"
	// ORTBContentCatCocktailsBeer is a content category for Cocktails/Beer
	ORTBContentCatCocktailsBeer ORTBContentCategory = "IAB8-5"
	// ORTBContentCatCoffeeTea is a content category for Coffee/Tea
	ORTBContentCatCoffeeTea ORTBContentCategory = "IAB8-6"
	// ORTBContentCatCuisineSpecific is a content category for Cuisine-Specific
	ORTBContentCatCuisineSpecific ORTBContentCategory = "IAB8-7"
	// ORTBContentCatDessertsAndBaking is a content category for Desserts & Baking
	ORTBContentCatDessertsAndBaking ORTBContentCategory = "IAB8-8"
	// ORTBContentCatDiningOut is a content category for Dining Out
	ORTBContentCatDiningOut ORTBContentCategory = "IAB8-9"
	// ORTBContentCatFoodAllergies is a content category for Food Allergies
	ORTBContentCatFoodAllergies ORTBContentCategory = "IAB8-10"
	// ORTBContentCatFrenchCuisine is a content category for French Cuisine
	ORTBContentCatFrenchCuisine ORTBContentCategory = "IAB8-11"
	// ORTBContentCatHealthLowFatCooking is a content category for Health/Low-Fat Cooking
	ORTBContentCatHealthLowFatCooking ORTBContentCategory = "IAB8-12"
	// ORTBContentCatItalianCuisine is a content category for Italian Cuisine
	ORTBContentCatItalianCuisine ORTBContentCategory = "IAB8-13"
	// ORTBContentCatJapaneseCuisine is a content category for Japanese Cuisine
	ORTBContentCatJapaneseCuisine ORTBContentCategory = "IAB8-14"
	// ORTBContentCatMexicanCuisine is a content category for Mexican Cuisine
	ORTBContentCatMexicanCuisine ORTBContentCategory = "IAB8-15"
	// ORTBContentCatVegan is a content category for Vegan
	ORTBContentCatVegan ORTBContentCategory = "IAB8-16"
	// ORTBContentCatVegetarian is a content category for Vegetarian
	ORTBContentCatVegetarian ORTBContentCategory = "IAB8-17"
	// ORTBContentCatWine is a content category for Wine
	ORTBContentCatWine ORTBContentCategory = "IAB8-18"
	// ORTBContentCatHobbiesAndInterests is a content category for Hobbies & Interests
	ORTBContentCatHobbiesAndInterests ORTBContentCategory = "IAB9"
	// ORTBContentCatArtTechnology is a content category for Art/Technology
	ORTBContentCatArtTechnology ORTBContentCategory = "IAB9-1"
	// ORTBContentCatArtsAndCrafts is a content category for Arts & Crafts
	ORTBContentCatArtsAndCrafts ORTBContentCategory = "IAB9-2"
	// ORTBContentCatBeadwork is a content category for Beadwork
	ORTBContentCatBeadwork ORTBContentCategory = "IAB9-3"
	// ORTBContentCatBirdWatching is a content category for Bird-Watching
	ORTBContentCatBirdWatching ORTBContentCategory = "IAB9-4"
	// ORTBContentCatBoardGamesPuzzles is a content category for Board Games/Puzzles
	ORTBContentCatBoardGamesPuzzles ORTBContentCategory = "IAB9-5"
	// ORTBContentCatCandleAndSoapMaking is a content category for Candle & Soap Making
	ORTBContentCatCandleAndSoapMaking ORTBContentCategory = "IAB9-6"
	// ORTBContentCatCardGames is a content category for Card Games
	ORTBContentCatCardGames ORTBContentCategory = "IAB9-7"
	// ORTBContentCatChess is a content category for Chess
	ORTBContentCatChess ORTBContentCategory = "IAB9-8"
	// ORTBContentCatCigars is a content category for Cigars
	ORTBContentCatCigars ORTBContentCategory = "IAB9-9"
	// ORTBContentCatCollecting is a content category for Collecting
	ORTBContentCatCollecting ORTBContentCategory = "IAB9-10"
	// ORTBContentCatComicBooks is a content category for Comic Books
	ORTBContentCatComicBooks ORTBContentCategory = "IAB9-11"
	// ORTBContentCatDrawingSketching is a content category for Drawing/Sketching
	ORTBContentCatDrawingSketching ORTBContentCategory = "IAB9-12"
	// ORTBContentCatFreelanceWriting is a content category for Freelance Writing
	ORTBContentCatFreelanceWriting ORTBContentCategory = "IAB9-13"
	// ORTBContentCatGenealogy is a content category for Genealogy
	ORTBContentCatGenealogy ORTBContentCategory = "IAB9-14"
	// ORTBContentCatGettingPublished is a content category for Getting Published
	ORTBContentCatGettingPublished ORTBContentCategory = "IAB9-15"
	// ORTBContentCatGuitar is a content category for Guitar
	ORTBContentCatGuitar ORTBContentCategory = "IAB9-16"
	// ORTBContentCatHomeRecording is a content category for Home Recording
	ORTBContentCatHomeRecording ORTBContentCategory = "IAB9-17"
	// ORTBContentCatInvestorsAndPatents is a content category for Investors & Patents
	ORTBContentCatInvestorsAndPatents ORTBContentCategory = "IAB9-18"
	// ORTBContentCatJewelryMaking is a content category for Jewelry Making
	ORTBContentCatJewelryMaking ORTBContentCategory = "IAB9-19"
	// ORTBContentCatMagicAndIllusion is a content category for Magic & Illusion
	ORTBContentCatMagicAndIllusion ORTBContentCategory = "IAB9-20"
	// ORTBContentCatNeedlework is a content category for Needlework
	ORTBContentCatNeedlework ORTBContentCategory = "IAB9-21"
	// ORTBContentCatPainting is a content category for Painting
	ORTBContentCatPainting ORTBContentCategory = "IAB9-22"
	// ORTBContentCatPhotography is a content category for Photography
	ORTBContentCatPhotography ORTBContentCategory = "IAB9-23"
	// ORTBContentCatRadio is a content category for Radio
	ORTBContentCatRadio ORTBContentCategory = "IAB9-24"
	// ORTBContentCatRoleplayingGames is a content category for Roleplaying Games
	ORTBContentCatRoleplayingGames ORTBContentCategory = "IAB9-25"
	// ORTBContentCatSciFiAndFantasy is a content category for Sci-Fi & Fantasy
	ORTBContentCatSciFiAndFantasy ORTBContentCategory = "IAB9-26"
	// ORTBContentCatScrapbooking is a content category for Scrapbooking
	ORTBContentCatScrapbooking ORTBContentCategory = "IAB9-27"
	// ORTBContentCatScreenwriting is a content category for Screenwriting
	ORTBContentCatScreenwriting ORTBContentCategory = "IAB9-28"
	// ORTBContentCatStampsAndCoins is a content category for Stamps & Coins
	ORTBContentCatStampsAndCoins ORTBContentCategory = "IAB9-29"
	// ORTBContentCatVideoAndComputerGames is a content category for Video & Computer Games
	ORTBContentCatVideoAndComputerGames ORTBContentCategory = "IAB9-30"
	// ORTBContentCatWoodworking is a content category for Woodworking
	ORTBContentCatWoodworking ORTBContentCategory = "IAB9-31"
	// ORTBContentCatHomeAndGarden is a content category for Home & Garden
	ORTBContentCatHomeAndGarden ORTBContentCategory = "IAB10"
	// ORTBContentCatAppliances is a content category for Appliances
	ORTBContentCatAppliances ORTBContentCategory = "IAB10-1"
	// ORTBContentCatEntertaining is a content category for Entertaining
	ORTBContentCatEntertaining ORTBContentCategory = "IAB10-2"
	// ORTBContentCatEnvironmentalSafety is a content category for Environmental Safety
	ORTBContentCatEnvironmentalSafety ORTBContentCategory = "IAB10-3"
	// ORTBContentCatGardening is a content category for Gardening
	ORTBContentCatGardening ORTBContentCategory = "IAB10-4"
	// ORTBContentCatHomeRepair is a content category for Home Repair
	ORTBContentCatHomeRepair ORTBContentCategory = "IAB10-5"
	// ORTBContentCatHomeTheater is a content category for Home Theater
	ORTBContentCatHomeTheater ORTBContentCategory = "IAB10-6"
	// ORTBContentCatInteriorDecorating is a content category for Interior Decorating
	ORTBContentCatInteriorDecorating ORTBContentCategory = "IAB10-7"
	// ORTBContentCatLandscaping is a content category for Landscaping
	ORTBContentCatLandscaping ORTBContentCategory = "IAB10-8"
	// ORTBContentCatRemodelingAndConstruction is a content category for Remodeling & Construction
	ORTBContentCatRemodelingAndConstruction ORTBContentCategory = "IAB10-9"
	// ORTBContentCatLawGovernmentAndPolitics is a content category for Law, Government, & Politics
	ORTBContentCatLawGovernmentAndPolitics ORTBContentCategory = "IAB11"
	// ORTBContentCatImmigration is a content category for Immigration
	ORTBContentCatImmigration ORTBContentCategory = "IAB11-1"
	// ORTBContentCatLegalIssues is a content category for Legal Issues
	ORTBContentCatLegalIssues ORTBContentCategory = "IAB11-2"
	// ORTBContentCatUSGovernmentResources is a content category for U.S. Government Resources
	ORTBContentCatUSGovernmentResources ORTBContentCategory = "IAB11-3"
	// ORTBContentCatPolitics is a content category for Politics
	ORTBContentCatPolitics ORTBContentCategory = "IAB11-4"
	// ORTBContentCatCommentary is a content category for Commentary
	ORTBContentCatCommentary ORTBContentCategory = "IAB11-5"
	// ORTBContentCatNews is a content category for News
	ORTBContentCatNews ORTBContentCategory = "IAB12"
	// ORTBContentCatInternationalNews is a content category for International News
	ORTBContentCatInternationalNews ORTBContentCategory = "IAB12-1"
	// ORTBContentCatNationalNews is a content category for National News
	ORTBContentCatNationalNews ORTBContentCategory = "IAB12-2"
	// ORTBContentCatLocalNews is a content category for Local News
	ORTBContentCatLocalNews ORTBContentCategory = "IAB12-3"
	// ORTBContentCatPersonalFinance is a content category for Personal Finance
	ORTBContentCatPersonalFinance ORTBContentCategory = "IAB13"
	// ORTBContentCatBeginningInvesting is a content category for Beginning Investing
	ORTBContentCatBeginningInvesting ORTBContentCategory = "IAB13-1"
	// ORTBContentCatCreditDebtAndLoans is a content category for Credit/Debt & Loans
	ORTBContentCatCreditDebtAndLoans ORTBContentCategory = "IAB13-2"
	// ORTBContentCatFinancialNews is a content category for Financial News
	ORTBContentCatFinancialNews ORTBContentCategory = "IAB13-3"
	// ORTBContentCatFinancialPlanning is a content category for Financial Planning
	ORTBContentCatFinancialPlanning ORTBContentCategory = "IAB13-4"
	// ORTBContentCatHedgeFund is a content category for Hedge Fund
	ORTBContentCatHedgeFund ORTBContentCategory = "IAB13-5"
	// ORTBContentCatInsurance is a content category for Insurance
	ORTBContentCatInsurance ORTBContentCategory = "IAB13-6"
	// ORTBContentCatInvesting is a content category for Investing
	ORTBContentCatInvesting ORTBContentCategory = "IAB13-7"
	// ORTBContentCatMutualFunds is a content category for Mutual Funds
	ORTBContentCatMutualFunds ORTBContentCategory = "IAB13-8"
	// ORTBContentCatOptions is a content category for Options
	ORTBContentCatOptions ORTBContentCategory = "IAB13-9"
	// ORTBContentCatRetirementPlanning is a content category for Retirement Planning
	ORTBContentCatRetirementPlanning ORTBContentCategory = "IAB13-10"
	// ORTBContentCatStocks is a content category for Stocks
	ORTBContentCatStocks ORTBContentCategory = "IAB13-11"
	// ORTBContentCatTaxPlanning is a content category for Tax Planning
	ORTBContentCatTaxPlanning ORTBContentCategory = "IAB13-12"
	// ORTBContentCatSociety is a content category for Society
	ORTBContentCatSociety ORTBContentCategory = "IAB14"
	// ORTBContentCatDating is a content category for Dating
	ORTBContentCatDating ORTBContentCategory = "IAB14-1"
	// ORTBContentCatDivorceSupport is a content category for Divorce Support
	ORTBContentCatDivorceSupport ORTBContentCategory = "IAB14-2"
	// ORTBContentCatGayLife is a content category for Gay Life
	ORTBContentCatGayLife ORTBContentCategory = "IAB14-3"
	// ORTBContentCatMarriage is a content category for Marriage
	ORTBContentCatMarriage ORTBContentCategory = "IAB14-4"
	// ORTBContentCatSeniorLiving is a content category for Senior Living
	ORTBContentCatSeniorLiving ORTBContentCategory = "IAB14-5"
	// ORTBContentCatTeens is a content category for Teens
	ORTBContentCatTeens ORTBContentCategory = "IAB14-6"
	// ORTBContentCatWeddings is a content category for Weddings
	ORTBContentCatWeddings ORTBContentCategory = "IAB14-7"
	// ORTBContentCatEthnicSpecific is a content category for Ethnic Specific
	ORTBContentCatEthnicSpecific ORTBContentCategory = "IAB14-8"
	// ORTBContentCatScience is a content category for Science
	ORTBContentCatScience ORTBContentCategory = "IAB15"
	// ORTBContentCatAstrology is a content category for Astrology
	ORTBContentCatAstrology ORTBContentCategory = "IAB15-1"
	// ORTBContentCatBiology is a content category for Biology
	ORTBContentCatBiology ORTBContentCategory = "IAB15-2"
	// ORTBContentCatChemistry is a content category for Chemistry
	ORTBContentCatChemistry ORTBContentCategory = "IAB15-3"
	// ORTBContentCatGeology is a content category for Geology
	ORTBContentCatGeology ORTBContentCategory = "IAB15-4"
	// ORTBContentCatParanormalPhenomena is a content category for Paranormal Phenomena
	ORTBContentCatParanormalPhenomena ORTBContentCategory = "IAB15-5"
	// ORTBContentCatPhysics is a content category for Physics
	ORTBContentCatPhysics ORTBContentCategory = "IAB15-6"
	// ORTBContentCatSpaceAstronomy is a content category for Space/Astronomy
	ORTBContentCatSpaceAstronomy ORTBContentCategory = "IAB15-7"
	// ORTBContentCatGeography is a content category for Geography
	ORTBContentCatGeography ORTBContentCategory = "IAB15-8"
	// ORTBContentCatBotany is a content category for Botany
	ORTBContentCatBotany ORTBContentCategory = "IAB15-9"
	// ORTBContentCatWeather is a content category for Weather
	ORTBContentCatWeather ORTBContentCategory = "IAB15-10"
	// ORTBContentCatPets is a content category for Pets
	ORTBContentCatPets ORTBContentCategory = "IAB16"
	// ORTBContentCatAquariums is a content category for Aquariums
	ORTBContentCatAquariums ORTBContentCategory = "IAB16-1"
	// ORTBContentCatBirds is a content category for Birds
	ORTBContentCatBirds ORTBContentCategory = "IAB16-2"
	// ORTBContentCatCats is a content category for Cats
	ORTBContentCatCats ORTBContentCategory = "IAB16-3"
	// ORTBContentCatDogs is a content category for Dogs
	ORTBContentCatDogs ORTBContentCategory = "IAB16-4"
	// ORTBContentCatLargeAnimals is a content category for Large Animals
	ORTBContentCatLargeAnimals ORTBContentCategory = "IAB16-5"
	// ORTBContentCatReptiles is a content category for Reptiles
	ORTBContentCatReptiles ORTBContentCategory = "IAB16-6"
	// ORTBContentCatVeterinaryMedicine is a content category for Veterinary Medicine
	ORTBContentCatVeterinaryMedicine ORTBContentCategory = "IAB16-7"
	// ORTBContentCatSports is a content category for Sports
	ORTBContentCatSports ORTBContentCategory = "IAB17"
	// ORTBContentCatAutoRacing is a content category for Auto Racing
	ORTBContentCatAutoRacing ORTBContentCategory = "IAB17-1"
	// ORTBContentCatBaseball is a content category for Baseball
	ORTBContentCatBaseball ORTBContentCategory = "IAB17-2"
	// ORTBContentCatBicycling is a content category for Bicycling
	ORTBContentCatBicycling ORTBContentCategory = "IAB17-3"
	// ORTBContentCatBodybuilding is a content category for Bodybuilding
	ORTBContentCatBodybuilding ORTBContentCategory = "IAB17-4"
	// ORTBContentCatBoxing is a content category for Boxing
	ORTBContentCatBoxing ORTBContentCategory = "IAB17-5"
	// ORTBContentCatCanoeingKayaking is a content category for Canoeing/Kayaking
	ORTBContentCatCanoeingKayaking ORTBContentCategory = "IAB17-6"
	// ORTBContentCatCheerleading is a content category for Cheerleading
	ORTBContentCatCheerleading ORTBContentCategory = "IAB17-7"
	// ORTBContentCatClimbing is a content category for Climbing
	ORTBContentCatClimbing ORTBContentCategory = "IAB17-8"
	// ORTBContentCatCricket is a content category for Cricket
	ORTBContentCatCricket ORTBContentCategory = "IAB17-9"
	// ORTBContentCatFigureSkating is a content category for Figure Skating
	ORTBContentCatFigureSkating ORTBContentCategory = "IAB17-10"
	// ORTBContentCatFlyFishing is a content category for Fly Fishing
	ORTBContentCatFlyFishing ORTBContentCategory = "IAB17-11"
	// ORTBContentCatFootball is a content category for Football
	ORTBContentCatFootball ORTBContentCategory = "IAB17-12"
	// ORTBContentCatFreshwaterFishing is a content category for Freshwater Fishing
	ORTBContentCatFreshwaterFishing ORTBContentCategory = "IAB17-13"
	// ORTBContentCatGameAndFish is a content category for Game & Fish
	ORTBContentCatGameAndFish ORTBContentCategory = "IAB17-14"
	// ORTBContentCatGolf is a content category for Golf
	ORTBContentCatGolf ORTBContentCategory = "IAB17-15"
	// ORTBContentCatHorseRacing is a content category for Horse Racing
	ORTBContentCatHorseRacing ORTBContentCategory = "IAB17-16"
	// ORTBContentCatHorses is a content category for Horses
	ORTBContentCatHorses ORTBContentCategory = "IAB17-17"
	// ORTBContentCatHuntingShooting is a content category for Hunting/Shooting
	ORTBContentCatHuntingShooting ORTBContentCategory = "IAB17-18"
	// ORTBContentCatInlineSkating is a content category for Inline Skating
	ORTBContentCatInlineSkating ORTBContentCategory = "IAB17-19"
	// ORTBContentCatMartialArts is a content category for Martial Arts
	ORTBContentCatMartialArts ORTBContentCategory = "IAB17-20"
	// ORTBContentCatMountainBiking is a content category for Mountain Biking
	ORTBContentCatMountainBiking ORTBContentCategory = "IAB17-21"
	// ORTBContentCatNASCARRacing is a content category for NASCAR Racing
	ORTBContentCatNASCARRacing ORTBContentCategory = "IAB17-22"
	// ORTBContentCatOlympics is a content category for Olympics
	ORTBContentCatOlympics ORTBContentCategory = "IAB17-23"
	// ORTBContentCatPaintball is a content category for Paintball
	ORTBContentCatPaintball ORTBContentCategory = "IAB17-24"
	// ORTBContentCatPowerAndMotorcycles is a content category for Power & Motorcycles
	ORTBContentCatPowerAndMotorcycles ORTBContentCategory = "IAB17-25"
	// ORTBContentCatProBasketball is a content category for Pro Basketball
	ORTBContentCatProBasketball ORTBContentCategory = "IAB17-26"
	// ORTBContentCatProIceHockey is a content category for Pro Ice Hockey
	ORTBContentCatProIceHockey ORTBContentCategory = "IAB17-27"
	// ORTBContentCatRodeo is a content category for Rodeo
	ORTBContentCatRodeo ORTBContentCategory = "IAB17-28"
	// ORTBContentCatRugby is a content category for Rugby
	ORTBContentCatRugby ORTBContentCategory = "IAB17-29"
	// ORTBContentCatRunningJogging is a content category for Running/Jogging
	ORTBContentCatRunningJogging ORTBContentCategory = "IAB17-30"
	// ORTBContentCatSailing is a content category for Sailing
	ORTBContentCatSailing ORTBContentCategory = "IAB17-31"
	// ORTBContentCatSaltwaterFishing is a content category for Saltwater Fishing
	ORTBContentCatSaltwaterFishing ORTBContentCategory = "IAB17-32"
	// ORTBContentCatScubaDiving is a content category for Scuba Diving
	ORTBContentCatScubaDiving ORTBContentCategory = "IAB17-33"
	// ORTBContentCatSkateboarding is a content category for Skateboarding
	ORTBContentCatSkateboarding ORTBContentCategory = "IAB17-34"
	// ORTBContentCatSkiing is a content category for Skiing
	ORTBContentCatSkiing ORTBContentCategory = "IAB17-35"
	// ORTBContentCatSnowboarding is a content category for Snowboarding
	ORTBContentCatSnowboarding ORTBContentCategory = "IAB17-36"
	// ORTBContentCatSurfingBodyBoarding is a content category for Surfing/Body-Boarding
	ORTBContentCatSurfingBodyBoarding ORTBContentCategory = "IAB17-37"
	// ORTBContentCatSwimming is a content category for Swimming
	ORTBContentCatSwimming ORTBContentCategory = "IAB17-38"
	// ORTBContentCatTableTennisPingPong is a content category for Table Tennis/Ping-Pong
	ORTBContentCatTableTennisPingPong ORTBContentCategory = "IAB17-39"
	// ORTBContentCatTennis is a content category for Tennis
	ORTBContentCatTennis ORTBContentCategory = "IAB17-40"
	// ORTBContentCatVolleyball is a content category for Volleyball
	ORTBContentCatVolleyball ORTBContentCategory = "IAB17-41"
	// ORTBContentCatWalking is a content category for Walking
	ORTBContentCatWalking ORTBContentCategory = "IAB17-42"
	// ORTBContentCatWaterskiWakeboard is a content category for Waterski/Wakeboard
	ORTBContentCatWaterskiWakeboard ORTBContentCategory = "IAB17-43"
	// ORTBContentCatWorldSoccer is a content category for World Soccer
	ORTBContentCatWorldSoccer ORTBContentCategory = "IAB17-44"
	// ORTBContentCatStyleAndFashion is a content category for Style & Fashion
	ORTBContentCatStyleAndFashion ORTBContentCategory = "IAB18"
	// ORTBContentCatBeauty is a content category for Beauty
	ORTBContentCatBeauty ORTBContentCategory = "IAB18-1"
	// ORTBContentCatBodyArt is a content category for Body Art
	ORTBContentCatBodyArt ORTBContentCategory = "IAB18-2"
	// ORTBContentCatFashion is a content category for Fashion
	ORTBContentCatFashion ORTBContentCategory = "IAB18-3"
	// ORTBContentCatJewelry is a content category for Jewelry
	ORTBContentCatJewelry ORTBContentCategory = "IAB18-4"
	// ORTBContentCatClothing is a content category for Clothing
	ORTBContentCatClothing ORTBContentCategory = "IAB18-5"
	// ORTBContentCatAccessories is a content category for Accessories
	ORTBContentCatAccessories ORTBContentCategory = "IAB18-6"
	// ORTBContentCatTechnologyAndComputing is a content category for Technology & Computing
	ORTBContentCatTechnologyAndComputing ORTBContentCategory = "IAB19"
	// ORTBContentCat3DGraphics is a content category for 3-D Graphics
	ORTBContentCat3DGraphics ORTBContentCategory = "IAB19-1"
	// ORTBContentCatAnimation is a content category for Animation
	ORTBContentCatAnimation ORTBContentCategory = "IAB19-2"
	// ORTBContentCatAntivirusSoftware is a content category for Antivirus Software
	ORTBContentCatAntivirusSoftware ORTBContentCategory = "IAB19-3"
	// ORTBContentCatCC is a content category for C/C++
	ORTBContentCatCC ORTBContentCategory = "IAB19-4"
	// ORTBContentCatCamerasAndCamcorders is a content category for Cameras & Camcorders
	ORTBContentCatCamerasAndCamcorders ORTBContentCategory = "IAB19-5"
	// ORTBContentCatCellPhones is a content category for Cell Phones
	ORTBContentCatCellPhones ORTBContentCategory = "IAB19-6"
	// ORTBContentCatComputerCertification is a content category for Computer Certification
	ORTBContentCatComputerCertification ORTBContentCategory = "IAB19-7"
	// ORTBContentCatComputerNetworking is a content category for Computer Networking
	ORTBContentCatComputerNetworking ORTBContentCategory = "IAB19-8"
	// ORTBContentCatComputerPeripherals is a content category for Computer Peripherals
	ORTBContentCatComputerPeripherals ORTBContentCategory = "IAB19-9"
	// ORTBContentCatComputerReviews is a content category for Computer Reviews
	ORTBContentCatComputerReviews ORTBContentCategory = "IAB19-10"
	// ORTBContentCatDataCenters is a content category for Data Centers
	ORTBContentCatDataCenters ORTBContentCategory = "IAB19-11"
	// ORTBContentCatDatabases is a content category for Databases
	ORTBContentCatDatabases ORTBContentCategory = "IAB19-12"
	// ORTBContentCatDesktopPublishing is a content category for Desktop Publishing
	ORTBContentCatDesktopPublishing ORTBContentCategory = "IAB19-13"
	// ORTBContentCatDesktopVideo is a content category for Desktop Video
	ORTBContentCatDesktopVideo ORTBContentCategory = "IAB19-14"
	// ORTBContentCatEmail is a content category for Email
	ORTBContentCatEmail ORTBContentCategory = "IAB19-15"
	// ORTBContentCatGraphicsSoftware is a content category for Graphics Software
	ORTBContentCatGraphicsSoftware ORTBContentCategory = "IAB19-16"
	// ORTBContentCatHomeVideoDVD is a content category for Home Video/DVD
	ORTBContentCatHomeVideoDVD ORTBContentCategory = "IAB19-17"
	// ORTBContentCatInternetTechnology is a content category for Internet Technology
	ORTBContentCatInternetTechnology ORTBContentCategory = "IAB19-18"
	// ORTBContentCatJava is a content category for Java
	ORTBContentCatJava ORTBContentCategory = "IAB19-19"
	// ORTBContentCatJavaScript is a content category for JavaScript
	ORTBContentCatJavaScript ORTBContentCategory = "IAB19-20"
	// ORTBContentCatMacSupport is a content category for Mac Support
	ORTBContentCatMacSupport ORTBContentCategory = "IAB19-21"
	// ORTBContentCatMP3MIDI is a content category for MP3/MIDI
	ORTBContentCatMP3MIDI ORTBContentCategory = "IAB19-22"
	// ORTBContentCatNetConferencing is a content category for Net Conferencing
	ORTBContentCatNetConferencing ORTBContentCategory = "IAB19-23"
	// ORTBContentCatNetforBeginners is a content category for Net for Beginners
	ORTBContentCatNetforBeginners ORTBContentCategory = "IAB19-24"
	// ORTBContentCatNetworkSecurity is a content category for Network Security
	ORTBContentCatNetworkSecurity ORTBContentCategory = "IAB19-25"
	// ORTBContentCatPalmtopsPDAs is a content category for Palmtops/PDAs
	ORTBContentCatPalmtopsPDAs ORTBContentCategory = "IAB19-26"
	// ORTBContentCatPCSupport is a content category for PC Support
	ORTBContentCatPCSupport ORTBContentCategory = "IAB19-27"
	// ORTBContentCatPortable is a content category for Portable
	ORTBContentCatPortable ORTBContentCategory = "IAB19-28"
	// ORTBContentCatEntertainment is a content category for Entertainment
	ORTBContentCatEntertainment ORTBContentCategory = "IAB19-29"
	// ORTBContentCatSharewareFreeware is a content category for Shareware/Freeware
	ORTBContentCatSharewareFreeware ORTBContentCategory = "IAB19-30"
	// ORTBContentCatUnix is a content category for Unix
	ORTBContentCatUnix ORTBContentCategory = "IAB19-31"
	// ORTBContentCatVisualBasic is a content category for Visual Basic
	ORTBContentCatVisualBasic ORTBContentCategory = "IAB19-32"
	// ORTBContentCatWebClipArt is a content category for Web Clip Art
	ORTBContentCatWebClipArt ORTBContentCategory = "IAB19-33"
	// ORTBContentCatWebDesignHTML is a content category for Web Design/HTML
	ORTBContentCatWebDesignHTML ORTBContentCategory = "IAB19-34"
	// ORTBContentCatWebSearch is a content category for Web Search
	ORTBContentCatWebSearch ORTBContentCategory = "IAB19-35"
	// ORTBContentCatWindows is a content category for Windows
	ORTBContentCatWindows ORTBContentCategory = "IAB19-36"
	// ORTBContentCatTravel is a content category for Travel
	ORTBContentCatTravel ORTBContentCategory = "IAB20"
	// ORTBContentCatAdventureTravel is a content category for Adventure Travel
	ORTBContentCatAdventureTravel ORTBContentCategory = "IAB20-1"
	// ORTBContentCatAfrica is a content category for Africa
	ORTBContentCatAfrica ORTBContentCategory = "IAB20-2"
	// ORTBContentCatAirTravel is a content category for Air Travel
	ORTBContentCatAirTravel ORTBContentCategory = "IAB20-3"
	// ORTBContentCatAustraliaAndNewZealand is a content category for Australia & New Zealand
	ORTBContentCatAustraliaAndNewZealand ORTBContentCategory = "IAB20-4"
	// ORTBContentCatBedAndBreakfasts is a content category for Bed & Breakfasts
	ORTBContentCatBedAndBreakfasts ORTBContentCategory = "IAB20-5"
	// ORTBContentCatBudgetTravel is a content category for Budget Travel
	ORTBContentCatBudgetTravel ORTBContentCategory = "IAB20-6"
	// ORTBContentCatBusinessTravel is a content category for Business Travel
	ORTBContentCatBusinessTravel ORTBContentCategory = "IAB20-7"
	// ORTBContentCatByUSLocale is a content category for By US Locale
	ORTBContentCatByUSLocale ORTBContentCategory = "IAB20-8"
	// ORTBContentCatCamping is a content category for Camping
	ORTBContentCatCamping ORTBContentCategory = "IAB20-9"
	// ORTBContentCatCanada is a content category for Canada
	ORTBContentCatCanada ORTBContentCategory = "IAB20-10"
	// ORTBContentCatCaribbean is a content category for Caribbean
	ORTBContentCatCaribbean ORTBContentCategory = "IAB20-11"
	// ORTBContentCatCruises is a content category for Cruises
	ORTBContentCatCruises ORTBContentCategory = "IAB20-12"
	// ORTBContentCatEasternEurope is a content category for Eastern Europe
	ORTBContentCatEasternEurope ORTBContentCategory = "IAB20-13"
	// ORTBContentCatEurope is a content category for Europe
	ORTBContentCatEurope ORTBContentCategory = "IAB20-14"
	// ORTBContentCatFrance is a content category for France
	ORTBContentCatFrance ORTBContentCategory = "IAB20-15"
	// ORTBContentCatGreece is a content category for Greece
	ORTBContentCatGreece ORTBContentCategory = "IAB20-16"
	// ORTBContentCatHoneymoonsGetaways is a content category for Honeymoons/Getaways
	ORTBContentCatHoneymoonsGetaways ORTBContentCategory = "IAB20-17"
	// ORTBContentCatHotels is a content category for Hotels
	ORTBContentCatHotels ORTBContentCategory = "IAB20-18"
	// ORTBContentCatItaly is a content category for Italy
	ORTBContentCatItaly ORTBContentCategory = "IAB20-19"
	// ORTBContentCatJapan is a content category for Japan
	ORTBContentCatJapan ORTBContentCategory = "IAB20-20"
	// ORTBContentCatMexicoAndCentralAmerica is a content category for Mexico & Central America
	ORTBContentCatMexicoAndCentralAmerica ORTBContentCategory = "IAB20-21"
	// ORTBContentCatNationalParks is a content category for National Parks
	ORTBContentCatNationalParks ORTBContentCategory = "IAB20-22"
	// ORTBContentCatSouthAmerica is a content category for South America
	ORTBContentCatSouthAmerica ORTBContentCategory = "IAB20-23"
	// ORTBContentCatSpas is a content category for Spas
	ORTBContentCatSpas ORTBContentCategory = "IAB20-24"
	// ORTBContentCatThemeParks is a content category for Theme Parks
	ORTBContentCatThemeParks ORTBContentCategory = "IAB20-25"
	// ORTBContentCatTravelingwithKids is a content category for Traveling with Kids
	ORTBContentCatTravelingwithKids ORTBContentCategory = "IAB20-26"
	// ORTBContentCatUnitedKingdom is a content category for United Kingdom
	ORTBContentCatUnitedKingdom ORTBContentCategory = "IAB20-27"
	// ORTBContentCatRealEstate is a content category for Real Estate
	ORTBContentCatRealEstate ORTBContentCategory = "IAB21"
	// ORTBContentCatApartments is a content category for Apartments
	ORTBContentCatApartments ORTBContentCategory = "IAB21-1"
	// ORTBContentCatArchitects is a content category for Architects
	ORTBContentCatArchitects ORTBContentCategory = "IAB21-2"
	// ORTBContentCatBuyingSellingHomes is a content category for Buying/Selling Homes
	ORTBContentCatBuyingSellingHomes ORTBContentCategory = "IAB21-3"
	// ORTBContentCatShopping is a content category for Shopping
	ORTBContentCatShopping ORTBContentCategory = "IAB22"
	// ORTBContentCatContestsAndFreebies is a content category for Contests & Freebies
	ORTBContentCatContestsAndFreebies ORTBContentCategory = "IAB22-1"
	// ORTBContentCatCouponing is a content category for Couponing
	ORTBContentCatCouponing ORTBContentCategory = "IAB22-2"
	// ORTBContentCatComparison is a content category for Comparison
	ORTBContentCatComparison ORTBContentCategory = "IAB22-3"
	// ORTBContentCatEngines is a content category for Engines
	ORTBContentCatEngines ORTBContentCategory = "IAB22-4"
	// ORTBContentCatReligionAndSpirituality is a content category for Religion & Spirituality
	ORTBContentCatReligionAndSpirituality ORTBContentCategory = "IAB23"
	// ORTBContentCatAlternativeReligions is a content category for Alternative Religions
	ORTBContentCatAlternativeReligions ORTBContentCategory = "IAB23-1"
	// ORTBContentCatAtheismAgnosticism is a content category for Atheism/Agnosticism
	ORTBContentCatAtheismAgnosticism ORTBContentCategory = "IAB23-2"
	// ORTBContentCatBuddhism is a content category for Buddhism
	ORTBContentCatBuddhism ORTBContentCategory = "IAB23-3"
	// ORTBContentCatCatholicism is a content category for Catholicism
	ORTBContentCatCatholicism ORTBContentCategory = "IAB23-4"
	// ORTBContentCatChristianity is a content category for Christianity
	ORTBContentCatChristianity ORTBContentCategory = "IAB23-5"
	// ORTBContentCatHinduism is a content category for Hinduism
	ORTBContentCatHinduism ORTBContentCategory = "IAB23-6"
	// ORTBContentCatIslam is a content category for Islam
	ORTBContentCatIslam ORTBContentCategory = "IAB23-7"
	// ORTBContentCatJudaism is a content category for Judaism
	ORTBContentCatJudaism ORTBContentCategory = "IAB23-8"
	// ORTBContentCatLatterDaySaints is a content category for Latter-Day Saints
	ORTBContentCatLatterDaySaints ORTBContentCategory = "IAB23-9"
	// ORTBContentCatPaganWiccan is a content category for Pagan/Wiccan
	ORTBContentCatPaganWiccan ORTBContentCategory = "IAB23-10"
	// ORTBContentCatUncategorized is a content category for Uncategorized
	ORTBContentCatUncategorized ORTBContentCategory = "IAB24"
	// ORTBContentCatNonStandardContent is a content category for Non-Standard Content
	ORTBContentCatNonStandardContent ORTBContentCategory = "IAB25"
	// ORTBContentCatUnmoderatedUGC is a content category for Unmoderated UGC
	ORTBContentCatUnmoderatedUGC ORTBContentCategory = "IAB25-1"
	// ORTBContentCatExtremeGraphicExplicitViolence is a content category for Extreme Graphic/Explicit Violence
	ORTBContentCatExtremeGraphicExplicitViolence ORTBContentCategory = "IAB25-2"
	// ORTBContentCatPornography is a content category for Pornography
	ORTBContentCatPornography ORTBContentCategory = "IAB25-3"
	// ORTBContentCatProfaneContent is a content category for Profane Content
	ORTBContentCatProfaneContent ORTBContentCategory = "IAB25-4"
	// ORTBContentCatHateContent is a content category for Hate Content
	ORTBContentCatHateContent ORTBContentCategory = "IAB25-5"
	// ORTBContentCatUnderConstruction is a content category for Under Construction
	ORTBContentCatUnderConstruction ORTBContentCategory = "IAB25-6"
	// ORTBContentCatIncentivized is a content category for Incentivized
	ORTBContentCatIncentivized ORTBContentCategory = "IAB25-7"
	// ORTBContentCatIllegalContents is a content category for Illegal Contents
	ORTBContentCatIllegalContents ORTBContentCategory = "IAB26"
	// ORTBContentCatIllegalContent is a content category for Illegal Content
	ORTBContentCatIllegalContent ORTBContentCategory = "IAB26-1"
	// ORTBContentCatWarez is a content category for Warez
	ORTBContentCatWarez ORTBContentCategory = "IAB26-2"
	// ORTBContentCatSpywareMalware is a content category for Spyware/Malware
	ORTBContentCatSpywareMalware ORTBContentCategory = "IAB26-3"
	// ORTBContentCatCopyrightInfringement is a content category for Copyright Infringement
	ORTBContentCatCopyrightInfringement ORTBContentCategory = "IAB26-4"
)
