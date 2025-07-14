package main

import (
	"strconv"

	"github.com/rivo/tview"
)

var options *tview.List

var question *tview.TextView

var questions = map[int]map[string]any{
	1: {
		"qtext": "Which AdTech component is primarily responsible for the real-time decisioning and delivery of ad creatives based on targeting parameters and campaign rules?",
		"options": []string{
			"Content Management System (CMS)",
			"Ad Server",
			"Customer Relationship Management (CRM)",
			"Email Marketing Platform",
		},
		"correct_answer": 1, // Index of "Ad Server"
	},
	2: {
		"qtext": "In programmatic advertising, which entity typically utilizes a Demand-Side Platform (DSP) to automate bid submissions and manage media buys across various ad exchanges?",
		"options": []string{
			"Publishers",
			"Data Management Platforms (DMPs)",
			"Advertisers/Agencies",
			"Supply-Side Platforms (SSPs)",
		},
		"correct_answer": 2, // Index of "Advertisers/Agencies"
	},
	3: {
		"qtext": "What is the core function of a Supply-Side Platform (SSP) in the programmatic ecosystem?",
		"options": []string{
			"To enable advertisers to purchase impressions programmatically",
			"To optimize the yield for publishers by managing and selling their ad inventory",
			"To provide audience segmentation data",
			"To prevent ad fraud by validating ad creatives",
		},
		"correct_answer": 1, // Index of "To optimize the yield for publishers by managing and selling their ad inventory"
	},
	4: {
		"qtext": "Which technical process facilitates the automated buying and selling of individual ad impressions through instantaneous auctions, typically within milliseconds?",
		"options": []string{
			"Direct Deal Negotiation",
			"Programmatic Guaranteed",
			"Real-Time Bidding (RTB)",
			"Private Marketplace (PMP)",
		},
		"correct_answer": 2, // Index of "Real-Time Bidding (RTB)"
	},
	5: {
		"qtext": "A publisher implements a pre-bid auction mechanism that allows multiple demand sources (SSPs, ad exchanges) to bid on inventory concurrently before the ad server is called. What is this mechanism known as?",
		"options": []string{
			"Waterfall Tagging",
			"Header Bidding (or Prebid.js implementation)",
			"Server-to-Server Integration",
			"Direct Ad Serving",
		},
		"correct_answer": 1, // Index of "Header Bidding (or Prebid.js implementation)"
	},
	6: {
		"qtext": "What is the primary role of a Data Management Platform (DMP) in the AdTech stack from a technical perspective?",
		"options": []string{
			"To serve ad impressions directly to users",
			"To collect, segment, and activate first-, second-, and third-party audience data for targeting",
			"To manage and optimize website content for SEO",
			"To process financial transactions between advertisers and publishers",
		},
		"correct_answer": 1, // Index of "To collect, segment, and activate first-, second-, and third-party audience data for targeting"
	},
	7: {
		"qtext": "Which method of user identification, commonly employed for cross-site tracking, relies on small text files stored in a user's web browser?",
		"options": []string{
			"IP Address Matching",
			"Device Fingerprinting",
			"HTTP Cookies",
			"Server-Side User IDs",
		},
		"correct_answer": 2, // Index of "HTTP Cookies"
	},
	8: {
		"qtext": "When an advertiser seeks to understand which specific touchpoints in a user's journey contributed to a conversion (e.g., a purchase), what analytical concept are they applying?",
		"options": []string{
			"Impression Measurement",
			"Click-Through Rate (CTR) Analysis",
			"Multi-Touch Attribution Modeling",
			"Viewability Metrics",
		},
		"correct_answer": 2, // Index of "Multi-Touch Attribution Modeling"
	},
	9: {
		"qtext": "What is \"Ad Fraud\" primarily defined as within the AdTech context?",
		"options": []string{
			"Non-viewable ad impressions",
			"Deliberate attempts to defraud advertisers or publishers, often through fake traffic or misrepresentation",
			"Technical glitches causing ads to load slowly",
			"The accidental misplacement of ad creatives on a webpage",
		},
		"correct_answer": 1,
	},
	10: {
		"qtext": "Which technical solution is often implemented by publishers to maximize their revenue by giving multiple demand partners the opportunity to bid on inventory before the ad server makes the final decision?",
		"options": []string{
			"Programmatic Direct Deals",
			"Open Auction Waterfall",
			"Header Bidding Wrappers",
			"Ad Block Recovery Software",
		},
		"correct_answer": 2,
	},
	11: {
		"qtext": "What critical challenge arises from the widespread collection and use of user data for targeted advertising, necessitating compliance with regulations like GDPR and CCPA?",
		"options": []string{
			"Bid Shading",
			"User Privacy and Data Governance",
			"Latency in Ad Delivery",
			"Impression Fraud",
		},
		"correct_answer": 1,
	},
	12: {
		"qtext": "In the context of ad serving, what does \"creative rendition\" refer to?",
		"options": []string{
			"The process of designing an ad creative",
			"The transformation and display of an ad creative within a user's browser or app",
			"The legal review of ad content",
			"The optimization of ad loading speed",
		},
		"correct_answer": 1,
	},
	13: {
		"qtext": "What is a \"Post-Bid Filter\" in an RTB environment?",
		"options": []string{
			"A mechanism used by DSPs to filter out unwanted impressions *before* bidding",
			"A security measure to prevent unauthorized access to ad servers",
			"A system that filters out bids or impressions *after* the auction based on brand safety or quality criteria",
			"A tool for publishers to pre-screen advertisers",
		},
		"correct_answer": 2,
	},
	14: {
		"qtext": "Which term describes the process by which a publisher's ad server evaluates available inventory against various demand sources (e.g., direct campaigns, programmatic deals) to determine the highest-paying ad to serve?",
		"options": []string{
			"Audience Segmentation",
			"Yield Optimization",
			"Data Onboarding",
			"Creative Asset Management",
		},
		"correct_answer": 1,
	},
	15: {
		"qtext": "What is the primary function of a \"Cookie Syncing\" (or User ID Syncing) mechanism between different AdTech platforms?",
		"options": []string{
			"To synchronize campaign budgets",
			"To enable different platforms to map and recognize the same user across their respective systems for unified targeting",
			"To update ad creatives simultaneously",
			"To ensure consistent ad loading times across various websites",
		},
		"correct_answer": 1,
	},
	16: {
		"qtext": "Which type of ad impression is deemed \"viewable\" according to common industry standards (e.g., IAB/MRC)?",
		"options": []string{
			"Any ad that has loaded on a webpage, regardless of its position",
			"An ad where at least 50% of its pixels are in view for at least one continuous second (for display) or two continuous seconds (for video)",
			"An ad that receives a click from a user",
			"An ad served but not yet fully loaded",
		},
		"correct_answer": 1,
	},
	17: {
		"qtext": "What is \"Server-Side Tagging\" in the context of data collection and AdTech integrations?",
		"options": []string{
			"Placing tracking tags directly in the HTML of a webpage",
			"Sending data from a website to a server-side container, which then distributes it to various vendor endpoints",
			"Using physical tags on server racks for inventory management",
			"A method for developers to debug ad server issues",
		},
		"correct_answer": 1,
	},
	18: {
		"qtext": "What is the primary technical challenge associated with \"cookieless tracking\" and identity resolution in a privacy-centric future?",
		"options": []string{
			"The inability to serve any ads to users",
			"Establishing persistent, privacy-compliant identifiers that work across different browsers and devices",
			"Increased reliance on third-party cookies",
			"The obsolescence of mobile advertising",
		},
		"correct_answer": 1,
	},
	19: {
		"qtext": "Which AdTech entity is responsible for providing tools and interfaces to define, manage, and optimize the delivery of specific ad creatives to target audiences, often including frequency capping and geo-targeting?",
		"options": []string{
			"Ad Server (Campaign Management Interface)",
			"Content Delivery Network (CDN)",
			"Web Analytics Platform",
			"Social Media Management Tool",
		},
		"correct_answer": 0,
	},
	20: {
		"qtext": "What is \"Bid Shading\" in programmatic advertising?",
		"options": []string{
			"A technique used by publishers to obscure their inventory",
			"An optimization strategy where DSPs intelligently reduce winning bids to the lowest possible price to still win the impression",
			"A form of ad fraud involving manipulating bid prices",
			"A method for advertisers to visually assess competitive bids.",
		},
		"correct_answer": 1,
	},
}

var qno = 1

func Quiz() *tview.Flex {
	QuizContainer := tview.NewFlex()

	question = tview.NewTextView()

	qText := strconv.Itoa(qno) + ".) " + questions[qno]["qtext"].(string)

	question.SetText(qText)

	options = tview.NewList()

	options.
		AddItem(questions[qno]["options"].([]string)[0], "", '1', nil).
		AddItem(questions[qno]["options"].([]string)[1], "", '2', nil).
		AddItem(questions[qno]["options"].([]string)[2], "", '3', nil).
		AddItem(questions[qno]["options"].([]string)[3], "", '4', nil).
		AddItem("Quit Quiz", "", 'q', func() { QuitApp() }).
		SetSelectedFunc(func(i int, answer, empty string, s rune) {
			if qno < 20 {
				answered(i)
			} else {
				submitPage := SubmitPage()
				pages.AddAndSwitchToPage("submitpage", submitPage, true)
			}
		})

	QuizContainer.
		AddItem(question, 3, 1, false).
		AddItem(options, 0, 1, true).
		SetBorder(true).SetTitle("GoQuiz")

	return QuizContainer.SetDirection(0)
}

func answered(answer int) {

	answers = append(answers, answer)
	qno++

	qText := strconv.Itoa(qno) + ".) " + questions[qno]["qtext"].(string)
	question.SetText(qText)

	options.SetItemText(0, questions[qno]["options"].([]string)[0], "")
	options.SetItemText(1, questions[qno]["options"].([]string)[1], "")
	options.SetItemText(2, questions[qno]["options"].([]string)[2], "")
	options.SetItemText(3, questions[qno]["options"].([]string)[3], "")

}
