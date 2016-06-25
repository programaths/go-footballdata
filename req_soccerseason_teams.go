package footballdata

type SoccerSeasonTeamsRequest struct{ request }

// Do Executes the request.
func (r SoccerSeasonTeamsRequest) Do() (s TeamList, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

// TeamsOfSoccerSeason Prepares a new request to fetch the league table of a given soccer season.
func (c *client) TeamsOfSoccerSeason(soccerSeasonId uint64) SoccerSeasonTeamsRequest {
	return SoccerSeasonTeamsRequest{c.req("soccerseasons/%d/leagueTable", soccerSeasonId)}
}
