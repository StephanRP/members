package Members

import "sort"

type Member struct {
	Id	int	 `json:"id,omitempty"`
	Name 	string 	 `json:"name,omitempty"`
	Lob 	string 	 `json:"lob,omitempty"`
	Pcp 	string 	 `json:"pcp,omitempty"`
	Address	*Address `json:"address,omitempty"`
}

type Address struct {
	City    string   `json:"city,omitempty"`
	State	string	 `json:"state,omitempty"`
}

//move to main make type with Iota as id
//update Itoa only works on constants
var memberArray []Member

func GetSingleMember(i int) Member {
	memberArray = (GetMembers())

	return memberArray[(i - 1)]
}

func DeleteSingleMember(i int) []Member {
	memberArray[i-1] = memberArray[len(memberArray)-1]
	memberArray[len(memberArray)-1] = Member {}
	memberArray = memberArray[:len(memberArray)-1]
	
	return memberArray
}

func GetMembers() []Member {
	sort.Slice(memberArray[:], func(i, j int) bool {
		return memberArray[i].Id < memberArray[j].Id
	})
	return memberArray
}

func SetMember(id int, name string, lob string, pcp string) Member {
	person := Member {Id: id, Name: name, Lob: lob, Pcp: pcp}
	memberArray = append(memberArray, person)
	return person
}
//check if creating a second member array
func DeleteMembers() []Member {
	memberArray = nil
	return memberArray
}

func  SetInitialMembers() {
	members := []Member {
		Member {Id: 1, Name: "Steven Strange", Lob: "Medicare", Pcp: "Dr Strange", Address: &Address {City: "123 Fake St", State: "NY"}},
		Member {Id: 2, Name: "Terry Crews", Lob: "Medicaid", Pcp: "Dr Totally Normal", Address: &Address {City:"321 For Real St"}},
		Member {Id: 3, Name: "John Smith", Lob: "Dual", Pcp: "Dr Who"},
	}

	for _, member := range members {
		memberArray = append(memberArray, member)
	}
}



