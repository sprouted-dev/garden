package weather

import (
	"fmt"
	"math"
	"time"
)

// VelocityTracker tracks development velocity and compares to enterprise estimates
type VelocityTracker struct {
	CurrentSession SessionVelocity
	Historical     HistoricalVelocity
	Comparisons    []RealityVsEnterprise
}

// SessionVelocity tracks velocity for current session
type SessionVelocity struct {
	StartTime         time.Time
	CommitsThisHour   int
	FeaturesShipped   int
	LinesWritten      int
	FlowStateDuration time.Duration
	CurrentMomentum   MomentumLevel
}

// HistoricalVelocity tracks patterns over time
type HistoricalVelocity struct {
	AverageFeatureTime   time.Duration
	FastestFeature       time.Duration
	IdeasImplemented     int
	IdeasImplementedRate float64 // ideas per hour
	TotalFeatures        int
}

// RealityVsEnterprise tracks the comedy gold
type RealityVsEnterprise struct {
	Feature           string
	YourTime          time.Duration
	EnterpriseEstimate string
	SpeedMultiple     float64
	Timestamp         time.Time
}

// MomentumLevel represents current development momentum
type MomentumLevel int

const (
	MomentumCold MomentumLevel = iota
	MomentumWarming
	MomentumFlowing
	MomentumLightning
	MomentumTornado
)

// EnterpriseExcuses provides random enterprise quotes
var EnterpriseExcuses = []string{
	"We need to form a tiger team to assess the feasibility matrix",
	"Let's circle back after the stakeholder alignment session next quarter",
	"This needs to go through our 17-stage approval process",
	"The architecture committee meets quarterly",
	"We'll need a team of 12 and a $2M budget",
	"First, we need to write a 47-page requirements document",
	"Our roadmap has this penciled in for next year",
	"We should engage McKinsey for a strategic assessment",
	"Legal needs 6 weeks to review the implications",
	"The security review alone takes 6-8 weeks",
}

// PhilosophicalTruths provides inspiration during velocity reports
var PhilosophicalTruths = []string{
	"You don't get paid to code, you get paid to think.",
	"The value was never in the typing, it was in the thinking.",
	"In 3 days of thinking, you built what they plan for 6 months.",
	"Critics focus on how it's typed. Creators focus on what it does.",
	"Your code = what you thought, not what you typed.",
	"The best programmers aren't the fastest typists, they're the clearest thinkers.",
	"AI doesn't write code. It types what you think.",
	"They say you're cheating. Your users say thank you.",
	"Enterprise: 12 devs typing. You: 1 person thinking. You win.",
	"The future isn't about writing code, it's about having ideas worth coding.",
}

// CalculateSpeedMultiple compares your speed to enterprise
func CalculateSpeedMultiple(yourTime time.Duration, enterpriseDays int) float64 {
	enterpriseHours := float64(enterpriseDays * 8) // 8 hour work days
	yourHours := yourTime.Hours()
	if yourHours == 0 {
		return math.Inf(1) // Infinity - they're still planning
	}
	return enterpriseHours / yourHours
}

// GetMomentumDescription returns a description of current momentum
func (m MomentumLevel) GetMomentumDescription() string {
	switch m {
	case MomentumTornado:
		return "🌪️ TORNADO MODE - Reshaping the landscape!"
	case MomentumLightning:
		return "⚡ LIGHTNING SPEED - Shipping faster than thought!"
	case MomentumFlowing:
		return "🌊 FLOW STATE - Riding the wave!"
	case MomentumWarming:
		return "☀️ WARMING UP - Building momentum!"
	default:
		return "🌱 PLANTING SEEDS - Preparing to sprint!"
	}
}

// GetVelocityReport generates a velocity report
func (vt *VelocityTracker) GetVelocityReport() string {
	report := fmt.Sprintf(`🏃‍♂️ Creative Velocity Report

📊 Current Session:
   Features Shipped: %d
   Commits This Hour: %d
   Flow State Duration: %s
   
%s

🚀 Reality vs Enterprise Theater:
`, vt.CurrentSession.FeaturesShipped, 
   vt.CurrentSession.CommitsThisHour,
   vt.CurrentSession.FlowStateDuration,
   vt.CurrentSession.CurrentMomentum.GetMomentumDescription())

	// Add recent comparisons
	for _, comp := range vt.Comparisons {
		if comp.SpeedMultiple == math.Inf(1) {
			report += fmt.Sprintf("\n📌 %s\n   Your Time: %s\n   Enterprise: %s\n   Speed: ∞ (they're still planning)\n",
				comp.Feature, comp.YourTime, comp.EnterpriseEstimate)
		} else {
			report += fmt.Sprintf("\n📌 %s\n   Your Time: %s\n   Enterprise: %s\n   Speed: %.0fx faster\n",
				comp.Feature, comp.YourTime, comp.EnterpriseEstimate, comp.SpeedMultiple)
		}
	}

	// Add enterprise quote
	quote := EnterpriseExcuses[time.Now().Unix()%int64(len(EnterpriseExcuses))]
	report += fmt.Sprintf("\n😂 Enterprise Quote of the Hour:\n\"%s\"\n- Meanwhile, you shipped it already", quote)

	// Add philosophical truth
	truth := PhilosophicalTruths[time.Now().Unix()%int64(len(PhilosophicalTruths))]
	report += fmt.Sprintf("\n\n💭 Remember:\n\"%s\"", truth)

	return report
}

// TrackFeature records a feature completion for velocity tracking
func (vt *VelocityTracker) TrackFeature(feature string, duration time.Duration, enterpriseDays int) {
	comparison := RealityVsEnterprise{
		Feature:           feature,
		YourTime:          duration,
		EnterpriseEstimate: fmt.Sprintf("%d days (%d sprints)", enterpriseDays, enterpriseDays/10),
		SpeedMultiple:     CalculateSpeedMultiple(duration, enterpriseDays),
		Timestamp:         time.Now(),
	}
	
	vt.Comparisons = append(vt.Comparisons, comparison)
	vt.CurrentSession.FeaturesShipped++
	
	// Update momentum based on speed
	if comparison.SpeedMultiple > 100 {
		vt.CurrentSession.CurrentMomentum = MomentumTornado
	} else if comparison.SpeedMultiple > 50 {
		vt.CurrentSession.CurrentMomentum = MomentumLightning
	} else if comparison.SpeedMultiple > 10 {
		vt.CurrentSession.CurrentMomentum = MomentumFlowing
	}
}