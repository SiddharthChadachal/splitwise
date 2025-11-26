package splitwise

import "errors"

func (b *Bill) AddParticipant(p Participant) {
	b.Participants = append(b.Participants, p)
}

func (b *Bill) RemoveParticipant(id string) error {
	for i, p := range b.Participants {
		if p.ID == id {
			b.Participants = append(b.Participants[:i], b.Participants[i+1:]...)
			return nil
		}
	}
	return errors.New("participant not found")
}

func (b *Bill) AddItem(i Item) {
	b.Items = append(b.Items, i)
}

func (b *Bill) RemoveItem(id string) error {
	for i, it := range b.Items {
		if it.ID == id {
			b.Items = append(b.Items[:i], b.Items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}
