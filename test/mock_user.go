package test

import (
	

	"github.com/davidKirshbom/cvSpecificator/models"
)



var MockUser = models.User{
	FirstName: "david",
	LastName:  "kirshbom",
	Title:     "F u l l S t a c k D e v e l o p e r",
	About: `Passionate and self-motivated full-stack developer, I have experience building and maintaining web
applications using various technologies. With a strong foundation in programming fundamentals
and an eagerness to learn new skills, I am committed to delivering high-quality code and
continuously improving my craft.`,
	Experience: []models.Experience{
		{
			StartYear:   "2021",
			EndYear:     "Present Day",
			CompanyName: "Spotower",
			Location:    "Tel-aviv, Israel",
			Role: models.Role{
				Title: "Full Stack Developer",
				Explain: `I spearheaded the migration and optimization of CMS sites to Next.js, achieving a
remarkable enhancement in both performance and user experience.
A linchpin in the team, I took charge of designing and upkeeping our internal data
management system, initially leveraging PHP Laravel and MySQL to guarantee data
integrity and streamlined handling. As the tech landscape evolved, I led the charge in
developing a cutting-edge system using Nest.js and React.
I actively contributed to the architectural design of our AWS infrastructure, fortifying our
cloud environment for optimal reliability. Simultaneously, I took ownership of overseeing
and refining the CI/CD process, orchestrating a seamless development and deployment
pipeline. Throughout, I consistently applied a diverse array of technologies, including
AWS, Git, Selenium, MongoDB, Redis, and CSS, to underpin these initiatives and uphold
the stellar performance of the application.`,
			},
		},
		{
			StartYear:   "2014",
			EndYear:     "2017",
			CompanyName: "IDF ramat-gan",
			Location:    "Israel",
			Role: models.Role{
				Title: "System Design Engineer",
				Explain: `Management of human resource systems.
The job included technical specification, testing, user guidance, and support.
Managing the systems updates, schedules.
Daily basis communication with system developers and software architects, civilian and
military users.`,
			},
		},
	},
	Education: []models.Education{
		{
			StartYear:   "2012",
			EndYear:     "2014",
			SchoolName:  "Ort Shapira",
			Location:    "Kfar saba, Israel",
			DegreeTitle: "Practical engineer",
			Learned: `My academic coursework provided me with a strong foundation in programming,
		including proficiency in C, Assembly, Java, and C#. I also gained expertise in data
		structures, databases, electronics, and operating systems.
		As part of the program, I completed a project using C#.`,
		},
		{
			StartYear:   "2020",
			EndYear:     "2021",
			SchoolName:  "Talpiot-hitech",
			Location:    "Ramat-gan, Israel",
			DegreeTitle: "Full stack developer",
			Learned: `Talpiot-hitech is an intensive practical program that delves into theories and
fundamentals of software development necessary for leading software developers. The
program covers:Procedural Programming, Object-Oriented Programming, Functional
Programming
-Cloud Structure, Microservice Architecture,
-Data Structures, Algorithms, Application Life Cycle
-Coding Standards, Testing, Reviewing Code, and Multi-Threaded Development`,
		},
	},
	Skills:    []string{"PHP", "Laravel", "Node", "ExpressJS", "Nest.js", "Java", "Javascript", "ReactJS", "Next.js", "CSS", "HTML", "SQL", "NoSQL", "AWS", "Kubernetes", "Docker", "Linux", "Git"},
	Languages: []string{"English", "Spanish", "Hebrew"},
	Link: models.Link{
		Label: "profile",
		Url:   "idKirshbom",
	},
	State: models.State{
		FinishEducationAsked:  true,
		MoreEducationAsked:    false,
		FinishForm:            true,
		MoreExperienceAsked:   false,
		FinishExperienceAsked: true,
		StartedForm:           true,
	},
}
var MocklastQuestionIndex = 22
