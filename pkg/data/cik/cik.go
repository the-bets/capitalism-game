package cik

// This package will handle managing CIK numbers, which are available via one massive text file here:
//		https://www.sec.gov/Archives/edgar/cik-lookup-data.txt

// These are required to perform data lookups, such as the following example:
//		https://data.sec.gov/api/xbrl/companyconcept/CIK0000320193/us-gaap/AccountsPayableCurrent.json

func getCIKs(url string) {
	// placeholder text for the time being because it's after 01:00 and I need to sleep
	return
}

//TODO: decide how to store this since it's massive... needs to be easily searchable
// error handling as well
// also consider transformation, possibly another function
// in theory this could replace the read and write packages since there won't be much of that to do
