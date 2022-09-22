package observer

type JobSite struct {
	subs    []Observer
	vacancy []string
}

func (J *JobSite) addVacancies(vacancy string) {
	J.vacancy = append(J.vacancy, vacancy)
	notifyAll()
}

func (J *JobSite) removeVacancies(vacancy string) {
	for i, v := range J.subs {
		if v == vacancy {
			copy(J.subs[i:], J.subs[i+1:])
			J.subs[len(J.subs)-1] = nil // обнуляем "хвост"
			J.subs = J.subs[:len(J.subs)-1]
		}
	}
	return J.subs
	notifyAll()
}

func (J *JobSite) subscribe(observer Observer) {
	J.subs = append(J.subs, observer)
}

func (J *JobSite) unsubscribe(observer Observer) {

}

func (J *JobSite) notifyAll() {
	for range J.subs {
		Person.HandleEvent(J.vacancy)
	}
	//for (Observer observer: this.subscribers) {
	//	observer.handleEvent(this.vacancies);
	//}
}
