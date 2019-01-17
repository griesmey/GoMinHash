package minhash

import (
	"github.com/stretchr/testify/assert"
	"math"

	"testing"
)

func TestGenerateCoeffs(t *testing.T) {
	k := 20
	assert.Equal(t, k, len(generateCoeffs()))

	// value is of size k
	coeffs := generateCoeffs()
	assert.Equal(t, k, len(coeffs))

	seen := map[int]bool{}

	// ensure that all coefficients are unique
	for _, coeff := range coeffs {
		_, ok := seen[coeff]

		assert.Equal(t, false, ok)
	}
}

func TestHardExample(t *testing.T) {
	left := "Fitness 19 Daly City is currently looking to expand its team of personal trainers. If you have a passion for fitness and helping others we have an amazing opportunity for you. Our developmental program ensures that personal trainers will be able to do what they want to do which is train clients! Scheduling and hours are flexible we are looking for both full time and part time positions.We promote a fun fast paced environment that allows our Trainers to help others and feel good while doing it. If you have worked as a trainer before or if this is your first time do not worry we have an online service that allows a renewal of expired certifications or a new certification for first time trainers.Requirements:- Passion for health and fitness.- Ability to design and execute workout programs that are safe, fun, and effective.- Ability to learn and grow in a fast paced environment.- Provide outstanding customer service.- Ability to work well in a team oriented atmosphere. If this interests you at all do not hesitate to reach out immediately. We are looking for 3 quality individuals  to on board and get started within the next 30 days. Company DescriptionFitness 19 as a company designs commercial style fitness centers with a local gym feel. We have over 250+ locations nationwide that strive to provide its members with the highest level of customer service in the industry. At Fitness 19 Daly City we offer a wide variety of services that appeal to all ages, demographics, and fitness levels. Our location includes multiple squat racks/ powerlifting stations, large free weight area, group exercise room/ aerobics room, cardio area, and Personal Training space with a wide range of functional tools. Our objective is to have the cleanest facility with the best/ friendliest staff for the most affordable cost. If you want to be apart of an awesome team in a successful health club Fitness 19 Daly City is the right place for you."
	right := "Do you have what it takes to change people’s lives and get them excited about fitness? Then get ready to start your career as a personal trainer! Fitness 19 Daly City has over 5000 members and hundreds are looking to get healthy and look great. So we’re looking for highly motivated individuals to help our members achieve the healthy lifestyle that they need. Our team will help you get started and give you the tools you need to become the best personal trainer you can be!- Beginner and master trainer programs available.- In house certification.- Hands on training from management to ensure success.- Flexible scheduling.- Highest pay in the industry.- Full time and part time positions available. We have 50+ clients ready and waiting, if this opportunity is of any interest to you do not hesitate to reach out immediately! We will be looking to onboard only three trainers and fill their schedules completely within the next 30 days.Company DescriptionFitness 19 as a company designs commercial style fitness centers with a local gym feel. We have over 250+ locations nationwide that strive to provide its members with the highest level of customer service in the industry. At Fitness 19 Daly City we offer a wide variety of services that appeal to all ages, demographics, and fitness levels. Our location includes multiple squat racks/ powerlifting stations, large free weight area, group exercise room/ aerobics room, cardio area, and Personal Training space with a wide range of functional tools. Our objective is to have the cleanest facility with the best/ friendliest staff for the most affordable cost. If you want to be apart of an awesome team in a successful health club Fitness 19 Daly City is the right place for you."

	sim := minHashSimilarity(GenerateMinHash(left), GenerateMinHash(right))
	assert.Equal(t, 0.3, sim)
}

func TestVerySimilar(t *testing.T) {
	right := "45884173 -4139326283 -4192052334 -4256750887 37899554 4896305 -800095331 46919616 23958645 28712869 11860993 -4291556783 -4214671869 -4203078934 -4184367405 -4284516719 -4190364341 -4293235802 11188214 -3937844396"
	left := "42564489 -4139326283 -4192052334 -4256750887 37899554 4896305 -800095331 6018046 23958645 28712869 11860993 -4291556783 -4223396949 -4203078934 -4184367405 -4284516719 -4190364341 -4293235802 23392302 -3937844396"
	l, _ := MinHashFromStr(left)
	r, _ := MinHashFromStr(right)

	similarity := minHashSimilarity(l, r)

	assert.Equal(t, 0.8, similarity)

	left = "7591606 -4260609193 -4284034930 -4239469257 815766 10601203 -3078644487 10741024 29164846 5042403 717987 -4293581964 -4270565387 -4261273528 -4191049522 -4272009785 -4198757552 -4284268361 4790323 -4262487924"
	right = "7591606 -4260609193 -4284034930 -4239469257 815766 10601203 -3078644487 10741024 29164846 5042403 717987 -4293581964 -4270565387 -4261273528 -4191049522 -4272009785 -4198757552 -4284268361 4790323 -4262487924"
	l, _ = MinHashFromStr(left)
	r, _ = MinHashFromStr(right)
	similarity = minHashSimilarity(l, r)

	assert.Equal(t, 1.0, similarity)

}

// Given a training set of documents and ground truth sets
// Compare all ground truth sets to find similarity and ensure that it's
// with our given tolerance of accuracy
func TestMinHashAccuracy(t *testing.T) {
	var groundTruthLeft = []string{"t1088", "t1297", "t1768", "t1952", "t980"}
	var groundTruthRight = []string{"t5015", "t4638", "t5248", "t3495", "t2023"}

	var documents = map[string]string{
		"t1088": `Russian Prime Minister Viktor Chernomyrdin on Thursday proposed a three-phase solution to end the three-month-old conflict in Chechnya, starting with the declaration of demilitarised zones. Derek Jeter hit a go-ahead homer and finished with four hits, Alex Rodriguez added a home run and the New York Yankees beat the slumping Mets 11-8 Saturday. India has added almost 100 million people to its list of the poor, a move that will give a total of 372 million access to state welfare schemes and subsidies, a government official said Monday. Maybe a job applicant claims that he earned a bachelor's degree when he was actually one semester shy of graduation. Or he boasts of winning an award from a trade group that doesn't exist. Police on Wednesday handcuffed and led away three children and seven adults who tried to take water into the hospice where brain-damaged Terri Schiavo is being cared for. He was an important political figure, arrested for engaging in lewd conduct in a public men's Married, with children, he told no one. Instead he pleaded guilty without even hiring a lawyer, hoping the problem would quietly disappear. French judges investigating a scandal involving cash payments for airline tickets moved a step closer to President Jacques Chirac on Wednesday, questioning his daughter in the case that dates back to Chirac's time as Paris mayor. A book based on a cancer patient's diary, which has recorded the emotions of his last days on earth, is being printed and will hit the shelves soon, said Monday's China Daily.`,
		"t1297": `Unexpectedly high producer price and industrial output rises fuelled fears the US economy was overheated, and that the Federal Reserve might raise interest rates, sending blue chips lower Wednesday. U.S. nuclear envoy Christopher Hill says the next round of North Korean nuclear disarmament talks could be held in early July. US President Barack Obama and Vice President Joe Biden will lead mourners on Sunday at a service in West Virginia for the dead of the worst US mining tragedy in decades. Digital video recorders are the rage. The picture is better, and units like the Replay TV _ which uses a hard disk instead of videotape _ have features not available with old-fashioned VCRs. European Union (EU) foreign policy chief Javier Solana said Wednesday that maintaining an arms embargo on China is "unfair" given the changes in China since it was imposed in 1989. "Justice League of America" is exactly the kind of movie Warner Bros. loves to make. Based on the classic DC Comics series, the script is filled with a dream team of recognizable superheroes -- Superman, Batman, Wonder Woman, the Flash -- and could not only become its own franchise, A prosecution expert said Tuesday that a former Ku Klux Klansman is faking mental problems and is competent to stand trial for the 1963 church bombing that killed four black girls. A total of 2.2 billion euros ( 1.870 billion U.S. dollars) of investments departed from Portugal's stock-market from January to August this year, according to a report released by Bank of Portugal Friday.`,
		"t1768": `The peseta nosedived to a new all-time low early Friday afternoon on the London forex market, hitting 93.30 to the German mark, Dresdner Bank analyst Elizabeth Legge said. Your work computer just suffered a major meltdown. Maybe the operating system failed, or a virus crashed the hard drive. News that banking giant Goldman Sachs has been charged with fraud sent Asian stocks tumbling Monday, while airlines were hit as northern European airspace was closed due to the Icelandic volcano. Stating that the ''foundation for economic expansion'' has been laid but that the strength and sustainability of the recovery is still uncertain, Alan Greenspan, the Federal Reserve's chairman, strongly suggested to Congress on Wednesday that monetary policy would remain unchanged for the foreseeable Prime Minister Ariel Sharon has told US officials there is no question of freezing Israel's planned expansion of Maale Adumim, the largest Jewish settlement in the West Bank, an aide said Thursday. er, darlings -- are back where they ought to be, make sure you keep an eye on their training for fall sports. The last thing you want is to have them injured and lounging on the couch where they have spent the past three months hollering for food. The heads of the West Coast chapter of the Hollywood performer unions have submitted a tentative contract settlement for a vote by the guilds' nearly 135,000 members. Overseas direct investment in China during the first 10 months this year increased 37 percent in contractual volume from the same period last year.`,
		"t1952": `Nippon Challenge came from behind at the last mark to defeat Spain's Rioja de Espana by 13 seconds and take the final position in the America's Cup challenger semifinals here Wednesday. Romania's prime minister held talks Tuesday with EU leaders on his country's troubled justice system, amid concerns that the government lacks commitment to further reforms. Next year's Six Nations will start on a Friday for the first time when Wales host England in Cardiff, organisers of Europe's leading international rugby union tournament said Wednesday. U.S. officials were scrambling Thursday to determine whether a 22-year-old captive from the war in Afghanistan is a U.S. citizen. The Ugandan army said Friday it killed seven Lord's Resistance Army (LRA) fighters and rescued 13 children from rebel captivity in clashes with the group in the war-scarred north of the country. Even as President Bush last week named four candidates to fill long-standing vacancies on federal appeals courts, conservative legal activists were spoiling for a fight over what they call the unfair treatment of the president's judicial nominees. Pudgy and shy, 13-year-old Matt Shafer had an outlet in rap but few to really share it with _ until a chance day he saw Bob Ritchie work the deejay's turntable. The Guatemalan government has asked a renewal of the U.N. Verification Mission in Guatemala (MINUGUA), the mission said in a communique issued in Guatemala City on Saturday.`,
		"t980":  `A man was shot dead and fifteen others injured when Zambian policemen clashed with citizens rioting in protest against alleged ritual murders by a local businessman, police said Monday. U.S. President George W. Bush expressed confidence on Monday about passing an immigration bill and said a Senate vote of no-confidence in Alberto Gonzales would have no bearing on his service as attorney general. French President Nicolas Sarkozy announced Tuesday in Washington that he would visit China later this month, joined by his wife, Carla Bruni-Sarkozy. These columns for release Tuesday, April 2, 2002 are moving today to clients of the New York Times News Service. Media and entertainment giant Viacom Inc. said Wednesday it may split into two divisions with one focussing on "growth" and the other, more traditional arm, aiming for "value". I don't know if ''Harry Potter and the Order of the Phoenix'' is a good movie -- I haven't seen it. But I'm pretty certain that it shouldn't be judged as a movie at all. It is a visual representation of a book in which millions of people are Two car bombs blew up Monday in the working class town of Yehud, just hours after three Palestinian militants were killed by missiles fired by an Israeli helicopter in a pinpointed attack. Australia's farmers should remember rising fuel prices were hitting farmers worldwide and not just them, Deputy Prime Minister and National Party Leader John Anderson said Tuesday.`,
		"t5015": `Russian Prime Minister Viktor Chernomyrdin on Thursday proposed a three-phase solution to end the three-month-old conflict in Chechnya, starting with the declaration of demilitarised zones. Derek Jeter hit a go-ahead homer and finished with four hits, Alex Rodriguez added a home run and the New York Yankees beat the slumping Mets 11-8 Saturday. India has added almost 100 million people to its list of the poor, a move that will give a total of 372 million access to state welfare schemes and subsidies, a government official said Monday. Maybe a job applicant claims that he earned a bachelor's degree when he was actually one semester shy of graduation. Or he boasts of winning an award from a trade group that doesn't exist. Police on Wednesday handcuffed and led away three children and seven adults who tried to take water into the hospice where brain-damaged Terri Schiavo is being cared for. He was an important political figure, arrested for engaging in lewd conduct in a public men's room. Married, with children, he told no one. Instead he pleaded guilty without even hiring a lawyer, hoping the problem would quietly disappear. French judges investigating a scandal involving cash payments for airline tickets moved a step closer to President Jacques Chirac on Wednesday, questioning his daughter in the case that dates back to Chirac's time as Paris mayor. A book based on a cancer patient's diary, which has recorded the emotions of his last days on earth, is being printed and will hit the shelves soon, said Monday's China Daily.`,
		"t4638": `Unexpectedly high producer price and industrial output rises fuelled fears the US economy was overheated, and that the Federal Reserve might raise interest rates, sending blue chips lower Wednesday. U.S. nuclear envoy Christopher Hill says the next round of North Korean nuclear disarmament talks could be held in early July. US President Barack Obama and Vice President Joe Biden will lead mourners on Sunday at a memorial service in West Virginia for the dead of the worst US mining tragedy in decades. Digital video recorders are the rage. The picture is better, and units like the Replay TV _ which uses a hard disk instead of videotape _ have features not available with old-fashioned VCRs. European Union (EU) foreign policy chief Javier Solana said Wednesday that maintaining an arms embargo on China is "unfair" given the changes in China since it was imposed in 1989. "Justice League of America" is exactly the kind of movie Warner Bros. loves to make. Based on the classic DC Comics series, the script is filled with a dream team of recognizable superheroes -- Superman, Batman, Wonder Woman, the Flash -- and could not only become its own franchise, A prosecution expert said Tuesday that a former Ku Klux Klansman is faking mental problems and is competent to stand trial for the 1963 church bombing that killed four black girls. A total of 2.2 billion euros ( 1.870 billion U.S. dollars) of investments departed from Portugal's stock-market from January to August this year, according to a report released by Bank of Portugal Friday.`,
		"t5248": `The peseta nosedived to a new all-time low early Friday afternoon on the London forex market, hitting 93.30 to the German mark, Dresdner Bank analyst Elizabeth Legge said. Your work computer just suffered a major meltdown. Maybe the operating system failed, or a virus crashed the hard drive. News that banking giant Goldman Sachs has been charged with fraud sent Asian stocks tumbling Monday, while airlines were hit as northern European airspace was closed due to the Icelandic volcano. Stating that the ''foundation for economic expansion'' has been laid but that the strength and sustainability of the recovery is still uncertain, Alan Greenspan, the Federal Reserve's chairman, strongly suggested to Congress on Wednesday that monetary policy would remain unchanged for the foreseeable future. Prime Minister Ariel Sharon has told US officials there is no question of freezing Israel's planned expansion of Maale Adumim, the largest Jewish settlement in the West Bank, an aide said Thursday. er, darlings -- are back where they ought to be, make sure you keep an eye on their training for fall sports. The last thing you want is to have them injured and lounging on the couch where they have spent the past three months hollering for food. The heads of the West Coast chapter of the Hollywood performer unions have submitted a tentative contract settlement for a vote by the guilds' nearly 135,000 members. Overseas direct investment in China during the first 10 months this year increased 37 percent in contractual volume from the same period last year.`,
		"t3495": `Nippon Challenge came from behind at the last mark to defeat Spain's Rioja de Espana by 13 seconds and take the final position in the America's Cup challenger semifinals here Wednesday. Romania's prime minister held talks Tuesday with EU leaders on his country's troubled justice system, amid concerns that the government lacks commitment to further reforms. Next year's Six Nations will start on a Friday for the first time when Wales host England in Cardiff, organisers of Europe's leading international rugby union tournament said Wednesday. U.S. officials were scrambling Thursday to determine whether a 22-year-old captive from the war in Afghanistan is a U.S. citizen. The Ugandan army said Friday it killed seven Lord's Resistance Army (LRA) fighters and rescued 13 children from rebel captivity in clashes with the group in the war-scarred north of the country. Even as President Bush last week named four to fill long-standing vacancies on federal appeals courts, conservative legal activists were spoiling for a fight over what they call the unfair treatment of the president's judicial nominees. Pudgy and shy, 13-year-old Matt Shafer had an outlet in rap but few to really share it with _ until a chance day he saw Bob Ritchie work the deejay's turntable. The Guatemalan government has asked a renewal of the U.N. Verification Mission in Guatemala (MINUGUA), the mission said in a communique issued in Guatemala City on Saturday.`,
		"t2023": `A man was shot dead and fifteen others injured when Zambian policemen clashed with citizens rioting in protest against alleged ritual murders by a local businessman, police said Monday. U.S. President George W. Bush expressed confidence on Monday about passing an immigration bill and said a Senate vote of no-confidence in Alberto Gonzales would have no bearing on his service as attorney general. French President Nicolas Sarkozy announced Tuesday in Washington that he would visit China later this month, joined by his wife, Carla Bruni-Sarkozy. These columns for release Tuesday, April 2, 2002 are moving today to clients of the New York Times News Service. Media and entertainment giant Viacom Inc. said Wednesday it may split into two divisions with one focussing on "growth" and the other, more traditional arm, aiming for "value". I don't know if ''Harry Potter and the Order of the Phoenix'' is a good movie -- I haven't seen it. But I'm pretty certain that it shouldn't be judged as a movie at all. It is a visual representation of a book in which of people are Two car bombs blew up Monday in the working class town of Yehud, just hours after three Palestinian militants were killed by missiles fired by an Israeli helicopter in a pinpointed attack. Australia's farmers should remember rising fuel prices were hitting farmers worldwide and not just them, Deputy Prime Minister and National Party Leader John Anderson said Tuesday.`,
	}

	// left and right should both be dups of each other
	// Jaccard distance for each pair of ground truth samples is 0.98
	// MinHash signature distance should be [0.95, 1] if using size 20 MinHash Signatures
	for i := range groundTruthLeft {
		similarity := minHashSimilarity(GenerateMinHash(documents[groundTruthLeft[i]]), GenerateMinHash(documents[groundTruthRight[i]]))
		assert.True(t, similarity >= 0.95)
		assert.True(t, similarity <= 1.0)
		assert.Equal(t, true, StringsSimilar(documents[groundTruthLeft[i]], documents[groundTruthRight[i]]))
	}
}

func TestGenerateMinHash(t *testing.T) {
	// ensure that Jaccard and MinHash are similar measures

	// Test exactly similar
	s := `An established legal solutions SaaS company located in Oakland is looking for a
	Senior Software Engineer to join their rapidly growing team. If having great work-life balance,
	a fun team environment, with great growth opportunities is something you are looking for- please apply
	Responsibilities:Responsible for end-to-end development on application developmentResponsible for communicating 
	with outside vendors and offshore development teamsResponsible for influencing process and development workflows
	and best technology practicesRequirements:Experience using C#, ASP.Net MVC, REST, and WCFStrong understanding of 
	object oriented programmingExcellent written and verbal communication skillsNice to Haves:Experience with enterprise
	class application developmentExperience with Azure and Azure SQLExperience architecting and designing applications
	from the ground upCompany Information / Benefits / Perks:Competitive medical/dental/vision plans (most monthly fees
	are covered)401k matching programCompetitive PTO plan with floating holidays and separate sick daysMonthly cell
	phone plan coverageCompany DescriptionWe Met Your Match.`

	// Should be 1.0 (exactly similar)
	jaccardSim := JaccardSimilarity(s, s)
	minHashSim := stringSimilarity(s, s)

	assert.Equal(t, 1.0, minHashSim)
	assert.Equal(t, 1.0, jaccardSim)

	// assert that sEdited and s are 95% similar
	sEdited := s[:len(s)-1] + ","
	assert.Equal(t, 0.95, stringSimilarity(sEdited, s), "These strings should be almost exactly the same")

	small := "Excellent job opportunity! need Node.js, MYSQL and resume"
	smallExtra := "Excellent job opportunity! need Node.js, MYSQL and resume."

	assert.Equal(t, 0.5, stringSimilarity(small, smallExtra), "These should be very close")

	// test slight variations of match
	// min hash should be within 5% of jaccardSimilarity
	// ensure that min hash and regular Jaccard are consistent within 5% error
	s1 := s[0 : len(s)/2]
	jaccardSim = JaccardSimilarity(s, s1)
	minHashSim = stringSimilarity(s, s1)

	diff := math.Abs(jaccardSim - minHashSim)
	assert.True(t, diff < 0.1)
}

func TestDoc2ShingleSet(t *testing.T) {
	s0 := "Excellent job opportunity!"
	shingle := string2Shingle(s0)
	assert.Equal(t, shingle, string2Shingle(s0))
}

func TestCalculateMinHash(t *testing.T) {
	s := "Excellent job opportunity!"
	assert.Equal(t, GenerateMinHash(s), GenerateMinHash(s))
}
