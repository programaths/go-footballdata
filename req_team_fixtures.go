package footballdata

import (
	"fmt"
	"time"
)

type TeamFixturesRequest struct{ request }

// TimeFrame Modifies the request to specify a specific time frame.
func (r TeamFixturesRequest) TimeFrame(timeframe time.Duration) TeamFixturesRequest {
	r.urlValues.Set("timeFrame", durationToTimeFrame(timeframe))
	return r
}

// Season Modifies the request to specify a list of leagues by their code.
func (r TeamFixturesRequest) Season(season uint64) TeamFixturesRequest {
	r.urlValues.Set("season", fmt.Sprintf("%d", season))
	return r
}

// Venue Modifies the request to specify a venue.
func (r TeamFixturesRequest) Venue(venue Venue) TeamFixturesRequest {
	r.urlValues.Set("venue", string(venue))
	return r
}

// Do Executes the request.
func (r TeamFixturesRequest) Do() (s FixturesResponse, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

// FixturesOfTeam Prepares a request to fetch the fixtures of a soccer season.
func (c *client) FixturesOfTeam(id uint64) TeamFixturesRequest {
	return TeamFixturesRequest{c.req("teams/%d/fixtures", id)}
}
