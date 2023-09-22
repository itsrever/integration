package server

const RETURN_ID = "return-1"

func getReturnResponse() map[string]string {
	return map[string]string{
		"return_id": RETURN_ID,
	}
}