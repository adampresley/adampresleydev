package home

import (
	"html/template"
	"net/http"

	"github.com/adampresley/rendering"
)

func (c HomeController) ExperiencePage(w http.ResponseWriter, r *http.Request) {
	pageName := "pages/experience"

	viewData := Experience{
		BaseViewModel: rendering.BaseViewModel{},
		Jobs: []Job{
			{
				Logo:    "umg.svg",
				LogoAlt: "Universal Music Group logo",
				Title:   "Universal Music Group - Lead Developer",
				Description: `<p>Universal Music Group is the largest music and entertainment company in the world;
					it is home to some of the largest musical artist ever. I work in their Global
					Commerce division, which provides software solutions for eCommerce store management,
					as well as integrations across storefront, order and PO fulfillment, and warehousing.</p>
					<p>I spend my time designing systems in Go and Typscript/Remix.</p>`,
				YearStarted: "2021",
				Highlights: []template.HTML{
					"Reduced processing time of the customer self-cancel process by 75% (95th percentile)",
					"Built a service and process to standardize chart reporting for US, Canada, Australia, Britian, France, and Ireland",
					"Built the foundation service for processing Shopify webhooks with an event architecture to the UMG services stack",
					"Built services to integrate OMS, WMS, and product data for hundreds of UK online stores",
				},
			},
			{
				Logo:    "app-nerds.svg",
				LogoAlt: "App Nerds logo",
				Title:   "Co-Owner and Developer of App Nerds LLC",
				Description: `<p>App Nerds is a company I helped form with a parter to build web sites and 
					applications. We specialized in custom websites in Wix, and custom software applications 
					and integrations. I sold my shares in 2024.</p>
					<p>I designed applications and system using Go, JavaScript, and Wix.</p>`,
				YearStarted: "2019",
				Highlights: []template.HTML{
					`Designed and built a workflow web application for <a href="https://www.in-lingo.com/" _target="_blank">In-Lingo</a> to manage paralegal document translations, with over 200 users processing over 48,000 documents.`,
					`Built a service responsible for handling vendor and payment management for Title & Title, completely automating a previously manual process`,
				},
			},
			{
				Logo:    "symplr.svg",
				LogoAlt: "Symplr Healthcare logo",
				Title:   "API Healthcare/GE Healthcare/Symplr Healthcare - Senior Engineer/Software Architect",
				Description: `<p>API Healthcare purchased Clearview Staffing Software, and I came back to work here 
					after the acquisition. API Healthcare, and later Symplr, provided solutions for staffing 
					agencies, direct hospital staffing, scheduling management, self scheduling, and vendor credentialing.</p>
					<p>We built systems using Typescript, Angular, and Java.</p>`,
				YearStarted: "2010",
				Highlights: []template.HTML{
					`Onboarded and trained a team of Java developers on JavaScript and Typescript fundamentals for an upcoming project`,
					`Spearheaded an effort to redesign an outdated Client Access portal application using newer, component-based client and server code`,
				},
			},
			{
				LogoAlt:  "Equator Financial Solutions logo",
				Initials: "EF",
				Title:    "Equator Financial Solutions - Senior Software Engineer",
				Description: `<p>Equator Financial Solutions built custom workflow web applications for financial institutions, 
					primarily banks, to better facilitate and track the short sale and foreclosure process. Here I 
					wrote ColdFusion web applications. I also served on the architecture team to help improve coding standards.</p>
				`,
				YearStarted: "2010",
				Highlights:  []template.HTML{},
			},
			{
				LogoAlt:  "eSports Partners, Inc. logo",
				Initials: "EP",
				Title:    "eSports Partners, Inc. - Senior Software Engineer",
				Description: `<p>eSports Partners provided an eCommerce solution for college and professional sports teams 
					websites to sell merchandise. I was brought in with a new team to build a new version of their eCommerce
					platform in ColdFusion and JavaScript, with an emphasis on improving speed and reliability.
					I also introduced Java and Groovy components for critical backend processing.<p>
				`,
				YearStarted: "2009",
				Highlights:  []template.HTML{},
			},
			{
				LogoAlt:  "Clearview Staffing Software logo",
				Initials: "CS",
				Title:    "Clearview Staffing Software - Software Architect",
				Description: `<p>Clearview Staffing Software built software for staffing agencies who staffed qualified nurses 
					to hospitals. This included temporary, on-demand staff, as well as long term "travel" nursing. ColdFusion 
					and Microsoft SQL was the primary tech stack. Over time, I introduced adding Java components to portions
					of the stack. I also built internal tooling for tasks like deployment and customer management using
					C#.</p>
				`,
				YearStarted: "2005",
				Highlights:  []template.HTML{},
			},
			{
				LogoAlt:  "IT Storm, Inc. logo",
				Initials: "IT",
				Title:    "IT Storm, Inc. - Software Engineer",
				Description: `<p>IT Storm was a small business formed by myself and two business partners. We built custom 
					web sites and ecommerce solutions in PHP and ColdFusion.</p>
				`,
				YearStarted: "2003",
				Highlights:  []template.HTML{},
			},
			{
				LogoAlt:  "Scan-Direct, Inc. logo",
				Initials: "SD",
				Title:    "Scan-Direct, Inc. - Programmer",
				Description: `<p>Scan-Direct was a company building custom document digitization and archival solutions. 
					They offered large-scale document scanning and digitization, along with a custom (early days) OCR 
					solution. I helped build internal tooling in C++ for splitting and saving multi-page TIFF files onto
					company file shares for archival and storage.</p>
				`,
				YearStarted: "2001",
				Highlights:  []template.HTML{},
			},
		},
	}

	c.renderer.Render(pageName, viewData, w)
}
