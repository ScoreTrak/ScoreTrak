package user

import (
	"context"
	"github.com/ory/kratos-client-go"
)

const CONTEXT_IDENTITY_KEY = "ory_user"

func NewContext(ctx context.Context, i *client.Identity) context.Context {
	return context.WithValue(ctx, CONTEXT_IDENTITY_KEY, i)
}

func FromContext(ctx context.Context) (*client.Identity, bool) {
	i, ok := ctx.Value(CONTEXT_IDENTITY_KEY).(*client.Identity)
	return i, ok
}

func IsAdmin(i *client.Identity) bool {
	if !i.HasMetadataPublic() {
		return false
	}
	if mp, found := i.MetadataPublic.(MetadataPublic); found {
		return mp.IsAdmin
	}
	return false
}

//func GetCompetitions(i *client.Identity) []string {
//	if !i.HasMetadataPublic() {
//		return nil
//	}
//	if comps, found := i.MetadataPublic.(map[string]interface{})["competitions"]; found {
//		log.Printf("%v", comps)
//		return comps.([]string)
//	}
//	return nil
//}

func GetTeams(i *client.Identity) []TeamID {
	if !i.HasMetadataPublic() {
		return nil
	}
	if mp, found := i.MetadataPublic.(MetadataPublic); found {
		return mp.Teams
	}
	return nil
}

//	func IsInCompetition(compId string, i *client.Identity) bool {
//		compIds := GetCompetitions(i)
//		for _, cid := range compIds {
//			if compId == cid {
//				return true
//			}
//		}
//		return false
//	}

func IsInTeam(teamId TeamID, i *client.Identity) bool {
	teamIds := GetTeams(i)
	for _, tid := range teamIds {
		if teamId == tid {
			return true
		}
	}
	return false
}
