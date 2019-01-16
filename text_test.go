package algo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJaccardDistance(t *testing.T) {

	// These are probably different jobs
	s1 := "Excellent job opportunity! need Node.js, MYSQL and resume"
	s2 := "Excellent job opportunity! need ASP.NET, MySQL and good skills in microsoft office"

	assert.False(t, Similar(s1, s2), "These strings should be distinct enough to not be similar")
	assert.True(t, Similar(s1, s1), "These strings should be exactly the same")

	s1 = "1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21"
	s2 = "1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20"

	assert.True(t, Similar(s1, s2), "These strings should be similar")

	assert.True(t, Similar(s1, s1), "These strings should be the same")
}

func TestRemoveNonAlphanumberic(t *testing.T) {
	s1 := "Excellent job opportunity! need ASP.NET, MySQL and good skills in microsoft office"

	// ensure determanistic behavior
	assert.Equal(t, removeNonAlphanumeric(s1), removeNonAlphanumeric(s1))
}

func TestNewWordSetFromText(t *testing.T) {
	s1 := "Excellent job opportunity! need Node.js, MYSQL and resume"
	ws := NewWordSetFromText(s1)

	assert.True(t, ws.Contains("Excellent job opportunity!"))
	assert.True(t, ws.Contains("MYSQL and resume"))
}

func TestSimilarity(t *testing.T) {
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
	phone plan coverageCompany DescriptionWe Met Your Match.

	Strategic Employment Partners (SEP, Inc.) is a premier Technology Recruiting and Placement firm that services
	the greater Southern California, San Francisco, New York, and Boston areas. Our mission is to connect with the 
	area's most innovative and well-respected companies. We provide placement services for Full-time/Direct-hire,
	Contract-to-Hire, and Contract positions.

		We strive to create matches which serve both our clients' business needs and our candidates' career goals.
	Doing so creates long term business relationships which benefit both parties equally. This is why companies 
	come back to us again to fulfill their Software and IT placement needs, and why candidates refer their most
	talented colleagues and friends to us. This is the secret to our success.

		We place Full-time/Direct-hire, Temporary, Part-time, Contract, and Contract-to-hire positions. We 
	represent and place the following types of Software and Information Technology professionals:

	Software/Web Application Architecture and Engineering
	* C#, VB.NET, ASP.NET
	* PHP, Ruby on Rails, Python
	* Java, J2EE
	* VC++, C++, C, and more

	Web Design and Development
	* HTML, JavaScript, CSS
	* User Experience (UX)

	Database Engineering and Administration
	* MS SQL Server, Oracle, MySQL
	* OLTP, Datawarehousing, Business Intelligence
	* Database Reporting

	Mobile Development
	*Objective C, Cocoa, Swift
	*Android, iPad/iPhone SDK, iOS

	Network Engineering and Systems Administration
	* Windows, Linux, Unix, Cisco and more

	Technical Support
	* Network/Hardware, Application/Software

	Technology Management
	* Managers, Directors, VP’s, CTO’s, and CIO’s`

	assert.Equal(t, NewWordSetFromText(s), NewWordSetFromText(s))

	assert.Equal(t, 1.0, JaccardSimilarity(s, s))
}
