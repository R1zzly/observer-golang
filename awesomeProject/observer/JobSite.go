package observer

type JobSite struct {
	subs    []Observer
	vacancy []string
}

func (J *JobSite) AddVacancies(vacancy string) {
	J.vacancy = append(J.vacancy, vacancy)
	J.notifyAll()
}

func (J *JobSite) RemoveVacancies(vacancy string) {
	J.vacancy = remove(J.vacancy, vacancy)
	J.notifyAll()
}

func (J *JobSite) Subscribe(observer Observer) {
	J.subs = append(J.subs, observer)
}

func (J *JobSite) Unsubscribe(observer Observer) {
	for ind := range J.subs {
		if J.subs[ind] == observer {
			J.subs = removeObserver(J.subs, ind)
			break
		}
	}
}

func (J *JobSite) notifyAll() {
	for ind := range J.subs {
		J.subs[ind].HandleEvent(J.vacancy)
	}
}

func remove(arr []string, element string) []string {
	for i, v := range arr {
		if v == element {
			copy(arr[i:], arr[i+1:])
			arr[len(arr)-1] = ""
			arr = arr[:len(arr)-1]
		}
	}
	return arr
}

func removeObserver(arr []Observer, i int) []Observer {

	copy(arr[i:], arr[i+1:])
	arr[len(arr)-1] = nil
	arr = arr[:len(arr)-1]

	return arr
}

func NewJobSite() *JobSite {
	jbs := new(JobSite)
	jbs.subs = make([]Observer, 0, 5)
	jbs.vacancy = make([]string, 0, 5)
	return jbs
}
